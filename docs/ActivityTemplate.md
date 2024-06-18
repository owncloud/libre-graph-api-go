# ActivityTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Activity description. | 
**Variables** | Pointer to **map[string]interface{}** | Activity description variables. | [optional] 

## Methods

### NewActivityTemplate

`func NewActivityTemplate(message string, ) *ActivityTemplate`

NewActivityTemplate instantiates a new ActivityTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTemplateWithDefaults

`func NewActivityTemplateWithDefaults() *ActivityTemplate`

NewActivityTemplateWithDefaults instantiates a new ActivityTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ActivityTemplate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActivityTemplate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActivityTemplate) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetVariables

`func (o *ActivityTemplate) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ActivityTemplate) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ActivityTemplate) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ActivityTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


