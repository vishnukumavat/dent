package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "elections"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_elections"

	AdminAddress = "dent1kcjmfa2rssd53uvtnsyzcme4njwntyya7xyhsw"

	VoterOTPLimit                          = 999999
	NewVoterAdminRequestExpirationDuration = time.Minute * 10
)

var (
	LastNewVoterRequestIDKey = []byte{0xa0}
	NewVoterRequestKeyPrefix = []byte{0xa1}

	LastVoterIDKey = []byte{0xa2}
	VoterKeyPrefix = []byte{0xa3}

	LastCandidateIDKey = []byte{0xa4}
	CandidateKeyPrefix = []byte{0xa5}

	LastPollIDKey = []byte{0xa6}
	PollKeyPrefix = []byte{0xa7}
)

// GetLastVoterRequestIDKey returns the store key to retrieve the last new voter request id.
func GetLastNewVoterRequestIDKey() []byte {
	return LastNewVoterRequestIDKey
}

// GetNewVoterRequestKey returns the store key to retrieve the AdminNewVoterRegistrationRequest object.
func GetNewVoterRequestKey(OTP uint64) []byte {
	return append(NewVoterRequestKeyPrefix, sdk.Uint64ToBigEndian(OTP)...)
}

// GetNewVoterRequestKey returns the store key to retrieve all requests.
func GetAllNewVoterRequestKey() []byte {
	return NewVoterRequestKeyPrefix
}

// GetVoterIDKey returns the store key to retrieve the last voter id.
func GetLastVoterIDKey() []byte {
	return LastVoterIDKey
}

// GetNewVoterRequestKey returns the store key to retrieve the AdminNewVoterRegistrationRequest object.
func GetVoterKey(walletAddress string) []byte {
	return append(VoterKeyPrefix, LengthPrefixString(walletAddress)...)
}

// GetAllVoterKey returns the store key to retrieve the all voters.
func GetAllVoterKey() []byte {
	return VoterKeyPrefix
}

// GetLastCandidateIDKey returns the store key to retrieve the last candidate id.
func GetLastCandidateIDKey() []byte {
	return LastCandidateIDKey
}

// GetCandidateKey returns the store key to retrieve the Candidate object.
func GetCandidateKey(candidateID uint64) []byte {
	return append(CandidateKeyPrefix, sdk.Uint64ToBigEndian(candidateID)...)
}

// GetAllCandidatesKey returns the store key to retrieve the all candidates.
func GetAllCandidatesKey() []byte {
	return CandidateKeyPrefix
}

// GetLastPollIDKey returns the store key to retrieve the last poll id.
func GetLastPollIDKey() []byte {
	return LastPollIDKey
}

// GetPollKey returns the store key to retrieve the poll object.
func GetPollKey(pollID uint64) []byte {
	return append(PollKeyPrefix, sdk.Uint64ToBigEndian(pollID)...)
}

// GetAllPollKey returns the store key to retrieve the all polls.
func GetAllPollKey() []byte {
	return PollKeyPrefix
}

// LengthPrefixString returns length-prefixed bytes representation
// of a string.
func LengthPrefixString(s string) []byte {
	bz := []byte(s)
	bzLen := len(bz)
	if bzLen == 0 {
		return bz
	}
	return append([]byte{byte(bzLen)}, bz...)
}
