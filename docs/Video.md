# Video

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioBitsPerSample** | Pointer to **int32** | Number of audio bits per sample. | [optional] 
**AudioChannels** | Pointer to **int32** | Number of audio channels. | [optional] 
**AudioFormat** | Pointer to **string** | Name of the audio format (AAC, MP3, etc.). | [optional] 
**AudioSamplesPerSecond** | Pointer to **int32** | Number of audio samples per second. | [optional] 
**Bitrate** | Pointer to **int32** | Bit rate of the video in bits per second. | [optional] 
**Duration** | Pointer to **int64** | Duration of the file in milliseconds. | [optional] 
**FourCC** | Pointer to **string** | \\\&quot;Four character code\\\&quot; name of the video format. | [optional] 
**FrameRate** | Pointer to **float64** | Frame rate of the video. | [optional] 
**Height** | Pointer to **int32** | Height of the video, in pixels. | [optional] 
**Width** | Pointer to **int32** | Width of the video, in pixels. | [optional] 

## Methods

### NewVideo

`func NewVideo() *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioBitsPerSample

`func (o *Video) GetAudioBitsPerSample() int32`

GetAudioBitsPerSample returns the AudioBitsPerSample field if non-nil, zero value otherwise.

### GetAudioBitsPerSampleOk

`func (o *Video) GetAudioBitsPerSampleOk() (*int32, bool)`

GetAudioBitsPerSampleOk returns a tuple with the AudioBitsPerSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBitsPerSample

`func (o *Video) SetAudioBitsPerSample(v int32)`

SetAudioBitsPerSample sets AudioBitsPerSample field to given value.

### HasAudioBitsPerSample

`func (o *Video) HasAudioBitsPerSample() bool`

HasAudioBitsPerSample returns a boolean if a field has been set.

### GetAudioChannels

`func (o *Video) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *Video) GetAudioChannelsOk() (*int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *Video) SetAudioChannels(v int32)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *Video) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### GetAudioFormat

`func (o *Video) GetAudioFormat() string`

GetAudioFormat returns the AudioFormat field if non-nil, zero value otherwise.

### GetAudioFormatOk

`func (o *Video) GetAudioFormatOk() (*string, bool)`

GetAudioFormatOk returns a tuple with the AudioFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioFormat

`func (o *Video) SetAudioFormat(v string)`

SetAudioFormat sets AudioFormat field to given value.

### HasAudioFormat

`func (o *Video) HasAudioFormat() bool`

HasAudioFormat returns a boolean if a field has been set.

### GetAudioSamplesPerSecond

`func (o *Video) GetAudioSamplesPerSecond() int32`

GetAudioSamplesPerSecond returns the AudioSamplesPerSecond field if non-nil, zero value otherwise.

### GetAudioSamplesPerSecondOk

`func (o *Video) GetAudioSamplesPerSecondOk() (*int32, bool)`

GetAudioSamplesPerSecondOk returns a tuple with the AudioSamplesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSamplesPerSecond

`func (o *Video) SetAudioSamplesPerSecond(v int32)`

SetAudioSamplesPerSecond sets AudioSamplesPerSecond field to given value.

### HasAudioSamplesPerSecond

`func (o *Video) HasAudioSamplesPerSecond() bool`

HasAudioSamplesPerSecond returns a boolean if a field has been set.

### GetBitrate

`func (o *Video) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *Video) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *Video) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *Video) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetDuration

`func (o *Video) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Video) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Video) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Video) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFourCC

`func (o *Video) GetFourCC() string`

GetFourCC returns the FourCC field if non-nil, zero value otherwise.

### GetFourCCOk

`func (o *Video) GetFourCCOk() (*string, bool)`

GetFourCCOk returns a tuple with the FourCC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFourCC

`func (o *Video) SetFourCC(v string)`

SetFourCC sets FourCC field to given value.

### HasFourCC

`func (o *Video) HasFourCC() bool`

HasFourCC returns a boolean if a field has been set.

### GetFrameRate

`func (o *Video) GetFrameRate() float64`

GetFrameRate returns the FrameRate field if non-nil, zero value otherwise.

### GetFrameRateOk

`func (o *Video) GetFrameRateOk() (*float64, bool)`

GetFrameRateOk returns a tuple with the FrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameRate

`func (o *Video) SetFrameRate(v float64)`

SetFrameRate sets FrameRate field to given value.

### HasFrameRate

`func (o *Video) HasFrameRate() bool`

HasFrameRate returns a boolean if a field has been set.

### GetHeight

`func (o *Video) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Video) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Video) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Video) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *Video) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Video) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Video) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Video) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


