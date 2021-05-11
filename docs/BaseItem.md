# BaseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**IdentitySet**](identitySet.md) | Identity of the user, device, or application which created the item. Read-only. | [optional] [readonly] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of item creation. Read-only. | [optional] [readonly] 
**Description** | Pointer to **string** | Provides a user-visible description of the item. Optional. | [optional] 
**ETag** | Pointer to **string** | ETag for the item. Read-only. | [optional] [readonly] 
**LastModifiedBy** | Pointer to [**IdentitySet**](identitySet.md) | Identity of the user, device, and application which last modified the item. Read-only. | [optional] [readonly] 
**LastModifiedDateTime** | Pointer to **time.Time** | Date and time the item was last modified. Read-only. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the item. Read-write. | [optional] 
**ParentReference** | Pointer to [**ItemReference**](itemReference.md) | Parent information, if the item has a parent. Read-write. | [optional] 
**WebUrl** | Pointer to **string** | URL that displays the resource in the browser. Read-only. | [optional] [readonly] 
**CreatedByUser** | Pointer to [**User**](user.md) | Identity of the user who created the item. Read-only. | [optional] [readonly] 
**LastModifiedByUser** | Pointer to [**User**](user.md) | Identity of the user who last modified the item. Read-only. | [optional] [readonly] 

## Methods

### NewBaseItem

`func NewBaseItem() *BaseItem`

NewBaseItem instantiates a new BaseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseItemWithDefaults

`func NewBaseItemWithDefaults() *BaseItem`

NewBaseItemWithDefaults instantiates a new BaseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *BaseItem) GetCreatedBy() IdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BaseItem) GetCreatedByOk() (*IdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BaseItem) SetCreatedBy(v IdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BaseItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *BaseItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *BaseItem) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *BaseItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *BaseItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *BaseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetETag

`func (o *BaseItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *BaseItem) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *BaseItem) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *BaseItem) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *BaseItem) GetLastModifiedBy() IdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *BaseItem) GetLastModifiedByOk() (*IdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *BaseItem) SetLastModifiedBy(v IdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *BaseItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *BaseItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *BaseItem) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *BaseItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *BaseItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *BaseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentReference

`func (o *BaseItem) GetParentReference() ItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *BaseItem) GetParentReferenceOk() (*ItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *BaseItem) SetParentReference(v ItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *BaseItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetWebUrl

`func (o *BaseItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *BaseItem) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *BaseItem) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *BaseItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *BaseItem) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *BaseItem) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *BaseItem) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *BaseItem) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetLastModifiedByUser

`func (o *BaseItem) GetLastModifiedByUser() User`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *BaseItem) GetLastModifiedByUserOk() (*User, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUser

`func (o *BaseItem) SetLastModifiedByUser(v User)`

SetLastModifiedByUser sets LastModifiedByUser field to given value.

### HasLastModifiedByUser

`func (o *BaseItem) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


