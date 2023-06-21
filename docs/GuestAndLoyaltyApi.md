# \GuestAndLoyaltyApi

All URIs are relative to *https://api.liteapi.travel/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GuestsGet**](GuestAndLoyaltyApi.md#GuestsGet) | **Get** /guests | guests



## GuestsGet

> map[string]interface{} GuestsGet(ctx).Email(email).Execute()

guests



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/liteapi-travel/go-sdk"
)

func main() {
    email := "johndoe@nlite.ml" // string | Email ID of the guest (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GuestAndLoyaltyApi.GuestsGet(context.Background()).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GuestAndLoyaltyApi.GuestsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GuestsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GuestAndLoyaltyApi.GuestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email ID of the guest | 

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

