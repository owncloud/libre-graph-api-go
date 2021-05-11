# OpenGraphDriveItemAllOf

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

### NewOpenGraphDriveItemAllOf

`func NewOpenGraphDriveItemAllOf() *OpenGraphDriveItemAllOf`

NewOpenGraphDriveItemAllOf instantiates a new OpenGraphDriveItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphDriveItemAllOfWithDefaults

`func NewOpenGraphDriveItemAllOfWithDefaults() *OpenGraphDriveItemAllOf`

NewOpenGraphDriveItemAllOfWithDefaults instantiates a new OpenGraphDriveItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *OpenGraphDriveItemAllOf) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OpenGraphDriveItemAllOf) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OpenGraphDriveItemAllOf) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OpenGraphDriveItemAllOf) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *OpenGraphDriveItemAllOf) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *OpenGraphDriveItemAllOf) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCTag

`func (o *OpenGraphDriveItemAllOf) GetCTag() string`

GetCTag returns the CTag field if non-nil, zero value otherwise.

### GetCTagOk

`func (o *OpenGraphDriveItemAllOf) GetCTagOk() (*string, bool)`

GetCTagOk returns a tuple with the CTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTag

`func (o *OpenGraphDriveItemAllOf) SetCTag(v string)`

SetCTag sets CTag field to given value.

### HasCTag

`func (o *OpenGraphDriveItemAllOf) HasCTag() bool`

HasCTag returns a boolean if a field has been set.

### SetCTagNil

`func (o *OpenGraphDriveItemAllOf) SetCTagNil(b bool)`

 SetCTagNil sets the value for CTag to be an explicit nil

### UnsetCTag
`func (o *OpenGraphDriveItemAllOf) UnsetCTag()`

UnsetCTag ensures that no value is present for CTag, not even an explicit nil
### GetDeleted

`func (o *OpenGraphDriveItemAllOf) GetDeleted() AnyOfopenGraphDeleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OpenGraphDriveItemAllOf) GetDeletedOk() (*AnyOfopenGraphDeleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OpenGraphDriveItemAllOf) SetDeleted(v AnyOfopenGraphDeleted)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OpenGraphDriveItemAllOf) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *OpenGraphDriveItemAllOf) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *OpenGraphDriveItemAllOf) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetFile

`func (o *OpenGraphDriveItemAllOf) GetFile() AnyOfopenGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *OpenGraphDriveItemAllOf) GetFileOk() (*AnyOfopenGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *OpenGraphDriveItemAllOf) SetFile(v AnyOfopenGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *OpenGraphDriveItemAllOf) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *OpenGraphDriveItemAllOf) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *OpenGraphDriveItemAllOf) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetFileSystemInfo

`func (o *OpenGraphDriveItemAllOf) GetFileSystemInfo() AnyOfopenGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *OpenGraphDriveItemAllOf) GetFileSystemInfoOk() (*AnyOfopenGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *OpenGraphDriveItemAllOf) SetFileSystemInfo(v AnyOfopenGraphFileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *OpenGraphDriveItemAllOf) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfoNil

`func (o *OpenGraphDriveItemAllOf) SetFileSystemInfoNil(b bool)`

 SetFileSystemInfoNil sets the value for FileSystemInfo to be an explicit nil

### UnsetFileSystemInfo
`func (o *OpenGraphDriveItemAllOf) UnsetFileSystemInfo()`

UnsetFileSystemInfo ensures that no value is present for FileSystemInfo, not even an explicit nil
### GetFolder

`func (o *OpenGraphDriveItemAllOf) GetFolder() AnyOfopenGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *OpenGraphDriveItemAllOf) GetFolderOk() (*AnyOfopenGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *OpenGraphDriveItemAllOf) SetFolder(v AnyOfopenGraphFolder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *OpenGraphDriveItemAllOf) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *OpenGraphDriveItemAllOf) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *OpenGraphDriveItemAllOf) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetImage

`func (o *OpenGraphDriveItemAllOf) GetImage() AnyOfopenGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OpenGraphDriveItemAllOf) GetImageOk() (*AnyOfopenGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OpenGraphDriveItemAllOf) SetImage(v AnyOfopenGraphImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *OpenGraphDriveItemAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *OpenGraphDriveItemAllOf) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *OpenGraphDriveItemAllOf) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetRoot

`func (o *OpenGraphDriveItemAllOf) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *OpenGraphDriveItemAllOf) GetRootOk() (*AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *OpenGraphDriveItemAllOf) SetRoot(v AnyOfobject)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *OpenGraphDriveItemAllOf) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *OpenGraphDriveItemAllOf) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *OpenGraphDriveItemAllOf) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSize

`func (o *OpenGraphDriveItemAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OpenGraphDriveItemAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OpenGraphDriveItemAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *OpenGraphDriveItemAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *OpenGraphDriveItemAllOf) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *OpenGraphDriveItemAllOf) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetWebDavUrl

`func (o *OpenGraphDriveItemAllOf) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *OpenGraphDriveItemAllOf) GetWebDavUrlOk() (*string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDavUrl

`func (o *OpenGraphDriveItemAllOf) SetWebDavUrl(v string)`

SetWebDavUrl sets WebDavUrl field to given value.

### HasWebDavUrl

`func (o *OpenGraphDriveItemAllOf) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### SetWebDavUrlNil

`func (o *OpenGraphDriveItemAllOf) SetWebDavUrlNil(b bool)`

 SetWebDavUrlNil sets the value for WebDavUrl to be an explicit nil

### UnsetWebDavUrl
`func (o *OpenGraphDriveItemAllOf) UnsetWebDavUrl()`

UnsetWebDavUrl ensures that no value is present for WebDavUrl, not even an explicit nil
### GetChildren

`func (o *OpenGraphDriveItemAllOf) GetChildren() []OpenGraphDriveItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *OpenGraphDriveItemAllOf) GetChildrenOk() (*[]OpenGraphDriveItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *OpenGraphDriveItemAllOf) SetChildren(v []OpenGraphDriveItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *OpenGraphDriveItemAllOf) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


