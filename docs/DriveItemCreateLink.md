# DriveItemCreateLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SharingLinkType**](SharingLinkType.md) |  | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | Optional. A String with format of yyyy-MM-ddTHH:mm:ssZ of DateTime indicates the expiration time of the permission. | [optional] 
**Password** | Pointer to **string** | Optional.The password of the sharing link that is set by the creator. | [optional] 
**DisplayName** | Pointer to **string** | Provides a user-visible display name of the link. Optional. Libregraph only. | [optional] 
**LibreGraphQuickLink** | Pointer to **bool** | The quicklink property can be assigned to only one link per resource. A quicklink can be used in the clients to provide a one-click copy to clipboard action. Optional. Libregraph only. | [optional] 

## Methods

### NewDriveItemCreateLink

`func NewDriveItemCreateLink() *DriveItemCreateLink`

NewDriveItemCreateLink instantiates a new DriveItemCreateLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveItemCreateLinkWithDefaults

`func NewDriveItemCreateLinkWithDefaults() *DriveItemCreateLink`

NewDriveItemCreateLinkWithDefaults instantiates a new DriveItemCreateLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DriveItemCreateLink) GetType() SharingLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DriveItemCreateLink) GetTypeOk() (*SharingLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DriveItemCreateLink) SetType(v SharingLinkType)`

SetType sets Type field to given value.

### HasType

`func (o *DriveItemCreateLink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *DriveItemCreateLink) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *DriveItemCreateLink) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *DriveItemCreateLink) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *DriveItemCreateLink) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetPassword

`func (o *DriveItemCreateLink) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DriveItemCreateLink) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DriveItemCreateLink) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DriveItemCreateLink) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDisplayName

`func (o *DriveItemCreateLink) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DriveItemCreateLink) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DriveItemCreateLink) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DriveItemCreateLink) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLibreGraphQuickLink

`func (o *DriveItemCreateLink) GetLibreGraphQuickLink() bool`

GetLibreGraphQuickLink returns the LibreGraphQuickLink field if non-nil, zero value otherwise.

### GetLibreGraphQuickLinkOk

`func (o *DriveItemCreateLink) GetLibreGraphQuickLinkOk() (*bool, bool)`

GetLibreGraphQuickLinkOk returns a tuple with the LibreGraphQuickLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphQuickLink

`func (o *DriveItemCreateLink) SetLibreGraphQuickLink(v bool)`

SetLibreGraphQuickLink sets LibreGraphQuickLink field to given value.

### HasLibreGraphQuickLink

`func (o *DriveItemCreateLink) HasLibreGraphQuickLink() bool`

HasLibreGraphQuickLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


