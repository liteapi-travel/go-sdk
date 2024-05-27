/*
liteAPI

Testing BookApiService

*/

package liteapi

import (
	"context"
	"testing"

	liteapiclient "github.com/liteapi-travel/go-sdk/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_liteapi_BookApiService(t *testing.T) {

	configuration := liteapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY")
	apiClient := liteapiclient.NewAPIClient(configuration)

	t.Run("Test BookApiService Book", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BookApi.Book(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BookApiService PreBook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BookApi.PreBook(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
