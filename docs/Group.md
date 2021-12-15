# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] [readonly] 
**DeletedDateTime** | Pointer to **time.Time** |  | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only. | [optional] 
**Description** | Pointer to **NullableString** | An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | [optional] 
**ExpirationDateTime** | Pointer to **NullableTime** | Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only. | [optional] 
**Mail** | Pointer to **NullableString** | The SMTP address for the group, for example, &#39;serviceadmins@owncloud.com&#39;. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**OnPremisesDomainName** | Pointer to **NullableString** | Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to **NullableTime** | Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in). | [optional] 
**OnPremisesSamAccountName** | Pointer to **NullableString** | Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith). Read-only. | [optional] 
**OnPremisesSyncEnabled** | Pointer to **NullableBool** | true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, not, in, and eq on null values). | [optional] 
**PreferredLanguage** | Pointer to **NullableString** | The preferred language for a group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**SecurityEnabled** | Pointer to **NullableBool** | Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, not, in). | [optional] 
**SecurityIdentifier** | Pointer to **NullableString** | Security identifier of the group, used in Windows scenarios. Returned by default. | [optional] 
**Visibility** | Pointer to **NullableString** | Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. It can&#39;t be updated later. Other values of visibility can be updated after group creation. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable. | [optional] 
**CreatedOnBehalfOf** | Pointer to [**NullableDirectoryObject**](DirectoryObject.md) |  | [optional] 
**MemberOf** | Pointer to [**[]DirectoryObject**](DirectoryObject.md) | Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand. | [optional] 
**Members** | Pointer to [**[]DirectoryObject**](DirectoryObject.md) | Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), Nullable. Supports $expand. | [optional] 
**Owners** | Pointer to [**[]DirectoryObject**](DirectoryObject.md) | The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. Supports $expand. | [optional] 
**Drive** | Pointer to [**Drive**](Drive.md) |  | [optional] 
**Drives** | Pointer to [**[]Drive**](Drive.md) | The group&#39;s drives. Read-only. | [optional] 
**IsArchived** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *Group) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *Group) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *Group) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *Group) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *Group) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Group) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Group) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Group) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Group) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Group) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Group) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Group) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *Group) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Group) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Group) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Group) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Group) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Group) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExpirationDateTime

`func (o *Group) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *Group) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *Group) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *Group) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *Group) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *Group) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetMail

`func (o *Group) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *Group) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *Group) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *Group) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMailNil

`func (o *Group) SetMailNil(b bool)`

 SetMailNil sets the value for Mail to be an explicit nil

### UnsetMail
`func (o *Group) UnsetMail()`

UnsetMail ensures that no value is present for Mail, not even an explicit nil
### GetOnPremisesDomainName

`func (o *Group) GetOnPremisesDomainName() string`

GetOnPremisesDomainName returns the OnPremisesDomainName field if non-nil, zero value otherwise.

### GetOnPremisesDomainNameOk

`func (o *Group) GetOnPremisesDomainNameOk() (*string, bool)`

GetOnPremisesDomainNameOk returns a tuple with the OnPremisesDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesDomainName

`func (o *Group) SetOnPremisesDomainName(v string)`

SetOnPremisesDomainName sets OnPremisesDomainName field to given value.

### HasOnPremisesDomainName

`func (o *Group) HasOnPremisesDomainName() bool`

HasOnPremisesDomainName returns a boolean if a field has been set.

### SetOnPremisesDomainNameNil

`func (o *Group) SetOnPremisesDomainNameNil(b bool)`

 SetOnPremisesDomainNameNil sets the value for OnPremisesDomainName to be an explicit nil

### UnsetOnPremisesDomainName
`func (o *Group) UnsetOnPremisesDomainName()`

UnsetOnPremisesDomainName ensures that no value is present for OnPremisesDomainName, not even an explicit nil
### GetOnPremisesLastSyncDateTime

`func (o *Group) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *Group) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *Group) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *Group) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *Group) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *Group) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesSamAccountName

`func (o *Group) GetOnPremisesSamAccountName() string`

GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field if non-nil, zero value otherwise.

### GetOnPremisesSamAccountNameOk

`func (o *Group) GetOnPremisesSamAccountNameOk() (*string, bool)`

GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSamAccountName

`func (o *Group) SetOnPremisesSamAccountName(v string)`

SetOnPremisesSamAccountName sets OnPremisesSamAccountName field to given value.

### HasOnPremisesSamAccountName

`func (o *Group) HasOnPremisesSamAccountName() bool`

HasOnPremisesSamAccountName returns a boolean if a field has been set.

### SetOnPremisesSamAccountNameNil

`func (o *Group) SetOnPremisesSamAccountNameNil(b bool)`

 SetOnPremisesSamAccountNameNil sets the value for OnPremisesSamAccountName to be an explicit nil

### UnsetOnPremisesSamAccountName
`func (o *Group) UnsetOnPremisesSamAccountName()`

UnsetOnPremisesSamAccountName ensures that no value is present for OnPremisesSamAccountName, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *Group) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *Group) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *Group) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *Group) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *Group) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *Group) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetPreferredLanguage

`func (o *Group) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *Group) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *Group) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *Group) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *Group) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *Group) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetSecurityEnabled

