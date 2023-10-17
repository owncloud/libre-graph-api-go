# Photo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CameraMake** | Pointer to **string** | Camera manufacturer. Read-only. | [optional] 
**CameraModel** | Pointer to **string** | Camera model. Read-only. | [optional] 
**ExposureDenominator** | Pointer to **float64** | The denominator for the exposure time fraction from the camera. Read-only. | [optional] 
**ExposureNumerator** | Pointer to **float64** | The numerator for the exposure time fraction from the camera. Read-only. | [optional] 
**FNumber** | Pointer to **float64** | The F-stop value from the camera. Read-only. | [optional] 
**FocalLength** | Pointer to **float64** | The focal length from the camera. Read-only. | [optional] 
**Iso** | Pointer to **int32** | The ISO value from the camera. Read-only. | [optional] 
**Orientation** | Pointer to **int32** | The orientation value from the camera. Read-only. | [optional] 
**TakenDateTime** | Pointer to **time.Time** | Represents the date and time the photo was taken. Read-only. | [optional] 

## Methods

### NewPhoto

`func NewPhoto() *Photo`

NewPhoto instantiates a new Photo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhotoWithDefaults

`func NewPhotoWithDefaults() *Photo`

NewPhotoWithDefaults instantiates a new Photo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCameraMake

`func (o *Photo) GetCameraMake() string`

GetCameraMake returns the CameraMake field if non-nil, zero value otherwise.

### GetCameraMakeOk

`func (o *Photo) GetCameraMakeOk() (*string, bool)`

GetCameraMakeOk returns a tuple with the CameraMake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraMake

`func (o *Photo) SetCameraMake(v string)`

SetCameraMake sets CameraMake field to given value.

### HasCameraMake

`func (o *Photo) HasCameraMake() bool`

HasCameraMake returns a boolean if a field has been set.

### GetCameraModel

`func (o *Photo) GetCameraModel() string`

GetCameraModel returns the CameraModel field if non-nil, zero value otherwise.

### GetCameraModelOk

`func (o *Photo) GetCameraModelOk() (*string, bool)`

GetCameraModelOk returns a tuple with the CameraModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraModel

`func (o *Photo) SetCameraModel(v string)`

SetCameraModel sets CameraModel field to given value.

### HasCameraModel

`func (o *Photo) HasCameraModel() bool`

HasCameraModel returns a boolean if a field has been set.

### GetExposureDenominator

`func (o *Photo) GetExposureDenominator() float64`

GetExposureDenominator returns the ExposureDenominator field if non-nil, zero value otherwise.

### GetExposureDenominatorOk

`func (o *Photo) GetExposureDenominatorOk() (*float64, bool)`

GetExposureDenominatorOk returns a tuple with the ExposureDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureDenominator

`func (o *Photo) SetExposureDenominator(v float64)`

SetExposureDenominator sets ExposureDenominator field to given value.

### HasExposureDenominator

`func (o *Photo) HasExposureDenominator() bool`

HasExposureDenominator returns a boolean if a field has been set.

### GetExposureNumerator

`func (o *Photo) GetExposureNumerator() float64`

GetExposureNumerator returns the ExposureNumerator field if non-nil, zero value otherwise.

### GetExposureNumeratorOk

`func (o *Photo) GetExposureNumeratorOk() (*float64, bool)`

GetExposureNumeratorOk returns a tuple with the ExposureNumerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureNumerator

`func (o *Photo) SetExposureNumerator(v float64)`

SetExposureNumerator sets ExposureNumerator field to given value.

### HasExposureNumerator

`func (o *Photo) HasExposureNumerator() bool`

HasExposureNumerator returns a boolean if a field has been set.

### GetFNumber

`func (o *Photo) GetFNumber() float64`

GetFNumber returns the FNumber field if non-nil, zero value otherwise.

### GetFNumberOk

`func (o *Photo) GetFNumberOk() (*float64, bool)`

GetFNumberOk returns a tuple with the FNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFNumber

`func (o *Photo) SetFNumber(v float64)`

SetFNumber sets FNumber field to given value.

### HasFNumber

`func (o *Photo) HasFNumber() bool`

HasFNumber returns a boolean if a field has been set.

### GetFocalLength

`func (o *Photo) GetFocalLength() float64`

GetFocalLength returns the FocalLength field if non-nil, zero value otherwise.

### GetFocalLengthOk

`func (o *Photo) GetFocalLengthOk() (*float64, bool)`

GetFocalLengthOk returns a tuple with the FocalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocalLength

`func (o *Photo) SetFocalLength(v float64)`

SetFocalLength sets FocalLength field to given value.

### HasFocalLength

`func (o *Photo) HasFocalLength() bool`

HasFocalLength returns a boolean if a field has been set.

### GetIso

`func (o *Photo) GetIso() int32`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *Photo) GetIsoOk() (*int32, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *Photo) SetIso(v int32)`

SetIso sets Iso field to given value.

### HasIso

`func (o *Photo) HasIso() bool`

HasIso returns a boolean if a field has been set.

### GetOrientation

`func (o *Photo) GetOrientation() int32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *Photo) GetOrientationOk() (*int32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *Photo) SetOrientation(v int32)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *Photo) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetTakenDateTime

`func (o *Photo) GetTakenDateTime() time.Time`

GetTakenDateTime returns the TakenDateTime field if non-nil, zero value otherwise.

### GetTakenDateTimeOk

`func (o *Photo) GetTakenDateTimeOk() (*time.Time, bool)`

GetTakenDateTimeOk returns a tuple with the TakenDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakenDateTime

`func (o *Photo) SetTakenDateTime(v time.Time)`

SetTakenDateTime sets TakenDateTime field to given value.

### HasTakenDateTime

`func (o *Photo) HasTakenDateTime() bool`

HasTakenDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


