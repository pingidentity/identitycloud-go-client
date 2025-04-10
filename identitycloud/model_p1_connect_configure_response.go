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

// checks if the P1ConnectConfigureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &P1ConnectConfigureResponse{}

// P1ConnectConfigureResponse struct for P1ConnectConfigureResponse
type P1ConnectConfigureResponse struct {
	AdminClientId *string `json:"adminClientId,omitempty"`
}

// NewP1ConnectConfigureResponse instantiates a new P1ConnectConfigureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP1ConnectConfigureResponse() *P1ConnectConfigureResponse {
	this := P1ConnectConfigureResponse{}
	return &this
}

// NewP1ConnectConfigureResponseWithDefaults instantiates a new P1ConnectConfigureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP1ConnectConfigureResponseWithDefaults() *P1ConnectConfigureResponse {
	this := P1ConnectConfigureResponse{}
	return &this
}

// GetAdminClientId returns the AdminClientId field value if set, zero value otherwise.
func (o *P1ConnectConfigureResponse) GetAdminClientId() string {
	if o == nil || IsNil(o.AdminClientId) {
		var ret string
		return ret
	}
	return *o.AdminClientId
}

// GetAdminClientIdOk returns a tuple with the AdminClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ConnectConfigureResponse) GetAdminClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdminClientId) {
		return nil, false
	}
	return o.AdminClientId, true
}

// HasAdminClientId returns a boolean if a field has been set.
func (o *P1ConnectConfigureResponse) HasAdminClientId() bool {
	if o != nil && !IsNil(o.AdminClientId) {
		return true
	}

	return false
}

// SetAdminClientId gets a reference to the given string and assigns it to the AdminClientId field.
func (o *P1ConnectConfigureResponse) SetAdminClientId(v string) {
	o.AdminClientId = &v
}

func (o P1ConnectConfigureResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o P1ConnectConfigureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminClientId) {
		toSerialize["adminClientId"] = o.AdminClientId
	}
	return toSerialize, nil
}

type NullableP1ConnectConfigureResponse struct {
	value *P1ConnectConfigureResponse
	isSet bool
}

func (v NullableP1ConnectConfigureResponse) Get() *P1ConnectConfigureResponse {
	return v.value
}

func (v *NullableP1ConnectConfigureResponse) Set(val *P1ConnectConfigureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableP1ConnectConfigureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableP1ConnectConfigureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP1ConnectConfigureResponse(val *P1ConnectConfigureResponse) *NullableP1ConnectConfigureResponse {
	return &NullableP1ConnectConfigureResponse{value: val, isSet: true}
}

func (v NullableP1ConnectConfigureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP1ConnectConfigureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
