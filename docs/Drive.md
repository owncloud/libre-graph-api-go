# Drive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveType** | Pointer to **string** | Describes the type of drive represented by this resource. Values are \&quot;personal\&quot; for users home spaces, \&quot;projectSpace\&quot; or \&quot;shares\&quot;. Read-only. | [optional] [readonly] 
**OCDriveStatus** | Pointer to **string** | Describes the status of the drive. | [optional] 
**Owner** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**OCCoOwner** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**Quota** | Pointer to [**Quota**](Quota.md) |  | [optional] 
**Items** | Pointer to [**[]DriveItem**](DriveItem.md) | All items contained in the drive. Read-only. Nullable. | [optional] [readonly] 
**Root** | Pointer to [**DriveItem**](DriveItem.md) |  | [optional] 

## Methods

### NewDrive

`func NewDrive() *Drive`

NewDrive instantiates a new Drive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveWithDefaults

`func NewDriveWithDefaults() *Drive`

NewDriveWithDefaults instantiates a new Drive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetOCDriveStatus

`func (o *Drive) GetOCDriveStatus() string`

GetOCDriveStatus returns the OCDriveStatus field if non-nil, zero value otherwise.

### GetOCDriveStatusOk

`func (o *Drive) GetOCDriveStatusOk() (*string, bool)`

GetOCDriveStatusOk returns a tuple with the OCDriveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCDriveStatus

`func (o *Drive) SetOCDriveStatus(v string)`

SetOCDriveStatus sets OCDriveStatus field to given value.

### HasOCDriveStatus

`func (o *Drive) HasOCDriveStatus() bool`

HasOCDriveStatus returns a boolean if a field has been set.

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

### GetOCCoOwner

`func (o *Drive) GetOCCoOwner() IdentitySet`

GetOCCoOwner returns the OCCoOwner field if non-nil, zero value otherwise.

### GetOCCoOwnerOk

`func (o *Drive) GetOCCoOwnerOk() (*IdentitySet, bool)`

GetOCCoOwnerOk returns a tuple with the OCCoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCCoOwner

`func (o *Drive) SetOCCoOwner(v IdentitySet)`

SetOCCoOwner sets OCCoOwner field to given value.

### HasOCCoOwner

`func (o *Drive) HasOCCoOwner() bool`

HasOCCoOwner returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


