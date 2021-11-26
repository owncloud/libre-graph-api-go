# DriveInline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveType** | Pointer to **string** | Describes the type of drive represented by this resource. Values are \&quot;personal\&quot; for users home spaces, \&quot;project\&quot; or \&quot;share\&quot;. Read-only. | [optional] [readonly] 
**Owner** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**Quota** | Pointer to [**Quota**](Quota.md) |  | [optional] 
**Items** | Pointer to [**[]DriveItem**](DriveItem.md) | All items contained in the drive. Read-only. Nullable. | [optional] [readonly] 
**Root** | Pointer to [**DriveItem**](DriveItem.md) |  | [optional] 

## Methods

### NewDriveInline

`func NewDriveInline() *DriveInline`

NewDriveInline instantiates a new DriveInline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveInlineWithDefaults

`func NewDriveInlineWithDefaults() *DriveInline`

NewDriveInlineWithDefaults instantiates a new DriveInline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveType

`func (o *DriveInline) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *DriveInline) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *DriveInline) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *DriveInline) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetOwner

`func (o *DriveInline) GetOwner() IdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DriveInline) GetOwnerOk() (*IdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DriveInline) SetOwner(v IdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DriveInline) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetQuota

`func (o *DriveInline) GetQuota() Quota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *DriveInline) GetQuotaOk() (*Quota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *DriveInline) SetQuota(v Quota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *DriveInline) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetItems

`func (o *DriveInline) GetItems() []DriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DriveInline) GetItemsOk() (*[]DriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DriveInline) SetItems(v []DriveItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *DriveInline) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRoot

`func (o *DriveInline) GetRoot() DriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *DriveInline) GetRootOk() (*DriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *DriveInline) SetRoot(v DriveItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *DriveInline) HasRoot() bool`

HasRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


