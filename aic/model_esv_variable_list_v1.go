/*
PingOne Advanced Identity Cloud API

## Introduction The PingOne Advanced Identity Cloud REST API lets you manage your Advanced Identity Cloud tenants. The API exposes access management and identity management endpoints, with additional endpoints specific to Advanced Identity Cloud tenant environments.<br /><br /> We are now publishing the API spec in OpenAPI 3.0. For the legacy Swagger 2.0 spec, please download [swagger.yaml](swagger.yaml), but note that it may not contain all new functionality.<br /><br /> For full PingOne Advanced Identity Cloud documentation, please visit [the docs website](https://docs.pingidentity.com/pingoneaic/latest/). ## Authenticating to the API The PingOne Advanced Identity Cloud REST API has two different authentication methods:   - [API key and secret](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-with-api-key-and-secret.html): used for tenant read-only operations  - [Access token](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-with-access-token.html): used for access management operations, identity management operations or tenant write operations  For a summary of how to use these authentication methods, refer to [Authenticate to Advanced Identity Cloud REST API](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-overview.html).

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aic

import (
	"encoding/json"
)

// checks if the EsvVariableListV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsvVariableListV1{}

// EsvVariableListV1 Variable list response when requested with empty Accept-API-Version or Accept-API-Version: resource=1.0
type EsvVariableListV1 struct {
	PagedResultsCookie      NullableString        `json:"pagedResultsCookie"`
	RemainingPagedResults   int64                 `json:"remainingPagedResults"`
	Result                  []EsvVariableResponse `json:"result"`
	ResultCount             int64                 `json:"resultCount"`
	TotalPagedResults       int64                 `json:"totalPagedResults"`
	TotalPagedResultsPolicy string                `json:"totalPagedResultsPolicy"`
}

// NewEsvVariableListV1 instantiates a new EsvVariableListV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsvVariableListV1(pagedResultsCookie NullableString, remainingPagedResults int64, result []EsvVariableResponse, resultCount int64, totalPagedResults int64, totalPagedResultsPolicy string) *EsvVariableListV1 {
	this := EsvVariableListV1{}
	this.PagedResultsCookie = pagedResultsCookie
	this.RemainingPagedResults = remainingPagedResults
	this.Result = result
	this.ResultCount = resultCount
	this.TotalPagedResults = totalPagedResults
	this.TotalPagedResultsPolicy = totalPagedResultsPolicy
	return &this
}

// NewEsvVariableListV1WithDefaults instantiates a new EsvVariableListV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsvVariableListV1WithDefaults() *EsvVariableListV1 {
	this := EsvVariableListV1{}
	return &this
}

// GetPagedResultsCookie returns the PagedResultsCookie field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EsvVariableListV1) GetPagedResultsCookie() string {
	if o == nil || o.PagedResultsCookie.Get() == nil {
		var ret string
		return ret
	}

	return *o.PagedResultsCookie.Get()
}

// GetPagedResultsCookieOk returns a tuple with the PagedResultsCookie field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsvVariableListV1) GetPagedResultsCookieOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PagedResultsCookie.Get(), o.PagedResultsCookie.IsSet()
}

// SetPagedResultsCookie sets field value
func (o *EsvVariableListV1) SetPagedResultsCookie(v string) {
	o.PagedResultsCookie.Set(&v)
}

// GetRemainingPagedResults returns the RemainingPagedResults field value
func (o *EsvVariableListV1) GetRemainingPagedResults() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RemainingPagedResults
}

// GetRemainingPagedResultsOk returns a tuple with the RemainingPagedResults field value
// and a boolean to check if the value has been set.
func (o *EsvVariableListV1) GetRemainingPagedResultsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingPagedResults, true
}

// SetRemainingPagedResults sets field value
func (o *EsvVariableListV1) SetRemainingPagedResults(v int64) {
	o.RemainingPagedResults = v
}

// GetResult returns the Result field value
func (o *EsvVariableListV1) GetResult() []EsvVariableResponse {
	if o == nil {
		var ret []EsvVariableResponse
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *EsvVariableListV1) GetResultOk() ([]EsvVariableResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *EsvVariableListV1) SetResult(v []EsvVariableResponse) {
	o.Result = v
}

// GetResultCount returns the ResultCount field value
func (o *EsvVariableListV1) GetResultCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ResultCount
}

// GetResultCountOk returns a tuple with the ResultCount field value
// and a boolean to check if the value has been set.
func (o *EsvVariableListV1) GetResultCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCount, true
}

// SetResultCount sets field value
func (o *EsvVariableListV1) SetResultCount(v int64) {
	o.ResultCount = v
}

// GetTotalPagedResults returns the TotalPagedResults field value
func (o *EsvVariableListV1) GetTotalPagedResults() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalPagedResults
}

// GetTotalPagedResultsOk returns a tuple with the TotalPagedResults field value
// and a boolean to check if the value has been set.
func (o *EsvVariableListV1) GetTotalPagedResultsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPagedResults, true
}

// SetTotalPagedResults sets field value
func (o *EsvVariableListV1) SetTotalPagedResults(v int64) {
	o.TotalPagedResults = v
}

// GetTotalPagedResultsPolicy returns the TotalPagedResultsPolicy field value
func (o *EsvVariableListV1) GetTotalPagedResultsPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalPagedResultsPolicy
}

// GetTotalPagedResultsPolicyOk returns a tuple with the TotalPagedResultsPolicy field value
// and a boolean to check if the value has been set.
func (o *EsvVariableListV1) GetTotalPagedResultsPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPagedResultsPolicy, true
}

// SetTotalPagedResultsPolicy sets field value
func (o *EsvVariableListV1) SetTotalPagedResultsPolicy(v string) {
	o.TotalPagedResultsPolicy = v
}

func (o EsvVariableListV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsvVariableListV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagedResultsCookie"] = o.PagedResultsCookie.Get()
	toSerialize["remainingPagedResults"] = o.RemainingPagedResults
	toSerialize["result"] = o.Result
	toSerialize["resultCount"] = o.ResultCount
	toSerialize["totalPagedResults"] = o.TotalPagedResults
	toSerialize["totalPagedResultsPolicy"] = o.TotalPagedResultsPolicy
	return toSerialize, nil
}

type NullableEsvVariableListV1 struct {
	value *EsvVariableListV1
	isSet bool
}

func (v NullableEsvVariableListV1) Get() *EsvVariableListV1 {
	return v.value
}

func (v *NullableEsvVariableListV1) Set(val *EsvVariableListV1) {
	v.value = val
	v.isSet = true
}

func (v NullableEsvVariableListV1) IsSet() bool {
	return v.isSet
}

func (v *NullableEsvVariableListV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsvVariableListV1(val *EsvVariableListV1) *NullableEsvVariableListV1 {
	return &NullableEsvVariableListV1{value: val, isSet: true}
}

func (v NullableEsvVariableListV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsvVariableListV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
