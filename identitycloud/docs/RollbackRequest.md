# RollbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promoter** | Pointer to **string** | The name of the admin user who initiated the promotion | [optional] 
**PromotionDescription** | Pointer to **string** | Promoter specified note to make it easier for them to identify what was promoted | [optional] 
**UnlockEnvironmentsAfterPromotion** | Pointer to **bool** | A boolean denoting whether or not to automatically unlock the environments after a successful promotion | [optional] 
**ZendeskTicketReference** | Pointer to **string** | A string denoting the Zendesk ticket reference to be added to the promotion commit message | [optional] 

## Methods

### NewRollbackRequest

`func NewRollbackRequest() *RollbackRequest`

NewRollbackRequest instantiates a new RollbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackRequestWithDefaults

`func NewRollbackRequestWithDefaults() *RollbackRequest`

NewRollbackRequestWithDefaults instantiates a new RollbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromoter

`func (o *RollbackRequest) GetPromoter() string`

GetPromoter returns the Promoter field if non-nil, zero value otherwise.

### GetPromoterOk

`func (o *RollbackRequest) GetPromoterOk() (*string, bool)`

GetPromoterOk returns a tuple with the Promoter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoter

`func (o *RollbackRequest) SetPromoter(v string)`

SetPromoter sets Promoter field to given value.

### HasPromoter

`func (o *RollbackRequest) HasPromoter() bool`

HasPromoter returns a boolean if a field has been set.

### GetPromotionDescription

`func (o *RollbackRequest) GetPromotionDescription() string`

GetPromotionDescription returns the PromotionDescription field if non-nil, zero value otherwise.

### GetPromotionDescriptionOk

`func (o *RollbackRequest) GetPromotionDescriptionOk() (*string, bool)`

GetPromotionDescriptionOk returns a tuple with the PromotionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionDescription

`func (o *RollbackRequest) SetPromotionDescription(v string)`

SetPromotionDescription sets PromotionDescription field to given value.

### HasPromotionDescription

`func (o *RollbackRequest) HasPromotionDescription() bool`

HasPromotionDescription returns a boolean if a field has been set.

### GetUnlockEnvironmentsAfterPromotion

`func (o *RollbackRequest) GetUnlockEnvironmentsAfterPromotion() bool`

GetUnlockEnvironmentsAfterPromotion returns the UnlockEnvironmentsAfterPromotion field if non-nil, zero value otherwise.

### GetUnlockEnvironmentsAfterPromotionOk

`func (o *RollbackRequest) GetUnlockEnvironmentsAfterPromotionOk() (*bool, bool)`

GetUnlockEnvironmentsAfterPromotionOk returns a tuple with the UnlockEnvironmentsAfterPromotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockEnvironmentsAfterPromotion

`func (o *RollbackRequest) SetUnlockEnvironmentsAfterPromotion(v bool)`

SetUnlockEnvironmentsAfterPromotion sets UnlockEnvironmentsAfterPromotion field to given value.

### HasUnlockEnvironmentsAfterPromotion

`func (o *RollbackRequest) HasUnlockEnvironmentsAfterPromotion() bool`

HasUnlockEnvironmentsAfterPromotion returns a boolean if a field has been set.

### GetZendeskTicketReference

`func (o *RollbackRequest) GetZendeskTicketReference() string`

GetZendeskTicketReference returns the ZendeskTicketReference field if non-nil, zero value otherwise.

### GetZendeskTicketReferenceOk

`func (o *RollbackRequest) GetZendeskTicketReferenceOk() (*string, bool)`

GetZendeskTicketReferenceOk returns a tuple with the ZendeskTicketReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZendeskTicketReference

`func (o *RollbackRequest) SetZendeskTicketReference(v string)`

SetZendeskTicketReference sets ZendeskTicketReference field to given value.

### HasZendeskTicketReference

`func (o *RollbackRequest) HasZendeskTicketReference() bool`

HasZendeskTicketReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


