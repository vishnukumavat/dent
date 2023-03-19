package types

import (
	time "time"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewPoll returns a new poll object.
func NewPoll(
	id uint64,
	pollName string,
	candidate_ids []uint64,
	startAt time.Time,
	pollDuration time.Duration,
) Poll {
	options := []*PollOptions{}
	for _, cid := range candidate_ids {
		options = append(options, &PollOptions{
			CandidateId:  cid,
			CurrentCount: 0,
		})
	}
	return Poll{
		Id:           id,
		PollName:     pollName,
		Options:      options,
		IsActive:     false,
		StartAt:      startAt,
		PollDuration: pollDuration,
		IsEnded:      false,
	}
}

// MustMarshalPoll returns the poll bytes.
// It throws panic if it fails.
func MustMarshalPoll(cdc codec.BinaryCodec, poll Poll) []byte {
	return cdc.MustMarshal(&poll)
}

// MustUnmarshalPoll return the unmarshalled Poll from bytes.
// It throws panic if it fails.
func MustUnmarshalPoll(cdc codec.BinaryCodec, value []byte) Poll {
	poll, err := UnmarshalPoll(cdc, value)
	if err != nil {
		panic(err)
	}

	return poll
}

// UnmarshalPoll returns the Poll from bytes.
func UnmarshalPoll(cdc codec.BinaryCodec, value []byte) (poll Poll, err error) {
	err = cdc.Unmarshal(value, &poll)
	return poll, err
}
