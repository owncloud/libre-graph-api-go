# EducationSchool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique idenfier for an entity. Read-only. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The organization name | [optional] 
**SchoolNumber** | Pointer to **string** | School number | [optional] 
**TerminationDate** | Pointer to **NullableTime** | Date and time at which the service for this organization is scheduled to be terminated | [optional] 

## Methods

### NewEducationSchool

`func NewEducationSchool() *EducationSchool`

NewEducationSchool instantiates a new EducationSchool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationSchoolWithDefaults

`func NewEducationSchoolWithDefaults() *EducationSchool`

NewEducationSchoolWithDefaults instantiates a new EducationSchool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EducationSchool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EducationSchool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EducationSchool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EducationSchool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *EducationSchool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EducationSchool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EducationSchool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EducationSchool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSchoolNumber

`func (o *EducationSchool) GetSchoolNumber() string`

GetSchoolNumber returns the SchoolNumber field if non-nil, zero value otherwise.

### GetSchoolNumberOk

`func (o *EducationSchool) GetSchoolNumberOk() (*string, bool)`

GetSchoolNumberOk returns a tuple with the SchoolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchoolNumber

`func (o *EducationSchool) SetSchoolNumber(v string)`

SetSchoolNumber sets SchoolNumber field to given value.

### HasSchoolNumber

`func (o *EducationSchool) HasSchoolNumber() bool`

HasSchoolNumber returns a boolean if a field has been set.

### GetTerminationDate

`func (o *EducationSchool) GetTerminationDate() time.Time`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *EducationSchool) GetTerminationDateOk() (*time.Time, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *EducationSchool) SetTerminationDate(v time.Time)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *EducationSchool) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *EducationSchool) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *EducationSchool) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


