/*
liteAPI

Testing StaticDataApiService

*/

package liteapi

import (
	"context"
	"testing"

	liteapiclient "github.com/liteapi-travel/go-sdk/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_liteapi_StaticDataApiService(t *testing.T) {

	configuration := liteapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY")
	apiClient := liteapiclient.NewAPIClient(configuration)

	t.Run("Test StaticDataApiService GetCitiesByCountryCode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StaticDataApi.GetCitiesByCountryCode(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StaticDataApiService GetCountries", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StaticDataApi.GetCountries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StaticDataApiService GetCurrencies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StaticDataApi.GetCurrencies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StaticDataApiService GetHotelDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StaticDataApi.GetHotelDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StaticDataApiService GetHotels", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StaticDataApi.GetHotels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StaticDataApiService GetIataCodes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StaticDataApi.GetIataCodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
