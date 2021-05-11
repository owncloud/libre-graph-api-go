# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int32** | Optional. Height of the image, in pixels. Read-only. | [optional] [readonly] 
**Width** | Pointer to **int32** | Optional. Width of the image, in pixels. Read-only. | [optional] [readonly] 

## Methods

### NewImage

`func NewImage() *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *Image) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Image) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Image) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Image) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *Image) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Image) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Image) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Image) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


