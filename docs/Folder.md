# Folder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildCount** | Pointer to **int32** | Number of children contained immediately within this container. | [optional] 
**View** | Pointer to [**FolderView**](folderView.md) | A collection of properties defining the recommended view for the folder. | [optional] 

## Methods

### NewFolder

`func NewFolder() *Folder`

NewFolder instantiates a new Folder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderWithDefaults

`func NewFolderWithDefaults() *Folder`

NewFolderWithDefaults instantiates a new Folder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildCount

`func (o *Folder) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *Folder) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *Folder) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *Folder) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### GetView

`func (o *Folder) GetView() FolderView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Folder) GetViewOk() (*FolderView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Folder) SetView(v FolderView)`

SetView sets View field to given value.

### HasView

`func (o *Folder) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


