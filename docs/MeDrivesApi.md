# \MeDrivesApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMyDrives**](MeDrivesApi.md#ListMyDrives) | **Get** /me/drives | Get all drives where the current user is a regular member of



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

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

