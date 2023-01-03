# \DefaultApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works/graph/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsApplicationIdGet**](DefaultApi.md#ApplicationsApplicationIdGet) | **Get** /applications/{application-id} | 



## ApplicationsApplicationIdGet

> Application ApplicationsApplicationIdGet(ctx, applicationId).Execute()





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
    resp, r, err := apiClient.DefaultApi.ApplicationsApplicationIdGet(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApplicationsApplicationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsApplicationIdGet`: Application
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApplicationsApplicationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationIdGetRequest struct via the builder pattern


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

