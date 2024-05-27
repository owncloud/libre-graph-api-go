# \MeDriveRootApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HomeGetRoot**](MeDriveRootApi.md#HomeGetRoot) | **Get** /v1.0/me/drive/root | Get root from personal space



## HomeGetRoot

> DriveItem HomeGetRoot(ctx).Execute()

Get root from personal space

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeDriveRootApi.HomeGetRoot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveRootApi.HomeGetRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HomeGetRoot`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `MeDriveRootApi.HomeGetRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHomeGetRootRequest struct via the builder pattern


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

