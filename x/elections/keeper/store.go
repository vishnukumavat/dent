package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"dent/x/elections/types"
)

// GetLastNewVoterRequestID returns the last request id for registeing new voter.
func (k Keeper) GetLastNewVoterRequestID(ctx sdk.Context) (id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastNewVoterRequestIDKey())
	if bz == nil {
		id = 0 // initialize the new voter request id
	} else {
		var val gogotypes.UInt64Value
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return
}

// SetLastNewVoterRequestID stores the last request id for registeing new voter.
func (k Keeper) SetLastNewVoterRequestID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.GetLastNewVoterRequestIDKey(), bz)
}

// SetVoterRequestByID stores the particular request by otp as the key.
func (k Keeper) SetVoterRequestByOTP(ctx sdk.Context, newVoterRequest types.AdminNewVoterRegistrationRequest) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalNewVoterRegistrationRequest(k.cdc, newVoterRequest)
	store.Set(types.GetNewVoterRequestKey(newVoterRequest.Otp), bz)
}

// GetVoterRequestByID returns new voter request object for the given otp.
func (k Keeper) GetVoterRequestByOTP(ctx sdk.Context, OTP uint64) (newVoterRequest types.AdminNewVoterRegistrationRequest, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetNewVoterRequestKey(OTP))
	if bz == nil {
		return
	}
	newVoterRequest = types.MustUnmarshalNewVoterRegistrationRequest(k.cdc, bz)
	return newVoterRequest, true
}

// DeleteVoterRequest deletes a request obj.
func (k Keeper) DeleteVoterRequest(ctx sdk.Context, req types.AdminNewVoterRegistrationRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetNewVoterRequestKey(req.Otp))
}

// IterateAllVoterRequests iterates over all the stored requests and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllVoterRequests(ctx sdk.Context, cb func(request types.AdminNewVoterRegistrationRequest) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetAllNewVoterRequestKey())
	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)
	for ; iter.Valid(); iter.Next() {
		request := types.MustUnmarshalNewVoterRegistrationRequest(k.cdc, iter.Value())
		stop, err := cb(request)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetRequests returns all requests in the store.
func (k Keeper) GetAllNewVoterRequests(ctx sdk.Context) (requests []types.AdminNewVoterRegistrationRequest) {
	requests = []types.AdminNewVoterRegistrationRequest{}
	_ = k.IterateAllVoterRequests(ctx, func(request types.AdminNewVoterRegistrationRequest) (stop bool, err error) {
		requests = append(requests, request)
		return false, nil
	})
	return requests
}

// GetLastVoterID returns the last voter id for new voter.
func (k Keeper) GetLastVoterID(ctx sdk.Context) (id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastVoterIDKey())
	if bz == nil {
		id = 0 // initialize the new voter request id
	} else {
		var val gogotypes.UInt64Value
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return
}

// SetLastVoterID stores the last voter id for  new voter.
func (k Keeper) SetLastVoterID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.GetLastVoterIDKey(), bz)
}

// SetVoterByWalletAddress stores the particular voter by wallet address as the key.
func (k Keeper) SetVoterByWalletAddress(ctx sdk.Context, voter types.Voter) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalVoter(k.cdc, voter)
	store.Set(types.GetVoterKey(voter.WalletAddress), bz)
}

// GetVoterByWalletAddress returns voter object for the given wallet address.
func (k Keeper) GetVoterByWalletAddress(ctx sdk.Context, walletAddress string) (voter types.Voter, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVoterKey(walletAddress))
	if bz == nil {
		return
	}
	voter = types.MustUnmarshalVoter(k.cdc, bz)
	return voter, true
}

