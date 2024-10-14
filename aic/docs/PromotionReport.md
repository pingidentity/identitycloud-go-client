# PromotionReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 
**MissingESVs** | Pointer to **[]string** |  | [optional] 
**PreviouslyIgnoredEncryptedSecrets** | Pointer to **[]string** |  | [optional] 
**Promoter** | Pointer to **string** | The name of the admin user who requested the promotion | [optional] 
**PromotionDescription** | Pointer to **string** | Promoter specified note to make it easier for them to identify what was promoted | [optional] 
**PromotionId** | Pointer to **string** |  | [optional] 
**Report** | Pointer to [**PromotionReportReport**](PromotionReportReport.md) |  | [optional] 
**ReportId** | Pointer to **string** |  | [optional] 
**ReportName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Either promotion or rollback depending on whether the report was for a config promotion or a config rollback | [optional] 

## Methods

### NewPromotionReport

`func NewPromotionReport() *PromotionReport`

NewPromotionReport instantiates a new PromotionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionReportWithDefaults

`func NewPromotionReportWithDefaults() *PromotionReport`

NewPromotionReportWithDefaults instantiates a new PromotionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *PromotionReport) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *PromotionReport) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *PromotionReport) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *PromotionReport) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDryRun

`func (o *PromotionReport) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PromotionReport) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PromotionReport) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PromotionReport) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetMissingESVs

`func (o *PromotionReport) GetMissingESVs() []string`

GetMissingESVs returns the MissingESVs field if non-nil, zero value otherwise.

### GetMissingESVsOk

`func (o *PromotionReport) GetMissingESVsOk() (*[]string, bool)`

GetMissingESVsOk returns a tuple with the MissingESVs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingESVs

`func (o *PromotionReport) SetMissingESVs(v []string)`

SetMissingESVs sets MissingESVs field to given value.

### HasMissingESVs

`func (o *PromotionReport) HasMissingESVs() bool`

HasMissingESVs returns a boolean if a field has been set.

### GetPreviouslyIgnoredEncryptedSecrets

`func (o *PromotionReport) GetPreviouslyIgnoredEncryptedSecrets() []string`

GetPreviouslyIgnoredEncryptedSecrets returns the PreviouslyIgnoredEncryptedSecrets field if non-nil, zero value otherwise.

### GetPreviouslyIgnoredEncryptedSecretsOk

`func (o *PromotionReport) GetPreviouslyIgnoredEncryptedSecretsOk() (*[]string, bool)`

GetPreviouslyIgnoredEncryptedSecretsOk returns a tuple with the PreviouslyIgnoredEncryptedSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviouslyIgnoredEncryptedSecrets

`func (o *PromotionReport) SetPreviouslyIgnoredEncryptedSecrets(v []string)`

SetPreviouslyIgnoredEncryptedSecrets sets PreviouslyIgnoredEncryptedSecrets field to given value.

### HasPreviouslyIgnoredEncryptedSecrets

`func (o *PromotionReport) HasPreviouslyIgnoredEncryptedSecrets() bool`

HasPreviouslyIgnoredEncryptedSecrets returns a boolean if a field has been set.

### GetPromoter

`func (o *PromotionReport) GetPromoter() string`

GetPromoter returns the Promoter field if non-nil, zero value otherwise.

### GetPromoterOk

`func (o *PromotionReport) GetPromoterOk() (*string, bool)`

GetPromoterOk returns a tuple with the Promoter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoter

`func (o *PromotionReport) SetPromoter(v string)`

SetPromoter sets Promoter field to given value.

### HasPromoter

`func (o *PromotionReport) HasPromoter() bool`

HasPromoter returns a boolean if a field has been set.

### GetPromotionDescription

`func (o *PromotionReport) GetPromotionDescription() string`

GetPromotionDescription returns the PromotionDescription field if non-nil, zero value otherwise.

### GetPromotionDescriptionOk

`func (o *PromotionReport) GetPromotionDescriptionOk() (*string, bool)`

GetPromotionDescriptionOk returns a tuple with the PromotionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionDescription

`func (o *PromotionReport) SetPromotionDescription(v string)`

SetPromotionDescription sets PromotionDescription field to given value.

### HasPromotionDescription

`func (o *PromotionReport) HasPromotionDescription() bool`

HasPromotionDescription returns a boolean if a field has been set.

### GetPromotionId

`func (o *PromotionReport) GetPromotionId() string`

GetPromotionId returns the PromotionId field if non-nil, zero value otherwise.

### GetPromotionIdOk

`func (o *PromotionReport) GetPromotionIdOk() (*string, bool)`

GetPromotionIdOk returns a tuple with the PromotionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionId

`func (o *PromotionReport) SetPromotionId(v string)`

SetPromotionId sets PromotionId field to given value.

### HasPromotionId

`func (o *PromotionReport) HasPromotionId() bool`

HasPromotionId returns a boolean if a field has been set.

### GetReport

`func (o *PromotionReport) GetReport() PromotionReportReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *PromotionReport) GetReportOk() (*PromotionReportReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *PromotionReport) SetReport(v PromotionReportReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *PromotionReport) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetReportId

`func (o *PromotionReport) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *PromotionReport) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *PromotionReport) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *PromotionReport) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetReportName

`func (o *PromotionReport) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *PromotionReport) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *PromotionReport) SetReportName(v string)`

SetReportName sets ReportName field to given value.

### HasReportName

`func (o *PromotionReport) HasReportName() bool`

HasReportName returns a boolean if a field has been set.

### GetType

`func (o *PromotionReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromotionReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromotionReport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromotionReport) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


