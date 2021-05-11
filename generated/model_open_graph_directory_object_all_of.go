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
	"time"
)

// OpenGraphDirectoryObjectAllOf Represents an Active Directory object. The directoryObject type is the base type for many other directory entity types.
type OpenGraphDirectoryObjectAllOf struct {
	DeletedDateTime NullableTime `json:"deletedDateTime,omitempty"`
}

// NewOpenGraphDirectoryObjectAllOf instantiates a new OpenGraphDirectoryObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphDirectoryObjectAllOf() *OpenGraphDirectoryObjectAllOf {
	this := OpenGraphDirectoryObjectAllOf{}
	return &this
}

// NewOpenGraphDirectoryObjectAllOfWithDefaults instantiates a new OpenGraphDirectoryObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphDirectoryObjectAllOfWithDefaults() *OpenGraphDirectoryObjectAllOf {
	this := OpenGraphDirectoryObjectAllOf{}
	return &this
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDirectoryObjectAllOf) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime.Get()
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDirectoryObjectAllOf) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedDateTime.Get(), o.DeletedDateTime.IsSet()
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *OpenGraphDirectoryObjectAllOf) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given NullableTime and assigns it to the DeletedDateTime field.
func (o *OpenGraphDirectoryObjectAllOf) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime.Set(&v)
}
// SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil
func (o *OpenGraphDirectoryObjectAllOf) SetDeletedDateTimeNil() {
	o.DeletedDateTime.Set(nil)
}

// UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
func (o *OpenGraphDirectoryObjectAllOf) UnsetDeletedDateTime() {
	o.DeletedDateTime.Unset()
}

func (o OpenGraphDirectoryObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeletedDateTime.IsSet() {
		toSerialize["deletedDateTime"] = o.DeletedDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphDirectoryObjectAllOf struct {
	value *OpenGraphDirectoryObjectAllOf
	isSet bool
}

func (v NullableOpenGraphDirectoryObjectAllOf) Get() *OpenGraphDirectoryObjectAllOf {
	return v.value
}

func (v *NullableOpenGraphDirectoryObjectAllOf) Set(val *OpenGraphDirectoryObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphDirectoryObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphDirectoryObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphDirectoryObjectAllOf(val *OpenGraphDirectoryObjectAllOf) *NullableOpenGraphDirectoryObjectAllOf {
	return &NullableOpenGraphDirectoryObjectAllOf{value: val, isSet: true}
}

func (v NullableOpenGraphDirectoryObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphDirectoryObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


