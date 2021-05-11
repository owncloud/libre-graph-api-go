# OpenGraphIdentitySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**OpenGraphIdentity**](open.graph.identity.md) | Optional. The application associated with this action. | [optional] 
**Device** | Pointer to [**OpenGraphIdentity**](open.graph.identity.md) | Optional. The device associated with this action. | [optional] 
**User** | Pointer to [**OpenGraphIdentity**](open.graph.identity.md) | Optional. The user associated with this action. | [optional] 

## Methods

### NewOpenGraphIdentitySet

`func NewOpenGraphIdentitySet() *OpenGraphIdentitySet`

NewOpenGraphIdentitySet instantiates a new OpenGraphIdentitySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphIdentitySetWithDefaults

`func NewOpenGraphIdentitySetWithDefaults() *OpenGraphIdentitySet`

NewOpenGraphIdentitySetWithDefaults instantiates a new OpenGraphIdentitySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *OpenGraphIdentitySet) GetApplication() OpenGraphIdentity`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *OpenGraphIdentitySet) GetApplicationOk() (*OpenGraphIdentity, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *OpenGraphIdentitySet) SetApplication(v OpenGraphIdentity)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *OpenGraphIdentitySet) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDevice

`func (o *OpenGraphIdentitySet) GetDevice() OpenGraphIdentity`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *OpenGraphIdentitySet) GetDeviceOk() (*OpenGraphIdentity, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *OpenGraphIdentitySet) SetDevice(v OpenGraphIdentity)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *OpenGraphIdentitySet) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetUser

`func (o *OpenGraphIdentitySet) GetUser() OpenGraphIdentity`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OpenGraphIdentitySet) GetUserOk() (*OpenGraphIdentity, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OpenGraphIdentitySet) SetUser(v OpenGraphIdentity)`

SetUser sets User field to given value.

### HasUser

`func (o *OpenGraphIdentitySet) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


