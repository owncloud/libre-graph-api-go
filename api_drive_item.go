/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DriveItemApiService DriveItemApi service
type DriveItemApiService service

type ApiDeleteDriveItemRequest struct {
	ctx        context.Context
	ApiService *DriveItemApiService
	driveId    string
	itemId     string
}

func (r ApiDeleteDriveItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDriveItemExecute(r)
}

/*
DeleteDriveItem Delete a DriveItem.

Delete a DriveItem by using its ID.

Deleting items using this method moves the items to the recycle bin instead of permanently deleting the item.

Mounted shares in the share jail are unmounted. The `@client.synchronize` property of the `driveItem` in the [sharedWithMe](#/me.drive/ListSharedWithMe) endpoint will change to false.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param driveId key: id of drive
	@param itemId key: id of item
	@return ApiDeleteDriveItemRequest
*/
func (a *DriveItemApiService) DeleteDriveItem(ctx context.Context, driveId string, itemId string) ApiDeleteDriveItemRequest {
	return ApiDeleteDriveItemRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
	}
}

// Execute executes the request
func (a *DriveItemApiService) DeleteDriveItemExecute(r ApiDeleteDriveItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DriveItemApiService.DeleteDriveItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

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
		var v OdataError
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

type ApiGetDriveItemRequest struct {
	ctx        context.Context
	ApiService *DriveItemApiService
	driveId    string
	itemId     string
}

func (r ApiGetDriveItemRequest) Execute() (*DriveItem, *http.Response, error) {
	return r.ApiService.GetDriveItemExecute(r)
}

/*
GetDriveItem Get a DriveItem.

Get a DriveItem by using its ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param driveId key: id of drive
	@param itemId key: id of item
	@return ApiGetDriveItemRequest
*/
func (a *DriveItemApiService) GetDriveItem(ctx context.Context, driveId string, itemId string) ApiGetDriveItemRequest {
	return ApiGetDriveItemRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
	}
}

// Execute executes the request
//
//	@return DriveItem
func (a *DriveItemApiService) GetDriveItemExecute(r ApiGetDriveItemRequest) (*DriveItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DriveItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DriveItemApiService.GetDriveItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

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
		var v OdataError
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

type ApiUpdateDriveItemRequest struct {
	ctx        context.Context
	ApiService *DriveItemApiService
	driveId    string
	itemId     string
	driveItem  *DriveItem
}

// DriveItem properties to update
func (r ApiUpdateDriveItemRequest) DriveItem(driveItem DriveItem) ApiUpdateDriveItemRequest {
	r.driveItem = &driveItem
	return r
}

func (r ApiUpdateDriveItemRequest) Execute() (*DriveItem, *http.Response, error) {
	return r.ApiService.UpdateDriveItemExecute(r)
}

/*
UpdateDriveItem Update a DriveItem.

Update a DriveItem.

The request body must include a JSON object with the properties to update.
Only the properties that are provided will be updated.

Currently it supports updating the following properties:

* `@UI.Hidden` - Hides the item from the UI.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param driveId key: id of drive
	@param itemId key: id of item
	@return ApiUpdateDriveItemRequest
*/
func (a *DriveItemApiService) UpdateDriveItem(ctx context.Context, driveId string, itemId string) ApiUpdateDriveItemRequest {
	return ApiUpdateDriveItemRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
	}
}

// Execute executes the request
//
//	@return DriveItem
func (a *DriveItemApiService) UpdateDriveItemExecute(r ApiUpdateDriveItemRequest) (*DriveItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DriveItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DriveItemApiService.UpdateDriveItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.driveItem == nil {
		return localVarReturnValue, nil, reportError("driveItem is required and must be specified")
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
	// body params
	localVarPostBody = r.driveItem
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
		var v OdataError
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