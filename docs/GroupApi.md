# \GroupApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMember**](GroupApi.md#AddMember) | **Post** /v1.0/groups/{group-id}/members/$ref | Add a member to a group
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /v1.0/groups/{group-id} | Delete entity from groups
[**DeleteMember**](GroupApi.md#DeleteMember) | **Delete** /v1.0/groups/{group-id}/members/{directory-object-id}/$ref | Delete member from a group
[**GetGroup**](GroupApi.md#GetGroup) | **Get** /v1.0/groups/{group-id} | Get entity from groups by key
[**ListMembers**](GroupApi.md#ListMembers) | **Get** /v1.0/groups/{group-id}/members | Get a list of the group&#39;s direct members
[**UpdateGroup**](GroupApi.md#UpdateGroup) | **Patch** /v1.0/groups/{group-id} | Update entity in groups



## AddMember

> AddMember(ctx, groupId).MemberReference(memberReference).Execute()

Add a member to a group

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
	groupId := "groupId_example" // string | key: id of group
	memberReference := *openapiclient.NewMemberReference() // MemberReference | Object to be added as member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupApi.AddMember(context.Background(), groupId).MemberReference(memberReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.AddMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberReference** | [**MemberReference**](MemberReference.md) | Object to be added as member | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).IfMatch(ifMatch).Execute()

Delete entity from groups

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
	groupId := "groupId_example" // string | key: id of group
	ifMatch := "ifMatch_example" // string | ETag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupApi.DeleteGroup(context.Background(), groupId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMember

> DeleteMember(ctx, groupId, directoryObjectId).IfMatch(ifMatch).Execute()

Delete member from a group

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
	groupId := "groupId_example" // string | key: id of group
	directoryObjectId := "directoryObjectId_example" // string | key: id of group member to remove
	ifMatch := "ifMatch_example" // string | ETag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupApi.DeleteMember(context.Background(), groupId, directoryObjectId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**directoryObjectId** | **string** | key: id of group member to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> Group GetGroup(ctx, groupId).Select_(select_).Expand(expand).Execute()

Get entity from groups by key

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
	groupId := "groupId_example" // string | key: id or name of group
	select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
	expand := []string{"Expand_example"} // []string | Expand related entities (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupApi.GetGroup(context.Background(), groupId).Select_(select_).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id or name of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**Group**](Group.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMembers

> CollectionOfUsers ListMembers(ctx, groupId).Execute()

Get a list of the group's direct members

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
	groupId := "86948e45-96a6-43df-b83d-46e92afd30de" // string | key: id or name of group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupApi.ListMembers(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.ListMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMembers`: CollectionOfUsers
	fmt.Fprintf(os.Stdout, "Response from `GroupApi.ListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id or name of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionOfUsers**](CollectionOfUsers.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> UpdateGroup(ctx, groupId).Group(group).Execute()

Update entity in groups

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
	groupId := "groupId_example" // string | key: id of group
	group := *openapiclient.NewGroup() // Group | New property values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupApi.UpdateGroup(context.Background(), groupId).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.UpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**Group**](Group.md) | New property values | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

