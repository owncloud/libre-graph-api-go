# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] [readonly] 
**AccountEnabled** | Pointer to **bool** | Set to \&quot;true\&quot; when the account is enabled. | [optional] 
**AppRoleAssignments** | Pointer to [**[]AppRoleAssignment**](AppRoleAssignment.md) | The apps and app roles which this user has been assigned. | [optional] [readonly] 
**DisplayName** | **string** | The name displayed in the address book for the user. This value is usually the combination of the user&#39;s first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Returned by default. Supports $orderby. | 
**Drives** | Pointer to [**[]Drive**](Drive.md) | A collection of drives available for this user. Read-only. | [optional] [readonly] 
**Drive** | Pointer to [**Drive**](Drive.md) |  | [optional] 
**Identities** | Pointer to [**[]ObjectIdentity**](ObjectIdentity.md) | Identities associated with this account. | [optional] 
**Mail** | Pointer to **string** | The SMTP address for the user, for example, &#39;jeff@contoso.onowncloud.com&#39;. Returned by default. | [optional] 
**MemberOf** | Pointer to [**[]Group**](Group.md) | Groups that this user is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand. | [optional] [readonly] 
**OnPremisesSamAccountName** | **string** | Contains the on-premises SAM account name synchronized from the on-premises directory. | 
**PasswordProfile** | Pointer to [**PasswordProfile**](PasswordProfile.md) |  | [optional] 
**Surname** | Pointer to **string** | The user&#39;s surname (family name or last name). Returned by default. | [optional] 
**GivenName** | Pointer to **string** | The user&#39;s givenName. Returned by default. | [optional] 
**UserType** | Pointer to **string** | The user&#x60;s type. This can be either \&quot;Member\&quot; for regular user, \&quot;Guest\&quot; for guest users or \&quot;Federated\&quot; for users imported from a federated instance. | [optional] [readonly] 
**PreferredLanguage** | Pointer to **string** | Represents the users language setting, ISO-639-1 Code | [optional] 
**SignInActivity** | Pointer to [**SignInActivity**](SignInActivity.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(displayName string, onPremisesSamAccountName string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountEnabled

`func (o *User) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *User) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *User) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *User) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### GetAppRoleAssignments

`func (o *User) GetAppRoleAssignments() []AppRoleAssignment`

GetAppRoleAssignments returns the AppRoleAssignments field if non-nil, zero value otherwise.

### GetAppRoleAssignmentsOk

`func (o *User) GetAppRoleAssignmentsOk() (*[]AppRoleAssignment, bool)`

GetAppRoleAssignmentsOk returns a tuple with the AppRoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleAssignments

`func (o *User) SetAppRoleAssignments(v []AppRoleAssignment)`

SetAppRoleAssignments sets AppRoleAssignments field to given value.

### HasAppRoleAssignments

`func (o *User) HasAppRoleAssignments() bool`

HasAppRoleAssignments returns a boolean if a field has been set.

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDrives

`func (o *User) GetDrives() []Drive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *User) GetDrivesOk() (*[]Drive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *User) SetDrives(v []Drive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *User) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetDrive

`func (o *User) GetDrive() Drive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *User) GetDriveOk() (*Drive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *User) SetDrive(v Drive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *User) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### GetIdentities

`func (o *User) GetIdentities() []ObjectIdentity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *User) GetIdentitiesOk() (*[]ObjectIdentity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *User) SetIdentities(v []ObjectIdentity)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *User) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetMail

`func (o *User) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *User) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *User) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *User) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetMemberOf

`func (o *User) GetMemberOf() []Group`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *User) GetMemberOfOk() (*[]Group, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *User) SetMemberOf(v []Group)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *User) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetOnPremisesSamAccountName

`func (o *User) GetOnPremisesSamAccountName() string`

GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field if non-nil, zero value otherwise.

### GetOnPremisesSamAccountNameOk

`func (o *User) GetOnPremisesSamAccountNameOk() (*string, bool)`

GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSamAccountName

`func (o *User) SetOnPremisesSamAccountName(v string)`

SetOnPremisesSamAccountName sets OnPremisesSamAccountName field to given value.


### GetPasswordProfile

`func (o *User) GetPasswordProfile() PasswordProfile`

GetPasswordProfile returns the PasswordProfile field if non-nil, zero value otherwise.

### GetPasswordProfileOk

`func (o *User) GetPasswordProfileOk() (*PasswordProfile, bool)`

GetPasswordProfileOk returns a tuple with the PasswordProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProfile

`func (o *User) SetPasswordProfile(v PasswordProfile)`

SetPasswordProfile sets PasswordProfile field to given value.

### HasPasswordProfile

`func (o *User) HasPasswordProfile() bool`

HasPasswordProfile returns a boolean if a field has been set.

### GetSurname

`func (o *User) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *User) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *User) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *User) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetGivenName

`func (o *User) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *User) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *User) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *User) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetUserType

`func (o *User) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *User) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *User) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *User) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetPreferredLanguage

`func (o *User) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *User) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *User) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *User) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### GetSignInActivity

`func (o *User) GetSignInActivity() SignInActivity`

GetSignInActivity returns the SignInActivity field if non-nil, zero value otherwise.

### GetSignInActivityOk

`func (o *User) GetSignInActivityOk() (*SignInActivity, bool)`

GetSignInActivityOk returns a tuple with the SignInActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInActivity

`func (o *User) SetSignInActivity(v SignInActivity)`

SetSignInActivity sets SignInActivity field to given value.

### HasSignInActivity

`func (o *User) HasSignInActivity() bool`

HasSignInActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


