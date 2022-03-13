package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/kodek16/redstone-consolidator/testutil/keeper"
	"github.com/kodek16/redstone-consolidator/testutil/nullify"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/keeper"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPriceFeed(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PriceFeed {
	items := make([]types.PriceFeed, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetPriceFeed(ctx, items[i])
	}
	return items
}

func TestPriceFeedGet(t *testing.T) {
	keeper, ctx := keepertest.RedstoneconsolidatorKeeper(t)
	items := createNPriceFeed(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPriceFeed(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPriceFeedRemove(t *testing.T) {
	keeper, ctx := keepertest.RedstoneconsolidatorKeeper(t)
	items := createNPriceFeed(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePriceFeed(ctx,
			item.Name,
		)
		_, found := keeper.GetPriceFeed(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestPriceFeedGetAll(t *testing.T) {
	keeper, ctx := keepertest.RedstoneconsolidatorKeeper(t)
	items := createNPriceFeed(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPriceFeed(ctx)),
	)
}
