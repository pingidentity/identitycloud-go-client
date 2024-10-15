# \SSOCookieAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSSOCookie**](SSOCookieAPI.md#GetSSOCookie) | **Get** /environment/sso-cookie | Get SSO cookie configuration
[**ResetSSOCookie**](SSOCookieAPI.md#ResetSSOCookie) | **Post** /environment/sso-cookie | Reset SSO cookie configuration
[**SetSSOCookie**](SSOCookieAPI.md#SetSSOCookie) | **Put** /environment/sso-cookie | Set SSO cookie configuration



## GetSSOCookie

> SSOCookie GetSSOCookie(ctx).Execute()

Get SSO cookie configuration



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSOCookieAPI.GetSSOCookie(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSOCookieAPI.GetSSOCookie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSOCookie`: SSOCookie
    fmt.Fprintf(os.Stdout, "Response from `SSOCookieAPI.GetSSOCookie`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSOCookieRequest struct via the builder pattern


### Return type

[**SSOCookie**](SSOCookie.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSSOCookie

> SSOCookie ResetSSOCookie(ctx).Action(action).Execute()

Reset SSO cookie configuration



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
    action := "action_example" // string | Reset SSO cookie configuration to default

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSOCookieAPI.ResetSSOCookie(context.Background()).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSOCookieAPI.ResetSSOCookie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetSSOCookie`: SSOCookie
    fmt.Fprintf(os.Stdout, "Response from `SSOCookieAPI.ResetSSOCookie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetSSOCookieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Reset SSO cookie configuration to default | 

### Return type

[**SSOCookie**](SSOCookie.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSSOCookie

> SSOCookie SetSSOCookie(ctx).SSOCookie(sSOCookie).Execute()

Set SSO cookie configuration



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
    sSOCookie := *openapiclient.NewSSOCookie("Name_example") // SSOCookie | SSO cookie configuration to apply to the tenant

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSOCookieAPI.SetSSOCookie(context.Background()).SSOCookie(sSOCookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSOCookieAPI.SetSSOCookie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSSOCookie`: SSOCookie
    fmt.Fprintf(os.Stdout, "Response from `SSOCookieAPI.SetSSOCookie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSSOCookieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sSOCookie** | [**SSOCookie**](SSOCookie.md) | SSO cookie configuration to apply to the tenant | 

### Return type

[**SSOCookie**](SSOCookie.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

