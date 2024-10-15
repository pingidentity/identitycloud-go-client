# P1Connect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**AdminClientId** | Pointer to **string** |  | [optional] 
**DefaultIssuer** | Pointer to **string** |  | [optional] 
**IdpAuthEndpoint** | Pointer to **string** |  | [optional] 
**IdpClientId** | Pointer to **string** |  | [optional] 
**IdpEnabled** | Pointer to **bool** |  | [optional] 
**IdpIssuer** | Pointer to **string** |  | [optional] 
**IdpTokenEndpoint** | Pointer to **string** |  | [optional] 
**IdpWellKnownEndpoint** | Pointer to **string** |  | [optional] 

## Methods

### NewP1Connect

`func NewP1Connect() *P1Connect`

NewP1Connect instantiates a new P1Connect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP1ConnectWithDefaults

`func NewP1ConnectWithDefaults() *P1Connect`

NewP1ConnectWithDefaults instantiates a new P1Connect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *P1Connect) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *P1Connect) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *P1Connect) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *P1Connect) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAdminClientId

`func (o *P1Connect) GetAdminClientId() string`

GetAdminClientId returns the AdminClientId field if non-nil, zero value otherwise.

### GetAdminClientIdOk

`func (o *P1Connect) GetAdminClientIdOk() (*string, bool)`

GetAdminClientIdOk returns a tuple with the AdminClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminClientId

`func (o *P1Connect) SetAdminClientId(v string)`

SetAdminClientId sets AdminClientId field to given value.

### HasAdminClientId

`func (o *P1Connect) HasAdminClientId() bool`

HasAdminClientId returns a boolean if a field has been set.

### GetDefaultIssuer

`func (o *P1Connect) GetDefaultIssuer() string`

GetDefaultIssuer returns the DefaultIssuer field if non-nil, zero value otherwise.

### GetDefaultIssuerOk

`func (o *P1Connect) GetDefaultIssuerOk() (*string, bool)`

GetDefaultIssuerOk returns a tuple with the DefaultIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIssuer

`func (o *P1Connect) SetDefaultIssuer(v string)`

SetDefaultIssuer sets DefaultIssuer field to given value.

### HasDefaultIssuer

`func (o *P1Connect) HasDefaultIssuer() bool`

HasDefaultIssuer returns a boolean if a field has been set.

### GetIdpAuthEndpoint

`func (o *P1Connect) GetIdpAuthEndpoint() string`

GetIdpAuthEndpoint returns the IdpAuthEndpoint field if non-nil, zero value otherwise.

### GetIdpAuthEndpointOk

`func (o *P1Connect) GetIdpAuthEndpointOk() (*string, bool)`

GetIdpAuthEndpointOk returns a tuple with the IdpAuthEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpAuthEndpoint

`func (o *P1Connect) SetIdpAuthEndpoint(v string)`

SetIdpAuthEndpoint sets IdpAuthEndpoint field to given value.

### HasIdpAuthEndpoint

`func (o *P1Connect) HasIdpAuthEndpoint() bool`

HasIdpAuthEndpoint returns a boolean if a field has been set.

### GetIdpClientId

`func (o *P1Connect) GetIdpClientId() string`

GetIdpClientId returns the IdpClientId field if non-nil, zero value otherwise.

### GetIdpClientIdOk

`func (o *P1Connect) GetIdpClientIdOk() (*string, bool)`

GetIdpClientIdOk returns a tuple with the IdpClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpClientId

`func (o *P1Connect) SetIdpClientId(v string)`

SetIdpClientId sets IdpClientId field to given value.

### HasIdpClientId

`func (o *P1Connect) HasIdpClientId() bool`

HasIdpClientId returns a boolean if a field has been set.

### GetIdpEnabled

`func (o *P1Connect) GetIdpEnabled() bool`

GetIdpEnabled returns the IdpEnabled field if non-nil, zero value otherwise.

### GetIdpEnabledOk

`func (o *P1Connect) GetIdpEnabledOk() (*bool, bool)`

GetIdpEnabledOk returns a tuple with the IdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEnabled

`func (o *P1Connect) SetIdpEnabled(v bool)`

SetIdpEnabled sets IdpEnabled field to given value.

### HasIdpEnabled

`func (o *P1Connect) HasIdpEnabled() bool`

HasIdpEnabled returns a boolean if a field has been set.

### GetIdpIssuer

`func (o *P1Connect) GetIdpIssuer() string`

GetIdpIssuer returns the IdpIssuer field if non-nil, zero value otherwise.

### GetIdpIssuerOk

`func (o *P1Connect) GetIdpIssuerOk() (*string, bool)`

GetIdpIssuerOk returns a tuple with the IdpIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpIssuer

`func (o *P1Connect) SetIdpIssuer(v string)`

SetIdpIssuer sets IdpIssuer field to given value.

### HasIdpIssuer

`func (o *P1Connect) HasIdpIssuer() bool`

HasIdpIssuer returns a boolean if a field has been set.

### GetIdpTokenEndpoint

`func (o *P1Connect) GetIdpTokenEndpoint() string`

GetIdpTokenEndpoint returns the IdpTokenEndpoint field if non-nil, zero value otherwise.

### GetIdpTokenEndpointOk

`func (o *P1Connect) GetIdpTokenEndpointOk() (*string, bool)`

GetIdpTokenEndpointOk returns a tuple with the IdpTokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpTokenEndpoint

`func (o *P1Connect) SetIdpTokenEndpoint(v string)`

SetIdpTokenEndpoint sets IdpTokenEndpoint field to given value.

### HasIdpTokenEndpoint

`func (o *P1Connect) HasIdpTokenEndpoint() bool`

HasIdpTokenEndpoint returns a boolean if a field has been set.

### GetIdpWellKnownEndpoint

`func (o *P1Connect) GetIdpWellKnownEndpoint() string`

GetIdpWellKnownEndpoint returns the IdpWellKnownEndpoint field if non-nil, zero value otherwise.

### GetIdpWellKnownEndpointOk

`func (o *P1Connect) GetIdpWellKnownEndpointOk() (*string, bool)`

GetIdpWellKnownEndpointOk returns a tuple with the IdpWellKnownEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpWellKnownEndpoint

`func (o *P1Connect) SetIdpWellKnownEndpoint(v string)`

SetIdpWellKnownEndpoint sets IdpWellKnownEndpoint field to given value.

### HasIdpWellKnownEndpoint

`func (o *P1Connect) HasIdpWellKnownEndpoint() bool`

HasIdpWellKnownEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


