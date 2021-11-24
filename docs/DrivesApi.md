# \DrivesApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesCreateDrive**](DrivesApi.md#DrivesCreateDrive) | **Post** /drives/{drive-id} | Create a new navigation property of a specific drive
[**DrivesDeleteDrive**](DrivesApi.md#DrivesDeleteDrive) | **Delete** /drives/{drive-id} | Delete a specific navigation property drive
[**DrivesGetDrive**](DrivesApi.md#DrivesGetDrive) | **Get** /drives/{drive-id} | Get drive by id
[**DrivesUpdateDrive**](DrivesApi.md#DrivesUpdateDrive) | **Patch** /drives/{drive-id} | Update the navigation property of a specific drive



## DrivesCreateDrive

> DrivesCreateDrive(ctx).Drive(drive).Execute()

Create a new navigation property of a specific drive

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
    drive := *openapiclient.NewDrive() // Drive | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesApi.DrivesCreateDrive(context.Background()).Drive(drive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.DrivesCreateDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDrivesCreateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **drive** | [**Drive**](Drive.md) | New navigation property values | 

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


## DrivesDeleteDrive

> DrivesDeleteDrive(ctx, driveId).IfMatch(ifMatch).Execute()

Delete a specific navigation property drive

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
    driveId := "driveId_example" // string | key: id of drive
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesApi.DrivesDeleteDrive(context.Background(), driveId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.DrivesDeleteDrive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDrivesDeleteDriveRequest struct via the builder pattern


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


## DrivesGetDrive

> Drive DrivesGetDrive(ctx, driveId).Select_(select_).Expand(expand).Execute()

Get drive by id

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
    driveId := "driveId_example" // string | key: id of drive
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesApi.DrivesGetDrive(context.Background(), driveId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.DrivesGetDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesGetDrive`: Drive
    fmt.Fprintf(os.Stdout, "Response from `DrivesApi.DrivesGetDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesGetDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**Drive**](Drive.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesUpdateDrive

> DrivesUpdateDrive(ctx, driveId).Drive(drive).Execute()

Update the navigation property of a specific drive

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
    driveId := "driveId_example" // string | key: id of drive
    drive := *openapiclient.NewDrive() // Drive | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesApi.DrivesUpdateDrive(context.Background(), driveId).Drive(drive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesApi.DrivesUpdateDrive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDrivesUpdateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **drive** | [**Drive**](Drive.md) | New navigation property values | 

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

