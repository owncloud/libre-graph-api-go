/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"time"
)

// Group struct for Group
type Group struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	// Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
	Description *string `json:"description,omitempty"`
	// The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
	DisplayName *string `json:"displayName,omitempty"`
	// Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// The SMTP address for the group, for example, 'serviceadmins@owncloud.com'. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	Mail *string `json:"mail,omitempty"`
	// Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory.
	OnPremisesDomainName *string `json:"onPremisesDomainName,omitempty"`
	// Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in).
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith). Read-only.
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	// true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// The preferred language for a group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values).
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, not, in).
	SecurityEnabled *bool `json:"securityEnabled,omitempty"`
	// Security identifier of the group, used in Windows scenarios. Returned by default.
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	// Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. It can't be updated later. Other values of visibility can be updated after group creation. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable.
	Visibility *string `json:"visibility,omitempty"`
	CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`
	// Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
	MemberOf *[]DirectoryObject `json:"memberOf,omitempty"`
	// Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), Nullable. Supports $expand.
	Members *[]DirectoryObject `json:"members,omitempty"`
	// The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. Supports $expand.
	Owners *[]DirectoryObject `json:"owners,omitempty"`
	Drive *Drive `json:"drive,omitempty"`
	// The group's drives. Read-only.
	Drives *[]Drive `json:"drives,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
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

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise.
func (o *Group) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.DeletedDateTime == nil {
		return nil, false
	}
	return o.DeletedDateTime, true
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *Group) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime != nil {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.
func (o *Group) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *Group) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *Group) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *Group) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
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

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *Group) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDateTime == nil {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *Group) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime != nil {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *Group) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *Group) GetMail() string {
	if o == nil || o.Mail == nil {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMailOk() (*string, bool) {
	if o == nil || o.Mail == nil {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *Group) HasMail() bool {
	if o != nil && o.Mail != nil {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *Group) SetMail(v string) {
	o.Mail = &v
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

// GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field value if set, zero value otherwise.
func (o *Group) GetOnPremisesLastSyncDateTime() time.Time {
	if o == nil || o.OnPremisesLastSyncDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.OnPremisesLastSyncDateTime
}

// GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool) {
	if o == nil || o.OnPremisesLastSyncDateTime == nil {
		return nil, false
	}
	return o.OnPremisesLastSyncDateTime, true
}

// HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.
func (o *Group) HasOnPremisesLastSyncDateTime() bool {
	if o != nil && o.OnPremisesLastSyncDateTime != nil {
		return true
	}

	return false
}

// SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.
func (o *Group) SetOnPremisesLastSyncDateTime(v time.Time) {
	o.OnPremisesLastSyncDateTime = &v
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

// GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field value if set, zero value otherwise.
func (o *Group) GetOnPremisesSyncEnabled() bool {
	if o == nil || o.OnPremisesSyncEnabled == nil {
		var ret bool
		return ret
	}
	return *o.OnPremisesSyncEnabled
}

// GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOnPremisesSyncEnabledOk() (*bool, bool) {
	if o == nil || o.OnPremisesSyncEnabled == nil {
		return nil, false
	}
	return o.OnPremisesSyncEnabled, true
}

// HasOnPremisesSyncEnabled returns a boolean if a field has been set.
func (o *Group) HasOnPremisesSyncEnabled() bool {
	if o != nil && o.OnPremisesSyncEnabled != nil {
		return true
	}

	return false
}

// SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.
func (o *Group) SetOnPremisesSyncEnabled(v bool) {
	o.OnPremisesSyncEnabled = &v
}

// GetPreferredLanguage returns the PreferredLanguage field value if set, zero value otherwise.
func (o *Group) GetPreferredLanguage() string {
	if o == nil || o.PreferredLanguage == nil {
		var ret string
		return ret
	}
	return *o.PreferredLanguage
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetPreferredLanguageOk() (*string, bool) {
	if o == nil || o.PreferredLanguage == nil {
		return nil, false
	}
	return o.PreferredLanguage, true
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *Group) HasPreferredLanguage() bool {
	if o != nil && o.PreferredLanguage != nil {
		return true
	}

	return false
}

// SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.
func (o *Group) SetPreferredLanguage(v string) {
	o.PreferredLanguage = &v
}

// GetSecurityEnabled returns the SecurityEnabled field value if set, zero value otherwise.
func (o *Group) GetSecurityEnabled() bool {
	if o == nil || o.SecurityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SecurityEnabled
}

// GetSecurityEnabledOk returns a tuple with the SecurityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSecurityEnabledOk() (*bool, bool) {
	if o == nil || o.SecurityEnabled == nil {
		return nil, false
	}
	return o.SecurityEnabled, true
}

// HasSecurityEnabled returns a boolean if a field has been set.
func (o *Group) HasSecurityEnabled() bool {
	if o != nil && o.SecurityEnabled != nil {
		return true
	}

	return false
}

// SetSecurityEnabled gets a reference to the given bool and assigns it to the SecurityEnabled field.
func (o *Group) SetSecurityEnabled(v bool) {
	o.SecurityEnabled = &v
}

// GetSecurityIdentifier returns the SecurityIdentifier field value if set, zero value otherwise.
func (o *Group) GetSecurityIdentifier() string {
	if o == nil || o.SecurityIdentifier == nil {
		var ret string
		return ret
	}
	return *o.SecurityIdentifier
}

// GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSecurityIdentifierOk() (*string, bool) {
	if o == nil || o.SecurityIdentifier == nil {
		return nil, false
	}
	return o.SecurityIdentifier, true
}

// HasSecurityIdentifier returns a boolean if a field has been set.
func (o *Group) HasSecurityIdentifier() bool {
	if o != nil && o.SecurityIdentifier != nil {
		return true
	}

	return false
}

// SetSecurityIdentifier gets a reference to the given string and assigns it to the SecurityIdentifier field.
func (o *Group) SetSecurityIdentifier(v string) {
	o.SecurityIdentifier = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *Group) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *Group) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *Group) SetVisibility(v string) {
	o.Visibility = &v
}

// GetCreatedOnBehalfOf returns the CreatedOnBehalfOf field value if set, zero value otherwise.
func (o *Group) GetCreatedOnBehalfOf() DirectoryObject {
	if o == nil || o.CreatedOnBehalfOf == nil {
		var ret DirectoryObject
		return ret
	}
	return *o.CreatedOnBehalfOf
}

// GetCreatedOnBehalfOfOk returns a tuple with the CreatedOnBehalfOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetCreatedOnBehalfOfOk() (*DirectoryObject, bool) {
	if o == nil || o.CreatedOnBehalfOf == nil {
		return nil, false
	}
	return o.CreatedOnBehalfOf, true
}

