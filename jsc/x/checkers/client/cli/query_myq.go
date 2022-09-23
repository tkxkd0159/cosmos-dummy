package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"jsc/x/checkers/types"
)

var _ = strconv.Itoa(0)

func CmdMyq() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "myq [req-1] [req-2]",
		Short: "Query myq",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReq1 := args[0]
			reqReq2 := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryMyqRequest{

				Req1: reqReq1,
				Req2: reqReq2,
			}

			res, err := queryClient.Myq(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
