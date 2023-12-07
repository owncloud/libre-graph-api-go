# DriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] [readonly] 
**CreatedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of item creation. Read-only. | [optional] [readonly] 
**Description** | Pointer to **string** | Provides a user-visible description of the item. Optional. | [optional] 
**ETag** | Pointer to **string** | ETag for the item. Read-only. | [optional] [readonly] 
**LastModifiedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Date and time the item was last modified. Read-only. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the item. Read-write. | [optional] 
**ParentReference** | Pointer to [**ItemReference**](ItemReference.md) |  | [optional] 
**WebUrl** | Pointer to **string** | URL that displays the resource in the browser. Read-only. | [optional] [readonly] 
**Content** | Pointer to **string** | The content stream, if the item represents a file. | [optional] 
**CTag** | Pointer to **string** | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only. | [optional] [readonly] 
**Deleted** | Pointer to [**Deleted**](Deleted.md) |  | [optional] 
**File** | Pointer to [**OpenGraphFile**](OpenGraphFile.md) |  | [optional] 
**FileSystemInfo** | Pointer to [**FileSystemInfo**](FileSystemInfo.md) |  | [optional] 
**Folder** | Pointer to [**Folder**](Folder.md) |  | [optional] 
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**Photo** | Pointer to [**Photo**](Photo.md) |  | [optional] 
**Location** | Pointer to [**GeoCoordinates**](GeoCoordinates.md) |  | [optional] 
**Root** | Pointer to **map[string]interface{}** | If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive. | [optional] 
**Trash** | Pointer to [**Trash**](Trash.md) |  | [optional] 
**SpecialFolder** | Pointer to [**SpecialFolder**](SpecialFolder.md) |  | [optional] 
**RemoteItem** | Pointer to [**RemoteItem**](RemoteItem.md) |  | [optional] 
**Size** | Pointer to **int64** | Size of the item in bytes. Read-only. | [optional] [readonly] 
**WebDavUrl** | Pointer to **string** | WebDAV compatible URL for the item. Read-only. | [optional] [readonly] 
**Children** | Pointer to [**[]DriveItem**](DriveItem.md) | Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable. | [optional] [readonly] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) | The set of permissions for the item. Read-only. Nullable. | [optional] [readonly] 
**Audio** | Pointer to [**Audio**](Audio.md) |  | [optional] 
**Video** | Pointer to [**Video**](Video.md) |  | [optional] 

## Methods

### NewDriveItem

`func NewDriveItem() *DriveItem`

NewDriveItem instantiates a new DriveItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveItemWithDefaults

`func NewDriveItemWithDefaults() *DriveItem`

NewDriveItemWithDefaults instantiates a new DriveItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriveItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriveItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriveItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DriveItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DriveItem) GetCreatedBy() IdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DriveItem) GetCreatedByOk() (*IdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DriveItem) SetCreatedBy(v IdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DriveItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *DriveItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *DriveItem) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *DriveItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *DriveItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *DriveItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DriveItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DriveItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DriveItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetETag

