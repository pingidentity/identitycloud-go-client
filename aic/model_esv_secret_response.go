/*
PingOne Advanced Identity Cloud API

## Introduction The PingOne Advanced Identity Cloud REST API lets you manage your Advanced Identity Cloud tenants. The API exposes access management and identity management endpoints, with additional endpoints specific to Advanced Identity Cloud tenant environments.<br /><br /> We are now publishing the API spec in OpenAPI 3.0. For the legacy Swagger 2.0 spec, please download [swagger.yaml](swagger.yaml), but note that it may not contain all new functionality.<br /><br /> For full PingOne Advanced Identity Cloud documentation, please visit [the docs website](https://docs.pingidentity.com/pingoneaic/latest/). ## Authenticating to the API The PingOne Advanced Identity Cloud REST API has two different authentication methods:   - [API key and secret](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-with-api-key-and-secret.html): used for tenant read-only operations  - [Access token](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-with-access-token.html): used for access management operations, identity management operations or tenant write operations  For a summary of how to use these authentication methods, refer to [Authenticate to Advanced Identity Cloud REST API](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-overview.html).

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aic

import (
	"encoding/json"
	"time"
)

// checks if the EsvSecretResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsvSecretResponse{}

// EsvSecretResponse struct for EsvSecretResponse
type EsvSecretResponse struct {
	Id                string    `json:"_id"`
	ActiveVersion     string    `json:"activeVersion"`
	Description       string    `json:"description"`
	Encoding          string    `json:"encoding"`
	LastChangeDate    time.Time `json:"lastChangeDate"`
	LastChangedBy     string    `json:"lastChangedBy"`
	Loaded            bool      `json:"loaded"`
	LoadedVersion     string    `json:"loadedVersion"`
	UseInPlaceholders bool      `json:"useInPlaceholders"`
}

// NewEsvSecretResponse instantiates a new EsvSecretResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsvSecretResponse(id string, activeVersion string, description string, encoding string, lastChangeDate time.Time, lastChangedBy string, loaded bool, loadedVersion string, useInPlaceholders bool) *EsvSecretResponse {
	this := EsvSecretResponse{}
	this.Id = id
	this.ActiveVersion = activeVersion
	this.Description = description
	this.Encoding = encoding
	this.LastChangeDate = lastChangeDate
	this.LastChangedBy = lastChangedBy
	this.Loaded = loaded
	this.LoadedVersion = loadedVersion
	this.UseInPlaceholders = useInPlaceholders
	return &this
}

// NewEsvSecretResponseWithDefaults instantiates a new EsvSecretResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsvSecretResponseWithDefaults() *EsvSecretResponse {
	this := EsvSecretResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EsvSecretResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EsvSecretResponse) SetId(v string) {
	o.Id = v
}

// GetActiveVersion returns the ActiveVersion field value
func (o *EsvSecretResponse) GetActiveVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveVersion
}

// GetActiveVersionOk returns a tuple with the ActiveVersion field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetActiveVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveVersion, true
}

// SetActiveVersion sets field value
func (o *EsvSecretResponse) SetActiveVersion(v string) {
	o.ActiveVersion = v
}

// GetDescription returns the Description field value
func (o *EsvSecretResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EsvSecretResponse) SetDescription(v string) {
	o.Description = v
}

// GetEncoding returns the Encoding field value
func (o *EsvSecretResponse) GetEncoding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value
func (o *EsvSecretResponse) SetEncoding(v string) {
	o.Encoding = v
}

// GetLastChangeDate returns the LastChangeDate field value
func (o *EsvSecretResponse) GetLastChangeDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastChangeDate
}

// GetLastChangeDateOk returns a tuple with the LastChangeDate field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetLastChangeDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastChangeDate, true
}

// SetLastChangeDate sets field value
func (o *EsvSecretResponse) SetLastChangeDate(v time.Time) {
	o.LastChangeDate = v
}

// GetLastChangedBy returns the LastChangedBy field value
func (o *EsvSecretResponse) GetLastChangedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastChangedBy
}

// GetLastChangedByOk returns a tuple with the LastChangedBy field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetLastChangedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastChangedBy, true
}

// SetLastChangedBy sets field value
func (o *EsvSecretResponse) SetLastChangedBy(v string) {
	o.LastChangedBy = v
}

// GetLoaded returns the Loaded field value
func (o *EsvSecretResponse) GetLoaded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Loaded
}

// GetLoadedOk returns a tuple with the Loaded field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetLoadedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loaded, true
}

// SetLoaded sets field value
func (o *EsvSecretResponse) SetLoaded(v bool) {
	o.Loaded = v
}

// GetLoadedVersion returns the LoadedVersion field value
func (o *EsvSecretResponse) GetLoadedVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadedVersion
}

// GetLoadedVersionOk returns a tuple with the LoadedVersion field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetLoadedVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadedVersion, true
}

// SetLoadedVersion sets field value
func (o *EsvSecretResponse) SetLoadedVersion(v string) {
	o.LoadedVersion = v
}

// GetUseInPlaceholders returns the UseInPlaceholders field value
func (o *EsvSecretResponse) GetUseInPlaceholders() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseInPlaceholders
}

// GetUseInPlaceholdersOk returns a tuple with the UseInPlaceholders field value
// and a boolean to check if the value has been set.
func (o *EsvSecretResponse) GetUseInPlaceholdersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseInPlaceholders, true
}

// SetUseInPlaceholders sets field value
func (o *EsvSecretResponse) SetUseInPlaceholders(v bool) {
	o.UseInPlaceholders = v
}

func (o EsvSecretResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsvSecretResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["activeVersion"] = o.ActiveVersion
	toSerialize["description"] = o.Description
	toSerialize["encoding"] = o.Encoding
	toSerialize["lastChangeDate"] = o.LastChangeDate
	toSerialize["lastChangedBy"] = o.LastChangedBy
	toSerialize["loaded"] = o.Loaded
	toSerialize["loadedVersion"] = o.LoadedVersion
	toSerialize["useInPlaceholders"] = o.UseInPlaceholders
	return toSerialize, nil
}

type NullableEsvSecretResponse struct {
	value *EsvSecretResponse
	isSet bool
}

func (v NullableEsvSecretResponse) Get() *EsvSecretResponse {
	return v.value
}

func (v *NullableEsvSecretResponse) Set(val *EsvSecretResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEsvSecretResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEsvSecretResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsvSecretResponse(val *EsvSecretResponse) *NullableEsvSecretResponse {
	return &NullableEsvSecretResponse{value: val, isSet: true}
}

func (v NullableEsvSecretResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsvSecretResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
