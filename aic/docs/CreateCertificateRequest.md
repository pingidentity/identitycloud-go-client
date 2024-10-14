# CreateCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | The active status of the certificate. Set this to true for the certificate to actively be served. | [optional] 
**Certificate** | **string** | The PEM formatted certificate. | 
**PrivateKey** | **string** | The private key for the certificate. For security reasons, only insert requests include this field. | 

## Methods

### NewCreateCertificateRequest

`func NewCreateCertificateRequest(certificate string, privateKey string, ) *CreateCertificateRequest`

NewCreateCertificateRequest instantiates a new CreateCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCertificateRequestWithDefaults

`func NewCreateCertificateRequestWithDefaults() *CreateCertificateRequest`

NewCreateCertificateRequestWithDefaults instantiates a new CreateCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateCertificateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateCertificateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateCertificateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateCertificateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCertificate

`func (o *CreateCertificateRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateCertificateRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateCertificateRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPrivateKey

`func (o *CreateCertificateRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateCertificateRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateCertificateRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


