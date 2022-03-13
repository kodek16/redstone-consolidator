package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/kodek16/redstone-consolidator/testutil/keeper"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/keeper"
	"github.com/kodek16/redstone-consolidator/x/redstoneconsolidator/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RedstoneconsolidatorKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
