# PromotionProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockingError** | Pointer to **bool** |  | [optional] [default to false]
**EncryptedSecrets** | Pointer to **[]string** |  | [optional] 
**GlobalLock** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MissingESVs** | Pointer to **[]string** |  | [optional] 
**PromotionId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Either promotion or rollback depending on whether the report was for a config promotion or a config rollback | [optional] 

## Methods

### NewPromotionProgress

`func NewPromotionProgress() *PromotionProgress`

NewPromotionProgress instantiates a new PromotionProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionProgressWithDefaults

`func NewPromotionProgressWithDefaults() *PromotionProgress`

NewPromotionProgressWithDefaults instantiates a new PromotionProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockingError

`func (o *PromotionProgress) GetBlockingError() bool`

GetBlockingError returns the BlockingError field if non-nil, zero value otherwise.

### GetBlockingErrorOk

`func (o *PromotionProgress) GetBlockingErrorOk() (*bool, bool)`

GetBlockingErrorOk returns a tuple with the BlockingError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingError

`func (o *PromotionProgress) SetBlockingError(v bool)`

SetBlockingError sets BlockingError field to given value.

### HasBlockingError

`func (o *PromotionProgress) HasBlockingError() bool`

HasBlockingError returns a boolean if a field has been set.

### GetEncryptedSecrets

`func (o *PromotionProgress) GetEncryptedSecrets() []string`

GetEncryptedSecrets returns the EncryptedSecrets field if non-nil, zero value otherwise.

### GetEncryptedSecretsOk

`func (o *PromotionProgress) GetEncryptedSecretsOk() (*[]string, bool)`

GetEncryptedSecretsOk returns a tuple with the EncryptedSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecrets

`func (o *PromotionProgress) SetEncryptedSecrets(v []string)`

SetEncryptedSecrets sets EncryptedSecrets field to given value.

### HasEncryptedSecrets

`func (o *PromotionProgress) HasEncryptedSecrets() bool`

HasEncryptedSecrets returns a boolean if a field has been set.

### GetGlobalLock

`func (o *PromotionProgress) GetGlobalLock() string`

GetGlobalLock returns the GlobalLock field if non-nil, zero value otherwise.

### GetGlobalLockOk

`func (o *PromotionProgress) GetGlobalLockOk() (*string, bool)`

GetGlobalLockOk returns a tuple with the GlobalLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLock

`func (o *PromotionProgress) SetGlobalLock(v string)`

SetGlobalLock sets GlobalLock field to given value.

### HasGlobalLock

`func (o *PromotionProgress) HasGlobalLock() bool`

HasGlobalLock returns a boolean if a field has been set.

### GetMessage

`func (o *PromotionProgress) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PromotionProgress) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PromotionProgress) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PromotionProgress) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMissingESVs

`func (o *PromotionProgress) GetMissingESVs() []string`

GetMissingESVs returns the MissingESVs field if non-nil, zero value otherwise.

### GetMissingESVsOk

`func (o *PromotionProgress) GetMissingESVsOk() (*[]string, bool)`

GetMissingESVsOk returns a tuple with the MissingESVs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingESVs

`func (o *PromotionProgress) SetMissingESVs(v []string)`

SetMissingESVs sets MissingESVs field to given value.

### HasMissingESVs

`func (o *PromotionProgress) HasMissingESVs() bool`

HasMissingESVs returns a boolean if a field has been set.

### GetPromotionId

`func (o *PromotionProgress) GetPromotionId() string`

GetPromotionId returns the PromotionId field if non-nil, zero value otherwise.

### GetPromotionIdOk

`func (o *PromotionProgress) GetPromotionIdOk() (*string, bool)`

GetPromotionIdOk returns a tuple with the PromotionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionId

`func (o *PromotionProgress) SetPromotionId(v string)`

SetPromotionId sets PromotionId field to given value.

### HasPromotionId

`func (o *PromotionProgress) HasPromotionId() bool`

HasPromotionId returns a boolean if a field has been set.

### GetStatus

`func (o *PromotionProgress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PromotionProgress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PromotionProgress) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PromotionProgress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeStamp

`func (o *PromotionProgress) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *PromotionProgress) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *PromotionProgress) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *PromotionProgress) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetType

`func (o *PromotionProgress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromotionProgress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromotionProgress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromotionProgress) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


