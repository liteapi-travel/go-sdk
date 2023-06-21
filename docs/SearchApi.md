# \SearchApi

All URIs are relative to *https://api.liteapi.travel/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HotelsGet**](SearchApi.md#HotelsGet) | **Get** /hotels | hotel minimum rates availability
[**HotelsRatesGet**](SearchApi.md#HotelsRatesGet) | **Get** /hotels/rates | hotel full rates availability



## HotelsGet

> map[string]interface{} HotelsGet(ctx).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).Currency(currency).GuestNationality(guestNationality).Adults(adults).Children(children).GuestId(guestId).Execute()

hotel minimum rates availability



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    hotelIds := "lp3803c,lp1f982,lp19b70,lp19e75" // string | List of hotelIds
    checkin := "2023-11-15" // string | Check in data in YYYY-MM-DD format
    checkout := "2023-11-16" // string | Check out data in YYYY-MM-DD format
    currency := "USD" // string | Currency code - example (USD)
    guestNationality := "US" // string | Guest nationality ISO-2 code - example (SG)
    adults := int32(1) // int32 | Number of adult guests staying
    children := "12,9" // string | Number of children staying if any (optional)
    guestId := "testtraveler1" // string | Unique traveler ID if available (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.HotelsGet(context.Background()).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).Currency(currency).GuestNationality(guestNationality).Adults(adults).Children(children).GuestId(guestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.HotelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HotelsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.HotelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHotelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hotelIds** | **string** | List of hotelIds | 
 **checkin** | **string** | Check in data in YYYY-MM-DD format | 
 **checkout** | **string** | Check out data in YYYY-MM-DD format | 
 **currency** | **string** | Currency code - example (USD) | 
 **guestNationality** | **string** | Guest nationality ISO-2 code - example (SG) | 
 **adults** | **int32** | Number of adult guests staying | 
 **children** | **string** | Number of children staying if any | 
 **guestId** | **string** | Unique traveler ID if available | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyAuth](../README.md#apikeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HotelsRatesGet

> map[string]interface{} HotelsRatesGet(ctx).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).GuestNationality(guestNationality).Currency(currency).Adults(adults).Children(children).GuestId(guestId).Execute()

hotel full rates availability



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    hotelIds := "lp3803c,lp1f982,lp19b70,lp19e75" // string | List of hotelIds
    checkin := "2023-11-15" // string | Check in data in YYYY-MM-DD format
    checkout := "2023-11-16" // string | Check out data in YYYY-MM-DD format
    guestNationality := "US" // string | Guest nationality ISO-2 code - example (SG)
    currency := "USD" // string | Currency code - example (USD)
    adults := int32(1) // int32 | Number of adult guests staying
    children := "12,9" // string | Number of children staying if any (optional)
    guestId := "traveler1" // string | Unique traveler ID if available (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.HotelsRatesGet(context.Background()).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).GuestNationality(guestNationality).Currency(currency).Adults(adults).Children(children).GuestId(guestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.HotelsRatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HotelsRatesGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.HotelsRatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHotelsRatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hotelIds** | **string** | List of hotelIds | 
 **checkin** | **string** | Check in data in YYYY-MM-DD format | 
 **checkout** | **string** | Check out data in YYYY-MM-DD format | 
 **guestNationality** | **string** | Guest nationality ISO-2 code - example (SG) | 
 **currency** | **string** | Currency code - example (USD) | 
 **adults** | **int32** | Number of adult guests staying | 
 **children** | **string** | Number of children staying if any | 
 **guestId** | **string** | Unique traveler ID if available | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyAuth](../README.md#apikeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

