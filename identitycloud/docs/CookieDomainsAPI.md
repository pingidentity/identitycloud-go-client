# \CookieDomainsAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCookieDomains**](CookieDomainsAPI.md#GetCookieDomains) | **Get** /environment/cookie-domains | Get cookie domains
[**SetCookieDomains**](CookieDomainsAPI.md#SetCookieDomains) | **Put** /environment/cookie-domains | Set cookie domains



## GetCookieDomains

> CookieDomains GetCookieDomains(ctx).Execute()

Get cookie domains



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
    resp, r, err := apiClient.CookieDomainsAPI.GetCookieDomains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CookieDomainsAPI.GetCookieDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCookieDomains`: CookieDomains
    fmt.Fprintf(os.Stdout, "Response from `CookieDomainsAPI.GetCookieDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCookieDomainsRequest struct via the builder pattern


### Return type

[**CookieDomains**](CookieDomains.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCookieDomains

> CookieDomains SetCookieDomains(ctx).CookieDomains(cookieDomains).Execute()

Set cookie domains



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
    cookieDomains := *openapiclient.NewCookieDomains() // CookieDomains | Cookie domains

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CookieDomainsAPI.SetCookieDomains(context.Background()).CookieDomains(cookieDomains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CookieDomainsAPI.SetCookieDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCookieDomains`: CookieDomains
    fmt.Fprintf(os.Stdout, "Response from `CookieDomainsAPI.SetCookieDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetCookieDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookieDomains** | [**CookieDomains**](CookieDomains.md) | Cookie domains | 

### Return type

[**CookieDomains**](CookieDomains.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

