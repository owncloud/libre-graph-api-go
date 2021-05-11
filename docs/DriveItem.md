# DriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The content stream, if the item represents a file. | [optional] 
**CTag** | Pointer to **string** | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only. | [optional] [readonly] 
**Deleted** | Pointer to [**Deleted**](deleted.md) | Information about the deleted state of the item. Read-only. | [optional] [readonly] 
**File** | Pointer to [**OpenGraphFile**](openGraphFile.md) | File metadata, if the item is a file. Read-only. | [optional] [readonly] 
**FileSystemInfo** | Pointer to [**FileSystemInfo**](fileSystemInfo.md) | File system information on client. Read-write. | [optional] 
**Folder** | Pointer to [**Folder**](folder.md) | Folder metadata, if the item is a folder. Read-only. | [optional] [readonly] 
**Image** | Pointer to [**Image**](image.md) | Image metadata, if the item is an image. Read-only. | [optional] [readonly] 
**Root** | Pointer to **map[string]interface{}** | If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive. | [optional] 
**Size** | Pointer to **int64** | Size of the item in bytes. Read-only. | [optional] [readonly] 
**WebDavUrl** | Pointer to **string** | WebDAV compatible URL for the item. Read-only. | [optional] [readonly] 
**Children** | Pointer to [**[]DriveItem**](DriveItem.md) | Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable. | [optional] [readonly] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


