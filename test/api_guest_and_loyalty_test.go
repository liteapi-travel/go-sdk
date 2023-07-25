/*
liteAPI

Testing GuestAndLoyaltyApiService

*/

package liteapi

import (
	"context"
	"testing"

	liteapiclient "github.com/liteapi-travel/go-sdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_liteapi_GuestAndLoyaltyApiService(t *testing.T) {

	configuration := liteapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY")
	apiClient := liteapiclient.NewAPIClient(configuration)

	t.Run("Test GuestAndLoyaltyApiService GetGuestsIds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GuestAndLoyaltyApi.GetGuestsIds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
