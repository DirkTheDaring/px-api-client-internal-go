# \ClusterAPI

All URIs are relative to *https://127.0.0.1:8006/api2/json*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterConfigNodes**](ClusterAPI.md#GetClusterConfigNodes) | **Get** /cluster/config/nodes | getClusterConfigNodes
[**GetClusterNextid**](ClusterAPI.md#GetClusterNextid) | **Get** /cluster/nextid | getClusterNextid
[**GetClusterResources**](ClusterAPI.md#GetClusterResources) | **Get** /cluster/resources | getClusterResources



## GetClusterConfigNodes

> GetClusterConfigNodes200Response GetClusterConfigNodes(ctx).Execute()

getClusterConfigNodes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-internal-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterAPI.GetClusterConfigNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.GetClusterConfigNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterConfigNodes`: GetClusterConfigNodes200Response
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.GetClusterConfigNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterConfigNodesRequest struct via the builder pattern


### Return type

[**GetClusterConfigNodes200Response**](GetClusterConfigNodes200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNextid

> GetClusterNextid200Response GetClusterNextid(ctx).Vmid(vmid).Execute()

getClusterNextid



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-internal-go"
)

func main() {
    vmid := int64(789) // int64 | The (unique) ID of the VM. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterAPI.GetClusterNextid(context.Background()).Vmid(vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.GetClusterNextid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNextid`: GetClusterNextid200Response
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.GetClusterNextid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNextidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vmid** | **int64** | The (unique) ID of the VM. | 

### Return type

[**GetClusterNextid200Response**](GetClusterNextid200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterResources

> GetClusterResources200Response GetClusterResources(ctx).Type_(type_).Execute()

getClusterResources



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-internal-go"
)

func main() {
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterAPI.GetClusterResources(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.GetClusterResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterResources`: GetClusterResources200Response
    fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.GetClusterResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 

### Return type

[**GetClusterResources200Response**](GetClusterResources200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

