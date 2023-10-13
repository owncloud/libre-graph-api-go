# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the permission among all permissions on the item. Read-only. | [optional] [readonly] 
**HasPassword** | Pointer to **bool** | Indicates whether the password is set for this permission. This property only appears in the response. Optional. Read-only.  | [optional] [readonly] 
**ExpirationDateTime** | Pointer to **time.Time** | An optional expiration date which limits the permission in time. | [optional] 
**GrantedToV2** | Pointer to [**SharePointIdentitySet**](SharePointIdentitySet.md) |  | [optional] 
**Link** | Pointer to [**SharingLink**](SharingLink.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**GrantedToIdentities** | Pointer to [**[]IdentitySet**](IdentitySet.md) | For link type permissions, the details of the identity to whom permission was granted. This could be used to grant access to a an external user that can be identified by email, aka guest accounts. | [optional] 
**LibreGraphPermissionsActions** | Pointer to **[]string** | Use this to create a permission with custom actions. | [optional] 
**UIHidden** | Pointer to **bool** | Properties or facets (see UI.Facet) annotated with this term will not be rendered if the annotation evaluates to true. Users can set this to hide permissons. | [optional] 

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

### GetId

`func (o *Permission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Permission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Permission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Permission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHasPassword

`func (o *Permission) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *Permission) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *Permission) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *Permission) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *Permission) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *Permission) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *Permission) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *Permission) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetGrantedToV2

`func (o *Permission) GetGrantedToV2() SharePointIdentitySet`

GetGrantedToV2 returns the GrantedToV2 field if non-nil, zero value otherwise.

### GetGrantedToV2Ok

`func (o *Permission) GetGrantedToV2Ok() (*SharePointIdentitySet, bool)`

GetGrantedToV2Ok returns a tuple with the GrantedToV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedToV2

`func (o *Permission) SetGrantedToV2(v SharePointIdentitySet)`

SetGrantedToV2 sets GrantedToV2 field to given value.

### HasGrantedToV2

`func (o *Permission) HasGrantedToV2() bool`

HasGrantedToV2 returns a boolean if a field has been set.

### GetLink

`func (o *Permission) GetLink() SharingLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Permission) GetLinkOk() (*SharingLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Permission) SetLink(v SharingLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *Permission) HasLink() bool`

HasLink returns a boolean if a field has been set.

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

### GetLibreGraphPermissionsActions

`func (o *Permission) GetLibreGraphPermissionsActions() []string`

GetLibreGraphPermissionsActions returns the LibreGraphPermissionsActions field if non-nil, zero value otherwise.

### GetLibreGraphPermissionsActionsOk

`func (o *Permission) GetLibreGraphPermissionsActionsOk() (*[]string, bool)`

GetLibreGraphPermissionsActionsOk returns a tuple with the LibreGraphPermissionsActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphPermissionsActions

`func (o *Permission) SetLibreGraphPermissionsActions(v []string)`

SetLibreGraphPermissionsActions sets LibreGraphPermissionsActions field to given value.

### HasLibreGraphPermissionsActions

`func (o *Permission) HasLibreGraphPermissionsActions() bool`

HasLibreGraphPermissionsActions returns a boolean if a field has been set.

### GetUIHidden

`func (o *Permission) GetUIHidden() bool`

GetUIHidden returns the UIHidden field if non-nil, zero value otherwise.

### GetUIHiddenOk

`func (o *Permission) GetUIHiddenOk() (*bool, bool)`

GetUIHiddenOk returns a tuple with the UIHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIHidden

`func (o *Permission) SetUIHidden(v bool)`

SetUIHidden sets UIHidden field to given value.

### HasUIHidden

`func (o *Permission) HasUIHidden() bool`

HasUIHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


