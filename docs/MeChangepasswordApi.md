# \MeChangepasswordApi

All URIs are relative to *https://ocis.ocis.rolling.owncloud.works/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeOwnPassword**](MeChangepasswordApi.md#ChangeOwnPassword) | **Post** /v1.0/me/changePassword | Chanage your own password



## ChangeOwnPassword

> ChangeOwnPassword(ctx).PasswordChange(passwordChange).Execute()

Chanage your own password

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
	passwordChange := *openapiclient.NewPasswordChange("CurrentPassword_example", "NewPassword_example") // PasswordChange | Password change request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeChangepasswordApi.ChangeOwnPassword(context.Background()).PasswordChange(passwordChange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeChangepasswordApi.ChangeOwnPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeOwnPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordChange** | [**PasswordChange**](PasswordChange.md) | Password change request | 

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

