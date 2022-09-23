package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"jsc/x/checkers/types"
)

var _ = strconv.Itoa(0)

func CmdPlayMove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "play-move [id-value] [from-x] [from-y] [to-x] [to-y]",
		Short: "Broadcast message playMove",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIdValue := args[0]
			argFromX, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argFromY, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argToX, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argToY, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPlayMove(
				clientCtx.GetFromAddress().String(),
				argIdValue,
				argFromX,
				argFromY,
				argToX,
				argToY,
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
