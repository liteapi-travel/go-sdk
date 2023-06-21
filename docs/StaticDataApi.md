# \StaticDataApi

All URIs are relative to *https://api.liteapi.travel/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataCitiesGet**](StaticDataApi.md#DataCitiesGet) | **Get** /data/cities | City list
[**DataCountriesGet**](StaticDataApi.md#DataCountriesGet) | **Get** /data/countries | Country list
[**DataCurrenciesGet**](StaticDataApi.md#DataCurrenciesGet) | **Get** /data/currencies | Currency list
[**DataHotelGet**](StaticDataApi.md#DataHotelGet) | **Get** /data/hotel | Hotel details
[**DataHotelsGet**](StaticDataApi.md#DataHotelsGet) | **Get** /data/hotels | Hotel list
[**DataIataCodesGet**](StaticDataApi.md#DataIataCodesGet) | **Get** /data/iataCodes | IATA code list



## DataCitiesGet

> map[string]interface{} DataCitiesGet(ctx).CountryCode(countryCode).Execute()

City list



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
    countryCode := "SG" // string | Country code in iso-2 format (example: SG)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticDataApi.DataCitiesGet(context.Background()).CountryCode(countryCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.DataCitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataCitiesGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaticDataApi.DataCitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataCitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | Country code in iso-2 format (example: SG) | 

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


## DataCountriesGet

> map[string]interface{} DataCountriesGet(ctx).Execute()

Country list



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticDataApi.DataCountriesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.DataCountriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataCountriesGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaticDataApi.DataCountriesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataCountriesGetRequest struct via the builder pattern


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


## DataCurrenciesGet

> map[string]interface{} DataCurrenciesGet(ctx).Execute()

Currency list



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticDataApi.DataCurrenciesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.DataCurrenciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataCurrenciesGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaticDataApi.DataCurrenciesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataCurrenciesGetRequest struct via the builder pattern


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


## DataHotelGet

> map[string]interface{} DataHotelGet(ctx).HotelId(hotelId).Execute()

Hotel details



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
    hotelId := int32(57871) // int32 | Unique ID of a hotel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticDataApi.DataHotelGet(context.Background()).HotelId(hotelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.DataHotelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataHotelGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaticDataApi.DataHotelGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataHotelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hotelId** | **int32** | Unique ID of a hotel | 

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


## DataHotelsGet

> map[string]interface{} DataHotelsGet(ctx).CountryCode(countryCode).CityName(cityName).Offset(offset).Limit(limit).Longitude(longitude).Latitude(latitude).Distance(distance).Execute()

Hotel list



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
    countryCode := "SG" // string | country code ISO-2 code - example (SG)
    cityName := "Singapore" // string | name of the city
    offset := int32(0) // int32 | specifies the number of rows to skip before starting to return rows (optional)
    limit := int32(1000) // int32 | limit number of results (max 1000) (optional)
    longitude := float32(-115.16988) // float32 | longitude geo coordinates (optional)
    latitude := float32(36.12510) // float32 | latitude geo coordinates (optional)
    distance := int32(1000) // int32 | distance in meters (min 1000m) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticDataApi.DataHotelsGet(context.Background()).CountryCode(countryCode).CityName(cityName).Offset(offset).Limit(limit).Longitude(longitude).Latitude(latitude).Distance(distance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.DataHotelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataHotelsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaticDataApi.DataHotelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataHotelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | country code ISO-2 code - example (SG) | 
 **cityName** | **string** | name of the city | 
 **offset** | **int32** | specifies the number of rows to skip before starting to return rows | 
 **limit** | **int32** | limit number of results (max 1000) | 
 **longitude** | **float32** | longitude geo coordinates | 
 **latitude** | **float32** | latitude geo coordinates | 
 **distance** | **int32** | distance in meters (min 1000m) | 

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


## DataIataCodesGet

> map[string]interface{} DataIataCodesGet(ctx).Execute()

IATA code list



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticDataApi.DataIataCodesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.DataIataCodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataIataCodesGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaticDataApi.DataIataCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataIataCodesGetRequest struct via the builder pattern


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

