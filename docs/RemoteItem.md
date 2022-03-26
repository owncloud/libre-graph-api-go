# RemoteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of item creation. Read-only. | [optional] 
**File** | Pointer to [**OpenGraphFile**](OpenGraphFile.md) |  | [optional] 
**FileSystemInfo** | Pointer to [**FileSystemInfo**](FileSystemInfo.md) |  | [optional] 
**Folder** | Pointer to [**Folder**](Folder.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier for the remote item in its drive. Read-only. | [optional] 
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**LastModifiedBy** | Pointer to [**IdentitySet**](IdentitySet.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Date and time the item was last modified. Read-only. | [optional] 
**Name** | Pointer to **string** | Optional. Filename of the remote item. Read-only. | [optional] 
**ETag** | Pointer to **string** | ETag for the item. Read-only. | [optional] [readonly] 
**CTag** | Pointer to **string** | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only. | [optional] [readonly] 
**ParentReference** | Pointer to [**ItemReference**](ItemReference.md) |  | [optional] 
**Shared** | Pointer to [**Shared**](Shared.md) |  | [optional] 
**Size** | Pointer to **int64** | Size of the remote item. Read-only. | [optional] 
**SpecialFolder** | Pointer to [**SpecialFolder**](SpecialFolder.md) |  | [optional] 
**WebDavUrl** | Pointer to **string** | DAV compatible URL for the item. | [optional] 
**WebUrl** | Pointer to **string** | URL that displays the resource in the browser. Read-only. | [optional] 

## Methods

### NewRemoteItem

`func NewRemoteItem() *RemoteItem`

NewRemoteItem instantiates a new RemoteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteItemWithDefaults

`func NewRemoteItemWithDefaults() *RemoteItem`

NewRemoteItemWithDefaults instantiates a new RemoteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *RemoteItem) GetCreatedBy() IdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RemoteItem) GetCreatedByOk() (*IdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RemoteItem) SetCreatedBy(v IdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RemoteItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *RemoteItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *RemoteItem) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *RemoteItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *RemoteItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetFile

`func (o *RemoteItem) GetFile() OpenGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *RemoteItem) GetFileOk() (*OpenGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *RemoteItem) SetFile(v OpenGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *RemoteItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFileSystemInfo

`func (o *RemoteItem) GetFileSystemInfo() FileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *RemoteItem) GetFileSystemInfoOk() (*FileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *RemoteItem) SetFileSystemInfo(v FileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *RemoteItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### GetFolder

`func (o *RemoteItem) GetFolder() Folder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *RemoteItem) GetFolderOk() (*Folder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *RemoteItem) SetFolder(v Folder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *RemoteItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetId

`func (o *RemoteItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *RemoteItem) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RemoteItem) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RemoteItem) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *RemoteItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *RemoteItem) GetLastModifiedBy() IdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *RemoteItem) GetLastModifiedByOk() (*IdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *RemoteItem) SetLastModifiedBy(v IdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *RemoteItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *RemoteItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *RemoteItem) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *RemoteItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *RemoteItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *RemoteItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetETag

`func (o *RemoteItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *RemoteItem) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *RemoteItem) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *RemoteItem) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetCTag

`func (o *RemoteItem) GetCTag() string`

GetCTag returns the CTag field if non-nil, zero value otherwise.

### GetCTagOk

`func (o *RemoteItem) GetCTagOk() (*string, bool)`

GetCTagOk returns a tuple with the CTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTag

`func (o *RemoteItem) SetCTag(v string)`

SetCTag sets CTag field to given value.

### HasCTag

`func (o *RemoteItem) HasCTag() bool`

HasCTag returns a boolean if a field has been set.

### GetParentReference

`func (o *RemoteItem) GetParentReference() ItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *RemoteItem) GetParentReferenceOk() (*ItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *RemoteItem) SetParentReference(v ItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *RemoteItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetShared

`func (o *RemoteItem) GetShared() Shared`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *RemoteItem) GetSharedOk() (*Shared, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *RemoteItem) SetShared(v Shared)`

SetShared sets Shared field to given value.

### HasShared

`func (o *RemoteItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *RemoteItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RemoteItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RemoteItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *RemoteItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSpecialFolder

`func (o *RemoteItem) GetSpecialFolder() SpecialFolder`

GetSpecialFolder returns the SpecialFolder field if non-nil, zero value otherwise.

### GetSpecialFolderOk

`func (o *RemoteItem) GetSpecialFolderOk() (*SpecialFolder, bool)`

GetSpecialFolderOk returns a tuple with the SpecialFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialFolder

`func (o *RemoteItem) SetSpecialFolder(v SpecialFolder)`

SetSpecialFolder sets SpecialFolder field to given value.

### HasSpecialFolder

`func (o *RemoteItem) HasSpecialFolder() bool`

HasSpecialFolder returns a boolean if a field has been set.

### GetWebDavUrl

`func (o *RemoteItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *RemoteItem) GetWebDavUrlOk() (*string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDavUrl

`func (o *RemoteItem) SetWebDavUrl(v string)`

SetWebDavUrl sets WebDavUrl field to given value.

### HasWebDavUrl

`func (o *RemoteItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### GetWebUrl

`func (o *RemoteItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *RemoteItem) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *RemoteItem) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *RemoteItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


