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

// checks if the CreateVMRequestEfidisk0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMRequestEfidisk0{}

// CreateVMRequestEfidisk0 struct for CreateVMRequestEfidisk0
type CreateVMRequestEfidisk0 struct {
	// Size and type of the OVMF EFI vars. '4m' is newer and recommended, and required for Secure Boot. For backwards compatibility, '2m' is used if not otherwise specified. Ignored for VMs with arch=aarc64 (ARM).
	Efitype *string `json:"efitype,omitempty"`
	// The drive's backing volume.
	File *string `json:"file,omitempty"`
	// The drive's backing file's data format.
	Format *string `json:"format,omitempty"`
	// Create a new disk, importing from this source (volume ID or absolute path). When an absolute path is specified, it's up to you to ensure that the source is not actively used by another process during the import!
	ImportFrom *string `json:"import-from,omitempty"`
	// Use am EFI vars template with distribution-specific and Microsoft Standard keys enrolled, if used with 'efitype=4m'. Note that this will enable Secure Boot by default, though it can still be turned off from within the VM.
	PreEnrolledKeys *int32 `json:"pre-enrolled-keys,omitempty"`
	// Disk size. This is purely informational and has no effect.
	Size *string `json:"size,omitempty"`
	Volume *string `json:"volume,omitempty"`
}

// NewCreateVMRequestEfidisk0 instantiates a new CreateVMRequestEfidisk0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMRequestEfidisk0() *CreateVMRequestEfidisk0 {
	this := CreateVMRequestEfidisk0{}
	return &this
}

// NewCreateVMRequestEfidisk0WithDefaults instantiates a new CreateVMRequestEfidisk0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMRequestEfidisk0WithDefaults() *CreateVMRequestEfidisk0 {
	this := CreateVMRequestEfidisk0{}
	return &this
}

// GetEfitype returns the Efitype field value if set, zero value otherwise.
func (o *CreateVMRequestEfidisk0) GetEfitype() string {
	if o == nil || IsNil(o.Efitype) {
		var ret string
		return ret
	}
	return *o.Efitype
}

// GetEfitypeOk returns a tuple with the Efitype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestEfidisk0) GetEfitypeOk() (*string, bool) {
	if o == nil || IsNil(o.Efitype) {
		return nil, false
	}
	return o.Efitype, true
}

// HasEfitype returns a boolean if a field has been set.
func (o *CreateVMRequestEfidisk0) HasEfitype() bool {
	if o != nil && !IsNil(o.Efitype) {
		return true
	}

	return false
}

// SetEfitype gets a reference to the given string and assigns it to the Efitype field.
func (o *CreateVMRequestEfidisk0) SetEfitype(v string) {
	o.Efitype = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *CreateVMRequestEfidisk0) GetFile() string {
	if o == nil || IsNil(o.File) {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestEfidisk0) GetFileOk() (*string, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *CreateVMRequestEfidisk0) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *CreateVMRequestEfidisk0) SetFile(v string) {
	o.File = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreateVMRequestEfidisk0) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestEfidisk0) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreateVMRequestEfidisk0) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreateVMRequestEfidisk0) SetFormat(v string) {
	o.Format = &v
}

// GetImportFrom returns the ImportFrom field value if set, zero value otherwise.
func (o *CreateVMRequestEfidisk0) GetImportFrom() string {
	if o == nil || IsNil(o.ImportFrom) {
		var ret string
		return ret
	}
	return *o.ImportFrom
}

// GetImportFromOk returns a tuple with the ImportFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestEfidisk0) GetImportFromOk() (*string, bool) {
	if o == nil || IsNil(o.ImportFrom) {
		return nil, false
	}
	return o.ImportFrom, true
}

// HasImportFrom returns a boolean if a field has been set.
func (o *CreateVMRequestEfidisk0) HasImportFrom() bool {
	if o != nil && !IsNil(o.ImportFrom) {
		return true
	}

	return false
}

// SetImportFrom gets a reference to the given string and assigns it to the ImportFrom field.
func (o *CreateVMRequestEfidisk0) SetImportFrom(v string) {
	o.ImportFrom = &v
}

// GetPreEnrolledKeys returns the PreEnrolledKeys field value if set, zero value otherwise.
func (o *CreateVMRequestEfidisk0) GetPreEnrolledKeys() int32 {
	if o == nil || IsNil(o.PreEnrolledKeys) {
		var ret int32
		return ret
	}
	return *o.PreEnrolledKeys
}

// GetPreEnrolledKeysOk returns a tuple with the PreEnrolledKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestEfidisk0) GetPreEnrolledKeysOk() (*int32, bool) {
	if o == nil || IsNil(o.PreEnrolledKeys) {
		return nil, false
	}
	return o.PreEnrolledKeys, true
}

// HasPreEnrolledKeys returns a boolean if a field has been set.
func (o *CreateVMRequestEfidisk0) HasPreEnrolledKeys() bool {
	if o != nil && !IsNil(o.PreEnrolledKeys) {
		return true
	}

	return false
}

// SetPreEnrolledKeys gets a reference to the given int32 and assigns it to the PreEnrolledKeys field.
func (o *CreateVMRequestEfidisk0) SetPreEnrolledKeys(v int32) {
	o.PreEnrolledKeys = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreateVMRequestEfidisk0) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestEfidisk0) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreateVMRequestEfidisk0) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *CreateVMRequestEfidisk0) SetSize(v string) {
	o.Size = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *CreateVMRequestEfidisk0) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestEfidisk0) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *CreateVMRequestEfidisk0) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *CreateVMRequestEfidisk0) SetVolume(v string) {
	o.Volume = &v
}

func (o CreateVMRequestEfidisk0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMRequestEfidisk0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Efitype) {
		toSerialize["efitype"] = o.Efitype
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.ImportFrom) {
		toSerialize["import-from"] = o.ImportFrom
	}
	if !IsNil(o.PreEnrolledKeys) {
		toSerialize["pre-enrolled-keys"] = o.PreEnrolledKeys
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableCreateVMRequestEfidisk0 struct {
	value *CreateVMRequestEfidisk0
	isSet bool
}

func (v NullableCreateVMRequestEfidisk0) Get() *CreateVMRequestEfidisk0 {
	return v.value
}

func (v *NullableCreateVMRequestEfidisk0) Set(val *CreateVMRequestEfidisk0) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMRequestEfidisk0) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMRequestEfidisk0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMRequestEfidisk0(val *CreateVMRequestEfidisk0) *NullableCreateVMRequestEfidisk0 {
	return &NullableCreateVMRequestEfidisk0{value: val, isSet: true}
}

func (v NullableCreateVMRequestEfidisk0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMRequestEfidisk0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


