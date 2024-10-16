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

// checks if the EsvRestartStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsvRestartStatus{}

// EsvRestartStatus struct for EsvRestartStatus
type EsvRestartStatus struct {
	RestartStatus string `json:"restartStatus"`
}

// NewEsvRestartStatus instantiates a new EsvRestartStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsvRestartStatus(restartStatus string) *EsvRestartStatus {
	this := EsvRestartStatus{}
	this.RestartStatus = restartStatus
	return &this
}

// NewEsvRestartStatusWithDefaults instantiates a new EsvRestartStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsvRestartStatusWithDefaults() *EsvRestartStatus {
	this := EsvRestartStatus{}
	return &this
}

// GetRestartStatus returns the RestartStatus field value
func (o *EsvRestartStatus) GetRestartStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestartStatus
}

// GetRestartStatusOk returns a tuple with the RestartStatus field value
// and a boolean to check if the value has been set.
func (o *EsvRestartStatus) GetRestartStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestartStatus, true
}

// SetRestartStatus sets field value
func (o *EsvRestartStatus) SetRestartStatus(v string) {
	o.RestartStatus = v
}

func (o EsvRestartStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsvRestartStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["restartStatus"] = o.RestartStatus
	return toSerialize, nil
}

type NullableEsvRestartStatus struct {
	value *EsvRestartStatus
	isSet bool
}

func (v NullableEsvRestartStatus) Get() *EsvRestartStatus {
	return v.value
}

func (v *NullableEsvRestartStatus) Set(val *EsvRestartStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEsvRestartStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEsvRestartStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsvRestartStatus(val *EsvRestartStatus) *NullableEsvRestartStatus {
	return &NullableEsvRestartStatus{value: val, isSet: true}
}

func (v NullableEsvRestartStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsvRestartStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}