# ThumbnailSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID within the item. Read-only. | [optional] 
**Large** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**Medium** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**Small** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**Source** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 

## Methods

### NewThumbnailSet

`func NewThumbnailSet() *ThumbnailSet`

NewThumbnailSet instantiates a new ThumbnailSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbnailSetWithDefaults

`func NewThumbnailSetWithDefaults() *ThumbnailSet`

NewThumbnailSetWithDefaults instantiates a new ThumbnailSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThumbnailSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThumbnailSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThumbnailSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThumbnailSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLarge

`func (o *ThumbnailSet) GetLarge() Thumbnail`

GetLarge returns the Large field if non-nil, zero value otherwise.

### GetLargeOk

`func (o *ThumbnailSet) GetLargeOk() (*Thumbnail, bool)`

GetLargeOk returns a tuple with the Large field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLarge

`func (o *ThumbnailSet) SetLarge(v Thumbnail)`

SetLarge sets Large field to given value.

### HasLarge

`func (o *ThumbnailSet) HasLarge() bool`

HasLarge returns a boolean if a field has been set.

### GetMedium

`func (o *ThumbnailSet) GetMedium() Thumbnail`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ThumbnailSet) GetMediumOk() (*Thumbnail, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ThumbnailSet) SetMedium(v Thumbnail)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *ThumbnailSet) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetSmall

`func (o *ThumbnailSet) GetSmall() Thumbnail`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *ThumbnailSet) GetSmallOk() (*Thumbnail, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *ThumbnailSet) SetSmall(v Thumbnail)`

SetSmall sets Small field to given value.

### HasSmall

`func (o *ThumbnailSet) HasSmall() bool`

HasSmall returns a boolean if a field has been set.

### GetSource

`func (o *ThumbnailSet) GetSource() Thumbnail`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ThumbnailSet) GetSourceOk() (*Thumbnail, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ThumbnailSet) SetSource(v Thumbnail)`

SetSource sets Source field to given value.

### HasSource

`func (o *ThumbnailSet) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


