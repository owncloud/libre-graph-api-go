/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// CollectionOfEducationClass struct for CollectionOfEducationClass
type CollectionOfEducationClass struct {
	Value []EducationClass `json:"value,omitempty"`
}

// NewCollectionOfEducationClass instantiates a new CollectionOfEducationClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfEducationClass() *CollectionOfEducationClass {
	this := CollectionOfEducationClass{}
	return &this
}

// NewCollectionOfEducationClassWithDefaults instantiates a new CollectionOfEducationClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfEducationClassWithDefaults() *CollectionOfEducationClass {
	this := CollectionOfEducationClass{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfEducationClass) GetValue() []EducationClass {
	if o == nil || o.Value == nil {
		var ret []EducationClass
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEducationClass) GetValueOk() ([]EducationClass, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfEducationClass) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []EducationClass and assigns it to the Value field.
func (o *CollectionOfEducationClass) SetValue(v []EducationClass) {
	o.Value = v
}

func (o CollectionOfEducationClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfEducationClass struct {
	value *CollectionOfEducationClass
	isSet bool
}

func (v NullableCollectionOfEducationClass) Get() *CollectionOfEducationClass {
	return v.value
}

func (v *NullableCollectionOfEducationClass) Set(val *CollectionOfEducationClass) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfEducationClass) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfEducationClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfEducationClass(val *CollectionOfEducationClass) *NullableCollectionOfEducationClass {
	return &NullableCollectionOfEducationClass{value: val, isSet: true}
}

func (v NullableCollectionOfEducationClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfEducationClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
