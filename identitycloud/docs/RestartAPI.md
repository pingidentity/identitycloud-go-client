# \RestartAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRestartStatus**](RestartAPI.md#GetRestartStatus) | **Get** /environment/startup | Get restart status
[**Restart**](RestartAPI.md#Restart) | **Post** /environment/startup | Initiate restart



## GetRestartStatus

> EsvRestartStatus GetRestartStatus(ctx).AcceptAPIVersion(acceptAPIVersion).Execute()

Get restart status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/identitycloud-go-client"
)

func main() {
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestartAPI.GetRestartStatus(context.Background()).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestartAPI.GetRestartStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestartStatus`: EsvRestartStatus
    fmt.Fprintf(os.Stdout, "Response from `RestartAPI.GetRestartStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRestartStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvRestartStatus**](EsvRestartStatus.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Restart

> EsvRestartStatus Restart(ctx).Action(action).AcceptAPIVersion(acceptAPIVersion).Execute()

Initiate restart



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/identitycloud-go-client"
)

func main() {
    action := "action_example" // string | Can only be restart
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestartAPI.Restart(context.Background()).Action(action).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestartAPI.Restart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Restart`: EsvRestartStatus
    fmt.Fprintf(os.Stdout, "Response from `RestartAPI.Restart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Can only be restart | 
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvRestartStatus**](EsvRestartStatus.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

