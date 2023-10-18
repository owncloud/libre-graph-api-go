# \DrivesGetDrivesAPI

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllDrives**](DrivesGetDrivesAPI.md#ListAllDrives) | **Get** /drives | Get all available drives



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
    resp, r, err := apiClient.DrivesGetDrivesAPI.ListAllDrives(context.Background()).Orderby(orderby).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesGetDrivesAPI.ListAllDrives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllDrives`: CollectionOfDrives1
    fmt.Fprintf(os.Stdout, "Response from `DrivesGetDrivesAPI.ListAllDrives`: %v\n", resp)
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

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

