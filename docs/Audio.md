# Audio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Album** | Pointer to **string** | The title of the album for this audio file. | [optional] 
**AlbumArtist** | Pointer to **string** | The artist named on the album for the audio file. | [optional] 
**Artist** | Pointer to **string** | The performing artist for the audio file. | [optional] 
**Bitrate** | Pointer to **int64** | Bitrate expressed in kbps. | [optional] 
**Composers** | Pointer to **string** | The name of the composer of the audio file. | [optional] 
**Copyright** | Pointer to **string** | Copyright information for the audio file. | [optional] 
**Disc** | Pointer to **int32** | The number of the disc this audio file came from. | [optional] 
**DiscCount** | Pointer to **int32** | The total number of discs in this album. | [optional] 
**Duration** | Pointer to **int64** | Duration of the audio file, expressed in milliseconds | [optional] 
**Genre** | Pointer to **string** | The genre of this audio file. | [optional] 
**HasDrm** | Pointer to **bool** | Indicates if the file is protected with digital rights management. | [optional] 
**IsVariableBitrate** | Pointer to **bool** | Indicates if the file is encoded with a variable bitrate. | [optional] 
**Title** | Pointer to **string** | The title of the audio file. | [optional] 
**Track** | Pointer to **int32** | The number of the track on the original disc for this audio file. | [optional] 
**TrackCount** | Pointer to **int32** | The total number of tracks on the original disc for this audio file. | [optional] 
**Year** | Pointer to **int32** | The year the audio file was recorded. | [optional] 

## Methods

### NewAudio

`func NewAudio() *Audio`

NewAudio instantiates a new Audio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioWithDefaults

`func NewAudioWithDefaults() *Audio`

NewAudioWithDefaults instantiates a new Audio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbum

`func (o *Audio) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *Audio) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *Audio) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *Audio) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetAlbumArtist

`func (o *Audio) GetAlbumArtist() string`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *Audio) GetAlbumArtistOk() (*string, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *Audio) SetAlbumArtist(v string)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *Audio) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### GetArtist

`func (o *Audio) GetArtist() string`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *Audio) GetArtistOk() (*string, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *Audio) SetArtist(v string)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *Audio) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### GetBitrate

`func (o *Audio) GetBitrate() int64`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *Audio) GetBitrateOk() (*int64, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *Audio) SetBitrate(v int64)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *Audio) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetComposers

`func (o *Audio) GetComposers() string`

GetComposers returns the Composers field if non-nil, zero value otherwise.

### GetComposersOk

`func (o *Audio) GetComposersOk() (*string, bool)`

GetComposersOk returns a tuple with the Composers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposers

`func (o *Audio) SetComposers(v string)`

SetComposers sets Composers field to given value.

### HasComposers

`func (o *Audio) HasComposers() bool`

HasComposers returns a boolean if a field has been set.

### GetCopyright

`func (o *Audio) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *Audio) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *Audio) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *Audio) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### GetDisc

`func (o *Audio) GetDisc() int32`

GetDisc returns the Disc field if non-nil, zero value otherwise.

### GetDiscOk

`func (o *Audio) GetDiscOk() (*int32, bool)`

GetDiscOk returns a tuple with the Disc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisc

`func (o *Audio) SetDisc(v int32)`

SetDisc sets Disc field to given value.

### HasDisc

`func (o *Audio) HasDisc() bool`

HasDisc returns a boolean if a field has been set.

### GetDiscCount

`func (o *Audio) GetDiscCount() int32`

GetDiscCount returns the DiscCount field if non-nil, zero value otherwise.

### GetDiscCountOk

`func (o *Audio) GetDiscCountOk() (*int32, bool)`

GetDiscCountOk returns a tuple with the DiscCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscCount

`func (o *Audio) SetDiscCount(v int32)`

SetDiscCount sets DiscCount field to given value.

### HasDiscCount

`func (o *Audio) HasDiscCount() bool`

HasDiscCount returns a boolean if a field has been set.

### GetDuration

`func (o *Audio) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Audio) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Audio) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Audio) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGenre

`func (o *Audio) GetGenre() string`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *Audio) GetGenreOk() (*string, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *Audio) SetGenre(v string)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *Audio) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetHasDrm

`func (o *Audio) GetHasDrm() bool`

GetHasDrm returns the HasDrm field if non-nil, zero value otherwise.

### GetHasDrmOk

`func (o *Audio) GetHasDrmOk() (*bool, bool)`

GetHasDrmOk returns a tuple with the HasDrm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDrm

`func (o *Audio) SetHasDrm(v bool)`

SetHasDrm sets HasDrm field to given value.

### HasHasDrm

`func (o *Audio) HasHasDrm() bool`

HasHasDrm returns a boolean if a field has been set.

### GetIsVariableBitrate

`func (o *Audio) GetIsVariableBitrate() bool`

GetIsVariableBitrate returns the IsVariableBitrate field if non-nil, zero value otherwise.

### GetIsVariableBitrateOk

`func (o *Audio) GetIsVariableBitrateOk() (*bool, bool)`

GetIsVariableBitrateOk returns a tuple with the IsVariableBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariableBitrate

`func (o *Audio) SetIsVariableBitrate(v bool)`

SetIsVariableBitrate sets IsVariableBitrate field to given value.

### HasIsVariableBitrate

`func (o *Audio) HasIsVariableBitrate() bool`

HasIsVariableBitrate returns a boolean if a field has been set.

### GetTitle

`func (o *Audio) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Audio) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Audio) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Audio) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTrack

`func (o *Audio) GetTrack() int32`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *Audio) GetTrackOk() (*int32, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *Audio) SetTrack(v int32)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *Audio) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetTrackCount

`func (o *Audio) GetTrackCount() int32`

GetTrackCount returns the TrackCount field if non-nil, zero value otherwise.

### GetTrackCountOk

`func (o *Audio) GetTrackCountOk() (*int32, bool)`

GetTrackCountOk returns a tuple with the TrackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackCount

`func (o *Audio) SetTrackCount(v int32)`

SetTrackCount sets TrackCount field to given value.

### HasTrackCount

`func (o *Audio) HasTrackCount() bool`

HasTrackCount returns a boolean if a field has been set.

### GetYear

`func (o *Audio) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Audio) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Audio) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *Audio) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


