package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdateStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-status [address] [status]",
		Short: "Update a user's membership status",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Update a user's membership status on The Denom.

Example:
  To update a member's status to Inactive
  $ %s tx %s update-status [address] inactive

The new status must be one of the following: %s
`,
				version.AppName, types.ModuleName, types.GetAllShortFormMembershipStatusesAsString())),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]

			// status
			status := types.ParseShortFormMembershipStatus(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateStatus(
				clientCtx.GetFromAddress().String(),
				argAddress,
				status,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
