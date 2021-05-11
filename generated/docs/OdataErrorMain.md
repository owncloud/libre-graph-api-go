# OdataErrorMain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Message** | **string** |  | 
**Target** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]OdataErrorDetail**](OdataErrorDetail.md) |  | [optional] 
**Innererror** | Pointer to **map[string]interface{}** | The structure of this object is service-specific | [optional] 

## Methods

### NewOdataErrorMain

`func NewOdataErrorMain(code string, message string, ) *OdataErrorMain`

NewOdataErrorMain instantiates a new OdataErrorMain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOdataErrorMainWithDefaults

`func NewOdataErrorMainWithDefaults() *OdataErrorMain`

NewOdataErrorMainWithDefaults instantiates a new OdataErrorMain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OdataErrorMain) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OdataErrorMain) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OdataErrorMain) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *OdataErrorMain) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OdataErrorMain) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OdataErrorMain) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTarget

`func (o *OdataErrorMain) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *OdataErrorMain) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *OdataErrorMain) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *OdataErrorMain) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetDetails

`func (o *OdataErrorMain) GetDetails() []OdataErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OdataErrorMain) GetDetailsOk() (*[]OdataErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *OdataErrorMain) SetDetails(v []OdataErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *OdataErrorMain) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInnererror

`func (o *OdataErrorMain) GetInnererror() map[string]interface{}`

GetInnererror returns the Innererror field if non-nil, zero value otherwise.

### GetInnererrorOk

`func (o *OdataErrorMain) GetInnererrorOk() (*map[string]interface{}, bool)`

GetInnererrorOk returns a tuple with the Innererror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnererror

`func (o *OdataErrorMain) SetInnererror(v map[string]interface{})`

SetInnererror sets Innererror field to given value.

### HasInnererror

`func (o *OdataErrorMain) HasInnererror() bool`

HasInnererror returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


