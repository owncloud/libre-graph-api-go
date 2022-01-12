# Trash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrashedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**TrashedDateTime** | Pointer to **time.Time** | The UTC date and time the folder was marked as trashed. | [optional] 

## Methods

### NewTrash

`func NewTrash() *Trash`

NewTrash instantiates a new Trash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrashWithDefaults

`func NewTrashWithDefaults() *Trash`

NewTrashWithDefaults instantiates a new Trash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrashedBy

`func (o *Trash) GetTrashedBy() IdentitySet`

GetTrashedBy returns the TrashedBy field if non-nil, zero value otherwise.

### GetTrashedByOk

`func (o *Trash) GetTrashedByOk() (*IdentitySet, bool)`

GetTrashedByOk returns a tuple with the TrashedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashedBy

`func (o *Trash) SetTrashedBy(v IdentitySet)`

SetTrashedBy sets TrashedBy field to given value.

### HasTrashedBy

`func (o *Trash) HasTrashedBy() bool`

HasTrashedBy returns a boolean if a field has been set.

### GetTrashedDateTime

`func (o *Trash) GetTrashedDateTime() time.Time`

GetTrashedDateTime returns the TrashedDateTime field if non-nil, zero value otherwise.

### GetTrashedDateTimeOk

`func (o *Trash) GetTrashedDateTimeOk() (*time.Time, bool)`

GetTrashedDateTimeOk returns a tuple with the TrashedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashedDateTime

`func (o *Trash) SetTrashedDateTime(v time.Time)`

SetTrashedDateTime sets TrashedDateTime field to given value.

### HasTrashedDateTime

`func (o *Trash) HasTrashedDateTime() bool`

HasTrashedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


