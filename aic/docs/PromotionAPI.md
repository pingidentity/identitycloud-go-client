# \PromotionAPI

All URIs are relative to *https://openam-example.id.forgerock.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckLock**](PromotionAPI.md#CheckLock) | **Get** /environment/promotion/lock/state | Get lock status
[**GetLastReport**](PromotionAPI.md#GetLastReport) | **Get** /environment/promotion/report | Get last promotion report
[**GetProvisionalReport**](PromotionAPI.md#GetProvisionalReport) | **Get** /environment/promotion/report/provisional | Get a provisional promotion report
[**GetProvisionalRollbackReport**](PromotionAPI.md#GetProvisionalRollbackReport) | **Get** /environment/promotion/report/provisional-rollback | Get a provisional rollback report
[**GetReport**](PromotionAPI.md#GetReport) | **Get** /environment/promotion/report/{reportId} | Get a promotion report
[**GetReportList**](PromotionAPI.md#GetReportList) | **Get** /environment/promotion/reports | Get promotion reports
[**Lock**](PromotionAPI.md#Lock) | **Post** /environment/promotion/lock | Lock environments
[**Progress**](PromotionAPI.md#Progress) | **Get** /environment/promotion/promote | Get promotion status
[**Rollback**](PromotionAPI.md#Rollback) | **Post** /environment/promotion/rollback | Rollback a promotion
[**Start**](PromotionAPI.md#Start) | **Post** /environment/promotion/promote | Run a promotion
[**Unlock**](PromotionAPI.md#Unlock) | **Delete** /environment/promotion/lock/{promotionId} | Unlock environments



## CheckLock

> PromotionLockState CheckLock(ctx).AcceptAPIVersion(acceptAPIVersion).LocalLockOnly(localLockOnly).Execute()

Get lock status



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0
    localLockOnly := true // bool | If set to true this parameter will only return the lock status of the environment the request is aimed at (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.CheckLock(context.Background()).AcceptAPIVersion(acceptAPIVersion).LocalLockOnly(localLockOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.CheckLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckLock`: PromotionLockState
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.CheckLock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 
 **localLockOnly** | **bool** | If set to true this parameter will only return the lock status of the environment the request is aimed at | 

### Return type

[**PromotionLockState**](PromotionLockState.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastReport

> PromotionReport GetLastReport(ctx).AcceptAPIVersion(acceptAPIVersion).Execute()

Get last promotion report



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.GetLastReport(context.Background()).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.GetLastReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastReport`: PromotionReport
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.GetLastReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLastReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 

### Return type

[**PromotionReport**](PromotionReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisionalReport

> PromotionReport GetProvisionalReport(ctx).AcceptAPIVersion(acceptAPIVersion).Execute()

Get a provisional promotion report



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.GetProvisionalReport(context.Background()).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.GetProvisionalReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvisionalReport`: PromotionReport
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.GetProvisionalReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisionalReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 

### Return type

[**PromotionReport**](PromotionReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisionalRollbackReport

> PromotionReport GetProvisionalRollbackReport(ctx).AcceptAPIVersion(acceptAPIVersion).Execute()

Get a provisional rollback report



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.GetProvisionalRollbackReport(context.Background()).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.GetProvisionalRollbackReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvisionalRollbackReport`: PromotionReport
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.GetProvisionalRollbackReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisionalRollbackReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 

### Return type

[**PromotionReport**](PromotionReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReport

> PromotionReport GetReport(ctx, reportId).AcceptAPIVersion(acceptAPIVersion).Execute()

Get a promotion report



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0
    reportId := "reportId_example" // string | Promotion report unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.GetReport(context.Background(), reportId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.GetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReport`: PromotionReport
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.GetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Promotion report unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 


### Return type

[**PromotionReport**](PromotionReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportList

> []PromotionReportMeta GetReportList(ctx).AcceptAPIVersion(acceptAPIVersion).Execute()

Get promotion reports



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.GetReportList(context.Background()).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.GetReportList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportList`: []PromotionReportMeta
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.GetReportList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 

### Return type

[**[]PromotionReportMeta**](PromotionReportMeta.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Lock

> PromotionStartLocking Lock(ctx).AcceptAPIVersion(acceptAPIVersion).Execute()

Lock environments



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.Lock(context.Background()).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.Lock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Lock`: PromotionStartLocking
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.Lock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 

### Return type

[**PromotionStartLocking**](PromotionStartLocking.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Progress

> PromotionProgress Progress(ctx).AcceptAPIVersion(acceptAPIVersion).Execute()

Get promotion status



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.Progress(context.Background()).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.Progress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Progress`: PromotionProgress
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.Progress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 

### Return type

[**PromotionProgress**](PromotionProgress.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rollback

> RollbackStart Rollback(ctx).AcceptAPIVersion(acceptAPIVersion).RollbackRequest(rollbackRequest).Execute()

Rollback a promotion



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0
    rollbackRequest := *openapiclient.NewRollbackRequest() // RollbackRequest | A request body with info required to initiate a rollback

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.Rollback(context.Background()).AcceptAPIVersion(acceptAPIVersion).RollbackRequest(rollbackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.Rollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Rollback`: RollbackStart
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.Rollback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 
 **rollbackRequest** | [**RollbackRequest**](RollbackRequest.md) | A request body with info required to initiate a rollback | 

### Return type

[**RollbackStart**](RollbackStart.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Start

> PromotionStart Start(ctx).AcceptAPIVersion(acceptAPIVersion).PromotionRequest(promotionRequest).Execute()

Run a promotion



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0
    promotionRequest := *openapiclient.NewPromotionRequest(false) // PromotionRequest | A request body with info required to initiate a promotion

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.Start(context.Background()).AcceptAPIVersion(acceptAPIVersion).PromotionRequest(promotionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.Start``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Start`: PromotionStart
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.Start`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 
 **promotionRequest** | [**PromotionRequest**](PromotionRequest.md) | A request body with info required to initiate a promotion | 

### Return type

[**PromotionStart**](PromotionStart.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unlock

> PromotionUnlocked Unlock(ctx, promotionId).AcceptAPIVersion(acceptAPIVersion).Execute()

Unlock environments



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
    acceptAPIVersion := "acceptAPIVersion_example" // string | protocol=1.0,resource=1.0
    promotionId := "promotionId_example" // string | Promotion unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionAPI.Unlock(context.Background(), promotionId).AcceptAPIVersion(acceptAPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionAPI.Unlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Unlock`: PromotionUnlocked
    fmt.Fprintf(os.Stdout, "Response from `PromotionAPI.Unlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promotionId** | **string** | Promotion unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAPIVersion** | **string** | protocol&#x3D;1.0,resource&#x3D;1.0 | 


### Return type

[**PromotionUnlocked**](PromotionUnlocked.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

