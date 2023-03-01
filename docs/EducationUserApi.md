# \EducationUserApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEducationUser**](EducationUserApi.md#CreateEducationUser) | **Post** /education/users | Add new education user
[**DeleteEducationUser**](EducationUserApi.md#DeleteEducationUser) | **Delete** /education/users/{user-id} | Delete educationUser
[**GetEducationUser**](EducationUserApi.md#GetEducationUser) | **Get** /education/users/{user-id} | Get properties of educationUser
[**ListEducationUsers**](EducationUserApi.md#ListEducationUsers) | **Get** /education/users | Get entities from education users
[**UpdateEducationUser**](EducationUserApi.md#UpdateEducationUser) | **Patch** /education/users/{user-id} | Update properties of educationUser



## CreateEducationUser

> EducationUser CreateEducationUser(ctx).EducationUser(educationUser).Execute()

Add new education user

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
    educationUser := *openapiclient.NewEducationUser() // EducationUser | New entity

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationUserApi.CreateEducationUser(context.Background()).EducationUser(educationUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationUserApi.CreateEducationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEducationUser`: EducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationUserApi.CreateEducationUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEducationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **educationUser** | [**EducationUser**](EducationUser.md) | New entity | 

### Return type

[**EducationUser**](EducationUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEducationUser

> DeleteEducationUser(ctx, userId).Execute()

Delete educationUser

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
    userId := "90eedea1-dea1-90ee-a1de-ee90a1deee90" // string | key: id or username of user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationUserApi.DeleteEducationUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationUserApi.DeleteEducationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id or username of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEducationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEducationUser

> EducationUser GetEducationUser(ctx, userId).Expand(expand).Execute()

Get properties of educationUser

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
    userId := "90eedea1-dea1-90ee-a1de-ee90a1deee90" // string | key: id or username of user
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationUserApi.GetEducationUser(context.Background(), userId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationUserApi.GetEducationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEducationUser`: EducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationUserApi.GetEducationUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id or username of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEducationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Expand related entities | 

### Return type

[**EducationUser**](EducationUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEducationUsers

> CollectionOfEducationUser ListEducationUsers(ctx).Orderby(orderby).Expand(expand).Execute()

Get entities from education users

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
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationUserApi.ListEducationUsers(context.Background()).Orderby(orderby).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationUserApi.ListEducationUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEducationUsers`: CollectionOfEducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationUserApi.ListEducationUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEducationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **[]string** | Order items by property values | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfEducationUser**](CollectionOfEducationUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEducationUser

> EducationUser UpdateEducationUser(ctx, userId).EducationUser(educationUser).Execute()

Update properties of educationUser

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
    userId := "90eedea1-dea1-90ee-a1de-ee90a1deee90" // string | key: id or username of user
    educationUser := *openapiclient.NewEducationUser() // EducationUser | New property values

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationUserApi.UpdateEducationUser(context.Background(), userId).EducationUser(educationUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationUserApi.UpdateEducationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEducationUser`: EducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationUserApi.UpdateEducationUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id or username of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEducationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **educationUser** | [**EducationUser**](EducationUser.md) | New property values | 

### Return type

[**EducationUser**](EducationUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

