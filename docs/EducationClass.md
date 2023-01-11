# EducationClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] [readonly] 
**Description** | Pointer to **string** | An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search. | [optional] 
**DisplayName** | **string** | The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | 
**Members** | Pointer to [**[]User**](User.md) | Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), Nullable. Supports $expand. | [optional] 
**MembersodataBind** | Pointer to **[]string** | A list of member references to the members to be added. Up to 20 members can be added with a single request | [optional] 
**Classification** | **string** | Classification of the group, i.e. \&quot;class\&quot; or \&quot;course\&quot; | 
**ExternalId** | Pointer to **string** | An external unique ID for the class | [optional] 

## Methods

### NewEducationClass

`func NewEducationClass(displayName string, classification string, ) *EducationClass`

NewEducationClass instantiates a new EducationClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationClassWithDefaults

`func NewEducationClassWithDefaults() *EducationClass`

NewEducationClassWithDefaults instantiates a new EducationClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EducationClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EducationClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EducationClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EducationClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *EducationClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EducationClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EducationClass) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EducationClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *EducationClass) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EducationClass) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EducationClass) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMembers

`func (o *EducationClass) GetMembers() []User`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *EducationClass) GetMembersOk() (*[]User, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *EducationClass) SetMembers(v []User)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *EducationClass) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMembersodataBind

`func (o *EducationClass) GetMembersodataBind() []string`

GetMembersodataBind returns the MembersodataBind field if non-nil, zero value otherwise.

### GetMembersodataBindOk

`func (o *EducationClass) GetMembersodataBindOk() (*[]string, bool)`

GetMembersodataBindOk returns a tuple with the MembersodataBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersodataBind

`func (o *EducationClass) SetMembersodataBind(v []string)`

SetMembersodataBind sets MembersodataBind field to given value.

### HasMembersodataBind

`func (o *EducationClass) HasMembersodataBind() bool`

HasMembersodataBind returns a boolean if a field has been set.

### GetClassification

`func (o *EducationClass) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *EducationClass) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *EducationClass) SetClassification(v string)`

SetClassification sets Classification field to given value.


### GetExternalId

`func (o *EducationClass) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EducationClass) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EducationClass) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EducationClass) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


