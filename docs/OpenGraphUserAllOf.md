# OpenGraphUserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountEnabled** | Pointer to **NullableBool** | true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter. | [optional] 
**BusinessPhones** | Pointer to **[]string** | The telephone numbers for the user. Only one number can be set for this property. Returned by default. Read-only for users synced from on-premises directory. | [optional] 
**City** | Pointer to **NullableString** | The city in which the user is located. Returned only on $select. Supports $filter. | [optional] 
**CompanyName** | Pointer to **NullableString** | The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. | [optional] 
**Country** | Pointer to **NullableString** | The country/region in which the user is located; for example, &#39;US&#39; or &#39;UK&#39;. Returned only on $select. Supports $filter. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time the user was created. The value cannot be modified and is automatically populated when the entity is created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. Property is nullable. A null value indicates that an accurate creation time couldn&#39;t be determined for the user. Returned only on $select. Read-only. Supports $filter. | [optional] [readonly] 
**Department** | Pointer to **NullableString** | The name for the department in which the user works. Returned only on $select. Supports $filter. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name displayed in the address book for the user. This value is usually the combination of the user&#39;s first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Returned by default. Supports $filter and $orderby. | [optional] 
**FaxNumber** | Pointer to **NullableString** | The fax number of the user. Returned only on $select. | [optional] 
**GivenName** | Pointer to **NullableString** | The given name (first name) of the user. Returned by default. Supports $filter. | [optional] 
**LastPasswordChangeDateTime** | Pointer to **NullableTime** | The time when this user last changed their password. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select. Read-only. | [optional] [readonly] 
**LegalAgeGroupClassification** | Pointer to **NullableString** | Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. | [optional] 
**Mail** | Pointer to **NullableString** | The SMTP address for the user, for example, &#39;jeff@contoso.onowncloud.com&#39;. Returned by default. Supports $filter and endsWith. | [optional] 
**MailNickname** | Pointer to **NullableString** | The mail alias for the user. This property must be specified when a user is created. Returned only on $select. Supports $filter. | [optional] 
**MobilePhone** | Pointer to **NullableString** | The primary cellular telephone number for the user. Returned by default. Read-only for users synced from on-premises directory. | [optional] 
**OfficeLocation** | Pointer to **NullableString** | The office location in the user&#39;s place of business. Returned by default. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code for the user&#39;s postal address. The postal code is specific to the user&#39;s country/region. In the United States of America, this attribute contains the ZIP code. Returned only on $select. | [optional] 
**PreferredLanguage** | Pointer to **NullableString** | The preferred language for the user. Should follow ISO 639-1 Code; for example &#39;en-US&#39;. Returned by default. | [optional] 
**State** | Pointer to **NullableString** | The state or province in the user&#39;s address. Returned only on $select. Supports $filter. | [optional] 
**StreetAddress** | Pointer to **NullableString** | The street address of the user&#39;s place of business. Returned only on $select. | [optional] 
**Surname** | Pointer to **NullableString** | The user&#39;s surname (family name or last name). Returned by default. Supports $filter. | [optional] 
**UsageLocation** | Pointer to **NullableString** | A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: &#39;US&#39;, &#39;JP&#39;, and &#39;GB&#39;. Not nullable. Returned only on $select. Supports $filter. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user&#39;s email name. The general format is alias@domain, where domain must be present in the tenant&#39;s collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization. Returned by default. Supports $filter, $orderby, and endsWith. | [optional] 
**UserType** | Pointer to **NullableString** | A string value that can be used to classify user types in your directory, such as &#39;Member&#39; and &#39;Guest&#39;. Returned only on $select. Supports $filter. | [optional] 
**AboutMe** | Pointer to **NullableString** | A freeform text entry field for the user to describe themselves. Returned only on $select. | [optional] 
**Birthday** | Pointer to **time.Time** | The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select. | [optional] 
**Drive** | Pointer to [**NullableAnyOfopenGraphDrive**](anyOf&lt;open.graph.drive&gt;.md) | The user&#39;s HomeDrive. Read-only. | [optional] [readonly] 
**Drives** | Pointer to [**[]OpenGraphDrive**](OpenGraphDrive.md) | A collection of drives available for this user. Read-only. | [optional] [readonly] 

