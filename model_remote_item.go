/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"time"
)

// checks if the RemoteItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteItem{}

// RemoteItem Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
type RemoteItem struct {
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time      `json:"createdDateTime,omitempty"`
	File            *OpenGraphFile  `json:"file,omitempty"`
	FileSystemInfo  *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	Folder          *Folder         `json:"folder,omitempty"`
	// The drive alias can be used in clients to make the urls user friendly. Example: 'personal/einstein'. This will be used to resolve to the correct driveID.
	DriveAlias *string `json:"driveAlias,omitempty"`
	// The relative path of the item in relation to its drive root.
	Path *string `json:"path,omitempty"`
	// Unique identifier for the drive root of this item. Read-only.
	RootId *string `json:"rootId,omitempty"`
	// Unique identifier for the remote item in its drive. Read-only.
	Id             *string      `json:"id,omitempty"`
	Image          *Image       `json:"image,omitempty"`
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Optional. Filename of the remote item. Read-only.
	Name *string `json:"name,omitempty"`
	// ETag for the item. Read-only.
	ETag *string `json:"eTag,omitempty"`
	// An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
	CTag            *string        `json:"cTag,omitempty"`
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	Shared          *Shared        `json:"shared,omitempty"`
	// The set of permissions for the item. Read-only. Nullable.
	Permissions []Permission `json:"permissions,omitempty"`
	// Size of the remote item. Read-only.
	Size          *int64         `json:"size,omitempty"`
	SpecialFolder *SpecialFolder `json:"specialFolder,omitempty"`
	// DAV compatible URL for the item.
	WebDavUrl *string `json:"webDavUrl,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
}

// NewRemoteItem instantiates a new RemoteItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteItem() *RemoteItem {
	this := RemoteItem{}
	return &this
}

// NewRemoteItemWithDefaults instantiates a new RemoteItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteItemWithDefaults() *RemoteItem {
	this := RemoteItem{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RemoteItem) GetCreatedBy() IdentitySet {
	if o == nil || IsNil(o.CreatedBy) {
		var ret IdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetCreatedByOk() (*IdentitySet, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RemoteItem) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given IdentitySet and assigns it to the CreatedBy field.
func (o *RemoteItem) SetCreatedBy(v IdentitySet) {
	o.CreatedBy = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *RemoteItem) GetCreatedDateTime() time.Time {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *RemoteItem) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *RemoteItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *RemoteItem) GetFile() OpenGraphFile {
	if o == nil || IsNil(o.File) {
		var ret OpenGraphFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetFileOk() (*OpenGraphFile, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *RemoteItem) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given OpenGraphFile and assigns it to the File field.
func (o *RemoteItem) SetFile(v OpenGraphFile) {
	o.File = &v
}

// GetFileSystemInfo returns the FileSystemInfo field value if set, zero value otherwise.
func (o *RemoteItem) GetFileSystemInfo() FileSystemInfo {
	if o == nil || IsNil(o.FileSystemInfo) {
		var ret FileSystemInfo
		return ret
	}
	return *o.FileSystemInfo
}

// GetFileSystemInfoOk returns a tuple with the FileSystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetFileSystemInfoOk() (*FileSystemInfo, bool) {
	if o == nil || IsNil(o.FileSystemInfo) {
		return nil, false
	}
	return o.FileSystemInfo, true
}

// HasFileSystemInfo returns a boolean if a field has been set.
func (o *RemoteItem) HasFileSystemInfo() bool {
	if o != nil && !IsNil(o.FileSystemInfo) {
		return true
	}

	return false
}

// SetFileSystemInfo gets a reference to the given FileSystemInfo and assigns it to the FileSystemInfo field.
func (o *RemoteItem) SetFileSystemInfo(v FileSystemInfo) {
	o.FileSystemInfo = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *RemoteItem) GetFolder() Folder {
	if o == nil || IsNil(o.Folder) {
		var ret Folder
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetFolderOk() (*Folder, bool) {
	if o == nil || IsNil(o.Folder) {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *RemoteItem) HasFolder() bool {
	if o != nil && !IsNil(o.Folder) {
		return true
	}

	return false
}

// SetFolder gets a reference to the given Folder and assigns it to the Folder field.
func (o *RemoteItem) SetFolder(v Folder) {
	o.Folder = &v
}

// GetDriveAlias returns the DriveAlias field value if set, zero value otherwise.
func (o *RemoteItem) GetDriveAlias() string {
	if o == nil || IsNil(o.DriveAlias) {
		var ret string
		return ret
	}
	return *o.DriveAlias
}

// GetDriveAliasOk returns a tuple with the DriveAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetDriveAliasOk() (*string, bool) {
	if o == nil || IsNil(o.DriveAlias) {
		return nil, false
	}
	return o.DriveAlias, true
}

// HasDriveAlias returns a boolean if a field has been set.
func (o *RemoteItem) HasDriveAlias() bool {
	if o != nil && !IsNil(o.DriveAlias) {
		return true
	}

	return false
}

// SetDriveAlias gets a reference to the given string and assigns it to the DriveAlias field.
func (o *RemoteItem) SetDriveAlias(v string) {
	o.DriveAlias = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RemoteItem) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RemoteItem) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RemoteItem) SetPath(v string) {
	o.Path = &v
}

// GetRootId returns the RootId field value if set, zero value otherwise.
func (o *RemoteItem) GetRootId() string {
	if o == nil || IsNil(o.RootId) {
		var ret string
		return ret
	}
	return *o.RootId
}

// GetRootIdOk returns a tuple with the RootId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetRootIdOk() (*string, bool) {
	if o == nil || IsNil(o.RootId) {
		return nil, false
	}
	return o.RootId, true
}

// HasRootId returns a boolean if a field has been set.
func (o *RemoteItem) HasRootId() bool {
	if o != nil && !IsNil(o.RootId) {
		return true
	}

	return false
}

// SetRootId gets a reference to the given string and assigns it to the RootId field.
func (o *RemoteItem) SetRootId(v string) {
	o.RootId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteItem) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *RemoteItem) GetImage() Image {
	if o == nil || IsNil(o.Image) {
		var ret Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetImageOk() (*Image, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *RemoteItem) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given Image and assigns it to the Image field.
func (o *RemoteItem) SetImage(v Image) {
	o.Image = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *RemoteItem) GetLastModifiedBy() IdentitySet {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret IdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetLastModifiedByOk() (*IdentitySet, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *RemoteItem) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given IdentitySet and assigns it to the LastModifiedBy field.
func (o *RemoteItem) SetLastModifiedBy(v IdentitySet) {
	o.LastModifiedBy = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *RemoteItem) GetLastModifiedDateTime() time.Time {
	if o == nil || IsNil(o.LastModifiedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDateTime) {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *RemoteItem) HasLastModifiedDateTime() bool {
	if o != nil && !IsNil(o.LastModifiedDateTime) {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *RemoteItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemoteItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemoteItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemoteItem) SetName(v string) {
	o.Name = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *RemoteItem) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *RemoteItem) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *RemoteItem) SetETag(v string) {
	o.ETag = &v
}

// GetCTag returns the CTag field value if set, zero value otherwise.
func (o *RemoteItem) GetCTag() string {
	if o == nil || IsNil(o.CTag) {
		var ret string
		return ret
	}
	return *o.CTag
}

// GetCTagOk returns a tuple with the CTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetCTagOk() (*string, bool) {
	if o == nil || IsNil(o.CTag) {
		return nil, false
	}
	return o.CTag, true
}

// HasCTag returns a boolean if a field has been set.
func (o *RemoteItem) HasCTag() bool {
	if o != nil && !IsNil(o.CTag) {
		return true
	}

	return false
}

// SetCTag gets a reference to the given string and assigns it to the CTag field.
func (o *RemoteItem) SetCTag(v string) {
	o.CTag = &v
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise.
func (o *RemoteItem) GetParentReference() ItemReference {
	if o == nil || IsNil(o.ParentReference) {
		var ret ItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetParentReferenceOk() (*ItemReference, bool) {
	if o == nil || IsNil(o.ParentReference) {
		return nil, false
	}
	return o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *RemoteItem) HasParentReference() bool {
	if o != nil && !IsNil(o.ParentReference) {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given ItemReference and assigns it to the ParentReference field.
func (o *RemoteItem) SetParentReference(v ItemReference) {
	o.ParentReference = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *RemoteItem) GetShared() Shared {
	if o == nil || IsNil(o.Shared) {
		var ret Shared
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetSharedOk() (*Shared, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *RemoteItem) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given Shared and assigns it to the Shared field.
func (o *RemoteItem) SetShared(v Shared) {
	o.Shared = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RemoteItem) GetPermissions() []Permission {
	if o == nil || IsNil(o.Permissions) {
		var ret []Permission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetPermissionsOk() ([]Permission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RemoteItem) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *RemoteItem) SetPermissions(v []Permission) {
	o.Permissions = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RemoteItem) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RemoteItem) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *RemoteItem) SetSize(v int64) {
	o.Size = &v
}

// GetSpecialFolder returns the SpecialFolder field value if set, zero value otherwise.
func (o *RemoteItem) GetSpecialFolder() SpecialFolder {
	if o == nil || IsNil(o.SpecialFolder) {
		var ret SpecialFolder
		return ret
	}
	return *o.SpecialFolder
}

// GetSpecialFolderOk returns a tuple with the SpecialFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetSpecialFolderOk() (*SpecialFolder, bool) {
	if o == nil || IsNil(o.SpecialFolder) {
		return nil, false
	}
	return o.SpecialFolder, true
}

// HasSpecialFolder returns a boolean if a field has been set.
func (o *RemoteItem) HasSpecialFolder() bool {
	if o != nil && !IsNil(o.SpecialFolder) {
		return true
	}

	return false
}

// SetSpecialFolder gets a reference to the given SpecialFolder and assigns it to the SpecialFolder field.
func (o *RemoteItem) SetSpecialFolder(v SpecialFolder) {
	o.SpecialFolder = &v
}

// GetWebDavUrl returns the WebDavUrl field value if set, zero value otherwise.
func (o *RemoteItem) GetWebDavUrl() string {
	if o == nil || IsNil(o.WebDavUrl) {
		var ret string
		return ret
	}
	return *o.WebDavUrl
}

// GetWebDavUrlOk returns a tuple with the WebDavUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetWebDavUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebDavUrl) {
		return nil, false
	}
	return o.WebDavUrl, true
}

// HasWebDavUrl returns a boolean if a field has been set.
func (o *RemoteItem) HasWebDavUrl() bool {
	if o != nil && !IsNil(o.WebDavUrl) {
		return true
	}

	return false
}

// SetWebDavUrl gets a reference to the given string and assigns it to the WebDavUrl field.
func (o *RemoteItem) SetWebDavUrl(v string) {
	o.WebDavUrl = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *RemoteItem) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *RemoteItem) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *RemoteItem) SetWebUrl(v string) {
	o.WebUrl = &v
}

func (o RemoteItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedDateTime) {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.FileSystemInfo) {
		toSerialize["fileSystemInfo"] = o.FileSystemInfo
	}
	if !IsNil(o.Folder) {
		toSerialize["folder"] = o.Folder
	}
	if !IsNil(o.DriveAlias) {
		toSerialize["driveAlias"] = o.DriveAlias
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.RootId) {
		toSerialize["rootId"] = o.RootId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !IsNil(o.LastModifiedDateTime) {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	if !IsNil(o.CTag) {
		toSerialize["cTag"] = o.CTag
	}
	if !IsNil(o.ParentReference) {
		toSerialize["parentReference"] = o.ParentReference
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.SpecialFolder) {
		toSerialize["specialFolder"] = o.SpecialFolder
	}
	if !IsNil(o.WebDavUrl) {
		toSerialize["webDavUrl"] = o.WebDavUrl
	}
	if !IsNil(o.WebUrl) {
		toSerialize["webUrl"] = o.WebUrl
	}
	return toSerialize, nil
}

type NullableRemoteItem struct {
	value *RemoteItem
	isSet bool
}

func (v NullableRemoteItem) Get() *RemoteItem {
	return v.value
}

func (v *NullableRemoteItem) Set(val *RemoteItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteItem(val *RemoteItem) *NullableRemoteItem {
	return &NullableRemoteItem{value: val, isSet: true}
}

func (v NullableRemoteItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
