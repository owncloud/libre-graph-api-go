# OpenGraphQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **int64** | Total space consumed by files in the recycle bin, in bytes. Read-only. | [optional] [readonly] 
**Remaining** | Pointer to **int64** | Total space remaining before reaching the quota limit, in bytes. Read-only. | [optional] [readonly] 
**State** | Pointer to **string** | Enumeration value that indicates the state of the storage space. Read-only. | [optional] [readonly] 
**StoragePlanInformation** | Pointer to [**OpenGraphStoragePlanInformation**](open.graph.storagePlanInformation.md) | Information about the drive&#39;s storage quota plans. | [optional] 
**Total** | Pointer to **int64** | Total allowed storage space, in bytes. Read-only. | [optional] [readonly] 
**Used** | Pointer to **int64** | Total space used, in bytes. Read-only. | [optional] [readonly] 

## Methods

### NewOpenGraphQuota

`func NewOpenGraphQuota() *OpenGraphQuota`

NewOpenGraphQuota instantiates a new OpenGraphQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphQuotaWithDefaults

`func NewOpenGraphQuotaWithDefaults() *OpenGraphQuota`

NewOpenGraphQuotaWithDefaults instantiates a new OpenGraphQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *OpenGraphQuota) GetDeleted() int64`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OpenGraphQuota) GetDeletedOk() (*int64, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OpenGraphQuota) SetDeleted(v int64)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OpenGraphQuota) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetRemaining

`func (o *OpenGraphQuota) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *OpenGraphQuota) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *OpenGraphQuota) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *OpenGraphQuota) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetState

`func (o *OpenGraphQuota) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OpenGraphQuota) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OpenGraphQuota) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OpenGraphQuota) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStoragePlanInformation

`func (o *OpenGraphQuota) GetStoragePlanInformation() OpenGraphStoragePlanInformation`

GetStoragePlanInformation returns the StoragePlanInformation field if non-nil, zero value otherwise.

### GetStoragePlanInformationOk

`func (o *OpenGraphQuota) GetStoragePlanInformationOk() (*OpenGraphStoragePlanInformation, bool)`

GetStoragePlanInformationOk returns a tuple with the StoragePlanInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePlanInformation

`func (o *OpenGraphQuota) SetStoragePlanInformation(v OpenGraphStoragePlanInformation)`

SetStoragePlanInformation sets StoragePlanInformation field to given value.

### HasStoragePlanInformation

`func (o *OpenGraphQuota) HasStoragePlanInformation() bool`

HasStoragePlanInformation returns a boolean if a field has been set.

### GetTotal

`func (o *OpenGraphQuota) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OpenGraphQuota) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OpenGraphQuota) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OpenGraphQuota) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *OpenGraphQuota) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *OpenGraphQuota) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *OpenGraphQuota) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *OpenGraphQuota) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


