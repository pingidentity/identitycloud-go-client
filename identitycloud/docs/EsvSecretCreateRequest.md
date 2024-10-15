# EsvSecretCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Encoding** | **string** |  | 
**UseInPlaceholders** | **bool** |  | 
**ValueBase64** | **string** |  | 

## Methods

### NewEsvSecretCreateRequest

`func NewEsvSecretCreateRequest(encoding string, useInPlaceholders bool, valueBase64 string, ) *EsvSecretCreateRequest`

NewEsvSecretCreateRequest instantiates a new EsvSecretCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsvSecretCreateRequestWithDefaults

`func NewEsvSecretCreateRequestWithDefaults() *EsvSecretCreateRequest`

NewEsvSecretCreateRequestWithDefaults instantiates a new EsvSecretCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EsvSecretCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsvSecretCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsvSecretCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsvSecretCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncoding

`func (o *EsvSecretCreateRequest) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *EsvSecretCreateRequest) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *EsvSecretCreateRequest) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetUseInPlaceholders

`func (o *EsvSecretCreateRequest) GetUseInPlaceholders() bool`

GetUseInPlaceholders returns the UseInPlaceholders field if non-nil, zero value otherwise.

### GetUseInPlaceholdersOk

`func (o *EsvSecretCreateRequest) GetUseInPlaceholdersOk() (*bool, bool)`

GetUseInPlaceholdersOk returns a tuple with the UseInPlaceholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseInPlaceholders

`func (o *EsvSecretCreateRequest) SetUseInPlaceholders(v bool)`

SetUseInPlaceholders sets UseInPlaceholders field to given value.


### GetValueBase64

`func (o *EsvSecretCreateRequest) GetValueBase64() string`

GetValueBase64 returns the ValueBase64 field if non-nil, zero value otherwise.

### GetValueBase64Ok

`func (o *EsvSecretCreateRequest) GetValueBase64Ok() (*string, bool)`

GetValueBase64Ok returns a tuple with the ValueBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBase64

`func (o *EsvSecretCreateRequest) SetValueBase64(v string)`

SetValueBase64 sets ValueBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


