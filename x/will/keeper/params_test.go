package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/GACWR/w3ll-dev-test/testutil/keeper"
	"github.com/GACWR/w3ll-dev-test/x/will/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.WillKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
