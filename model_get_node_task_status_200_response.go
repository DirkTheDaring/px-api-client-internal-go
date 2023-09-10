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

// checks if the GetNodeTaskStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodeTaskStatus200Response{}

// GetNodeTaskStatus200Response struct for GetNodeTaskStatus200Response
type GetNodeTaskStatus200Response struct {
	Data *GetNodeTaskStatus200ResponseData `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewGetNodeTaskStatus200Response instantiates a new GetNodeTaskStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodeTaskStatus200Response() *GetNodeTaskStatus200Response {
	this := GetNodeTaskStatus200Response{}
	return &this
}

// NewGetNodeTaskStatus200ResponseWithDefaults instantiates a new GetNodeTaskStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodeTaskStatus200ResponseWithDefaults() *GetNodeTaskStatus200Response {
	this := GetNodeTaskStatus200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetNodeTaskStatus200Response) GetData() GetNodeTaskStatus200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GetNodeTaskStatus200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTaskStatus200Response) GetDataOk() (*GetNodeTaskStatus200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetNodeTaskStatus200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetNodeTaskStatus200ResponseData and assigns it to the Data field.
func (o *GetNodeTaskStatus200Response) SetData(v GetNodeTaskStatus200ResponseData) {
	o.Data = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetNodeTaskStatus200Response) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTaskStatus200Response) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetNodeTaskStatus200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GetNodeTaskStatus200Response) SetErrors(v []string) {
	o.Errors = v
}

func (o GetNodeTaskStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodeTaskStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetNodeTaskStatus200Response struct {
	value *GetNodeTaskStatus200Response
	isSet bool
}

func (v NullableGetNodeTaskStatus200Response) Get() *GetNodeTaskStatus200Response {
	return v.value
}

func (v *NullableGetNodeTaskStatus200Response) Set(val *GetNodeTaskStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodeTaskStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodeTaskStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodeTaskStatus200Response(val *GetNodeTaskStatus200Response) *NullableGetNodeTaskStatus200Response {
	return &NullableGetNodeTaskStatus200Response{value: val, isSet: true}
}

func (v NullableGetNodeTaskStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodeTaskStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

