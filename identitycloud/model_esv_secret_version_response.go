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

// checks if the EsvSecretVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsvSecretVersionResponse{}

// EsvSecretVersionResponse struct for EsvSecretVersionResponse
type EsvSecretVersionResponse struct {
	CreateDate time.Time `json:"createDate"`
	Loaded     bool      `json:"loaded"`
	Status     string    `json:"status"`
	Version    string    `json:"version"`
}

// NewEsvSecretVersionResponse instantiates a new EsvSecretVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsvSecretVersionResponse(createDate time.Time, loaded bool, status string, version string) *EsvSecretVersionResponse {
	this := EsvSecretVersionResponse{}
	this.CreateDate = createDate
	this.Loaded = loaded
	this.Status = status
	this.Version = version
	return &this
}

// NewEsvSecretVersionResponseWithDefaults instantiates a new EsvSecretVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsvSecretVersionResponseWithDefaults() *EsvSecretVersionResponse {
	this := EsvSecretVersionResponse{}
	return &this
}

// GetCreateDate returns the CreateDate field value
func (o *EsvSecretVersionResponse) GetCreateDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value
// and a boolean to check if the value has been set.
func (o *EsvSecretVersionResponse) GetCreateDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateDate, true
}

// SetCreateDate sets field value
func (o *EsvSecretVersionResponse) SetCreateDate(v time.Time) {
	o.CreateDate = v
}

// GetLoaded returns the Loaded field value
func (o *EsvSecretVersionResponse) GetLoaded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Loaded
}

// GetLoadedOk returns a tuple with the Loaded field value
// and a boolean to check if the value has been set.
func (o *EsvSecretVersionResponse) GetLoadedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loaded, true
}

// SetLoaded sets field value
func (o *EsvSecretVersionResponse) SetLoaded(v bool) {
	o.Loaded = v
}

// GetStatus returns the Status field value
func (o *EsvSecretVersionResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EsvSecretVersionResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EsvSecretVersionResponse) SetStatus(v string) {
	o.Status = v
}

// GetVersion returns the Version field value
func (o *EsvSecretVersionResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EsvSecretVersionResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EsvSecretVersionResponse) SetVersion(v string) {
	o.Version = v
}

func (o EsvSecretVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsvSecretVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createDate"] = o.CreateDate
	toSerialize["loaded"] = o.Loaded
	toSerialize["status"] = o.Status
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableEsvSecretVersionResponse struct {
	value *EsvSecretVersionResponse
	isSet bool
}

func (v NullableEsvSecretVersionResponse) Get() *EsvSecretVersionResponse {
	return v.value
}

func (v *NullableEsvSecretVersionResponse) Set(val *EsvSecretVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEsvSecretVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEsvSecretVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsvSecretVersionResponse(val *EsvSecretVersionResponse) *NullableEsvSecretVersionResponse {
	return &NullableEsvSecretVersionResponse{value: val, isSet: true}
}

func (v NullableEsvSecretVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsvSecretVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}