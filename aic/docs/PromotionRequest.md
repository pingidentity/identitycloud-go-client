# PromotionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | A boolean denoting whether this a dry run | 
**IgnoreEncryptedSecrets** | Pointer to **bool** | A boolean denoting whether or not to ignore encrypted secrets detected in config during the promotion | [optional] 
**Promoter** | Pointer to **string** | The name of the admin user who initiated the promotion | [optional] 
**PromotionDescription** | Pointer to **string** | Promoter specified note to make it easier for them to identify what was promoted | [optional] 
**UnlockEnvironmentsAfterPromotion** | Pointer to **bool** | A boolean denoting whether or not to automatically unlock the environments after a successful promotion | [optional] 
**ZendeskTicketReference** | Pointer to **string** | A string denoting the Zendesk ticket reference to be added to the promotion commit message | [optional] 

## Methods

### NewPromotionRequest

`func NewPromotionRequest(dryRun bool, ) *PromotionRequest`

NewPromotionRequest instantiates a new PromotionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionRequestWithDefaults

`func NewPromotionRequestWithDefaults() *PromotionRequest`

NewPromotionRequestWithDefaults instantiates a new PromotionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *PromotionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PromotionRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PromotionRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetIgnoreEncryptedSecrets

`func (o *PromotionRequest) GetIgnoreEncryptedSecrets() bool`

GetIgnoreEncryptedSecrets returns the IgnoreEncryptedSecrets field if non-nil, zero value otherwise.

### GetIgnoreEncryptedSecretsOk

`func (o *PromotionRequest) GetIgnoreEncryptedSecretsOk() (*bool, bool)`

GetIgnoreEncryptedSecretsOk returns a tuple with the IgnoreEncryptedSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEncryptedSecrets

`func (o *PromotionRequest) SetIgnoreEncryptedSecrets(v bool)`

SetIgnoreEncryptedSecrets sets IgnoreEncryptedSecrets field to given value.

### HasIgnoreEncryptedSecrets

`func (o *PromotionRequest) HasIgnoreEncryptedSecrets() bool`

HasIgnoreEncryptedSecrets returns a boolean if a field has been set.

### GetPromoter

`func (o *PromotionRequest) GetPromoter() string`

GetPromoter returns the Promoter field if non-nil, zero value otherwise.

### GetPromoterOk

`func (o *PromotionRequest) GetPromoterOk() (*string, bool)`

GetPromoterOk returns a tuple with the Promoter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoter

`func (o *PromotionRequest) SetPromoter(v string)`

SetPromoter sets Promoter field to given value.

### HasPromoter

`func (o *PromotionRequest) HasPromoter() bool`

HasPromoter returns a boolean if a field has been set.

### GetPromotionDescription

`func (o *PromotionRequest) GetPromotionDescription() string`

GetPromotionDescription returns the PromotionDescription field if non-nil, zero value otherwise.

### GetPromotionDescriptionOk

`func (o *PromotionRequest) GetPromotionDescriptionOk() (*string, bool)`

GetPromotionDescriptionOk returns a tuple with the PromotionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionDescription

`func (o *PromotionRequest) SetPromotionDescription(v string)`

SetPromotionDescription sets PromotionDescription field to given value.

### HasPromotionDescription

`func (o *PromotionRequest) HasPromotionDescription() bool`

HasPromotionDescription returns a boolean if a field has been set.

### GetUnlockEnvironmentsAfterPromotion

`func (o *PromotionRequest) GetUnlockEnvironmentsAfterPromotion() bool`

GetUnlockEnvironmentsAfterPromotion returns the UnlockEnvironmentsAfterPromotion field if non-nil, zero value otherwise.

### GetUnlockEnvironmentsAfterPromotionOk

`func (o *PromotionRequest) GetUnlockEnvironmentsAfterPromotionOk() (*bool, bool)`

GetUnlockEnvironmentsAfterPromotionOk returns a tuple with the UnlockEnvironmentsAfterPromotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockEnvironmentsAfterPromotion

`func (o *PromotionRequest) SetUnlockEnvironmentsAfterPromotion(v bool)`

SetUnlockEnvironmentsAfterPromotion sets UnlockEnvironmentsAfterPromotion field to given value.

### HasUnlockEnvironmentsAfterPromotion

`func (o *PromotionRequest) HasUnlockEnvironmentsAfterPromotion() bool`

HasUnlockEnvironmentsAfterPromotion returns a boolean if a field has been set.

### GetZendeskTicketReference

`func (o *PromotionRequest) GetZendeskTicketReference() string`

GetZendeskTicketReference returns the ZendeskTicketReference field if non-nil, zero value otherwise.

### GetZendeskTicketReferenceOk

`func (o *PromotionRequest) GetZendeskTicketReferenceOk() (*string, bool)`

GetZendeskTicketReferenceOk returns a tuple with the ZendeskTicketReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskTicketReference

`func (o *PromotionRequest) SetZendeskTicketReference(v string)`

SetZendeskTicketReference sets ZendeskTicketReference field to given value.

### HasZendeskTicketReference

`func (o *PromotionRequest) HasZendeskTicketReference() bool`

HasZendeskTicketReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


