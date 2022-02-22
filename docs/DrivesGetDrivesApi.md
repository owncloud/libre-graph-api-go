# \DrivesGetDrivesApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllDrives**](DrivesGetDrivesApi.md#ListAllDrives) | **Get** /drives | Get All drives



## ListAllDrives

> CollectionOfDrives ListAllDrives(ctx).Top(top).Skip(skip).Orderby(orderby).Filter(filter).Count(count).Select_(select_).Expand(expand).Execute()

Get All drives

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
    orderby := "orderby_example" // string | The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc. (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DrivesGetDrivesApi.ListAllDrives(context.Background()).Top(top).Skip(skip).Orderby(orderby).Filter(filter).Count(count).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesGetDrivesApi.ListAllDrives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllDrives`: CollectionOfDrives
    fmt.Fprintf(os.Stdout, "Response from `DrivesGetDrivesApi.ListAllDrives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **orderby** | **string** | The $orderby system query option allows clients to request resources in either ascending order using asc or descending order using desc. | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDrives**](CollectionOfDrives.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

