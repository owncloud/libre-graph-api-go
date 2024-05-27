# \MeDriveApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHome**](MeDriveApi.md#GetHome) | **Get** /v1.0/me/drive | Get personal space for user
[**ListSharedByMe**](MeDriveApi.md#ListSharedByMe) | **Get** /v1beta1/me/drive/sharedByMe | Get a list of driveItem objects shared by the current user.
[**ListSharedWithMe**](MeDriveApi.md#ListSharedWithMe) | **Get** /v1beta1/me/drive/sharedWithMe | Get a list of driveItem objects shared with the owner of a drive.



## GetHome

> Drive GetHome(ctx).Execute()

Get personal space for user

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
	resp, r, err := apiClient.MeDriveApi.GetHome(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.GetHome``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHome`: Drive
	fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.GetHome`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHomeRequest struct via the builder pattern


### Return type

[**Drive**](Drive.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedByMe

> CollectionOfDriveItems1 ListSharedByMe(ctx).Execute()

Get a list of driveItem objects shared by the current user.



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
	resp, r, err := apiClient.MeDriveApi.ListSharedByMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.ListSharedByMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharedByMe`: CollectionOfDriveItems1
	fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.ListSharedByMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSharedByMeRequest struct via the builder pattern


### Return type

[**CollectionOfDriveItems1**](CollectionOfDriveItems1.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedWithMe

> CollectionOfDriveItems1 ListSharedWithMe(ctx).Execute()

Get a list of driveItem objects shared with the owner of a drive.



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
	resp, r, err := apiClient.MeDriveApi.ListSharedWithMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.ListSharedWithMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharedWithMe`: CollectionOfDriveItems1
	fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.ListSharedWithMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSharedWithMeRequest struct via the builder pattern


### Return type

[**CollectionOfDriveItems1**](CollectionOfDriveItems1.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

