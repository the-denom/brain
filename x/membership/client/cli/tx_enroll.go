package cli

import (
	"strconv"

	"github.com/cdbo/brain/x/membership/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

const (
	FlagNickname = "nickname"
)

var _ = strconv.Itoa(0)

func CmdEnroll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll",
		Short: "Enroll the caller as a new member of The Denom",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			flagNickname, err := cmd.Flags().GetString(FlagNickname)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnroll(
				clientCtx.GetFromAddress().String(),
				flagNickname,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagNickname, "", "The member's nickname")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
