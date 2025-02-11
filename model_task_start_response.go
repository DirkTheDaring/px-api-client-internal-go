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

// checks if the TaskStartResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskStartResponse{}

// TaskStartResponse struct for TaskStartResponse
type TaskStartResponse struct {
	Errors []string `json:"errors,omitempty"`
	// the task ID.
	Data *string `json:"data,omitempty"`
}

// NewTaskStartResponse instantiates a new TaskStartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskStartResponse() *TaskStartResponse {
	this := TaskStartResponse{}
	return &this
}

// NewTaskStartResponseWithDefaults instantiates a new TaskStartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskStartResponseWithDefaults() *TaskStartResponse {
	this := TaskStartResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *TaskStartResponse) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStartResponse) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *TaskStartResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *TaskStartResponse) SetErrors(v []string) {
	o.Errors = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TaskStartResponse) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStartResponse) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TaskStartResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *TaskStartResponse) SetData(v string) {
	o.Data = &v
}

func (o TaskStartResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskStartResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTaskStartResponse struct {
	value *TaskStartResponse
	isSet bool
}

func (v NullableTaskStartResponse) Get() *TaskStartResponse {
	return v.value
}

func (v *NullableTaskStartResponse) Set(val *TaskStartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStartResponse(val *TaskStartResponse) *NullableTaskStartResponse {
	return &NullableTaskStartResponse{value: val, isSet: true}
}

func (v NullableTaskStartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


