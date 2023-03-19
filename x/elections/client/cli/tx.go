package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"dent/x/elections/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCreateAdminNewVoterRequestCmd(),
		NewAdminCreateCandidateCmd(),
		NewAdminCreatePollCmd(),
		NewCreateNewVoterRequesCmd(),
	)

	return cmd
}

func NewCreateAdminNewVoterRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin-register-voter [voter_wallet_address] [otp]",
		Args:  cobra.ExactArgs(2),
		Short: "Create a new Request for registering new voter (admin only)",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new Request for registering new Voter.
Example:
$ %s tx %s admin-register-voter dent1qh6350n7ll9myakhy9l2u4wt6vtuwczhtlkhq8 987547 --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			voterWalletAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return fmt.Errorf("invalid address %w", err)
			}

			otp, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("parse otp: %w", err)
			}

			msg := types.NewMsgAdminNewVoterRegistrationRequest(clientCtx.GetFromAddress().String(), voterWalletAddress.String(), otp)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewAdminCreateCandidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin-create-candidate [name] [party]",
		Args:  cobra.ExactArgs(2),
		Short: "Create a new candidate for elections (admin only)",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new candidate for elections.
Example:
$ %s tx %s admin-create-candidate "Narendra Modi" "BJP" --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			party := args[1]

			msg := types.NewMsgAdminCreateCandidate(clientCtx.GetFromAddress().String(), name, party)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewAdminCreatePollCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin-create-poll [poll_name] [candidate_ids] [start_at] [poll_duration]",
		Args:  cobra.ExactArgs(4),
		Short: "Create new poll (admin only)",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create new poll for elections.
Example:
$ %s tx %s admin-create-poll "Lok Sabha" 1,2,3  --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pollName := args[0]

			var candidateIds []uint64
			if args[1] != "" {
				candidateIdsStr := strings.Split(args[1], ",")
				if len(candidateIdsStr) > 1 {
					for _, candidateIDStr := range candidateIdsStr {
						candidateID, err := strconv.ParseUint(candidateIDStr, 10, 64)
						if err != nil {
							return err
						}
						candidateIds = append(candidateIds, candidateID)
					}
				} else {
					return fmt.Errorf("atleast 2 candidates required")
				}
			} else {
				return fmt.Errorf("candidate IDs cannot be empty")
			}

			var startTime time.Time
			timeStr := args[2]
			if err != nil {
				return err
			}
			if timeStr == "" { // empty start time
				return fmt.Errorf("start_at cannot be empty string")
			} else if timeUnix, err := strconv.ParseInt(timeStr, 10, 64); err == nil { // unix time
				startTime = time.Unix(timeUnix, 0)
			} else if timeRFC, err := time.Parse(time.RFC3339, timeStr); err == nil { // RFC time
				startTime = timeRFC
			} else { // invalid input
				return errors.New("invalid start time format")
			}

			pollDuration, err := time.ParseDuration(args[3])
			if err != nil {
				return fmt.Errorf("parse trigger-duration: %w", err)
			}

			msg := types.NewMsgAdminCreatePoll(pollName, candidateIds, startTime, pollDuration, clientCtx.GetFromAddress().String())
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewCreateNewVoterRequesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-voter [otp]",
		Args:  cobra.ExactArgs(1),
		Short: "Register a new voter",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Register a voter.
Example:
$ %s tx %s register-voter 987547 --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			otp, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("parse otp: %w", err)
			}

			msg := types.NewMsgNewVoterRegisterationRequest(clientCtx.GetFromAddress().String(), otp)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
