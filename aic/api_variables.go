/*
PingOne Advanced Identity Cloud API

## Introduction The PingOne Advanced Identity Cloud REST API lets you manage your Advanced Identity Cloud tenants. The API exposes access management and identity management endpoints, with additional endpoints specific to Advanced Identity Cloud tenant environments.<br /><br /> We are now publishing the API spec in OpenAPI 3.0. For the legacy Swagger 2.0 spec, please download [swagger.yaml](swagger.yaml), but note that it may not contain all new functionality.<br /><br /> For full PingOne Advanced Identity Cloud documentation, please visit [the docs website](https://docs.pingidentity.com/pingoneaic/latest/). ## Authenticating to the API The PingOne Advanced Identity Cloud REST API has two different authentication methods:   - [API key and secret](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-with-api-key-and-secret.html): used for tenant read-only operations  - [Access token](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-with-access-token.html): used for access management operations, identity management operations or tenant write operations  For a summary of how to use these authentication methods, refer to [Authenticate to Advanced Identity Cloud REST API](https://docs.pingidentity.com/pingoneaic/latest/developer-docs/authenticate-to-rest-api-overview.html).

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aic

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// VariablesAPIService VariablesAPI service
type VariablesAPIService service

type ApiActionVariableRequest struct {
	ctx                      context.Context
	ApiService               *VariablesAPIService
	variableId               string
	action                   *string
	esvSetDescriptionRequest *EsvSetDescriptionRequest
	acceptAPIVersion         *string
}

func (r ApiActionVariableRequest) Action(action string) ApiActionVariableRequest {
	r.action = &action
	return r
}

// The description of this variable
func (r ApiActionVariableRequest) EsvSetDescriptionRequest(esvSetDescriptionRequest EsvSetDescriptionRequest) ApiActionVariableRequest {
	r.esvSetDescriptionRequest = &esvSetDescriptionRequest
	return r
}

// resource&#x3D;2.0
func (r ApiActionVariableRequest) AcceptAPIVersion(acceptAPIVersion string) ApiActionVariableRequest {
	r.acceptAPIVersion = &acceptAPIVersion
	return r
}

func (r ApiActionVariableRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActionVariableExecute(r)
}

/*
ActionVariable Set a variable description

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param variableId ID of the variable
	@return ApiActionVariableRequest
*/
func (a *VariablesAPIService) ActionVariable(ctx context.Context, variableId string) ApiActionVariableRequest {
	return ApiActionVariableRequest{
		ApiService: a,
		ctx:        ctx,
		variableId: variableId,
	}
}

