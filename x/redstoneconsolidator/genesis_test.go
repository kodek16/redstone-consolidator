package redstoneconsolidator_test

import (
	"testing"

	keepertest "github.com/kodek16/redstone-consolidator/testutil/keeper"
	"github.com/kodek16/redstone-consolidator/testutil/nullify"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RedstoneconsolidatorKeeper(t)
	redstoneconsolidator.InitGenesis(ctx, *k, genesisState)
	got := redstoneconsolidator.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
