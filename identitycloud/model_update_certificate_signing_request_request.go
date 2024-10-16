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

// checks if the UpdateCertificateSigningRequestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCertificateSigningRequestRequest{}

// UpdateCertificateSigningRequestRequest struct for UpdateCertificateSigningRequestRequest
type UpdateCertificateSigningRequestRequest struct {
	// The matching signed certificate for the request. This should only be set on update requests to upload the certificate.
	Certificate string `json:"certificate"`
}

// NewUpdateCertificateSigningRequestRequest instantiates a new UpdateCertificateSigningRequestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCertificateSigningRequestRequest(certificate string) *UpdateCertificateSigningRequestRequest {
	this := UpdateCertificateSigningRequestRequest{}
	this.Certificate = certificate
	return &this
}

// NewUpdateCertificateSigningRequestRequestWithDefaults instantiates a new UpdateCertificateSigningRequestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCertificateSigningRequestRequestWithDefaults() *UpdateCertificateSigningRequestRequest {
	this := UpdateCertificateSigningRequestRequest{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *UpdateCertificateSigningRequestRequest) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *UpdateCertificateSigningRequestRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *UpdateCertificateSigningRequestRequest) SetCertificate(v string) {
	o.Certificate = v
}

func (o UpdateCertificateSigningRequestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCertificateSigningRequestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificate"] = o.Certificate
	return toSerialize, nil
}

type NullableUpdateCertificateSigningRequestRequest struct {
	value *UpdateCertificateSigningRequestRequest
	isSet bool
}

func (v NullableUpdateCertificateSigningRequestRequest) Get() *UpdateCertificateSigningRequestRequest {
	return v.value
}

func (v *NullableUpdateCertificateSigningRequestRequest) Set(val *UpdateCertificateSigningRequestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCertificateSigningRequestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCertificateSigningRequestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCertificateSigningRequestRequest(val *UpdateCertificateSigningRequestRequest) *NullableUpdateCertificateSigningRequestRequest {
	return &NullableUpdateCertificateSigningRequestRequest{value: val, isSet: true}
}

func (v NullableUpdateCertificateSigningRequestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCertificateSigningRequestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}