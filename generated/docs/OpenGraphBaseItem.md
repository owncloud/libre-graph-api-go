# OpenGraphBaseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**NullableAnyOfopenGraphIdentitySet**](anyOf&lt;open.graph.identitySet&gt;.md) | Identity of the user, device, or application which created the item. Read-only. | [optional] [readonly] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of item creation. Read-only. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Provides a user-visible description of the item. Optional. | [optional] 
**ETag** | Pointer to **NullableString** | ETag for the item. Read-only. | [optional] [readonly] 
**LastModifiedBy** | Pointer to [**NullableAnyOfopenGraphIdentitySet**](anyOf&lt;open.graph.identitySet&gt;.md) | Identity of the user, device, and application which last modified the item. Read-only. | [optional] [readonly] 
**LastModifiedDateTime** | Pointer to **time.Time** | Date and time the item was last modified. Read-only. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The name of the item. Read-write. | [optional] 
**ParentReference** | Pointer to [**NullableAnyOfopenGraphItemReference**](anyOf&lt;open.graph.itemReference&gt;.md) | Parent information, if the item has a parent. Read-write. | [optional] 
**WebUrl** | Pointer to **NullableString** | URL that displays the resource in the browser. Read-only. | [optional] [readonly] 
**CreatedByUser** | Pointer to [**NullableAnyOfopenGraphUser**](anyOf&lt;open.graph.user&gt;.md) | Identity of the user who created the item. Read-only. | [optional] [readonly] 
**LastModifiedByUser** | Pointer to [**NullableAnyOfopenGraphUser**](anyOf&lt;open.graph.user&gt;.md) | Identity of the user who last modified the item. Read-only. | [optional] [readonly] 

## Methods

### NewOpenGraphBaseItem

`func NewOpenGraphBaseItem() *OpenGraphBaseItem`

NewOpenGraphBaseItem instantiates a new OpenGraphBaseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphBaseItemWithDefaults

`func NewOpenGraphBaseItemWithDefaults() *OpenGraphBaseItem`

NewOpenGraphBaseItemWithDefaults instantiates a new OpenGraphBaseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *OpenGraphBaseItem) GetCreatedBy() AnyOfopenGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OpenGraphBaseItem) GetCreatedByOk() (*AnyOfopenGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OpenGraphBaseItem) SetCreatedBy(v AnyOfopenGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OpenGraphBaseItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *OpenGraphBaseItem) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *OpenGraphBaseItem) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *OpenGraphBaseItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OpenGraphBaseItem) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *OpenGraphBaseItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *OpenGraphBaseItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *OpenGraphBaseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenGraphBaseItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenGraphBaseItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenGraphBaseItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OpenGraphBaseItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OpenGraphBaseItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetETag

`func (o *OpenGraphBaseItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *OpenGraphBaseItem) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *OpenGraphBaseItem) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *OpenGraphBaseItem) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *OpenGraphBaseItem) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *OpenGraphBaseItem) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil
### GetLastModifiedBy

`func (o *OpenGraphBaseItem) GetLastModifiedBy() AnyOfopenGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *OpenGraphBaseItem) GetLastModifiedByOk() (*AnyOfopenGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *OpenGraphBaseItem) SetLastModifiedBy(v AnyOfopenGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *OpenGraphBaseItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *OpenGraphBaseItem) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *OpenGraphBaseItem) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *OpenGraphBaseItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *OpenGraphBaseItem) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *OpenGraphBaseItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *OpenGraphBaseItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *OpenGraphBaseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenGraphBaseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenGraphBaseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenGraphBaseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpenGraphBaseItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpenGraphBaseItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentReference

`func (o *OpenGraphBaseItem) GetParentReference() AnyOfopenGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *OpenGraphBaseItem) GetParentReferenceOk() (*AnyOfopenGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *OpenGraphBaseItem) SetParentReference(v AnyOfopenGraphItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *OpenGraphBaseItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReferenceNil

`func (o *OpenGraphBaseItem) SetParentReferenceNil(b bool)`

 SetParentReferenceNil sets the value for ParentReference to be an explicit nil

### UnsetParentReference
`func (o *OpenGraphBaseItem) UnsetParentReference()`

UnsetParentReference ensures that no value is present for ParentReference, not even an explicit nil
### GetWebUrl

`func (o *OpenGraphBaseItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *OpenGraphBaseItem) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *OpenGraphBaseItem) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *OpenGraphBaseItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *OpenGraphBaseItem) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *OpenGraphBaseItem) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetCreatedByUser

`func (o *OpenGraphBaseItem) GetCreatedByUser() AnyOfopenGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *OpenGraphBaseItem) GetCreatedByUserOk() (*AnyOfopenGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *OpenGraphBaseItem) SetCreatedByUser(v AnyOfopenGraphUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *OpenGraphBaseItem) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUserNil

`func (o *OpenGraphBaseItem) SetCreatedByUserNil(b bool)`

 SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil

### UnsetCreatedByUser
`func (o *OpenGraphBaseItem) UnsetCreatedByUser()`

UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
### GetLastModifiedByUser

`func (o *OpenGraphBaseItem) GetLastModifiedByUser() AnyOfopenGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *OpenGraphBaseItem) GetLastModifiedByUserOk() (*AnyOfopenGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUser

`func (o *OpenGraphBaseItem) SetLastModifiedByUser(v AnyOfopenGraphUser)`

SetLastModifiedByUser sets LastModifiedByUser field to given value.

### HasLastModifiedByUser

`func (o *OpenGraphBaseItem) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUserNil

`func (o *OpenGraphBaseItem) SetLastModifiedByUserNil(b bool)`

 SetLastModifiedByUserNil sets the value for LastModifiedByUser to be an explicit nil

### UnsetLastModifiedByUser
`func (o *OpenGraphBaseItem) UnsetLastModifiedByUser()`

UnsetLastModifiedByUser ensures that no value is present for LastModifiedByUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


