# \DrivesRootApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDriveItem**](DrivesRootApi.md#CreateDriveItem) | **Post** /v1beta1/drives/{drive-id}/root/children | Create a drive item
[**GetRoot**](DrivesRootApi.md#GetRoot) | **Get** /v1.0/drives/{drive-id}/root | Get root from arbitrary space



## CreateDriveItem

> DriveItem CreateDriveItem(ctx, driveId).DriveItem(driveItem).Execute()

Create a drive item



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
    driveItem := *openapiclient.NewDriveItem() // DriveItem | In the request body, provide a JSON object with the following parameters. For mounting a share the necessary remoteItem id and permission id can be taken from the [sharedWithMe](#/me.drive/ListSharedWithMe) endpoint. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DrivesRootApi.CreateDriveItem(context.Background(), driveId).DriveItem(driveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.CreateDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDriveItem`: DriveItem
    fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.CreateDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **driveItem** | [**DriveItem**](DriveItem.md) | In the request body, provide a JSON object with the following parameters. For mounting a share the necessary remoteItem id and permission id can be taken from the [sharedWithMe](#/me.drive/ListSharedWithMe) endpoint. | 

### Return type

[**DriveItem**](DriveItem.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoot

> DriveItem GetRoot(ctx, driveId).Execute()

Get root from arbitrary space

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
    driveId := "driveId_example" // string | key: id of drive

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DrivesRootApi.GetRoot(context.Background(), driveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.GetRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoot`: DriveItem
    fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.GetRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriveItem**](DriveItem.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

