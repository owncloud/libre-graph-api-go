# SignInActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSuccessfulSignInDateTime** | Pointer to **time.Time** | The date and time of the last successful sign-in for the user. | [optional] 

## Methods

### NewSignInActivity

`func NewSignInActivity() *SignInActivity`

NewSignInActivity instantiates a new SignInActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInActivityWithDefaults

`func NewSignInActivityWithDefaults() *SignInActivity`

NewSignInActivityWithDefaults instantiates a new SignInActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSuccessfulSignInDateTime

`func (o *SignInActivity) GetLastSuccessfulSignInDateTime() time.Time`

GetLastSuccessfulSignInDateTime returns the LastSuccessfulSignInDateTime field if non-nil, zero value otherwise.

### GetLastSuccessfulSignInDateTimeOk

`func (o *SignInActivity) GetLastSuccessfulSignInDateTimeOk() (*time.Time, bool)`

GetLastSuccessfulSignInDateTimeOk returns a tuple with the LastSuccessfulSignInDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulSignInDateTime

`func (o *SignInActivity) SetLastSuccessfulSignInDateTime(v time.Time)`

SetLastSuccessfulSignInDateTime sets LastSuccessfulSignInDateTime field to given value.

### HasLastSuccessfulSignInDateTime

`func (o *SignInActivity) HasLastSuccessfulSignInDateTime() bool`

HasLastSuccessfulSignInDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


