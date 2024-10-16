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

// checks if the EsvSecretVersionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsvSecretVersionCreateRequest{}

// EsvSecretVersionCreateRequest struct for EsvSecretVersionCreateRequest
type EsvSecretVersionCreateRequest struct {
	ValueBase64 string `json:"valueBase64"`
}

// NewEsvSecretVersionCreateRequest instantiates a new EsvSecretVersionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsvSecretVersionCreateRequest(valueBase64 string) *EsvSecretVersionCreateRequest {
	this := EsvSecretVersionCreateRequest{}
	this.ValueBase64 = valueBase64
	return &this
}

// NewEsvSecretVersionCreateRequestWithDefaults instantiates a new EsvSecretVersionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsvSecretVersionCreateRequestWithDefaults() *EsvSecretVersionCreateRequest {
	this := EsvSecretVersionCreateRequest{}
	return &this
}

// GetValueBase64 returns the ValueBase64 field value
func (o *EsvSecretVersionCreateRequest) GetValueBase64() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueBase64
}

// GetValueBase64Ok returns a tuple with the ValueBase64 field value
// and a boolean to check if the value has been set.
func (o *EsvSecretVersionCreateRequest) GetValueBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueBase64, true
}

// SetValueBase64 sets field value
func (o *EsvSecretVersionCreateRequest) SetValueBase64(v string) {
	o.ValueBase64 = v
}

func (o EsvSecretVersionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsvSecretVersionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valueBase64"] = o.ValueBase64
	return toSerialize, nil
}

type NullableEsvSecretVersionCreateRequest struct {
	value *EsvSecretVersionCreateRequest
	isSet bool
}

func (v NullableEsvSecretVersionCreateRequest) Get() *EsvSecretVersionCreateRequest {
	return v.value
}

func (v *NullableEsvSecretVersionCreateRequest) Set(val *EsvSecretVersionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEsvSecretVersionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEsvSecretVersionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsvSecretVersionCreateRequest(val *EsvSecretVersionCreateRequest) *NullableEsvSecretVersionCreateRequest {
	return &NullableEsvSecretVersionCreateRequest{value: val, isSet: true}
}

func (v NullableEsvSecretVersionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsvSecretVersionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}