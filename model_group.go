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

// Group struct for Group
type Group struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
	Description *string `json:"description,omitempty"`
	// The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
	DisplayName *string `json:"displayName,omitempty"`
	// Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), Nullable. Supports $expand.
	Members []User `json:"members,omitempty"`
	// Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select.
	OnPremisesDomainName *string `json:"onPremisesDomainName,omitempty"`
	// Contains the on-premises SAM account name synchronized from the on-premises directory. Read-only.
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	// A list of member references to the members to be added. Up to 20 members can be added with a single request
	MembersodataBind []string `json:"members@odata.bind,omitempty"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Group) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Group) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Group) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Group) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Group) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Group) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Group) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Group) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Group) GetMembers() []User {
	if o == nil || o.Members == nil {
		var ret []User
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMembersOk() ([]User, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *Group) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []User and assigns it to the Members field.
func (o *Group) SetMembers(v []User) {
	o.Members = v
}

// GetOnPremisesDomainName returns the OnPremisesDomainName field value if set, zero value otherwise.
func (o *Group) GetOnPremisesDomainName() string {
	if o == nil || o.OnPremisesDomainName == nil {
		var ret string
		return ret
	}
	return *o.OnPremisesDomainName
}

// GetOnPremisesDomainNameOk returns a tuple with the OnPremisesDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOnPremisesDomainNameOk() (*string, bool) {
	if o == nil || o.OnPremisesDomainName == nil {
		return nil, false
	}
	return o.OnPremisesDomainName, true
}

// HasOnPremisesDomainName returns a boolean if a field has been set.
func (o *Group) HasOnPremisesDomainName() bool {
	if o != nil && o.OnPremisesDomainName != nil {
		return true
	}

	return false
}

// SetOnPremisesDomainName gets a reference to the given string and assigns it to the OnPremisesDomainName field.
func (o *Group) SetOnPremisesDomainName(v string) {
	o.OnPremisesDomainName = &v
}

// GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field value if set, zero value otherwise.
func (o *Group) GetOnPremisesSamAccountName() string {
	if o == nil || o.OnPremisesSamAccountName == nil {
		var ret string
		return ret
	}
	return *o.OnPremisesSamAccountName
}

// GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOnPremisesSamAccountNameOk() (*string, bool) {
	if o == nil || o.OnPremisesSamAccountName == nil {
		return nil, false
	}
	return o.OnPremisesSamAccountName, true
}

// HasOnPremisesSamAccountName returns a boolean if a field has been set.
func (o *Group) HasOnPremisesSamAccountName() bool {
	if o != nil && o.OnPremisesSamAccountName != nil {
		return true
	}

	return false
}

// SetOnPremisesSamAccountName gets a reference to the given string and assigns it to the OnPremisesSamAccountName field.
func (o *Group) SetOnPremisesSamAccountName(v string) {
	o.OnPremisesSamAccountName = &v
}

// GetMembersodataBind returns the MembersodataBind field value if set, zero value otherwise.
func (o *Group) GetMembersodataBind() []string {
	if o == nil || o.MembersodataBind == nil {
		var ret []string
		return ret
	}
	return o.MembersodataBind
}

// GetMembersodataBindOk returns a tuple with the MembersodataBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMembersodataBindOk() ([]string, bool) {
	if o == nil || o.MembersodataBind == nil {
		return nil, false
	}
	return o.MembersodataBind, true
}

// HasMembersodataBind returns a boolean if a field has been set.
func (o *Group) HasMembersodataBind() bool {
	if o != nil && o.MembersodataBind != nil {
		return true
	}

	return false
}

// SetMembersodataBind gets a reference to the given []string and assigns it to the MembersodataBind field.
func (o *Group) SetMembersodataBind(v []string) {
	o.MembersodataBind = v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.OnPremisesDomainName != nil {
		toSerialize["onPremisesDomainName"] = o.OnPremisesDomainName
	}
	if o.OnPremisesSamAccountName != nil {
		toSerialize["onPremisesSamAccountName"] = o.OnPremisesSamAccountName
	}
	if o.MembersodataBind != nil {
		toSerialize["members@odata.bind"] = o.MembersodataBind
	}
	return json.Marshal(toSerialize)
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


