# \VariablesAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionVariable**](VariablesAPI.md#ActionVariable) | **Post** /environment/variables/{variableId} | Set a variable description
[**CreateVariables**](VariablesAPI.md#CreateVariables) | **Put** /environment/variables/{variableId} | Create or update a variable
[**DeleteVariable**](VariablesAPI.md#DeleteVariable) | **Delete** /environment/variables/{variableId} | Delete a variable
[**GetAllVariables**](VariablesAPI.md#GetAllVariables) | **Get** /environment/variables | Get all variables
[**GetVariable**](VariablesAPI.md#GetVariable) | **Get** /environment/variables/{variableId} | Get a variable



## ActionVariable

> ActionVariable(ctx, variableId).Action(action).EsvSetDescriptionRequest(esvSetDescriptionRequest).AcceptAPIVersion(acceptAPIVersion).Execute()

Set a variable description

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
    variableId := "variableId_example" // string | ID of the variable
    action := "action_example" // string | 
    esvSetDescriptionRequest := *openapiclient.NewEsvSetDescriptionRequest("My secret") // EsvSetDescriptionRequest | The description of this variable
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VariablesAPI.ActionVariable(context.Background(), variableId).Action(action).EsvSetDescriptionRequest(esvSetDescriptionRequest).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.ActionVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | ID of the variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** |  | 
 **esvSetDescriptionRequest** | [**EsvSetDescriptionRequest**](EsvSetDescriptionRequest.md) | The description of this variable | 
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

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


## CreateVariables

> EsvVariableResponse CreateVariables(ctx, variableId).EsvVariableCreateUpdateRequest(esvVariableCreateUpdateRequest).AcceptAPIVersion(acceptAPIVersion).Execute()

Create or update a variable



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
    variableId := "variableId_example" // string | ID of the variable
    esvVariableCreateUpdateRequest := *openapiclient.NewEsvVariableCreateUpdateRequest(string([B@18245eb0)) // EsvVariableCreateUpdateRequest | JSON body of the new variable
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesAPI.CreateVariables(context.Background(), variableId).EsvVariableCreateUpdateRequest(esvVariableCreateUpdateRequest).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.CreateVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariables`: EsvVariableResponse
    fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.CreateVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | ID of the variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **esvVariableCreateUpdateRequest** | [**EsvVariableCreateUpdateRequest**](EsvVariableCreateUpdateRequest.md) | JSON body of the new variable | 
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvVariableResponse**](EsvVariableResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVariable

> EsvVariableResponse DeleteVariable(ctx, variableId).AcceptAPIVersion(acceptAPIVersion).Execute()

Delete a variable

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
    variableId := "variableId_example" // string | ID of the variable
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesAPI.DeleteVariable(context.Background(), variableId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.DeleteVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVariable`: EsvVariableResponse
    fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.DeleteVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | ID of the variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvVariableResponse**](EsvVariableResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllVariables

> GetAllVariables200Response GetAllVariables(ctx).AcceptAPIVersion(acceptAPIVersion).PageSize(pageSize).PagedResultsCookie(pagedResultsCookie).PagedResultsOffset(pagedResultsOffset).OnlyPending(onlyPending).Execute()

Get all variables

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
    pageSize := int64(56) // int64 | _Accept-API-Version: resource=2.0_ only, maximum number of results returned by endpoint before paging (optional) (default to 25)
    pagedResultsCookie := "pagedResultsCookie_example" // string | _Accept-API-Version: resource=2.0_ only, opaque data used for paging result data, can be used for paging instead of having to track pageSize and pagedResultsOffset (optional) (default to "")
    pagedResultsOffset := int64(56) // int64 | _Accept-API-Version: resource=2.0_ only, offset of the first result to be returned by endpoint (optional) (default to 0)
    onlyPending := true // bool | _Accept-API-Version: resource=2.0_ only, returns ESVs with unapplied changes (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesAPI.GetAllVariables(context.Background()).AcceptAPIVersion(acceptAPIVersion).PageSize(pageSize).PagedResultsCookie(pagedResultsCookie).PagedResultsOffset(pagedResultsOffset).OnlyPending(onlyPending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.GetAllVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllVariables`: GetAllVariables200Response
    fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.GetAllVariables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 
 **pageSize** | **int64** | _Accept-API-Version: resource&#x3D;2.0_ only, maximum number of results returned by endpoint before paging | [default to 25]
 **pagedResultsCookie** | **string** | _Accept-API-Version: resource&#x3D;2.0_ only, opaque data used for paging result data, can be used for paging instead of having to track pageSize and pagedResultsOffset | [default to &quot;&quot;]
 **pagedResultsOffset** | **int64** | _Accept-API-Version: resource&#x3D;2.0_ only, offset of the first result to be returned by endpoint | [default to 0]
 **onlyPending** | **bool** | _Accept-API-Version: resource&#x3D;2.0_ only, returns ESVs with unapplied changes | [default to false]

### Return type

[**GetAllVariables200Response**](GetAllVariables200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariable

> EsvVariableResponse GetVariable(ctx, variableId).AcceptAPIVersion(acceptAPIVersion).Execute()

Get a variable

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
    variableId := "variableId_example" // string | ID of the variable
    acceptAPIVersion := "acceptAPIVersion_example" // string | resource=2.0 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesAPI.GetVariable(context.Background(), variableId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.GetVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariable`: EsvVariableResponse
    fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.GetVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | ID of the variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptAPIVersion** | **string** | resource&#x3D;2.0 | 

### Return type

[**EsvVariableResponse**](EsvVariableResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

