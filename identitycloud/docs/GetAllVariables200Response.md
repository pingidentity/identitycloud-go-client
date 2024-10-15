# GetAllVariables200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PagedResultsCookie** | **NullableString** |  | 
**RemainingPagedResults** | **int64** |  | 
**Result** | [**[]EsvVariableResponseV2**](EsvVariableResponseV2.md) |  | 
**ResultCount** | **int64** |  | 
**TotalPagedResults** | **int64** |  | 
**TotalPagedResultsPolicy** | **string** |  | 

## Methods

### NewGetAllVariables200Response

`func NewGetAllVariables200Response(pagedResultsCookie NullableString, remainingPagedResults int64, result []EsvVariableResponseV2, resultCount int64, totalPagedResults int64, totalPagedResultsPolicy string, ) *GetAllVariables200Response`

NewGetAllVariables200Response instantiates a new GetAllVariables200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllVariables200ResponseWithDefaults

`func NewGetAllVariables200ResponseWithDefaults() *GetAllVariables200Response`

NewGetAllVariables200ResponseWithDefaults instantiates a new GetAllVariables200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagedResultsCookie

`func (o *GetAllVariables200Response) GetPagedResultsCookie() string`

GetPagedResultsCookie returns the PagedResultsCookie field if non-nil, zero value otherwise.

### GetPagedResultsCookieOk

`func (o *GetAllVariables200Response) GetPagedResultsCookieOk() (*string, bool)`

GetPagedResultsCookieOk returns a tuple with the PagedResultsCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagedResultsCookie

`func (o *GetAllVariables200Response) SetPagedResultsCookie(v string)`

SetPagedResultsCookie sets PagedResultsCookie field to given value.


### SetPagedResultsCookieNil

`func (o *GetAllVariables200Response) SetPagedResultsCookieNil(b bool)`

 SetPagedResultsCookieNil sets the value for PagedResultsCookie to be an explicit nil

### UnsetPagedResultsCookie
`func (o *GetAllVariables200Response) UnsetPagedResultsCookie()`

UnsetPagedResultsCookie ensures that no value is present for PagedResultsCookie, not even an explicit nil
### GetRemainingPagedResults

`func (o *GetAllVariables200Response) GetRemainingPagedResults() int64`

GetRemainingPagedResults returns the RemainingPagedResults field if non-nil, zero value otherwise.

### GetRemainingPagedResultsOk

`func (o *GetAllVariables200Response) GetRemainingPagedResultsOk() (*int64, bool)`

GetRemainingPagedResultsOk returns a tuple with the RemainingPagedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingPagedResults

`func (o *GetAllVariables200Response) SetRemainingPagedResults(v int64)`

SetRemainingPagedResults sets RemainingPagedResults field to given value.


### GetResult

`func (o *GetAllVariables200Response) GetResult() []EsvVariableResponseV2`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAllVariables200Response) GetResultOk() (*[]EsvVariableResponseV2, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAllVariables200Response) SetResult(v []EsvVariableResponseV2)`

SetResult sets Result field to given value.


### GetResultCount

`func (o *GetAllVariables200Response) GetResultCount() int64`

GetResultCount returns the ResultCount field if non-nil, zero value otherwise.

### GetResultCountOk

`func (o *GetAllVariables200Response) GetResultCountOk() (*int64, bool)`

GetResultCountOk returns a tuple with the ResultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCount

`func (o *GetAllVariables200Response) SetResultCount(v int64)`

SetResultCount sets ResultCount field to given value.


### GetTotalPagedResults

`func (o *GetAllVariables200Response) GetTotalPagedResults() int64`

GetTotalPagedResults returns the TotalPagedResults field if non-nil, zero value otherwise.

### GetTotalPagedResultsOk

`func (o *GetAllVariables200Response) GetTotalPagedResultsOk() (*int64, bool)`

GetTotalPagedResultsOk returns a tuple with the TotalPagedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPagedResults

`func (o *GetAllVariables200Response) SetTotalPagedResults(v int64)`

SetTotalPagedResults sets TotalPagedResults field to given value.


### GetTotalPagedResultsPolicy

`func (o *GetAllVariables200Response) GetTotalPagedResultsPolicy() string`

GetTotalPagedResultsPolicy returns the TotalPagedResultsPolicy field if non-nil, zero value otherwise.

### GetTotalPagedResultsPolicyOk

`func (o *GetAllVariables200Response) GetTotalPagedResultsPolicyOk() (*string, bool)`

GetTotalPagedResultsPolicyOk returns a tuple with the TotalPagedResultsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPagedResultsPolicy

`func (o *GetAllVariables200Response) SetTotalPagedResultsPolicy(v string)`

SetTotalPagedResultsPolicy sets TotalPagedResultsPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


