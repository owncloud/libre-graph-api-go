# OpenGraphHashes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crc32Hash** | Pointer to **NullableString** | The CRC32 value of the file (if available). Read-only. | [optional] 
**QuickXorHash** | Pointer to **NullableString** | A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only. | [optional] 
**Sha1Hash** | Pointer to **NullableString** | SHA1 hash for the contents of the file (if available). Read-only. | [optional] 
**Sha256Hash** | Pointer to **NullableString** | SHA256 hash for the contents of the file (if available). Read-only. | [optional] 

## Methods

### NewOpenGraphHashes

`func NewOpenGraphHashes() *OpenGraphHashes`

NewOpenGraphHashes instantiates a new OpenGraphHashes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphHashesWithDefaults

`func NewOpenGraphHashesWithDefaults() *OpenGraphHashes`

NewOpenGraphHashesWithDefaults instantiates a new OpenGraphHashes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrc32Hash

`func (o *OpenGraphHashes) GetCrc32Hash() string`

GetCrc32Hash returns the Crc32Hash field if non-nil, zero value otherwise.

### GetCrc32HashOk

`func (o *OpenGraphHashes) GetCrc32HashOk() (*string, bool)`

GetCrc32HashOk returns a tuple with the Crc32Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrc32Hash

`func (o *OpenGraphHashes) SetCrc32Hash(v string)`

SetCrc32Hash sets Crc32Hash field to given value.

### HasCrc32Hash

`func (o *OpenGraphHashes) HasCrc32Hash() bool`

HasCrc32Hash returns a boolean if a field has been set.

### SetCrc32HashNil

`func (o *OpenGraphHashes) SetCrc32HashNil(b bool)`

 SetCrc32HashNil sets the value for Crc32Hash to be an explicit nil

### UnsetCrc32Hash
`func (o *OpenGraphHashes) UnsetCrc32Hash()`

UnsetCrc32Hash ensures that no value is present for Crc32Hash, not even an explicit nil
### GetQuickXorHash

`func (o *OpenGraphHashes) GetQuickXorHash() string`

GetQuickXorHash returns the QuickXorHash field if non-nil, zero value otherwise.

### GetQuickXorHashOk

`func (o *OpenGraphHashes) GetQuickXorHashOk() (*string, bool)`

GetQuickXorHashOk returns a tuple with the QuickXorHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickXorHash

`func (o *OpenGraphHashes) SetQuickXorHash(v string)`

SetQuickXorHash sets QuickXorHash field to given value.

### HasQuickXorHash

`func (o *OpenGraphHashes) HasQuickXorHash() bool`

HasQuickXorHash returns a boolean if a field has been set.

### SetQuickXorHashNil

`func (o *OpenGraphHashes) SetQuickXorHashNil(b bool)`

 SetQuickXorHashNil sets the value for QuickXorHash to be an explicit nil

### UnsetQuickXorHash
`func (o *OpenGraphHashes) UnsetQuickXorHash()`

UnsetQuickXorHash ensures that no value is present for QuickXorHash, not even an explicit nil
### GetSha1Hash

`func (o *OpenGraphHashes) GetSha1Hash() string`

GetSha1Hash returns the Sha1Hash field if non-nil, zero value otherwise.

### GetSha1HashOk

`func (o *OpenGraphHashes) GetSha1HashOk() (*string, bool)`

GetSha1HashOk returns a tuple with the Sha1Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Hash

`func (o *OpenGraphHashes) SetSha1Hash(v string)`

SetSha1Hash sets Sha1Hash field to given value.

### HasSha1Hash

`func (o *OpenGraphHashes) HasSha1Hash() bool`

HasSha1Hash returns a boolean if a field has been set.

### SetSha1HashNil

`func (o *OpenGraphHashes) SetSha1HashNil(b bool)`

 SetSha1HashNil sets the value for Sha1Hash to be an explicit nil

### UnsetSha1Hash
`func (o *OpenGraphHashes) UnsetSha1Hash()`

UnsetSha1Hash ensures that no value is present for Sha1Hash, not even an explicit nil
### GetSha256Hash

`func (o *OpenGraphHashes) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *OpenGraphHashes) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *OpenGraphHashes) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.

### HasSha256Hash

`func (o *OpenGraphHashes) HasSha256Hash() bool`

HasSha256Hash returns a boolean if a field has been set.

### SetSha256HashNil

`func (o *OpenGraphHashes) SetSha256HashNil(b bool)`

 SetSha256HashNil sets the value for Sha256Hash to be an explicit nil

### UnsetSha256Hash
`func (o *OpenGraphHashes) UnsetSha256Hash()`

UnsetSha256Hash ensures that no value is present for Sha256Hash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


