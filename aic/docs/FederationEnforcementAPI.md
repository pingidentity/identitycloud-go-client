# \FederationEnforcementAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnforcement**](FederationEnforcementAPI.md#GetEnforcement) | **Get** /environment/federation/enforcement | Get enforcement of federation
[**SetEnforcement**](FederationEnforcementAPI.md#SetEnforcement) | **Put** /environment/federation/enforcement | Set enforcement of federation



## GetEnforcement

> FederationEnforcement GetEnforcement(ctx).Execute()

Get enforcement of federation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/aic-temp-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederationEnforcementAPI.GetEnforcement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationEnforcementAPI.GetEnforcement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnforcement`: FederationEnforcement
    fmt.Fprintf(os.Stdout, "Response from `FederationEnforcementAPI.GetEnforcement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnforcementRequest struct via the builder pattern


### Return type

[**FederationEnforcement**](FederationEnforcement.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEnforcement

> FederationEnforcement SetEnforcement(ctx).FederationEnforcement(federationEnforcement).Execute()

Set enforcement of federation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/aic-temp-go-client"
)

func main() {
    federationEnforcement := *openapiclient.NewFederationEnforcement("Groups_example") // FederationEnforcement | Enforcement settings to apply to the tenant

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederationEnforcementAPI.SetEnforcement(context.Background()).FederationEnforcement(federationEnforcement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationEnforcementAPI.SetEnforcement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEnforcement`: FederationEnforcement
    fmt.Fprintf(os.Stdout, "Response from `FederationEnforcementAPI.SetEnforcement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetEnforcementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **federationEnforcement** | [**FederationEnforcement**](FederationEnforcement.md) | Enforcement settings to apply to the tenant | 

### Return type

[**FederationEnforcement**](FederationEnforcement.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

