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
	"time"
)

// checks if the EsvVariableResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsvVariableResponse{}

// EsvVariableResponse struct for EsvVariableResponse
type EsvVariableResponse struct {
	Id             string    `json:"_id"`
	Description    string    `json:"description"`
	ExpressionType string    `json:"expressionType"`
	LastChangeDate time.Time `json:"lastChangeDate"`
	LastChangedBy  string    `json:"lastChangedBy"`
	Loaded         bool      `json:"loaded"`
	ValueBase64    string    `json:"valueBase64"`
}

// NewEsvVariableResponse instantiates a new EsvVariableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsvVariableResponse(id string, description string, expressionType string, lastChangeDate time.Time, lastChangedBy string, loaded bool, valueBase64 string) *EsvVariableResponse {
	this := EsvVariableResponse{}
	this.Id = id
	this.Description = description
	this.ExpressionType = expressionType
	this.LastChangeDate = lastChangeDate
	this.LastChangedBy = lastChangedBy
	this.Loaded = loaded
	this.ValueBase64 = valueBase64
	return &this
}

// NewEsvVariableResponseWithDefaults instantiates a new EsvVariableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsvVariableResponseWithDefaults() *EsvVariableResponse {
	this := EsvVariableResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EsvVariableResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EsvVariableResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EsvVariableResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *EsvVariableResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EsvVariableResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EsvVariableResponse) SetDescription(v string) {
	o.Description = v
}

// GetExpressionType returns the ExpressionType field value
func (o *EsvVariableResponse) GetExpressionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value
// and a boolean to check if the value has been set.
func (o *EsvVariableResponse) GetExpressionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpressionType, true
}

// SetExpressionType sets field value
func (o *EsvVariableResponse) SetExpressionType(v string) {
	o.ExpressionType = v
}

// GetLastChangeDate returns the LastChangeDate field value
func (o *EsvVariableResponse) GetLastChangeDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastChangeDate
}

// GetLastChangeDateOk returns a tuple with the LastChangeDate field value
// and a boolean to check if the value has been set.
func (o *EsvVariableResponse) GetLastChangeDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastChangeDate, true
}

// SetLastChangeDate sets field value
func (o *EsvVariableResponse) SetLastChangeDate(v time.Time) {
	o.LastChangeDate = v
}

// GetLastChangedBy returns the LastChangedBy field value
func (o *EsvVariableResponse) GetLastChangedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastChangedBy
}

// GetLastChangedByOk returns a tuple with the LastChangedBy field value
// and a boolean to check if the value has been set.
func (o *EsvVariableResponse) GetLastChangedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastChangedBy, true
}

// SetLastChangedBy sets field value
func (o *EsvVariableResponse) SetLastChangedBy(v string) {
	o.LastChangedBy = v
}

// GetLoaded returns the Loaded field value
func (o *EsvVariableResponse) GetLoaded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Loaded
}

// GetLoadedOk returns a tuple with the Loaded field value
// and a boolean to check if the value has been set.
func (o *EsvVariableResponse) GetLoadedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loaded, true
}

// SetLoaded sets field value
func (o *EsvVariableResponse) SetLoaded(v bool) {
	o.Loaded = v
}

// GetValueBase64 returns the ValueBase64 field value
func (o *EsvVariableResponse) GetValueBase64() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueBase64
}

// GetValueBase64Ok returns a tuple with the ValueBase64 field value
// and a boolean to check if the value has been set.
func (o *EsvVariableResponse) GetValueBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueBase64, true
}

// SetValueBase64 sets field value
func (o *EsvVariableResponse) SetValueBase64(v string) {
	o.ValueBase64 = v
}

func (o EsvVariableResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsvVariableResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["description"] = o.Description
	toSerialize["expressionType"] = o.ExpressionType
	toSerialize["lastChangeDate"] = o.LastChangeDate
	toSerialize["lastChangedBy"] = o.LastChangedBy
	toSerialize["loaded"] = o.Loaded
	toSerialize["valueBase64"] = o.ValueBase64
	return toSerialize, nil
}

type NullableEsvVariableResponse struct {
	value *EsvVariableResponse
	isSet bool
}

func (v NullableEsvVariableResponse) Get() *EsvVariableResponse {
	return v.value
}

func (v *NullableEsvVariableResponse) Set(val *EsvVariableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEsvVariableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEsvVariableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsvVariableResponse(val *EsvVariableResponse) *NullableEsvVariableResponse {
	return &NullableEsvVariableResponse{value: val, isSet: true}
}

func (v NullableEsvVariableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsvVariableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
