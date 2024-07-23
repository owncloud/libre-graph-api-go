# \ActivitiesApi

All URIs are relative to *https://ocis.ocis.rolling.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivities**](ActivitiesApi.md#GetActivities) | **Get** /v1beta1/extensions/org.libregraph/activities | Get activities



## GetActivities

> CollectionOfActivities GetActivities(ctx).Kql(kql).Execute()

Get activities

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
	kql := "resourceid:a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668 depth:2" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivitiesApi.GetActivities(context.Background()).Kql(kql).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivitiesApi.GetActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivities`: CollectionOfActivities
	fmt.Fprintf(os.Stdout, "Response from `ActivitiesApi.GetActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kql** | **string** |  | 

### Return type

[**CollectionOfActivities**](CollectionOfActivities.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

