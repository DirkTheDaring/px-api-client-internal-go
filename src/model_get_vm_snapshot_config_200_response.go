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

// checks if the GetVMSnapshotConfig200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVMSnapshotConfig200Response{}

// GetVMSnapshotConfig200Response struct for GetVMSnapshotConfig200Response
type GetVMSnapshotConfig200Response struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewGetVMSnapshotConfig200Response instantiates a new GetVMSnapshotConfig200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVMSnapshotConfig200Response() *GetVMSnapshotConfig200Response {
	this := GetVMSnapshotConfig200Response{}
	return &this
}

// NewGetVMSnapshotConfig200ResponseWithDefaults instantiates a new GetVMSnapshotConfig200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVMSnapshotConfig200ResponseWithDefaults() *GetVMSnapshotConfig200Response {
	this := GetVMSnapshotConfig200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetVMSnapshotConfig200Response) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMSnapshotConfig200Response) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetVMSnapshotConfig200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *GetVMSnapshotConfig200Response) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetVMSnapshotConfig200Response) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMSnapshotConfig200Response) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetVMSnapshotConfig200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GetVMSnapshotConfig200Response) SetErrors(v []string) {
	o.Errors = v
}

func (o GetVMSnapshotConfig200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVMSnapshotConfig200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetVMSnapshotConfig200Response struct {
	value *GetVMSnapshotConfig200Response
	isSet bool
}

func (v NullableGetVMSnapshotConfig200Response) Get() *GetVMSnapshotConfig200Response {
	return v.value
}

func (v *NullableGetVMSnapshotConfig200Response) Set(val *GetVMSnapshotConfig200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVMSnapshotConfig200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVMSnapshotConfig200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVMSnapshotConfig200Response(val *GetVMSnapshotConfig200Response) *NullableGetVMSnapshotConfig200Response {
	return &NullableGetVMSnapshotConfig200Response{value: val, isSet: true}
}

func (v NullableGetVMSnapshotConfig200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVMSnapshotConfig200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