// IterateAllVoters iterates over all the stored voters and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllVoters(ctx sdk.Context, cb func(voter types.Voter) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetAllVoterKey())
	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)
	for ; iter.Valid(); iter.Next() {
		voter := types.MustUnmarshalVoter(k.cdc, iter.Value())
		stop, err := cb(voter)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllVoters returns all voters in the store.
func (k Keeper) GetAllVoters(ctx sdk.Context) (voters []types.Voter) {
	voters = []types.Voter{}
	_ = k.IterateAllVoters(ctx, func(voter types.Voter) (stop bool, err error) {
		voters = append(voters, voter)
		return false, nil
	})
	return voters
}

// GetLastCandidateID returns the last candidate id for new voter.
func (k Keeper) GetLastCandidateID(ctx sdk.Context) (id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastCandidateIDKey())
	if bz == nil {
		id = 0 // initialize the new candidate id
	} else {
		var val gogotypes.UInt64Value
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return
}

// SetLastVoterID stores the last voter id for  new voter.
func (k Keeper) SetLastCandidateID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.GetLastCandidateIDKey(), bz)
}

// SetCandidateByID stores the particular candidate by candidate id as the key.
func (k Keeper) SetCandidateByID(ctx sdk.Context, candidate types.Candidate) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalCandidate(k.cdc, candidate)
	store.Set(types.GetCandidateKey(candidate.Id), bz)
}

// GetCandidateByWalletId returns candidate object for the given id.
func (k Keeper) GetCandidateBytID(ctx sdk.Context, id uint64) (candidate types.Candidate, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetCandidateKey(id))
	if bz == nil {
		return
	}
	candidate = types.MustUnmarshalCandidate(k.cdc, bz)
	return candidate, true
}

// IterateAllCandidates iterates over all the stored candidates and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllCandidates(ctx sdk.Context, cb func(candidate types.Candidate) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetAllCandidatesKey())
	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)
	for ; iter.Valid(); iter.Next() {
		candidate := types.MustUnmarshalCandidate(k.cdc, iter.Value())
		stop, err := cb(candidate)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllCandidates returns all candidates in the store.
func (k Keeper) GetAllCandidates(ctx sdk.Context) (candidates []types.Candidate) {
	candidates = []types.Candidate{}
	_ = k.IterateAllCandidates(ctx, func(candidate types.Candidate) (stop bool, err error) {
		candidates = append(candidates, candidate)
		return false, nil
	})
	return candidates
}

// GetLastPollID returns the last poll id for new voter.
func (k Keeper) GetLastPollID(ctx sdk.Context) (id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastPollIDKey())
	if bz == nil {
		id = 0 // initialize the new poll id
	} else {
		var val gogotypes.UInt64Value
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return
}

// SetLastPollID stores the last poll id for  new voter.
func (k Keeper) SetLastPollID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.GetLastPollIDKey(), bz)
}

// SetPollByID stores the particular poll by poll id as the key.
func (k Keeper) SetPollByID(ctx sdk.Context, poll types.Poll) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalPoll(k.cdc, poll)
	store.Set(types.GetPollKey(poll.Id), bz)
}

// GetPollByID returns poll object for the given id.
func (k Keeper) GetPollBytID(ctx sdk.Context, id uint64) (poll types.Poll, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPollKey(id))
	if bz == nil {
		return
	}
	poll = types.MustUnmarshalPoll(k.cdc, bz)
	return poll, true
}

// IterateAllPools iterates over all the stored polls and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllPools(ctx sdk.Context, cb func(poll types.Poll) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetAllPollKey())
	defer func(iter sdk.Iterator) {
		err := iter.Close()
		if err != nil {
			return
		}
	}(iter)
	for ; iter.Valid(); iter.Next() {
		poll := types.MustUnmarshalPoll(k.cdc, iter.Value())
		stop, err := cb(poll)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}

// GetAllPools returns all polls in the store.
func (k Keeper) GetAllPools(ctx sdk.Context) (pools []types.Poll) {
	pools = []types.Poll{}
	_ = k.IterateAllPools(ctx, func(poll types.Poll) (stop bool, err error) {
		pools = append(pools, poll)
		return false, nil
	})
	return pools
}
