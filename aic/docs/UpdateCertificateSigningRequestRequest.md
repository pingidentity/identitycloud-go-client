# UpdateCertificateSigningRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | The matching signed certificate for the request. This should only be set on update requests to upload the certificate. | 

## Methods

### NewUpdateCertificateSigningRequestRequest

`func NewUpdateCertificateSigningRequestRequest(certificate string, ) *UpdateCertificateSigningRequestRequest`

NewUpdateCertificateSigningRequestRequest instantiates a new UpdateCertificateSigningRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCertificateSigningRequestRequestWithDefaults

`func NewUpdateCertificateSigningRequestRequestWithDefaults() *UpdateCertificateSigningRequestRequest`

NewUpdateCertificateSigningRequestRequestWithDefaults instantiates a new UpdateCertificateSigningRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *UpdateCertificateSigningRequestRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateCertificateSigningRequestRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateCertificateSigningRequestRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


