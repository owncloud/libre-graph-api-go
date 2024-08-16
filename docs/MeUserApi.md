# \MeUserApi

All URIs are relative to *https://ocis.ocis.rolling.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOwnUser**](MeUserApi.md#GetOwnUser) | **Get** /v1.0/me | Get current user
[**UpdateOwnUser**](MeUserApi.md#UpdateOwnUser) | **Patch** /v1.0/me | Update the current user



## GetOwnUser

> User GetOwnUser(ctx).Expand(expand).Execute()

Get current user

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
	expand := []string{"Expand_example"} // []string | Expand related entities (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeUserApi.GetOwnUser(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeUserApi.GetOwnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnUser`: User
	fmt.Fprintf(os.Stdout, "Response from `MeUserApi.GetOwnUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## UpdateOwnUser

> User UpdateOwnUser(ctx).UserUpdate(userUpdate).Execute()

Update the current user

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
	userUpdate := *openapiclient.NewUserUpdate() // UserUpdate | New user values (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeUserApi.UpdateOwnUser(context.Background()).UserUpdate(userUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeUserApi.UpdateOwnUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOwnUser`: User
	fmt.Fprintf(os.Stdout, "Response from `MeUserApi.UpdateOwnUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userUpdate** | [**UserUpdate**](UserUpdate.md) | New user values | 

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

