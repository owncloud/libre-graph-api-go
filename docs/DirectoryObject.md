# DirectoryObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedDateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDirectoryObject

`func NewDirectoryObject() *DirectoryObject`

NewDirectoryObject instantiates a new DirectoryObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryObjectWithDefaults

`func NewDirectoryObjectWithDefaults() *DirectoryObject`

NewDirectoryObjectWithDefaults instantiates a new DirectoryObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedDateTime

`func (o *DirectoryObject) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *DirectoryObject) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *DirectoryObject) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *DirectoryObject) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


