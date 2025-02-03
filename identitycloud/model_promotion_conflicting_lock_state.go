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

// checks if the PromotionConflictingLockState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionConflictingLockState{}

// PromotionConflictingLockState struct for PromotionConflictingLockState
type PromotionConflictingLockState struct {
	Message *string `json:"message,omitempty"`
}

// NewPromotionConflictingLockState instantiates a new PromotionConflictingLockState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionConflictingLockState() *PromotionConflictingLockState {
	this := PromotionConflictingLockState{}
	return &this
}

// NewPromotionConflictingLockStateWithDefaults instantiates a new PromotionConflictingLockState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionConflictingLockStateWithDefaults() *PromotionConflictingLockState {
	this := PromotionConflictingLockState{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PromotionConflictingLockState) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionConflictingLockState) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PromotionConflictingLockState) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PromotionConflictingLockState) SetMessage(v string) {
	o.Message = &v
}

func (o PromotionConflictingLockState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionConflictingLockState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePromotionConflictingLockState struct {
	value *PromotionConflictingLockState
	isSet bool
}

func (v NullablePromotionConflictingLockState) Get() *PromotionConflictingLockState {
	return v.value
}

func (v *NullablePromotionConflictingLockState) Set(val *PromotionConflictingLockState) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionConflictingLockState) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionConflictingLockState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionConflictingLockState(val *PromotionConflictingLockState) *NullablePromotionConflictingLockState {
	return &NullablePromotionConflictingLockState{value: val, isSet: true}
}

func (v NullablePromotionConflictingLockState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionConflictingLockState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
