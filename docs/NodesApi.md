# \NodesApi

All URIs are relative to *https://127.0.0.1:8006/api2/json*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainer**](NodesApi.md#CreateContainer) | **Post** /nodes/{node}/lxc | createContainer
[**CreateContainerSnapshot**](NodesApi.md#CreateContainerSnapshot) | **Post** /nodes/{node}/lxc/{vmid}/snapshot | createContainerSnapshot
[**CreateNodesSingleStorageSingleContent**](NodesApi.md#CreateNodesSingleStorageSingleContent) | **Post** /nodes/{node}/storage/{storage}/content | createNodesSingleStorageSingleContent
[**CreateVM**](NodesApi.md#CreateVM) | **Post** /nodes/{node}/qemu | createVM
[**CreateVMSnapshot**](NodesApi.md#CreateVMSnapshot) | **Post** /nodes/{node}/qemu/{vmid}/snapshot | createVMSnapshot
[**DeleteContainer**](NodesApi.md#DeleteContainer) | **Delete** /nodes/{node}/lxc/{vmid} | deleteContainer
[**DeleteContainerSnapshot**](NodesApi.md#DeleteContainerSnapshot) | **Delete** /nodes/{node}/lxc/{vmid}/snapshot/{snapname} | deleteContainerSnapshot
[**DeleteVM**](NodesApi.md#DeleteVM) | **Delete** /nodes/{node}/qemu/{vmid} | deleteVM
[**DeleteVMSnapshot**](NodesApi.md#DeleteVMSnapshot) | **Delete** /nodes/{node}/qemu/{vmid}/snapshot/{snapname} | deleteVMSnapshot
[**GetContainer**](NodesApi.md#GetContainer) | **Get** /nodes/{node}/lxc/{vmid} | getContainer
[**GetContainerConfig**](NodesApi.md#GetContainerConfig) | **Get** /nodes/{node}/lxc/{vmid}/config | getContainerConfig
[**GetContainerConfigPending**](NodesApi.md#GetContainerConfigPending) | **Get** /nodes/{node}/lxc/{vmid}/pending | getContainerConfigPending
[**GetContainerSnapshot**](NodesApi.md#GetContainerSnapshot) | **Get** /nodes/{node}/lxc/{vmid}/snapshot/{snapname} | getContainerSnapshot
[**GetContainerSnapshotConfig**](NodesApi.md#GetContainerSnapshotConfig) | **Get** /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config | getContainerSnapshotConfig
[**GetContainerSnapshots**](NodesApi.md#GetContainerSnapshots) | **Get** /nodes/{node}/lxc/{vmid}/snapshot | getContainerSnapshots
[**GetContainerStatus**](NodesApi.md#GetContainerStatus) | **Get** /nodes/{node}/lxc/{vmid}/status | getContainerStatus
[**GetContainers**](NodesApi.md#GetContainers) | **Get** /nodes/{node}/lxc | getContainers
[**GetCurrentContainerStatus**](NodesApi.md#GetCurrentContainerStatus) | **Get** /nodes/{node}/lxc/{vmid}/status/current | getCurrentContainerStatus
[**GetCurrentVMStatus**](NodesApi.md#GetCurrentVMStatus) | **Get** /nodes/{node}/qemu/{vmid}/status/current | getCurrentVMStatus
[**GetNodeTask**](NodesApi.md#GetNodeTask) | **Get** /nodes/{node}/tasks/{upid} | getNodeTask
[**GetNodeTaskLog**](NodesApi.md#GetNodeTaskLog) | **Get** /nodes/{node}/tasks/{upid}/log | getNodeTaskLog
[**GetNodeTaskStatus**](NodesApi.md#GetNodeTaskStatus) | **Get** /nodes/{node}/tasks/{upid}/status | getNodeTaskStatus
[**GetNodeTasks**](NodesApi.md#GetNodeTasks) | **Get** /nodes/{node}/tasks | getNodeTasks
[**GetStorageContent**](NodesApi.md#GetStorageContent) | **Get** /nodes/{node}/storage/{storage}/content | getStorageContent
[**GetStorages**](NodesApi.md#GetStorages) | **Get** /nodes/{node}/storage | getStorages
[**GetVM**](NodesApi.md#GetVM) | **Get** /nodes/{node}/qemu/{vmid} | getVM
[**GetVMConfig**](NodesApi.md#GetVMConfig) | **Get** /nodes/{node}/qemu/{vmid}/config | getVMConfig
[**GetVMConfigPending**](NodesApi.md#GetVMConfigPending) | **Get** /nodes/{node}/qemu/{vmid}/pending | getVMConfigPending
[**GetVMSnapshot**](NodesApi.md#GetVMSnapshot) | **Get** /nodes/{node}/qemu/{vmid}/snapshot/{snapname} | getVMSnapshot
[**GetVMSnapshotConfig**](NodesApi.md#GetVMSnapshotConfig) | **Get** /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/config | getVMSnapshotConfig
[**GetVMSnapshots**](NodesApi.md#GetVMSnapshots) | **Get** /nodes/{node}/qemu/{vmid}/snapshot | getVMSnapshots
[**GetVMs**](NodesApi.md#GetVMs) | **Get** /nodes/{node}/qemu | getVMs
[**RebootContainer**](NodesApi.md#RebootContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/reboot | rebootContainer
[**RebootVM**](NodesApi.md#RebootVM) | **Post** /nodes/{node}/qemu/{vmid}/status/reboot | rebootVM
[**ResizeContainerDisk**](NodesApi.md#ResizeContainerDisk) | **Put** /nodes/{node}/lxc/{vmid}/resize | resizeContainerDisk
[**ResizeVMDisk**](NodesApi.md#ResizeVMDisk) | **Put** /nodes/{node}/qemu/{vmid}/resize | resizeVMDisk
[**ResumeContainer**](NodesApi.md#ResumeContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/resume | resumeContainer
[**ResumeVM**](NodesApi.md#ResumeVM) | **Post** /nodes/{node}/qemu/{vmid}/status/resume | resumeVM
[**RollbackContainerSnapshot**](NodesApi.md#RollbackContainerSnapshot) | **Post** /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/rollback | rollbackContainerSnapshot
[**RollbackVMSnapshot**](NodesApi.md#RollbackVMSnapshot) | **Post** /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/rollback | rollbackVMSnapshot
[**ShutdownContainer**](NodesApi.md#ShutdownContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/shutdown | shutdownContainer
[**ShutdownVM**](NodesApi.md#ShutdownVM) | **Post** /nodes/{node}/qemu/{vmid}/status/shutdown | shutdownVM
[**StartContainer**](NodesApi.md#StartContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/start | startContainer
[**StartVM**](NodesApi.md#StartVM) | **Post** /nodes/{node}/qemu/{vmid}/status/start | startVM
[**StopContainer**](NodesApi.md#StopContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/stop | stopContainer
[**StopNodeTask**](NodesApi.md#StopNodeTask) | **Delete** /nodes/{node}/tasks/{upid} | stopNodeTask
[**StopVM**](NodesApi.md#StopVM) | **Post** /nodes/{node}/qemu/{vmid}/status/stop | stopVM
[**SuspendContainer**](NodesApi.md#SuspendContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/suspend | suspendContainer
[**SuspendVM**](NodesApi.md#SuspendVM) | **Post** /nodes/{node}/qemu/{vmid}/status/suspend | suspendVM
[**UpdateContainerConfigSync**](NodesApi.md#UpdateContainerConfigSync) | **Put** /nodes/{node}/lxc/{vmid}/config | updateContainerConfigSync
[**UpdateContainerSnapshotConfig**](NodesApi.md#UpdateContainerSnapshotConfig) | **Put** /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config | updateContainerSnapshotConfig
[**UpdateVMConfig**](NodesApi.md#UpdateVMConfig) | **Post** /nodes/{node}/qemu/{vmid}/config | updateVMConfig
[**UpdateVMConfigSync**](NodesApi.md#UpdateVMConfigSync) | **Put** /nodes/{node}/qemu/{vmid}/config | updateVMConfigSync
[**UpdateVMSnapshotConfig**](NodesApi.md#UpdateVMSnapshotConfig) | **Put** /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/config | updateVMSnapshotConfig
[**UploadFile**](NodesApi.md#UploadFile) | **Post** /nodes/{node}/storage/{storage}/upload | uploadFile



## CreateContainer

> CreateVM200Response CreateContainer(ctx, node).CreateContainerRequest(createContainerRequest).Execute()

createContainer



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
    node := "node_example" // string | node
    createContainerRequest := *openapiclient.NewCreateContainerRequest("Ostemplate_example", int64(123)) // CreateContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CreateContainer(context.Background(), node).CreateContainerRequest(createContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CreateContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CreateContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createContainerRequest** | [**CreateContainerRequest**](CreateContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainerSnapshot

> TaskStartResponse CreateContainerSnapshot(ctx, node, vmid).CreateContainerSnapshotRequest(createContainerSnapshotRequest).Execute()

createContainerSnapshot



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    createContainerSnapshotRequest := *openapiclient.NewCreateContainerSnapshotRequest("Snapname_example") // CreateContainerSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CreateContainerSnapshot(context.Background(), node, vmid).CreateContainerSnapshotRequest(createContainerSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CreateContainerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CreateContainerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createContainerSnapshotRequest** | [**CreateContainerSnapshotRequest**](CreateContainerSnapshotRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNodesSingleStorageSingleContent

> CreateNodesSingleStorageSingleContent200Response CreateNodesSingleStorageSingleContent(ctx, node, storage).CreateNodesSingleStorageSingleContentRequest(createNodesSingleStorageSingleContentRequest).Execute()

createNodesSingleStorageSingleContent



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
    node := "node_example" // string | node
    storage := "storage_example" // string | storage
    createNodesSingleStorageSingleContentRequest := *openapiclient.NewCreateNodesSingleStorageSingleContentRequest("TODO", "Size_example", int64(123)) // CreateNodesSingleStorageSingleContentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CreateNodesSingleStorageSingleContent(context.Background(), node, storage).CreateNodesSingleStorageSingleContentRequest(createNodesSingleStorageSingleContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CreateNodesSingleStorageSingleContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNodesSingleStorageSingleContent`: CreateNodesSingleStorageSingleContent200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CreateNodesSingleStorageSingleContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**storage** | **string** | storage | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodesSingleStorageSingleContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNodesSingleStorageSingleContentRequest** | [**CreateNodesSingleStorageSingleContentRequest**](CreateNodesSingleStorageSingleContentRequest.md) |  | 

### Return type

[**CreateNodesSingleStorageSingleContent200Response**](CreateNodesSingleStorageSingleContent200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVM

> CreateVM200Response CreateVM(ctx, node).CreateVMRequest(createVMRequest).Execute()

createVM



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
    node := "node_example" // string | node
    createVMRequest := *openapiclient.NewCreateVMRequest(int64(123)) // CreateVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CreateVM(context.Background(), node).CreateVMRequest(createVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CreateVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CreateVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVMRequest** | [**CreateVMRequest**](CreateVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVMSnapshot

> TaskStartResponse CreateVMSnapshot(ctx, node, vmid).CreateVMSnapshotRequest(createVMSnapshotRequest).Execute()

createVMSnapshot



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    createVMSnapshotRequest := *openapiclient.NewCreateVMSnapshotRequest("Snapname_example") // CreateVMSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.CreateVMSnapshot(context.Background(), node, vmid).CreateVMSnapshotRequest(createVMSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.CreateVMSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVMSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.CreateVMSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVMSnapshotRequest** | [**CreateVMSnapshotRequest**](CreateVMSnapshotRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainer

> CreateVM200Response DeleteContainer(ctx, node, vmid).Execute()

deleteContainer



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.DeleteContainer(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.DeleteContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.DeleteContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerSnapshot

> TaskStartResponse DeleteContainerSnapshot(ctx, node, snapname, vmid).Execute()

deleteContainerSnapshot



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.DeleteContainerSnapshot(context.Background(), node, snapname, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.DeleteContainerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContainerSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.DeleteContainerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVM

> CreateVM200Response DeleteVM(ctx, node, vmid).Execute()

deleteVM



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.DeleteVM(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.DeleteVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.DeleteVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVMSnapshot

> TaskStartResponse DeleteVMSnapshot(ctx, node, snapname, vmid).Execute()

deleteVMSnapshot



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.DeleteVMSnapshot(context.Background(), node, snapname, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.DeleteVMSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVMSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.DeleteVMSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainer

> GetVM200Response GetContainer(ctx, node, vmid).Execute()

getContainer



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetContainer(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainer`: GetVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVM200Response**](GetVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerConfig

> GetContainerConfig200Response GetContainerConfig(ctx, node, vmid).Execute()

getContainerConfig



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetContainerConfig(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetContainerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerConfig`: GetContainerConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetContainerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetContainerConfig200Response**](GetContainerConfig200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerConfigPending

> GetContainerConfigPending200Response GetContainerConfigPending(ctx, node, vmid).Execute()

getContainerConfigPending



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetContainerConfigPending(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetContainerConfigPending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerConfigPending`: GetContainerConfigPending200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetContainerConfigPending`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerConfigPendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetContainerConfigPending200Response**](GetContainerConfigPending200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSnapshot

> GetVMSnapshot200Response GetContainerSnapshot(ctx, node, snapname, vmid).Execute()

getContainerSnapshot



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetContainerSnapshot(context.Background(), node, snapname, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetContainerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSnapshot`: GetVMSnapshot200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetContainerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVMSnapshot200Response**](GetVMSnapshot200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSnapshotConfig

> GetVMSnapshotConfig200Response GetContainerSnapshotConfig(ctx, node, snapname, vmid).Execute()

getContainerSnapshotConfig



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetContainerSnapshotConfig(context.Background(), node, snapname, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetContainerSnapshotConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSnapshotConfig`: GetVMSnapshotConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetContainerSnapshotConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSnapshotConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVMSnapshotConfig200Response**](GetVMSnapshotConfig200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSnapshots

> GetContainerSnapshots200Response GetContainerSnapshots(ctx, node, vmid).Execute()

getContainerSnapshots



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetContainerSnapshots(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetContainerSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSnapshots`: GetContainerSnapshots200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetContainerSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetContainerSnapshots200Response**](GetContainerSnapshots200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerStatus

> GetVM200Response GetContainerStatus(ctx, node, vmid).Execute()

getContainerStatus



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetContainerStatus(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetContainerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerStatus`: GetVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetContainerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVM200Response**](GetVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainers

> GetContainers200Response GetContainers(ctx, node).Execute()

getContainers



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
    node := "node_example" // string | node

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetContainers(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainers`: GetContainers200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContainers200Response**](GetContainers200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentContainerStatus

> GetCurrentContainerStatus200Response GetCurrentContainerStatus(ctx, node, vmid).Execute()

getCurrentContainerStatus



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetCurrentContainerStatus(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetCurrentContainerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentContainerStatus`: GetCurrentContainerStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetCurrentContainerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentContainerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCurrentContainerStatus200Response**](GetCurrentContainerStatus200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentVMStatus

> GetCurrentVMStatus200Response GetCurrentVMStatus(ctx, node, vmid).Execute()

getCurrentVMStatus



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetCurrentVMStatus(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetCurrentVMStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentVMStatus`: GetCurrentVMStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetCurrentVMStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentVMStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCurrentVMStatus200Response**](GetCurrentVMStatus200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTask

> GetVMSnapshot200Response GetNodeTask(ctx, node, upid).Execute()

getNodeTask



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
    node := "node_example" // string | node
    upid := "upid_example" // string | upid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetNodeTask(context.Background(), node, upid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetNodeTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeTask`: GetVMSnapshot200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetNodeTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**upid** | **string** | upid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVMSnapshot200Response**](GetVMSnapshot200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTaskLog

> GetNodeTaskLog200Response GetNodeTaskLog(ctx, node, upid).Execute()

getNodeTaskLog



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
    node := "node_example" // string | node
    upid := "upid_example" // string | upid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetNodeTaskLog(context.Background(), node, upid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetNodeTaskLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeTaskLog`: GetNodeTaskLog200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetNodeTaskLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**upid** | **string** | upid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNodeTaskLog200Response**](GetNodeTaskLog200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTaskStatus

> GetNodeTaskStatus200Response GetNodeTaskStatus(ctx, node, upid).Execute()

getNodeTaskStatus



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
    node := "node_example" // string | node
    upid := "upid_example" // string | upid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetNodeTaskStatus(context.Background(), node, upid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetNodeTaskStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeTaskStatus`: GetNodeTaskStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetNodeTaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**upid** | **string** | upid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNodeTaskStatus200Response**](GetNodeTaskStatus200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTasks

> GetNodeTasks200Response GetNodeTasks(ctx, node).Execute()

getNodeTasks



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
    node := "node_example" // string | node

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetNodeTasks(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetNodeTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeTasks`: GetNodeTasks200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetNodeTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNodeTasks200Response**](GetNodeTasks200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageContent

> GetStorageContent200Response GetStorageContent(ctx, node, storage).Execute()

getStorageContent



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
    node := "node_example" // string | node
    storage := "storage_example" // string | storage

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetStorageContent(context.Background(), node, storage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetStorageContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageContent`: GetStorageContent200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetStorageContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**storage** | **string** | storage | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetStorageContent200Response**](GetStorageContent200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorages

> GetStorages200Response GetStorages(ctx, node).Execute()

getStorages



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
    node := "node_example" // string | node

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetStorages(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetStorages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorages`: GetStorages200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStorages200Response**](GetStorages200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVM

> GetVM200Response GetVM(ctx, node, vmid).Execute()

getVM



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetVM(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVM`: GetVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVM200Response**](GetVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMConfig

> GetVMConfig200Response GetVMConfig(ctx, node, vmid).Execute()

getVMConfig



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetVMConfig(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetVMConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMConfig`: GetVMConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetVMConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVMConfig200Response**](GetVMConfig200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMConfigPending

> GetVMConfigPending200Response GetVMConfigPending(ctx, node, vmid).Execute()

getVMConfigPending



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetVMConfigPending(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetVMConfigPending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMConfigPending`: GetVMConfigPending200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetVMConfigPending`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMConfigPendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVMConfigPending200Response**](GetVMConfigPending200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMSnapshot

> GetVMSnapshot200Response GetVMSnapshot(ctx, node, snapname, vmid).Execute()

getVMSnapshot



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetVMSnapshot(context.Background(), node, snapname, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetVMSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMSnapshot`: GetVMSnapshot200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetVMSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVMSnapshot200Response**](GetVMSnapshot200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMSnapshotConfig

> GetVMSnapshotConfig200Response GetVMSnapshotConfig(ctx, node, snapname, vmid).Execute()

getVMSnapshotConfig



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetVMSnapshotConfig(context.Background(), node, snapname, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetVMSnapshotConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMSnapshotConfig`: GetVMSnapshotConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetVMSnapshotConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMSnapshotConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVMSnapshotConfig200Response**](GetVMSnapshotConfig200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMSnapshots

> GetVMSnapshots200Response GetVMSnapshots(ctx, node, vmid).Execute()

getVMSnapshots



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetVMSnapshots(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetVMSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMSnapshots`: GetVMSnapshots200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetVMSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVMSnapshots200Response**](GetVMSnapshots200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMs

> GetVMs200Response GetVMs(ctx, node).Execute()

getVMs



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
    node := "node_example" // string | node

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.GetVMs(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.GetVMs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMs`: GetVMs200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.GetVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVMs200Response**](GetVMs200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootContainer

> CreateVM200Response RebootContainer(ctx, node, vmid).RebootContainerRequest(rebootContainerRequest).Execute()

rebootContainer



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    rebootContainerRequest := *openapiclient.NewRebootContainerRequest() // RebootContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.RebootContainer(context.Background(), node, vmid).RebootContainerRequest(rebootContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.RebootContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.RebootContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rebootContainerRequest** | [**RebootContainerRequest**](RebootContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootVM

> CreateVM200Response RebootVM(ctx, node, vmid).RebootVMRequest(rebootVMRequest).Execute()

rebootVM



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    rebootVMRequest := *openapiclient.NewRebootVMRequest() // RebootVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.RebootVM(context.Background(), node, vmid).RebootVMRequest(rebootVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.RebootVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.RebootVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rebootVMRequest** | [**RebootVMRequest**](RebootVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeContainerDisk

> TaskStartResponse ResizeContainerDisk(ctx, node, vmid).ResizeContainerDiskRequest(resizeContainerDiskRequest).Execute()

resizeContainerDisk



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    resizeContainerDiskRequest := *openapiclient.NewResizeContainerDiskRequest("Disk_example", "Size_example") // ResizeContainerDiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ResizeContainerDisk(context.Background(), node, vmid).ResizeContainerDiskRequest(resizeContainerDiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ResizeContainerDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeContainerDisk`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ResizeContainerDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeContainerDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resizeContainerDiskRequest** | [**ResizeContainerDiskRequest**](ResizeContainerDiskRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeVMDisk

> TaskStartResponse ResizeVMDisk(ctx, node, vmid).ResizeVMDiskRequest(resizeVMDiskRequest).Execute()

resizeVMDisk



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    resizeVMDiskRequest := *openapiclient.NewResizeVMDiskRequest("Disk_example", "Size_example") // ResizeVMDiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ResizeVMDisk(context.Background(), node, vmid).ResizeVMDiskRequest(resizeVMDiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ResizeVMDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeVMDisk`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ResizeVMDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeVMDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resizeVMDiskRequest** | [**ResizeVMDiskRequest**](ResizeVMDiskRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeContainer

> CreateVM200Response ResumeContainer(ctx, node, vmid).Body(body).Execute()

resumeContainer



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ResumeContainer(context.Background(), node, vmid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ResumeContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ResumeContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeVM

> CreateVM200Response ResumeVM(ctx, node, vmid).ResumeVMRequest(resumeVMRequest).Execute()

resumeVM



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    resumeVMRequest := *openapiclient.NewResumeVMRequest() // ResumeVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ResumeVM(context.Background(), node, vmid).ResumeVMRequest(resumeVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ResumeVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ResumeVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resumeVMRequest** | [**ResumeVMRequest**](ResumeVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackContainerSnapshot

> TaskStartResponse RollbackContainerSnapshot(ctx, node, snapname, vmid).RollbackContainerSnapshotRequest(rollbackContainerSnapshotRequest).Execute()

rollbackContainerSnapshot



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid
    rollbackContainerSnapshotRequest := *openapiclient.NewRollbackContainerSnapshotRequest() // RollbackContainerSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.RollbackContainerSnapshot(context.Background(), node, snapname, vmid).RollbackContainerSnapshotRequest(rollbackContainerSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.RollbackContainerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackContainerSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.RollbackContainerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackContainerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rollbackContainerSnapshotRequest** | [**RollbackContainerSnapshotRequest**](RollbackContainerSnapshotRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackVMSnapshot

> TaskStartResponse RollbackVMSnapshot(ctx, node, snapname, vmid).RollbackVMSnapshotRequest(rollbackVMSnapshotRequest).Execute()

rollbackVMSnapshot



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid
    rollbackVMSnapshotRequest := *openapiclient.NewRollbackVMSnapshotRequest() // RollbackVMSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.RollbackVMSnapshot(context.Background(), node, snapname, vmid).RollbackVMSnapshotRequest(rollbackVMSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.RollbackVMSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackVMSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.RollbackVMSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackVMSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rollbackVMSnapshotRequest** | [**RollbackVMSnapshotRequest**](RollbackVMSnapshotRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShutdownContainer

> CreateVM200Response ShutdownContainer(ctx, node, vmid).ShutdownContainerRequest(shutdownContainerRequest).Execute()

shutdownContainer



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    shutdownContainerRequest := *openapiclient.NewShutdownContainerRequest() // ShutdownContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ShutdownContainer(context.Background(), node, vmid).ShutdownContainerRequest(shutdownContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ShutdownContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShutdownContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ShutdownContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shutdownContainerRequest** | [**ShutdownContainerRequest**](ShutdownContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShutdownVM

> CreateVM200Response ShutdownVM(ctx, node, vmid).ShutdownVMRequest(shutdownVMRequest).Execute()

shutdownVM



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    shutdownVMRequest := *openapiclient.NewShutdownVMRequest() // ShutdownVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.ShutdownVM(context.Background(), node, vmid).ShutdownVMRequest(shutdownVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.ShutdownVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShutdownVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.ShutdownVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shutdownVMRequest** | [**ShutdownVMRequest**](ShutdownVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartContainer

> CreateVM200Response StartContainer(ctx, node, vmid).StartContainerRequest(startContainerRequest).Execute()

startContainer



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    startContainerRequest := *openapiclient.NewStartContainerRequest() // StartContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.StartContainer(context.Background(), node, vmid).StartContainerRequest(startContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.StartContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.StartContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startContainerRequest** | [**StartContainerRequest**](StartContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVM

> CreateVM200Response StartVM(ctx, node, vmid).StartVMRequest(startVMRequest).Execute()

startVM



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    startVMRequest := *openapiclient.NewStartVMRequest() // StartVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.StartVM(context.Background(), node, vmid).StartVMRequest(startVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.StartVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.StartVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startVMRequest** | [**StartVMRequest**](StartVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopContainer

> CreateVM200Response StopContainer(ctx, node, vmid).StopContainerRequest(stopContainerRequest).Execute()

stopContainer



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    stopContainerRequest := *openapiclient.NewStopContainerRequest() // StopContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.StopContainer(context.Background(), node, vmid).StopContainerRequest(stopContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.StopContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.StopContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stopContainerRequest** | [**StopContainerRequest**](StopContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopNodeTask

> CreateVM200Response StopNodeTask(ctx, node, upid).Execute()

stopNodeTask



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
    node := "node_example" // string | node
    upid := "upid_example" // string | upid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.StopNodeTask(context.Background(), node, upid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.StopNodeTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopNodeTask`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.StopNodeTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**upid** | **string** | upid | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopNodeTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVM

> CreateVM200Response StopVM(ctx, node, vmid).StopVMRequest(stopVMRequest).Execute()

stopVM



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    stopVMRequest := *openapiclient.NewStopVMRequest() // StopVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.StopVM(context.Background(), node, vmid).StopVMRequest(stopVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.StopVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.StopVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stopVMRequest** | [**StopVMRequest**](StopVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendContainer

> CreateVM200Response SuspendContainer(ctx, node, vmid).Body(body).Execute()

suspendContainer



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.SuspendContainer(context.Background(), node, vmid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.SuspendContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.SuspendContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendVM

> CreateVM200Response SuspendVM(ctx, node, vmid).SuspendVMRequest(suspendVMRequest).Execute()

suspendVM



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    suspendVMRequest := *openapiclient.NewSuspendVMRequest() // SuspendVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.SuspendVM(context.Background(), node, vmid).SuspendVMRequest(suspendVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.SuspendVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.SuspendVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suspendVMRequest** | [**SuspendVMRequest**](SuspendVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerConfigSync

> CreateVM200Response UpdateContainerConfigSync(ctx, node, vmid).UpdateContainerConfigSyncRequest(updateContainerConfigSyncRequest).Execute()

updateContainerConfigSync



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    updateContainerConfigSyncRequest := *openapiclient.NewUpdateContainerConfigSyncRequest() // UpdateContainerConfigSyncRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UpdateContainerConfigSync(context.Background(), node, vmid).UpdateContainerConfigSyncRequest(updateContainerConfigSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UpdateContainerConfigSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContainerConfigSync`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UpdateContainerConfigSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerConfigSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateContainerConfigSyncRequest** | [**UpdateContainerConfigSyncRequest**](UpdateContainerConfigSyncRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerSnapshotConfig

> CreateVM200Response UpdateContainerSnapshotConfig(ctx, node, snapname, vmid).UpdateContainerSnapshotConfigRequest(updateContainerSnapshotConfigRequest).Execute()

updateContainerSnapshotConfig



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid
    updateContainerSnapshotConfigRequest := *openapiclient.NewUpdateContainerSnapshotConfigRequest() // UpdateContainerSnapshotConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UpdateContainerSnapshotConfig(context.Background(), node, snapname, vmid).UpdateContainerSnapshotConfigRequest(updateContainerSnapshotConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UpdateContainerSnapshotConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContainerSnapshotConfig`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UpdateContainerSnapshotConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerSnapshotConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateContainerSnapshotConfigRequest** | [**UpdateContainerSnapshotConfigRequest**](UpdateContainerSnapshotConfigRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMConfig

> CreateVM200Response UpdateVMConfig(ctx, node, vmid).UpdateVMConfigRequest(updateVMConfigRequest).Execute()

updateVMConfig



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    updateVMConfigRequest := *openapiclient.NewUpdateVMConfigRequest() // UpdateVMConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UpdateVMConfig(context.Background(), node, vmid).UpdateVMConfigRequest(updateVMConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UpdateVMConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVMConfig`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UpdateVMConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMConfigRequest** | [**UpdateVMConfigRequest**](UpdateVMConfigRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMConfigSync

> CreateVM200Response UpdateVMConfigSync(ctx, node, vmid).UpdateVMConfigSyncRequest(updateVMConfigSyncRequest).Execute()

updateVMConfigSync



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
    node := "node_example" // string | node
    vmid := TODO // int64 | vmid
    updateVMConfigSyncRequest := *openapiclient.NewUpdateVMConfigSyncRequest() // UpdateVMConfigSyncRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UpdateVMConfigSync(context.Background(), node, vmid).UpdateVMConfigSyncRequest(updateVMConfigSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UpdateVMConfigSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVMConfigSync`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UpdateVMConfigSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMConfigSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMConfigSyncRequest** | [**UpdateVMConfigSyncRequest**](UpdateVMConfigSyncRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMSnapshotConfig

> CreateVM200Response UpdateVMSnapshotConfig(ctx, node, snapname, vmid).UpdateVMSnapshotConfigRequest(updateVMSnapshotConfigRequest).Execute()

updateVMSnapshotConfig



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
    node := "node_example" // string | node
    snapname := "snapname_example" // string | snapname
    vmid := TODO // int64 | vmid
    updateVMSnapshotConfigRequest := *openapiclient.NewUpdateVMSnapshotConfigRequest() // UpdateVMSnapshotConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UpdateVMSnapshotConfig(context.Background(), node, snapname, vmid).UpdateVMSnapshotConfigRequest(updateVMSnapshotConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UpdateVMSnapshotConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVMSnapshotConfig`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UpdateVMSnapshotConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**snapname** | **string** | snapname | 
**vmid** | [**int64**](.md) | vmid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMSnapshotConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateVMSnapshotConfigRequest** | [**UpdateVMSnapshotConfigRequest**](UpdateVMSnapshotConfigRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> CreateVM200Response UploadFile(ctx, node, storage).Content(content).Filename(filename).Checksum(checksum).ChecksumAlgorithm(checksumAlgorithm).Tmpfilename(tmpfilename).Execute()

uploadFile



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
    node := "node_example" // string | node
    storage := "storage_example" // string | storage
    content := "content_example" // string | 
    filename := os.NewFile(1234, "some_file") // *os.File | 
    checksum := "checksum_example" // string |  (optional)
    checksumAlgorithm := "checksumAlgorithm_example" // string |  (optional)
    tmpfilename := "tmpfilename_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesApi.UploadFile(context.Background(), node, storage).Content(content).Filename(filename).Checksum(checksum).ChecksumAlgorithm(checksumAlgorithm).Tmpfilename(tmpfilename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesApi.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesApi.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | node | 
**storage** | **string** | storage | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | **string** |  | 
 **filename** | ***os.File** |  | 
 **checksum** | **string** |  | 
 **checksumAlgorithm** | **string** |  | 
 **tmpfilename** | **string** |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

