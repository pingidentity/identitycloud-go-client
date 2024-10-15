# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **string** |  | [optional] 
**NextUpgrade** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewRelease

`func NewRelease() *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *Release) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *Release) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *Release) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *Release) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetNextUpgrade

`func (o *Release) GetNextUpgrade() time.Time`

GetNextUpgrade returns the NextUpgrade field if non-nil, zero value otherwise.

### GetNextUpgradeOk

`func (o *Release) GetNextUpgradeOk() (*time.Time, bool)`

GetNextUpgradeOk returns a tuple with the NextUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpgrade

`func (o *Release) SetNextUpgrade(v time.Time)`

SetNextUpgrade sets NextUpgrade field to given value.

### HasNextUpgrade

`func (o *Release) HasNextUpgrade() bool`

HasNextUpgrade returns a boolean if a field has been set.

### SetNextUpgradeNil

`func (o *Release) SetNextUpgradeNil(b bool)`

 SetNextUpgradeNil sets the value for NextUpgrade to be an explicit nil

### UnsetNextUpgrade
`func (o *Release) UnsetNextUpgrade()`

UnsetNextUpgrade ensures that no value is present for NextUpgrade, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


