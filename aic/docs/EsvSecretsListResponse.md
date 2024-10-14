# EsvSecretsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PagedResultsCookie** | **NullableString** |  | 
**RemainingPagedResults** | **int64** |  | 
**Result** | [**[]EsvSecretResponse**](EsvSecretResponse.md) |  | 
**ResultCount** | **int64** |  | 
**TotalPagedResults** | **int64** |  | 
**TotalPagedResultsPolicy** | **string** |  | 

## Methods

### NewEsvSecretsListResponse

`func NewEsvSecretsListResponse(pagedResultsCookie NullableString, remainingPagedResults int64, result []EsvSecretResponse, resultCount int64, totalPagedResults int64, totalPagedResultsPolicy string, ) *EsvSecretsListResponse`

NewEsvSecretsListResponse instantiates a new EsvSecretsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsvSecretsListResponseWithDefaults

`func NewEsvSecretsListResponseWithDefaults() *EsvSecretsListResponse`

NewEsvSecretsListResponseWithDefaults instantiates a new EsvSecretsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagedResultsCookie

`func (o *EsvSecretsListResponse) GetPagedResultsCookie() string`

GetPagedResultsCookie returns the PagedResultsCookie field if non-nil, zero value otherwise.

### GetPagedResultsCookieOk

`func (o *EsvSecretsListResponse) GetPagedResultsCookieOk() (*string, bool)`

GetPagedResultsCookieOk returns a tuple with the PagedResultsCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagedResultsCookie

`func (o *EsvSecretsListResponse) SetPagedResultsCookie(v string)`

SetPagedResultsCookie sets PagedResultsCookie field to given value.


### SetPagedResultsCookieNil

`func (o *EsvSecretsListResponse) SetPagedResultsCookieNil(b bool)`

 SetPagedResultsCookieNil sets the value for PagedResultsCookie to be an explicit nil

### UnsetPagedResultsCookie
`func (o *EsvSecretsListResponse) UnsetPagedResultsCookie()`

UnsetPagedResultsCookie ensures that no value is present for PagedResultsCookie, not even an explicit nil
### GetRemainingPagedResults

`func (o *EsvSecretsListResponse) GetRemainingPagedResults() int64`

GetRemainingPagedResults returns the RemainingPagedResults field if non-nil, zero value otherwise.

### GetRemainingPagedResultsOk

`func (o *EsvSecretsListResponse) GetRemainingPagedResultsOk() (*int64, bool)`

GetRemainingPagedResultsOk returns a tuple with the RemainingPagedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingPagedResults

`func (o *EsvSecretsListResponse) SetRemainingPagedResults(v int64)`

SetRemainingPagedResults sets RemainingPagedResults field to given value.


### GetResult

`func (o *EsvSecretsListResponse) GetResult() []EsvSecretResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EsvSecretsListResponse) GetResultOk() (*[]EsvSecretResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EsvSecretsListResponse) SetResult(v []EsvSecretResponse)`

SetResult sets Result field to given value.


### GetResultCount

`func (o *EsvSecretsListResponse) GetResultCount() int64`

GetResultCount returns the ResultCount field if non-nil, zero value otherwise.

### GetResultCountOk

`func (o *EsvSecretsListResponse) GetResultCountOk() (*int64, bool)`

GetResultCountOk returns a tuple with the ResultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCount

`func (o *EsvSecretsListResponse) SetResultCount(v int64)`

SetResultCount sets ResultCount field to given value.


### GetTotalPagedResults

`func (o *EsvSecretsListResponse) GetTotalPagedResults() int64`

GetTotalPagedResults returns the TotalPagedResults field if non-nil, zero value otherwise.

### GetTotalPagedResultsOk

`func (o *EsvSecretsListResponse) GetTotalPagedResultsOk() (*int64, bool)`

GetTotalPagedResultsOk returns a tuple with the TotalPagedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPagedResults

`func (o *EsvSecretsListResponse) SetTotalPagedResults(v int64)`

SetTotalPagedResults sets TotalPagedResults field to given value.


### GetTotalPagedResultsPolicy

`func (o *EsvSecretsListResponse) GetTotalPagedResultsPolicy() string`

GetTotalPagedResultsPolicy returns the TotalPagedResultsPolicy field if non-nil, zero value otherwise.

### GetTotalPagedResultsPolicyOk

`func (o *EsvSecretsListResponse) GetTotalPagedResultsPolicyOk() (*string, bool)`

GetTotalPagedResultsPolicyOk returns a tuple with the TotalPagedResultsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPagedResultsPolicy

`func (o *EsvSecretsListResponse) SetTotalPagedResultsPolicy(v string)`

SetTotalPagedResultsPolicy sets TotalPagedResultsPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


