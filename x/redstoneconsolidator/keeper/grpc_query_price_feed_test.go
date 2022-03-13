package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/kodek16/redstone-consolidator/testutil/keeper"
	"github.com/kodek16/redstone-consolidator/testutil/nullify"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPriceFeedQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RedstoneconsolidatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPriceFeed(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPriceFeedRequest
		response *types.QueryGetPriceFeedResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPriceFeedRequest{
				Name: msgs[0].Name,
			},
			response: &types.QueryGetPriceFeedResponse{PriceFeed: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPriceFeedRequest{
				Name: msgs[1].Name,
			},
			response: &types.QueryGetPriceFeedResponse{PriceFeed: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPriceFeedRequest{
				Name: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PriceFeed(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPriceFeedQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RedstoneconsolidatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPriceFeed(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPriceFeedRequest {
		return &types.QueryAllPriceFeedRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PriceFeedAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PriceFeed), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PriceFeed),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PriceFeedAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PriceFeed), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PriceFeed),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PriceFeedAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PriceFeed),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PriceFeedAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
