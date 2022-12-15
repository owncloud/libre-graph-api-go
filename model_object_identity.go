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

// ObjectIdentity Represents an identity used to sign in to a user account
type ObjectIdentity struct {
	// domain of the Provider issuing the identity
	Issuer *string `json:"issuer,omitempty"`
	// The unique id assigned by the issuer to the account
	IssuerAssignedId *string `json:"issuerAssignedId,omitempty"`
}

// NewObjectIdentity instantiates a new ObjectIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectIdentity() *ObjectIdentity {
	this := ObjectIdentity{}
	return &this
}

// NewObjectIdentityWithDefaults instantiates a new ObjectIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectIdentityWithDefaults() *ObjectIdentity {
	this := ObjectIdentity{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *ObjectIdentity) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectIdentity) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *ObjectIdentity) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *ObjectIdentity) SetIssuer(v string) {
	o.Issuer = &v
}

// GetIssuerAssignedId returns the IssuerAssignedId field value if set, zero value otherwise.
func (o *ObjectIdentity) GetIssuerAssignedId() string {
	if o == nil || o.IssuerAssignedId == nil {
		var ret string
		return ret
	}
	return *o.IssuerAssignedId
}

// GetIssuerAssignedIdOk returns a tuple with the IssuerAssignedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectIdentity) GetIssuerAssignedIdOk() (*string, bool) {
	if o == nil || o.IssuerAssignedId == nil {
		return nil, false
	}
	return o.IssuerAssignedId, true
}

// HasIssuerAssignedId returns a boolean if a field has been set.
func (o *ObjectIdentity) HasIssuerAssignedId() bool {
	if o != nil && o.IssuerAssignedId != nil {
		return true
	}

	return false
}

// SetIssuerAssignedId gets a reference to the given string and assigns it to the IssuerAssignedId field.
func (o *ObjectIdentity) SetIssuerAssignedId(v string) {
	o.IssuerAssignedId = &v
}

func (o ObjectIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.IssuerAssignedId != nil {
		toSerialize["issuerAssignedId"] = o.IssuerAssignedId
	}
	return json.Marshal(toSerialize)
}

type NullableObjectIdentity struct {
	value *ObjectIdentity
	isSet bool
}

func (v NullableObjectIdentity) Get() *ObjectIdentity {
	return v.value
}

func (v *NullableObjectIdentity) Set(val *ObjectIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectIdentity(val *ObjectIdentity) *NullableObjectIdentity {
	return &NullableObjectIdentity{value: val, isSet: true}
}

func (v NullableObjectIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

