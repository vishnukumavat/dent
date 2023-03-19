package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgAdminNewVoterRegistrationRequest)(nil)
	_ sdk.Msg = (*MsgCreateCandidateRequest)(nil)
	_ sdk.Msg = (*MsgCreatePollRequest)(nil)
	_ sdk.Msg = (*MsgNewVoterRegisterationRequest)(nil)
)

// Message types for the elections module.
const (
	TypeMsgAdminVoterRegisterationRequest = "admin_voter_registeration_request"
	TypeMsgCreateCandidate                = "admin_create_candidate"
	TypeMsgCreatePoll                     = "admin_create_poll"
	TypeMsgVoterRegisteration             = "voter_registeration"
)

func NewMsgAdminNewVoterRegistrationRequest(
	adminAddress, voterAddress string,
	otp uint64,
) *MsgAdminNewVoterRegistrationRequest {
	return &MsgAdminNewVoterRegistrationRequest{
		AdminAddress:       adminAddress,
		VoterWalletAddress: voterAddress,
		Otp:                otp,
	}
}

func (msg MsgAdminNewVoterRegistrationRequest) Route() string { return RouterKey }

func (msg MsgAdminNewVoterRegistrationRequest) Type() string {
	return TypeMsgAdminVoterRegisterationRequest
}

func (msg MsgAdminNewVoterRegistrationRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address: %v", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.VoterWalletAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid voter address: %v", err)
	}
	if msg.Otp > VoterOTPLimit {
		return fmt.Errorf("invalid OTP length")
	}
	return nil
}

func (msg MsgAdminNewVoterRegistrationRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgAdminNewVoterRegistrationRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgAdminCreateCandidate(
	adminAddress,
	name, party string,
) *MsgCreateCandidateRequest {
	return &MsgCreateCandidateRequest{
		AdminAddress: adminAddress,
		Name:         name,
		Party:        party,
	}
}

func (msg MsgCreateCandidateRequest) Route() string { return RouterKey }

func (msg MsgCreateCandidateRequest) Type() string {
	return TypeMsgAdminVoterRegisterationRequest
}

func (msg MsgCreateCandidateRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address: %v", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrCandidateNamePartyEmpty, "candidate name cannot be empty")
	}
	if msg.Party == "" {
		return sdkerrors.Wrapf(ErrCandidateNamePartyEmpty, "candidate party cannot be empty")
	}
	return nil
}

func (msg MsgCreateCandidateRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateCandidateRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgAdminCreatePoll(
	pollName string,
	candidateIds []uint64,
	startAt time.Time,
	pollDuration time.Duration,
	adminAddress string,
) *MsgCreatePollRequest {
	return &MsgCreatePollRequest{
		PollName:     pollName,
		CandidateIds: candidateIds,
		StartAt:      startAt,
		PollDuration: pollDuration,
		AdminAddress: adminAddress,
	}
}

func (msg MsgCreatePollRequest) Route() string { return RouterKey }

func (msg MsgCreatePollRequest) Type() string {
	return TypeMsgCreatePoll
}

func (msg MsgCreatePollRequest) ValidateBasic() error {
	if msg.PollName == "" {
		return sdkerrors.Wrapf(ErrCandidateNamePartyEmpty, "candidate name cannot be empty")
	}
	if msg.PollDuration < 0 {
		return sdkerrors.Wrapf(ErrInvalidPollDuration, "poll duration should be positive")
	}
	return nil
}

func (msg MsgCreatePollRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreatePollRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgNewVoterRegisterationRequest(
	voterAddress string,
	otp uint64,
) *MsgNewVoterRegisterationRequest {
	return &MsgNewVoterRegisterationRequest{
		VoterWalletAddress: voterAddress,
		Otp:                otp,
	}
}

func (msg MsgNewVoterRegisterationRequest) Route() string { return RouterKey }

func (msg MsgNewVoterRegisterationRequest) Type() string {
	return TypeMsgAdminVoterRegisterationRequest
}

func (msg MsgNewVoterRegisterationRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.VoterWalletAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid voter address: %v", err)
	}
	if msg.Otp > VoterOTPLimit {
		return fmt.Errorf("invalid OTP length")
	}
	return nil
}

func (msg MsgNewVoterRegisterationRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgNewVoterRegisterationRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.VoterWalletAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
