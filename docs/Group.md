# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] [readonly] 
**Description** | Pointer to **string** | An optional description for the group. Returned by default. | [optional] 
**DisplayName** | Pointer to **string** | The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $search and $orderBy. | [optional] 
**GroupTypes** | Pointer to **[]string** | Specifies the group types. In MS Graph a group can have multiple types, so this is an array. In libreGraph the possible group types deviate from the MS Graph. The only group type that we currently support is \&quot;ReadOnly\&quot;, which is set for groups that cannot be modified on the current instance. | [optional] 
**Members** | Pointer to [**[]User**](User.md) | Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), Nullable. Supports $expand. | [optional] 
**MembersodataBind** | Pointer to **[]string** | A list of member references to the members to be added. Up to 20 members can be added with a single request | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Group) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Group) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Group) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Group) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGroupTypes

`func (o *Group) GetGroupTypes() []string`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *Group) GetGroupTypesOk() (*[]string, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypes

`func (o *Group) SetGroupTypes(v []string)`

SetGroupTypes sets GroupTypes field to given value.

### HasGroupTypes

`func (o *Group) HasGroupTypes() bool`

HasGroupTypes returns a boolean if a field has been set.

### GetMembers

`func (o *Group) GetMembers() []User`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Group) GetMembersOk() (*[]User, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Group) SetMembers(v []User)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Group) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMembersodataBind

`func (o *Group) GetMembersodataBind() []string`

GetMembersodataBind returns the MembersodataBind field if non-nil, zero value otherwise.

### GetMembersodataBindOk

`func (o *Group) GetMembersodataBindOk() (*[]string, bool)`

GetMembersodataBindOk returns a tuple with the MembersodataBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersodataBind

`func (o *Group) SetMembersodataBind(v []string)`

SetMembersodataBind sets MembersodataBind field to given value.

### HasMembersodataBind

`func (o *Group) HasMembersodataBind() bool`

HasMembersodataBind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


