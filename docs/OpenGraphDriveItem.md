# OpenGraphDriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** | The content stream, if the item represents a file. | [optional] 
**CTag** | Pointer to **NullableString** | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only. | [optional] [readonly] 
**Deleted** | Pointer to [**NullableAnyOfopenGraphDeleted**](anyOf&lt;open.graph.deleted&gt;.md) | Information about the deleted state of the item. Read-only. | [optional] [readonly] 
**File** | Pointer to [**NullableAnyOfopenGraphFile**](anyOf&lt;open.graph.file&gt;.md) | File metadata, if the item is a file. Read-only. | [optional] [readonly] 
**FileSystemInfo** | Pointer to [**NullableAnyOfopenGraphFileSystemInfo**](anyOf&lt;open.graph.fileSystemInfo&gt;.md) | File system information on client. Read-write. | [optional] 
**Folder** | Pointer to [**NullableAnyOfopenGraphFolder**](anyOf&lt;open.graph.folder&gt;.md) | Folder metadata, if the item is a folder. Read-only. | [optional] [readonly] 
**Image** | Pointer to [**NullableAnyOfopenGraphImage**](anyOf&lt;open.graph.image&gt;.md) | Image metadata, if the item is an image. Read-only. | [optional] [readonly] 
**Root** | Pointer to [**NullableAnyOfobject**](anyOf&lt;object&gt;.md) | If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive. | [optional] 
**Size** | Pointer to **NullableInt64** | Size of the item in bytes. Read-only. | [optional] [readonly] 
**WebDavUrl** | Pointer to **NullableString** | WebDAV compatible URL for the item. Read-only. | [optional] [readonly] 
**Children** | Pointer to [**[]OpenGraphDriveItem**](OpenGraphDriveItem.md) | Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable. | [optional] [readonly] 

## Methods

### NewOpenGraphDriveItem

`func NewOpenGraphDriveItem() *OpenGraphDriveItem`

NewOpenGraphDriveItem instantiates a new OpenGraphDriveItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphDriveItemWithDefaults

`func NewOpenGraphDriveItemWithDefaults() *OpenGraphDriveItem`

NewOpenGraphDriveItemWithDefaults instantiates a new OpenGraphDriveItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *OpenGraphDriveItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OpenGraphDriveItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OpenGraphDriveItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OpenGraphDriveItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *OpenGraphDriveItem) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *OpenGraphDriveItem) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCTag

`func (o *OpenGraphDriveItem) GetCTag() string`

GetCTag returns the CTag field if non-nil, zero value otherwise.

### GetCTagOk

`func (o *OpenGraphDriveItem) GetCTagOk() (*string, bool)`

GetCTagOk returns a tuple with the CTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTag

`func (o *OpenGraphDriveItem) SetCTag(v string)`

SetCTag sets CTag field to given value.

### HasCTag

`func (o *OpenGraphDriveItem) HasCTag() bool`

HasCTag returns a boolean if a field has been set.

### SetCTagNil

`func (o *OpenGraphDriveItem) SetCTagNil(b bool)`

 SetCTagNil sets the value for CTag to be an explicit nil

### UnsetCTag
`func (o *OpenGraphDriveItem) UnsetCTag()`

UnsetCTag ensures that no value is present for CTag, not even an explicit nil
### GetDeleted

`func (o *OpenGraphDriveItem) GetDeleted() AnyOfopenGraphDeleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OpenGraphDriveItem) GetDeletedOk() (*AnyOfopenGraphDeleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OpenGraphDriveItem) SetDeleted(v AnyOfopenGraphDeleted)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OpenGraphDriveItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *OpenGraphDriveItem) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *OpenGraphDriveItem) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetFile

`func (o *OpenGraphDriveItem) GetFile() AnyOfopenGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *OpenGraphDriveItem) GetFileOk() (*AnyOfopenGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *OpenGraphDriveItem) SetFile(v AnyOfopenGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *OpenGraphDriveItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *OpenGraphDriveItem) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *OpenGraphDriveItem) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetFileSystemInfo

`func (o *OpenGraphDriveItem) GetFileSystemInfo() AnyOfopenGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *OpenGraphDriveItem) GetFileSystemInfoOk() (*AnyOfopenGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *OpenGraphDriveItem) SetFileSystemInfo(v AnyOfopenGraphFileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *OpenGraphDriveItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfoNil

`func (o *OpenGraphDriveItem) SetFileSystemInfoNil(b bool)`

 SetFileSystemInfoNil sets the value for FileSystemInfo to be an explicit nil

### UnsetFileSystemInfo
`func (o *OpenGraphDriveItem) UnsetFileSystemInfo()`

UnsetFileSystemInfo ensures that no value is present for FileSystemInfo, not even an explicit nil
### GetFolder

`func (o *OpenGraphDriveItem) GetFolder() AnyOfopenGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *OpenGraphDriveItem) GetFolderOk() (*AnyOfopenGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *OpenGraphDriveItem) SetFolder(v AnyOfopenGraphFolder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *OpenGraphDriveItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *OpenGraphDriveItem) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *OpenGraphDriveItem) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetImage

`func (o *OpenGraphDriveItem) GetImage() AnyOfopenGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OpenGraphDriveItem) GetImageOk() (*AnyOfopenGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OpenGraphDriveItem) SetImage(v AnyOfopenGraphImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *OpenGraphDriveItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *OpenGraphDriveItem) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *OpenGraphDriveItem) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetRoot

`func (o *OpenGraphDriveItem) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *OpenGraphDriveItem) GetRootOk() (*AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *OpenGraphDriveItem) SetRoot(v AnyOfobject)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *OpenGraphDriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *OpenGraphDriveItem) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *OpenGraphDriveItem) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSize

`func (o *OpenGraphDriveItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OpenGraphDriveItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OpenGraphDriveItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *OpenGraphDriveItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *OpenGraphDriveItem) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *OpenGraphDriveItem) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetWebDavUrl

`func (o *OpenGraphDriveItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *OpenGraphDriveItem) GetWebDavUrlOk() (*string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDavUrl

`func (o *OpenGraphDriveItem) SetWebDavUrl(v string)`

SetWebDavUrl sets WebDavUrl field to given value.

### HasWebDavUrl

`func (o *OpenGraphDriveItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### SetWebDavUrlNil

`func (o *OpenGraphDriveItem) SetWebDavUrlNil(b bool)`

 SetWebDavUrlNil sets the value for WebDavUrl to be an explicit nil

### UnsetWebDavUrl
`func (o *OpenGraphDriveItem) UnsetWebDavUrl()`

UnsetWebDavUrl ensures that no value is present for WebDavUrl, not even an explicit nil
### GetChildren

`func (o *OpenGraphDriveItem) GetChildren() []OpenGraphDriveItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *OpenGraphDriveItem) GetChildrenOk() (*[]OpenGraphDriveItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *OpenGraphDriveItem) SetChildren(v []OpenGraphDriveItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *OpenGraphDriveItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


