# \DrivesApi

All URIs are relative to *https://ocis.ocis.rolling.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrive**](DrivesApi.md#CreateDrive) | **Post** /v1.0/drives | Create a new drive of a specific type
[**CreateDriveBeta**](DrivesApi.md#CreateDriveBeta) | **Post** /v1beta1/drives | Create a new drive of a specific type. Alias for &#39;/v1.0/drives&#39;, the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles.
[**DeleteDrive**](DrivesApi.md#DeleteDrive) | **Delete** /v1.0/drives/{drive-id} | Delete a specific space
[**DeleteDriveBeta**](DrivesApi.md#DeleteDriveBeta) | **Delete** /v1beta1/drives/{drive-id} | Delete a specific space. Alias for &#39;/v1.0/drives&#39;.
[**GetDrive**](DrivesApi.md#GetDrive) | **Get** /v1.0/drives/{drive-id} | Get drive by id
[**GetDriveBeta**](DrivesApi.md#GetDriveBeta) | **Get** /v1beta1/drives/{drive-id} | Get drive by id. Alias for &#39;/v1.0/drives&#39;, the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles
[**UpdateDrive**](DrivesApi.md#UpdateDrive) | **Patch** /v1.0/drives/{drive-id} | Update the drive
[**UpdateDriveBeta**](DrivesApi.md#UpdateDriveBeta) | **Patch** /v1beta1/drives/{drive-id} | Update the drive. Alias for &#39;/v1.0/drives&#39;, the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles



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


## CreateDriveBeta

> Drive CreateDriveBeta(ctx).Drive(drive).Execute()

Create a new drive of a specific type. Alias for '/v1.0/drives', the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles.

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
	resp, r, err := apiClient.DrivesApi.CreateDriveBeta(context.Background()).Drive(drive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.CreateDriveBeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDriveBeta`: Drive
	fmt.Fprintf(os.Stdout, "Response from `DrivesApi.CreateDriveBeta`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriveBetaRequest struct via the builder pattern


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


## DeleteDriveBeta

> DeleteDriveBeta(ctx, driveId).IfMatch(ifMatch).Execute()

Delete a specific space. Alias for '/v1.0/drives'.

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
	r, err := apiClient.DrivesApi.DeleteDriveBeta(context.Background(), driveId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.DeleteDriveBeta``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDriveBetaRequest struct via the builder pattern


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


## GetDriveBeta

> Drive GetDriveBeta(ctx, driveId).Execute()

Get drive by id. Alias for '/v1.0/drives', the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles

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
	resp, r, err := apiClient.DrivesApi.GetDriveBeta(context.Background(), driveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.GetDriveBeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveBeta`: Drive
	fmt.Fprintf(os.Stdout, "Response from `DrivesApi.GetDriveBeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveBetaRequest struct via the builder pattern


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

> Drive UpdateDrive(ctx, driveId).DriveUpdate(driveUpdate).Execute()

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
	driveUpdate := *openapiclient.NewDriveUpdate() // DriveUpdate | New space values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesApi.UpdateDrive(context.Background(), driveId).DriveUpdate(driveUpdate).Execute()
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

 **driveUpdate** | [**DriveUpdate**](DriveUpdate.md) | New space values | 

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


## UpdateDriveBeta

> Drive UpdateDriveBeta(ctx, driveId).DriveUpdate(driveUpdate).Execute()

Update the drive. Alias for '/v1.0/drives', the difference is that grantedtoV2 is used and roles contain unified roles instead of cs3 roles

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
	driveUpdate := *openapiclient.NewDriveUpdate() // DriveUpdate | New space values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesApi.UpdateDriveBeta(context.Background(), driveId).DriveUpdate(driveUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.UpdateDriveBeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDriveBeta`: Drive
	fmt.Fprintf(os.Stdout, "Response from `DrivesApi.UpdateDriveBeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriveBetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **driveUpdate** | [**DriveUpdate**](DriveUpdate.md) | New space values | 

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

