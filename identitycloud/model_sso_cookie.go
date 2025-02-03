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

// checks if the SSOCookie type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSOCookie{}

// SSOCookie struct for SSOCookie
type SSOCookie struct {
	Name string `json:"name"`
}

// NewSSOCookie instantiates a new SSOCookie object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSOCookie(name string) *SSOCookie {
	this := SSOCookie{}
	this.Name = name
	return &this
}

// NewSSOCookieWithDefaults instantiates a new SSOCookie object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSOCookieWithDefaults() *SSOCookie {
	this := SSOCookie{}
	return &this
}

// GetName returns the Name field value
func (o *SSOCookie) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SSOCookie) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SSOCookie) SetName(v string) {
	o.Name = v
}

func (o SSOCookie) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSOCookie) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableSSOCookie struct {
	value *SSOCookie
	isSet bool
}

func (v NullableSSOCookie) Get() *SSOCookie {
	return v.value
}

func (v *NullableSSOCookie) Set(val *SSOCookie) {
	v.value = val
	v.isSet = true
}

func (v NullableSSOCookie) IsSet() bool {
	return v.isSet
}

func (v *NullableSSOCookie) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSOCookie(val *SSOCookie) *NullableSSOCookie {
	return &NullableSSOCookie{value: val, isSet: true}
}

func (v NullableSSOCookie) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSOCookie) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
