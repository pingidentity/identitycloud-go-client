# EsvVariableListV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PagedResultsCookie** | **NullableString** |  | 
**RemainingPagedResults** | **int64** |  | 
**Result** | [**[]EsvVariableResponse**](EsvVariableResponse.md) |  | 
**ResultCount** | **int64** |  | 
**TotalPagedResults** | **int64** |  | 
**TotalPagedResultsPolicy** | **string** |  | 

## Methods

### NewEsvVariableListV1

`func NewEsvVariableListV1(pagedResultsCookie NullableString, remainingPagedResults int64, result []EsvVariableResponse, resultCount int64, totalPagedResults int64, totalPagedResultsPolicy string, ) *EsvVariableListV1`

NewEsvVariableListV1 instantiates a new EsvVariableListV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsvVariableListV1WithDefaults

`func NewEsvVariableListV1WithDefaults() *EsvVariableListV1`

NewEsvVariableListV1WithDefaults instantiates a new EsvVariableListV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagedResultsCookie

`func (o *EsvVariableListV1) GetPagedResultsCookie() string`

GetPagedResultsCookie returns the PagedResultsCookie field if non-nil, zero value otherwise.

### GetPagedResultsCookieOk

`func (o *EsvVariableListV1) GetPagedResultsCookieOk() (*string, bool)`

GetPagedResultsCookieOk returns a tuple with the PagedResultsCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagedResultsCookie

`func (o *EsvVariableListV1) SetPagedResultsCookie(v string)`

SetPagedResultsCookie sets PagedResultsCookie field to given value.


### SetPagedResultsCookieNil

`func (o *EsvVariableListV1) SetPagedResultsCookieNil(b bool)`

 SetPagedResultsCookieNil sets the value for PagedResultsCookie to be an explicit nil

### UnsetPagedResultsCookie
`func (o *EsvVariableListV1) UnsetPagedResultsCookie()`

UnsetPagedResultsCookie ensures that no value is present for PagedResultsCookie, not even an explicit nil
### GetRemainingPagedResults

`func (o *EsvVariableListV1) GetRemainingPagedResults() int64`

GetRemainingPagedResults returns the RemainingPagedResults field if non-nil, zero value otherwise.

### GetRemainingPagedResultsOk

`func (o *EsvVariableListV1) GetRemainingPagedResultsOk() (*int64, bool)`

GetRemainingPagedResultsOk returns a tuple with the RemainingPagedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingPagedResults

`func (o *EsvVariableListV1) SetRemainingPagedResults(v int64)`

SetRemainingPagedResults sets RemainingPagedResults field to given value.


### GetResult

`func (o *EsvVariableListV1) GetResult() []EsvVariableResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EsvVariableListV1) GetResultOk() (*[]EsvVariableResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EsvVariableListV1) SetResult(v []EsvVariableResponse)`

SetResult sets Result field to given value.


### GetResultCount

`func (o *EsvVariableListV1) GetResultCount() int64`

GetResultCount returns the ResultCount field if non-nil, zero value otherwise.

### GetResultCountOk

`func (o *EsvVariableListV1) GetResultCountOk() (*int64, bool)`

GetResultCountOk returns a tuple with the ResultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCount

`func (o *EsvVariableListV1) SetResultCount(v int64)`

SetResultCount sets ResultCount field to given value.


### GetTotalPagedResults

`func (o *EsvVariableListV1) GetTotalPagedResults() int64`

GetTotalPagedResults returns the TotalPagedResults field if non-nil, zero value otherwise.

### GetTotalPagedResultsOk

`func (o *EsvVariableListV1) GetTotalPagedResultsOk() (*int64, bool)`

GetTotalPagedResultsOk returns a tuple with the TotalPagedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPagedResults

`func (o *EsvVariableListV1) SetTotalPagedResults(v int64)`

SetTotalPagedResults sets TotalPagedResults field to given value.


### GetTotalPagedResultsPolicy

`func (o *EsvVariableListV1) GetTotalPagedResultsPolicy() string`

GetTotalPagedResultsPolicy returns the TotalPagedResultsPolicy field if non-nil, zero value otherwise.

### GetTotalPagedResultsPolicyOk

`func (o *EsvVariableListV1) GetTotalPagedResultsPolicyOk() (*string, bool)`

GetTotalPagedResultsPolicyOk returns a tuple with the TotalPagedResultsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPagedResultsPolicy

`func (o *EsvVariableListV1) SetTotalPagedResultsPolicy(v string)`

SetTotalPagedResultsPolicy sets TotalPagedResultsPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


