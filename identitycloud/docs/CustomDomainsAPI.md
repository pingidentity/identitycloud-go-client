# \CustomDomainsAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomDomains**](CustomDomainsAPI.md#GetCustomDomains) | **Get** /environment/custom-domains/{realm} | Get custom domains
[**SetCustomDomains**](CustomDomainsAPI.md#SetCustomDomains) | **Put** /environment/custom-domains/{realm} | Set custom domains
[**VerifyCustomDomains**](CustomDomainsAPI.md#VerifyCustomDomains) | **Post** /environment/custom-domains | Verify a CNAME



## GetCustomDomains

> CustomDomains GetCustomDomains(ctx, realm).Execute()

Get custom domains



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
    realm := "realm_example" // string | Realm for the domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsAPI.GetCustomDomains(context.Background(), realm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsAPI.GetCustomDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomDomains`: CustomDomains
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainsAPI.GetCustomDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | Realm for the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomDomains**](CustomDomains.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCustomDomains

> CustomDomains SetCustomDomains(ctx, realm).Body(body).Execute()

Set custom domains



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
    realm := "realm_example" // string | Realm for the domain
    body := *openapiclient.NewCustomDomains() // CustomDomains | Custom domains

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsAPI.SetCustomDomains(context.Background(), realm).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsAPI.SetCustomDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCustomDomains`: CustomDomains
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainsAPI.SetCustomDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | Realm for the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCustomDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CustomDomains**](CustomDomains.md) | Custom domains | 

### Return type

[**CustomDomains**](CustomDomains.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCustomDomains

> VerifyCustomDomains(ctx).Action(action).Body(body).Execute()

Verify a CNAME



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
    action := "action_example" // string | Requested action type
    body := *openapiclient.NewCName("Name_example") // CName | Custom domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomDomainsAPI.VerifyCustomDomains(context.Background()).Action(action).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsAPI.VerifyCustomDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyCustomDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Requested action type | 
 **body** | [**CName**](CName.md) | Custom domain | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

