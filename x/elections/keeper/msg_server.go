package keeper

import (
	"context"
	"dent/x/elections/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// AdminVoterRegisteration defines a method to create request for registering new voter.
func (m msgServer) AdminVoterRegisteration(goCtx context.Context, msg *types.MsgAdminNewVoterRegistrationRequest) (*types.MsgAdminNewVoterRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.CreateNewAdminVoterRegisterationRequest(ctx, msg); err != nil {
		return nil, err
	}
	return &types.MsgAdminNewVoterRegistrationResponse{}, nil
}

// AdminVoterRegisteration defines a method to create request for registering new voter.
func (m msgServer) AdminCreateCandidate(goCtx context.Context, msg *types.MsgCreateCandidateRequest) (*types.MsgCreateCandidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.CreateCandidate(ctx, msg); err != nil {
		return nil, err
	}
	return &types.MsgCreateCandidateResponse{}, nil
}

// AdminCreatePoll defines a method to create new poll.
func (m msgServer) AdminCreatePoll(goCtx context.Context, msg *types.MsgCreatePollRequest) (*types.MsgCreatPollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.CreatePoll(ctx, msg); err != nil {
		return nil, err
	}
	return &types.MsgCreatPollResponse{}, nil
}

// VoterRegisteration defines a method for registering new voter.
func (m msgServer) VoterRegisteration(goCtx context.Context, msg *types.MsgNewVoterRegisterationRequest) (*types.MsgNewVoterRegisterationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.CreateNewVoter(ctx, msg); err != nil {
		return nil, err
	}
	return &types.MsgNewVoterRegisterationResponse{}, nil
}

func (m msgServer) Vote(goCtx context.Context, msg *types.MsgVoteRequest) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.Vote(ctx, msg); err != nil {
		return nil, err
	}
	return &types.MsgVoteResponse{}, nil
}
