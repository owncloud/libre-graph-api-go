# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] [readonly] 
**DeletedDateTime** | Pointer to **time.Time** |  | [optional] 
**AccountEnabled** | Pointer to **bool** | true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter. | [optional] 
**BusinessPhones** | Pointer to **[]string** | The telephone numbers for the user. Only one number can be set for this property. Returned by default. Read-only for users synced from on-premises directory. | [optional] 
**City** | Pointer to **string** | The city in which the user is located. Returned only on $select. Supports $filter. | [optional] 
**CompanyName** | Pointer to **string** | The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. | [optional] 
**Country** | Pointer to **string** | The country/region in which the user is located; for example, &#39;US&#39; or &#39;UK&#39;. Returned only on $select. Supports $filter. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The date and time the user was created. The value cannot be modified and is automatically populated when the entity is created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. Property is nullable. A null value indicates that an accurate creation time couldn&#39;t be determined for the user. Returned only on $select. Read-only. Supports $filter. | [optional] [readonly] 
**Department** | Pointer to **string** | The name for the department in which the user works. Returned only on $select. Supports $filter. | [optional] 
**DisplayName** | Pointer to **string** | The name displayed in the address book for the user. This value is usually the combination of the user&#39;s first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Returned by default. Supports $filter and $orderby. | [optional] 
**FaxNumber** | Pointer to **string** | The fax number of the user. Returned only on $select. | [optional] 
**GivenName** | Pointer to **string** | The given name (first name) of the user. Returned by default. Supports $filter. | [optional] 
**LastPasswordChangeDateTime** | Pointer to **time.Time** | The time when this user last changed their password. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select. Read-only. | [optional] [readonly] 
**LegalAgeGroupClassification** | Pointer to **string** | Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. | [optional] 
**Mail** | Pointer to **string** | The SMTP address for the user, for example, &#39;jeff@contoso.onowncloud.com&#39;. Returned by default. Supports $filter and endsWith. | [optional] 
**MailNickname** | Pointer to **string** | The mail alias for the user. This property must be specified when a user is created. Returned only on $select. Supports $filter. | [optional] 
**MobilePhone** | Pointer to **string** | The primary cellular telephone number for the user. Returned by default. Read-only for users synced from on-premises directory. | [optional] 
**OnPremisesDistinguishedName** | Pointer to **NullableString** | Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. | [optional] 
**OnPremisesDomainName** | Pointer to **NullableString** | Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. | [optional] 
**OnPremisesImmutableId** | Pointer to **NullableString** | This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user&#39;s userPrincipalName (UPN) property. NOTE: The $ and _ characters cannot be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).. | [optional] 
**OnPremisesSyncEnabled** | Pointer to **NullableBool** | true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to **NullableTime** | Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in). | [optional] 
**OnPremisesSamAccountName** | Pointer to **NullableString** | Contains the on-premises SAM account name synchronized from the on-premises directory. Read-only. | [optional] 
**OnPremisesUserPrincipalName** | Pointer to **NullableString** | Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith). | [optional] 
**OfficeLocation** | Pointer to **string** | The office location in the user&#39;s place of business. Returned by default. | [optional] 
**PostalCode** | Pointer to **string** | The postal code for the user&#39;s postal address. The postal code is specific to the user&#39;s country/region. In the United States of America, this attribute contains the ZIP code. Returned only on $select. | [optional] 
**PreferredLanguage** | Pointer to **string** | The preferred language for the user. Should follow ISO 639-1 Code; for example &#39;en-US&#39;. Returned by default. | [optional] 
**State** | Pointer to **string** | The state or province in the user&#39;s address. Returned only on $select. Supports $filter. | [optional] 
**StreetAddress** | Pointer to **string** | The street address of the user&#39;s place of business. Returned only on $select. | [optional] 
**Surname** | Pointer to **string** | The user&#39;s surname (family name or last name). Returned by default. Supports $filter. | [optional] 
**UsageLocation** | Pointer to **string** | A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: &#39;US&#39;, &#39;JP&#39;, and &#39;GB&#39;. Not nullable. Returned only on $select. Supports $filter. | [optional] 
**UserPrincipalName** | Pointer to **string** | The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user&#39;s email name. The general format is alias@domain, where domain must be present in the tenant&#39;s collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization. Returned by default. Supports $filter, $orderby, and endsWith. | [optional] 
**UserType** | Pointer to **string** | A string value that can be used to classify user types in your directory, such as &#39;Member&#39; and &#39;Guest&#39;. Returned only on $select. Supports $filter. | [optional] 
**AboutMe** | Pointer to **string** | A freeform text entry field for the user to describe themselves. Returned only on $select. | [optional] 
**Birthday** | Pointer to **time.Time** | The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select. | [optional] 
**PreferredName** | Pointer to **NullableString** | The preferred name for the user. Returned only on $select. | [optional] 
**Drive** | Pointer to [**Drive**](Drive.md) |  | [optional] 
**Drives** | Pointer to [**[]Drive**](Drive.md) | A collection of drives available for this user. Read-only. | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

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

