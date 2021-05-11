# DriveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveType** | Pointer to **string** | Describes the type of drive represented by this resource. Values are \&quot;personal\&quot; for users home spaces, \&quot;projectSpace\&quot; or \&quot;shares\&quot;. Read-only. | [optional] [readonly] 
**OCDriveStatus** | Pointer to **string** | Describes the status of the drive. | [optional] 
**Owner** | Pointer to [**IdentitySet**](identitySet.md) | Optional. The user account that owns the drive. Read-only. | [optional] [readonly] 
**OCCoOwner** | Pointer to [**IdentitySet**](identitySet.md) | Optional. The user account that owns the drive. | [optional] 
**Quota** | Pointer to [**Quota**](quota.md) | Optional. Information about the drive&#39;s storage space quota. Read-only. | [optional] [readonly] 
**Items** | Pointer to [**[]DriveItem**](DriveItem.md) | All items contained in the drive. Read-only. Nullable. | [optional] [readonly] 
**Root** | Pointer to [**DriveItem**](driveItem.md) | The root folder of the drive. Read-only. | [optional] [readonly] 

## Methods

### NewDriveAllOf

`func NewDriveAllOf() *DriveAllOf`

NewDriveAllOf instantiates a new DriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveAllOfWithDefaults

`func NewDriveAllOfWithDefaults() *DriveAllOf`

NewDriveAllOfWithDefaults instantiates a new DriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveType

`func (o *DriveAllOf) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *DriveAllOf) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *DriveAllOf) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *DriveAllOf) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetOCDriveStatus

`func (o *DriveAllOf) GetOCDriveStatus() string`

GetOCDriveStatus returns the OCDriveStatus field if non-nil, zero value otherwise.

### GetOCDriveStatusOk

`func (o *DriveAllOf) GetOCDriveStatusOk() (*string, bool)`

GetOCDriveStatusOk returns a tuple with the OCDriveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCDriveStatus

`func (o *DriveAllOf) SetOCDriveStatus(v string)`

SetOCDriveStatus sets OCDriveStatus field to given value.

### HasOCDriveStatus

`func (o *DriveAllOf) HasOCDriveStatus() bool`

HasOCDriveStatus returns a boolean if a field has been set.

### GetOwner

`func (o *DriveAllOf) GetOwner() IdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DriveAllOf) GetOwnerOk() (*IdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DriveAllOf) SetOwner(v IdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DriveAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOCCoOwner

`func (o *DriveAllOf) GetOCCoOwner() IdentitySet`

GetOCCoOwner returns the OCCoOwner field if non-nil, zero value otherwise.

### GetOCCoOwnerOk

`func (o *DriveAllOf) GetOCCoOwnerOk() (*IdentitySet, bool)`

GetOCCoOwnerOk returns a tuple with the OCCoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCCoOwner

`func (o *DriveAllOf) SetOCCoOwner(v IdentitySet)`

SetOCCoOwner sets OCCoOwner field to given value.

### HasOCCoOwner

`func (o *DriveAllOf) HasOCCoOwner() bool`

HasOCCoOwner returns a boolean if a field has been set.

### GetQuota

`func (o *DriveAllOf) GetQuota() Quota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *DriveAllOf) GetQuotaOk() (*Quota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *DriveAllOf) SetQuota(v Quota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *DriveAllOf) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetItems

`func (o *DriveAllOf) GetItems() []DriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DriveAllOf) GetItemsOk() (*[]DriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DriveAllOf) SetItems(v []DriveItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *DriveAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRoot

`func (o *DriveAllOf) GetRoot() DriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *DriveAllOf) GetRootOk() (*DriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *DriveAllOf) SetRoot(v DriveItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *DriveAllOf) HasRoot() bool`

HasRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


