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

// checks if the GetVMConfig200ResponseDataUsb0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVMConfig200ResponseDataUsb0{}

// GetVMConfig200ResponseDataUsb0 struct for GetVMConfig200ResponseDataUsb0
type GetVMConfig200ResponseDataUsb0 struct {
	// The Host USB device or port or the value 'spice'. HOSTUSBDEVICE syntax is:   'bus-port(.port)*' (decimal numbers) or  'vendor_id:product_id' (hexadeciaml numbers) or  'spice'  You can use the 'lsusb -t' command to list existing usb devices.  NOTE: This option allows direct access to host hardware. So it is no longer possible to migrate such machines - use with special care.  The value 'spice' can be used to add a usb redirection devices for spice.  Either this or the 'mapping' key must be set. 
	Host *string `json:"host,omitempty"`
	// The ID of a cluster wide mapping. Either this or the default-key 'host' must be set.
	Mapping *string `json:"mapping,omitempty"`
	// Specifies whether if given host option is a USB3 device or port. For modern guests (machine version >= 7.1 and ostype l26 and windows > 7), this flag is irrelevant (all devices are plugged into a xhci controller).
	Usb3 *int32 `json:"usb3,omitempty"`
}

// NewGetVMConfig200ResponseDataUsb0 instantiates a new GetVMConfig200ResponseDataUsb0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVMConfig200ResponseDataUsb0() *GetVMConfig200ResponseDataUsb0 {
	this := GetVMConfig200ResponseDataUsb0{}
	return &this
}

// NewGetVMConfig200ResponseDataUsb0WithDefaults instantiates a new GetVMConfig200ResponseDataUsb0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVMConfig200ResponseDataUsb0WithDefaults() *GetVMConfig200ResponseDataUsb0 {
	this := GetVMConfig200ResponseDataUsb0{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataUsb0) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataUsb0) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataUsb0) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GetVMConfig200ResponseDataUsb0) SetHost(v string) {
	o.Host = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataUsb0) GetMapping() string {
	if o == nil || IsNil(o.Mapping) {
		var ret string
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataUsb0) GetMappingOk() (*string, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataUsb0) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given string and assigns it to the Mapping field.
func (o *GetVMConfig200ResponseDataUsb0) SetMapping(v string) {
	o.Mapping = &v
}

// GetUsb3 returns the Usb3 field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataUsb0) GetUsb3() int32 {
	if o == nil || IsNil(o.Usb3) {
		var ret int32
		return ret
	}
	return *o.Usb3
}

// GetUsb3Ok returns a tuple with the Usb3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataUsb0) GetUsb3Ok() (*int32, bool) {
	if o == nil || IsNil(o.Usb3) {
		return nil, false
	}
	return o.Usb3, true
}

// HasUsb3 returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataUsb0) HasUsb3() bool {
	if o != nil && !IsNil(o.Usb3) {
		return true
	}

	return false
}

// SetUsb3 gets a reference to the given int32 and assigns it to the Usb3 field.
func (o *GetVMConfig200ResponseDataUsb0) SetUsb3(v int32) {
	o.Usb3 = &v
}

func (o GetVMConfig200ResponseDataUsb0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVMConfig200ResponseDataUsb0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	if !IsNil(o.Usb3) {
		toSerialize["usb3"] = o.Usb3
	}
	return toSerialize, nil
}

type NullableGetVMConfig200ResponseDataUsb0 struct {
	value *GetVMConfig200ResponseDataUsb0
	isSet bool
}

func (v NullableGetVMConfig200ResponseDataUsb0) Get() *GetVMConfig200ResponseDataUsb0 {
	return v.value
}

func (v *NullableGetVMConfig200ResponseDataUsb0) Set(val *GetVMConfig200ResponseDataUsb0) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVMConfig200ResponseDataUsb0) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVMConfig200ResponseDataUsb0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVMConfig200ResponseDataUsb0(val *GetVMConfig200ResponseDataUsb0) *NullableGetVMConfig200ResponseDataUsb0 {
	return &NullableGetVMConfig200ResponseDataUsb0{value: val, isSet: true}
}

func (v NullableGetVMConfig200ResponseDataUsb0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVMConfig200ResponseDataUsb0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

