# \DrivesApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrive**](DrivesApi.md#CreateDrive) | **Post** /v1.0/drives | Create a new drive of a specific type
[**DeleteDrive**](DrivesApi.md#DeleteDrive) | **Delete** /v1.0/drives/{drive-id} | Delete a specific space
[**GetDrive**](DrivesApi.md#GetDrive) | **Get** /v1.0/drives/{drive-id} | Get drive by id
[**UpdateDrive**](DrivesApi.md#UpdateDrive) | **Patch** /v1.0/drives/{drive-id} | Update the drive



## CreateDrive

> Drive CreateDrive(ctx).Drive(drive).Execute()

Create a new drive of a specific type

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
	drive := *openapiclient.NewDrive("Name_example") // Drive | New space property values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesApi.CreateDrive(context.Background()).Drive(drive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.CreateDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDrive`: Drive
	fmt.Fprintf(os.Stdout, "Response from `DrivesApi.CreateDrive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **drive** | [**Drive**](Drive.md) | New space property values | 

### Return type

[**Drive**](Drive.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDrive

> DeleteDrive(ctx, driveId).IfMatch(ifMatch).Execute()

Delete a specific space

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
	ifMatch := "ifMatch_example" // string | ETag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DrivesApi.DeleteDrive(context.Background(), driveId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.DeleteDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | ETag | 

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


## GetDrive

> Drive GetDrive(ctx, driveId).Execute()

Get drive by id

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
	resp, r, err := apiClient.DrivesApi.GetDrive(context.Background(), driveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.GetDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDrive`: Drive
	fmt.Fprintf(os.Stdout, "Response from `DrivesApi.GetDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Drive**](Drive.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDrive

> Drive UpdateDrive(ctx, driveId).Drive(drive).Execute()

Update the drive

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
	drive := *openapiclient.NewDrive("Name_example") // Drive | New space values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesApi.UpdateDrive(context.Background(), driveId).Drive(drive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.UpdateDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDrive`: Drive
	fmt.Fprintf(os.Stdout, "Response from `DrivesApi.UpdateDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **drive** | [**Drive**](Drive.md) | New space values | 

### Return type

[**Drive**](Drive.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

