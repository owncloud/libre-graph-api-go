# \DefaultApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationApplicationIdGet**](DefaultApi.md#ApplicationApplicationIdGet) | **Get** /application/{application-id} | 



## ApplicationApplicationIdGet

> Application ApplicationApplicationIdGet(ctx, applicationId).Execute()





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
    applicationId := "applicationId_example" // string | key: id of application

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApplicationApplicationIdGet(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApplicationApplicationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationApplicationIdGet`: Application
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApplicationApplicationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationApplicationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

