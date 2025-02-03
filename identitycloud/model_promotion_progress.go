// Copyright © 2025 Ping Identity Corporation

/*
PingOne Advanced Identity Cloud API

## Introduction The PingOne Advanced Identity Cloud REST API lets you manage your Advanced Identity Cloud tenants. The API exposes access management and identity management endpoints, with additional endpoints specific to Advanced Identity Cloud tenant environments.<br /><br /> We are now publishing the API spec in OpenAPI 3.0. For the legacy Swagger 2.0 spec, please download [swagger.yaml](swagger.yaml), but note that it may not contain all new functionality.<br /><br /> For full PingOne Advanced Identity Cloud documentation, please visit [the docs website](https://docs.pingidentity.com/pingoneaic/latest/). ## Authenticating to the API The PingOne Advanced Identity Cloud REST API has two different authentication methods:   - [API key and secret](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-with-api-key-and-secret.html): used for tenant read-only operations  - [Access token](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-with-access-token.html): used for access management operations, identity management operations or tenant write operations  For a summary of how to use these authentication methods, refer to [Authenticate to Advanced Identity Cloud REST API](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-overview.html).

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package identitycloud

import (
	"encoding/json"
)

// checks if the PromotionProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionProgress{}

// PromotionProgress struct for PromotionProgress
type PromotionProgress struct {
	BlockingError    *bool    `json:"blockingError,omitempty"`
	EncryptedSecrets []string `json:"encryptedSecrets,omitempty"`
	GlobalLock       *string  `json:"globalLock,omitempty"`
	Message          *string  `json:"message,omitempty"`
	MissingESVs      []string `json:"missingESVs,omitempty"`
	PromotionId      *string  `json:"promotionId,omitempty"`
	Status           *string  `json:"status,omitempty"`
	TimeStamp        *string  `json:"timeStamp,omitempty"`
	// Either promotion or rollback depending on whether the report was for a config promotion or a config rollback
	Type *string `json:"type,omitempty"`
}

// NewPromotionProgress instantiates a new PromotionProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionProgress() *PromotionProgress {
	this := PromotionProgress{}
	var blockingError bool = false
	this.BlockingError = &blockingError
	return &this
}

// NewPromotionProgressWithDefaults instantiates a new PromotionProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionProgressWithDefaults() *PromotionProgress {
	this := PromotionProgress{}
	var blockingError bool = false
	this.BlockingError = &blockingError
	return &this
}

// GetBlockingError returns the BlockingError field value if set, zero value otherwise.
func (o *PromotionProgress) GetBlockingError() bool {
	if o == nil || IsNil(o.BlockingError) {
		var ret bool
		return ret
	}
	return *o.BlockingError
}

// GetBlockingErrorOk returns a tuple with the BlockingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetBlockingErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockingError) {
		return nil, false
	}
	return o.BlockingError, true
}

// HasBlockingError returns a boolean if a field has been set.
func (o *PromotionProgress) HasBlockingError() bool {
	if o != nil && !IsNil(o.BlockingError) {
		return true
	}

	return false
}

// SetBlockingError gets a reference to the given bool and assigns it to the BlockingError field.
func (o *PromotionProgress) SetBlockingError(v bool) {
	o.BlockingError = &v
}

// GetEncryptedSecrets returns the EncryptedSecrets field value if set, zero value otherwise.
func (o *PromotionProgress) GetEncryptedSecrets() []string {
	if o == nil || IsNil(o.EncryptedSecrets) {
		var ret []string
		return ret
	}
	return o.EncryptedSecrets
}

// GetEncryptedSecretsOk returns a tuple with the EncryptedSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetEncryptedSecretsOk() ([]string, bool) {
	if o == nil || IsNil(o.EncryptedSecrets) {
		return nil, false
	}
	return o.EncryptedSecrets, true
}

// HasEncryptedSecrets returns a boolean if a field has been set.
func (o *PromotionProgress) HasEncryptedSecrets() bool {
	if o != nil && !IsNil(o.EncryptedSecrets) {
		return true
	}

	return false
}

// SetEncryptedSecrets gets a reference to the given []string and assigns it to the EncryptedSecrets field.
func (o *PromotionProgress) SetEncryptedSecrets(v []string) {
	o.EncryptedSecrets = v
}

// GetGlobalLock returns the GlobalLock field value if set, zero value otherwise.
func (o *PromotionProgress) GetGlobalLock() string {
	if o == nil || IsNil(o.GlobalLock) {
		var ret string
		return ret
	}
	return *o.GlobalLock
}

// GetGlobalLockOk returns a tuple with the GlobalLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetGlobalLockOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalLock) {
		return nil, false
	}
	return o.GlobalLock, true
}

// HasGlobalLock returns a boolean if a field has been set.
func (o *PromotionProgress) HasGlobalLock() bool {
	if o != nil && !IsNil(o.GlobalLock) {
		return true
	}

	return false
}

// SetGlobalLock gets a reference to the given string and assigns it to the GlobalLock field.
func (o *PromotionProgress) SetGlobalLock(v string) {
	o.GlobalLock = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PromotionProgress) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PromotionProgress) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PromotionProgress) SetMessage(v string) {
	o.Message = &v
}

// GetMissingESVs returns the MissingESVs field value if set, zero value otherwise.
func (o *PromotionProgress) GetMissingESVs() []string {
	if o == nil || IsNil(o.MissingESVs) {
		var ret []string
		return ret
	}
	return o.MissingESVs
}

// GetMissingESVsOk returns a tuple with the MissingESVs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetMissingESVsOk() ([]string, bool) {
	if o == nil || IsNil(o.MissingESVs) {
		return nil, false
	}
	return o.MissingESVs, true
}

// HasMissingESVs returns a boolean if a field has been set.
func (o *PromotionProgress) HasMissingESVs() bool {
	if o != nil && !IsNil(o.MissingESVs) {
		return true
	}

	return false
}

// SetMissingESVs gets a reference to the given []string and assigns it to the MissingESVs field.
func (o *PromotionProgress) SetMissingESVs(v []string) {
	o.MissingESVs = v
}

// GetPromotionId returns the PromotionId field value if set, zero value otherwise.
func (o *PromotionProgress) GetPromotionId() string {
	if o == nil || IsNil(o.PromotionId) {
		var ret string
		return ret
	}
	return *o.PromotionId
}

// GetPromotionIdOk returns a tuple with the PromotionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetPromotionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionId) {
		return nil, false
	}
	return o.PromotionId, true
}

// HasPromotionId returns a boolean if a field has been set.
func (o *PromotionProgress) HasPromotionId() bool {
	if o != nil && !IsNil(o.PromotionId) {
		return true
	}

	return false
}

// SetPromotionId gets a reference to the given string and assigns it to the PromotionId field.
func (o *PromotionProgress) SetPromotionId(v string) {
	o.PromotionId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PromotionProgress) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PromotionProgress) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PromotionProgress) SetStatus(v string) {
	o.Status = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *PromotionProgress) GetTimeStamp() string {
	if o == nil || IsNil(o.TimeStamp) {
		var ret string
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetTimeStampOk() (*string, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *PromotionProgress) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given string and assigns it to the TimeStamp field.
func (o *PromotionProgress) SetTimeStamp(v string) {
	o.TimeStamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PromotionProgress) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionProgress) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PromotionProgress) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PromotionProgress) SetType(v string) {
	o.Type = &v
}

func (o PromotionProgress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockingError) {
		toSerialize["blockingError"] = o.BlockingError
	}
	if !IsNil(o.EncryptedSecrets) {
		toSerialize["encryptedSecrets"] = o.EncryptedSecrets
	}
	if !IsNil(o.GlobalLock) {
		toSerialize["globalLock"] = o.GlobalLock
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.MissingESVs) {
		toSerialize["missingESVs"] = o.MissingESVs
	}
	if !IsNil(o.PromotionId) {
		toSerialize["promotionId"] = o.PromotionId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePromotionProgress struct {
	value *PromotionProgress
	isSet bool
}

func (v NullablePromotionProgress) Get() *PromotionProgress {
	return v.value
}

func (v *NullablePromotionProgress) Set(val *PromotionProgress) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionProgress) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionProgress(val *PromotionProgress) *NullablePromotionProgress {
	return &NullablePromotionProgress{value: val, isSet: true}
}

func (v NullablePromotionProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
