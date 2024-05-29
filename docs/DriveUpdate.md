# DriveUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique idenfier for this drive. | [optional] [readonly] 
**CreatedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of item creation. Read-only. | [optional] [readonly] 
**Description** | Pointer to **string** | Provides a user-visible description of the item. Optional. | [optional] 
**ETag** | Pointer to **string** | ETag for the item. Read-only. | [optional] [readonly] 
**LastModifiedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Date and time the item was last modified. Read-only. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the item. Read-write. | [optional] 
**ParentReference** | Pointer to [**ItemReference**](ItemReference.md) |  | [optional] 
**WebUrl** | Pointer to **string** | URL that displays the resource in the browser. Read-only. | [optional] [readonly] 
**DriveType** | Pointer to **string** | Describes the type of drive represented by this resource. Values are \&quot;personal\&quot; for users home spaces, \&quot;project\&quot;, \&quot;virtual\&quot; or \&quot;share\&quot;. Read-only. | [optional] [readonly] 
**DriveAlias** | Pointer to **string** | The drive alias can be used in clients to make the urls user friendly. Example: &#39;personal/einstein&#39;. This will be used to resolve to the correct driveID. | [optional] 
**Owner** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**Quota** | Pointer to [**Quota**](Quota.md) |  | [optional] 
**Items** | Pointer to [**[]DriveItem**](DriveItem.md) | All items contained in the drive. Read-only. Nullable. | [optional] [readonly] 
**Root** | Pointer to [**DriveItem**](DriveItem.md) |  | [optional] 
**Special** | Pointer to [**[]DriveItem**](DriveItem.md) | A collection of special drive resources. | [optional] 

## Methods

### NewDriveUpdate

`func NewDriveUpdate() *DriveUpdate`

NewDriveUpdate instantiates a new DriveUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveUpdateWithDefaults

`func NewDriveUpdateWithDefaults() *DriveUpdate`

NewDriveUpdateWithDefaults instantiates a new DriveUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriveUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriveUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriveUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DriveUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DriveUpdate) GetCreatedBy() IdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DriveUpdate) GetCreatedByOk() (*IdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DriveUpdate) SetCreatedBy(v IdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DriveUpdate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *DriveUpdate) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *DriveUpdate) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *DriveUpdate) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *DriveUpdate) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *DriveUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DriveUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DriveUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DriveUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetETag

`func (o *DriveUpdate) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *DriveUpdate) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *DriveUpdate) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *DriveUpdate) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *DriveUpdate) GetLastModifiedBy() IdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *DriveUpdate) GetLastModifiedByOk() (*IdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *DriveUpdate) SetLastModifiedBy(v IdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *DriveUpdate) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *DriveUpdate) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *DriveUpdate) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *DriveUpdate) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *DriveUpdate) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *DriveUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DriveUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DriveUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DriveUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentReference

`func (o *DriveUpdate) GetParentReference() ItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *DriveUpdate) GetParentReferenceOk() (*ItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *DriveUpdate) SetParentReference(v ItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *DriveUpdate) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetWebUrl

`func (o *DriveUpdate) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *DriveUpdate) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *DriveUpdate) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *DriveUpdate) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetDriveType

`func (o *DriveUpdate) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *DriveUpdate) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *DriveUpdate) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *DriveUpdate) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetDriveAlias

`func (o *DriveUpdate) GetDriveAlias() string`

GetDriveAlias returns the DriveAlias field if non-nil, zero value otherwise.

### GetDriveAliasOk

`func (o *DriveUpdate) GetDriveAliasOk() (*string, bool)`

GetDriveAliasOk returns a tuple with the DriveAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveAlias

`func (o *DriveUpdate) SetDriveAlias(v string)`

SetDriveAlias sets DriveAlias field to given value.

### HasDriveAlias

`func (o *DriveUpdate) HasDriveAlias() bool`

HasDriveAlias returns a boolean if a field has been set.

### GetOwner

`func (o *DriveUpdate) GetOwner() IdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DriveUpdate) GetOwnerOk() (*IdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DriveUpdate) SetOwner(v IdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DriveUpdate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetQuota

`func (o *DriveUpdate) GetQuota() Quota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *DriveUpdate) GetQuotaOk() (*Quota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *DriveUpdate) SetQuota(v Quota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *DriveUpdate) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetItems

`func (o *DriveUpdate) GetItems() []DriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DriveUpdate) GetItemsOk() (*[]DriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DriveUpdate) SetItems(v []DriveItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *DriveUpdate) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRoot

`func (o *DriveUpdate) GetRoot() DriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *DriveUpdate) GetRootOk() (*DriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *DriveUpdate) SetRoot(v DriveItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *DriveUpdate) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetSpecial

`func (o *DriveUpdate) GetSpecial() []DriveItem`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *DriveUpdate) GetSpecialOk() (*[]DriveItem, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *DriveUpdate) SetSpecial(v []DriveItem)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *DriveUpdate) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


