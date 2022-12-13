# \EducationSchoolApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClassToSchool**](EducationSchoolApi.md#AddClassToSchool) | **Post** /education/schools/{school-id}/classes/$ref | Assign a class to a school
[**AddUserToSchool**](EducationSchoolApi.md#AddUserToSchool) | **Post** /education/schools/{school-id}/users/$ref | Assign a user to a school
[**CreateSchool**](EducationSchoolApi.md#CreateSchool) | **Post** /education/schools | Add new school
[**DeleteClassFromSchool**](EducationSchoolApi.md#DeleteClassFromSchool) | **Delete** /education/schools/{school-id}/classes/{class-id}/$ref | Unassign class from a school
[**DeleteSchool**](EducationSchoolApi.md#DeleteSchool) | **Delete** /education/schools/{school-id} | Delete school
[**DeleteUserFromSchool**](EducationSchoolApi.md#DeleteUserFromSchool) | **Delete** /education/schools/{school-id}/users/{user-id}/$ref | Unassign user from a school
[**GetSchool**](EducationSchoolApi.md#GetSchool) | **Get** /education/schools/{school-id} | Get the properties of a specific school
[**ListSchools**](EducationSchoolApi.md#ListSchools) | **Get** /education/schools | Get a list of schools and their properties
[**UpdateSchool**](EducationSchoolApi.md#UpdateSchool) | **Patch** /education/schools/{school-id} | Update properties of a school



## AddClassToSchool

> AddClassToSchool(ctx, schoolId).ClassReference(classReference).Execute()

Assign a class to a school

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
    schoolId := "43b879c4-14c6-4e0a-9b3f-b1b33c5a4bd4" // string | key: id of school
    classReference := *openapiclient.NewClassReference() // ClassReference | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.AddClassToSchool(context.Background(), schoolId).ClassReference(classReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.AddClassToSchool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schoolId** | **string** | key: id of school | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddClassToSchoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classReference** | [**ClassReference**](ClassReference.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserToSchool

> AddUserToSchool(ctx, schoolId).EducationUserReference(educationUserReference).Execute()

Assign a user to a school

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
    schoolId := "43b879c4-14c6-4e0a-9b3f-b1b33c5a4bd4" // string | key: id of school
    educationUserReference := *openapiclient.NewEducationUserReference() // EducationUserReference | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.AddUserToSchool(context.Background(), schoolId).EducationUserReference(educationUserReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.AddUserToSchool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schoolId** | **string** | key: id of school | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToSchoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **educationUserReference** | [**EducationUserReference**](EducationUserReference.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSchool

> EducationSchool CreateSchool(ctx).EducationSchool(educationSchool).Execute()

Add new school

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
    educationSchool := *openapiclient.NewEducationSchool() // EducationSchool | New school

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.CreateSchool(context.Background()).EducationSchool(educationSchool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.CreateSchool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSchool`: EducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationSchoolApi.CreateSchool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSchoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **educationSchool** | [**EducationSchool**](EducationSchool.md) | New school | 

### Return type

[**EducationSchool**](EducationSchool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClassFromSchool

> DeleteClassFromSchool(ctx, schoolId, userId).Execute()

Unassign class from a school

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
    schoolId := "43b879c4-14c6-4e0a-9b3f-b1b33c5a4bd4" // string | key: id of school
    userId := "7e84a069-f374-479b-817d-71590117d443" // string | key: id of the class to unassign from school

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.DeleteClassFromSchool(context.Background(), schoolId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.DeleteClassFromSchool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schoolId** | **string** | key: id of school | 
**userId** | **string** | key: id of the class to unassign from school | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClassFromSchoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSchool

> DeleteSchool(ctx, schoolId).Execute()

Delete school

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
    schoolId := "43b879c4-14c6-4e0a-9b3f-b1b33c5a4bd4" // string | key: id of school

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.DeleteSchool(context.Background(), schoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.DeleteSchool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schoolId** | **string** | key: id of school | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSchoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserFromSchool

> DeleteUserFromSchool(ctx, schoolId, userId).Execute()

Unassign user from a school

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
    schoolId := "43b879c4-14c6-4e0a-9b3f-b1b33c5a4bd4" // string | key: id of school
    userId := "90eedea1-dea1-90ee-a1de-ee90a1deee90" // string | key: id of the user to unassign from school

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.DeleteUserFromSchool(context.Background(), schoolId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.DeleteUserFromSchool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schoolId** | **string** | key: id of school | 
**userId** | **string** | key: id of the user to unassign from school | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserFromSchoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchool

> EducationSchool GetSchool(ctx, schoolId).Execute()

Get the properties of a specific school

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
    schoolId := "43b879c4-14c6-4e0a-9b3f-b1b33c5a4bd4" // string | key: id of school

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.GetSchool(context.Background(), schoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.GetSchool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchool`: EducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationSchoolApi.GetSchool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schoolId** | **string** | key: id of school | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EducationSchool**](EducationSchool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchools

> CollectionOfSchools ListSchools(ctx).Execute()

Get a list of schools and their properties

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.ListSchools(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.ListSchools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchools`: CollectionOfSchools
    fmt.Fprintf(os.Stdout, "Response from `EducationSchoolApi.ListSchools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSchoolsRequest struct via the builder pattern


### Return type

[**CollectionOfSchools**](CollectionOfSchools.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSchool

> EducationSchool UpdateSchool(ctx, schoolId).EducationSchool(educationSchool).Execute()

Update properties of a school

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
    schoolId := "43b879c4-14c6-4e0a-9b3f-b1b33c5a4bd4" // string | key: id of school
    educationSchool := *openapiclient.NewEducationSchool() // EducationSchool | New property values

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationSchoolApi.UpdateSchool(context.Background(), schoolId).EducationSchool(educationSchool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationSchoolApi.UpdateSchool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSchool`: EducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationSchoolApi.UpdateSchool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schoolId** | **string** | key: id of school | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSchoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **educationSchool** | [**EducationSchool**](EducationSchool.md) | New property values | 

### Return type

[**EducationSchool**](EducationSchool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