### GetDeletedDateTime

`func (o *User) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *User) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *User) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *User) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

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

### GetBusinessPhones

`func (o *User) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *User) GetBusinessPhonesOk() (*[]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *User) SetBusinessPhones(v []string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *User) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetCity

`func (o *User) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *User) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *User) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *User) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *User) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *User) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *User) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *User) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *User) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *User) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *User) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *User) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *User) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *User) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *User) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *User) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDepartment

`func (o *User) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *User) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *User) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *User) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

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

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFaxNumber

`func (o *User) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *User) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *User) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *User) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

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

### GetLastPasswordChangeDateTime

`func (o *User) GetLastPasswordChangeDateTime() time.Time`

GetLastPasswordChangeDateTime returns the LastPasswordChangeDateTime field if non-nil, zero value otherwise.

### GetLastPasswordChangeDateTimeOk

`func (o *User) GetLastPasswordChangeDateTimeOk() (*time.Time, bool)`

GetLastPasswordChangeDateTimeOk returns a tuple with the LastPasswordChangeDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChangeDateTime

`func (o *User) SetLastPasswordChangeDateTime(v time.Time)`

SetLastPasswordChangeDateTime sets LastPasswordChangeDateTime field to given value.

### HasLastPasswordChangeDateTime

`func (o *User) HasLastPasswordChangeDateTime() bool`

HasLastPasswordChangeDateTime returns a boolean if a field has been set.

### GetLegalAgeGroupClassification

`func (o *User) GetLegalAgeGroupClassification() string`

GetLegalAgeGroupClassification returns the LegalAgeGroupClassification field if non-nil, zero value otherwise.

### GetLegalAgeGroupClassificationOk

`func (o *User) GetLegalAgeGroupClassificationOk() (*string, bool)`

GetLegalAgeGroupClassificationOk returns a tuple with the LegalAgeGroupClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAgeGroupClassification

`func (o *User) SetLegalAgeGroupClassification(v string)`

SetLegalAgeGroupClassification sets LegalAgeGroupClassification field to given value.

### HasLegalAgeGroupClassification

`func (o *User) HasLegalAgeGroupClassification() bool`

HasLegalAgeGroupClassification returns a boolean if a field has been set.

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

### GetMailNickname

`func (o *User) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *User) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *User) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *User) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### GetMobilePhone

`func (o *User) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *User) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *User) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *User) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetOnPremisesDistinguishedName

`func (o *User) GetOnPremisesDistinguishedName() string`

GetOnPremisesDistinguishedName returns the OnPremisesDistinguishedName field if non-nil, zero value otherwise.

### GetOnPremisesDistinguishedNameOk

`func (o *User) GetOnPremisesDistinguishedNameOk() (*string, bool)`

GetOnPremisesDistinguishedNameOk returns a tuple with the OnPremisesDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesDistinguishedName

`func (o *User) SetOnPremisesDistinguishedName(v string)`

SetOnPremisesDistinguishedName sets OnPremisesDistinguishedName field to given value.

### HasOnPremisesDistinguishedName

`func (o *User) HasOnPremisesDistinguishedName() bool`

HasOnPremisesDistinguishedName returns a boolean if a field has been set.

### SetOnPremisesDistinguishedNameNil

`func (o *User) SetOnPremisesDistinguishedNameNil(b bool)`

 SetOnPremisesDistinguishedNameNil sets the value for OnPremisesDistinguishedName to be an explicit nil

### UnsetOnPremisesDistinguishedName
`func (o *User) UnsetOnPremisesDistinguishedName()`

UnsetOnPremisesDistinguishedName ensures that no value is present for OnPremisesDistinguishedName, not even an explicit nil
### GetOnPremisesDomainName

`func (o *User) GetOnPremisesDomainName() string`

GetOnPremisesDomainName returns the OnPremisesDomainName field if non-nil, zero value otherwise.

### GetOnPremisesDomainNameOk

`func (o *User) GetOnPremisesDomainNameOk() (*string, bool)`

GetOnPremisesDomainNameOk returns a tuple with the OnPremisesDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesDomainName

`func (o *User) SetOnPremisesDomainName(v string)`

SetOnPremisesDomainName sets OnPremisesDomainName field to given value.

### HasOnPremisesDomainName

`func (o *User) HasOnPremisesDomainName() bool`

HasOnPremisesDomainName returns a boolean if a field has been set.

### SetOnPremisesDomainNameNil

`func (o *User) SetOnPremisesDomainNameNil(b bool)`

 SetOnPremisesDomainNameNil sets the value for OnPremisesDomainName to be an explicit nil

### UnsetOnPremisesDomainName
`func (o *User) UnsetOnPremisesDomainName()`

UnsetOnPremisesDomainName ensures that no value is present for OnPremisesDomainName, not even an explicit nil
### GetOnPremisesImmutableId

`func (o *User) GetOnPremisesImmutableId() string`

GetOnPremisesImmutableId returns the OnPremisesImmutableId field if non-nil, zero value otherwise.

### GetOnPremisesImmutableIdOk

`func (o *User) GetOnPremisesImmutableIdOk() (*string, bool)`

GetOnPremisesImmutableIdOk returns a tuple with the OnPremisesImmutableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesImmutableId

`func (o *User) SetOnPremisesImmutableId(v string)`

SetOnPremisesImmutableId sets OnPremisesImmutableId field to given value.

### HasOnPremisesImmutableId

`func (o *User) HasOnPremisesImmutableId() bool`

HasOnPremisesImmutableId returns a boolean if a field has been set.

### SetOnPremisesImmutableIdNil

`func (o *User) SetOnPremisesImmutableIdNil(b bool)`

 SetOnPremisesImmutableIdNil sets the value for OnPremisesImmutableId to be an explicit nil

### UnsetOnPremisesImmutableId
`func (o *User) UnsetOnPremisesImmutableId()`

UnsetOnPremisesImmutableId ensures that no value is present for OnPremisesImmutableId, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *User) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *User) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *User) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *User) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *User) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *User) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetOnPremisesLastSyncDateTime

`func (o *User) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *User) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *User) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *User) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *User) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *User) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
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

