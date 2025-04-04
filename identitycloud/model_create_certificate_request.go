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

// checks if the CreateCertificateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCertificateRequest{}

// CreateCertificateRequest struct for CreateCertificateRequest
type CreateCertificateRequest struct {
	// The active status of the certificate. Set this to true for the certificate to actively be served.
	Active *bool `json:"active,omitempty"`
	// The PEM formatted certificate.
	Certificate string `json:"certificate"`
	// The private key for the certificate. For security reasons, only insert requests include this field.
	PrivateKey string `json:"privateKey"`
}

// NewCreateCertificateRequest instantiates a new CreateCertificateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCertificateRequest(certificate string, privateKey string) *CreateCertificateRequest {
	this := CreateCertificateRequest{}
	this.Certificate = certificate
	this.PrivateKey = privateKey
	return &this
}

// NewCreateCertificateRequestWithDefaults instantiates a new CreateCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCertificateRequestWithDefaults() *CreateCertificateRequest {
	this := CreateCertificateRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateCertificateRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateCertificateRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateCertificateRequest) SetActive(v bool) {
	o.Active = &v
}

// GetCertificate returns the Certificate field value
func (o *CreateCertificateRequest) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *CreateCertificateRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *CreateCertificateRequest) SetCertificate(v string) {
	o.Certificate = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *CreateCertificateRequest) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *CreateCertificateRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *CreateCertificateRequest) SetPrivateKey(v string) {
	o.PrivateKey = v
}

func (o CreateCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCertificateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["certificate"] = o.Certificate
	toSerialize["privateKey"] = o.PrivateKey
	return toSerialize, nil
}

type NullableCreateCertificateRequest struct {
	value *CreateCertificateRequest
	isSet bool
}

func (v NullableCreateCertificateRequest) Get() *CreateCertificateRequest {
	return v.value
}

func (v *NullableCreateCertificateRequest) Set(val *CreateCertificateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCertificateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCertificateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCertificateRequest(val *CreateCertificateRequest) *NullableCreateCertificateRequest {
	return &NullableCreateCertificateRequest{value: val, isSet: true}
}

func (v NullableCreateCertificateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCertificateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
