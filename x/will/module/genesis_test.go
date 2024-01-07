package will_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "w3ll/testutil/keeper"
	"w3ll/testutil/nullify"
	"w3ll/x/will/module"
	"w3ll/x/will/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WillKeeper(t)
	will.InitGenesis(ctx, k, genesisState)
	got := will.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
