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

// checks if the P1ConnectConfigureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &P1ConnectConfigureRequest{}

// P1ConnectConfigureRequest struct for P1ConnectConfigureRequest
type P1ConnectConfigureRequest struct {
	// The OIDC well-known endpoint for the PingOne environment where admin identities exist
	AdminDiscoveryEndpoint string `json:"adminDiscoveryEndpoint"`
	// The OIDC well-known endpoint for the PingOne consumer environment being linked
	ConsumerDiscoveryEndpoint string `json:"consumerDiscoveryEndpoint"`
	// The JWKS endpoint for the dynamically-registered admin client
	PingOneAdminJwksEndpoint string `json:"pingOneAdminJwksEndpoint"`
	// The client ID to use in the PingOne IDP definition for SSO
	SsoAppId string `json:"ssoAppId"`
}

// NewP1ConnectConfigureRequest instantiates a new P1ConnectConfigureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP1ConnectConfigureRequest(adminDiscoveryEndpoint string, consumerDiscoveryEndpoint string, pingOneAdminJwksEndpoint string, ssoAppId string) *P1ConnectConfigureRequest {
	this := P1ConnectConfigureRequest{}
	this.AdminDiscoveryEndpoint = adminDiscoveryEndpoint
	this.ConsumerDiscoveryEndpoint = consumerDiscoveryEndpoint
	this.PingOneAdminJwksEndpoint = pingOneAdminJwksEndpoint
	this.SsoAppId = ssoAppId
	return &this
}

// NewP1ConnectConfigureRequestWithDefaults instantiates a new P1ConnectConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP1ConnectConfigureRequestWithDefaults() *P1ConnectConfigureRequest {
	this := P1ConnectConfigureRequest{}
	return &this
}

// GetAdminDiscoveryEndpoint returns the AdminDiscoveryEndpoint field value
func (o *P1ConnectConfigureRequest) GetAdminDiscoveryEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdminDiscoveryEndpoint
}

// GetAdminDiscoveryEndpointOk returns a tuple with the AdminDiscoveryEndpoint field value
// and a boolean to check if the value has been set.
func (o *P1ConnectConfigureRequest) GetAdminDiscoveryEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminDiscoveryEndpoint, true
}

// SetAdminDiscoveryEndpoint sets field value
func (o *P1ConnectConfigureRequest) SetAdminDiscoveryEndpoint(v string) {
	o.AdminDiscoveryEndpoint = v
}

// GetConsumerDiscoveryEndpoint returns the ConsumerDiscoveryEndpoint field value
func (o *P1ConnectConfigureRequest) GetConsumerDiscoveryEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerDiscoveryEndpoint
}

// GetConsumerDiscoveryEndpointOk returns a tuple with the ConsumerDiscoveryEndpoint field value
// and a boolean to check if the value has been set.
func (o *P1ConnectConfigureRequest) GetConsumerDiscoveryEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerDiscoveryEndpoint, true
}

// SetConsumerDiscoveryEndpoint sets field value
func (o *P1ConnectConfigureRequest) SetConsumerDiscoveryEndpoint(v string) {
	o.ConsumerDiscoveryEndpoint = v
}

// GetPingOneAdminJwksEndpoint returns the PingOneAdminJwksEndpoint field value
func (o *P1ConnectConfigureRequest) GetPingOneAdminJwksEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PingOneAdminJwksEndpoint
}

// GetPingOneAdminJwksEndpointOk returns a tuple with the PingOneAdminJwksEndpoint field value
// and a boolean to check if the value has been set.
func (o *P1ConnectConfigureRequest) GetPingOneAdminJwksEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PingOneAdminJwksEndpoint, true
}

// SetPingOneAdminJwksEndpoint sets field value
func (o *P1ConnectConfigureRequest) SetPingOneAdminJwksEndpoint(v string) {
	o.PingOneAdminJwksEndpoint = v
}

// GetSsoAppId returns the SsoAppId field value
func (o *P1ConnectConfigureRequest) GetSsoAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsoAppId
}

// GetSsoAppIdOk returns a tuple with the SsoAppId field value
// and a boolean to check if the value has been set.
func (o *P1ConnectConfigureRequest) GetSsoAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsoAppId, true
}

// SetSsoAppId sets field value
func (o *P1ConnectConfigureRequest) SetSsoAppId(v string) {
	o.SsoAppId = v
}

func (o P1ConnectConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o P1ConnectConfigureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adminDiscoveryEndpoint"] = o.AdminDiscoveryEndpoint
	toSerialize["consumerDiscoveryEndpoint"] = o.ConsumerDiscoveryEndpoint
	toSerialize["pingOneAdminJwksEndpoint"] = o.PingOneAdminJwksEndpoint
	toSerialize["ssoAppId"] = o.SsoAppId
	return toSerialize, nil
}

type NullableP1ConnectConfigureRequest struct {
	value *P1ConnectConfigureRequest
	isSet bool
}

func (v NullableP1ConnectConfigureRequest) Get() *P1ConnectConfigureRequest {
	return v.value
}

func (v *NullableP1ConnectConfigureRequest) Set(val *P1ConnectConfigureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableP1ConnectConfigureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableP1ConnectConfigureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP1ConnectConfigureRequest(val *P1ConnectConfigureRequest) *NullableP1ConnectConfigureRequest {
	return &NullableP1ConnectConfigureRequest{value: val, isSet: true}
}

func (v NullableP1ConnectConfigureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP1ConnectConfigureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
