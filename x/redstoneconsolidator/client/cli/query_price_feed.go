package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/types"
	"github.com/spf13/cobra"
)

func CmdListPriceFeed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-price-feed",
		Short: "list all priceFeed",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPriceFeedRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PriceFeedAll(context.Background(), params)
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

func CmdShowPriceFeed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-price-feed [name]",
		Short: "shows a priceFeed",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argName := args[0]

			params := &types.QueryGetPriceFeedRequest{
				Name: argName,
			}

			res, err := queryClient.PriceFeed(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
