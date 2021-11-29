# \MeDriveRootChildrenApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HomeGetChildren**](MeDriveRootChildrenApi.md#HomeGetChildren) | **Get** /me/drive/root/children | Get children from drive



## HomeGetChildren

> CollectionOfDriveItems HomeGetChildren(ctx).Execute()

Get children from drive

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
    resp, r, err := api_client.MeDriveRootChildrenApi.HomeGetChildren(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDriveRootChildrenApi.HomeGetChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HomeGetChildren`: CollectionOfDriveItems
    fmt.Fprintf(os.Stdout, "Response from `MeDriveRootChildrenApi.HomeGetChildren`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHomeGetChildrenRequest struct via the builder pattern


### Return type

[**CollectionOfDriveItems**](CollectionOfDriveItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

