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

// checks if the PromotionReportContentConfigChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionReportContentConfigChange{}

// PromotionReportContentConfigChange struct for PromotionReportContentConfigChange
type PromotionReportContentConfigChange struct {
	Added    []PromotionCIMeta `json:"added,omitempty"`
	Deleted  []PromotionCIMeta `json:"deleted,omitempty"`
	Modified []PromotionCIMeta `json:"modified,omitempty"`
}

// NewPromotionReportContentConfigChange instantiates a new PromotionReportContentConfigChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionReportContentConfigChange() *PromotionReportContentConfigChange {
	this := PromotionReportContentConfigChange{}
	return &this
}

// NewPromotionReportContentConfigChangeWithDefaults instantiates a new PromotionReportContentConfigChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionReportContentConfigChangeWithDefaults() *PromotionReportContentConfigChange {
	this := PromotionReportContentConfigChange{}
	return &this
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *PromotionReportContentConfigChange) GetAdded() []PromotionCIMeta {
	if o == nil || IsNil(o.Added) {
		var ret []PromotionCIMeta
		return ret
	}
	return o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportContentConfigChange) GetAddedOk() ([]PromotionCIMeta, bool) {
	if o == nil || IsNil(o.Added) {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *PromotionReportContentConfigChange) HasAdded() bool {
	if o != nil && !IsNil(o.Added) {
		return true
	}

	return false
}

// SetAdded gets a reference to the given []PromotionCIMeta and assigns it to the Added field.
func (o *PromotionReportContentConfigChange) SetAdded(v []PromotionCIMeta) {
	o.Added = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *PromotionReportContentConfigChange) GetDeleted() []PromotionCIMeta {
	if o == nil || IsNil(o.Deleted) {
		var ret []PromotionCIMeta
		return ret
	}
	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportContentConfigChange) GetDeletedOk() ([]PromotionCIMeta, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *PromotionReportContentConfigChange) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given []PromotionCIMeta and assigns it to the Deleted field.
func (o *PromotionReportContentConfigChange) SetDeleted(v []PromotionCIMeta) {
	o.Deleted = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *PromotionReportContentConfigChange) GetModified() []PromotionCIMeta {
	if o == nil || IsNil(o.Modified) {
		var ret []PromotionCIMeta
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportContentConfigChange) GetModifiedOk() ([]PromotionCIMeta, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *PromotionReportContentConfigChange) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given []PromotionCIMeta and assigns it to the Modified field.
func (o *PromotionReportContentConfigChange) SetModified(v []PromotionCIMeta) {
	o.Modified = v
}

func (o PromotionReportContentConfigChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionReportContentConfigChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Added) {
		toSerialize["added"] = o.Added
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	return toSerialize, nil
}

type NullablePromotionReportContentConfigChange struct {
	value *PromotionReportContentConfigChange
	isSet bool
}

func (v NullablePromotionReportContentConfigChange) Get() *PromotionReportContentConfigChange {
	return v.value
}

func (v *NullablePromotionReportContentConfigChange) Set(val *PromotionReportContentConfigChange) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionReportContentConfigChange) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionReportContentConfigChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionReportContentConfigChange(val *PromotionReportContentConfigChange) *NullablePromotionReportContentConfigChange {
	return &NullablePromotionReportContentConfigChange{value: val, isSet: true}
}

func (v NullablePromotionReportContentConfigChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionReportContentConfigChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
