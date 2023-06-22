# \BookApi

All URIs are relative to *https://api.liteapi.travel/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RatesBookPost**](BookApi.md#RatesBookPost) | **Post** /rates/book | hotel rate book
[**RatesPrebookPost**](BookApi.md#RatesPrebookPost) | **Post** /rates/prebook | hotel rate prebook



## RatesBookPost

> map[string]interface{} RatesBookPost(ctx).RatesBookPostRequest(ratesBookPostRequest).Execute()

hotel rate book



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
    ratesBookPostRequest := *openapiclient.NewRatesBookPostRequest("PrebookId_example", *openapiclient.NewRatesBookPostRequestGuestInfo("GuestFirstName_example", "GuestLastName_example", "GuestEmail_example")) // RatesBookPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY") 
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookApi.RatesBookPost(context.Background()).RatesBookPostRequest(ratesBookPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.RatesBookPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RatesBookPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BookApi.RatesBookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatesBookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratesBookPostRequest** | [**RatesBookPostRequest**](RatesBookPostRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyAuth](../README.md#apikeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RatesPrebookPost

> map[string]interface{} RatesPrebookPost(ctx).Body(body).Execute()

hotel rate prebook



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY") 
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookApi.RatesPrebookPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookApi.RatesPrebookPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RatesPrebookPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BookApi.RatesPrebookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatesPrebookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[apikeyAuth](../README.md#apikeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

