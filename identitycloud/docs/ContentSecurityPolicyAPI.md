# \ContentSecurityPolicyAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnforcedContentSecurityPolicy**](ContentSecurityPolicyAPI.md#GetEnforcedContentSecurityPolicy) | **Get** /environment/content-security-policy/enforced | Get enforced content security policy
[**GetReportOnlyContentSecurityPolicy**](ContentSecurityPolicyAPI.md#GetReportOnlyContentSecurityPolicy) | **Get** /environment/content-security-policy/report-only | Get report-only content security policy
[**SetEnforcedContentSecurityPolicy**](ContentSecurityPolicyAPI.md#SetEnforcedContentSecurityPolicy) | **Put** /environment/content-security-policy/enforced | Set enforced content security policy
[**SetReportOnlyContentSecurityPolicy**](ContentSecurityPolicyAPI.md#SetReportOnlyContentSecurityPolicy) | **Put** /environment/content-security-policy/report-only | Set report-only content security policy



## GetEnforcedContentSecurityPolicy

> ContentSecurityPolicy GetEnforcedContentSecurityPolicy(ctx).Execute()

Get enforced content security policy



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
    resp, r, err := apiClient.ContentSecurityPolicyAPI.GetEnforcedContentSecurityPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentSecurityPolicyAPI.GetEnforcedContentSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnforcedContentSecurityPolicy`: ContentSecurityPolicy
    fmt.Fprintf(os.Stdout, "Response from `ContentSecurityPolicyAPI.GetEnforcedContentSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnforcedContentSecurityPolicyRequest struct via the builder pattern


### Return type

[**ContentSecurityPolicy**](ContentSecurityPolicy.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportOnlyContentSecurityPolicy

> ContentSecurityPolicy GetReportOnlyContentSecurityPolicy(ctx).Execute()

Get report-only content security policy



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
    resp, r, err := apiClient.ContentSecurityPolicyAPI.GetReportOnlyContentSecurityPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentSecurityPolicyAPI.GetReportOnlyContentSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportOnlyContentSecurityPolicy`: ContentSecurityPolicy
    fmt.Fprintf(os.Stdout, "Response from `ContentSecurityPolicyAPI.GetReportOnlyContentSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportOnlyContentSecurityPolicyRequest struct via the builder pattern


### Return type

[**ContentSecurityPolicy**](ContentSecurityPolicy.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEnforcedContentSecurityPolicy

> ContentSecurityPolicy SetEnforcedContentSecurityPolicy(ctx).Body(body).Execute()

Set enforced content security policy



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
    body := *openapiclient.NewContentSecurityPolicy() // ContentSecurityPolicy | Enforced content security policy to apply to the tenant

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentSecurityPolicyAPI.SetEnforcedContentSecurityPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentSecurityPolicyAPI.SetEnforcedContentSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEnforcedContentSecurityPolicy`: ContentSecurityPolicy
    fmt.Fprintf(os.Stdout, "Response from `ContentSecurityPolicyAPI.SetEnforcedContentSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetEnforcedContentSecurityPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ContentSecurityPolicy**](ContentSecurityPolicy.md) | Enforced content security policy to apply to the tenant | 

### Return type

[**ContentSecurityPolicy**](ContentSecurityPolicy.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReportOnlyContentSecurityPolicy

> ContentSecurityPolicy SetReportOnlyContentSecurityPolicy(ctx).Body(body).Execute()

Set report-only content security policy



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
    body := *openapiclient.NewContentSecurityPolicy() // ContentSecurityPolicy | Enforced content security policy to apply to the tenant

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetReportOnlyContentSecurityPolicy`: ContentSecurityPolicy
    fmt.Fprintf(os.Stdout, "Response from `ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetReportOnlyContentSecurityPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ContentSecurityPolicy**](ContentSecurityPolicy.md) | Enforced content security policy to apply to the tenant | 

### Return type

[**ContentSecurityPolicy**](ContentSecurityPolicy.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

