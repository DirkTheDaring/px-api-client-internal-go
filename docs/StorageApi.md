# \StorageApi

All URIs are relative to *https://127.0.0.1:8006/api2/json*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorage**](StorageApi.md#CreateStorage) | **Post** /storage | createStorage
[**GetStorage**](StorageApi.md#GetStorage) | **Get** /storage | getStorage



## CreateStorage

> CreateStorage200Response CreateStorage(ctx).CreateStorageRequest(createStorageRequest).Execute()

createStorage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/proxmox-openapi-go"
)

func main() {
    createStorageRequest := *openapiclient.NewCreateStorageRequest("Storage_example", "Type_example") // CreateStorageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageApi.CreateStorage(context.Background()).CreateStorageRequest(createStorageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.CreateStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorage`: CreateStorage200Response
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.CreateStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStorageRequest** | [**CreateStorageRequest**](CreateStorageRequest.md) |  | 

### Return type

[**CreateStorage200Response**](CreateStorage200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorage

> GetStorage200Response GetStorage(ctx).Execute()

getStorage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/proxmox-openapi-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageApi.GetStorage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageApi.GetStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorage`: GetStorage200Response
    fmt.Fprintf(os.Stdout, "Response from `StorageApi.GetStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageRequest struct via the builder pattern


### Return type

[**GetStorage200Response**](GetStorage200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

