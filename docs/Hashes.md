# Hashes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crc32Hash** | Pointer to **string** | The CRC32 value of the file (if available). Read-only. | [optional] 
**QuickXorHash** | Pointer to **string** | A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only. | [optional] 
**Sha1Hash** | Pointer to **string** | SHA1 hash for the contents of the file (if available). Read-only. | [optional] 
**Sha256Hash** | Pointer to **string** | SHA256 hash for the contents of the file (if available). Read-only. | [optional] 

## Methods

### NewHashes

`func NewHashes() *Hashes`

NewHashes instantiates a new Hashes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashesWithDefaults

`func NewHashesWithDefaults() *Hashes`

NewHashesWithDefaults instantiates a new Hashes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrc32Hash

`func (o *Hashes) GetCrc32Hash() string`

GetCrc32Hash returns the Crc32Hash field if non-nil, zero value otherwise.

### GetCrc32HashOk

`func (o *Hashes) GetCrc32HashOk() (*string, bool)`

GetCrc32HashOk returns a tuple with the Crc32Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrc32Hash

`func (o *Hashes) SetCrc32Hash(v string)`

SetCrc32Hash sets Crc32Hash field to given value.

### HasCrc32Hash

`func (o *Hashes) HasCrc32Hash() bool`

HasCrc32Hash returns a boolean if a field has been set.

### GetQuickXorHash

`func (o *Hashes) GetQuickXorHash() string`

GetQuickXorHash returns the QuickXorHash field if non-nil, zero value otherwise.

### GetQuickXorHashOk

`func (o *Hashes) GetQuickXorHashOk() (*string, bool)`

GetQuickXorHashOk returns a tuple with the QuickXorHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickXorHash

`func (o *Hashes) SetQuickXorHash(v string)`

SetQuickXorHash sets QuickXorHash field to given value.

### HasQuickXorHash

`func (o *Hashes) HasQuickXorHash() bool`

HasQuickXorHash returns a boolean if a field has been set.

### GetSha1Hash

`func (o *Hashes) GetSha1Hash() string`

GetSha1Hash returns the Sha1Hash field if non-nil, zero value otherwise.

### GetSha1HashOk

`func (o *Hashes) GetSha1HashOk() (*string, bool)`

GetSha1HashOk returns a tuple with the Sha1Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Hash

`func (o *Hashes) SetSha1Hash(v string)`

SetSha1Hash sets Sha1Hash field to given value.

### HasSha1Hash

`func (o *Hashes) HasSha1Hash() bool`

HasSha1Hash returns a boolean if a field has been set.

### GetSha256Hash

`func (o *Hashes) GetSha256Hash() string`

GetSha256Hash returns the Sha256Hash field if non-nil, zero value otherwise.

### GetSha256HashOk

`func (o *Hashes) GetSha256HashOk() (*string, bool)`

GetSha256HashOk returns a tuple with the Sha256Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Hash

`func (o *Hashes) SetSha256Hash(v string)`

SetSha256Hash sets Sha256Hash field to given value.

### HasSha256Hash

`func (o *Hashes) HasSha256Hash() bool`

HasSha256Hash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


