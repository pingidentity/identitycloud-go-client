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

// checks if the CertificateSigningRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateSigningRequest{}

// CertificateSigningRequest struct for CertificateSigningRequest
type CertificateSigningRequest struct {
	// The algorithm for the private key. The encryption algorithm will either be RSA-2048 or ECDSA P-256 depending on the algorithm choice. The default is RSA-2048.
	Algorithm *string `json:"algorithm,omitempty"`
	// The ID of the certificate created from this CSR if the CSR has been completed.
	CertificateID *string `json:"certificateID,omitempty"`
	CreatedDate   *string `json:"createdDate,omitempty"`
	// The unique identifier for the CSR
	Id *string `json:"id,omitempty"`
	// PEM formatted CSR.
	Request *string `json:"request,omitempty" datastore:",noindex"`
	// the CSR subject
	Subject *string `json:"subject,omitempty"`
	// Additional domain or domains that the SSL certificate is securing
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty"`
}

// NewCertificateSigningRequest instantiates a new CertificateSigningRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateSigningRequest() *CertificateSigningRequest {
	this := CertificateSigningRequest{}
	return &this
}

// NewCertificateSigningRequestWithDefaults instantiates a new CertificateSigningRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateSigningRequestWithDefaults() *CertificateSigningRequest {
	this := CertificateSigningRequest{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *CertificateSigningRequest) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSigningRequest) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *CertificateSigningRequest) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *CertificateSigningRequest) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetCertificateID returns the CertificateID field value if set, zero value otherwise.
func (o *CertificateSigningRequest) GetCertificateID() string {
	if o == nil || IsNil(o.CertificateID) {
		var ret string
		return ret
	}
	return *o.CertificateID
}

// GetCertificateIDOk returns a tuple with the CertificateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSigningRequest) GetCertificateIDOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateID) {
		return nil, false
	}
	return o.CertificateID, true
}

// HasCertificateID returns a boolean if a field has been set.
func (o *CertificateSigningRequest) HasCertificateID() bool {
	if o != nil && !IsNil(o.CertificateID) {
		return true
	}

	return false
}

// SetCertificateID gets a reference to the given string and assigns it to the CertificateID field.
func (o *CertificateSigningRequest) SetCertificateID(v string) {
	o.CertificateID = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *CertificateSigningRequest) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSigningRequest) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CertificateSigningRequest) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *CertificateSigningRequest) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateSigningRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSigningRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateSigningRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CertificateSigningRequest) SetId(v string) {
	o.Id = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *CertificateSigningRequest) GetRequest() string {
	if o == nil || IsNil(o.Request) {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSigningRequest) GetRequestOk() (*string, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *CertificateSigningRequest) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *CertificateSigningRequest) SetRequest(v string) {
	o.Request = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CertificateSigningRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSigningRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CertificateSigningRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CertificateSigningRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetSubjectAlternativeNames returns the SubjectAlternativeNames field value if set, zero value otherwise.
func (o *CertificateSigningRequest) GetSubjectAlternativeNames() []string {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		var ret []string
		return ret
	}
	return o.SubjectAlternativeNames
}

// GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSigningRequest) GetSubjectAlternativeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		return nil, false
	}
	return o.SubjectAlternativeNames, true
}

// HasSubjectAlternativeNames returns a boolean if a field has been set.
func (o *CertificateSigningRequest) HasSubjectAlternativeNames() bool {
	if o != nil && !IsNil(o.SubjectAlternativeNames) {
		return true
	}

	return false
}

// SetSubjectAlternativeNames gets a reference to the given []string and assigns it to the SubjectAlternativeNames field.
func (o *CertificateSigningRequest) SetSubjectAlternativeNames(v []string) {
	o.SubjectAlternativeNames = v
}

func (o CertificateSigningRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateSigningRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.CertificateID) {
		toSerialize["certificateID"] = o.CertificateID
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.SubjectAlternativeNames) {
		toSerialize["subjectAlternativeNames"] = o.SubjectAlternativeNames
	}
	return toSerialize, nil
}

type NullableCertificateSigningRequest struct {
	value *CertificateSigningRequest
	isSet bool
}

func (v NullableCertificateSigningRequest) Get() *CertificateSigningRequest {
	return v.value
}

func (v *NullableCertificateSigningRequest) Set(val *CertificateSigningRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateSigningRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateSigningRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateSigningRequest(val *CertificateSigningRequest) *NullableCertificateSigningRequest {
	return &NullableCertificateSigningRequest{value: val, isSet: true}
}

func (v NullableCertificateSigningRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateSigningRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}