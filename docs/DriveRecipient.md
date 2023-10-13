# DriveRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **string** | The unique identifier for the recipient in the directory. | [optional] 
**LibreGraphRecipientType** | Pointer to **string** | When the recipient is referenced by objectId this annotation is used to differentiate &#x60;user&#x60; and &#x60;group&#x60; recipients. | [optional] [default to "user"]

## Methods

### NewDriveRecipient

`func NewDriveRecipient() *DriveRecipient`

NewDriveRecipient instantiates a new DriveRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveRecipientWithDefaults

`func NewDriveRecipientWithDefaults() *DriveRecipient`

NewDriveRecipientWithDefaults instantiates a new DriveRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *DriveRecipient) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *DriveRecipient) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *DriveRecipient) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *DriveRecipient) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetLibreGraphRecipientType

`func (o *DriveRecipient) GetLibreGraphRecipientType() string`

GetLibreGraphRecipientType returns the LibreGraphRecipientType field if non-nil, zero value otherwise.

### GetLibreGraphRecipientTypeOk

`func (o *DriveRecipient) GetLibreGraphRecipientTypeOk() (*string, bool)`

GetLibreGraphRecipientTypeOk returns a tuple with the LibreGraphRecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphRecipientType

`func (o *DriveRecipient) SetLibreGraphRecipientType(v string)`

SetLibreGraphRecipientType sets LibreGraphRecipientType field to given value.

### HasLibreGraphRecipientType

`func (o *DriveRecipient) HasLibreGraphRecipientType() bool`

HasLibreGraphRecipientType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


