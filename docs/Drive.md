# Drive

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
**Name** | **string** | The name of the item. Read-write. | 
**ParentReference** | Pointer to [**ItemReference**](ItemReference.md) |  | [optional] 
**WebUrl** | Pointer to **string** | URL that displays the resource in the browser. Read-only. | [optional] [readonly] 
**CreatedByUser** | Pointer to [**User**](User.md) |  | [optional] 
**LastModifiedByUser** | Pointer to [**User**](User.md) |  | [optional] 
**DriveType** | Pointer to **string** | Describes the type of drive represented by this resource. Values are \&quot;personal\&quot; for users home spaces, \&quot;project\&quot;, \&quot;virtual\&quot; or \&quot;share\&quot;. Read-only. | [optional] [readonly] 
**DriveAlias** | Pointer to **string** | The drive alias can be used in clients to make the urls user friendly. Example: &#39;personal/einstein&#39;. This will be used to resolve to the correct driveID. | [optional] 
**Owner** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**Quota** | Pointer to [**Quota**](Quota.md) |  | [optional] 
**Items** | Pointer to [**[]DriveItem**](DriveItem.md) | All items contained in the drive. Read-only. Nullable. | [optional] [readonly] 
**Root** | Pointer to [**DriveItem**](DriveItem.md) |  | [optional] 
**Special** | Pointer to [**[]DriveItem**](DriveItem.md) | A collection of special drive resources. | [optional] 

## Methods

### NewDrive

`func NewDrive(name string, ) *Drive`

NewDrive instantiates a new Drive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveWithDefaults

`func NewDriveWithDefaults() *Drive`

NewDriveWithDefaults instantiates a new Drive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Drive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Drive) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Drive) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Drive) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Drive) GetCreatedBy() IdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Drive) GetCreatedByOk() (*IdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Drive) SetCreatedBy(v IdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Drive) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *Drive) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Drive) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Drive) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Drive) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *Drive) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Drive) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Drive) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Drive) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetETag

`func (o *Drive) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *Drive) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *Drive) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *Drive) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *Drive) GetLastModifiedBy() IdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *Drive) GetLastModifiedByOk() (*IdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *Drive) SetLastModifiedBy(v IdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *Drive) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *Drive) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *Drive) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *Drive) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *Drive) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *Drive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Drive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Drive) SetName(v string)`

SetName sets Name field to given value.


### GetParentReference

`func (o *Drive) GetParentReference() ItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *Drive) GetParentReferenceOk() (*ItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *Drive) SetParentReference(v ItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *Drive) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetWebUrl

`func (o *Drive) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *Drive) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *Drive) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *Drive) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *Drive) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *Drive) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *Drive) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *Drive) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetLastModifiedByUser

`func (o *Drive) GetLastModifiedByUser() User`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *Drive) GetLastModifiedByUserOk() (*User, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUser

`func (o *Drive) SetLastModifiedByUser(v User)`

SetLastModifiedByUser sets LastModifiedByUser field to given value.

### HasLastModifiedByUser

`func (o *Drive) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### GetDriveType

`func (o *Drive) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *Drive) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *Drive) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *Drive) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetDriveAlias

`func (o *Drive) GetDriveAlias() string`

GetDriveAlias returns the DriveAlias field if non-nil, zero value otherwise.

### GetDriveAliasOk

`func (o *Drive) GetDriveAliasOk() (*string, bool)`

GetDriveAliasOk returns a tuple with the DriveAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveAlias

`func (o *Drive) SetDriveAlias(v string)`

SetDriveAlias sets DriveAlias field to given value.

### HasDriveAlias

`func (o *Drive) HasDriveAlias() bool`

HasDriveAlias returns a boolean if a field has been set.

### GetOwner

`func (o *Drive) GetOwner() IdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Drive) GetOwnerOk() (*IdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Drive) SetOwner(v IdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Drive) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetQuota

`func (o *Drive) GetQuota() Quota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Drive) GetQuotaOk() (*Quota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Drive) SetQuota(v Quota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Drive) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetItems

`func (o *Drive) GetItems() []DriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Drive) GetItemsOk() (*[]DriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Drive) SetItems(v []DriveItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Drive) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRoot

`func (o *Drive) GetRoot() DriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *Drive) GetRootOk() (*DriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *Drive) SetRoot(v DriveItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *Drive) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetSpecial

`func (o *Drive) GetSpecial() []DriveItem`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *Drive) GetSpecialOk() (*[]DriveItem, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *Drive) SetSpecial(v []DriveItem)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *Drive) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


