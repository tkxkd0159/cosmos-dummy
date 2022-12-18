package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "jschain/testutil/keeper"
	"jschain/x/jschain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JschainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
