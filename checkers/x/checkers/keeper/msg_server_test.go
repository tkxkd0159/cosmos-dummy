package keeper_test

import (
	"context"
	"testing"

	keepertest "checkers/testutil/keeper"
	"checkers/x/checkers/keeper"
	"checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
