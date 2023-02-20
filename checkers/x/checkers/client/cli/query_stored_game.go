package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"checkers/x/checkers/types"
)

func CmdListStoredGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stored-game",
		Short: "list all storedGame",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			if clientCtx.Height == 0 || cmd.Flags().Changed(flags.FlagHeight) {
				height, _ := cmd.Flags().GetInt64(flags.FlagHeight)
				clientCtx = clientCtx.WithHeight(height)
			}
			err := client.SetCmdClientContext(cmd, clientCtx)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStoredGameRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.StoredGameAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowStoredGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stored-game [index]",
		Short: "shows a storedGame",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetStoredGameRequest{
				Index: argIndex,
			}

			res, err := queryClient.StoredGame(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
