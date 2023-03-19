package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/elections module sentinel errors
var (
	ErrUnauthorized                = sdkerrors.Register(ModuleName, 101, "address unauthorized to add new request")
	ErrVoterAlreadyRegistered      = sdkerrors.Register(ModuleName, 102, "given voter address is already registered")
	ErrOTPLenghtInvalid            = sdkerrors.Register(ModuleName, 103, "only 6 digits OTP is allowed")
	ErrRequestExpiredOrInvalidOTP  = sdkerrors.Register(ModuleName, 104, "invalid OTP or the voter registeration request expired")
	ErrRequestAlreadyExistsWithOTP = sdkerrors.Register(ModuleName, 105, "request already exists with givn OTP")
	ErrCandidateAlreadyExists      = sdkerrors.Register(ModuleName, 106, "candidate already exists with name and party")
	ErrCandidateNamePartyEmpty     = sdkerrors.Register(ModuleName, 107, "candidate name and party cannot be empty")
	ErrInvalidStartTime            = sdkerrors.Register(ModuleName, 108, "invalid start time")
	ErrInvalidPollName             = sdkerrors.Register(ModuleName, 109, "poll name cannot be empty")
	ErrInvalidPollDuration         = sdkerrors.Register(ModuleName, 110, "poll duration should be positive")
)
