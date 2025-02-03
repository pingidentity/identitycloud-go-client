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

// checks if the ContentSecurityPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSecurityPolicy{}

// ContentSecurityPolicy struct for ContentSecurityPolicy
type ContentSecurityPolicy struct {
	Active     *bool                `json:"active,omitempty"`
	Directives *map[string][]string `json:"directives,omitempty"`
}

// NewContentSecurityPolicy instantiates a new ContentSecurityPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSecurityPolicy() *ContentSecurityPolicy {
	this := ContentSecurityPolicy{}
	return &this
}

// NewContentSecurityPolicyWithDefaults instantiates a new ContentSecurityPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSecurityPolicyWithDefaults() *ContentSecurityPolicy {
	this := ContentSecurityPolicy{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ContentSecurityPolicy) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSecurityPolicy) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ContentSecurityPolicy) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ContentSecurityPolicy) SetActive(v bool) {
	o.Active = &v
}

// GetDirectives returns the Directives field value if set, zero value otherwise.
func (o *ContentSecurityPolicy) GetDirectives() map[string][]string {
	if o == nil || IsNil(o.Directives) {
		var ret map[string][]string
		return ret
	}
	return *o.Directives
}

// GetDirectivesOk returns a tuple with the Directives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSecurityPolicy) GetDirectivesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Directives) {
		return nil, false
	}
	return o.Directives, true
}

// HasDirectives returns a boolean if a field has been set.
func (o *ContentSecurityPolicy) HasDirectives() bool {
	if o != nil && !IsNil(o.Directives) {
		return true
	}

	return false
}

// SetDirectives gets a reference to the given map[string][]string and assigns it to the Directives field.
func (o *ContentSecurityPolicy) SetDirectives(v map[string][]string) {
	o.Directives = &v
}

func (o ContentSecurityPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSecurityPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Directives) {
		toSerialize["directives"] = o.Directives
	}
	return toSerialize, nil
}

type NullableContentSecurityPolicy struct {
	value *ContentSecurityPolicy
	isSet bool
}

func (v NullableContentSecurityPolicy) Get() *ContentSecurityPolicy {
	return v.value
}

func (v *NullableContentSecurityPolicy) Set(val *ContentSecurityPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSecurityPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSecurityPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSecurityPolicy(val *ContentSecurityPolicy) *NullableContentSecurityPolicy {
	return &NullableContentSecurityPolicy{value: val, isSet: true}
}

func (v NullableContentSecurityPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSecurityPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
