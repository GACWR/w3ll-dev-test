package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "w3ll/testutil/keeper"
	"w3ll/x/will/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.WillKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
