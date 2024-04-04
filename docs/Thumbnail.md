# Thumbnail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The content stream for the thumbnail. | [optional] 
**Height** | Pointer to **int32** | The height of the thumbnail, in pixels. | [optional] 
**SourceItemId** | Pointer to **string** | The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested. | [optional] 
**Url** | Pointer to **string** | The URL used to fetch the thumbnail content. | [optional] 
**Width** | Pointer to **int32** | The width of the thumbnail, in pixels. | [optional] 

## Methods

### NewThumbnail

`func NewThumbnail() *Thumbnail`

NewThumbnail instantiates a new Thumbnail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbnailWithDefaults

`func NewThumbnailWithDefaults() *Thumbnail`

NewThumbnailWithDefaults instantiates a new Thumbnail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Thumbnail) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Thumbnail) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Thumbnail) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Thumbnail) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetHeight

`func (o *Thumbnail) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Thumbnail) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Thumbnail) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Thumbnail) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetSourceItemId

`func (o *Thumbnail) GetSourceItemId() string`

GetSourceItemId returns the SourceItemId field if non-nil, zero value otherwise.

### GetSourceItemIdOk

`func (o *Thumbnail) GetSourceItemIdOk() (*string, bool)`

GetSourceItemIdOk returns a tuple with the SourceItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceItemId

`func (o *Thumbnail) SetSourceItemId(v string)`

SetSourceItemId sets SourceItemId field to given value.

### HasSourceItemId

`func (o *Thumbnail) HasSourceItemId() bool`

HasSourceItemId returns a boolean if a field has been set.

### GetUrl

`func (o *Thumbnail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Thumbnail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Thumbnail) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Thumbnail) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWidth

`func (o *Thumbnail) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Thumbnail) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Thumbnail) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Thumbnail) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


