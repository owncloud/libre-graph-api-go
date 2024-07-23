# \DrivesGetDrivesApi

All URIs are relative to *https://ocis.ocis.rolling.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllDrives**](DrivesGetDrivesApi.md#ListAllDrives) | **Get** /v1.0/drives | Get all available drives
[**ListAllDrivesBeta**](DrivesGetDrivesApi.md#ListAllDrivesBeta) | **Get** /v1beta1/drives | Alias for &#39;/v1.0/drives&#39;, the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles



## ListAllDrives

> CollectionOfDrives1 ListAllDrives(ctx).Orderby(orderby).Filter(filter).Execute()

Get all available drives

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
	resp, r, err := apiClient.DrivesGetDrivesApi.ListAllDrives(context.Background()).Orderby(orderby).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesGetDrivesApi.ListAllDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllDrives`: CollectionOfDrives1
	fmt.Fprintf(os.Stdout, "Response from `DrivesGetDrivesApi.ListAllDrives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc. | 
 **filter** | **string** | Filter items by property values | 

### Return type

[**CollectionOfDrives1**](CollectionOfDrives1.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllDrivesBeta

> CollectionOfDrives1 ListAllDrivesBeta(ctx).Orderby(orderby).Filter(filter).Execute()

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
	resp, r, err := apiClient.DrivesGetDrivesApi.ListAllDrivesBeta(context.Background()).Orderby(orderby).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesGetDrivesApi.ListAllDrivesBeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllDrivesBeta`: CollectionOfDrives1
	fmt.Fprintf(os.Stdout, "Response from `DrivesGetDrivesApi.ListAllDrivesBeta`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllDrivesBetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc. | 
 **filter** | **string** | Filter items by property values | 

### Return type

[**CollectionOfDrives1**](CollectionOfDrives1.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

