# \MeDriveRootApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDriveGetRoot**](MeDriveRootApi.md#MeDriveGetRoot) | **Get** /me/drive/root | Get root from home drive



## MeDriveGetRoot

> DriveItem MeDriveGetRoot(ctx).Execute()

Get root from home drive

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDriveRootApi.MeDriveGetRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDriveRootApi.MeDriveGetRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeDriveGetRoot`: DriveItem
    fmt.Fprintf(os.Stdout, "Response from `MeDriveRootApi.MeDriveGetRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMeDriveGetRootRequest struct via the builder pattern


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

