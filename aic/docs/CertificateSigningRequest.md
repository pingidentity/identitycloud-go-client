# CertificateSigningRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | The algorithm for the private key. The encryption algorithm will either be RSA-2048 or ECDSA P-256 depending on the algorithm choice. The default is RSA-2048. | [optional] 
**CertificateID** | Pointer to **string** | The ID of the certificate created from this CSR if the CSR has been completed. | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | The unique identifier for the CSR | [optional] 
**Request** | Pointer to **string** | PEM formatted CSR. | [optional] 
**Subject** | Pointer to **string** | the CSR subject | [optional] 
**SubjectAlternativeNames** | Pointer to **[]string** | Additional domain or domains that the SSL certificate is securing | [optional] 

## Methods

### NewCertificateSigningRequest

`func NewCertificateSigningRequest() *CertificateSigningRequest`

NewCertificateSigningRequest instantiates a new CertificateSigningRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSigningRequestWithDefaults

`func NewCertificateSigningRequestWithDefaults() *CertificateSigningRequest`

NewCertificateSigningRequestWithDefaults instantiates a new CertificateSigningRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *CertificateSigningRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *CertificateSigningRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *CertificateSigningRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *CertificateSigningRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCertificateID

`func (o *CertificateSigningRequest) GetCertificateID() string`

GetCertificateID returns the CertificateID field if non-nil, zero value otherwise.

### GetCertificateIDOk

`func (o *CertificateSigningRequest) GetCertificateIDOk() (*string, bool)`

GetCertificateIDOk returns a tuple with the CertificateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateID

`func (o *CertificateSigningRequest) SetCertificateID(v string)`

SetCertificateID sets CertificateID field to given value.

### HasCertificateID

`func (o *CertificateSigningRequest) HasCertificateID() bool`

HasCertificateID returns a boolean if a field has been set.

### GetCreatedDate

`func (o *CertificateSigningRequest) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CertificateSigningRequest) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CertificateSigningRequest) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CertificateSigningRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *CertificateSigningRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateSigningRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateSigningRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateSigningRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequest

`func (o *CertificateSigningRequest) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *CertificateSigningRequest) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *CertificateSigningRequest) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *CertificateSigningRequest) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSubject

`func (o *CertificateSigningRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateSigningRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateSigningRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertificateSigningRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *CertificateSigningRequest) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *CertificateSigningRequest) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *CertificateSigningRequest) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *CertificateSigningRequest) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


