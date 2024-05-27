/*
liteAPI

Testing BookingManagementApiService

*/

package liteapi

import (
	"context"
	"testing"

	liteapiclient "github.com/liteapi-travel/go-sdk/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_liteapi_BookingManagementApiService(t *testing.T) {

	configuration := liteapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY")
	apiClient := liteapiclient.NewAPIClient(configuration)

	t.Run("Test BookingManagementApiService RetrievedBooking", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bookingId string

		resp, httpRes, err := apiClient.BookingManagementApi.RetrievedBooking(context.Background(), bookingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookingManagementApiService CancelBooking", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bookingId string

		resp, httpRes, err := apiClient.BookingManagementApi.CancelBooking(context.Background(), bookingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookingManagementApiService GetBookingListByGuestId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BookingManagementApi.GetBookingListByGuestId(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
