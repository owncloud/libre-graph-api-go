/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// DrivesGetDrivesApiService DrivesGetDrivesApi service
type DrivesGetDrivesApiService service

type ApiListAllDrivesRequest struct {
	ctx context.Context
	ApiService *DrivesGetDrivesApiService
	top *int32
	skip *int32
	orderby *string
	filter *string
	count *bool
	select_ *[]string
	expand *[]string
}

// Show only the first n items
func (r ApiListAllDrivesRequest) Top(top int32) ApiListAllDrivesRequest {
	r.top = &top
	return r
}

// Skip the first n items
func (r ApiListAllDrivesRequest) Skip(skip int32) ApiListAllDrivesRequest {
	r.skip = &skip
	return r
}

// The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc.
func (r ApiListAllDrivesRequest) Orderby(orderby string) ApiListAllDrivesRequest {
	r.orderby = &orderby
	return r
}

// Filter items by property values
func (r ApiListAllDrivesRequest) Filter(filter string) ApiListAllDrivesRequest {
	r.filter = &filter
	return r
}

// Include count of items
func (r ApiListAllDrivesRequest) Count(count bool) ApiListAllDrivesRequest {
	r.count = &count
	return r
}

// Select properties to be returned
func (r ApiListAllDrivesRequest) Select_(select_ []string) ApiListAllDrivesRequest {
	r.select_ = &select_
	return r
}

// Expand related entities
func (r ApiListAllDrivesRequest) Expand(expand []string) ApiListAllDrivesRequest {
	r.expand = &expand
	return r
}

func (r ApiListAllDrivesRequest) Execute() (*CollectionOfDrives, *http.Response, error) {
	return r.ApiService.ListAllDrivesExecute(r)
}

/*
ListAllDrives Get All drives

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAllDrivesRequest
*/
func (a *DrivesGetDrivesApiService) ListAllDrives(ctx context.Context) ApiListAllDrivesRequest {
	return ApiListAllDrivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfDrives
func (a *DrivesGetDrivesApiService) ListAllDrivesExecute(r ApiListAllDrivesRequest) (*CollectionOfDrives, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfDrives
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesGetDrivesApiService.ListAllDrives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drives"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.top != nil {
		localVarQueryParams.Add("$top", parameterToString(*r.top, ""))
	}
	if r.skip != nil {
		localVarQueryParams.Add("$skip", parameterToString(*r.skip, ""))
	}
	if r.orderby != nil {
		localVarQueryParams.Add("$orderby", parameterToString(*r.orderby, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("$filter", parameterToString(*r.filter, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("$count", parameterToString(*r.count, ""))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("$select", parameterToString(*r.select_, "csv"))
	}
	if r.expand != nil {
		localVarQueryParams.Add("$expand", parameterToString(*r.expand, "csv"))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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
