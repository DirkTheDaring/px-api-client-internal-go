/*
ProxMox VE API

ProxMox VE API

API version: 8.0
Contact: baldur@email.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pxapiobject

import (
	"encoding/json"
)

// checks if the CreateVMRequestTpmstate0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMRequestTpmstate0{}

// CreateVMRequestTpmstate0 struct for CreateVMRequestTpmstate0
type CreateVMRequestTpmstate0 struct {
	// The drive's backing volume.
	File *string `json:"file,omitempty"`
	// Create a new disk, importing from this source (volume ID or absolute path). When an absolute path is specified, it's up to you to ensure that the source is not actively used by another process during the import!
	ImportFrom *string `json:"import-from,omitempty"`
	// Disk size. This is purely informational and has no effect.
	Size *string `json:"size,omitempty"`
	// The TPM interface version. v2.0 is newer and should be preferred. Note that this cannot be changed later on.
	Version *string `json:"version,omitempty"`
	Volume *string `json:"volume,omitempty"`
}

// NewCreateVMRequestTpmstate0 instantiates a new CreateVMRequestTpmstate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMRequestTpmstate0() *CreateVMRequestTpmstate0 {
	this := CreateVMRequestTpmstate0{}
	return &this
}

// NewCreateVMRequestTpmstate0WithDefaults instantiates a new CreateVMRequestTpmstate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMRequestTpmstate0WithDefaults() *CreateVMRequestTpmstate0 {
	this := CreateVMRequestTpmstate0{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *CreateVMRequestTpmstate0) GetFile() string {
	if o == nil || IsNil(o.File) {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestTpmstate0) GetFileOk() (*string, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *CreateVMRequestTpmstate0) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *CreateVMRequestTpmstate0) SetFile(v string) {
	o.File = &v
}

// GetImportFrom returns the ImportFrom field value if set, zero value otherwise.
func (o *CreateVMRequestTpmstate0) GetImportFrom() string {
	if o == nil || IsNil(o.ImportFrom) {
		var ret string
		return ret
	}
	return *o.ImportFrom
}

// GetImportFromOk returns a tuple with the ImportFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestTpmstate0) GetImportFromOk() (*string, bool) {
	if o == nil || IsNil(o.ImportFrom) {
		return nil, false
	}
	return o.ImportFrom, true
}

// HasImportFrom returns a boolean if a field has been set.
func (o *CreateVMRequestTpmstate0) HasImportFrom() bool {
	if o != nil && !IsNil(o.ImportFrom) {
		return true
	}

	return false
}

// SetImportFrom gets a reference to the given string and assigns it to the ImportFrom field.
func (o *CreateVMRequestTpmstate0) SetImportFrom(v string) {
	o.ImportFrom = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreateVMRequestTpmstate0) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestTpmstate0) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreateVMRequestTpmstate0) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *CreateVMRequestTpmstate0) SetSize(v string) {
	o.Size = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateVMRequestTpmstate0) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestTpmstate0) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateVMRequestTpmstate0) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CreateVMRequestTpmstate0) SetVersion(v string) {
	o.Version = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *CreateVMRequestTpmstate0) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestTpmstate0) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *CreateVMRequestTpmstate0) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *CreateVMRequestTpmstate0) SetVolume(v string) {
	o.Volume = &v
}

func (o CreateVMRequestTpmstate0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMRequestTpmstate0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.ImportFrom) {
		toSerialize["import-from"] = o.ImportFrom
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableCreateVMRequestTpmstate0 struct {
	value *CreateVMRequestTpmstate0
	isSet bool
}

func (v NullableCreateVMRequestTpmstate0) Get() *CreateVMRequestTpmstate0 {
	return v.value
}

func (v *NullableCreateVMRequestTpmstate0) Set(val *CreateVMRequestTpmstate0) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMRequestTpmstate0) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMRequestTpmstate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMRequestTpmstate0(val *CreateVMRequestTpmstate0) *NullableCreateVMRequestTpmstate0 {
	return &NullableCreateVMRequestTpmstate0{value: val, isSet: true}
}

func (v NullableCreateVMRequestTpmstate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMRequestTpmstate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

