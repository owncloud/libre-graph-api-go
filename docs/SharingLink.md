# SharingLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SharingLinkType**](SharingLinkType.md) |  | [optional] 
**PreventsDownload** | Pointer to **bool** | If &#x60;true&#x60; then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. | [optional] [readonly] 
**WebUrl** | Pointer to **string** | A URL that opens the item in the browser on the website. | [optional] [readonly] 
**LibreGraphDisplayName** | Pointer to **string** | Provides a user-visible display name of the link. Optional. Libregraph only. | [optional] 

## Methods

### NewSharingLink

`func NewSharingLink() *SharingLink`

NewSharingLink instantiates a new SharingLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingLinkWithDefaults

`func NewSharingLinkWithDefaults() *SharingLink`

NewSharingLinkWithDefaults instantiates a new SharingLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SharingLink) GetType() SharingLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SharingLink) GetTypeOk() (*SharingLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SharingLink) SetType(v SharingLinkType)`

SetType sets Type field to given value.

### HasType

`func (o *SharingLink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPreventsDownload

`func (o *SharingLink) GetPreventsDownload() bool`

GetPreventsDownload returns the PreventsDownload field if non-nil, zero value otherwise.

### GetPreventsDownloadOk

`func (o *SharingLink) GetPreventsDownloadOk() (*bool, bool)`

GetPreventsDownloadOk returns a tuple with the PreventsDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventsDownload

`func (o *SharingLink) SetPreventsDownload(v bool)`

SetPreventsDownload sets PreventsDownload field to given value.

### HasPreventsDownload

`func (o *SharingLink) HasPreventsDownload() bool`

HasPreventsDownload returns a boolean if a field has been set.

### GetWebUrl

`func (o *SharingLink) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *SharingLink) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *SharingLink) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *SharingLink) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetLibreGraphDisplayName

`func (o *SharingLink) GetLibreGraphDisplayName() string`

GetLibreGraphDisplayName returns the LibreGraphDisplayName field if non-nil, zero value otherwise.

### GetLibreGraphDisplayNameOk

`func (o *SharingLink) GetLibreGraphDisplayNameOk() (*string, bool)`

GetLibreGraphDisplayNameOk returns a tuple with the LibreGraphDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphDisplayName

`func (o *SharingLink) SetLibreGraphDisplayName(v string)`

SetLibreGraphDisplayName sets LibreGraphDisplayName field to given value.

### HasLibreGraphDisplayName

`func (o *SharingLink) HasLibreGraphDisplayName() bool`

HasLibreGraphDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


