# UnifiedRoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description for the unifiedRoleDefinition. | [optional] 
**DisplayName** | Pointer to **string** | The display name for the unifiedRoleDefinition. Required. Supports $filter (&#x60;eq&#x60;, &#x60;in&#x60;). | [optional] 
**Id** | Pointer to **string** | The unique identifier for the role definition. Key, not nullable, Read-only. Inherited from entity. Supports $filter (&#x60;eq&#x60;, &#x60;in&#x60;). | [optional] 
**RolePermissions** | Pointer to [**[]UnifiedRolePermission**](UnifiedRolePermission.md) | List of permissions included in the role. | [optional] 
**LibreGraphWeight** | Pointer to **int32** | When presenting a list of roles the weight can be used to order them in a meaningful way. Lower weight gets higher precedence. So content with lower weight will come first. If set, weights should be non-zero, as 0 is interpreted as an unset weight.  | [optional] 

## Methods

### NewUnifiedRoleDefinition

`func NewUnifiedRoleDefinition() *UnifiedRoleDefinition`

NewUnifiedRoleDefinition instantiates a new UnifiedRoleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnifiedRoleDefinitionWithDefaults

`func NewUnifiedRoleDefinitionWithDefaults() *UnifiedRoleDefinition`

NewUnifiedRoleDefinitionWithDefaults instantiates a new UnifiedRoleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UnifiedRoleDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnifiedRoleDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnifiedRoleDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnifiedRoleDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *UnifiedRoleDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnifiedRoleDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnifiedRoleDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnifiedRoleDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *UnifiedRoleDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnifiedRoleDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnifiedRoleDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnifiedRoleDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRolePermissions

`func (o *UnifiedRoleDefinition) GetRolePermissions() []UnifiedRolePermission`

GetRolePermissions returns the RolePermissions field if non-nil, zero value otherwise.

### GetRolePermissionsOk

`func (o *UnifiedRoleDefinition) GetRolePermissionsOk() (*[]UnifiedRolePermission, bool)`

GetRolePermissionsOk returns a tuple with the RolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolePermissions

`func (o *UnifiedRoleDefinition) SetRolePermissions(v []UnifiedRolePermission)`

SetRolePermissions sets RolePermissions field to given value.

### HasRolePermissions

`func (o *UnifiedRoleDefinition) HasRolePermissions() bool`

HasRolePermissions returns a boolean if a field has been set.

### GetLibreGraphWeight

`func (o *UnifiedRoleDefinition) GetLibreGraphWeight() int32`

GetLibreGraphWeight returns the LibreGraphWeight field if non-nil, zero value otherwise.

### GetLibreGraphWeightOk

`func (o *UnifiedRoleDefinition) GetLibreGraphWeightOk() (*int32, bool)`

GetLibreGraphWeightOk returns a tuple with the LibreGraphWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphWeight

`func (o *UnifiedRoleDefinition) SetLibreGraphWeight(v int32)`

SetLibreGraphWeight sets LibreGraphWeight field to given value.

### HasLibreGraphWeight

`func (o *UnifiedRoleDefinition) HasLibreGraphWeight() bool`

HasLibreGraphWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


