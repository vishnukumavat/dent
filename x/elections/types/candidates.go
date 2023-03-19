package types

import (
	time "time"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewCandidate returns a new Candidate object.
func NewCandidate(id uint64, name, party string, createdAt time.Time) Candidate {
	return Candidate{
		Id:        id,
		Name:      name,
		Party:     party,
		CreatedAt: createdAt,
	}
}

// MustMarshalCandidate returns the candidate bytes.
// It throws panic if it fails.
func MustMarshalCandidate(cdc codec.BinaryCodec, candidate Candidate) []byte {
	return cdc.MustMarshal(&candidate)
}

// MustUnmarshalCandidate return the unmarshalled Candidate from bytes.
// It throws panic if it fails.
func MustUnmarshalCandidate(cdc codec.BinaryCodec, value []byte) Candidate {
	candidate, err := UnmarshalCandidate(cdc, value)
	if err != nil {
		panic(err)
	}

	return candidate
}

// UnmarshalCandidate returns the Candidate from bytes.
func UnmarshalCandidate(cdc codec.BinaryCodec, value []byte) (candidate Candidate, err error) {
	err = cdc.Unmarshal(value, &candidate)
	return candidate, err
}
