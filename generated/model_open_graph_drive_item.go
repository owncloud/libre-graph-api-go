/*
 * Open Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OpenGraphDriveItem struct for OpenGraphDriveItem
type OpenGraphDriveItem struct {
	OpenGraphBaseItem
	// The content stream, if the item represents a file.
	Content NullableString `json:"content,omitempty"`
	// An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
	CTag NullableString `json:"cTag,omitempty"`
	// Information about the deleted state of the item. Read-only.
	Deleted NullableAnyOfopenGraphDeleted `json:"deleted,omitempty"`
	// File metadata, if the item is a file. Read-only.
	File NullableAnyOfopenGraphFile `json:"file,omitempty"`
	// File system information on client. Read-write.
	FileSystemInfo NullableAnyOfopenGraphFileSystemInfo `json:"fileSystemInfo,omitempty"`
	// Folder metadata, if the item is a folder. Read-only.
	Folder NullableAnyOfopenGraphFolder `json:"folder,omitempty"`
	// Image metadata, if the item is an image. Read-only.
	Image NullableAnyOfopenGraphImage `json:"image,omitempty"`
	// If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
	Root NullableAnyOfobject `json:"root,omitempty"`
	// Size of the item in bytes. Read-only.
	Size NullableInt64 `json:"size,omitempty"`
	// WebDAV compatible URL for the item. Read-only.
	WebDavUrl NullableString `json:"webDavUrl,omitempty"`
	// Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
	Children *[]OpenGraphDriveItem `json:"children,omitempty"`
}

// NewOpenGraphDriveItem instantiates a new OpenGraphDriveItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphDriveItem() *OpenGraphDriveItem {
	this := OpenGraphDriveItem{}
	return &this
}

// NewOpenGraphDriveItemWithDefaults instantiates a new OpenGraphDriveItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphDriveItemWithDefaults() *OpenGraphDriveItem {
	this := OpenGraphDriveItem{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *OpenGraphDriveItem) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *OpenGraphDriveItem) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetContent() {
	o.Content.Unset()
}

// GetCTag returns the CTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetCTag() string {
	if o == nil || o.CTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.CTag.Get()
}

// GetCTagOk returns a tuple with the CTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetCTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CTag.Get(), o.CTag.IsSet()
}

// HasCTag returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasCTag() bool {
	if o != nil && o.CTag.IsSet() {
		return true
	}

	return false
}

// SetCTag gets a reference to the given NullableString and assigns it to the CTag field.
func (o *OpenGraphDriveItem) SetCTag(v string) {
	o.CTag.Set(&v)
}
// SetCTagNil sets the value for CTag to be an explicit nil
func (o *OpenGraphDriveItem) SetCTagNil() {
	o.CTag.Set(nil)
}

// UnsetCTag ensures that no value is present for CTag, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetCTag() {
	o.CTag.Unset()
}

// GetDeleted returns the Deleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetDeleted() AnyOfopenGraphDeleted {
	if o == nil || o.Deleted.Get() == nil {
		var ret AnyOfopenGraphDeleted
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetDeletedOk() (*AnyOfopenGraphDeleted, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// HasDeleted returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasDeleted() bool {
	if o != nil && o.Deleted.IsSet() {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given NullableAnyOfopenGraphDeleted and assigns it to the Deleted field.
func (o *OpenGraphDriveItem) SetDeleted(v AnyOfopenGraphDeleted) {
	o.Deleted.Set(&v)
}
// SetDeletedNil sets the value for Deleted to be an explicit nil
func (o *OpenGraphDriveItem) SetDeletedNil() {
	o.Deleted.Set(nil)
}

// UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetDeleted() {
	o.Deleted.Unset()
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetFile() AnyOfopenGraphFile {
	if o == nil || o.File.Get() == nil {
		var ret AnyOfopenGraphFile
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetFileOk() (*AnyOfopenGraphFile, bool) {
	if o == nil  {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableAnyOfopenGraphFile and assigns it to the File field.
func (o *OpenGraphDriveItem) SetFile(v AnyOfopenGraphFile) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *OpenGraphDriveItem) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetFile() {
	o.File.Unset()
}

// GetFileSystemInfo returns the FileSystemInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetFileSystemInfo() AnyOfopenGraphFileSystemInfo {
	if o == nil || o.FileSystemInfo.Get() == nil {
		var ret AnyOfopenGraphFileSystemInfo
		return ret
	}
	return *o.FileSystemInfo.Get()
}

// GetFileSystemInfoOk returns a tuple with the FileSystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetFileSystemInfoOk() (*AnyOfopenGraphFileSystemInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileSystemInfo.Get(), o.FileSystemInfo.IsSet()
}

// HasFileSystemInfo returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasFileSystemInfo() bool {
	if o != nil && o.FileSystemInfo.IsSet() {
		return true
	}

	return false
}

// SetFileSystemInfo gets a reference to the given NullableAnyOfopenGraphFileSystemInfo and assigns it to the FileSystemInfo field.
func (o *OpenGraphDriveItem) SetFileSystemInfo(v AnyOfopenGraphFileSystemInfo) {
	o.FileSystemInfo.Set(&v)
}
// SetFileSystemInfoNil sets the value for FileSystemInfo to be an explicit nil
func (o *OpenGraphDriveItem) SetFileSystemInfoNil() {
	o.FileSystemInfo.Set(nil)
}

// UnsetFileSystemInfo ensures that no value is present for FileSystemInfo, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetFileSystemInfo() {
	o.FileSystemInfo.Unset()
}

// GetFolder returns the Folder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetFolder() AnyOfopenGraphFolder {
	if o == nil || o.Folder.Get() == nil {
		var ret AnyOfopenGraphFolder
		return ret
	}
	return *o.Folder.Get()
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetFolderOk() (*AnyOfopenGraphFolder, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Folder.Get(), o.Folder.IsSet()
}

// HasFolder returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasFolder() bool {
	if o != nil && o.Folder.IsSet() {
		return true
	}

	return false
}

// SetFolder gets a reference to the given NullableAnyOfopenGraphFolder and assigns it to the Folder field.
func (o *OpenGraphDriveItem) SetFolder(v AnyOfopenGraphFolder) {
	o.Folder.Set(&v)
}
// SetFolderNil sets the value for Folder to be an explicit nil
func (o *OpenGraphDriveItem) SetFolderNil() {
	o.Folder.Set(nil)
}

// UnsetFolder ensures that no value is present for Folder, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetFolder() {
	o.Folder.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetImage() AnyOfopenGraphImage {
	if o == nil || o.Image.Get() == nil {
		var ret AnyOfopenGraphImage
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetImageOk() (*AnyOfopenGraphImage, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableAnyOfopenGraphImage and assigns it to the Image field.
func (o *OpenGraphDriveItem) SetImage(v AnyOfopenGraphImage) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *OpenGraphDriveItem) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetImage() {
	o.Image.Unset()
}

// GetRoot returns the Root field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetRoot() AnyOfobject {
	if o == nil || o.Root.Get() == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Root.Get()
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetRootOk() (*AnyOfobject, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Root.Get(), o.Root.IsSet()
}

// HasRoot returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasRoot() bool {
	if o != nil && o.Root.IsSet() {
		return true
	}

	return false
}

// SetRoot gets a reference to the given NullableAnyOfobject and assigns it to the Root field.
func (o *OpenGraphDriveItem) SetRoot(v AnyOfobject) {
	o.Root.Set(&v)
}
// SetRootNil sets the value for Root to be an explicit nil
func (o *OpenGraphDriveItem) SetRootNil() {
	o.Root.Set(nil)
}

// UnsetRoot ensures that no value is present for Root, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetRoot() {
	o.Root.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetSize() int64 {
	if o == nil || o.Size.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt64 and assigns it to the Size field.
func (o *OpenGraphDriveItem) SetSize(v int64) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *OpenGraphDriveItem) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetSize() {
	o.Size.Unset()
}

// GetWebDavUrl returns the WebDavUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenGraphDriveItem) GetWebDavUrl() string {
	if o == nil || o.WebDavUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebDavUrl.Get()
}

// GetWebDavUrlOk returns a tuple with the WebDavUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenGraphDriveItem) GetWebDavUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebDavUrl.Get(), o.WebDavUrl.IsSet()
}

// HasWebDavUrl returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasWebDavUrl() bool {
	if o != nil && o.WebDavUrl.IsSet() {
		return true
	}

	return false
}

// SetWebDavUrl gets a reference to the given NullableString and assigns it to the WebDavUrl field.
func (o *OpenGraphDriveItem) SetWebDavUrl(v string) {
	o.WebDavUrl.Set(&v)
}
// SetWebDavUrlNil sets the value for WebDavUrl to be an explicit nil
func (o *OpenGraphDriveItem) SetWebDavUrlNil() {
	o.WebDavUrl.Set(nil)
}

// UnsetWebDavUrl ensures that no value is present for WebDavUrl, not even an explicit nil
func (o *OpenGraphDriveItem) UnsetWebDavUrl() {
	o.WebDavUrl.Unset()
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *OpenGraphDriveItem) GetChildren() []OpenGraphDriveItem {
	if o == nil || o.Children == nil {
		var ret []OpenGraphDriveItem
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphDriveItem) GetChildrenOk() (*[]OpenGraphDriveItem, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *OpenGraphDriveItem) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []OpenGraphDriveItem and assigns it to the Children field.
func (o *OpenGraphDriveItem) SetChildren(v []OpenGraphDriveItem) {
	o.Children = &v
}

func (o OpenGraphDriveItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOpenGraphBaseItem, errOpenGraphBaseItem := json.Marshal(o.OpenGraphBaseItem)
	if errOpenGraphBaseItem != nil {
		return []byte{}, errOpenGraphBaseItem
	}
	errOpenGraphBaseItem = json.Unmarshal([]byte(serializedOpenGraphBaseItem), &toSerialize)
	if errOpenGraphBaseItem != nil {
		return []byte{}, errOpenGraphBaseItem
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.CTag.IsSet() {
		toSerialize["cTag"] = o.CTag.Get()
	}
	if o.Deleted.IsSet() {
		toSerialize["deleted"] = o.Deleted.Get()
	}
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}
	if o.FileSystemInfo.IsSet() {
		toSerialize["fileSystemInfo"] = o.FileSystemInfo.Get()
	}
	if o.Folder.IsSet() {
		toSerialize["folder"] = o.Folder.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.Root.IsSet() {
		toSerialize["root"] = o.Root.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.WebDavUrl.IsSet() {
		toSerialize["webDavUrl"] = o.WebDavUrl.Get()
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphDriveItem struct {
	value *OpenGraphDriveItem
	isSet bool
}

func (v NullableOpenGraphDriveItem) Get() *OpenGraphDriveItem {
	return v.value
}

func (v *NullableOpenGraphDriveItem) Set(val *OpenGraphDriveItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphDriveItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphDriveItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphDriveItem(val *OpenGraphDriveItem) *NullableOpenGraphDriveItem {
	return &NullableOpenGraphDriveItem{value: val, isSet: true}
}

func (v NullableOpenGraphDriveItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphDriveItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


