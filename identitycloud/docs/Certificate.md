# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | The active status of the certificate. Set this to true for the certificate to be made live on the loadbalancer. | [optional] 
**Certificate** | Pointer to **string** | Deprecated. Contains no data. | [optional] 
**ExpireTime** | Pointer to **string** | The expiry time of the certificate. Not used in insert requests. | [optional] 
**Id** | Pointer to **string** | The unique identifier for the certificate resource. Not used in insert requests. | [optional] 
**Issuer** | Pointer to **string** | The issuer DN of the certificate | [optional] 
**Live** | Pointer to **bool** | The live status of the certificate. This is automatically set by the system and indicates if the certificate is currently live externally. | [optional] 
**Subject** | Pointer to **string** | The subject DN of the certificate | [optional] 
**SubjectAlternativeNames** | Pointer to **[]string** | Domains associated by with the certificate via the Subject Alternative Name extension. The common name should be included in the SANs as well. Not used in insert requests. | [optional] 
**ValidFromTime** | Pointer to **string** | The notBefore time of the certificate. | [optional] 

## Methods

### NewCertificate

`func NewCertificate() *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Certificate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Certificate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Certificate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Certificate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCertificate

`func (o *Certificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Certificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Certificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *Certificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetExpireTime

`func (o *Certificate) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *Certificate) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *Certificate) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *Certificate) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetId

`func (o *Certificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Certificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *Certificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Certificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Certificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Certificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLive

`func (o *Certificate) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *Certificate) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *Certificate) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *Certificate) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetSubject

`func (o *Certificate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Certificate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Certificate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Certificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *Certificate) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *Certificate) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *Certificate) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *Certificate) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetValidFromTime

`func (o *Certificate) GetValidFromTime() string`

GetValidFromTime returns the ValidFromTime field if non-nil, zero value otherwise.

### GetValidFromTimeOk

`func (o *Certificate) GetValidFromTimeOk() (*string, bool)`

GetValidFromTimeOk returns a tuple with the ValidFromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFromTime

`func (o *Certificate) SetValidFromTime(v string)`

SetValidFromTime sets ValidFromTime field to given value.

### HasValidFromTime

`func (o *Certificate) HasValidFromTime() bool`

HasValidFromTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


