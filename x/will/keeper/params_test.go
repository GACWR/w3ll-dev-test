package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "w3ll/testutil/keeper"
	"w3ll/x/will/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.WillKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
