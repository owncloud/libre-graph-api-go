# UserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Drive** | Pointer to [**Drive**](Drive.md) |  | [optional] 
**Drives** | Pointer to [**[]Drive**](Drive.md) | A collection of drives available for this user. Read-only. | [optional] [readonly] 

## Methods

### NewUserAllOf

`func NewUserAllOf() *UserAllOf`

NewUserAllOf instantiates a new UserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAllOfWithDefaults

`func NewUserAllOfWithDefaults() *UserAllOf`

NewUserAllOfWithDefaults instantiates a new UserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountEnabled

`func (o *UserAllOf) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *UserAllOf) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *UserAllOf) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *UserAllOf) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### GetBusinessPhones

`func (o *UserAllOf) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *UserAllOf) GetBusinessPhonesOk() (*[]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *UserAllOf) SetBusinessPhones(v []string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *UserAllOf) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetCity

`func (o *UserAllOf) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserAllOf) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserAllOf) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserAllOf) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *UserAllOf) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UserAllOf) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UserAllOf) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UserAllOf) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *UserAllOf) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserAllOf) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserAllOf) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserAllOf) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *UserAllOf) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *UserAllOf) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *UserAllOf) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *UserAllOf) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDepartment

`func (o *UserAllOf) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *UserAllOf) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *UserAllOf) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *UserAllOf) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserAllOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFaxNumber

`func (o *UserAllOf) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *UserAllOf) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *UserAllOf) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *UserAllOf) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### GetGivenName

`func (o *UserAllOf) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *UserAllOf) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *UserAllOf) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *UserAllOf) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetLastPasswordChangeDateTime

`func (o *UserAllOf) GetLastPasswordChangeDateTime() time.Time`

GetLastPasswordChangeDateTime returns the LastPasswordChangeDateTime field if non-nil, zero value otherwise.

### GetLastPasswordChangeDateTimeOk

`func (o *UserAllOf) GetLastPasswordChangeDateTimeOk() (*time.Time, bool)`

GetLastPasswordChangeDateTimeOk returns a tuple with the LastPasswordChangeDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChangeDateTime

`func (o *UserAllOf) SetLastPasswordChangeDateTime(v time.Time)`

SetLastPasswordChangeDateTime sets LastPasswordChangeDateTime field to given value.

### HasLastPasswordChangeDateTime

`func (o *UserAllOf) HasLastPasswordChangeDateTime() bool`

HasLastPasswordChangeDateTime returns a boolean if a field has been set.

### GetLegalAgeGroupClassification

`func (o *UserAllOf) GetLegalAgeGroupClassification() string`

GetLegalAgeGroupClassification returns the LegalAgeGroupClassification field if non-nil, zero value otherwise.

### GetLegalAgeGroupClassificationOk

`func (o *UserAllOf) GetLegalAgeGroupClassificationOk() (*string, bool)`

GetLegalAgeGroupClassificationOk returns a tuple with the LegalAgeGroupClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAgeGroupClassification

`func (o *UserAllOf) SetLegalAgeGroupClassification(v string)`

SetLegalAgeGroupClassification sets LegalAgeGroupClassification field to given value.

### HasLegalAgeGroupClassification

`func (o *UserAllOf) HasLegalAgeGroupClassification() bool`

HasLegalAgeGroupClassification returns a boolean if a field has been set.

### GetMail

`func (o *UserAllOf) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *UserAllOf) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *UserAllOf) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *UserAllOf) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetMailNickname

`func (o *UserAllOf) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *UserAllOf) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *UserAllOf) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *UserAllOf) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### GetMobilePhone

`func (o *UserAllOf) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *UserAllOf) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *UserAllOf) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *UserAllOf) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetOfficeLocation

`func (o *UserAllOf) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *UserAllOf) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *UserAllOf) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *UserAllOf) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### GetPostalCode

`func (o *UserAllOf) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UserAllOf) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UserAllOf) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UserAllOf) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPreferredLanguage

`func (o *UserAllOf) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *UserAllOf) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *UserAllOf) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *UserAllOf) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### GetState

`func (o *UserAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreetAddress

`func (o *UserAllOf) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UserAllOf) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UserAllOf) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *UserAllOf) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetSurname

`func (o *UserAllOf) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *UserAllOf) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *UserAllOf) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *UserAllOf) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetUsageLocation

`func (o *UserAllOf) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *UserAllOf) GetUsageLocationOk() (*string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLocation

`func (o *UserAllOf) SetUsageLocation(v string)`

SetUsageLocation sets UsageLocation field to given value.

### HasUsageLocation

`func (o *UserAllOf) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *UserAllOf) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *UserAllOf) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *UserAllOf) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *UserAllOf) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetUserType

`func (o *UserAllOf) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserAllOf) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserAllOf) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserAllOf) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetAboutMe

`func (o *UserAllOf) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *UserAllOf) GetAboutMeOk() (*string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutMe

`func (o *UserAllOf) SetAboutMe(v string)`

SetAboutMe sets AboutMe field to given value.

### HasAboutMe

`func (o *UserAllOf) HasAboutMe() bool`

HasAboutMe returns a boolean if a field has been set.

### GetBirthday

`func (o *UserAllOf) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UserAllOf) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UserAllOf) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UserAllOf) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetDrive

`func (o *UserAllOf) GetDrive() Drive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *UserAllOf) GetDriveOk() (*Drive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *UserAllOf) SetDrive(v Drive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *UserAllOf) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### GetDrives

`func (o *UserAllOf) GetDrives() []Drive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *UserAllOf) GetDrivesOk() (*[]Drive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *UserAllOf) SetDrives(v []Drive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *UserAllOf) HasDrives() bool`

HasDrives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


