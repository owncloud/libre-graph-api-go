# \MeDriveApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateDrives**](MeDriveApi.md#MeCreateDrives) | **Post** /me/drives | Create new navigation property to drives for me
[**MeListDrives**](MeDriveApi.md#MeListDrives) | **Get** /me/drives | Get drives from me



## MeCreateDrives

> Drive MeCreateDrives(ctx).Drive(drive).Execute()

Create new navigation property to drives for me

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
    drive := *openapiclient.NewDrive() // Drive | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDriveApi.MeCreateDrives(context.Background()).Drive(drive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.MeCreateDrives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateDrives`: Drive
    fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.MeCreateDrives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **drive** | [**Drive**](Drive.md) | New navigation property | 

### Return type

[**Drive**](Drive.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListDrives

> CollectionOfDrive MeListDrives(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Select_(select_).Expand(expand).Execute()

Get drives from me

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
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDriveApi.MeListDrives(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.MeListDrives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListDrives`: CollectionOfDrive
    fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.MeListDrives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDrive**](CollectionOfDrive.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

