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

// checks if the CreateVMRequestIpconfig0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMRequestIpconfig0{}

// CreateVMRequestIpconfig0 struct for CreateVMRequestIpconfig0
type CreateVMRequestIpconfig0 struct {
	Gw *string `json:"gw,omitempty"`
	Gw6 *string `json:"gw6,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Ip6 *string `json:"ip6,omitempty"`
}

// NewCreateVMRequestIpconfig0 instantiates a new CreateVMRequestIpconfig0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMRequestIpconfig0() *CreateVMRequestIpconfig0 {
	this := CreateVMRequestIpconfig0{}
	return &this
}

// NewCreateVMRequestIpconfig0WithDefaults instantiates a new CreateVMRequestIpconfig0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMRequestIpconfig0WithDefaults() *CreateVMRequestIpconfig0 {
	this := CreateVMRequestIpconfig0{}
	return &this
}

// GetGw returns the Gw field value if set, zero value otherwise.
func (o *CreateVMRequestIpconfig0) GetGw() string {
	if o == nil || IsNil(o.Gw) {
		var ret string
		return ret
	}
	return *o.Gw
}

// GetGwOk returns a tuple with the Gw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestIpconfig0) GetGwOk() (*string, bool) {
	if o == nil || IsNil(o.Gw) {
		return nil, false
	}
	return o.Gw, true
}

// HasGw returns a boolean if a field has been set.
func (o *CreateVMRequestIpconfig0) HasGw() bool {
	if o != nil && !IsNil(o.Gw) {
		return true
	}

	return false
}

// SetGw gets a reference to the given string and assigns it to the Gw field.
func (o *CreateVMRequestIpconfig0) SetGw(v string) {
	o.Gw = &v
}

// GetGw6 returns the Gw6 field value if set, zero value otherwise.
func (o *CreateVMRequestIpconfig0) GetGw6() string {
	if o == nil || IsNil(o.Gw6) {
		var ret string
		return ret
	}
	return *o.Gw6
}

// GetGw6Ok returns a tuple with the Gw6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestIpconfig0) GetGw6Ok() (*string, bool) {
	if o == nil || IsNil(o.Gw6) {
		return nil, false
	}
	return o.Gw6, true
}

// HasGw6 returns a boolean if a field has been set.
func (o *CreateVMRequestIpconfig0) HasGw6() bool {
	if o != nil && !IsNil(o.Gw6) {
		return true
	}

	return false
}

// SetGw6 gets a reference to the given string and assigns it to the Gw6 field.
func (o *CreateVMRequestIpconfig0) SetGw6(v string) {
	o.Gw6 = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CreateVMRequestIpconfig0) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestIpconfig0) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CreateVMRequestIpconfig0) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CreateVMRequestIpconfig0) SetIp(v string) {
	o.Ip = &v
}

// GetIp6 returns the Ip6 field value if set, zero value otherwise.
func (o *CreateVMRequestIpconfig0) GetIp6() string {
	if o == nil || IsNil(o.Ip6) {
		var ret string
		return ret
	}
	return *o.Ip6
}

// GetIp6Ok returns a tuple with the Ip6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMRequestIpconfig0) GetIp6Ok() (*string, bool) {
	if o == nil || IsNil(o.Ip6) {
		return nil, false
	}
	return o.Ip6, true
}

// HasIp6 returns a boolean if a field has been set.
func (o *CreateVMRequestIpconfig0) HasIp6() bool {
	if o != nil && !IsNil(o.Ip6) {
		return true
	}

	return false
}

// SetIp6 gets a reference to the given string and assigns it to the Ip6 field.
func (o *CreateVMRequestIpconfig0) SetIp6(v string) {
	o.Ip6 = &v
}

func (o CreateVMRequestIpconfig0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMRequestIpconfig0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gw) {
		toSerialize["gw"] = o.Gw
	}
	if !IsNil(o.Gw6) {
		toSerialize["gw6"] = o.Gw6
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Ip6) {
		toSerialize["ip6"] = o.Ip6
	}
	return toSerialize, nil
}

type NullableCreateVMRequestIpconfig0 struct {
	value *CreateVMRequestIpconfig0
	isSet bool
}

func (v NullableCreateVMRequestIpconfig0) Get() *CreateVMRequestIpconfig0 {
	return v.value
}

func (v *NullableCreateVMRequestIpconfig0) Set(val *CreateVMRequestIpconfig0) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMRequestIpconfig0) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMRequestIpconfig0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMRequestIpconfig0(val *CreateVMRequestIpconfig0) *NullableCreateVMRequestIpconfig0 {
	return &NullableCreateVMRequestIpconfig0{value: val, isSet: true}
}

func (v NullableCreateVMRequestIpconfig0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMRequestIpconfig0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


