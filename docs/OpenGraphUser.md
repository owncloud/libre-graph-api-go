# OpenGraphUser

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
**Drive** | Pointer to [**OpenGraphDrive**](open.graph.drive.md) | The user&#39;s HomeDrive. Read-only. | [optional] [readonly] 
**Drives** | Pointer to [**[]OpenGraphDrive**](OpenGraphDrive.md) | A collection of drives available for this user. Read-only. | [optional] [readonly] 

## Methods

### NewOpenGraphUser

`func NewOpenGraphUser() *OpenGraphUser`

NewOpenGraphUser instantiates a new OpenGraphUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphUserWithDefaults

`func NewOpenGraphUserWithDefaults() *OpenGraphUser`

NewOpenGraphUserWithDefaults instantiates a new OpenGraphUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountEnabled

`func (o *OpenGraphUser) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *OpenGraphUser) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *OpenGraphUser) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *OpenGraphUser) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### GetBusinessPhones

`func (o *OpenGraphUser) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *OpenGraphUser) GetBusinessPhonesOk() (*[]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *OpenGraphUser) SetBusinessPhones(v []string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *OpenGraphUser) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetCity

`func (o *OpenGraphUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OpenGraphUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OpenGraphUser) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OpenGraphUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *OpenGraphUser) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OpenGraphUser) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OpenGraphUser) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OpenGraphUser) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *OpenGraphUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OpenGraphUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OpenGraphUser) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OpenGraphUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *OpenGraphUser) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OpenGraphUser) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *OpenGraphUser) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *OpenGraphUser) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDepartment

`func (o *OpenGraphUser) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *OpenGraphUser) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *OpenGraphUser) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *OpenGraphUser) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDisplayName

`func (o *OpenGraphUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OpenGraphUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OpenGraphUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OpenGraphUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFaxNumber

`func (o *OpenGraphUser) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *OpenGraphUser) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *OpenGraphUser) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *OpenGraphUser) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### GetGivenName

`func (o *OpenGraphUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *OpenGraphUser) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *OpenGraphUser) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *OpenGraphUser) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetLastPasswordChangeDateTime

`func (o *OpenGraphUser) GetLastPasswordChangeDateTime() time.Time`

GetLastPasswordChangeDateTime returns the LastPasswordChangeDateTime field if non-nil, zero value otherwise.

### GetLastPasswordChangeDateTimeOk

`func (o *OpenGraphUser) GetLastPasswordChangeDateTimeOk() (*time.Time, bool)`

GetLastPasswordChangeDateTimeOk returns a tuple with the LastPasswordChangeDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChangeDateTime

`func (o *OpenGraphUser) SetLastPasswordChangeDateTime(v time.Time)`

SetLastPasswordChangeDateTime sets LastPasswordChangeDateTime field to given value.

### HasLastPasswordChangeDateTime

`func (o *OpenGraphUser) HasLastPasswordChangeDateTime() bool`

HasLastPasswordChangeDateTime returns a boolean if a field has been set.

### GetLegalAgeGroupClassification

`func (o *OpenGraphUser) GetLegalAgeGroupClassification() string`

GetLegalAgeGroupClassification returns the LegalAgeGroupClassification field if non-nil, zero value otherwise.

### GetLegalAgeGroupClassificationOk

`func (o *OpenGraphUser) GetLegalAgeGroupClassificationOk() (*string, bool)`

GetLegalAgeGroupClassificationOk returns a tuple with the LegalAgeGroupClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAgeGroupClassification

`func (o *OpenGraphUser) SetLegalAgeGroupClassification(v string)`

SetLegalAgeGroupClassification sets LegalAgeGroupClassification field to given value.

### HasLegalAgeGroupClassification

`func (o *OpenGraphUser) HasLegalAgeGroupClassification() bool`

HasLegalAgeGroupClassification returns a boolean if a field has been set.

### GetMail

`func (o *OpenGraphUser) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *OpenGraphUser) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *OpenGraphUser) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *OpenGraphUser) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetMailNickname

`func (o *OpenGraphUser) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *OpenGraphUser) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *OpenGraphUser) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *OpenGraphUser) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### GetMobilePhone

`func (o *OpenGraphUser) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *OpenGraphUser) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *OpenGraphUser) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *OpenGraphUser) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetOfficeLocation

`func (o *OpenGraphUser) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *OpenGraphUser) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *OpenGraphUser) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *OpenGraphUser) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### GetPostalCode

`func (o *OpenGraphUser) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OpenGraphUser) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OpenGraphUser) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OpenGraphUser) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPreferredLanguage

`func (o *OpenGraphUser) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *OpenGraphUser) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *OpenGraphUser) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *OpenGraphUser) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### GetState

`func (o *OpenGraphUser) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OpenGraphUser) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OpenGraphUser) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OpenGraphUser) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreetAddress

`func (o *OpenGraphUser) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *OpenGraphUser) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *OpenGraphUser) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *OpenGraphUser) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetSurname

`func (o *OpenGraphUser) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *OpenGraphUser) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *OpenGraphUser) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *OpenGraphUser) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetUsageLocation

`func (o *OpenGraphUser) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *OpenGraphUser) GetUsageLocationOk() (*string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLocation

`func (o *OpenGraphUser) SetUsageLocation(v string)`

SetUsageLocation sets UsageLocation field to given value.

### HasUsageLocation

`func (o *OpenGraphUser) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *OpenGraphUser) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *OpenGraphUser) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *OpenGraphUser) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *OpenGraphUser) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetUserType

`func (o *OpenGraphUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *OpenGraphUser) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *OpenGraphUser) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *OpenGraphUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetAboutMe

`func (o *OpenGraphUser) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *OpenGraphUser) GetAboutMeOk() (*string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutMe

`func (o *OpenGraphUser) SetAboutMe(v string)`

SetAboutMe sets AboutMe field to given value.

### HasAboutMe

`func (o *OpenGraphUser) HasAboutMe() bool`

HasAboutMe returns a boolean if a field has been set.

### GetBirthday

`func (o *OpenGraphUser) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *OpenGraphUser) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *OpenGraphUser) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *OpenGraphUser) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetDrive

`func (o *OpenGraphUser) GetDrive() OpenGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *OpenGraphUser) GetDriveOk() (*OpenGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *OpenGraphUser) SetDrive(v OpenGraphDrive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *OpenGraphUser) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### GetDrives

`func (o *OpenGraphUser) GetDrives() []OpenGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *OpenGraphUser) GetDrivesOk() (*[]OpenGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *OpenGraphUser) SetDrives(v []OpenGraphDrive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *OpenGraphUser) HasDrives() bool`

HasDrives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


