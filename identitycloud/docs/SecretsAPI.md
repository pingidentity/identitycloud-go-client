# \SecretsAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionSecret**](SecretsAPI.md#ActionSecret) | **Post** /environment/secrets/{secretId} | Set a secret description
[**ChangeSecretVersion**](SecretsAPI.md#ChangeSecretVersion) | **Post** /environment/secrets/{secretId}/versions/{versionId} | Update the status of a version of a secret
[**CreateSecret**](SecretsAPI.md#CreateSecret) | **Put** /environment/secrets/{secretId} | Create a secret
[**CreateSecretVersion**](SecretsAPI.md#CreateSecretVersion) | **Post** /environment/secrets/{secretId}/versions | Create a new version of a secret
[**DeleteSecret**](SecretsAPI.md#DeleteSecret) | **Delete** /environment/secrets/{secretId} | Delete a secret
[**DeleteSecretVersion**](SecretsAPI.md#DeleteSecretVersion) | **Delete** /environment/secrets/{secretId}/versions/{versionId} | Delete a version of a secret
[**GetAllSecrets**](SecretsAPI.md#GetAllSecrets) | **Get** /environment/secrets | Get all secrets
[**GetSecret**](SecretsAPI.md#GetSecret) | **Get** /environment/secrets/{secretId} | Get a secret
[**GetSecretVersion**](SecretsAPI.md#GetSecretVersion) | **Get** /environment/secrets/{secretId}/versions/{versionId} | Get a version of a secret
[**GetSecretVersions**](SecretsAPI.md#GetSecretVersions) | **Get** /environment/secrets/{secretId}/versions | Get all versions of a secret



## ActionSecret

> ActionSecret(ctx, secretId).Action(action).EsvSetDescriptionRequest(esvSetDescriptionRequest).Execute()

Set a secret description

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
    secretId := "secretId_example" // string | ID of the secret
    action := "action_example" // string | 
    esvSetDescriptionRequest := *openapiclient.NewEsvSetDescriptionRequest("My secret") // EsvSetDescriptionRequest | The description of this secret

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecretsAPI.ActionSecret(context.Background(), secretId).Action(action).EsvSetDescriptionRequest(esvSetDescriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.ActionSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** |  | 
 **esvSetDescriptionRequest** | [**EsvSetDescriptionRequest**](EsvSetDescriptionRequest.md) | The description of this secret | 

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


## ChangeSecretVersion

> EsvSecretVersionResponse ChangeSecretVersion(ctx, secretId, versionId).Action(action).EsvSecretVersionStatusRequest(esvSecretVersionStatusRequest).AcceptAPIVersion(acceptAPIVersion).Execute()

Update the status of a version of a secret



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
    action := "action_example" // string | Can only be changestatus
    secretId := "secretId_example" // string | ID of the secret
    versionId := "versionId_example" // string | ID of the secret version
    esvSecretVersionStatusRequest := *openapiclient.NewEsvSecretVersionStatusRequest("DISABLED") // EsvSecretVersionStatusRequest | JSON body of the new status of the secret version
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.ChangeSecretVersion(context.Background(), secretId, versionId).Action(action).EsvSecretVersionStatusRequest(esvSecretVersionStatusRequest).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.ChangeSecretVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeSecretVersion`: EsvSecretVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.ChangeSecretVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 
**versionId** | **string** | ID of the secret version | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeSecretVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Can only be changestatus | 


 **esvSecretVersionStatusRequest** | [**EsvSecretVersionStatusRequest**](EsvSecretVersionStatusRequest.md) | JSON body of the new status of the secret version | 
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvSecretVersionResponse**](EsvSecretVersionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecret

> EsvSecretResponse CreateSecret(ctx, secretId).EsvSecretCreateRequest(esvSecretCreateRequest).AcceptAPIVersion(acceptAPIVersion).Execute()

Create a secret



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
    secretId := "secretId_example" // string | ID of the secret
    esvSecretCreateRequest := *openapiclient.NewEsvSecretCreateRequest("generic", false, string([B@5be067de)) // EsvSecretCreateRequest | JSON body of the new secret
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.CreateSecret(context.Background(), secretId).EsvSecretCreateRequest(esvSecretCreateRequest).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.CreateSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecret`: EsvSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.CreateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esvSecretCreateRequest** | [**EsvSecretCreateRequest**](EsvSecretCreateRequest.md) | JSON body of the new secret | 
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvSecretResponse**](EsvSecretResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecretVersion

> EsvSecretVersionResponse CreateSecretVersion(ctx, secretId).Action(action).EsvSecretVersionCreateRequest(esvSecretVersionCreateRequest).AcceptAPIVersion(acceptAPIVersion).Execute()

Create a new version of a secret

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
    action := "action_example" // string | Can only be create
    secretId := "secretId_example" // string | ID of the secret
    esvSecretVersionCreateRequest := *openapiclient.NewEsvSecretVersionCreateRequest(string([B@7383eae2)) // EsvSecretVersionCreateRequest | JSON body of the new secret version
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.CreateSecretVersion(context.Background(), secretId).Action(action).EsvSecretVersionCreateRequest(esvSecretVersionCreateRequest).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.CreateSecretVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecretVersion`: EsvSecretVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.CreateSecretVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Can only be create | 

 **esvSecretVersionCreateRequest** | [**EsvSecretVersionCreateRequest**](EsvSecretVersionCreateRequest.md) | JSON body of the new secret version | 
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvSecretVersionResponse**](EsvSecretVersionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecret

> EsvSecretResponse DeleteSecret(ctx, secretId).AcceptAPIVersion(acceptAPIVersion).Execute()

Delete a secret



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
    secretId := "secretId_example" // string | ID of the secret
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.DeleteSecret(context.Background(), secretId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.DeleteSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSecret`: EsvSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.DeleteSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvSecretResponse**](EsvSecretResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecretVersion

> EsvSecretVersionResponse DeleteSecretVersion(ctx, secretId, versionId).AcceptAPIVersion(acceptAPIVersion).Execute()

Delete a version of a secret

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
    secretId := "secretId_example" // string | ID of the secret
    versionId := "versionId_example" // string | ID of the secret version
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.DeleteSecretVersion(context.Background(), secretId, versionId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.DeleteSecretVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSecretVersion`: EsvSecretVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.DeleteSecretVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 
**versionId** | **string** | ID of the secret version | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvSecretVersionResponse**](EsvSecretVersionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSecrets

> EsvSecretsListResponse GetAllSecrets(ctx).AcceptAPIVersion(acceptAPIVersion).PageSize(pageSize).PagedResultsCookie(pagedResultsCookie).PagedResultsOffset(pagedResultsOffset).OnlyPending(onlyPending).Execute()

Get all secrets



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
    pageSize := int64(56) // int64 | _Accept-API-Version: resource=2.0_ only, maximum number of results returned by endpoint before paging (optional) (default to 25)
    pagedResultsCookie := "pagedResultsCookie_example" // string | _Accept-API-Version: resource=2.0_ only, opaque data used for paging result data, can be used for paging instead of having to track pageSize and pagedResultsOffset (optional) (default to "")
    pagedResultsOffset := int64(56) // int64 | _Accept-API-Version: resource=2.0_ only, offset of the first result to be returned by endpoint (optional) (default to 0)
    onlyPending := true // bool | _Accept-API-Version: resource=2.0_ only, returns ESVs with unapplied changes (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.GetAllSecrets(context.Background()).AcceptAPIVersion(acceptAPIVersion).PageSize(pageSize).PagedResultsCookie(pagedResultsCookie).PagedResultsOffset(pagedResultsOffset).OnlyPending(onlyPending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetAllSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSecrets`: EsvSecretsListResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetAllSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 
 **pageSize** | **int64** | _Accept-API-Version: resource&#x3D;2.0_ only, maximum number of results returned by endpoint before paging | [default to 25]
 **pagedResultsCookie** | **string** | _Accept-API-Version: resource&#x3D;2.0_ only, opaque data used for paging result data, can be used for paging instead of having to track pageSize and pagedResultsOffset | [default to &quot;&quot;]
 **pagedResultsOffset** | **int64** | _Accept-API-Version: resource&#x3D;2.0_ only, offset of the first result to be returned by endpoint | [default to 0]
 **onlyPending** | **bool** | _Accept-API-Version: resource&#x3D;2.0_ only, returns ESVs with unapplied changes | [default to false]

### Return type

[**EsvSecretsListResponse**](EsvSecretsListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecret

> EsvSecretResponse GetSecret(ctx, secretId).AcceptAPIVersion(acceptAPIVersion).Execute()

Get a secret



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
    secretId := "secretId_example" // string | ID of the secret
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.GetSecret(context.Background(), secretId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecret`: EsvSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvSecretResponse**](EsvSecretResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretVersion

> EsvSecretVersionResponse GetSecretVersion(ctx, secretId, versionId).AcceptAPIVersion(acceptAPIVersion).Execute()

Get a version of a secret



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
    secretId := "secretId_example" // string | ID of the secret
    versionId := "versionId_example" // string | ID of the secret version
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.GetSecretVersion(context.Background(), secretId, versionId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretVersion`: EsvSecretVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 
**versionId** | **string** | ID of the secret version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvSecretVersionResponse**](EsvSecretVersionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretVersions

> []EsvSecretVersionResponse GetSecretVersions(ctx, secretId).AcceptAPIVersion(acceptAPIVersion).Execute()

Get all versions of a secret



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
    secretId := "secretId_example" // string | ID of the secret
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsAPI.GetSecretVersions(context.Background(), secretId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretVersions`: []EsvSecretVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | ID of the secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**[]EsvSecretVersionResponse**](EsvSecretVersionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