// Execute executes the request
func (a *VariablesAPIService) ActionVariableExecute(r ApiActionVariableRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.ActionVariable")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/variables/{variableId}"
	localVarPath = strings.Replace(localVarPath, "{"+"variableId"+"}", url.PathEscape(parameterValueToString(r.variableId, "variableId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.action == nil {
		return nil, reportError("action is required and must be specified")
	}
	if r.esvSetDescriptionRequest == nil {
		return nil, reportError("esvSetDescriptionRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "_action", r.action, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptAPIVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-API-Version", r.acceptAPIVersion, "")
	}
	// body params
	localVarPostBody = r.esvSetDescriptionRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v EsvError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v EsvError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateVariablesRequest struct {
	ctx                            context.Context
	ApiService                     *VariablesAPIService
	variableId                     string
	esvVariableCreateUpdateRequest *EsvVariableCreateUpdateRequest
	acceptAPIVersion               *string
}

// JSON body of the new variable
func (r ApiCreateVariablesRequest) EsvVariableCreateUpdateRequest(esvVariableCreateUpdateRequest EsvVariableCreateUpdateRequest) ApiCreateVariablesRequest {
	r.esvVariableCreateUpdateRequest = &esvVariableCreateUpdateRequest
	return r
}

// resource&#x3D;2.0
func (r ApiCreateVariablesRequest) AcceptAPIVersion(acceptAPIVersion string) ApiCreateVariablesRequest {
	r.acceptAPIVersion = &acceptAPIVersion
	return r
}

func (r ApiCreateVariablesRequest) Execute() (*EsvVariableResponse, *http.Response, error) {
	return r.ApiService.CreateVariablesExecute(r)
}

/*
CreateVariables Create or update a variable

Create or update a variable using a predefined name. Once created, a variable name cannot be changed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param variableId ID of the variable
	@return ApiCreateVariablesRequest
*/
func (a *VariablesAPIService) CreateVariables(ctx context.Context, variableId string) ApiCreateVariablesRequest {
	return ApiCreateVariablesRequest{
		ApiService: a,
		ctx:        ctx,
		variableId: variableId,
	}
}

// Execute executes the request
//
//	@return EsvVariableResponse
func (a *VariablesAPIService) CreateVariablesExecute(r ApiCreateVariablesRequest) (*EsvVariableResponse, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *EsvVariableResponse
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreateVariablesExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *VariablesAPIService) internalCreateVariablesExecute(r ApiCreateVariablesRequest) (*EsvVariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EsvVariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.CreateVariables")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/variables/{variableId}"
	localVarPath = strings.Replace(localVarPath, "{"+"variableId"+"}", url.PathEscape(parameterValueToString(r.variableId, "variableId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.esvVariableCreateUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("esvVariableCreateUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptAPIVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-API-Version", r.acceptAPIVersion, "")
	}
	// body params
	localVarPostBody = r.esvVariableCreateUpdateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v EsvError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v EsvError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteVariableRequest struct {
	ctx              context.Context
	ApiService       *VariablesAPIService
	variableId       string
	acceptAPIVersion *string
}

// resource&#x3D;2.0
func (r ApiDeleteVariableRequest) AcceptAPIVersion(acceptAPIVersion string) ApiDeleteVariableRequest {
	r.acceptAPIVersion = &acceptAPIVersion
	return r
}

func (r ApiDeleteVariableRequest) Execute() (*EsvVariableResponse, *http.Response, error) {
	return r.ApiService.DeleteVariableExecute(r)
}

/*
DeleteVariable Delete a variable

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param variableId ID of the variable
	@return ApiDeleteVariableRequest
*/
func (a *VariablesAPIService) DeleteVariable(ctx context.Context, variableId string) ApiDeleteVariableRequest {
	return ApiDeleteVariableRequest{
		ApiService: a,
		ctx:        ctx,
		variableId: variableId,
	}
}

// Execute executes the request
//
//	@return EsvVariableResponse
func (a *VariablesAPIService) DeleteVariableExecute(r ApiDeleteVariableRequest) (*EsvVariableResponse, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *EsvVariableResponse
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalDeleteVariableExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *VariablesAPIService) internalDeleteVariableExecute(r ApiDeleteVariableRequest) (*EsvVariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EsvVariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.DeleteVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/variables/{variableId}"
	localVarPath = strings.Replace(localVarPath, "{"+"variableId"+"}", url.PathEscape(parameterValueToString(r.variableId, "variableId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptAPIVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-API-Version", r.acceptAPIVersion, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v EsvError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v EsvError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAllVariablesRequest struct {
	ctx                context.Context
	ApiService         *VariablesAPIService
	acceptAPIVersion   *string
	pageSize           *int64
	pagedResultsCookie *string
	pagedResultsOffset *int64
	onlyPending        *bool
}

// resource&#x3D;2.0
func (r ApiGetAllVariablesRequest) AcceptAPIVersion(acceptAPIVersion string) ApiGetAllVariablesRequest {
	r.acceptAPIVersion = &acceptAPIVersion
	return r
}

// _Accept-API-Version: resource&#x3D;2.0_ only, maximum number of results returned by endpoint before paging
func (r ApiGetAllVariablesRequest) PageSize(pageSize int64) ApiGetAllVariablesRequest {
	r.pageSize = &pageSize
	return r
}

// _Accept-API-Version: resource&#x3D;2.0_ only, opaque data used for paging result data, can be used for paging instead of having to track pageSize and pagedResultsOffset
func (r ApiGetAllVariablesRequest) PagedResultsCookie(pagedResultsCookie string) ApiGetAllVariablesRequest {
	r.pagedResultsCookie = &pagedResultsCookie
	return r
}

// _Accept-API-Version: resource&#x3D;2.0_ only, offset of the first result to be returned by endpoint
func (r ApiGetAllVariablesRequest) PagedResultsOffset(pagedResultsOffset int64) ApiGetAllVariablesRequest {
	r.pagedResultsOffset = &pagedResultsOffset
	return r
}

// _Accept-API-Version: resource&#x3D;2.0_ only, returns ESVs with unapplied changes
func (r ApiGetAllVariablesRequest) OnlyPending(onlyPending bool) ApiGetAllVariablesRequest {
	r.onlyPending = &onlyPending
	return r
}

func (r ApiGetAllVariablesRequest) Execute() (*GetAllVariables200Response, *http.Response, error) {
	return r.ApiService.GetAllVariablesExecute(r)
}

/*
GetAllVariables Get all variables

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllVariablesRequest
*/
func (a *VariablesAPIService) GetAllVariables(ctx context.Context) ApiGetAllVariablesRequest {
	return ApiGetAllVariablesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAllVariables200Response
func (a *VariablesAPIService) GetAllVariablesExecute(r ApiGetAllVariablesRequest) (*GetAllVariables200Response, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *GetAllVariables200Response
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetAllVariablesExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *VariablesAPIService) internalGetAllVariablesExecute(r ApiGetAllVariablesRequest) (*GetAllVariables200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetAllVariables200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.GetAllVariables")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/variables"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "_pageSize", r.pageSize, "")
	} else {
		var defaultValue int64 = 25
		r.pageSize = &defaultValue
	}
	if r.pagedResultsCookie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "_pagedResultsCookie", r.pagedResultsCookie, "")
	} else {
		var defaultValue string = ""
		r.pagedResultsCookie = &defaultValue
	}
	if r.pagedResultsOffset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "_pagedResultsOffset", r.pagedResultsOffset, "")
	} else {
		var defaultValue int64 = 0
		r.pagedResultsOffset = &defaultValue
	}
	if r.onlyPending != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "_onlyPending", r.onlyPending, "")
	} else {
		var defaultValue bool = false
		r.onlyPending = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptAPIVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-API-Version", r.acceptAPIVersion, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v EsvError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v EsvError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetVariableRequest struct {
	ctx              context.Context
	ApiService       *VariablesAPIService
	variableId       string
	acceptAPIVersion *string
}

// resource&#x3D;2.0
func (r ApiGetVariableRequest) AcceptAPIVersion(acceptAPIVersion string) ApiGetVariableRequest {
	r.acceptAPIVersion = &acceptAPIVersion
	return r
}

func (r ApiGetVariableRequest) Execute() (*EsvVariableResponse, *http.Response, error) {
	return r.ApiService.GetVariableExecute(r)
}

/*
GetVariable Get a variable

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param variableId ID of the variable
	@return ApiGetVariableRequest
*/
func (a *VariablesAPIService) GetVariable(ctx context.Context, variableId string) ApiGetVariableRequest {
	return ApiGetVariableRequest{
		ApiService: a,
		ctx:        ctx,
		variableId: variableId,
	}
}

// Execute executes the request
//
//	@return EsvVariableResponse
func (a *VariablesAPIService) GetVariableExecute(r ApiGetVariableRequest) (*EsvVariableResponse, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *EsvVariableResponse
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetVariableExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *VariablesAPIService) internalGetVariableExecute(r ApiGetVariableRequest) (*EsvVariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EsvVariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.GetVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/variables/{variableId}"
	localVarPath = strings.Replace(localVarPath, "{"+"variableId"+"}", url.PathEscape(parameterValueToString(r.variableId, "variableId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptAPIVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-API-Version", r.acceptAPIVersion, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v EsvError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v EsvError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
