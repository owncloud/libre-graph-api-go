# \EducationClassApi

All URIs are relative to *https://ocis.ocis.rolling.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToClass**](EducationClassApi.md#AddUserToClass) | **Post** /v1.0/education/classes/{class-id}/members/$ref | Assign a user to a class
[**CreateClass**](EducationClassApi.md#CreateClass) | **Post** /v1.0/education/classes | Add new education class
[**DeleteClass**](EducationClassApi.md#DeleteClass) | **Delete** /v1.0/education/classes/{class-id} | Delete education class
[**DeleteUserFromClass**](EducationClassApi.md#DeleteUserFromClass) | **Delete** /v1.0/education/classes/{class-id}/members/{user-id}/$ref | Unassign user from a class
[**GetClass**](EducationClassApi.md#GetClass) | **Get** /v1.0/education/classes/{class-id} | Get class by key
[**ListClassMembers**](EducationClassApi.md#ListClassMembers) | **Get** /v1.0/education/classes/{class-id}/members | Get the educationClass resources owned by an educationSchool
[**ListClasses**](EducationClassApi.md#ListClasses) | **Get** /v1.0/education/classes | list education classes
[**UpdateClass**](EducationClassApi.md#UpdateClass) | **Patch** /v1.0/education/classes/{class-id} | Update properties of a education class



## AddUserToClass

> AddUserToClass(ctx, classId).ClassMemberReference(classMemberReference).Execute()

Assign a user to a class

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
	classId := "86948e45-96a6-43df-b83d-46e92afd30de" // string | key: id or externalId of class
	classMemberReference := *openapiclient.NewClassMemberReference() // ClassMemberReference | educationUser to be added as member

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EducationClassApi.AddUserToClass(context.Background(), classId).ClassMemberReference(classMemberReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducationClassApi.AddUserToClass``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddUserToClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classMemberReference** | [**ClassMemberReference**](ClassMemberReference.md) | educationUser to be added as member | 

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


## CreateClass

> EducationClass CreateClass(ctx).EducationClass(educationClass).Execute()

Add new education class

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
	educationClass := *openapiclient.NewEducationClass() // EducationClass | New entity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EducationClassApi.CreateClass(context.Background()).EducationClass(educationClass).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducationClassApi.CreateClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClass`: EducationClass
	fmt.Fprintf(os.Stdout, "Response from `EducationClassApi.CreateClass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **educationClass** | [**EducationClass**](EducationClass.md) | New entity | 

### Return type

[**EducationClass**](EducationClass.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClass

> DeleteClass(ctx, classId).Execute()

Delete education class

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
	classId := "86948e45-96a6-43df-b83d-46e92afd30de" // string | key: id or externalId of class

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EducationClassApi.DeleteClass(context.Background(), classId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducationClassApi.DeleteClass``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteClassRequest struct via the builder pattern


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


## DeleteUserFromClass

> DeleteUserFromClass(ctx, classId, userId).Execute()

Unassign user from a class

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
	classId := "classId_example" // string | key: id or externalId of class
	userId := "90eedea1-dea1-90ee-a1de-ee90a1deee90" // string | key: id or username of the user to unassign from class

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EducationClassApi.DeleteUserFromClass(context.Background(), classId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducationClassApi.DeleteUserFromClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**classId** | **string** | key: id or externalId of class | 
**userId** | **string** | key: id or username of the user to unassign from class | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserFromClassRequest struct via the builder pattern


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


## GetClass

> EducationClass GetClass(ctx, classId).Execute()

Get class by key

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
	classId := "86948e45-96a6-43df-b83d-46e92afd30de" // string | key: id or externalId of class

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EducationClassApi.GetClass(context.Background(), classId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducationClassApi.GetClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClass`: EducationClass
	fmt.Fprintf(os.Stdout, "Response from `EducationClassApi.GetClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**classId** | **string** | key: id or externalId of class | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EducationClass**](EducationClass.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClassMembers

> CollectionOfEducationUser ListClassMembers(ctx, classId).Execute()

Get the educationClass resources owned by an educationSchool

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
	classId := "86948e45-96a6-43df-b83d-46e92afd30de" // string | key: id or externalId of class

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EducationClassApi.ListClassMembers(context.Background(), classId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducationClassApi.ListClassMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClassMembers`: CollectionOfEducationUser
	fmt.Fprintf(os.Stdout, "Response from `EducationClassApi.ListClassMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**classId** | **string** | key: id or externalId of class | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClassMembersRequest struct via the builder pattern


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


## ListClasses

> CollectionOfClass ListClasses(ctx).Execute()

list education classes

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EducationClassApi.ListClasses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducationClassApi.ListClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClasses`: CollectionOfClass
	fmt.Fprintf(os.Stdout, "Response from `EducationClassApi.ListClasses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClassesRequest struct via the builder pattern


### Return type

[**CollectionOfClass**](CollectionOfClass.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClass

> EducationClass UpdateClass(ctx, classId).EducationClass(educationClass).Execute()

Update properties of a education class

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
	classId := "86948e45-96a6-43df-b83d-46e92afd30de" // string | key: id or externalId of class
	educationClass := *openapiclient.NewEducationClass() // EducationClass | New property values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EducationClassApi.UpdateClass(context.Background(), classId).EducationClass(educationClass).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EducationClassApi.UpdateClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClass`: EducationClass
	fmt.Fprintf(os.Stdout, "Response from `EducationClassApi.UpdateClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**classId** | **string** | key: id or externalId of class | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **educationClass** | [**EducationClass**](EducationClass.md) | New property values | 

### Return type

[**EducationClass**](EducationClass.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

