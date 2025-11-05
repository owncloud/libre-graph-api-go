# EducationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] [readonly] 
**AccountEnabled** | Pointer to **bool** | Set to \&quot;true\&quot; when the account is enabled. | [optional] 
**DisplayName** | Pointer to **string** | The name displayed in the address book for the user. This value is usually the combination of the user&#39;s first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Returned by default. Supports $orderby. | [optional] 
**Drives** | Pointer to [**[]Drive**](Drive.md) | A collection of drives available for this user. Read-only. | [optional] [readonly] 
**Drive** | Pointer to [**Drive**](Drive.md) |  | [optional] 
**Identities** | Pointer to [**[]ObjectIdentity**](ObjectIdentity.md) | Identities associated with this account. | [optional] 
**Mail** | Pointer to **string** | The SMTP address for the user, for example, &#39;jeff@contoso.onowncloud.com&#39;. Returned by default. | [optional] 
**MemberOf** | Pointer to [**[]Group**](Group.md) | Groups that this user is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand. | [optional] 
**OnPremisesSamAccountName** | Pointer to **string** | Contains the on-premises SAM account name synchronized from the on-premises directory. Read-only. | [optional] 
**PasswordProfile** | Pointer to [**PasswordProfile**](PasswordProfile.md) |  | [optional] 
**Surname** | Pointer to **string** | The user&#39;s surname (family name or last name). Returned by default. | [optional] 
**GivenName** | Pointer to **string** | The user&#39;s givenName. Returned by default. | [optional] 
**PrimaryRole** | Pointer to **string** | The user&#x60;s default role. Such as \&quot;student\&quot; or \&quot;teacher\&quot; | [optional] 
**UserType** | Pointer to **string** | The user&#x60;s type. This can be either \&quot;Member\&quot; for regular user, \&quot;Guest\&quot; for guest users or \&quot;Federated\&quot; for users imported from a federated instance. | [optional] 
**ExternalID** | Pointer to **string** | A unique identifier for the user assigned by the school or institution. | [optional] 

## Methods

### NewEducationUser

`func NewEducationUser() *EducationUser`

NewEducationUser instantiates a new EducationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationUserWithDefaults

`func NewEducationUserWithDefaults() *EducationUser`

NewEducationUserWithDefaults instantiates a new EducationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EducationUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EducationUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EducationUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EducationUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountEnabled

`func (o *EducationUser) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *EducationUser) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *EducationUser) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *EducationUser) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### GetDisplayName

`func (o *EducationUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EducationUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EducationUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EducationUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDrives

`func (o *EducationUser) GetDrives() []Drive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *EducationUser) GetDrivesOk() (*[]Drive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *EducationUser) SetDrives(v []Drive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *EducationUser) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetDrive

`func (o *EducationUser) GetDrive() Drive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *EducationUser) GetDriveOk() (*Drive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *EducationUser) SetDrive(v Drive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *EducationUser) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### GetIdentities

`func (o *EducationUser) GetIdentities() []ObjectIdentity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *EducationUser) GetIdentitiesOk() (*[]ObjectIdentity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *EducationUser) SetIdentities(v []ObjectIdentity)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *EducationUser) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetMail

`func (o *EducationUser) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *EducationUser) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *EducationUser) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *EducationUser) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetMemberOf

`func (o *EducationUser) GetMemberOf() []Group`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *EducationUser) GetMemberOfOk() (*[]Group, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *EducationUser) SetMemberOf(v []Group)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *EducationUser) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetOnPremisesSamAccountName

`func (o *EducationUser) GetOnPremisesSamAccountName() string`

GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field if non-nil, zero value otherwise.

### GetOnPremisesSamAccountNameOk

`func (o *EducationUser) GetOnPremisesSamAccountNameOk() (*string, bool)`

GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSamAccountName

`func (o *EducationUser) SetOnPremisesSamAccountName(v string)`

SetOnPremisesSamAccountName sets OnPremisesSamAccountName field to given value.

### HasOnPremisesSamAccountName

`func (o *EducationUser) HasOnPremisesSamAccountName() bool`

HasOnPremisesSamAccountName returns a boolean if a field has been set.

### GetPasswordProfile

`func (o *EducationUser) GetPasswordProfile() PasswordProfile`

GetPasswordProfile returns the PasswordProfile field if non-nil, zero value otherwise.

### GetPasswordProfileOk

`func (o *EducationUser) GetPasswordProfileOk() (*PasswordProfile, bool)`

GetPasswordProfileOk returns a tuple with the PasswordProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProfile

`func (o *EducationUser) SetPasswordProfile(v PasswordProfile)`

SetPasswordProfile sets PasswordProfile field to given value.

### HasPasswordProfile

`func (o *EducationUser) HasPasswordProfile() bool`

HasPasswordProfile returns a boolean if a field has been set.

### GetSurname

`func (o *EducationUser) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *EducationUser) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *EducationUser) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *EducationUser) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetGivenName

`func (o *EducationUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *EducationUser) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *EducationUser) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *EducationUser) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetPrimaryRole

`func (o *EducationUser) GetPrimaryRole() string`

GetPrimaryRole returns the PrimaryRole field if non-nil, zero value otherwise.

### GetPrimaryRoleOk

`func (o *EducationUser) GetPrimaryRoleOk() (*string, bool)`

GetPrimaryRoleOk returns a tuple with the PrimaryRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRole

`func (o *EducationUser) SetPrimaryRole(v string)`

SetPrimaryRole sets PrimaryRole field to given value.

### HasPrimaryRole

`func (o *EducationUser) HasPrimaryRole() bool`

HasPrimaryRole returns a boolean if a field has been set.

### GetUserType

`func (o *EducationUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *EducationUser) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *EducationUser) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *EducationUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetExternalID

`func (o *EducationUser) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *EducationUser) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *EducationUser) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *EducationUser) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


