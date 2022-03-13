package keeper_test

import (
	"testing"

	testkeeper "github.com/kodek16/redstone-consolidator/testutil/keeper"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RedstoneconsolidatorKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
