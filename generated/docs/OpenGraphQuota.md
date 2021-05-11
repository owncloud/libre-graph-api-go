# OpenGraphQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **NullableInt64** | Total space consumed by files in the recycle bin, in bytes. Read-only. | [optional] [readonly] 
**Remaining** | Pointer to **NullableInt64** | Total space remaining before reaching the quota limit, in bytes. Read-only. | [optional] [readonly] 
**State** | Pointer to **NullableString** | Enumeration value that indicates the state of the storage space. Read-only. | [optional] [readonly] 
**StoragePlanInformation** | Pointer to [**NullableAnyOfopenGraphStoragePlanInformation**](anyOf&lt;open.graph.storagePlanInformation&gt;.md) | Information about the drive&#39;s storage quota plans. | [optional] 
**Total** | Pointer to **NullableInt64** | Total allowed storage space, in bytes. Read-only. | [optional] [readonly] 
**Used** | Pointer to **NullableInt64** | Total space used, in bytes. Read-only. | [optional] [readonly] 

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

### SetDeletedNil

`func (o *OpenGraphQuota) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *OpenGraphQuota) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
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

### SetRemainingNil

`func (o *OpenGraphQuota) SetRemainingNil(b bool)`

 SetRemainingNil sets the value for Remaining to be an explicit nil

### UnsetRemaining
`func (o *OpenGraphQuota) UnsetRemaining()`

UnsetRemaining ensures that no value is present for Remaining, not even an explicit nil
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

### SetStateNil

`func (o *OpenGraphQuota) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *OpenGraphQuota) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStoragePlanInformation

`func (o *OpenGraphQuota) GetStoragePlanInformation() AnyOfopenGraphStoragePlanInformation`

GetStoragePlanInformation returns the StoragePlanInformation field if non-nil, zero value otherwise.

### GetStoragePlanInformationOk

`func (o *OpenGraphQuota) GetStoragePlanInformationOk() (*AnyOfopenGraphStoragePlanInformation, bool)`

GetStoragePlanInformationOk returns a tuple with the StoragePlanInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePlanInformation

`func (o *OpenGraphQuota) SetStoragePlanInformation(v AnyOfopenGraphStoragePlanInformation)`

SetStoragePlanInformation sets StoragePlanInformation field to given value.

### HasStoragePlanInformation

`func (o *OpenGraphQuota) HasStoragePlanInformation() bool`

HasStoragePlanInformation returns a boolean if a field has been set.

### SetStoragePlanInformationNil

`func (o *OpenGraphQuota) SetStoragePlanInformationNil(b bool)`

 SetStoragePlanInformationNil sets the value for StoragePlanInformation to be an explicit nil

### UnsetStoragePlanInformation
`func (o *OpenGraphQuota) UnsetStoragePlanInformation()`

UnsetStoragePlanInformation ensures that no value is present for StoragePlanInformation, not even an explicit nil
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

### SetTotalNil

`func (o *OpenGraphQuota) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *OpenGraphQuota) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
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

### SetUsedNil

`func (o *OpenGraphQuota) SetUsedNil(b bool)`

 SetUsedNil sets the value for Used to be an explicit nil

### UnsetUsed
`func (o *OpenGraphQuota) UnsetUsed()`

UnsetUsed ensures that no value is present for Used, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


