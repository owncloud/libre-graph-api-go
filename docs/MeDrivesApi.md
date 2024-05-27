# \MeDrivesApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMyDrives**](MeDrivesApi.md#ListMyDrives) | **Get** /v1.0/me/drives | Get all drives where the current user is a regular member of
[**ListMyDrivesBeta**](MeDrivesApi.md#ListMyDrivesBeta) | **Get** /v1beta1/me/drives | Alias for &#39;/v1.0/drives&#39;, the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles



## ListMyDrives

> CollectionOfDrives ListMyDrives(ctx).Orderby(orderby).Filter(filter).Execute()

Get all drives where the current user is a regular member of

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
	orderby := "lastModifiedDateTime desc" // string | The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc. (optional)
	filter := "driveType eq 'project'" // string | Filter items by property values (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeDrivesApi.ListMyDrives(context.Background()).Orderby(orderby).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDrivesApi.ListMyDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyDrives`: CollectionOfDrives
	fmt.Fprintf(os.Stdout, "Response from `MeDrivesApi.ListMyDrives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc. | 
 **filter** | **string** | Filter items by property values | 

### Return type

[**CollectionOfDrives**](CollectionOfDrives.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyDrivesBeta

> CollectionOfDrives ListMyDrivesBeta(ctx).Orderby(orderby).Filter(filter).Execute()

Alias for '/v1.0/drives', the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles

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
	orderby := "lastModifiedDateTime desc" // string | The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc. (optional)
	filter := "driveType eq 'project'" // string | Filter items by property values (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeDrivesApi.ListMyDrivesBeta(context.Background()).Orderby(orderby).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDrivesApi.ListMyDrivesBeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyDrivesBeta`: CollectionOfDrives
	fmt.Fprintf(os.Stdout, "Response from `MeDrivesApi.ListMyDrivesBeta`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyDrivesBetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc. | 
 **filter** | **string** | Filter items by property values | 

### Return type

[**CollectionOfDrives**](CollectionOfDrives.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

