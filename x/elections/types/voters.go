package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewVoterRequest returns a new AdminNewVoterRegistrationRequest object.
func NewVoterRequest(requestID, otp uint64, walletAddress string, expireAt time.Time) AdminNewVoterRegistrationRequest {
	return AdminNewVoterRegistrationRequest{
		RequestId:     requestID,
		Otp:           otp,
		WalletAddress: walletAddress,
		ExpireAt:      expireAt,
	}
}

// NewVoter returns a new Voter object.
func NewVoter(voterID uint64, walletAddress string) Voter {
	return Voter{
		VoterId:       voterID,
		WalletAddress: walletAddress,
	}
}

// MustMarshalNewVoterRegistrationRequest returns the AdminNewVoterRegistrationRequest bytes.
// It throws panic if it fails.
func MustMarshalNewVoterRegistrationRequest(cdc codec.BinaryCodec, newVoterRequest AdminNewVoterRegistrationRequest) []byte {
	return cdc.MustMarshal(&newVoterRequest)
}

// MustUnmarshalNewVoterRegistrationRequest return the unmarshalled AdminNewVoterRegistrationRequest from bytes.
// It throws panic if it fails.
func MustUnmarshalNewVoterRegistrationRequest(cdc codec.BinaryCodec, value []byte) AdminNewVoterRegistrationRequest {
	request, err := UnmarshalNewVoterRegistrationRequest(cdc, value)
	if err != nil {
		panic(err)
	}

	return request
}

// UnmarshalNewVoterRegistrationRequest returns the AdminNewVoterRegistrationRequest from bytes.
func UnmarshalNewVoterRegistrationRequest(cdc codec.BinaryCodec, value []byte) (request AdminNewVoterRegistrationRequest, err error) {
	err = cdc.Unmarshal(value, &request)
	return request, err
}

// MustMarshalVoter returns the Voter bytes.
// It throws panic if it fails.
func MustMarshalVoter(cdc codec.BinaryCodec, voter Voter) []byte {
	return cdc.MustMarshal(&voter)
}

// MustUnmarshalVoter return the unmarshalled Voter from bytes.
// It throws panic if it fails.
func MustUnmarshalVoter(cdc codec.BinaryCodec, value []byte) Voter {
	voter, err := UnmarshalVoter(cdc, value)
	if err != nil {
		panic(err)
	}

	return voter
}

// UnmarshalVoter returns the Voter from bytes.
func UnmarshalVoter(cdc codec.BinaryCodec, value []byte) (voter Voter, err error) {
	err = cdc.Unmarshal(value, &voter)
	return voter, err
}
