package keeper

import (
	"dent/x/elections/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func removeDuplicateInt(intSlice []uint64) []uint64 {
	allKeys := make(map[uint64]bool)
	list := []uint64{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// getNextPairIdWithUpdate increments request id by one and set it.
func (k Keeper) getNextPollIDWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastPollID(ctx) + 1
	k.SetLastPollID(ctx, id)
	return id
}

func (k Keeper) ValidatCreatePoll(ctx sdk.Context, msg *types.MsgCreatePollRequest) error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return fmt.Errorf("invalid address %w", err)
	}

	if msg.AdminAddress != types.AdminAddress {
		return types.ErrUnauthorized
	}

	if ctx.BlockTime().After(msg.StartAt) {
		return types.ErrInvalidStartTime

	}
	if msg.PollName == "" {
		return types.ErrInvalidPollName
	}

	if msg.PollDuration < 0 {
		return types.ErrInvalidPollDuration
	}

	for _, cid := range msg.CandidateIds {
		_, found := k.GetCandidateBytID(ctx, cid)
		if !found {
			return status.Errorf(codes.NotFound, "candidate with id  %d doesn't exist", cid)
		}
	}
	return nil
}

func (k Keeper) CreatePoll(ctx sdk.Context, msg *types.MsgCreatePollRequest) error {
	msg.CandidateIds = removeDuplicateInt(msg.CandidateIds)
	if err := k.ValidatCreatePoll(ctx, msg); err != nil {
		return err
	}
	newPollID := k.getNextPollIDWithUpdate(ctx)
	poll := types.NewPoll(newPollID, msg.PollName, msg.CandidateIds, msg.StartAt, msg.PollDuration)
	k.SetPollByID(ctx, poll)
	return nil
}

func (k Keeper) UpdatePoll(ctx sdk.Context) {
	allPolls := k.GetAllPools(ctx)

	for _, poll := range allPolls {
		if !poll.IsEnded {
			if poll.IsActive {
				if ctx.BlockTime().After(poll.StartAt.Add(poll.PollDuration)) {
					poll.IsActive = false
					poll.IsEnded = true
				}
			} else {
				if ctx.BlockTime().After(poll.StartAt) {
					poll.IsActive = true
				}
			}
			k.SetPollByID(ctx, poll)
		}
	}
}
