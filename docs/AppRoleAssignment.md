# AppRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the object. 12345678-9abc-def0-1234-56789abcde. The value of the ID property is often, but not exclusively, in the form of a GUID. The value should be treated as an opaque identifier and not based in being a GUID. Null values are not allowed. Read-only. | [optional] [readonly] 
**DeletedDateTime** | Pointer to **time.Time** |  | [optional] 
**AppRoleId** | **string** | The identifier (id) for the app role which is assigned to the user. Required on create. | 
**CreatedDateTime** | Pointer to **NullableTime** | The time when the app role assignment was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**PrincipalDisplayName** | Pointer to **NullableString** | The display name of the user, group, or service principal that was granted the app role assignment. Read-only. | [optional] 
**PrincipalId** | **NullableString** | The unique identifier (id) for the user, security group, or service principal being granted the app role. Security groups with dynamic memberships are supported. Required on create. | 
**PrincipalType** | Pointer to **NullableString** | The type of the assigned principal. This can either be User, Group, or ServicePrincipal. Read-only. | [optional] 
**ResourceDisplayName** | Pointer to **NullableString** | The display name of the resource app&#39;s service principal to which the assignment is made. | [optional] 
**ResourceId** | **NullableString** | The unique identifier (id) for the resource service principal for which the assignment is made. Required on create. | 

## Methods

### NewAppRoleAssignment

`func NewAppRoleAssignment(appRoleId string, principalId NullableString, resourceId NullableString, ) *AppRoleAssignment`

NewAppRoleAssignment instantiates a new AppRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleAssignmentWithDefaults

`func NewAppRoleAssignmentWithDefaults() *AppRoleAssignment`

NewAppRoleAssignmentWithDefaults instantiates a new AppRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRoleAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppRoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *AppRoleAssignment) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *AppRoleAssignment) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *AppRoleAssignment) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *AppRoleAssignment) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### GetAppRoleId

`func (o *AppRoleAssignment) GetAppRoleId() string`

GetAppRoleId returns the AppRoleId field if non-nil, zero value otherwise.

### GetAppRoleIdOk

`func (o *AppRoleAssignment) GetAppRoleIdOk() (*string, bool)`

GetAppRoleIdOk returns a tuple with the AppRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleId

`func (o *AppRoleAssignment) SetAppRoleId(v string)`

SetAppRoleId sets AppRoleId field to given value.


### GetCreatedDateTime

`func (o *AppRoleAssignment) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AppRoleAssignment) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AppRoleAssignment) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AppRoleAssignment) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *AppRoleAssignment) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *AppRoleAssignment) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetPrincipalDisplayName

`func (o *AppRoleAssignment) GetPrincipalDisplayName() string`

GetPrincipalDisplayName returns the PrincipalDisplayName field if non-nil, zero value otherwise.

### GetPrincipalDisplayNameOk

`func (o *AppRoleAssignment) GetPrincipalDisplayNameOk() (*string, bool)`

GetPrincipalDisplayNameOk returns a tuple with the PrincipalDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalDisplayName

`func (o *AppRoleAssignment) SetPrincipalDisplayName(v string)`

SetPrincipalDisplayName sets PrincipalDisplayName field to given value.

### HasPrincipalDisplayName

`func (o *AppRoleAssignment) HasPrincipalDisplayName() bool`

HasPrincipalDisplayName returns a boolean if a field has been set.

### SetPrincipalDisplayNameNil

`func (o *AppRoleAssignment) SetPrincipalDisplayNameNil(b bool)`

 SetPrincipalDisplayNameNil sets the value for PrincipalDisplayName to be an explicit nil

### UnsetPrincipalDisplayName
`func (o *AppRoleAssignment) UnsetPrincipalDisplayName()`

UnsetPrincipalDisplayName ensures that no value is present for PrincipalDisplayName, not even an explicit nil
### GetPrincipalId

`func (o *AppRoleAssignment) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *AppRoleAssignment) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *AppRoleAssignment) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### SetPrincipalIdNil

`func (o *AppRoleAssignment) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *AppRoleAssignment) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetPrincipalType

`func (o *AppRoleAssignment) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *AppRoleAssignment) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *AppRoleAssignment) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *AppRoleAssignment) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### SetPrincipalTypeNil

`func (o *AppRoleAssignment) SetPrincipalTypeNil(b bool)`

 SetPrincipalTypeNil sets the value for PrincipalType to be an explicit nil

### UnsetPrincipalType
`func (o *AppRoleAssignment) UnsetPrincipalType()`

UnsetPrincipalType ensures that no value is present for PrincipalType, not even an explicit nil
### GetResourceDisplayName

`func (o *AppRoleAssignment) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *AppRoleAssignment) GetResourceDisplayNameOk() (*string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *AppRoleAssignment) SetResourceDisplayName(v string)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *AppRoleAssignment) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### SetResourceDisplayNameNil

`func (o *AppRoleAssignment) SetResourceDisplayNameNil(b bool)`

 SetResourceDisplayNameNil sets the value for ResourceDisplayName to be an explicit nil

### UnsetResourceDisplayName
`func (o *AppRoleAssignment) UnsetResourceDisplayName()`

UnsetResourceDisplayName ensures that no value is present for ResourceDisplayName, not even an explicit nil
### GetResourceId

`func (o *AppRoleAssignment) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AppRoleAssignment) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AppRoleAssignment) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### SetResourceIdNil

`func (o *AppRoleAssignment) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AppRoleAssignment) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


