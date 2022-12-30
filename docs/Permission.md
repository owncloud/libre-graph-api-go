# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantedToIdentities** | Pointer to [**[]IdentitySet**](IdentitySet.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantedToIdentities

`func (o *Permission) GetGrantedToIdentities() []IdentitySet`

GetGrantedToIdentities returns the GrantedToIdentities field if non-nil, zero value otherwise.

### GetGrantedToIdentitiesOk

`func (o *Permission) GetGrantedToIdentitiesOk() (*[]IdentitySet, bool)`

GetGrantedToIdentitiesOk returns a tuple with the GrantedToIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedToIdentities

`func (o *Permission) SetGrantedToIdentities(v []IdentitySet)`

SetGrantedToIdentities sets GrantedToIdentities field to given value.

### HasGrantedToIdentities

`func (o *Permission) HasGrantedToIdentities() bool`

HasGrantedToIdentities returns a boolean if a field has been set.

### GetRoles

`func (o *Permission) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Permission) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Permission) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Permission) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


