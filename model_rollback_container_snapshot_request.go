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

// checks if the RollbackContainerSnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollbackContainerSnapshotRequest{}

// RollbackContainerSnapshotRequest struct for RollbackContainerSnapshotRequest
type RollbackContainerSnapshotRequest struct {
	// Whether the container should get started after rolling back successfully
	Start *bool `json:"start,omitempty"`
}

// NewRollbackContainerSnapshotRequest instantiates a new RollbackContainerSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackContainerSnapshotRequest() *RollbackContainerSnapshotRequest {
	this := RollbackContainerSnapshotRequest{}
	return &this
}

// NewRollbackContainerSnapshotRequestWithDefaults instantiates a new RollbackContainerSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackContainerSnapshotRequestWithDefaults() *RollbackContainerSnapshotRequest {
	this := RollbackContainerSnapshotRequest{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *RollbackContainerSnapshotRequest) GetStart() bool {
	if o == nil || IsNil(o.Start) {
		var ret bool
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackContainerSnapshotRequest) GetStartOk() (*bool, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *RollbackContainerSnapshotRequest) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given bool and assigns it to the Start field.
func (o *RollbackContainerSnapshotRequest) SetStart(v bool) {
	o.Start = &v
}

func (o RollbackContainerSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollbackContainerSnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return toSerialize, nil
}

type NullableRollbackContainerSnapshotRequest struct {
	value *RollbackContainerSnapshotRequest
	isSet bool
}

func (v NullableRollbackContainerSnapshotRequest) Get() *RollbackContainerSnapshotRequest {
	return v.value
}

func (v *NullableRollbackContainerSnapshotRequest) Set(val *RollbackContainerSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackContainerSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackContainerSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackContainerSnapshotRequest(val *RollbackContainerSnapshotRequest) *NullableRollbackContainerSnapshotRequest {
	return &NullableRollbackContainerSnapshotRequest{value: val, isSet: true}
}

func (v NullableRollbackContainerSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackContainerSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


