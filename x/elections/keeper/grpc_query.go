package keeper

import (
	"context"
	"dent/x/elections/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the liquidity module.
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.Keeper.paramstore.GetParamSet(ctx, &params)
	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Querier) IsVoterRegistered(c context.Context, req *types.QueryIsVoterRegisteredRequest) (*types.QueryIsVoterRegisteredResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if _, err := sdk.AccAddressFromBech32(req.WalletAddress); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid wallet address")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryIsVoterRegisteredResponse{WalletAddress: req.WalletAddress, IsVoterRegistered: k.IsValidVoter(ctx, req.WalletAddress)}, nil
}

func (k Querier) Candidate(c context.Context, req *types.QueryCandidateRequest) (*types.QueryCandidateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.CandidateId == 0 {
		return nil, status.Error(codes.InvalidArgument, "candidate id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	candidate, found := k.GetCandidateBytID(ctx, req.CandidateId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "candidate with id  %d doesn't exist", req.CandidateId)
	}
	return &types.QueryCandidateResponse{Candidate: &candidate}, nil
}

func (k Querier) Candidates(c context.Context, req *types.QueryCandidatesRequest) (*types.QueryCandidatesResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	var candidatePointers []*types.Candidate

	candidates := k.GetAllCandidates(ctx)
	for i := range candidates {
		candidatePointers = append(candidatePointers, &candidates[i])
	}
	return &types.QueryCandidatesResponse{Candidates: candidatePointers}, nil
}

func (k Querier) Poll(c context.Context, req *types.QueryPollRequest) (*types.QueryPollResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PollId == 0 {
		return nil, status.Error(codes.InvalidArgument, "poll id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	poll, found := k.GetPollBytID(ctx, req.PollId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "poll with id  %d doesn't exist", req.PollId)
	}
	return &types.QueryPollResponse{Poll: &poll}, nil
}

func (k Querier) Polls(c context.Context, req *types.QueryPollsRequest) (*types.QueryPollsResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	var pollPointers []*types.Poll

	polls := k.GetAllPools(ctx)
	for i := range polls {
		if req.IsActive {
			if polls[i].IsActive {
				pollPointers = append(pollPointers, &polls[i])
			}
			continue
		}
		pollPointers = append(pollPointers, &polls[i])
	}
	return &types.QueryPollsResponse{Pools: pollPointers}, nil
}
