# \EducationClassTeachersApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeacherToClass**](EducationClassTeachersApi.md#AddTeacherToClass) | **Post** /education/classes/{class-id}/teachers/$ref | Assign a teacher to a class
[**DeleteTeacherFromClass**](EducationClassTeachersApi.md#DeleteTeacherFromClass) | **Delete** /education/classes/{class-id}/teachers/{user-id}/$ref | Unassign user as teacher of a class
[**GetTeachers**](EducationClassTeachersApi.md#GetTeachers) | **Get** /education/classes/{class-id}/teachers | Get the teachers for a class



## AddTeacherToClass

> AddTeacherToClass(ctx, classId).ClassTeacherReference(classTeacherReference).Execute()

Assign a teacher to a class

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
    classId := "86948e45-96a6-43df-b83d-46e92afd30de" // string | key: id or externalId of class
    classTeacherReference := *openapiclient.NewClassTeacherReference() // ClassTeacherReference | educationUser to be added as teacher

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationClassTeachersApi.AddTeacherToClass(context.Background(), classId).ClassTeacherReference(classTeacherReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationClassTeachersApi.AddTeacherToClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**classId** | **string** | key: id or externalId of class | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTeacherToClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classTeacherReference** | [**ClassTeacherReference**](ClassTeacherReference.md) | educationUser to be added as teacher | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeacherFromClass

> DeleteTeacherFromClass(ctx, classId, userId).Execute()

Unassign user as teacher of a class

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
    classId := "classId_example" // string | key: id or externalId of class
    userId := "90eedea1-dea1-90ee-a1de-ee90a1deee90" // string | key: id or username of the user to unassign as teacher

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationClassTeachersApi.DeleteTeacherFromClass(context.Background(), classId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationClassTeachersApi.DeleteTeacherFromClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**classId** | **string** | key: id or externalId of class | 
**userId** | **string** | key: id or username of the user to unassign as teacher | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeacherFromClassRequest struct via the builder pattern


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


## GetTeachers

> CollectionOfEducationUser GetTeachers(ctx, classId).Execute()

Get the teachers for a class

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
    classId := "86948e45-96a6-43df-b83d-46e92afd30de" // string | key: id or externalId of class

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EducationClassTeachersApi.GetTeachers(context.Background(), classId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationClassTeachersApi.GetTeachers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeachers`: CollectionOfEducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationClassTeachersApi.GetTeachers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**classId** | **string** | key: id or externalId of class | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeachersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