### HasOnPremisesSamAccountName

`func (o *User) HasOnPremisesSamAccountName() bool`

HasOnPremisesSamAccountName returns a boolean if a field has been set.

### SetOnPremisesSamAccountNameNil

`func (o *User) SetOnPremisesSamAccountNameNil(b bool)`

 SetOnPremisesSamAccountNameNil sets the value for OnPremisesSamAccountName to be an explicit nil

### UnsetOnPremisesSamAccountName
`func (o *User) UnsetOnPremisesSamAccountName()`

UnsetOnPremisesSamAccountName ensures that no value is present for OnPremisesSamAccountName, not even an explicit nil
### GetOnPremisesUserPrincipalName

`func (o *User) GetOnPremisesUserPrincipalName() string`

GetOnPremisesUserPrincipalName returns the OnPremisesUserPrincipalName field if non-nil, zero value otherwise.

### GetOnPremisesUserPrincipalNameOk

`func (o *User) GetOnPremisesUserPrincipalNameOk() (*string, bool)`

GetOnPremisesUserPrincipalNameOk returns a tuple with the OnPremisesUserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesUserPrincipalName

`func (o *User) SetOnPremisesUserPrincipalName(v string)`

SetOnPremisesUserPrincipalName sets OnPremisesUserPrincipalName field to given value.

### HasOnPremisesUserPrincipalName

`func (o *User) HasOnPremisesUserPrincipalName() bool`

HasOnPremisesUserPrincipalName returns a boolean if a field has been set.

### SetOnPremisesUserPrincipalNameNil

`func (o *User) SetOnPremisesUserPrincipalNameNil(b bool)`

 SetOnPremisesUserPrincipalNameNil sets the value for OnPremisesUserPrincipalName to be an explicit nil

### UnsetOnPremisesUserPrincipalName
`func (o *User) UnsetOnPremisesUserPrincipalName()`

UnsetOnPremisesUserPrincipalName ensures that no value is present for OnPremisesUserPrincipalName, not even an explicit nil
### GetOfficeLocation

`func (o *User) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *User) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *User) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *User) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### GetPostalCode

`func (o *User) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *User) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *User) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *User) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

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

### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreetAddress

`func (o *User) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *User) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *User) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *User) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

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

### GetUsageLocation

`func (o *User) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *User) GetUsageLocationOk() (*string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLocation

`func (o *User) SetUsageLocation(v string)`

SetUsageLocation sets UsageLocation field to given value.

### HasUsageLocation

`func (o *User) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *User) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *User) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *User) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *User) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

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

### GetAboutMe

`func (o *User) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *User) GetAboutMeOk() (*string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutMe

`func (o *User) SetAboutMe(v string)`

SetAboutMe sets AboutMe field to given value.

### HasAboutMe

`func (o *User) HasAboutMe() bool`

HasAboutMe returns a boolean if a field has been set.

### GetBirthday

`func (o *User) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *User) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *User) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *User) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetPreferredName

`func (o *User) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *User) GetPreferredNameOk() (*string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredName

`func (o *User) SetPreferredName(v string)`

SetPreferredName sets PreferredName field to given value.

### HasPreferredName

`func (o *User) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### SetPreferredNameNil

`func (o *User) SetPreferredNameNil(b bool)`

 SetPreferredNameNil sets the value for PreferredName to be an explicit nil

### UnsetPreferredName
`func (o *User) UnsetPreferredName()`

UnsetPreferredName ensures that no value is present for PreferredName, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


