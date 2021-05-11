/*
 * Open Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OpenGraphQuota struct for OpenGraphQuota
type OpenGraphQuota struct {
	// Total space consumed by files in the recycle bin, in bytes. Read-only.
	Deleted NullableInt64 `json:"deleted,omitempty"`
	// Total space remaining before reaching the quota limit, in bytes. Read-only.
	Remaining NullableInt64 `json:"remaining,omitempty"`
	// Enumeration value that indicates the state of the storage space. Read-only.
	State NullableString `json:"state,omitempty"`
	// Information about the drive's storage quota plans.
	StoragePlanInformation NullableAnyOfopenGraphStoragePlanInformation `json:"storagePlanInformation,omitempty"`
	// Total allowed storage space, in bytes. Read-only.
	Total NullableInt64 `json:"total,omitempty"`
	// Total space used, in bytes. Read-only.
	Used NullableInt64 `json:"used,omitempty"`
}

// NewOpenGraphQuota instantiates a new OpenGraphQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphQuota() *OpenGraphQuota {
	this := OpenGraphQuota{}
	return &this
}

// NewOpenGraphQuotaWithDefaults instantiates a new OpenGraphQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphQuotaWithDefaults() *OpenGraphQuota {
	this := OpenGraphQuota{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphQuota) GetDeleted() int64 {
	if o == nil || o.Deleted.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphQuota) GetDeletedOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// HasDeleted returns a boolean if a field has been set.
func (o *OpenGraphQuota) HasDeleted() bool {
	if o != nil && o.Deleted.IsSet() {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given NullableInt64 and assigns it to the Deleted field.
func (o *OpenGraphQuota) SetDeleted(v int64) {
	o.Deleted.Set(&v)
}
// SetDeletedNil sets the value for Deleted to be an explicit nil
func (o *OpenGraphQuota) SetDeletedNil() {
	o.Deleted.Set(nil)
}

// UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
func (o *OpenGraphQuota) UnsetDeleted() {
	o.Deleted.Unset()
}

// GetRemaining returns the Remaining field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphQuota) GetRemaining() int64 {
	if o == nil || o.Remaining.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Remaining.Get()
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphQuota) GetRemainingOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Remaining.Get(), o.Remaining.IsSet()
}

// HasRemaining returns a boolean if a field has been set.
func (o *OpenGraphQuota) HasRemaining() bool {
	if o != nil && o.Remaining.IsSet() {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given NullableInt64 and assigns it to the Remaining field.
func (o *OpenGraphQuota) SetRemaining(v int64) {
	o.Remaining.Set(&v)
}
// SetRemainingNil sets the value for Remaining to be an explicit nil
func (o *OpenGraphQuota) SetRemainingNil() {
	o.Remaining.Set(nil)
}

// UnsetRemaining ensures that no value is present for Remaining, not even an explicit nil
func (o *OpenGraphQuota) UnsetRemaining() {
	o.Remaining.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphQuota) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphQuota) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *OpenGraphQuota) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *OpenGraphQuota) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *OpenGraphQuota) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *OpenGraphQuota) UnsetState() {
	o.State.Unset()
}

// GetStoragePlanInformation returns the StoragePlanInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphQuota) GetStoragePlanInformation() AnyOfopenGraphStoragePlanInformation {
	if o == nil || o.StoragePlanInformation.Get() == nil {
		var ret AnyOfopenGraphStoragePlanInformation
		return ret
	}
	return *o.StoragePlanInformation.Get()
}

// GetStoragePlanInformationOk returns a tuple with the StoragePlanInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphQuota) GetStoragePlanInformationOk() (*AnyOfopenGraphStoragePlanInformation, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StoragePlanInformation.Get(), o.StoragePlanInformation.IsSet()
}

// HasStoragePlanInformation returns a boolean if a field has been set.
func (o *OpenGraphQuota) HasStoragePlanInformation() bool {
	if o != nil && o.StoragePlanInformation.IsSet() {
		return true
	}

	return false
}

// SetStoragePlanInformation gets a reference to the given NullableAnyOfopenGraphStoragePlanInformation and assigns it to the StoragePlanInformation field.
func (o *OpenGraphQuota) SetStoragePlanInformation(v AnyOfopenGraphStoragePlanInformation) {
	o.StoragePlanInformation.Set(&v)
}
// SetStoragePlanInformationNil sets the value for StoragePlanInformation to be an explicit nil
func (o *OpenGraphQuota) SetStoragePlanInformationNil() {
	o.StoragePlanInformation.Set(nil)
}

// UnsetStoragePlanInformation ensures that no value is present for StoragePlanInformation, not even an explicit nil
func (o *OpenGraphQuota) UnsetStoragePlanInformation() {
	o.StoragePlanInformation.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphQuota) GetTotal() int64 {
	if o == nil || o.Total.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphQuota) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *OpenGraphQuota) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableInt64 and assigns it to the Total field.
func (o *OpenGraphQuota) SetTotal(v int64) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *OpenGraphQuota) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *OpenGraphQuota) UnsetTotal() {
	o.Total.Unset()
}

// GetUsed returns the Used field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphQuota) GetUsed() int64 {
	if o == nil || o.Used.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Used.Get()
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphQuota) GetUsedOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Used.Get(), o.Used.IsSet()
}

// HasUsed returns a boolean if a field has been set.
func (o *OpenGraphQuota) HasUsed() bool {
	if o != nil && o.Used.IsSet() {
		return true
	}

	return false
}

// SetUsed gets a reference to the given NullableInt64 and assigns it to the Used field.
func (o *OpenGraphQuota) SetUsed(v int64) {
	o.Used.Set(&v)
}
// SetUsedNil sets the value for Used to be an explicit nil
func (o *OpenGraphQuota) SetUsedNil() {
	o.Used.Set(nil)
}

// UnsetUsed ensures that no value is present for Used, not even an explicit nil
func (o *OpenGraphQuota) UnsetUsed() {
	o.Used.Unset()
}

func (o OpenGraphQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deleted.IsSet() {
		toSerialize["deleted"] = o.Deleted.Get()
	}
	if o.Remaining.IsSet() {
		toSerialize["remaining"] = o.Remaining.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.StoragePlanInformation.IsSet() {
		toSerialize["storagePlanInformation"] = o.StoragePlanInformation.Get()
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}
	if o.Used.IsSet() {
		toSerialize["used"] = o.Used.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphQuota struct {
	value *OpenGraphQuota
	isSet bool
}

func (v NullableOpenGraphQuota) Get() *OpenGraphQuota {
	return v.value
}

func (v *NullableOpenGraphQuota) Set(val *OpenGraphQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphQuota(val *OpenGraphQuota) *NullableOpenGraphQuota {
	return &NullableOpenGraphQuota{value: val, isSet: true}
}

func (v NullableOpenGraphQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


