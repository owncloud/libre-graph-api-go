# OpenGraphItemReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveId** | Pointer to **NullableString** | Unique identifier of the drive instance that contains the item. Read-only. | [optional] [readonly] 
**DriveType** | Pointer to **NullableString** | Identifies the type of drive. See [drive][] resource for values. Read-only. | [optional] [readonly] 
**Id** | Pointer to **NullableString** | Unique identifier of the item in the drive. Read-only. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The name of the item being referenced. Read-only. | [optional] [readonly] 
**Path** | Pointer to **NullableString** | Path that can be used to navigate to the item. Read-only. | [optional] [readonly] 
**ShareId** | Pointer to **NullableString** | A unique identifier for a shared resource that can be accessed via the [Shares][] API. | [optional] 

## Methods

### NewOpenGraphItemReference

`func NewOpenGraphItemReference() *OpenGraphItemReference`

NewOpenGraphItemReference instantiates a new OpenGraphItemReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphItemReferenceWithDefaults

`func NewOpenGraphItemReferenceWithDefaults() *OpenGraphItemReference`

NewOpenGraphItemReferenceWithDefaults instantiates a new OpenGraphItemReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveId

`func (o *OpenGraphItemReference) GetDriveId() string`

GetDriveId returns the DriveId field if non-nil, zero value otherwise.

### GetDriveIdOk

`func (o *OpenGraphItemReference) GetDriveIdOk() (*string, bool)`

GetDriveIdOk returns a tuple with the DriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveId

`func (o *OpenGraphItemReference) SetDriveId(v string)`

SetDriveId sets DriveId field to given value.

### HasDriveId

`func (o *OpenGraphItemReference) HasDriveId() bool`

HasDriveId returns a boolean if a field has been set.

### SetDriveIdNil

`func (o *OpenGraphItemReference) SetDriveIdNil(b bool)`

 SetDriveIdNil sets the value for DriveId to be an explicit nil

### UnsetDriveId
`func (o *OpenGraphItemReference) UnsetDriveId()`

UnsetDriveId ensures that no value is present for DriveId, not even an explicit nil
### GetDriveType

`func (o *OpenGraphItemReference) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *OpenGraphItemReference) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *OpenGraphItemReference) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *OpenGraphItemReference) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### SetDriveTypeNil

`func (o *OpenGraphItemReference) SetDriveTypeNil(b bool)`

 SetDriveTypeNil sets the value for DriveType to be an explicit nil

### UnsetDriveType
`func (o *OpenGraphItemReference) UnsetDriveType()`

UnsetDriveType ensures that no value is present for DriveType, not even an explicit nil
### GetId

`func (o *OpenGraphItemReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenGraphItemReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenGraphItemReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenGraphItemReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OpenGraphItemReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OpenGraphItemReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *OpenGraphItemReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenGraphItemReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenGraphItemReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenGraphItemReference) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpenGraphItemReference) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpenGraphItemReference) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *OpenGraphItemReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OpenGraphItemReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OpenGraphItemReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *OpenGraphItemReference) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *OpenGraphItemReference) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *OpenGraphItemReference) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetShareId

`func (o *OpenGraphItemReference) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *OpenGraphItemReference) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *OpenGraphItemReference) SetShareId(v string)`

SetShareId sets ShareId field to given value.

### HasShareId

`func (o *OpenGraphItemReference) HasShareId() bool`

HasShareId returns a boolean if a field has been set.

### SetShareIdNil

`func (o *OpenGraphItemReference) SetShareIdNil(b bool)`

 SetShareIdNil sets the value for ShareId to be an explicit nil

### UnsetShareId
`func (o *OpenGraphItemReference) UnsetShareId()`

UnsetShareId ensures that no value is present for ShareId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


