# PromotionReportMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 
**Promoter** | Pointer to **string** |  | [optional] 
**PromotionId** | Pointer to **string** |  | [optional] 
**ReportId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Either promotion or rollback depending on whether the report was for a config promotion or a config rollback | [optional] 

## Methods

### NewPromotionReportMeta

`func NewPromotionReportMeta() *PromotionReportMeta`

NewPromotionReportMeta instantiates a new PromotionReportMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionReportMetaWithDefaults

`func NewPromotionReportMetaWithDefaults() *PromotionReportMeta`

NewPromotionReportMetaWithDefaults instantiates a new PromotionReportMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *PromotionReportMeta) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *PromotionReportMeta) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *PromotionReportMeta) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *PromotionReportMeta) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDryRun

`func (o *PromotionReportMeta) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PromotionReportMeta) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PromotionReportMeta) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PromotionReportMeta) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetPromoter

`func (o *PromotionReportMeta) GetPromoter() string`

GetPromoter returns the Promoter field if non-nil, zero value otherwise.

### GetPromoterOk

`func (o *PromotionReportMeta) GetPromoterOk() (*string, bool)`

GetPromoterOk returns a tuple with the Promoter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoter

`func (o *PromotionReportMeta) SetPromoter(v string)`

SetPromoter sets Promoter field to given value.

### HasPromoter

`func (o *PromotionReportMeta) HasPromoter() bool`

HasPromoter returns a boolean if a field has been set.

### GetPromotionId

`func (o *PromotionReportMeta) GetPromotionId() string`

GetPromotionId returns the PromotionId field if non-nil, zero value otherwise.

### GetPromotionIdOk

`func (o *PromotionReportMeta) GetPromotionIdOk() (*string, bool)`

GetPromotionIdOk returns a tuple with the PromotionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionId

`func (o *PromotionReportMeta) SetPromotionId(v string)`

SetPromotionId sets PromotionId field to given value.

### HasPromotionId

`func (o *PromotionReportMeta) HasPromotionId() bool`

HasPromotionId returns a boolean if a field has been set.

### GetReportId

`func (o *PromotionReportMeta) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *PromotionReportMeta) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *PromotionReportMeta) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *PromotionReportMeta) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetType

`func (o *PromotionReportMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromotionReportMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromotionReportMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromotionReportMeta) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


