# P1ConnectConfigureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminDiscoveryEndpoint** | **string** | The OIDC well-known endpoint for the PingOne environment where admin identities exist | 
**ConsumerDiscoveryEndpoint** | **string** | The OIDC well-known endpoint for the PingOne consumer environment being linked | 
**PingOneAdminJwksEndpoint** | **string** | The JWKS endpoint for the dynamically-registered admin client | 
**SsoAppId** | **string** | The client ID to use in the PingOne IDP definition for SSO | 

## Methods

### NewP1ConnectConfigureRequest

`func NewP1ConnectConfigureRequest(adminDiscoveryEndpoint string, consumerDiscoveryEndpoint string, pingOneAdminJwksEndpoint string, ssoAppId string, ) *P1ConnectConfigureRequest`

NewP1ConnectConfigureRequest instantiates a new P1ConnectConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP1ConnectConfigureRequestWithDefaults

`func NewP1ConnectConfigureRequestWithDefaults() *P1ConnectConfigureRequest`

NewP1ConnectConfigureRequestWithDefaults instantiates a new P1ConnectConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminDiscoveryEndpoint

`func (o *P1ConnectConfigureRequest) GetAdminDiscoveryEndpoint() string`

GetAdminDiscoveryEndpoint returns the AdminDiscoveryEndpoint field if non-nil, zero value otherwise.

### GetAdminDiscoveryEndpointOk

`func (o *P1ConnectConfigureRequest) GetAdminDiscoveryEndpointOk() (*string, bool)`

GetAdminDiscoveryEndpointOk returns a tuple with the AdminDiscoveryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDiscoveryEndpoint

`func (o *P1ConnectConfigureRequest) SetAdminDiscoveryEndpoint(v string)`

SetAdminDiscoveryEndpoint sets AdminDiscoveryEndpoint field to given value.


### GetConsumerDiscoveryEndpoint

`func (o *P1ConnectConfigureRequest) GetConsumerDiscoveryEndpoint() string`

GetConsumerDiscoveryEndpoint returns the ConsumerDiscoveryEndpoint field if non-nil, zero value otherwise.

### GetConsumerDiscoveryEndpointOk

`func (o *P1ConnectConfigureRequest) GetConsumerDiscoveryEndpointOk() (*string, bool)`

GetConsumerDiscoveryEndpointOk returns a tuple with the ConsumerDiscoveryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerDiscoveryEndpoint

`func (o *P1ConnectConfigureRequest) SetConsumerDiscoveryEndpoint(v string)`

SetConsumerDiscoveryEndpoint sets ConsumerDiscoveryEndpoint field to given value.


### GetPingOneAdminJwksEndpoint

`func (o *P1ConnectConfigureRequest) GetPingOneAdminJwksEndpoint() string`

GetPingOneAdminJwksEndpoint returns the PingOneAdminJwksEndpoint field if non-nil, zero value otherwise.

### GetPingOneAdminJwksEndpointOk

`func (o *P1ConnectConfigureRequest) GetPingOneAdminJwksEndpointOk() (*string, bool)`

GetPingOneAdminJwksEndpointOk returns a tuple with the PingOneAdminJwksEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneAdminJwksEndpoint

`func (o *P1ConnectConfigureRequest) SetPingOneAdminJwksEndpoint(v string)`

SetPingOneAdminJwksEndpoint sets PingOneAdminJwksEndpoint field to given value.


### GetSsoAppId

`func (o *P1ConnectConfigureRequest) GetSsoAppId() string`

GetSsoAppId returns the SsoAppId field if non-nil, zero value otherwise.

### GetSsoAppIdOk

`func (o *P1ConnectConfigureRequest) GetSsoAppIdOk() (*string, bool)`

GetSsoAppIdOk returns a tuple with the SsoAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAppId

`func (o *P1ConnectConfigureRequest) SetSsoAppId(v string)`

SetSsoAppId sets SsoAppId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