`func (o *Group) GetSecurityEnabled() bool`

GetSecurityEnabled returns the SecurityEnabled field if non-nil, zero value otherwise.

### GetSecurityEnabledOk

`func (o *Group) GetSecurityEnabledOk() (*bool, bool)`

GetSecurityEnabledOk returns a tuple with the SecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEnabled

`func (o *Group) SetSecurityEnabled(v bool)`

SetSecurityEnabled sets SecurityEnabled field to given value.

### HasSecurityEnabled

`func (o *Group) HasSecurityEnabled() bool`

HasSecurityEnabled returns a boolean if a field has been set.

### SetSecurityEnabledNil

`func (o *Group) SetSecurityEnabledNil(b bool)`

 SetSecurityEnabledNil sets the value for SecurityEnabled to be an explicit nil

### UnsetSecurityEnabled
`func (o *Group) UnsetSecurityEnabled()`

UnsetSecurityEnabled ensures that no value is present for SecurityEnabled, not even an explicit nil
### GetSecurityIdentifier

`func (o *Group) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *Group) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *Group) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *Group) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### SetSecurityIdentifierNil

`func (o *Group) SetSecurityIdentifierNil(b bool)`

 SetSecurityIdentifierNil sets the value for SecurityIdentifier to be an explicit nil

### UnsetSecurityIdentifier
`func (o *Group) UnsetSecurityIdentifier()`

UnsetSecurityIdentifier ensures that no value is present for SecurityIdentifier, not even an explicit nil
### GetVisibility

`func (o *Group) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Group) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Group) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Group) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *Group) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *Group) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetCreatedOnBehalfOf

`func (o *Group) GetCreatedOnBehalfOf() DirectoryObject`

GetCreatedOnBehalfOf returns the CreatedOnBehalfOf field if non-nil, zero value otherwise.

### GetCreatedOnBehalfOfOk

`func (o *Group) GetCreatedOnBehalfOfOk() (*DirectoryObject, bool)`

GetCreatedOnBehalfOfOk returns a tuple with the CreatedOnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOnBehalfOf

`func (o *Group) SetCreatedOnBehalfOf(v DirectoryObject)`

SetCreatedOnBehalfOf sets CreatedOnBehalfOf field to given value.

### HasCreatedOnBehalfOf

`func (o *Group) HasCreatedOnBehalfOf() bool`

HasCreatedOnBehalfOf returns a boolean if a field has been set.

### SetCreatedOnBehalfOfNil

`func (o *Group) SetCreatedOnBehalfOfNil(b bool)`

 SetCreatedOnBehalfOfNil sets the value for CreatedOnBehalfOf to be an explicit nil

### UnsetCreatedOnBehalfOf
`func (o *Group) UnsetCreatedOnBehalfOf()`

UnsetCreatedOnBehalfOf ensures that no value is present for CreatedOnBehalfOf, not even an explicit nil
### GetMemberOf

`func (o *Group) GetMemberOf() []DirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *Group) GetMemberOfOk() (*[]DirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *Group) SetMemberOf(v []DirectoryObject)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *Group) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetMembers

`func (o *Group) GetMembers() []DirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Group) GetMembersOk() (*[]DirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Group) SetMembers(v []DirectoryObject)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Group) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetOwners

`func (o *Group) GetOwners() []DirectoryObject`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *Group) GetOwnersOk() (*[]DirectoryObject, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *Group) SetOwners(v []DirectoryObject)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *Group) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetDrive

`func (o *Group) GetDrive() Drive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *Group) GetDriveOk() (*Drive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *Group) SetDrive(v Drive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *Group) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### GetDrives

`func (o *Group) GetDrives() []Drive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *Group) GetDrivesOk() (*[]Drive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *Group) SetDrives(v []Drive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *Group) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetIsArchived

`func (o *Group) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Group) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Group) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *Group) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchivedNil

`func (o *Group) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *Group) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