## Methods

### NewOpenGraphUserAllOf

`func NewOpenGraphUserAllOf() *OpenGraphUserAllOf`

NewOpenGraphUserAllOf instantiates a new OpenGraphUserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphUserAllOfWithDefaults

`func NewOpenGraphUserAllOfWithDefaults() *OpenGraphUserAllOf`

NewOpenGraphUserAllOfWithDefaults instantiates a new OpenGraphUserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountEnabled

`func (o *OpenGraphUserAllOf) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *OpenGraphUserAllOf) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *OpenGraphUserAllOf) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *OpenGraphUserAllOf) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabledNil

`func (o *OpenGraphUserAllOf) SetAccountEnabledNil(b bool)`

 SetAccountEnabledNil sets the value for AccountEnabled to be an explicit nil

### UnsetAccountEnabled
`func (o *OpenGraphUserAllOf) UnsetAccountEnabled()`

UnsetAccountEnabled ensures that no value is present for AccountEnabled, not even an explicit nil
### GetBusinessPhones

`func (o *OpenGraphUserAllOf) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *OpenGraphUserAllOf) GetBusinessPhonesOk() (*[]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *OpenGraphUserAllOf) SetBusinessPhones(v []string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *OpenGraphUserAllOf) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetCity

`func (o *OpenGraphUserAllOf) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OpenGraphUserAllOf) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OpenGraphUserAllOf) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OpenGraphUserAllOf) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *OpenGraphUserAllOf) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *OpenGraphUserAllOf) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCompanyName

`func (o *OpenGraphUserAllOf) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OpenGraphUserAllOf) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OpenGraphUserAllOf) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OpenGraphUserAllOf) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *OpenGraphUserAllOf) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *OpenGraphUserAllOf) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCountry

`func (o *OpenGraphUserAllOf) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OpenGraphUserAllOf) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OpenGraphUserAllOf) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OpenGraphUserAllOf) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *OpenGraphUserAllOf) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *OpenGraphUserAllOf) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCreatedDateTime

`func (o *OpenGraphUserAllOf) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OpenGraphUserAllOf) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *OpenGraphUserAllOf) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *OpenGraphUserAllOf) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *OpenGraphUserAllOf) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *OpenGraphUserAllOf) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDepartment

`func (o *OpenGraphUserAllOf) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *OpenGraphUserAllOf) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *OpenGraphUserAllOf) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *OpenGraphUserAllOf) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *OpenGraphUserAllOf) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *OpenGraphUserAllOf) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDisplayName

`func (o *OpenGraphUserAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OpenGraphUserAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OpenGraphUserAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OpenGraphUserAllOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *OpenGraphUserAllOf) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *OpenGraphUserAllOf) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFaxNumber

`func (o *OpenGraphUserAllOf) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *OpenGraphUserAllOf) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *OpenGraphUserAllOf) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *OpenGraphUserAllOf) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### SetFaxNumberNil

`func (o *OpenGraphUserAllOf) SetFaxNumberNil(b bool)`

 SetFaxNumberNil sets the value for FaxNumber to be an explicit nil

### UnsetFaxNumber
`func (o *OpenGraphUserAllOf) UnsetFaxNumber()`

UnsetFaxNumber ensures that no value is present for FaxNumber, not even an explicit nil
### GetGivenName

`func (o *OpenGraphUserAllOf) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *OpenGraphUserAllOf) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *OpenGraphUserAllOf) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *OpenGraphUserAllOf) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *OpenGraphUserAllOf) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *OpenGraphUserAllOf) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetLastPasswordChangeDateTime

`func (o *OpenGraphUserAllOf) GetLastPasswordChangeDateTime() time.Time`

GetLastPasswordChangeDateTime returns the LastPasswordChangeDateTime field if non-nil, zero value otherwise.

### GetLastPasswordChangeDateTimeOk

`func (o *OpenGraphUserAllOf) GetLastPasswordChangeDateTimeOk() (*time.Time, bool)`

GetLastPasswordChangeDateTimeOk returns a tuple with the LastPasswordChangeDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChangeDateTime

`func (o *OpenGraphUserAllOf) SetLastPasswordChangeDateTime(v time.Time)`

