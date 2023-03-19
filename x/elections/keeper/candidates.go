package keeper

import (
	"dent/x/elections/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getNexCandidateIDWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastCandidateID(ctx) + 1
	k.SetLastCandidateID(ctx, id)
	return id
}

func (k Keeper) ValidateCreateCandidate(ctx sdk.Context, msg *types.MsgCreateCandidateRequest) error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return fmt.Errorf("invalid address %w", err)
	}

	if msg.AdminAddress != types.AdminAddress {
		return types.ErrUnauthorized
	}

	if msg.Name == "" || msg.Party == "" {
		return types.ErrCandidateNamePartyEmpty
	}

	allCandidates := k.GetAllCandidates(ctx)
	for _, candidate := range allCandidates {
		if msg.Name == candidate.Name && msg.Party == candidate.Party {
			return types.ErrCandidateAlreadyExists
		}
	}
	return nil
}

func (k Keeper) CreateCandidate(ctx sdk.Context, msg *types.MsgCreateCandidateRequest) error {
	if err := k.ValidateCreateCandidate(ctx, msg); err != nil {
		return err
	}
	newCandidateID := k.getNexCandidateIDWithUpdate(ctx)
	newCandidate := types.NewCandidate(newCandidateID, msg.Name, msg.Party, ctx.BlockTime())
	k.SetCandidateByID(ctx, newCandidate)
	return nil
}