`func (o *DriveItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *DriveItem) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *DriveItem) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *DriveItem) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *DriveItem) GetLastModifiedBy() IdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *DriveItem) GetLastModifiedByOk() (*IdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *DriveItem) SetLastModifiedBy(v IdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *DriveItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *DriveItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *DriveItem) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *DriveItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *DriveItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *DriveItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DriveItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DriveItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DriveItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentReference

`func (o *DriveItem) GetParentReference() ItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *DriveItem) GetParentReferenceOk() (*ItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *DriveItem) SetParentReference(v ItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *DriveItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetWebUrl

`func (o *DriveItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *DriveItem) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *DriveItem) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *DriveItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetContent

`func (o *DriveItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DriveItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DriveItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DriveItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCTag

`func (o *DriveItem) GetCTag() string`

GetCTag returns the CTag field if non-nil, zero value otherwise.

### GetCTagOk

`func (o *DriveItem) GetCTagOk() (*string, bool)`

GetCTagOk returns a tuple with the CTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTag

`func (o *DriveItem) SetCTag(v string)`

SetCTag sets CTag field to given value.

### HasCTag

`func (o *DriveItem) HasCTag() bool`

HasCTag returns a boolean if a field has been set.

### GetDeleted

`func (o *DriveItem) GetDeleted() Deleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DriveItem) GetDeletedOk() (*Deleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DriveItem) SetDeleted(v Deleted)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DriveItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFile

`func (o *DriveItem) GetFile() OpenGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *DriveItem) GetFileOk() (*OpenGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *DriveItem) SetFile(v OpenGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *DriveItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFileSystemInfo

`func (o *DriveItem) GetFileSystemInfo() FileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *DriveItem) GetFileSystemInfoOk() (*FileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *DriveItem) SetFileSystemInfo(v FileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *DriveItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### GetFolder

`func (o *DriveItem) GetFolder() Folder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *DriveItem) GetFolderOk() (*Folder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *DriveItem) SetFolder(v Folder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *DriveItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetImage

`func (o *DriveItem) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DriveItem) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DriveItem) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *DriveItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPhoto

`func (o *DriveItem) GetPhoto() Photo`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *DriveItem) GetPhotoOk() (*Photo, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *DriveItem) SetPhoto(v Photo)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *DriveItem) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetLocation

`func (o *DriveItem) GetLocation() GeoCoordinates`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DriveItem) GetLocationOk() (*GeoCoordinates, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DriveItem) SetLocation(v GeoCoordinates)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DriveItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRoot

`func (o *DriveItem) GetRoot() map[string]interface{}`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *DriveItem) GetRootOk() (*map[string]interface{}, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *DriveItem) SetRoot(v map[string]interface{})`

SetRoot sets Root field to given value.

### HasRoot

`func (o *DriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetTrash

`func (o *DriveItem) GetTrash() Trash`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *DriveItem) GetTrashOk() (*Trash, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *DriveItem) SetTrash(v Trash)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *DriveItem) HasTrash() bool`

HasTrash returns a boolean if a field has been set.

### GetSpecialFolder

`func (o *DriveItem) GetSpecialFolder() SpecialFolder`

GetSpecialFolder returns the SpecialFolder field if non-nil, zero value otherwise.

### GetSpecialFolderOk

`func (o *DriveItem) GetSpecialFolderOk() (*SpecialFolder, bool)`

GetSpecialFolderOk returns a tuple with the SpecialFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialFolder

`func (o *DriveItem) SetSpecialFolder(v SpecialFolder)`

SetSpecialFolder sets SpecialFolder field to given value.

### HasSpecialFolder

`func (o *DriveItem) HasSpecialFolder() bool`

HasSpecialFolder returns a boolean if a field has been set.

### GetRemoteItem

`func (o *DriveItem) GetRemoteItem() RemoteItem`

GetRemoteItem returns the RemoteItem field if non-nil, zero value otherwise.

### GetRemoteItemOk

`func (o *DriveItem) GetRemoteItemOk() (*RemoteItem, bool)`

GetRemoteItemOk returns a tuple with the RemoteItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteItem

`func (o *DriveItem) SetRemoteItem(v RemoteItem)`

SetRemoteItem sets RemoteItem field to given value.

### HasRemoteItem

`func (o *DriveItem) HasRemoteItem() bool`

HasRemoteItem returns a boolean if a field has been set.

### GetSize

`func (o *DriveItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DriveItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DriveItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DriveItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetWebDavUrl

`func (o *DriveItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *DriveItem) GetWebDavUrlOk() (*string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDavUrl

`func (o *DriveItem) SetWebDavUrl(v string)`

SetWebDavUrl sets WebDavUrl field to given value.

### HasWebDavUrl

`func (o *DriveItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### GetChildren

`func (o *DriveItem) GetChildren() []DriveItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *DriveItem) GetChildrenOk() (*[]DriveItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *DriveItem) SetChildren(v []DriveItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *DriveItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetPermissions

`func (o *DriveItem) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DriveItem) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DriveItem) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DriveItem) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAudio

`func (o *DriveItem) GetAudio() Audio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *DriveItem) GetAudioOk() (*Audio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *DriveItem) SetAudio(v Audio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *DriveItem) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetVideo

`func (o *DriveItem) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *DriveItem) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *DriveItem) SetVideo(v Video)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *DriveItem) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


