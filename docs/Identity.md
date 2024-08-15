# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The identity&#39;s display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won&#39;t show up as having changed when using delta. | 
**Id** | Pointer to **string** | Unique identifier for the identity. | [optional] 
**LibreGraphUserType** | Pointer to **string** | The type of the identity. This can be either \&quot;Member\&quot; for regular user, \&quot;Guest\&quot; for guest users or \&quot;Federated\&quot; for users imported from a federated instance. Can be used by clients to indicate the type of user. For more details, clients should look up and cache the user at the /users enpoint. | [optional] 

## Methods

### NewIdentity

`func NewIdentity(displayName string, ) *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *Identity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Identity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Identity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *Identity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Identity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Identity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Identity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLibreGraphUserType

`func (o *Identity) GetLibreGraphUserType() string`

GetLibreGraphUserType returns the LibreGraphUserType field if non-nil, zero value otherwise.

### GetLibreGraphUserTypeOk

`func (o *Identity) GetLibreGraphUserTypeOk() (*string, bool)`

GetLibreGraphUserTypeOk returns a tuple with the LibreGraphUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphUserType

`func (o *Identity) SetLibreGraphUserType(v string)`

SetLibreGraphUserType sets LibreGraphUserType field to given value.

### HasLibreGraphUserType

`func (o *Identity) HasLibreGraphUserType() bool`

HasLibreGraphUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


