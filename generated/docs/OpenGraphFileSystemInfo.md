# OpenGraphFileSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | The UTC date and time the file was created on a client. | [optional] 
**LastAccessedDateTime** | Pointer to **NullableTime** | The UTC date and time the file was last accessed. Available for the recent file list only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The UTC date and time the file was last modified on a client. | [optional] 

## Methods

### NewOpenGraphFileSystemInfo

`func NewOpenGraphFileSystemInfo() *OpenGraphFileSystemInfo`

NewOpenGraphFileSystemInfo instantiates a new OpenGraphFileSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphFileSystemInfoWithDefaults

`func NewOpenGraphFileSystemInfoWithDefaults() *OpenGraphFileSystemInfo`

NewOpenGraphFileSystemInfoWithDefaults instantiates a new OpenGraphFileSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *OpenGraphFileSystemInfo) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OpenGraphFileSystemInfo) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *OpenGraphFileSystemInfo) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *OpenGraphFileSystemInfo) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *OpenGraphFileSystemInfo) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *OpenGraphFileSystemInfo) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastAccessedDateTime

`func (o *OpenGraphFileSystemInfo) GetLastAccessedDateTime() time.Time`

GetLastAccessedDateTime returns the LastAccessedDateTime field if non-nil, zero value otherwise.

### GetLastAccessedDateTimeOk

`func (o *OpenGraphFileSystemInfo) GetLastAccessedDateTimeOk() (*time.Time, bool)`

GetLastAccessedDateTimeOk returns a tuple with the LastAccessedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedDateTime

`func (o *OpenGraphFileSystemInfo) SetLastAccessedDateTime(v time.Time)`

SetLastAccessedDateTime sets LastAccessedDateTime field to given value.

### HasLastAccessedDateTime

`func (o *OpenGraphFileSystemInfo) HasLastAccessedDateTime() bool`

HasLastAccessedDateTime returns a boolean if a field has been set.

### SetLastAccessedDateTimeNil

`func (o *OpenGraphFileSystemInfo) SetLastAccessedDateTimeNil(b bool)`

 SetLastAccessedDateTimeNil sets the value for LastAccessedDateTime to be an explicit nil

### UnsetLastAccessedDateTime
`func (o *OpenGraphFileSystemInfo) UnsetLastAccessedDateTime()`

UnsetLastAccessedDateTime ensures that no value is present for LastAccessedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *OpenGraphFileSystemInfo) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *OpenGraphFileSystemInfo) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *OpenGraphFileSystemInfo) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *OpenGraphFileSystemInfo) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *OpenGraphFileSystemInfo) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *OpenGraphFileSystemInfo) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


