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

// checks if the RollbackStart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollbackStart{}

// RollbackStart struct for RollbackStart
type RollbackStart struct {
	Result *string `json:"result,omitempty"`
}

// NewRollbackStart instantiates a new RollbackStart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackStart() *RollbackStart {
	this := RollbackStart{}
	return &this
}

// NewRollbackStartWithDefaults instantiates a new RollbackStart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackStartWithDefaults() *RollbackStart {
	this := RollbackStart{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RollbackStart) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackStart) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RollbackStart) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *RollbackStart) SetResult(v string) {
	o.Result = &v
}

func (o RollbackStart) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollbackStart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableRollbackStart struct {
	value *RollbackStart
	isSet bool
}

func (v NullableRollbackStart) Get() *RollbackStart {
	return v.value
}

func (v *NullableRollbackStart) Set(val *RollbackStart) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackStart) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackStart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackStart(val *RollbackStart) *NullableRollbackStart {
	return &NullableRollbackStart{value: val, isSet: true}
}

func (v NullableRollbackStart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackStart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
