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

// checks if the PromotionLockState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionLockState{}

// PromotionLockState struct for PromotionLockState
type PromotionLockState struct {
	Description *string                     `json:"description,omitempty"`
	LowerEnv    *PromotionLockStateLowerEnv `json:"lowerEnv,omitempty"`
	PromotionId *string                     `json:"promotionId,omitempty"`
	Result      *string                     `json:"result,omitempty"`
	UpperEnv    *PromotionLockStateLowerEnv `json:"upperEnv,omitempty"`
}

// NewPromotionLockState instantiates a new PromotionLockState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionLockState() *PromotionLockState {
	this := PromotionLockState{}
	return &this
}

// NewPromotionLockStateWithDefaults instantiates a new PromotionLockState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionLockStateWithDefaults() *PromotionLockState {
	this := PromotionLockState{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PromotionLockState) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionLockState) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PromotionLockState) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PromotionLockState) SetDescription(v string) {
	o.Description = &v
}

// GetLowerEnv returns the LowerEnv field value if set, zero value otherwise.
func (o *PromotionLockState) GetLowerEnv() PromotionLockStateLowerEnv {
	if o == nil || IsNil(o.LowerEnv) {
		var ret PromotionLockStateLowerEnv
		return ret
	}
	return *o.LowerEnv
}

// GetLowerEnvOk returns a tuple with the LowerEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionLockState) GetLowerEnvOk() (*PromotionLockStateLowerEnv, bool) {
	if o == nil || IsNil(o.LowerEnv) {
		return nil, false
	}
	return o.LowerEnv, true
}

// HasLowerEnv returns a boolean if a field has been set.
func (o *PromotionLockState) HasLowerEnv() bool {
	if o != nil && !IsNil(o.LowerEnv) {
		return true
	}

	return false
}

// SetLowerEnv gets a reference to the given PromotionLockStateLowerEnv and assigns it to the LowerEnv field.
func (o *PromotionLockState) SetLowerEnv(v PromotionLockStateLowerEnv) {
	o.LowerEnv = &v
}

// GetPromotionId returns the PromotionId field value if set, zero value otherwise.
func (o *PromotionLockState) GetPromotionId() string {
	if o == nil || IsNil(o.PromotionId) {
		var ret string
		return ret
	}
	return *o.PromotionId
}

// GetPromotionIdOk returns a tuple with the PromotionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionLockState) GetPromotionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionId) {
		return nil, false
	}
	return o.PromotionId, true
}

// HasPromotionId returns a boolean if a field has been set.
func (o *PromotionLockState) HasPromotionId() bool {
	if o != nil && !IsNil(o.PromotionId) {
		return true
	}

	return false
}

// SetPromotionId gets a reference to the given string and assigns it to the PromotionId field.
func (o *PromotionLockState) SetPromotionId(v string) {
	o.PromotionId = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *PromotionLockState) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionLockState) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PromotionLockState) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *PromotionLockState) SetResult(v string) {
	o.Result = &v
}

// GetUpperEnv returns the UpperEnv field value if set, zero value otherwise.
func (o *PromotionLockState) GetUpperEnv() PromotionLockStateLowerEnv {
	if o == nil || IsNil(o.UpperEnv) {
		var ret PromotionLockStateLowerEnv
		return ret
	}
	return *o.UpperEnv
}

// GetUpperEnvOk returns a tuple with the UpperEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionLockState) GetUpperEnvOk() (*PromotionLockStateLowerEnv, bool) {
	if o == nil || IsNil(o.UpperEnv) {
		return nil, false
	}
	return o.UpperEnv, true
}

// HasUpperEnv returns a boolean if a field has been set.
func (o *PromotionLockState) HasUpperEnv() bool {
	if o != nil && !IsNil(o.UpperEnv) {
		return true
	}

	return false
}

// SetUpperEnv gets a reference to the given PromotionLockStateLowerEnv and assigns it to the UpperEnv field.
func (o *PromotionLockState) SetUpperEnv(v PromotionLockStateLowerEnv) {
	o.UpperEnv = &v
}

func (o PromotionLockState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionLockState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LowerEnv) {
		toSerialize["lowerEnv"] = o.LowerEnv
	}
	if !IsNil(o.PromotionId) {
		toSerialize["promotionId"] = o.PromotionId
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.UpperEnv) {
		toSerialize["upperEnv"] = o.UpperEnv
	}
	return toSerialize, nil
}

type NullablePromotionLockState struct {
	value *PromotionLockState
	isSet bool
}

func (v NullablePromotionLockState) Get() *PromotionLockState {
	return v.value
}

func (v *NullablePromotionLockState) Set(val *PromotionLockState) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionLockState) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionLockState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionLockState(val *PromotionLockState) *NullablePromotionLockState {
	return &NullablePromotionLockState{value: val, isSet: true}
}

func (v NullablePromotionLockState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionLockState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
