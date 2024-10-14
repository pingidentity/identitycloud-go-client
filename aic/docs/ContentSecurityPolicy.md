# ContentSecurityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Directives** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewContentSecurityPolicy

`func NewContentSecurityPolicy() *ContentSecurityPolicy`

NewContentSecurityPolicy instantiates a new ContentSecurityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSecurityPolicyWithDefaults

`func NewContentSecurityPolicyWithDefaults() *ContentSecurityPolicy`

NewContentSecurityPolicyWithDefaults instantiates a new ContentSecurityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ContentSecurityPolicy) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ContentSecurityPolicy) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ContentSecurityPolicy) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ContentSecurityPolicy) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDirectives

`func (o *ContentSecurityPolicy) GetDirectives() map[string][]string`

GetDirectives returns the Directives field if non-nil, zero value otherwise.

### GetDirectivesOk

`func (o *ContentSecurityPolicy) GetDirectivesOk() (*map[string][]string, bool)`

GetDirectivesOk returns a tuple with the Directives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectives

`func (o *ContentSecurityPolicy) SetDirectives(v map[string][]string)`

SetDirectives sets Directives field to given value.

### HasDirectives

`func (o *ContentSecurityPolicy) HasDirectives() bool`

HasDirectives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


