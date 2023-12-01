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

// checks if the GetContainerConfig200ResponseDataDev0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContainerConfig200ResponseDataDev0{}

// GetContainerConfig200ResponseDataDev0 struct for GetContainerConfig200ResponseDataDev0
type GetContainerConfig200ResponseDataDev0 struct {
	// Group ID to be assigned to the device node
	Gid *int64 `json:"gid,omitempty"`
	// Access mode to be set on the device node
	Mode *string `json:"mode,omitempty"`
	// Device to pass through to the container
	Path *string `json:"path,omitempty"`
	// User ID to be assigned to the device node
	Uid *int64 `json:"uid,omitempty"`
}

// NewGetContainerConfig200ResponseDataDev0 instantiates a new GetContainerConfig200ResponseDataDev0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContainerConfig200ResponseDataDev0() *GetContainerConfig200ResponseDataDev0 {
	this := GetContainerConfig200ResponseDataDev0{}
	return &this
}

// NewGetContainerConfig200ResponseDataDev0WithDefaults instantiates a new GetContainerConfig200ResponseDataDev0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContainerConfig200ResponseDataDev0WithDefaults() *GetContainerConfig200ResponseDataDev0 {
	this := GetContainerConfig200ResponseDataDev0{}
	return &this
}

// GetGid returns the Gid field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataDev0) GetGid() int64 {
	if o == nil || IsNil(o.Gid) {
		var ret int64
		return ret
	}
	return *o.Gid
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataDev0) GetGidOk() (*int64, bool) {
	if o == nil || IsNil(o.Gid) {
		return nil, false
	}
	return o.Gid, true
}

// HasGid returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataDev0) HasGid() bool {
	if o != nil && !IsNil(o.Gid) {
		return true
	}

	return false
}

// SetGid gets a reference to the given int64 and assigns it to the Gid field.
func (o *GetContainerConfig200ResponseDataDev0) SetGid(v int64) {
	o.Gid = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataDev0) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataDev0) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataDev0) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GetContainerConfig200ResponseDataDev0) SetMode(v string) {
	o.Mode = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataDev0) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataDev0) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataDev0) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GetContainerConfig200ResponseDataDev0) SetPath(v string) {
	o.Path = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataDev0) GetUid() int64 {
	if o == nil || IsNil(o.Uid) {
		var ret int64
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataDev0) GetUidOk() (*int64, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataDev0) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int64 and assigns it to the Uid field.
func (o *GetContainerConfig200ResponseDataDev0) SetUid(v int64) {
	o.Uid = &v
}

func (o GetContainerConfig200ResponseDataDev0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContainerConfig200ResponseDataDev0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gid) {
		toSerialize["gid"] = o.Gid
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableGetContainerConfig200ResponseDataDev0 struct {
	value *GetContainerConfig200ResponseDataDev0
	isSet bool
}

func (v NullableGetContainerConfig200ResponseDataDev0) Get() *GetContainerConfig200ResponseDataDev0 {
	return v.value
}

func (v *NullableGetContainerConfig200ResponseDataDev0) Set(val *GetContainerConfig200ResponseDataDev0) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContainerConfig200ResponseDataDev0) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContainerConfig200ResponseDataDev0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContainerConfig200ResponseDataDev0(val *GetContainerConfig200ResponseDataDev0) *NullableGetContainerConfig200ResponseDataDev0 {
	return &NullableGetContainerConfig200ResponseDataDev0{value: val, isSet: true}
}

func (v NullableGetContainerConfig200ResponseDataDev0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContainerConfig200ResponseDataDev0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

