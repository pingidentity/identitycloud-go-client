# EsvVariableCreateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ExpressionType** | Pointer to **string** |  | [optional] 
**ValueBase64** | **string** |  | 

## Methods

### NewEsvVariableCreateUpdateRequest

`func NewEsvVariableCreateUpdateRequest(valueBase64 string, ) *EsvVariableCreateUpdateRequest`

NewEsvVariableCreateUpdateRequest instantiates a new EsvVariableCreateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsvVariableCreateUpdateRequestWithDefaults

`func NewEsvVariableCreateUpdateRequestWithDefaults() *EsvVariableCreateUpdateRequest`

NewEsvVariableCreateUpdateRequestWithDefaults instantiates a new EsvVariableCreateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EsvVariableCreateUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsvVariableCreateUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsvVariableCreateUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsvVariableCreateUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpressionType

`func (o *EsvVariableCreateUpdateRequest) GetExpressionType() string`

GetExpressionType returns the ExpressionType field if non-nil, zero value otherwise.

### GetExpressionTypeOk

`func (o *EsvVariableCreateUpdateRequest) GetExpressionTypeOk() (*string, bool)`

GetExpressionTypeOk returns a tuple with the ExpressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionType

`func (o *EsvVariableCreateUpdateRequest) SetExpressionType(v string)`

SetExpressionType sets ExpressionType field to given value.

### HasExpressionType

`func (o *EsvVariableCreateUpdateRequest) HasExpressionType() bool`

HasExpressionType returns a boolean if a field has been set.

### GetValueBase64

`func (o *EsvVariableCreateUpdateRequest) GetValueBase64() string`

GetValueBase64 returns the ValueBase64 field if non-nil, zero value otherwise.

### GetValueBase64Ok

`func (o *EsvVariableCreateUpdateRequest) GetValueBase64Ok() (*string, bool)`

GetValueBase64Ok returns a tuple with the ValueBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBase64

`func (o *EsvVariableCreateUpdateRequest) SetValueBase64(v string)`

SetValueBase64 sets ValueBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


