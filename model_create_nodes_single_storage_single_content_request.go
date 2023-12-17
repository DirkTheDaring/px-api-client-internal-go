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
	"os"
	"fmt"
)

// checks if the CreateNodesSingleStorageSingleContentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNodesSingleStorageSingleContentRequest{}

// CreateNodesSingleStorageSingleContentRequest struct for CreateNodesSingleStorageSingleContentRequest
type CreateNodesSingleStorageSingleContentRequest struct {
	// The name of the file to create.
	Filename *os.File `json:"filename"`
	Format *string `json:"format,omitempty"`
	// Size in kilobyte (1024 bytes). Optional suffixes 'M' (megabyte, 1024K) and 'G' (gigabyte, 1024M)
	Size string `json:"size"`
	// Specify owner VM
	Vmid int64 `json:"vmid"`
}

type _CreateNodesSingleStorageSingleContentRequest CreateNodesSingleStorageSingleContentRequest

// NewCreateNodesSingleStorageSingleContentRequest instantiates a new CreateNodesSingleStorageSingleContentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNodesSingleStorageSingleContentRequest(filename *os.File, size string, vmid int64) *CreateNodesSingleStorageSingleContentRequest {
	this := CreateNodesSingleStorageSingleContentRequest{}
	this.Filename = filename
	this.Size = size
	this.Vmid = vmid
	return &this
}

// NewCreateNodesSingleStorageSingleContentRequestWithDefaults instantiates a new CreateNodesSingleStorageSingleContentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNodesSingleStorageSingleContentRequestWithDefaults() *CreateNodesSingleStorageSingleContentRequest {
	this := CreateNodesSingleStorageSingleContentRequest{}
	return &this
}

// GetFilename returns the Filename field value
func (o *CreateNodesSingleStorageSingleContentRequest) GetFilename() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *CreateNodesSingleStorageSingleContentRequest) GetFilenameOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *CreateNodesSingleStorageSingleContentRequest) SetFilename(v *os.File) {
	o.Filename = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreateNodesSingleStorageSingleContentRequest) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNodesSingleStorageSingleContentRequest) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreateNodesSingleStorageSingleContentRequest) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreateNodesSingleStorageSingleContentRequest) SetFormat(v string) {
	o.Format = &v
}

// GetSize returns the Size field value
func (o *CreateNodesSingleStorageSingleContentRequest) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CreateNodesSingleStorageSingleContentRequest) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CreateNodesSingleStorageSingleContentRequest) SetSize(v string) {
	o.Size = v
}

// GetVmid returns the Vmid field value
func (o *CreateNodesSingleStorageSingleContentRequest) GetVmid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Vmid
}

// GetVmidOk returns a tuple with the Vmid field value
// and a boolean to check if the value has been set.
func (o *CreateNodesSingleStorageSingleContentRequest) GetVmidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vmid, true
}

// SetVmid sets field value
func (o *CreateNodesSingleStorageSingleContentRequest) SetVmid(v int64) {
	o.Vmid = v
}

func (o CreateNodesSingleStorageSingleContentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNodesSingleStorageSingleContentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filename"] = o.Filename
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	toSerialize["size"] = o.Size
	toSerialize["vmid"] = o.Vmid
	return toSerialize, nil
}

func (o *CreateNodesSingleStorageSingleContentRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filename",
		"size",
		"vmid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNodesSingleStorageSingleContentRequest := _CreateNodesSingleStorageSingleContentRequest{}

	err = json.Unmarshal(bytes, &varCreateNodesSingleStorageSingleContentRequest)

	if err != nil {
		return err
	}

	*o = CreateNodesSingleStorageSingleContentRequest(varCreateNodesSingleStorageSingleContentRequest)

	return err
}

type NullableCreateNodesSingleStorageSingleContentRequest struct {
	value *CreateNodesSingleStorageSingleContentRequest
	isSet bool
}

func (v NullableCreateNodesSingleStorageSingleContentRequest) Get() *CreateNodesSingleStorageSingleContentRequest {
	return v.value
}

func (v *NullableCreateNodesSingleStorageSingleContentRequest) Set(val *CreateNodesSingleStorageSingleContentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNodesSingleStorageSingleContentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNodesSingleStorageSingleContentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNodesSingleStorageSingleContentRequest(val *CreateNodesSingleStorageSingleContentRequest) *NullableCreateNodesSingleStorageSingleContentRequest {
	return &NullableCreateNodesSingleStorageSingleContentRequest{value: val, isSet: true}
}

func (v NullableCreateNodesSingleStorageSingleContentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNodesSingleStorageSingleContentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


