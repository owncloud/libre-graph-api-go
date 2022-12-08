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

// CollectionOfSchools struct for CollectionOfSchools
type CollectionOfSchools struct {
	Value []EducationSchool `json:"value,omitempty"`
}

// NewCollectionOfSchools instantiates a new CollectionOfSchools object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfSchools() *CollectionOfSchools {
	this := CollectionOfSchools{}
	return &this
}

// NewCollectionOfSchoolsWithDefaults instantiates a new CollectionOfSchools object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfSchoolsWithDefaults() *CollectionOfSchools {
	this := CollectionOfSchools{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfSchools) GetValue() []EducationSchool {
	if o == nil || o.Value == nil {
		var ret []EducationSchool
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSchools) GetValueOk() ([]EducationSchool, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfSchools) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []EducationSchool and assigns it to the Value field.
func (o *CollectionOfSchools) SetValue(v []EducationSchool) {
	o.Value = v
}

func (o CollectionOfSchools) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfSchools struct {
	value *CollectionOfSchools
	isSet bool
}

func (v NullableCollectionOfSchools) Get() *CollectionOfSchools {
	return v.value
}

func (v *NullableCollectionOfSchools) Set(val *CollectionOfSchools) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfSchools) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfSchools) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfSchools(val *CollectionOfSchools) *NullableCollectionOfSchools {
	return &NullableCollectionOfSchools{value: val, isSet: true}
}

func (v NullableCollectionOfSchools) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfSchools) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


