/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OdataErrorDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OdataErrorDetail{}

// OdataErrorDetail struct for OdataErrorDetail
type OdataErrorDetail struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Target  *string `json:"target,omitempty"`
}

type _OdataErrorDetail OdataErrorDetail

// NewOdataErrorDetail instantiates a new OdataErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOdataErrorDetail(code string, message string) *OdataErrorDetail {
	this := OdataErrorDetail{}
	this.Code = code
	this.Message = message
	return &this
}

// NewOdataErrorDetailWithDefaults instantiates a new OdataErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOdataErrorDetailWithDefaults() *OdataErrorDetail {
	this := OdataErrorDetail{}
	return &this
}

// GetCode returns the Code field value
func (o *OdataErrorDetail) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OdataErrorDetail) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *OdataErrorDetail) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OdataErrorDetail) SetMessage(v string) {
	o.Message = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *OdataErrorDetail) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *OdataErrorDetail) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *OdataErrorDetail) SetTarget(v string) {
	o.Target = &v
}

func (o OdataErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OdataErrorDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

func (o *OdataErrorDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOdataErrorDetail := _OdataErrorDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOdataErrorDetail)

	if err != nil {
		return err
	}

	*o = OdataErrorDetail(varOdataErrorDetail)

	return err
}

type NullableOdataErrorDetail struct {
	value *OdataErrorDetail
	isSet bool
}

func (v NullableOdataErrorDetail) Get() *OdataErrorDetail {
	return v.value
}

func (v *NullableOdataErrorDetail) Set(val *OdataErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableOdataErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableOdataErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdataErrorDetail(val *OdataErrorDetail) *NullableOdataErrorDetail {
	return &NullableOdataErrorDetail{value: val, isSet: true}
}

func (v NullableOdataErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdataErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
