# DirectoryObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the object. 12345678-9abc-def0-1234-56789abcde. The value of the ID property is often, but not exclusively, in the form of a GUID. The value should be treated as an opaque identifier and not based in being a GUID. Null values are not allowed. Read-only. | [optional] [readonly] 
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

### GetId

`func (o *DirectoryObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DirectoryObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DirectoryObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DirectoryObject) HasId() bool`

HasId returns a boolean if a field has been set.

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


