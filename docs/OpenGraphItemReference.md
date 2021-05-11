# OpenGraphItemReference

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


