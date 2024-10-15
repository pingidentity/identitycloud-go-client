# CreateCertificateSigningRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | The algorithm for the private key. The encryption algorithm will either be RSA-2048 or ECDSA P-256 depending on the algorithm choice. The default is RSA-2048. | [optional] 
**BusinessCategory** | Pointer to **string** | Category of business, such as \&quot;Private Organization\&quot;, “Government Entity”, “Business Entity”, or “Non-Commercial Entity”. Relevant for EV certificates. | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** | Domain name that the SSL certificate is securing | [optional] 
**Country** | Pointer to **string** | Two-letter ISO-3166 country code | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**JurisdictionCity** | Pointer to **string** | This field contains only information relevant to the Jurisdiction of Incorporation or Registration. Relevant for EV certificates. | [optional] 
**JurisdictionCountry** | Pointer to **string** | This field contains only information relevant to the Jurisdiction of Incorporation or Registration. Relevant for EV certificates. | [optional] 
**JurisdictionState** | Pointer to **string** | This field contains only information relevant to the Jurisdiction of Incorporation or Registration. Relevant for EV certificates. | [optional] 
**Organization** | Pointer to **string** | Full name of company | [optional] 
**OrganizationalUnit** | Pointer to **string** | Company section or department | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** | The Registration (or similar) Number assigned to the Subject by the Incorporating or Registration Agency in its Jurisdiction of Incorporation or Registration. Relevant for EV certificates. | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StreetAddress** | Pointer to **string** |  | [optional] 
**SubjectAlternativeNames** | Pointer to **[]string** | Additional domain or domains that the SSL certificate is securing | [optional] 

## Methods

### NewCreateCertificateSigningRequestRequest

`func NewCreateCertificateSigningRequestRequest() *CreateCertificateSigningRequestRequest`

NewCreateCertificateSigningRequestRequest instantiates a new CreateCertificateSigningRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCertificateSigningRequestRequestWithDefaults

`func NewCreateCertificateSigningRequestRequestWithDefaults() *CreateCertificateSigningRequestRequest`

NewCreateCertificateSigningRequestRequestWithDefaults instantiates a new CreateCertificateSigningRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *CreateCertificateSigningRequestRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *CreateCertificateSigningRequestRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *CreateCertificateSigningRequestRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *CreateCertificateSigningRequestRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetBusinessCategory

`func (o *CreateCertificateSigningRequestRequest) GetBusinessCategory() string`

GetBusinessCategory returns the BusinessCategory field if non-nil, zero value otherwise.

### GetBusinessCategoryOk

`func (o *CreateCertificateSigningRequestRequest) GetBusinessCategoryOk() (*string, bool)`

GetBusinessCategoryOk returns a tuple with the BusinessCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCategory

`func (o *CreateCertificateSigningRequestRequest) SetBusinessCategory(v string)`

SetBusinessCategory sets BusinessCategory field to given value.

### HasBusinessCategory

`func (o *CreateCertificateSigningRequestRequest) HasBusinessCategory() bool`

HasBusinessCategory returns a boolean if a field has been set.

### GetCity

`func (o *CreateCertificateSigningRequestRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreateCertificateSigningRequestRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CreateCertificateSigningRequestRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CreateCertificateSigningRequestRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCommonName

`func (o *CreateCertificateSigningRequestRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CreateCertificateSigningRequestRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CreateCertificateSigningRequestRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CreateCertificateSigningRequestRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *CreateCertificateSigningRequestRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateCertificateSigningRequestRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateCertificateSigningRequestRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateCertificateSigningRequestRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *CreateCertificateSigningRequestRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateCertificateSigningRequestRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateCertificateSigningRequestRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateCertificateSigningRequestRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetJurisdictionCity

`func (o *CreateCertificateSigningRequestRequest) GetJurisdictionCity() string`

GetJurisdictionCity returns the JurisdictionCity field if non-nil, zero value otherwise.

### GetJurisdictionCityOk

`func (o *CreateCertificateSigningRequestRequest) GetJurisdictionCityOk() (*string, bool)`

GetJurisdictionCityOk returns a tuple with the JurisdictionCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionCity

`func (o *CreateCertificateSigningRequestRequest) SetJurisdictionCity(v string)`

SetJurisdictionCity sets JurisdictionCity field to given value.

### HasJurisdictionCity

`func (o *CreateCertificateSigningRequestRequest) HasJurisdictionCity() bool`

HasJurisdictionCity returns a boolean if a field has been set.

### GetJurisdictionCountry

`func (o *CreateCertificateSigningRequestRequest) GetJurisdictionCountry() string`

GetJurisdictionCountry returns the JurisdictionCountry field if non-nil, zero value otherwise.

### GetJurisdictionCountryOk

`func (o *CreateCertificateSigningRequestRequest) GetJurisdictionCountryOk() (*string, bool)`

GetJurisdictionCountryOk returns a tuple with the JurisdictionCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionCountry

`func (o *CreateCertificateSigningRequestRequest) SetJurisdictionCountry(v string)`

SetJurisdictionCountry sets JurisdictionCountry field to given value.

### HasJurisdictionCountry

`func (o *CreateCertificateSigningRequestRequest) HasJurisdictionCountry() bool`

HasJurisdictionCountry returns a boolean if a field has been set.

### GetJurisdictionState

`func (o *CreateCertificateSigningRequestRequest) GetJurisdictionState() string`

GetJurisdictionState returns the JurisdictionState field if non-nil, zero value otherwise.

### GetJurisdictionStateOk

`func (o *CreateCertificateSigningRequestRequest) GetJurisdictionStateOk() (*string, bool)`

GetJurisdictionStateOk returns a tuple with the JurisdictionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionState

`func (o *CreateCertificateSigningRequestRequest) SetJurisdictionState(v string)`

SetJurisdictionState sets JurisdictionState field to given value.

### HasJurisdictionState

`func (o *CreateCertificateSigningRequestRequest) HasJurisdictionState() bool`

HasJurisdictionState returns a boolean if a field has been set.

### GetOrganization

`func (o *CreateCertificateSigningRequestRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CreateCertificateSigningRequestRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CreateCertificateSigningRequestRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CreateCertificateSigningRequestRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *CreateCertificateSigningRequestRequest) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *CreateCertificateSigningRequestRequest) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *CreateCertificateSigningRequestRequest) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *CreateCertificateSigningRequestRequest) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetPostalCode

`func (o *CreateCertificateSigningRequestRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CreateCertificateSigningRequestRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CreateCertificateSigningRequestRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CreateCertificateSigningRequestRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CreateCertificateSigningRequestRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CreateCertificateSigningRequestRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CreateCertificateSigningRequestRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CreateCertificateSigningRequestRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetState

`func (o *CreateCertificateSigningRequestRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateCertificateSigningRequestRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateCertificateSigningRequestRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateCertificateSigningRequestRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreetAddress

`func (o *CreateCertificateSigningRequestRequest) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *CreateCertificateSigningRequestRequest) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *CreateCertificateSigningRequestRequest) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *CreateCertificateSigningRequestRequest) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *CreateCertificateSigningRequestRequest) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *CreateCertificateSigningRequestRequest) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *CreateCertificateSigningRequestRequest) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *CreateCertificateSigningRequestRequest) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


