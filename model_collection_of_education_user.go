/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// CollectionOfEducationUser struct for CollectionOfEducationUser
type CollectionOfEducationUser struct {
	Value []EducationUser `json:"value,omitempty"`
}

// NewCollectionOfEducationUser instantiates a new CollectionOfEducationUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfEducationUser() *CollectionOfEducationUser {
	this := CollectionOfEducationUser{}
	return &this
}

// NewCollectionOfEducationUserWithDefaults instantiates a new CollectionOfEducationUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfEducationUserWithDefaults() *CollectionOfEducationUser {
	this := CollectionOfEducationUser{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfEducationUser) GetValue() []EducationUser {
	if o == nil || o.Value == nil {
		var ret []EducationUser
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEducationUser) GetValueOk() ([]EducationUser, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfEducationUser) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []EducationUser and assigns it to the Value field.
func (o *CollectionOfEducationUser) SetValue(v []EducationUser) {
	o.Value = v
}

func (o CollectionOfEducationUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfEducationUser struct {
	value *CollectionOfEducationUser
	isSet bool
}

func (v NullableCollectionOfEducationUser) Get() *CollectionOfEducationUser {
	return v.value
}

func (v *NullableCollectionOfEducationUser) Set(val *CollectionOfEducationUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfEducationUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfEducationUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfEducationUser(val *CollectionOfEducationUser) *NullableCollectionOfEducationUser {
	return &NullableCollectionOfEducationUser{value: val, isSet: true}
}

func (v NullableCollectionOfEducationUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfEducationUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


