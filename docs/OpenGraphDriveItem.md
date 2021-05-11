# OpenGraphDriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The content stream, if the item represents a file. | [optional] 
**CTag** | Pointer to **string** | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only. | [optional] [readonly] 
**Deleted** | Pointer to [**OpenGraphDeleted**](open.graph.deleted.md) | Information about the deleted state of the item. Read-only. | [optional] [readonly] 
**File** | Pointer to [**OpenGraphFile**](open.graph.file.md) | File metadata, if the item is a file. Read-only. | [optional] [readonly] 
**FileSystemInfo** | Pointer to [**OpenGraphFileSystemInfo**](open.graph.fileSystemInfo.md) | File system information on client. Read-write. | [optional] 
**Folder** | Pointer to [**OpenGraphFolder**](open.graph.folder.md) | Folder metadata, if the item is a folder. Read-only. | [optional] [readonly] 
**Image** | Pointer to [**OpenGraphImage**](open.graph.image.md) | Image metadata, if the item is an image. Read-only. | [optional] [readonly] 
**Root** | Pointer to **map[string]interface{}** | If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive. | [optional] 
**Size** | Pointer to **int64** | Size of the item in bytes. Read-only. | [optional] [readonly] 
**WebDavUrl** | Pointer to **string** | WebDAV compatible URL for the item. Read-only. | [optional] [readonly] 
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

### GetDeleted

`func (o *OpenGraphDriveItem) GetDeleted() OpenGraphDeleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OpenGraphDriveItem) GetDeletedOk() (*OpenGraphDeleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OpenGraphDriveItem) SetDeleted(v OpenGraphDeleted)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OpenGraphDriveItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFile

`func (o *OpenGraphDriveItem) GetFile() OpenGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *OpenGraphDriveItem) GetFileOk() (*OpenGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *OpenGraphDriveItem) SetFile(v OpenGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *OpenGraphDriveItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFileSystemInfo

`func (o *OpenGraphDriveItem) GetFileSystemInfo() OpenGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *OpenGraphDriveItem) GetFileSystemInfoOk() (*OpenGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *OpenGraphDriveItem) SetFileSystemInfo(v OpenGraphFileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *OpenGraphDriveItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### GetFolder

`func (o *OpenGraphDriveItem) GetFolder() OpenGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *OpenGraphDriveItem) GetFolderOk() (*OpenGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *OpenGraphDriveItem) SetFolder(v OpenGraphFolder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *OpenGraphDriveItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetImage

`func (o *OpenGraphDriveItem) GetImage() OpenGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OpenGraphDriveItem) GetImageOk() (*OpenGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OpenGraphDriveItem) SetImage(v OpenGraphImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *OpenGraphDriveItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRoot

`func (o *OpenGraphDriveItem) GetRoot() map[string]interface{}`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *OpenGraphDriveItem) GetRootOk() (*map[string]interface{}, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *OpenGraphDriveItem) SetRoot(v map[string]interface{})`

SetRoot sets Root field to given value.

### HasRoot

`func (o *OpenGraphDriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

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


