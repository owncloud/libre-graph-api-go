# IdentitySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Device** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**User** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Group** | Pointer to [**Identity**](Identity.md) |  | [optional] 

## Methods

### NewIdentitySet

`func NewIdentitySet() *IdentitySet`

NewIdentitySet instantiates a new IdentitySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySetWithDefaults

`func NewIdentitySetWithDefaults() *IdentitySet`

NewIdentitySetWithDefaults instantiates a new IdentitySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *IdentitySet) GetApplication() Identity`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IdentitySet) GetApplicationOk() (*Identity, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IdentitySet) SetApplication(v Identity)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *IdentitySet) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDevice

`func (o *IdentitySet) GetDevice() Identity`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *IdentitySet) GetDeviceOk() (*Identity, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *IdentitySet) SetDevice(v Identity)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *IdentitySet) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetUser

`func (o *IdentitySet) GetUser() Identity`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IdentitySet) GetUserOk() (*Identity, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IdentitySet) SetUser(v Identity)`

SetUser sets User field to given value.

### HasUser

`func (o *IdentitySet) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetGroup

`func (o *IdentitySet) GetGroup() Identity`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IdentitySet) GetGroupOk() (*Identity, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IdentitySet) SetGroup(v Identity)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IdentitySet) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


