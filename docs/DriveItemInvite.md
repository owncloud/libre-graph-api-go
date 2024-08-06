# DriveItemInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | Pointer to [**[]DriveRecipient**](DriveRecipient.md) | A collection of recipients who will receive access and the sharing invitation. Currently, only internal users or groups are supported. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the roles that are to be granted to the recipients of the sharing invitation. | [optional] 
**LibreGraphPermissionsActions** | Pointer to **[]string** | Specifies the actions that are to be granted to the recipients of the sharing invitation, in effect creating a custom role. | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | Specifies the dateTime after which the permission expires. | [optional] 

## Methods

### NewDriveItemInvite

`func NewDriveItemInvite() *DriveItemInvite`

NewDriveItemInvite instantiates a new DriveItemInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveItemInviteWithDefaults

`func NewDriveItemInviteWithDefaults() *DriveItemInvite`

NewDriveItemInviteWithDefaults instantiates a new DriveItemInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *DriveItemInvite) GetRecipients() []DriveRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *DriveItemInvite) GetRecipientsOk() (*[]DriveRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *DriveItemInvite) SetRecipients(v []DriveRecipient)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *DriveItemInvite) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetRoles

`func (o *DriveItemInvite) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DriveItemInvite) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DriveItemInvite) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DriveItemInvite) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetLibreGraphPermissionsActions

`func (o *DriveItemInvite) GetLibreGraphPermissionsActions() []string`

GetLibreGraphPermissionsActions returns the LibreGraphPermissionsActions field if non-nil, zero value otherwise.

### GetLibreGraphPermissionsActionsOk

`func (o *DriveItemInvite) GetLibreGraphPermissionsActionsOk() (*[]string, bool)`

GetLibreGraphPermissionsActionsOk returns a tuple with the LibreGraphPermissionsActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphPermissionsActions

`func (o *DriveItemInvite) SetLibreGraphPermissionsActions(v []string)`

SetLibreGraphPermissionsActions sets LibreGraphPermissionsActions field to given value.

### HasLibreGraphPermissionsActions

`func (o *DriveItemInvite) HasLibreGraphPermissionsActions() bool`

HasLibreGraphPermissionsActions returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *DriveItemInvite) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *DriveItemInvite) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *DriveItemInvite) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *DriveItemInvite) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


