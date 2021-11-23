# BaseItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of item creation. Read-only. | [optional] [readonly] 
**Description** | Pointer to **string** | Provides a user-visible description of the item. Optional. | [optional] 
**ETag** | Pointer to **string** | ETag for the item. Read-only. | [optional] [readonly] 
**LastModifiedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Date and time the item was last modified. Read-only. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the item. Read-write. | [optional] 
**ParentReference** | Pointer to [**ItemReference**](ItemReference.md) |  | [optional] 
**WebUrl** | Pointer to **string** | URL that displays the resource in the browser. Read-only. | [optional] [readonly] 
**CreatedByUser** | Pointer to [**User**](User.md) |  | [optional] 
**LastModifiedByUser** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewBaseItemAllOf

`func NewBaseItemAllOf() *BaseItemAllOf`

NewBaseItemAllOf instantiates a new BaseItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseItemAllOfWithDefaults

`func NewBaseItemAllOfWithDefaults() *BaseItemAllOf`

NewBaseItemAllOfWithDefaults instantiates a new BaseItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *BaseItemAllOf) GetCreatedBy() IdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BaseItemAllOf) GetCreatedByOk() (*IdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BaseItemAllOf) SetCreatedBy(v IdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BaseItemAllOf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *BaseItemAllOf) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *BaseItemAllOf) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *BaseItemAllOf) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *BaseItemAllOf) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *BaseItemAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseItemAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseItemAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseItemAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetETag

`func (o *BaseItemAllOf) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *BaseItemAllOf) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *BaseItemAllOf) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *BaseItemAllOf) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *BaseItemAllOf) GetLastModifiedBy() IdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *BaseItemAllOf) GetLastModifiedByOk() (*IdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *BaseItemAllOf) SetLastModifiedBy(v IdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *BaseItemAllOf) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *BaseItemAllOf) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *BaseItemAllOf) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *BaseItemAllOf) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *BaseItemAllOf) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *BaseItemAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseItemAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseItemAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseItemAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentReference

`func (o *BaseItemAllOf) GetParentReference() ItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *BaseItemAllOf) GetParentReferenceOk() (*ItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *BaseItemAllOf) SetParentReference(v ItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *BaseItemAllOf) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetWebUrl

`func (o *BaseItemAllOf) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *BaseItemAllOf) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *BaseItemAllOf) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *BaseItemAllOf) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *BaseItemAllOf) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *BaseItemAllOf) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *BaseItemAllOf) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *BaseItemAllOf) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetLastModifiedByUser

`func (o *BaseItemAllOf) GetLastModifiedByUser() User`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *BaseItemAllOf) GetLastModifiedByUserOk() (*User, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUser

`func (o *BaseItemAllOf) SetLastModifiedByUser(v User)`

SetLastModifiedByUser sets LastModifiedByUser field to given value.

### HasLastModifiedByUser

`func (o *BaseItemAllOf) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


