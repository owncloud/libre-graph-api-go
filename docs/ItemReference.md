# ItemReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveId** | Pointer to **string** | Unique identifier of the drive instance that contains the item. Read-only. | [optional] [readonly] 
**DriveType** | Pointer to **string** | Identifies the type of drive. See [drive][] resource for values. Read-only. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier of the item in the drive. Read-only. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the item being referenced. Read-only. | [optional] [readonly] 
**Path** | Pointer to **string** | Path that can be used to navigate to the item. Read-only. | [optional] [readonly] 
**ShareId** | Pointer to **string** | A unique identifier for a shared resource that can be accessed via the [Shares][] API. | [optional] 

## Methods

### NewItemReference

`func NewItemReference() *ItemReference`

NewItemReference instantiates a new ItemReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemReferenceWithDefaults

`func NewItemReferenceWithDefaults() *ItemReference`

NewItemReferenceWithDefaults instantiates a new ItemReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveId

`func (o *ItemReference) GetDriveId() string`

GetDriveId returns the DriveId field if non-nil, zero value otherwise.

### GetDriveIdOk

`func (o *ItemReference) GetDriveIdOk() (*string, bool)`

GetDriveIdOk returns a tuple with the DriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveId

`func (o *ItemReference) SetDriveId(v string)`

SetDriveId sets DriveId field to given value.

### HasDriveId

`func (o *ItemReference) HasDriveId() bool`

HasDriveId returns a boolean if a field has been set.

### GetDriveType

`func (o *ItemReference) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *ItemReference) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *ItemReference) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *ItemReference) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetId

`func (o *ItemReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ItemReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ItemReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ItemReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ItemReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ItemReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ItemReference) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetShareId

`func (o *ItemReference) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *ItemReference) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *ItemReference) SetShareId(v string)`

SetShareId sets ShareId field to given value.

### HasShareId

`func (o *ItemReference) HasShareId() bool`

HasShareId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


