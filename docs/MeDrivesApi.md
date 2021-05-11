# \MeDrivesApi

All URIs are relative to *https://owncloud.dev/open-graph-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDrivesDeleteDrive**](MeDrivesApi.md#MeDrivesDeleteDrive) | **Delete** /me/drives/{drive-id} | Delete a specific navigation property drive
[**MeDrivesGetDrive**](MeDrivesApi.md#MeDrivesGetDrive) | **Get** /me/drives/{drive-id} | Get drive by id
[**MeDrivesUpdateDrive**](MeDrivesApi.md#MeDrivesUpdateDrive) | **Patch** /me/drives/{drive-id} | Update the navigation property of a specific drive



## MeDrivesDeleteDrive

> MeDrivesDeleteDrive(ctx, driveId).IfMatch(ifMatch).Execute()

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
    resp, r, err := api_client.MeDrivesApi.MeDrivesDeleteDrive(context.Background(), driveId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDrivesApi.MeDrivesDeleteDrive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeDrivesDeleteDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDrivesGetDrive

> OpenGraphDrive MeDrivesGetDrive(ctx, driveId).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.MeDrivesApi.MeDrivesGetDrive(context.Background(), driveId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDrivesApi.MeDrivesGetDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeDrivesGetDrive`: OpenGraphDrive
    fmt.Fprintf(os.Stdout, "Response from `MeDrivesApi.MeDrivesGetDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDrivesGetDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**OpenGraphDrive**](OpenGraphDrive.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDrivesUpdateDrive

> MeDrivesUpdateDrive(ctx, driveId).OpenGraphDrive(openGraphDrive).Execute()

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
    openGraphDrive := *openapiclient.NewOpenGraphDrive() // OpenGraphDrive | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDrivesApi.MeDrivesUpdateDrive(context.Background(), driveId).OpenGraphDrive(openGraphDrive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDrivesApi.MeDrivesUpdateDrive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeDrivesUpdateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openGraphDrive** | [**OpenGraphDrive**](OpenGraphDrive.md) | New navigation property values | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

