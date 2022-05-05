# \MeChangepasswordApi

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeOwnPassword**](MeChangepasswordApi.md#ChangeOwnPassword) | **Post** /me/changePassword | Chanage your own password



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
    openapiclient "./openapi"
)

func main() {
    passwordChange := *openapiclient.NewPasswordChange() // PasswordChange | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MeChangepasswordApi.ChangeOwnPassword(context.Background()).PasswordChange(passwordChange).Execute()
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
 **passwordChange** | [**PasswordChange**](PasswordChange.md) |  | 

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

