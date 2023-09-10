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

// checks if the GetVMSnapshot200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVMSnapshot200Response{}

// GetVMSnapshot200Response struct for GetVMSnapshot200Response
type GetVMSnapshot200Response struct {
	Data []map[string]interface{} `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewGetVMSnapshot200Response instantiates a new GetVMSnapshot200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVMSnapshot200Response() *GetVMSnapshot200Response {
	this := GetVMSnapshot200Response{}
	return &this
}

// NewGetVMSnapshot200ResponseWithDefaults instantiates a new GetVMSnapshot200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVMSnapshot200ResponseWithDefaults() *GetVMSnapshot200Response {
	this := GetVMSnapshot200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetVMSnapshot200Response) GetData() []map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMSnapshot200Response) GetDataOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetVMSnapshot200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []map[string]interface{} and assigns it to the Data field.
func (o *GetVMSnapshot200Response) SetData(v []map[string]interface{}) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetVMSnapshot200Response) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMSnapshot200Response) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetVMSnapshot200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GetVMSnapshot200Response) SetErrors(v []string) {
	o.Errors = v
}

func (o GetVMSnapshot200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVMSnapshot200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetVMSnapshot200Response struct {
	value *GetVMSnapshot200Response
	isSet bool
}

func (v NullableGetVMSnapshot200Response) Get() *GetVMSnapshot200Response {
	return v.value
}

func (v *NullableGetVMSnapshot200Response) Set(val *GetVMSnapshot200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVMSnapshot200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVMSnapshot200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVMSnapshot200Response(val *GetVMSnapshot200Response) *NullableGetVMSnapshot200Response {
	return &NullableGetVMSnapshot200Response{value: val, isSet: true}
}

func (v NullableGetVMSnapshot200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVMSnapshot200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

