# SharePointIdentitySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Group** | Pointer to [**Identity**](Identity.md) |  | [optional] 

## Methods

### NewSharePointIdentitySet

`func NewSharePointIdentitySet() *SharePointIdentitySet`

NewSharePointIdentitySet instantiates a new SharePointIdentitySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharePointIdentitySetWithDefaults

`func NewSharePointIdentitySetWithDefaults() *SharePointIdentitySet`

NewSharePointIdentitySetWithDefaults instantiates a new SharePointIdentitySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *SharePointIdentitySet) GetUser() Identity`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SharePointIdentitySet) GetUserOk() (*Identity, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SharePointIdentitySet) SetUser(v Identity)`

SetUser sets User field to given value.

### HasUser

`func (o *SharePointIdentitySet) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetGroup

`func (o *SharePointIdentitySet) GetGroup() Identity`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SharePointIdentitySet) GetGroupOk() (*Identity, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SharePointIdentitySet) SetGroup(v Identity)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SharePointIdentitySet) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


