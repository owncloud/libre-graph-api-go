# \GroupsApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /v1.0/groups | Add new entity to groups
[**ListGroups**](GroupsApi.md#ListGroups) | **Get** /v1.0/groups | Get entities from groups



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
	openapiclient "github.com/owncloud/libre-graph-api-go"
)

func main() {
	group := *openapiclient.NewGroup() // Group | New entity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsApi.CreateGroup(context.Background()).Group(group).Execute()
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

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> CollectionOfGroup ListGroups(ctx).Search(search).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from groups

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
	search := "search_example" // string | Search items by search phrases (optional)
	orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
	select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
	expand := []string{"Expand_example"} // []string | Expand related entities (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsApi.ListGroups(context.Background()).Search(search).Orderby(orderby).Select_(select_).Expand(expand).Execute()
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
 **search** | **string** | Search items by search phrases | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfGroup**](CollectionOfGroup.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

