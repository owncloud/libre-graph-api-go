# ExportPersonalDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageLocation** | Pointer to **string** | the path where the file should be created in the users personal space | [optional] 

## Methods

### NewExportPersonalDataRequest

`func NewExportPersonalDataRequest() *ExportPersonalDataRequest`

NewExportPersonalDataRequest instantiates a new ExportPersonalDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportPersonalDataRequestWithDefaults

`func NewExportPersonalDataRequestWithDefaults() *ExportPersonalDataRequest`

NewExportPersonalDataRequestWithDefaults instantiates a new ExportPersonalDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageLocation

`func (o *ExportPersonalDataRequest) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *ExportPersonalDataRequest) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *ExportPersonalDataRequest) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *ExportPersonalDataRequest) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


