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

// checks if the PromotionReportMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionReportMeta{}

// PromotionReportMeta struct for PromotionReportMeta
type PromotionReportMeta struct {
	CreatedDate *string `json:"createdDate,omitempty"`
	DryRun      *bool   `json:"dryRun,omitempty"`
	Promoter    *string `json:"promoter,omitempty"`
	PromotionId *string `json:"promotionId,omitempty"`
	ReportId    *string `json:"reportId,omitempty"`
	// Either promotion or rollback depending on whether the report was for a config promotion or a config rollback
	Type *string `json:"type,omitempty"`
}

// NewPromotionReportMeta instantiates a new PromotionReportMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionReportMeta() *PromotionReportMeta {
	this := PromotionReportMeta{}
	return &this
}

// NewPromotionReportMetaWithDefaults instantiates a new PromotionReportMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionReportMetaWithDefaults() *PromotionReportMeta {
	this := PromotionReportMeta{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *PromotionReportMeta) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportMeta) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *PromotionReportMeta) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *PromotionReportMeta) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *PromotionReportMeta) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportMeta) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *PromotionReportMeta) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *PromotionReportMeta) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPromoter returns the Promoter field value if set, zero value otherwise.
func (o *PromotionReportMeta) GetPromoter() string {
	if o == nil || IsNil(o.Promoter) {
		var ret string
		return ret
	}
	return *o.Promoter
}

// GetPromoterOk returns a tuple with the Promoter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportMeta) GetPromoterOk() (*string, bool) {
	if o == nil || IsNil(o.Promoter) {
		return nil, false
	}
	return o.Promoter, true
}

// HasPromoter returns a boolean if a field has been set.
func (o *PromotionReportMeta) HasPromoter() bool {
	if o != nil && !IsNil(o.Promoter) {
		return true
	}

	return false
}

// SetPromoter gets a reference to the given string and assigns it to the Promoter field.
func (o *PromotionReportMeta) SetPromoter(v string) {
	o.Promoter = &v
}

// GetPromotionId returns the PromotionId field value if set, zero value otherwise.
func (o *PromotionReportMeta) GetPromotionId() string {
	if o == nil || IsNil(o.PromotionId) {
		var ret string
		return ret
	}
	return *o.PromotionId
}

// GetPromotionIdOk returns a tuple with the PromotionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportMeta) GetPromotionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionId) {
		return nil, false
	}
	return o.PromotionId, true
}

// HasPromotionId returns a boolean if a field has been set.
func (o *PromotionReportMeta) HasPromotionId() bool {
	if o != nil && !IsNil(o.PromotionId) {
		return true
	}

	return false
}

// SetPromotionId gets a reference to the given string and assigns it to the PromotionId field.
func (o *PromotionReportMeta) SetPromotionId(v string) {
	o.PromotionId = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *PromotionReportMeta) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportMeta) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *PromotionReportMeta) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *PromotionReportMeta) SetReportId(v string) {
	o.ReportId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PromotionReportMeta) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionReportMeta) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PromotionReportMeta) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PromotionReportMeta) SetType(v string) {
	o.Type = &v
}

func (o PromotionReportMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionReportMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	if !IsNil(o.Promoter) {
		toSerialize["promoter"] = o.Promoter
	}
	if !IsNil(o.PromotionId) {
		toSerialize["promotionId"] = o.PromotionId
	}
	if !IsNil(o.ReportId) {
		toSerialize["reportId"] = o.ReportId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePromotionReportMeta struct {
	value *PromotionReportMeta
	isSet bool
}

func (v NullablePromotionReportMeta) Get() *PromotionReportMeta {
	return v.value
}

func (v *NullablePromotionReportMeta) Set(val *PromotionReportMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionReportMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionReportMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionReportMeta(val *PromotionReportMeta) *NullablePromotionReportMeta {
	return &NullablePromotionReportMeta{value: val, isSet: true}
}

func (v NullablePromotionReportMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionReportMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}