SetLastPasswordChangeDateTime sets LastPasswordChangeDateTime field to given value.

### HasLastPasswordChangeDateTime

`func (o *OpenGraphUserAllOf) HasLastPasswordChangeDateTime() bool`

HasLastPasswordChangeDateTime returns a boolean if a field has been set.

### SetLastPasswordChangeDateTimeNil

`func (o *OpenGraphUserAllOf) SetLastPasswordChangeDateTimeNil(b bool)`

 SetLastPasswordChangeDateTimeNil sets the value for LastPasswordChangeDateTime to be an explicit nil

### UnsetLastPasswordChangeDateTime
`func (o *OpenGraphUserAllOf) UnsetLastPasswordChangeDateTime()`

UnsetLastPasswordChangeDateTime ensures that no value is present for LastPasswordChangeDateTime, not even an explicit nil
### GetLegalAgeGroupClassification

`func (o *OpenGraphUserAllOf) GetLegalAgeGroupClassification() string`

GetLegalAgeGroupClassification returns the LegalAgeGroupClassification field if non-nil, zero value otherwise.

### GetLegalAgeGroupClassificationOk

`func (o *OpenGraphUserAllOf) GetLegalAgeGroupClassificationOk() (*string, bool)`

GetLegalAgeGroupClassificationOk returns a tuple with the LegalAgeGroupClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAgeGroupClassification

`func (o *OpenGraphUserAllOf) SetLegalAgeGroupClassification(v string)`

SetLegalAgeGroupClassification sets LegalAgeGroupClassification field to given value.

### HasLegalAgeGroupClassification

`func (o *OpenGraphUserAllOf) HasLegalAgeGroupClassification() bool`

HasLegalAgeGroupClassification returns a boolean if a field has been set.

### SetLegalAgeGroupClassificationNil

`func (o *OpenGraphUserAllOf) SetLegalAgeGroupClassificationNil(b bool)`

 SetLegalAgeGroupClassificationNil sets the value for LegalAgeGroupClassification to be an explicit nil

### UnsetLegalAgeGroupClassification
`func (o *OpenGraphUserAllOf) UnsetLegalAgeGroupClassification()`

UnsetLegalAgeGroupClassification ensures that no value is present for LegalAgeGroupClassification, not even an explicit nil
### GetMail

`func (o *OpenGraphUserAllOf) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *OpenGraphUserAllOf) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *OpenGraphUserAllOf) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *OpenGraphUserAllOf) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMailNil

`func (o *OpenGraphUserAllOf) SetMailNil(b bool)`

 SetMailNil sets the value for Mail to be an explicit nil

### UnsetMail
`func (o *OpenGraphUserAllOf) UnsetMail()`

UnsetMail ensures that no value is present for Mail, not even an explicit nil
### GetMailNickname

`func (o *OpenGraphUserAllOf) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *OpenGraphUserAllOf) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *OpenGraphUserAllOf) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *OpenGraphUserAllOf) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNicknameNil

`func (o *OpenGraphUserAllOf) SetMailNicknameNil(b bool)`

 SetMailNicknameNil sets the value for MailNickname to be an explicit nil

### UnsetMailNickname
`func (o *OpenGraphUserAllOf) UnsetMailNickname()`

UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
### GetMobilePhone

`func (o *OpenGraphUserAllOf) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *OpenGraphUserAllOf) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *OpenGraphUserAllOf) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *OpenGraphUserAllOf) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *OpenGraphUserAllOf) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *OpenGraphUserAllOf) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetOfficeLocation

`func (o *OpenGraphUserAllOf) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *OpenGraphUserAllOf) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *OpenGraphUserAllOf) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *OpenGraphUserAllOf) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocationNil

`func (o *OpenGraphUserAllOf) SetOfficeLocationNil(b bool)`

 SetOfficeLocationNil sets the value for OfficeLocation to be an explicit nil

### UnsetOfficeLocation
`func (o *OpenGraphUserAllOf) UnsetOfficeLocation()`

UnsetOfficeLocation ensures that no value is present for OfficeLocation, not even an explicit nil
### GetPostalCode

