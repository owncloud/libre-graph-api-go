# Quota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **int64** | Total space consumed by files in the recycle bin, in bytes. Read-only. | [optional] [readonly] 
**Remaining** | Pointer to **int64** | Total space remaining before reaching the quota limit, in bytes. Read-only. | [optional] [readonly] 
**State** | Pointer to **string** | Enumeration value that indicates the state of the storage space. Read-only. | [optional] [readonly] 
**Total** | Pointer to **int64** | Total allowed storage space, in bytes. Read-only. | [optional] [readonly] 
**Used** | Pointer to **int64** | Total space used, in bytes. Read-only. | [optional] [readonly] 

## Methods

### NewQuota

`func NewQuota() *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *Quota) GetDeleted() int64`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Quota) GetDeletedOk() (*int64, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Quota) SetDeleted(v int64)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Quota) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetRemaining

`func (o *Quota) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *Quota) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *Quota) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *Quota) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetState

`func (o *Quota) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Quota) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Quota) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Quota) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTotal

`func (o *Quota) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Quota) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Quota) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Quota) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *Quota) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Quota) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Quota) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Quota) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


