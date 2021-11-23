# OpenGraphFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashes** | Pointer to [**Hashes**](Hashes.md) |  | [optional] 
**MimeType** | Pointer to **string** | The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only. | [optional] [readonly] 
**ProcessingMetadata** | Pointer to **bool** |  | [optional] 

## Methods

### NewOpenGraphFile

`func NewOpenGraphFile() *OpenGraphFile`

NewOpenGraphFile instantiates a new OpenGraphFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphFileWithDefaults

`func NewOpenGraphFileWithDefaults() *OpenGraphFile`

NewOpenGraphFileWithDefaults instantiates a new OpenGraphFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashes

`func (o *OpenGraphFile) GetHashes() Hashes`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *OpenGraphFile) GetHashesOk() (*Hashes, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *OpenGraphFile) SetHashes(v Hashes)`

SetHashes sets Hashes field to given value.

### HasHashes

`func (o *OpenGraphFile) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### GetMimeType

`func (o *OpenGraphFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *OpenGraphFile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *OpenGraphFile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *OpenGraphFile) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetProcessingMetadata

`func (o *OpenGraphFile) GetProcessingMetadata() bool`

GetProcessingMetadata returns the ProcessingMetadata field if non-nil, zero value otherwise.

### GetProcessingMetadataOk

`func (o *OpenGraphFile) GetProcessingMetadataOk() (*bool, bool)`

GetProcessingMetadataOk returns a tuple with the ProcessingMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMetadata

`func (o *OpenGraphFile) SetProcessingMetadata(v bool)`

SetProcessingMetadata sets ProcessingMetadata field to given value.

### HasProcessingMetadata

`func (o *OpenGraphFile) HasProcessingMetadata() bool`

HasProcessingMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


