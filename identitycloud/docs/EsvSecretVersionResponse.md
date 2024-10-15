# EsvSecretVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateDate** | **time.Time** |  | 
**Loaded** | **bool** |  | 
**Status** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewEsvSecretVersionResponse

`func NewEsvSecretVersionResponse(createDate time.Time, loaded bool, status string, version string, ) *EsvSecretVersionResponse`

NewEsvSecretVersionResponse instantiates a new EsvSecretVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsvSecretVersionResponseWithDefaults

`func NewEsvSecretVersionResponseWithDefaults() *EsvSecretVersionResponse`

NewEsvSecretVersionResponseWithDefaults instantiates a new EsvSecretVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateDate

`func (o *EsvSecretVersionResponse) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *EsvSecretVersionResponse) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *EsvSecretVersionResponse) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.


### GetLoaded

`func (o *EsvSecretVersionResponse) GetLoaded() bool`

GetLoaded returns the Loaded field if non-nil, zero value otherwise.

### GetLoadedOk

`func (o *EsvSecretVersionResponse) GetLoadedOk() (*bool, bool)`

GetLoadedOk returns a tuple with the Loaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoaded

`func (o *EsvSecretVersionResponse) SetLoaded(v bool)`

SetLoaded sets Loaded field to given value.


### GetStatus

`func (o *EsvSecretVersionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsvSecretVersionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsvSecretVersionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *EsvSecretVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EsvSecretVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EsvSecretVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


