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

// checks if the GetCurrentContainerStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCurrentContainerStatus200Response{}

// GetCurrentContainerStatus200Response struct for GetCurrentContainerStatus200Response
type GetCurrentContainerStatus200Response struct {
	Data *GetCurrentContainerStatus200ResponseData `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewGetCurrentContainerStatus200Response instantiates a new GetCurrentContainerStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrentContainerStatus200Response() *GetCurrentContainerStatus200Response {
	this := GetCurrentContainerStatus200Response{}
	return &this
}

// NewGetCurrentContainerStatus200ResponseWithDefaults instantiates a new GetCurrentContainerStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrentContainerStatus200ResponseWithDefaults() *GetCurrentContainerStatus200Response {
	this := GetCurrentContainerStatus200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCurrentContainerStatus200Response) GetData() GetCurrentContainerStatus200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GetCurrentContainerStatus200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentContainerStatus200Response) GetDataOk() (*GetCurrentContainerStatus200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCurrentContainerStatus200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetCurrentContainerStatus200ResponseData and assigns it to the Data field.
func (o *GetCurrentContainerStatus200Response) SetData(v GetCurrentContainerStatus200ResponseData) {
	o.Data = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetCurrentContainerStatus200Response) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentContainerStatus200Response) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetCurrentContainerStatus200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GetCurrentContainerStatus200Response) SetErrors(v []string) {
	o.Errors = v
}

func (o GetCurrentContainerStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrentContainerStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetCurrentContainerStatus200Response struct {
	value *GetCurrentContainerStatus200Response
	isSet bool
}

func (v NullableGetCurrentContainerStatus200Response) Get() *GetCurrentContainerStatus200Response {
	return v.value
}

func (v *NullableGetCurrentContainerStatus200Response) Set(val *GetCurrentContainerStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrentContainerStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrentContainerStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrentContainerStatus200Response(val *GetCurrentContainerStatus200Response) *NullableGetCurrentContainerStatus200Response {
	return &NullableGetCurrentContainerStatus200Response{value: val, isSet: true}
}

func (v NullableGetCurrentContainerStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrentContainerStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

