# \BookingManagementApi

All URIs are relative to *https://api.liteapi.travel/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookingsBookingIdGet**](BookingManagementApi.md#BookingsBookingIdGet) | **Get** /bookings/{bookingId} | Booking retrieve
[**BookingsBookingIdPut**](BookingManagementApi.md#BookingsBookingIdPut) | **Put** /bookings/{bookingId} | Booking cancel
[**BookingsGet**](BookingManagementApi.md#BookingsGet) | **Get** /bookings | Booking list



## BookingsBookingIdGet

> map[string]interface{} BookingsBookingIdGet(ctx, bookingId).Execute()

Booking retrieve



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
    bookingId := " hSq2gVDrf" // string | The Booking Id that needs to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookingManagementApi.BookingsBookingIdGet(context.Background(), bookingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingManagementApi.BookingsBookingIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingsBookingIdGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BookingManagementApi.BookingsBookingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingId** | **string** | The Booking Id that needs to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingsBookingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BookingsBookingIdPut

> map[string]interface{} BookingsBookingIdPut(ctx, bookingId).Execute()

Booking cancel



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
    bookingId := "hSq2gVDrf" // string | (Required) The unique identifier of the booking you would like to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookingManagementApi.BookingsBookingIdPut(context.Background(), bookingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingManagementApi.BookingsBookingIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingsBookingIdPut`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BookingManagementApi.BookingsBookingIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingId** | **string** | (Required) The unique identifier of the booking you would like to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookingsBookingIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BookingsGet

> map[string]interface{} BookingsGet(ctx).GuestId(guestId).Execute()

Booking list



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
    guestId := "FrT56hfty" // string | The Guest Id of the user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookingManagementApi.BookingsGet(context.Background()).GuestId(guestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingManagementApi.BookingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookingsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BookingManagementApi.BookingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guestId** | **string** | The Guest Id of the user | 

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

