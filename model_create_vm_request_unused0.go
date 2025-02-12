/*
ProxMox VE API

ProxMox VE API

API version: 8.3
Contact: baldur@email.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pxapiobject

import (
	"encoding/json"
)

// checks if the CreateVMRequestUnused0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMRequestUnused0{}

// CreateVMRequestUnused0 struct for CreateVMRequestUnused0
type CreateVMRequestUnused0 struct {
	File string `json:"file"`
}

// NewCreateVMRequestUnused0 instantiates a new CreateVMRequestUnused0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMRequestUnused0(file string) *CreateVMRequestUnused0 {
	this := CreateVMRequestUnused0{}
	this.File = file
	return &this
}

// NewCreateVMRequestUnused0WithDefaults instantiates a new CreateVMRequestUnused0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMRequestUnused0WithDefaults() *CreateVMRequestUnused0 {
	this := CreateVMRequestUnused0{}
	return &this
}

// GetFile returns the File field value
func (o *CreateVMRequestUnused0) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *CreateVMRequestUnused0) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *CreateVMRequestUnused0) SetFile(v string) {
	o.File = v
}

func (o CreateVMRequestUnused0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMRequestUnused0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	return toSerialize, nil
}

type NullableCreateVMRequestUnused0 struct {
	value *CreateVMRequestUnused0
	isSet bool
}

func (v NullableCreateVMRequestUnused0) Get() *CreateVMRequestUnused0 {
	return v.value
}

func (v *NullableCreateVMRequestUnused0) Set(val *CreateVMRequestUnused0) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMRequestUnused0) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMRequestUnused0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMRequestUnused0(val *CreateVMRequestUnused0) *NullableCreateVMRequestUnused0 {
	return &NullableCreateVMRequestUnused0{value: val, isSet: true}
}

func (v NullableCreateVMRequestUnused0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMRequestUnused0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


