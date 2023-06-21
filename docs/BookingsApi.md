# \BookingsApi

All URIs are relative to _http://localhost:8080_

| Method                                        | HTTP request      | Description |
| --------------------------------------------- | ----------------- | ----------- |
| [**BookPost**](BookingsApi.md#BookPost)       | **Post** /book    | Book        |
| [**PrebookPost**](BookingsApi.md#PrebookPost) | **Post** /prebook | Prebook     |

## BookPost

> map[string]any BookPost(ctx).Body(body).Execute()

Book

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
    body := map[string]any{ ... } // map[string]any |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookingsApi.BookPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingsApi.BookPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BookPost`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `BookingsApi.BookPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiBookPostRequest struct via the builder pattern

| Name     | Type                       | Description | Notes |
| -------- | -------------------------- | ----------- | ----- |
| **body** | **map[string]any** |             |

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

## PrebookPost

> map[string]any PrebookPost(ctx).Body(body).Execute()

Prebook

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
    body := map[string]any{ ... } // map[string]any |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BookingsApi.PrebookPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingsApi.PrebookPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrebookPost`: map[string]any
    fmt.Fprintf(os.Stdout, "Response from `BookingsApi.PrebookPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPrebookPostRequest struct via the builder pattern

| Name     | Type                       | Description | Notes |
| -------- | -------------------------- | ----------- | ----- |
| **body** | **map[string]any** |             |

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