// HasCreatedOnBehalfOf returns a boolean if a field has been set.
func (o *Group) HasCreatedOnBehalfOf() bool {
	if o != nil && o.CreatedOnBehalfOf != nil {
		return true
	}

	return false
}

// SetCreatedOnBehalfOf gets a reference to the given DirectoryObject and assigns it to the CreatedOnBehalfOf field.
func (o *Group) SetCreatedOnBehalfOf(v DirectoryObject) {
	o.CreatedOnBehalfOf = &v
}

// GetMemberOf returns the MemberOf field value if set, zero value otherwise.
func (o *Group) GetMemberOf() []DirectoryObject {
	if o == nil || o.MemberOf == nil {
		var ret []DirectoryObject
		return ret
	}
	return *o.MemberOf
}

// GetMemberOfOk returns a tuple with the MemberOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMemberOfOk() (*[]DirectoryObject, bool) {
	if o == nil || o.MemberOf == nil {
		return nil, false
	}
	return o.MemberOf, true
}

// HasMemberOf returns a boolean if a field has been set.
func (o *Group) HasMemberOf() bool {
	if o != nil && o.MemberOf != nil {
		return true
	}

	return false
}

// SetMemberOf gets a reference to the given []DirectoryObject and assigns it to the MemberOf field.
func (o *Group) SetMemberOf(v []DirectoryObject) {
	o.MemberOf = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Group) GetMembers() []DirectoryObject {
	if o == nil || o.Members == nil {
		var ret []DirectoryObject
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetMembersOk() (*[]DirectoryObject, bool) {
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

// SetMembers gets a reference to the given []DirectoryObject and assigns it to the Members field.
func (o *Group) SetMembers(v []DirectoryObject) {
	o.Members = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *Group) GetOwners() []DirectoryObject {
	if o == nil || o.Owners == nil {
		var ret []DirectoryObject
		return ret
	}
	return *o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOwnersOk() (*[]DirectoryObject, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *Group) HasOwners() bool {
	if o != nil && o.Owners != nil {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []DirectoryObject and assigns it to the Owners field.
func (o *Group) SetOwners(v []DirectoryObject) {
	o.Owners = &v
}

// GetDrive returns the Drive field value if set, zero value otherwise.
func (o *Group) GetDrive() Drive {
	if o == nil || o.Drive == nil {
		var ret Drive
		return ret
	}
	return *o.Drive
}

// GetDriveOk returns a tuple with the Drive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDriveOk() (*Drive, bool) {
	if o == nil || o.Drive == nil {
		return nil, false
	}
	return o.Drive, true
}

// HasDrive returns a boolean if a field has been set.
func (o *Group) HasDrive() bool {
	if o != nil && o.Drive != nil {
		return true
	}

	return false
}

// SetDrive gets a reference to the given Drive and assigns it to the Drive field.
func (o *Group) SetDrive(v Drive) {
	o.Drive = &v
}

// GetDrives returns the Drives field value if set, zero value otherwise.
func (o *Group) GetDrives() []Drive {
	if o == nil || o.Drives == nil {
		var ret []Drive
		return ret
	}
	return *o.Drives
}

// GetDrivesOk returns a tuple with the Drives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDrivesOk() (*[]Drive, bool) {
	if o == nil || o.Drives == nil {
		return nil, false
	}
	return o.Drives, true
}

// HasDrives returns a boolean if a field has been set.
func (o *Group) HasDrives() bool {
	if o != nil && o.Drives != nil {
		return true
	}

	return false
}

// SetDrives gets a reference to the given []Drive and assigns it to the Drives field.
func (o *Group) SetDrives(v []Drive) {
	o.Drives = &v
}

// GetIsArchived returns the IsArchived field value if set, zero value otherwise.
func (o *Group) GetIsArchived() bool {
	if o == nil || o.IsArchived == nil {
		var ret bool
		return ret
	}
	return *o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIsArchivedOk() (*bool, bool) {
	if o == nil || o.IsArchived == nil {
		return nil, false
	}
	return o.IsArchived, true
}

// HasIsArchived returns a boolean if a field has been set.
func (o *Group) HasIsArchived() bool {
	if o != nil && o.IsArchived != nil {
		return true
	}

	return false
}

// SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.
func (o *Group) SetIsArchived(v bool) {
	o.IsArchived = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeletedDateTime != nil {
		toSerialize["deletedDateTime"] = o.DeletedDateTime
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ExpirationDateTime != nil {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if o.Mail != nil {
		toSerialize["mail"] = o.Mail
	}
	if o.OnPremisesDomainName != nil {
		toSerialize["onPremisesDomainName"] = o.OnPremisesDomainName
	}
	if o.OnPremisesLastSyncDateTime != nil {
		toSerialize["onPremisesLastSyncDateTime"] = o.OnPremisesLastSyncDateTime
	}
	if o.OnPremisesSamAccountName != nil {
		toSerialize["onPremisesSamAccountName"] = o.OnPremisesSamAccountName
	}
	if o.OnPremisesSyncEnabled != nil {
		toSerialize["onPremisesSyncEnabled"] = o.OnPremisesSyncEnabled
	}
	if o.PreferredLanguage != nil {
		toSerialize["preferredLanguage"] = o.PreferredLanguage
	}
	if o.SecurityEnabled != nil {
		toSerialize["securityEnabled"] = o.SecurityEnabled
	}
	if o.SecurityIdentifier != nil {
		toSerialize["securityIdentifier"] = o.SecurityIdentifier
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.CreatedOnBehalfOf != nil {
		toSerialize["createdOnBehalfOf"] = o.CreatedOnBehalfOf
	}
	if o.MemberOf != nil {
		toSerialize["memberOf"] = o.MemberOf
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if o.Drive != nil {
		toSerialize["drive"] = o.Drive
	}
	if o.Drives != nil {
		toSerialize["drives"] = o.Drives
	}
	if o.IsArchived != nil {
		toSerialize["isArchived"] = o.IsArchived
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


