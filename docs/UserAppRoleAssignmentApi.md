# \UserAppRoleAssignmentApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCreateAppRoleAssignments**](UserAppRoleAssignmentApi.md#UserCreateAppRoleAssignments) | **Post** /users/{user-id}/appRoleAssignments | Grant an appRoleAssignment to a user
[**UserDeleteAppRoleAssignments**](UserAppRoleAssignmentApi.md#UserDeleteAppRoleAssignments) | **Delete** /users/{user-id}/appRoleAssignments/{appRoleAssignment-id} | Delete the appRoleAssignment from a user
[**UserListAppRoleAssignments**](UserAppRoleAssignmentApi.md#UserListAppRoleAssignments) | **Get** /users/{user-id}/appRoleAssignments | Get appRoleAssignments from a user



## UserCreateAppRoleAssignments

> AppRoleAssignment UserCreateAppRoleAssignments(ctx, userId).AppRoleAssignment(appRoleAssignment).Execute()

Grant an appRoleAssignment to a user



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
    userId := "userId_example" // string | key: id of user
    appRoleAssignment := *openapiclient.NewAppRoleAssignment("AppRoleId_example", "PrincipalId_example", "ResourceId_example") // AppRoleAssignment | New app role assignment value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAppRoleAssignmentApi.UserCreateAppRoleAssignments(context.Background(), userId).AppRoleAssignment(appRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAppRoleAssignmentApi.UserCreateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCreateAppRoleAssignments`: AppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `UserAppRoleAssignmentApi.UserCreateAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateAppRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRoleAssignment** | [**AppRoleAssignment**](AppRoleAssignment.md) | New app role assignment value | 

### Return type

[**AppRoleAssignment**](AppRoleAssignment.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDeleteAppRoleAssignments

> UserDeleteAppRoleAssignments(ctx, userId, appRoleAssignmentId).IfMatch(ifMatch).Execute()

Delete the appRoleAssignment from a user

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
    userId := "userId_example" // string | key: id of user
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment. This is the concatenated {user-id}:{appRole-id} separated by a colon.
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAppRoleAssignmentApi.UserDeleteAppRoleAssignments(context.Background(), userId, appRoleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAppRoleAssignmentApi.UserDeleteAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment. This is the concatenated {user-id}:{appRole-id} separated by a colon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteAppRoleAssignmentsRequest struct via the builder pattern


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


## UserListAppRoleAssignments

> CollectionOfAppRoleAssignments UserListAppRoleAssignments(ctx, userId).Execute()

Get appRoleAssignments from a user



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
    userId := "userId_example" // string | key: id of user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAppRoleAssignmentApi.UserListAppRoleAssignments(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAppRoleAssignmentApi.UserListAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserListAppRoleAssignments`: CollectionOfAppRoleAssignments
    fmt.Fprintf(os.Stdout, "Response from `UserAppRoleAssignmentApi.UserListAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserListAppRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionOfAppRoleAssignments**](CollectionOfAppRoleAssignments.md)

### Authorization

[openId](../README.md#openId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

