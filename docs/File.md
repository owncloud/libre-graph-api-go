# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashes** | Pointer to [**Hashes**](hashes.md) | Hashes of the file&#39;s binary content, if available. Read-only. | [optional] [readonly] 
**MimeType** | Pointer to **string** | The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only. | [optional] [readonly] 
**ProcessingMetadata** | Pointer to **bool** |  | [optional] 

## Methods

### NewFile

`func NewFile() *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashes

`func (o *File) GetHashes() Hashes`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *File) GetHashesOk() (*Hashes, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *File) SetHashes(v Hashes)`

SetHashes sets Hashes field to given value.

### HasHashes

`func (o *File) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### GetMimeType

`func (o *File) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *File) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *File) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *File) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetProcessingMetadata

`func (o *File) GetProcessingMetadata() bool`

GetProcessingMetadata returns the ProcessingMetadata field if non-nil, zero value otherwise.

### GetProcessingMetadataOk

`func (o *File) GetProcessingMetadataOk() (*bool, bool)`

GetProcessingMetadataOk returns a tuple with the ProcessingMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMetadata

`func (o *File) SetProcessingMetadata(v bool)`

SetProcessingMetadata sets ProcessingMetadata field to given value.

### HasProcessingMetadata

`func (o *File) HasProcessingMetadata() bool`

HasProcessingMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


