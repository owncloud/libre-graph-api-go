# AppRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedMemberTypes** | Pointer to **[]string** | Specifies whether this app role can be assigned to users and groups (by setting to [&#39;User&#39;]), to other application&#39;s (by setting to [&#39;Application&#39;], or both (by setting to [&#39;User&#39;, &#39;Application&#39;]). App roles supporting assignment to other applications&#39; service principals are also known as application permissions. The &#39;Application&#39; value is only supported for app roles defined on application entities. | [optional] 
**Description** | Pointer to **NullableString** | The description for the app role. This is displayed when the app role is being assigned and, if the app role functions as an application permission, during  consent experiences. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for the permission that appears in the app role assignment and consent experiences. | [optional] 
**Id** | **string** | Unique role identifier inside the appRoles collection. When creating a new app role, a new GUID identifier must be provided. | 

## Methods

### NewAppRole

`func NewAppRole(id string, ) *AppRole`

NewAppRole instantiates a new AppRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleWithDefaults

`func NewAppRoleWithDefaults() *AppRole`

NewAppRoleWithDefaults instantiates a new AppRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedMemberTypes

`func (o *AppRole) GetAllowedMemberTypes() []string`

GetAllowedMemberTypes returns the AllowedMemberTypes field if non-nil, zero value otherwise.

### GetAllowedMemberTypesOk

`func (o *AppRole) GetAllowedMemberTypesOk() (*[]string, bool)`

GetAllowedMemberTypesOk returns a tuple with the AllowedMemberTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMemberTypes

`func (o *AppRole) SetAllowedMemberTypes(v []string)`

SetAllowedMemberTypes sets AllowedMemberTypes field to given value.

### HasAllowedMemberTypes

`func (o *AppRole) HasAllowedMemberTypes() bool`

HasAllowedMemberTypes returns a boolean if a field has been set.

### GetDescription

`func (o *AppRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *AppRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AppRole) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AppRole) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AppRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AppRole) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AppRole) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *AppRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRole) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


