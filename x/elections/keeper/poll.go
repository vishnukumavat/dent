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

func ItemExists(array []uint64, item uint64) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
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

func (k Keeper) UpdateCount(ctx sdk.Context, pollID, candidateID uint64) {
	poll, _ := k.GetPollBytID(ctx, pollID)
	for _, option := range poll.Options {
		if option.CandidateId == candidateID {
			option.CurrentCount = option.CurrentCount + 1
			break
		}
	}
	fmt.Println(poll)
	k.SetPollByID(ctx, poll)

}

func (k Keeper) ValidateVote(ctx sdk.Context, msg *types.MsgVoteRequest) error {
	voterAddr, err := sdk.AccAddressFromBech32(msg.VoterAddress)
	if err != nil {
		return fmt.Errorf("invalid address %w", err)
	}

	if !k.IsValidVoter(ctx, voterAddr.String()) {
		return types.ErrVoterNotRegistered
	}

	poll, found := k.GetPollBytID(ctx, msg.PollId)
	if !found {
		return types.ErrInvalidPollOrCandidateID
	}

	if poll.IsEnded || !poll.IsActive {
		return types.ErrInactivePoll
	}

	found = false
	for _, pollOption := range poll.Options {
		if pollOption.CandidateId == msg.CandidateId {
			found = true
			break
		}
	}
	if !found {
		return types.ErrInvalidOption
	}

	_, found = k.GetCandidateBytID(ctx, msg.CandidateId)
	if !found {
		return types.ErrInvalidPollOrCandidateID
	}
	_, found = k.GetVote(ctx, msg.PollId, voterAddr)
	if found {
		return types.ErrAlreadyVoteAdded
	}
	return nil
}

func (k Keeper) Vote(ctx sdk.Context, msg *types.MsgVoteRequest) error {
	if err := k.ValidateVote(ctx, msg); err != nil {
		return err
	}
	newVote := types.NewVote(msg.PollId, sdk.MustAccAddressFromBech32(msg.VoterAddress))
	k.SetVote(ctx, newVote)
	k.UpdateCount(ctx, msg.PollId, msg.CandidateId)
	return nil
}
