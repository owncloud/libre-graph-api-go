# FileSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **time.Time** | The UTC date and time the file was created on a client. | [optional] 
**LastAccessedDateTime** | Pointer to **time.Time** | The UTC date and time the file was last accessed. Available for the recent file list only. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The UTC date and time the file was last modified on a client. | [optional] 

## Methods

### NewFileSystemInfo

`func NewFileSystemInfo() *FileSystemInfo`

NewFileSystemInfo instantiates a new FileSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemInfoWithDefaults

`func NewFileSystemInfoWithDefaults() *FileSystemInfo`

NewFileSystemInfoWithDefaults instantiates a new FileSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *FileSystemInfo) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *FileSystemInfo) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *FileSystemInfo) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *FileSystemInfo) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetLastAccessedDateTime

`func (o *FileSystemInfo) GetLastAccessedDateTime() time.Time`

GetLastAccessedDateTime returns the LastAccessedDateTime field if non-nil, zero value otherwise.

### GetLastAccessedDateTimeOk

`func (o *FileSystemInfo) GetLastAccessedDateTimeOk() (*time.Time, bool)`

GetLastAccessedDateTimeOk returns a tuple with the LastAccessedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedDateTime

`func (o *FileSystemInfo) SetLastAccessedDateTime(v time.Time)`

SetLastAccessedDateTime sets LastAccessedDateTime field to given value.

### HasLastAccessedDateTime

`func (o *FileSystemInfo) HasLastAccessedDateTime() bool`

HasLastAccessedDateTime returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *FileSystemInfo) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *FileSystemInfo) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *FileSystemInfo) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *FileSystemInfo) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


