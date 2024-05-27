# \UserApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /v1.0/users/{user-id} | Delete entity from users
[**ExportPersonalData**](UserApi.md#ExportPersonalData) | **Post** /v1.0/users/{user-id}/exportPersonalData | export personal data of a user
[**GetUser**](UserApi.md#GetUser) | **Get** /v1.0/users/{user-id} | Get entity from users by key
[**UpdateUser**](UserApi.md#UpdateUser) | **Patch** /v1.0/users/{user-id} | Update entity in users



## DeleteUser

> DeleteUser(ctx, userId).IfMatch(ifMatch).Execute()

Delete entity from users

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
	userId := "userId_example" // string | key: id or name of user
	ifMatch := "ifMatch_example" // string | ETag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserApi.DeleteUser(context.Background(), userId).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id or name of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPersonalData

> ExportPersonalData(ctx, userId).ExportPersonalDataRequest(exportPersonalDataRequest).Execute()

export personal data of a user

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
	userId := "userId_example" // string | key: id or name of user
	exportPersonalDataRequest := *openapiclient.NewExportPersonalDataRequest() // ExportPersonalDataRequest | destination the file should be created at (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserApi.ExportPersonalData(context.Background(), userId).ExportPersonalDataRequest(exportPersonalDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ExportPersonalData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id or name of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPersonalDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportPersonalDataRequest** | [**ExportPersonalDataRequest**](ExportPersonalDataRequest.md) | destination the file should be created at | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId).Select_(select_).Expand(expand).Execute()

Get entity from users by key

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
	userId := "userId_example" // string | key: id or name of user
	select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
	expand := []string{"Expand_example"} // []string | Expand related entities (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApi.GetUser(context.Background(), userId).Select_(select_).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id or name of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**User**](User.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId).User(user).Execute()

Update entity in users

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
	user := *openapiclient.NewUser() // User | New property values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApi.UpdateUser(context.Background(), userId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UserApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) | New property values | 

### Return type

[**User**](User.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

