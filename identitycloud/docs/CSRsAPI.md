# \CSRsAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificateSigningRequest**](CSRsAPI.md#CreateCertificateSigningRequest) | **Post** /environment/csrs | Create CSR
[**DeleteCertificateSigningRequestById**](CSRsAPI.md#DeleteCertificateSigningRequestById) | **Delete** /environment/csrs/{id} | Delete CSR by ID
[**GetCertificateSigningRequestById**](CSRsAPI.md#GetCertificateSigningRequestById) | **Get** /environment/csrs/{id} | Get CSR by ID
[**GetCertificateSigningRequests**](CSRsAPI.md#GetCertificateSigningRequests) | **Get** /environment/csrs | Get all CSRs
[**UpdateCertificateSigningRequestById**](CSRsAPI.md#UpdateCertificateSigningRequestById) | **Patch** /environment/csrs/{id} | Update CSR by ID



## CreateCertificateSigningRequest

> CertificateSigningRequest CreateCertificateSigningRequest(ctx).Body(body).Execute()

Create CSR



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
    body := *openapiclient.NewCreateCertificateSigningRequestRequest() // CreateCertificateSigningRequestRequest | JSON body of the new certificate signing request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRsAPI.CreateCertificateSigningRequest(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRsAPI.CreateCertificateSigningRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateSigningRequest`: CertificateSigningRequest
    fmt.Fprintf(os.Stdout, "Response from `CSRsAPI.CreateCertificateSigningRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateSigningRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateCertificateSigningRequestRequest**](CreateCertificateSigningRequestRequest.md) | JSON body of the new certificate signing request | 

### Return type

[**CertificateSigningRequest**](CertificateSigningRequest.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificateSigningRequestById

> DeleteCertificateSigningRequestById(ctx, id).Execute()

Delete CSR by ID



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
    id := "id_example" // string | ID of the certificateSigningRequest

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CSRsAPI.DeleteCertificateSigningRequestById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRsAPI.DeleteCertificateSigningRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the certificateSigningRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateSigningRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateSigningRequestById

> CertificateSigningRequest GetCertificateSigningRequestById(ctx, id).Execute()

Get CSR by ID



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
    id := "id_example" // string | ID of the CSR

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRsAPI.GetCertificateSigningRequestById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRsAPI.GetCertificateSigningRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateSigningRequestById`: CertificateSigningRequest
    fmt.Fprintf(os.Stdout, "Response from `CSRsAPI.GetCertificateSigningRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateSigningRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateSigningRequest**](CertificateSigningRequest.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateSigningRequests

> []CertificateSigningRequest GetCertificateSigningRequests(ctx).Execute()

Get all CSRs



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
    resp, r, err := apiClient.CSRsAPI.GetCertificateSigningRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRsAPI.GetCertificateSigningRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateSigningRequests`: []CertificateSigningRequest
    fmt.Fprintf(os.Stdout, "Response from `CSRsAPI.GetCertificateSigningRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateSigningRequestsRequest struct via the builder pattern


### Return type

[**[]CertificateSigningRequest**](CertificateSigningRequest.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateSigningRequestById

> CertificateSigningRequest UpdateCertificateSigningRequestById(ctx, id).Body(body).Execute()

Update CSR by ID



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
    id := "id_example" // string | ID of the certificateSigningRequest
    body := *openapiclient.NewUpdateCertificateSigningRequestRequest("Certificate_example") // UpdateCertificateSigningRequestRequest | JSON body of the new certificate signing request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSRsAPI.UpdateCertificateSigningRequestById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSRsAPI.UpdateCertificateSigningRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateSigningRequestById`: CertificateSigningRequest
    fmt.Fprintf(os.Stdout, "Response from `CSRsAPI.UpdateCertificateSigningRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the certificateSigningRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateSigningRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateCertificateSigningRequestRequest**](UpdateCertificateSigningRequestRequest.md) | JSON body of the new certificate signing request | 

### Return type

[**CertificateSigningRequest**](CertificateSigningRequest.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