`func (o *OpenGraphUserAllOf) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OpenGraphUserAllOf) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OpenGraphUserAllOf) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OpenGraphUserAllOf) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *OpenGraphUserAllOf) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *OpenGraphUserAllOf) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPreferredLanguage

`func (o *OpenGraphUserAllOf) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *OpenGraphUserAllOf) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *OpenGraphUserAllOf) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *OpenGraphUserAllOf) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *OpenGraphUserAllOf) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *OpenGraphUserAllOf) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetState

`func (o *OpenGraphUserAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OpenGraphUserAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OpenGraphUserAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OpenGraphUserAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *OpenGraphUserAllOf) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *OpenGraphUserAllOf) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreetAddress

`func (o *OpenGraphUserAllOf) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *OpenGraphUserAllOf) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *OpenGraphUserAllOf) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *OpenGraphUserAllOf) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *OpenGraphUserAllOf) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *OpenGraphUserAllOf) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetSurname

`func (o *OpenGraphUserAllOf) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *OpenGraphUserAllOf) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *OpenGraphUserAllOf) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *OpenGraphUserAllOf) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *OpenGraphUserAllOf) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *OpenGraphUserAllOf) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetUsageLocation

`func (o *OpenGraphUserAllOf) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *OpenGraphUserAllOf) GetUsageLocationOk() (*string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLocation

`func (o *OpenGraphUserAllOf) SetUsageLocation(v string)`

SetUsageLocation sets UsageLocation field to given value.

### HasUsageLocation

`func (o *OpenGraphUserAllOf) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### SetUsageLocationNil

`func (o *OpenGraphUserAllOf) SetUsageLocationNil(b bool)`

 SetUsageLocationNil sets the value for UsageLocation to be an explicit nil

### UnsetUsageLocation
`func (o *OpenGraphUserAllOf) UnsetUsageLocation()`

UnsetUsageLocation ensures that no value is present for UsageLocation, not even an explicit nil
### GetUserPrincipalName

`func (o *OpenGraphUserAllOf) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *OpenGraphUserAllOf) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *OpenGraphUserAllOf) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *OpenGraphUserAllOf) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *OpenGraphUserAllOf) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *OpenGraphUserAllOf) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
### GetUserType

`func (o *OpenGraphUserAllOf) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *OpenGraphUserAllOf) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *OpenGraphUserAllOf) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *OpenGraphUserAllOf) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *OpenGraphUserAllOf) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *OpenGraphUserAllOf) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetAboutMe

`func (o *OpenGraphUserAllOf) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *OpenGraphUserAllOf) GetAboutMeOk() (*string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutMe

`func (o *OpenGraphUserAllOf) SetAboutMe(v string)`

SetAboutMe sets AboutMe field to given value.

### HasAboutMe

`func (o *OpenGraphUserAllOf) HasAboutMe() bool`

HasAboutMe returns a boolean if a field has been set.

### SetAboutMeNil

`func (o *OpenGraphUserAllOf) SetAboutMeNil(b bool)`

 SetAboutMeNil sets the value for AboutMe to be an explicit nil

### UnsetAboutMe
`func (o *OpenGraphUserAllOf) UnsetAboutMe()`

UnsetAboutMe ensures that no value is present for AboutMe, not even an explicit nil
### GetBirthday

`func (o *OpenGraphUserAllOf) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *OpenGraphUserAllOf) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *OpenGraphUserAllOf) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *OpenGraphUserAllOf) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetDrive

`func (o *OpenGraphUserAllOf) GetDrive() AnyOfopenGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *OpenGraphUserAllOf) GetDriveOk() (*AnyOfopenGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *OpenGraphUserAllOf) SetDrive(v AnyOfopenGraphDrive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *OpenGraphUserAllOf) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDriveNil

`func (o *OpenGraphUserAllOf) SetDriveNil(b bool)`

 SetDriveNil sets the value for Drive to be an explicit nil

### UnsetDrive
`func (o *OpenGraphUserAllOf) UnsetDrive()`

UnsetDrive ensures that no value is present for Drive, not even an explicit nil
### GetDrives

`func (o *OpenGraphUserAllOf) GetDrives() []OpenGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *OpenGraphUserAllOf) GetDrivesOk() (*[]OpenGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *OpenGraphUserAllOf) SetDrives(v []OpenGraphDrive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *OpenGraphUserAllOf) HasDrives() bool`

HasDrives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


