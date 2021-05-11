# OpenGraphDriveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveType** | Pointer to **string** | Describes the type of drive represented by this resource. Values are \&quot;personal\&quot; for users home spaces, \&quot;projectSpace\&quot; or \&quot;shares\&quot;. Read-only. | [optional] [readonly] 
**OCDriveStatus** | Pointer to **string** | Describes the status of the drive. | [optional] 
**Owner** | Pointer to [**OpenGraphIdentitySet**](open.graph.identitySet.md) | Optional. The user account that owns the drive. Read-only. | [optional] [readonly] 
**OCCoOwner** | Pointer to [**OpenGraphIdentitySet**](open.graph.identitySet.md) | Optional. The user account that owns the drive. | [optional] 
**Quota** | Pointer to [**OpenGraphQuota**](open.graph.quota.md) | Optional. Information about the drive&#39;s storage space quota. Read-only. | [optional] [readonly] 
**Items** | Pointer to [**[]OpenGraphDriveItem**](OpenGraphDriveItem.md) | All items contained in the drive. Read-only. Nullable. | [optional] [readonly] 
**Root** | Pointer to [**OpenGraphDriveItem**](open.graph.driveItem.md) | The root folder of the drive. Read-only. | [optional] [readonly] 

## Methods

### NewOpenGraphDriveAllOf

`func NewOpenGraphDriveAllOf() *OpenGraphDriveAllOf`

NewOpenGraphDriveAllOf instantiates a new OpenGraphDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphDriveAllOfWithDefaults

`func NewOpenGraphDriveAllOfWithDefaults() *OpenGraphDriveAllOf`

NewOpenGraphDriveAllOfWithDefaults instantiates a new OpenGraphDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveType

`func (o *OpenGraphDriveAllOf) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *OpenGraphDriveAllOf) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *OpenGraphDriveAllOf) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *OpenGraphDriveAllOf) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetOCDriveStatus

`func (o *OpenGraphDriveAllOf) GetOCDriveStatus() string`

GetOCDriveStatus returns the OCDriveStatus field if non-nil, zero value otherwise.

### GetOCDriveStatusOk

`func (o *OpenGraphDriveAllOf) GetOCDriveStatusOk() (*string, bool)`

GetOCDriveStatusOk returns a tuple with the OCDriveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCDriveStatus

`func (o *OpenGraphDriveAllOf) SetOCDriveStatus(v string)`

SetOCDriveStatus sets OCDriveStatus field to given value.

### HasOCDriveStatus

`func (o *OpenGraphDriveAllOf) HasOCDriveStatus() bool`

HasOCDriveStatus returns a boolean if a field has been set.

### GetOwner

`func (o *OpenGraphDriveAllOf) GetOwner() OpenGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OpenGraphDriveAllOf) GetOwnerOk() (*OpenGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OpenGraphDriveAllOf) SetOwner(v OpenGraphIdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *OpenGraphDriveAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOCCoOwner

`func (o *OpenGraphDriveAllOf) GetOCCoOwner() OpenGraphIdentitySet`

GetOCCoOwner returns the OCCoOwner field if non-nil, zero value otherwise.

### GetOCCoOwnerOk

`func (o *OpenGraphDriveAllOf) GetOCCoOwnerOk() (*OpenGraphIdentitySet, bool)`

GetOCCoOwnerOk returns a tuple with the OCCoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOCCoOwner

`func (o *OpenGraphDriveAllOf) SetOCCoOwner(v OpenGraphIdentitySet)`

SetOCCoOwner sets OCCoOwner field to given value.

### HasOCCoOwner

`func (o *OpenGraphDriveAllOf) HasOCCoOwner() bool`

HasOCCoOwner returns a boolean if a field has been set.

### GetQuota

`func (o *OpenGraphDriveAllOf) GetQuota() OpenGraphQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *OpenGraphDriveAllOf) GetQuotaOk() (*OpenGraphQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *OpenGraphDriveAllOf) SetQuota(v OpenGraphQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *OpenGraphDriveAllOf) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetItems

`func (o *OpenGraphDriveAllOf) GetItems() []OpenGraphDriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OpenGraphDriveAllOf) GetItemsOk() (*[]OpenGraphDriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OpenGraphDriveAllOf) SetItems(v []OpenGraphDriveItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *OpenGraphDriveAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRoot

`func (o *OpenGraphDriveAllOf) GetRoot() OpenGraphDriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *OpenGraphDriveAllOf) GetRootOk() (*OpenGraphDriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *OpenGraphDriveAllOf) SetRoot(v OpenGraphDriveItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *OpenGraphDriveAllOf) HasRoot() bool`

HasRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


