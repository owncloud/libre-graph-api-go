# \MeDriveRootApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDriveRootDeleteRoot**](MeDriveRootApi.md#MeDriveRootDeleteRoot) | **Delete** /drive/root | Delete navigation property root for drive
[**MeDriveRootGetRoot**](MeDriveRootApi.md#MeDriveRootGetRoot) | **Get** /drive/root | Get root from drive
[**MeDriveRootUpdateRoot**](MeDriveRootApi.md#MeDriveRootUpdateRoot) | **Patch** /drive/root | Update the navigation property root in drive



## MeDriveRootDeleteRoot

> MeDriveRootDeleteRoot(ctx).IfMatch(ifMatch).Execute()

Delete navigation property root for drive

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDriveRootApi.MeDriveRootDeleteRoot(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDriveRootApi.MeDriveRootDeleteRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDriveRootDeleteRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDriveRootGetRoot

> DriveItem MeDriveRootGetRoot(ctx).Select_(select_).Execute()

Get root from drive

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDriveRootApi.MeDriveRootGetRoot(context.Background()).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDriveRootApi.MeDriveRootGetRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeDriveRootGetRoot`: DriveItem
    fmt.Fprintf(os.Stdout, "Response from `MeDriveRootApi.MeDriveRootGetRoot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDriveRootGetRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**DriveItem**](DriveItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDriveRootUpdateRoot

> MeDriveRootUpdateRoot(ctx).DriveItem(driveItem).Execute()

Update the navigation property root in drive

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    driveItem := *openapiclient.NewDriveItem() // DriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDriveRootApi.MeDriveRootUpdateRoot(context.Background()).DriveItem(driveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDriveRootApi.MeDriveRootUpdateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDriveRootUpdateRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **driveItem** | [**DriveItem**](DriveItem.md) | New navigation property values | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

