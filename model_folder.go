/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.14.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// Folder Folder metadata, if the item is a folder. Read-only.
type Folder struct {
	// Number of children contained immediately within this container.
	ChildCount *int32 `json:"childCount,omitempty"`
	View *FolderView `json:"view,omitempty"`
}

// NewFolder instantiates a new Folder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolder() *Folder {
	this := Folder{}
	return &this
}

// NewFolderWithDefaults instantiates a new Folder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderWithDefaults() *Folder {
	this := Folder{}
	return &this
}

// GetChildCount returns the ChildCount field value if set, zero value otherwise.
func (o *Folder) GetChildCount() int32 {
	if o == nil || o.ChildCount == nil {
		var ret int32
		return ret
	}
	return *o.ChildCount
}

// GetChildCountOk returns a tuple with the ChildCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetChildCountOk() (*int32, bool) {
	if o == nil || o.ChildCount == nil {
		return nil, false
	}
	return o.ChildCount, true
}

// HasChildCount returns a boolean if a field has been set.
func (o *Folder) HasChildCount() bool {
	if o != nil && o.ChildCount != nil {
		return true
	}

	return false
}

// SetChildCount gets a reference to the given int32 and assigns it to the ChildCount field.
func (o *Folder) SetChildCount(v int32) {
	o.ChildCount = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *Folder) GetView() FolderView {
	if o == nil || o.View == nil {
		var ret FolderView
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetViewOk() (*FolderView, bool) {
	if o == nil || o.View == nil {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *Folder) HasView() bool {
	if o != nil && o.View != nil {
		return true
	}

	return false
}

// SetView gets a reference to the given FolderView and assigns it to the View field.
func (o *Folder) SetView(v FolderView) {
	o.View = &v
}

func (o Folder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChildCount != nil {
		toSerialize["childCount"] = o.ChildCount
	}
	if o.View != nil {
		toSerialize["view"] = o.View
	}
	return json.Marshal(toSerialize)
}

type NullableFolder struct {
	value *Folder
	isSet bool
}

func (v NullableFolder) Get() *Folder {
	return v.value
}

func (v *NullableFolder) Set(val *Folder) {
	v.value = val
	v.isSet = true
}

func (v NullableFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolder(val *Folder) *NullableFolder {
	return &NullableFolder{value: val, isSet: true}
}

func (v NullableFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


