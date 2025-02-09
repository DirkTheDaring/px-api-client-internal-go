# \ClusterApi

All URIs are relative to *https://127.0.0.1:8006/api2/json*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterConfigNodes**](ClusterApi.md#GetClusterConfigNodes) | **Get** /cluster/config/nodes | getClusterConfigNodes
[**GetClusterNextid**](ClusterApi.md#GetClusterNextid) | **Get** /cluster/nextid | getClusterNextid
[**GetClusterResources**](ClusterApi.md#GetClusterResources) | **Get** /cluster/resources | getClusterResources



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
    openapiclient "github.com/DirkTheDaring/proxmox-openapi-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.GetClusterConfigNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetClusterConfigNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterConfigNodes`: GetClusterConfigNodes200Response
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetClusterConfigNodes`: %v\n", resp)
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

> GetClusterNextid200Response GetClusterNextid(ctx).Execute()

getClusterNextid



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
    resp, r, err := apiClient.ClusterApi.GetClusterNextid(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetClusterNextid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNextid`: GetClusterNextid200Response
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetClusterNextid`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNextidRequest struct via the builder pattern


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

> GetClusterResources200Response GetClusterResources(ctx).Execute()

getClusterResources



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
    resp, r, err := apiClient.ClusterApi.GetClusterResources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetClusterResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterResources`: GetClusterResources200Response
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetClusterResources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterResourcesRequest struct via the builder pattern


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

