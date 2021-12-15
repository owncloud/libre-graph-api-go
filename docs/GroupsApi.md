# \GroupsApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /groups | Add new entity to groups
[**ListGroups**](GroupsApi.md#ListGroups) | **Get** /groups | Get entities from groups



## CreateGroup

> Group CreateGroup(ctx).Group(group).Execute()

Add new entity to groups

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
    group := *openapiclient.NewGroup() // Group | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.CreateGroup(context.Background()).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**Group**](Group.md) | New entity | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> CollectionOfGroup ListGroups(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get entities from groups

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
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.ListGroups(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.ListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroups`: CollectionOfGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.ListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfGroup**](CollectionOfGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

