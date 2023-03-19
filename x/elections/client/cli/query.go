package cli

import (
	"context"
	"fmt"
	"strconv"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"dent/x/elections/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group elections queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryIsVoterRegistered(),
		CmdQueryCandidate(),
		CmdQueryCandidates(),
		CmdQueryPoll(),
		CmdQueryPolls(),
	)

	return cmd
}

func CmdQueryIsVoterRegistered() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-voter-registered [wallet_address]",
		Short: "check if the give wallet is registered as voter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			voterWalletAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return fmt.Errorf("invalid address %w", err)
			}

			res, err := queryClient.IsVoterRegistered(context.Background(), &types.QueryIsVoterRegisteredRequest{WalletAddress: voterWalletAddress.String()})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryCandidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "candidate [candidate_id]",
		Short: "query listed candidate by candidate ID ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			candidate_id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("parse candidate id: %w", err)
			}

			res, err := queryClient.Candidate(context.Background(), &types.QueryCandidateRequest{CandidateId: candidate_id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryCandidates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "candidates",
		Short: "query all listed candidate",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Candidates(context.Background(), &types.QueryCandidatesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "poll [poll_id]",
		Short: "query listed polls by poll ID ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			poll_id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("parse poll id: %w", err)
			}

			res, err := queryClient.Poll(context.Background(), &types.QueryPollRequest{PollId: poll_id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPolls() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "polls",
		Short: "query all listed polls",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			isActive, err := cmd.Flags().GetBool(FlagIsActive)
			if err != nil {
				return err
			}
			res, err := queryClient.Polls(context.Background(), &types.QueryPollsRequest{IsActive: isActive})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetQueryPolls())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
