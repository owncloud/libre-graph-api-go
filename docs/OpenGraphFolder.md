# OpenGraphFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildCount** | Pointer to **int32** | Number of children contained immediately within this container. | [optional] 
**View** | Pointer to [**OpenGraphFolderView**](open.graph.folderView.md) | A collection of properties defining the recommended view for the folder. | [optional] 

## Methods

### NewOpenGraphFolder

`func NewOpenGraphFolder() *OpenGraphFolder`

NewOpenGraphFolder instantiates a new OpenGraphFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphFolderWithDefaults

`func NewOpenGraphFolderWithDefaults() *OpenGraphFolder`

NewOpenGraphFolderWithDefaults instantiates a new OpenGraphFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildCount

`func (o *OpenGraphFolder) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *OpenGraphFolder) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *OpenGraphFolder) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *OpenGraphFolder) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### GetView

`func (o *OpenGraphFolder) GetView() OpenGraphFolderView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *OpenGraphFolder) GetViewOk() (*OpenGraphFolderView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *OpenGraphFolder) SetView(v OpenGraphFolderView)`

SetView sets View field to given value.

### HasView

`func (o *OpenGraphFolder) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


