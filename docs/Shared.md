# Shared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**Scope** | Pointer to **NullableString** | Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only. | [optional] 
**SharedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**SharedDateTime** | Pointer to **NullableTime** | The UTC date and time when the item was shared. Read-only. | [optional] 

## Methods

### NewShared

`func NewShared() *Shared`

NewShared instantiates a new Shared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedWithDefaults

`func NewSharedWithDefaults() *Shared`

NewSharedWithDefaults instantiates a new Shared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *Shared) GetOwner() IdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Shared) GetOwnerOk() (*IdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Shared) SetOwner(v IdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Shared) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetScope

`func (o *Shared) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Shared) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Shared) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Shared) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *Shared) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *Shared) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetSharedBy

`func (o *Shared) GetSharedBy() IdentitySet`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *Shared) GetSharedByOk() (*IdentitySet, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *Shared) SetSharedBy(v IdentitySet)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *Shared) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetSharedDateTime

`func (o *Shared) GetSharedDateTime() time.Time`

GetSharedDateTime returns the SharedDateTime field if non-nil, zero value otherwise.

### GetSharedDateTimeOk

`func (o *Shared) GetSharedDateTimeOk() (*time.Time, bool)`

GetSharedDateTimeOk returns a tuple with the SharedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDateTime

`func (o *Shared) SetSharedDateTime(v time.Time)`

SetSharedDateTime sets SharedDateTime field to given value.

### HasSharedDateTime

`func (o *Shared) HasSharedDateTime() bool`

HasSharedDateTime returns a boolean if a field has been set.

### SetSharedDateTimeNil

`func (o *Shared) SetSharedDateTimeNil(b bool)`

 SetSharedDateTimeNil sets the value for SharedDateTime to be an explicit nil

### UnsetSharedDateTime
`func (o *Shared) UnsetSharedDateTime()`

UnsetSharedDateTime ensures that no value is present for SharedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


