# OpenGraphDriveItemAllOf

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

### GetDeleted

`func (o *OpenGraphDriveItemAllOf) GetDeleted() OpenGraphDeleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OpenGraphDriveItemAllOf) GetDeletedOk() (*OpenGraphDeleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OpenGraphDriveItemAllOf) SetDeleted(v OpenGraphDeleted)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OpenGraphDriveItemAllOf) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFile

`func (o *OpenGraphDriveItemAllOf) GetFile() OpenGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *OpenGraphDriveItemAllOf) GetFileOk() (*OpenGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *OpenGraphDriveItemAllOf) SetFile(v OpenGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *OpenGraphDriveItemAllOf) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFileSystemInfo

`func (o *OpenGraphDriveItemAllOf) GetFileSystemInfo() OpenGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *OpenGraphDriveItemAllOf) GetFileSystemInfoOk() (*OpenGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *OpenGraphDriveItemAllOf) SetFileSystemInfo(v OpenGraphFileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *OpenGraphDriveItemAllOf) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### GetFolder

`func (o *OpenGraphDriveItemAllOf) GetFolder() OpenGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *OpenGraphDriveItemAllOf) GetFolderOk() (*OpenGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *OpenGraphDriveItemAllOf) SetFolder(v OpenGraphFolder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *OpenGraphDriveItemAllOf) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetImage

`func (o *OpenGraphDriveItemAllOf) GetImage() OpenGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OpenGraphDriveItemAllOf) GetImageOk() (*OpenGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OpenGraphDriveItemAllOf) SetImage(v OpenGraphImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *OpenGraphDriveItemAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRoot

`func (o *OpenGraphDriveItemAllOf) GetRoot() map[string]interface{}`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *OpenGraphDriveItemAllOf) GetRootOk() (*map[string]interface{}, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *OpenGraphDriveItemAllOf) SetRoot(v map[string]interface{})`

SetRoot sets Root field to given value.

### HasRoot

`func (o *OpenGraphDriveItemAllOf) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

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


