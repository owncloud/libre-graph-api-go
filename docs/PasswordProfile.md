# PasswordProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceChangePasswordNextSignIn** | Pointer to **bool** | If true the user is required to change their password upon the next login | [optional] [default to false]
**Password** | Pointer to **string** | The user&#39;s password | [optional] 

## Methods

### NewPasswordProfile

`func NewPasswordProfile() *PasswordProfile`

NewPasswordProfile instantiates a new PasswordProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordProfileWithDefaults

`func NewPasswordProfileWithDefaults() *PasswordProfile`

NewPasswordProfileWithDefaults instantiates a new PasswordProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceChangePasswordNextSignIn

`func (o *PasswordProfile) GetForceChangePasswordNextSignIn() bool`

GetForceChangePasswordNextSignIn returns the ForceChangePasswordNextSignIn field if non-nil, zero value otherwise.

### GetForceChangePasswordNextSignInOk

`func (o *PasswordProfile) GetForceChangePasswordNextSignInOk() (*bool, bool)`

GetForceChangePasswordNextSignInOk returns a tuple with the ForceChangePasswordNextSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChangePasswordNextSignIn

`func (o *PasswordProfile) SetForceChangePasswordNextSignIn(v bool)`

SetForceChangePasswordNextSignIn sets ForceChangePasswordNextSignIn field to given value.

### HasForceChangePasswordNextSignIn

`func (o *PasswordProfile) HasForceChangePasswordNextSignIn() bool`

HasForceChangePasswordNextSignIn returns a boolean if a field has been set.

### GetPassword

`func (o *PasswordProfile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordProfile) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordProfile) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PasswordProfile) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


