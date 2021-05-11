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

// OpenGraphBaseItemAllOf struct for OpenGraphBaseItemAllOf
type OpenGraphBaseItemAllOf struct {
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy NullableAnyOfopenGraphIdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description NullableString `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag NullableString `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy NullableAnyOfopenGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name NullableString `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference NullableAnyOfopenGraphItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl NullableString `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser NullableAnyOfopenGraphUser `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser NullableAnyOfopenGraphUser `json:"lastModifiedByUser,omitempty"`
}

// NewOpenGraphBaseItemAllOf instantiates a new OpenGraphBaseItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphBaseItemAllOf() *OpenGraphBaseItemAllOf {
	this := OpenGraphBaseItemAllOf{}
	return &this
}

// NewOpenGraphBaseItemAllOfWithDefaults instantiates a new OpenGraphBaseItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphBaseItemAllOfWithDefaults() *OpenGraphBaseItemAllOf {
	this := OpenGraphBaseItemAllOf{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetCreatedBy() AnyOfopenGraphIdentitySet {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret AnyOfopenGraphIdentitySet
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetCreatedByOk() (*AnyOfopenGraphIdentitySet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableAnyOfopenGraphIdentitySet and assigns it to the CreatedBy field.
func (o *OpenGraphBaseItemAllOf) SetCreatedBy(v AnyOfopenGraphIdentitySet) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *OpenGraphBaseItemAllOf) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OpenGraphBaseItemAllOf) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetDescription() {
	o.Description.Unset()
}

// GetETag returns the ETag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetETag() string {
	if o == nil || o.ETag.Get() == nil {
		var ret string
		return ret
	}
	return *o.ETag.Get()
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetETagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ETag.Get(), o.ETag.IsSet()
}

// HasETag returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasETag() bool {
	if o != nil && o.ETag.IsSet() {
		return true
	}

	return false
}

// SetETag gets a reference to the given NullableString and assigns it to the ETag field.
func (o *OpenGraphBaseItemAllOf) SetETag(v string) {
	o.ETag.Set(&v)
}
// SetETagNil sets the value for ETag to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetETagNil() {
	o.ETag.Set(nil)
}

// UnsetETag ensures that no value is present for ETag, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetETag() {
	o.ETag.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetLastModifiedBy() AnyOfopenGraphIdentitySet {
	if o == nil || o.LastModifiedBy.Get() == nil {
		var ret AnyOfopenGraphIdentitySet
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetLastModifiedByOk() (*AnyOfopenGraphIdentitySet, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableAnyOfopenGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *OpenGraphBaseItemAllOf) SetLastModifiedBy(v AnyOfopenGraphIdentitySet) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *OpenGraphBaseItemAllOf) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OpenGraphBaseItemAllOf) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetName() {
	o.Name.Unset()
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetParentReference() AnyOfopenGraphItemReference {
	if o == nil || o.ParentReference.Get() == nil {
		var ret AnyOfopenGraphItemReference
		return ret
	}
	return *o.ParentReference.Get()
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetParentReferenceOk() (*AnyOfopenGraphItemReference, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentReference.Get(), o.ParentReference.IsSet()
}

// HasParentReference returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasParentReference() bool {
	if o != nil && o.ParentReference.IsSet() {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given NullableAnyOfopenGraphItemReference and assigns it to the ParentReference field.
func (o *OpenGraphBaseItemAllOf) SetParentReference(v AnyOfopenGraphItemReference) {
	o.ParentReference.Set(&v)
}
// SetParentReferenceNil sets the value for ParentReference to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetParentReferenceNil() {
	o.ParentReference.Set(nil)
}

// UnsetParentReference ensures that no value is present for ParentReference, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetParentReference() {
	o.ParentReference.Unset()
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetWebUrl() string {
	if o == nil || o.WebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasWebUrl() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *OpenGraphBaseItemAllOf) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetWebUrl() {
	o.WebUrl.Unset()
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetCreatedByUser() AnyOfopenGraphUser {
	if o == nil || o.CreatedByUser.Get() == nil {
		var ret AnyOfopenGraphUser
		return ret
	}
	return *o.CreatedByUser.Get()
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetCreatedByUserOk() (*AnyOfopenGraphUser, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedByUser.Get(), o.CreatedByUser.IsSet()
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser.IsSet() {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given NullableAnyOfopenGraphUser and assigns it to the CreatedByUser field.
func (o *OpenGraphBaseItemAllOf) SetCreatedByUser(v AnyOfopenGraphUser) {
	o.CreatedByUser.Set(&v)
}
// SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetCreatedByUserNil() {
	o.CreatedByUser.Set(nil)
}

// UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetCreatedByUser() {
	o.CreatedByUser.Unset()
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphBaseItemAllOf) GetLastModifiedByUser() AnyOfopenGraphUser {
	if o == nil || o.LastModifiedByUser.Get() == nil {
		var ret AnyOfopenGraphUser
		return ret
	}
	return *o.LastModifiedByUser.Get()
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphBaseItemAllOf) GetLastModifiedByUserOk() (*AnyOfopenGraphUser, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedByUser.Get(), o.LastModifiedByUser.IsSet()
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given NullableAnyOfopenGraphUser and assigns it to the LastModifiedByUser field.
func (o *OpenGraphBaseItemAllOf) SetLastModifiedByUser(v AnyOfopenGraphUser) {
	o.LastModifiedByUser.Set(&v)
}
// SetLastModifiedByUserNil sets the value for LastModifiedByUser to be an explicit nil
func (o *OpenGraphBaseItemAllOf) SetLastModifiedByUserNil() {
	o.LastModifiedByUser.Set(nil)
}

// UnsetLastModifiedByUser ensures that no value is present for LastModifiedByUser, not even an explicit nil
func (o *OpenGraphBaseItemAllOf) UnsetLastModifiedByUser() {
	o.LastModifiedByUser.Unset()
}

func (o OpenGraphBaseItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ETag.IsSet() {
		toSerialize["eTag"] = o.ETag.Get()
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ParentReference.IsSet() {
		toSerialize["parentReference"] = o.ParentReference.Get()
	}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	if o.CreatedByUser.IsSet() {
		toSerialize["createdByUser"] = o.CreatedByUser.Get()
	}
	if o.LastModifiedByUser.IsSet() {
		toSerialize["lastModifiedByUser"] = o.LastModifiedByUser.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphBaseItemAllOf struct {
	value *OpenGraphBaseItemAllOf
	isSet bool
}

func (v NullableOpenGraphBaseItemAllOf) Get() *OpenGraphBaseItemAllOf {
	return v.value
}

func (v *NullableOpenGraphBaseItemAllOf) Set(val *OpenGraphBaseItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphBaseItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphBaseItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphBaseItemAllOf(val *OpenGraphBaseItemAllOf) *NullableOpenGraphBaseItemAllOf {
	return &NullableOpenGraphBaseItemAllOf{value: val, isSet: true}
}

func (v NullableOpenGraphBaseItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphBaseItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


