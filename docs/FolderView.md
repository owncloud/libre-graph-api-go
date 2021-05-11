# FolderView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortBy** | Pointer to **string** | The method by which the folder should be sorted. | [optional] 
**SortOrder** | Pointer to **string** | If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending. | [optional] 
**ViewType** | Pointer to **string** | The type of view that should be used to represent the folder. | [optional] 

## Methods

### NewFolderView

`func NewFolderView() *FolderView`

NewFolderView instantiates a new FolderView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderViewWithDefaults

`func NewFolderViewWithDefaults() *FolderView`

NewFolderViewWithDefaults instantiates a new FolderView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortBy

`func (o *FolderView) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *FolderView) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *FolderView) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *FolderView) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSortOrder

`func (o *FolderView) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FolderView) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FolderView) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *FolderView) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetViewType

`func (o *FolderView) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *FolderView) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *FolderView) SetViewType(v string)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *FolderView) HasViewType() bool`

HasViewType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


