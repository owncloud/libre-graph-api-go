# \DriveItemApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDriveItem**](DriveItemApi.md#DeleteDriveItem) | **Delete** /v1beta1/drives/{drive-id}/items/{item-id} | Delete a DriveItem.
[**GetDriveItem**](DriveItemApi.md#GetDriveItem) | **Get** /v1beta1/drives/{drive-id}/items/{item-id} | Get a DriveItem.
[**UpdateDriveItem**](DriveItemApi.md#UpdateDriveItem) | **Patch** /v1beta1/drives/{drive-id}/items/{item-id} | Update a DriveItem.



## DeleteDriveItem

> DeleteDriveItem(ctx, driveId, itemId).Execute()

Delete a DriveItem.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/owncloud/libre-graph-api-go"
)

func main() {
	driveId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668" // string | key: id of drive
	itemId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668!share-id" // string | key: id of item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DriveItemApi.DeleteDriveItem(context.Background(), driveId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.DeleteDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDriveItem

> DriveItem GetDriveItem(ctx, driveId, itemId).Execute()

Get a DriveItem.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/owncloud/libre-graph-api-go"
)

func main() {
	driveId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668" // string | key: id of drive
	itemId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668!share-id" // string | key: id of item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveItemApi.GetDriveItem(context.Background(), driveId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.GetDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveItem`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `DriveItemApi.GetDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DriveItem**](DriveItem.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDriveItem

> DriveItem UpdateDriveItem(ctx, driveId, itemId).DriveItem(driveItem).Execute()

Update a DriveItem.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/owncloud/libre-graph-api-go"
)

func main() {
	driveId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668" // string | key: id of drive
	itemId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668!share-id" // string | key: id of item
	driveItem := *openapiclient.NewDriveItem() // DriveItem | DriveItem properties to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveItemApi.UpdateDriveItem(context.Background(), driveId, itemId).DriveItem(driveItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.UpdateDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDriveItem`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `DriveItemApi.UpdateDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **driveItem** | [**DriveItem**](DriveItem.md) | DriveItem properties to update | 

### Return type

[**DriveItem**](DriveItem.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

