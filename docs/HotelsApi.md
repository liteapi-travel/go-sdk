# \HotelsApi

All URIs are relative to _http://localhost:8080_

| Method                                                | HTTP request              | Description                                      |
| ----------------------------------------------------- | ------------------------- | ------------------------------------------------ |
| [**GetData**](HotelsApi.md#GetData)                   | **Get** /data             | Search by Destination/Hotel                      |
| [**HotelsHotelIdGet**](HotelsApi.md#HotelsHotelIdGet) | **Get** /hotels/{hotelId} | Get Room Availability &amp; Rates for a Hotel ID |
| [**HotelsPost**](HotelsApi.md#HotelsPost)             | **Post** /hotels          | Get Minimum Price for Available Hotels           |

## GetData

> GetData200Response GetData(ctx).CountryCode(countryCode).HotelName(hotelName).CityName(cityName).Limit(limit).Offset(offset).Latitude(latitude).Longitude(longitude).Distance(distance).Execute()

Search by Destination/Hotel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    countryCode := "countryCode_example" // string | country code Alpha-2 code (example US, RU, CN)
    hotelName := "hotelName_example" // string | hotel name (optional)
    cityName := "cityName_example" // string | city name (optional)
    limit := float32(8.14) // float32 | limit results (max value 1000) (optional)
    offset := float32(8.14) // float32 | results offset (optional)
    latitude := float32(8.14) // float32 | latitude geo coordinates (optional)
    longitude := float32(8.14) // float32 | longtude geo coordinates (optional)
    distance := float32(8.14) // float32 | the distance starting from the selected geopgraphic point (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HotelsApi.GetData(context.Background()).CountryCode(countryCode).HotelName(hotelName).CityName(cityName).Limit(limit).Offset(offset).Latitude(latitude).Longitude(longitude).Distance(distance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HotelsApi.GetData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetData`: GetData200Response
    fmt.Fprintf(os.Stdout, "Response from `HotelsApi.GetData`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRequest struct via the builder pattern

| Name            | Type        | Description                                               | Notes |
| --------------- | ----------- | --------------------------------------------------------- | ----- |
| **countryCode** | **string**  | country code Alpha-2 code (example US, RU, CN)            |
| **hotelName**   | **string**  | hotel name                                                |
| **cityName**    | **string**  | city name                                                 |
| **limit**       | **float32** | limit results (max value 1000)                            |
| **offset**      | **float32** | results offset                                            |
| **latitude**    | **float32** | latitude geo coordinates                                  |
| **longitude**   | **float32** | longtude geo coordinates                                  |
| **distance**    | **float32** | the distance starting from the selected geopgraphic point |

### Return type

[**GetData200Response**](GetData200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## HotelsHotelIdGet

> map[string]any HotelsHotelIdGet(ctx, hotelId).Rid(rid).Xid(xid).Checkin(checkin).Checkout(checkout).Adults(adults).GuestNationality(guestNationality).Currency(currency).SessionId(sessionId).Execute()

Get Room Availability & Rates for a Hotel ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    hotelId := int32(436827) // int32 |
    rid := "888-8888-8888-888" // string |  (optional)
    xid := "xid_example" // string |  (optional)
    checkin := "2023-01-15" // string |  (optional)
    checkout := "2023-01-16" // string |  (optional)
    adults := int32(1) // int32 |  (optional)
    guestNationality := "MA" // string |  (optional)
    currency := "USD" // string |  (optional)
    sessionId := "GIYDEMZNGAYS2MJVPQZDAMRTFUYDCLJRGZ6DC7D4JVAXY7DZM5MDMQKKJJAXKZY" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HotelsApi.HotelsHotelIdGet(context.Background(), hotelId).Rid(rid).Xid(xid).Checkin(checkin).Checkout(checkout).Adults(adults).GuestNationality(guestNationality).Currency(currency).SessionId(sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HotelsApi.HotelsHotelIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HotelsHotelIdGet`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `HotelsApi.HotelsHotelIdGet`: %v\n", resp)
}
```

### Path Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **hotelId** | **int32**           |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiHotelsHotelIdGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**rid** | **string** | |
**xid** | **string** | |
**checkin** | **string** | |
**checkout** | **string** | |
**adults** | **int32** | |
**guestNationality** | **string** | |
**currency** | **string** | |
**sessionId** | **string** | |

### Return type

**map[string]any**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## HotelsPost

> map[string]any HotelsPost(ctx).Rid(rid).HotelsPostRequest(hotelsPostRequest).Execute()

Get Minimum Price for Available Hotels

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    rid := "888-8888-8888-888" // string |  (optional)
    hotelsPostRequest := *openapiclient.NewHotelsPostRequest() // HotelsPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HotelsApi.HotelsPost(context.Background()).Rid(rid).HotelsPostRequest(hotelsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HotelsApi.HotelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HotelsPost`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `HotelsApi.HotelsPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiHotelsPostRequest struct via the builder pattern

| Name                  | Type                                          | Description | Notes |
| --------------------- | --------------------------------------------- | ----------- | ----- |
| **rid**               | **string**                                    |             |
| **hotelsPostRequest** | [**HotelsPostRequest**](HotelsPostRequest.md) |             |

### Return type

**map[string]any**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
