/*
Lite API

Testing HotelsApiService

*/

package openapi

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_HotelsApiService(t *testing.T) {

	apiClient, ctx := apiInit()

	t.Run("Test HotelsApiService GetData", func(t *testing.T) {

		countryCode := "TR"
		limit := 5

		resp, httpRes, err := apiClient.HotelsApi.GetData(ctx).CountryCode(countryCode).Limit(float32(limit)).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HotelsApiService HotelsPost", func(t *testing.T) {

		checkIn := "2023-03-15"
		checkOut := "2023-03-16"
		currency := "USD"
		guestNationality := "MA"
		var adults int32 = 1

		hotelIDs := []string{
			"1000018",
			"26191",
			"248093",
			"57871",
			"268206",
			"28906",
			"497829",
			"436827",
			"1000091",
			"1000876",
			"1001301",
			"1001325",
			"1001464",
			"99249",
			"99122",
			"99121",
			"99119",
		}

		hotelPrebookRequest := HotelsPostRequest{
			Checkin:          &checkIn,
			Checkout:         &checkOut,
			Adults:           &adults,
			Currency:         &currency,
			GuestNationality: &guestNationality,
			HotelIDs:         &hotelIDs,
		}
		resp, httpRes, err := apiClient.HotelsApi.HotelsPost(ctx).HotelsPostRequest(hotelPrebookRequest).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HotelsApiService HotelsHotelIdGet", func(t *testing.T) {

		var hotelID int32 = 57871
		checkIn := "2023-03-15"
		checkOut := "2023-03-16"
		currency := "USD"
		guestNationality := "MA"
		adults := 1

		resp, httpRes, err := apiClient.HotelsApi.HotelsHotelIdGet(ctx, hotelID).
			Checkin(checkIn).Checkout(checkOut).Currency(currency).
			GuestNationality(guestNationality).Adults(int32(adults)).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})
}

func authContext() context.Context {
	keys := map[string]APIKey{}
	keys["ApiKeyAuth"] = APIKey{Key: os.Getenv("API_KEY")}
	ctx := context.WithValue(context.TODO(), ContextAPIKeys, keys)
	return ctx
}

func apiInit() (*APIClient, context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)
	ctx := authContext()
	return apiClient, ctx
}
