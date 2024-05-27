/*
liteAPI

Testing SearchApiService

*/

package liteapi

import (
	"context"
	"testing"

	liteapiclient "github.com/liteapi-travel/go-sdk/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_liteapi_SearchApiService(t *testing.T) {

	configuration := liteapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY")
	apiClient := liteapiclient.NewAPIClient(configuration)

	/* t.Run("Test SearchApiService GetMinimumRates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SearchApi.GetMinimumRates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	}) */

	t.Run("Test SearchApiService GetFullRates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SearchApi.GetFullRates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
