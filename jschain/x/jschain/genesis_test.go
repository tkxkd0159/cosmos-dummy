package jschain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "jschain/testutil/keeper"
	"jschain/testutil/nullify"
	"jschain/x/jschain"
	"jschain/x/jschain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JschainKeeper(t)
	jschain.InitGenesis(ctx, *k, genesisState)
	got := jschain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
