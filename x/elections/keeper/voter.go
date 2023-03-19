package keeper

import (
	"dent/x/elections/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// getNextPairIdWithUpdate increments request id by one and set it.
func (k Keeper) getNextNewVoterRequestIDWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastNewVoterRequestID(ctx) + 1
	k.SetLastNewVoterRequestID(ctx, id)
	return id
}

// getNexVoterIDWithUpdate increments voter id by one and set it.
func (k Keeper) getNexVoterIDWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastVoterID(ctx) + 1
	k.SetLastVoterID(ctx, id)
	return id
}

func (k Keeper) ValidateAdminVoterRegisteration(ctx sdk.Context, msg *types.MsgAdminNewVoterRegistrationRequest) error {

	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return fmt.Errorf("invalid address %w", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.VoterWalletAddress)
	if err != nil {
		return fmt.Errorf("invalid address %w", err)
	}

	if msg.AdminAddress != types.AdminAddress {
		return types.ErrUnauthorized
	}

	if msg.Otp > types.VoterOTPLimit {
		return types.ErrOTPLenghtInvalid

	}
	_, found := k.GetVoterByWalletAddress(ctx, msg.VoterWalletAddress)
	if found {
		return types.ErrVoterAlreadyRegistered
	}
	_, found = k.GetVoterRequestByOTP(ctx, msg.Otp)
	if found {
		return types.ErrRequestAlreadyExistsWithOTP
	}

	return nil
}

func (k Keeper) CreateNewAdminVoterRegisterationRequest(ctx sdk.Context, msg *types.MsgAdminNewVoterRegistrationRequest) error {
	if err := k.ValidateAdminVoterRegisteration(ctx, msg); err != nil {
		return err
	}
	// initialize voter account, so voter can initiate tx for account registeration
	err := k.bankKeeper.SendCoins(ctx, sdk.MustAccAddressFromBech32(msg.AdminAddress), sdk.MustAccAddressFromBech32(msg.VoterWalletAddress), sdk.NewCoins(sdk.NewCoin("udent", sdk.OneInt())))
	if err != nil {
		return err
	}
	newRequestID := k.getNextNewVoterRequestIDWithUpdate(ctx)
	requestExpireAt := ctx.BlockTime().Add(types.NewVoterAdminRequestExpirationDuration)
	newVoterRequest := types.NewVoterRequest(newRequestID, msg.Otp, msg.VoterWalletAddress, requestExpireAt)
	k.SetVoterRequestByOTP(ctx, newVoterRequest)
	return nil
}

func (k Keeper) ValidateVoterRegisteration(ctx sdk.Context, msg *types.MsgNewVoterRegisterationRequest) (error, types.AdminNewVoterRegistrationRequest) {
	_, err := sdk.AccAddressFromBech32(msg.VoterWalletAddress)
	if err != nil {
		return fmt.Errorf("invalid address %w", err), types.AdminNewVoterRegistrationRequest{}
	}

	if msg.Otp > types.VoterOTPLimit {
		return types.ErrOTPLenghtInvalid, types.AdminNewVoterRegistrationRequest{}

	}
	_, found := k.GetVoterByWalletAddress(ctx, msg.VoterWalletAddress)
	if found {
		return types.ErrVoterAlreadyRegistered, types.AdminNewVoterRegistrationRequest{}
	}

	voterRequest, found := k.GetVoterRequestByOTP(ctx, msg.Otp)
	if !found {
		return types.ErrRequestExpiredOrInvalidOTP, types.AdminNewVoterRegistrationRequest{}
	}
	return nil, voterRequest
}

func (k Keeper) CreateNewVoter(ctx sdk.Context, msg *types.MsgNewVoterRegisterationRequest) error {
	err, voterRequest := k.ValidateVoterRegisteration(ctx, msg)
	if err != nil {
		return err
	}
	newVoterID := k.getNexVoterIDWithUpdate(ctx)
	newVoter := types.NewVoter(newVoterID, msg.VoterWalletAddress)
	k.SetVoterByWalletAddress(ctx, newVoter)
	k.DeleteVoterRequest(ctx, voterRequest)
	return nil
}

func (k Keeper) DeleteAllOutdatedNewVoterRequests(ctx sdk.Context) {
	allRequests := k.GetAllNewVoterRequests(ctx)
	for _, request := range allRequests {
		if ctx.BlockTime().After(request.ExpireAt) {
			k.DeleteVoterRequest(ctx, request)
		}
	}
}

func (k Keeper) IsValidVoter(ctx sdk.Context, walletAddress string) bool {
	_, found := k.GetVoterByWalletAddress(ctx, walletAddress)
	if found {
		return true
	}
	return false
}
