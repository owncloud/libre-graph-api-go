# GeoCoordinates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Altitude** | Pointer to **float64** | The altitude (height), in feet, above sea level for the item. Read-only. | [optional] 
**Latitude** | Pointer to **float64** | The latitude, in decimal, for the item. Read-only. | [optional] 
**Longitude** | Pointer to **float64** | The longitude, in decimal, for the item. Read-only. | [optional] 

## Methods

### NewGeoCoordinates

`func NewGeoCoordinates() *GeoCoordinates`

NewGeoCoordinates instantiates a new GeoCoordinates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoCoordinatesWithDefaults

`func NewGeoCoordinatesWithDefaults() *GeoCoordinates`

NewGeoCoordinatesWithDefaults instantiates a new GeoCoordinates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAltitude

`func (o *GeoCoordinates) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *GeoCoordinates) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *GeoCoordinates) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *GeoCoordinates) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetLatitude

`func (o *GeoCoordinates) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GeoCoordinates) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GeoCoordinates) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GeoCoordinates) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *GeoCoordinates) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GeoCoordinates) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GeoCoordinates) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GeoCoordinates) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


