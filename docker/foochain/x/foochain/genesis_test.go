package foochain_test

import (
	"testing"

	keepertest "foochain/testutil/keeper"
	"foochain/testutil/nullify"
	"foochain/x/foochain"
	"foochain/x/foochain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FoochainKeeper(t)
	foochain.InitGenesis(ctx, *k, genesisState)
	got := foochain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
