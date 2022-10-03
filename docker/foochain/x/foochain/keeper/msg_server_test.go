package keeper_test

import (
	"context"
	"testing"

	keepertest "foochain/testutil/keeper"
	"foochain/x/foochain/keeper"
	"foochain/x/foochain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FoochainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
