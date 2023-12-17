# \AccessAPI

All URIs are relative to *https://127.0.0.1:8006/api2/json*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessTicket**](AccessAPI.md#CreateAccessTicket) | **Post** /access/ticket | createAccessTicket
[**GetAccessTicket**](AccessAPI.md#GetAccessTicket) | **Get** /access/ticket | getAccessTicket



## CreateAccessTicket

> CreateAccessTicket200Response CreateAccessTicket(ctx).CreateAccessTicketRequest(createAccessTicketRequest).Execute()

createAccessTicket



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
    createAccessTicketRequest := *openapiclient.NewCreateAccessTicketRequest("Password_example", "Username_example") // CreateAccessTicketRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessAPI.CreateAccessTicket(context.Background()).CreateAccessTicketRequest(createAccessTicketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.CreateAccessTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessTicket`: CreateAccessTicket200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessAPI.CreateAccessTicket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessTicketRequest** | [**CreateAccessTicketRequest**](CreateAccessTicketRequest.md) |  | 

### Return type

[**CreateAccessTicket200Response**](CreateAccessTicket200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessTicket

> CreateVM200Response GetAccessTicket(ctx).Execute()

getAccessTicket



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
    resp, r, err := apiClient.AccessAPI.GetAccessTicket(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessAPI.GetAccessTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessTicket`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessAPI.GetAccessTicket`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTicketRequest struct via the builder pattern


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

