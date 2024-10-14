# \CountAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCount**](CountAPI.md#GetCount) | **Get** /environment/count | Get count of ESVs



## GetCount

> EsvCount GetCount(ctx).AcceptAPIVersion(acceptAPIVersion).OnlyPending(onlyPending).Execute()

Get count of ESVs



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)
    onlyPending := true // bool | _Accept-API-Version: resource=2.0_ only, returns ESVs with unapplied changes (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountAPI.GetCount(context.Background()).AcceptAPIVersion(acceptAPIVersion).OnlyPending(onlyPending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountAPI.GetCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount`: EsvCount
    fmt.Fprintf(os.Stdout, "Response from `CountAPI.GetCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 
 **onlyPending** | **bool** | _Accept-API-Version: resource&#x3D;2.0_ only, returns ESVs with unapplied changes | [default to false]

### Return type

[**EsvCount**](EsvCount.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

