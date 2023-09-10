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

// checks if the ShutdownContainerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShutdownContainerRequest{}

// ShutdownContainerRequest struct for ShutdownContainerRequest
type ShutdownContainerRequest struct {
	// Make sure the Container stops.
	ForceStop *int32 `json:"forceStop,omitempty"`
	// Wait maximal timeout seconds.
	Timeout *int64 `json:"timeout,omitempty"`
}

// NewShutdownContainerRequest instantiates a new ShutdownContainerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShutdownContainerRequest() *ShutdownContainerRequest {
	this := ShutdownContainerRequest{}
	return &this
}

// NewShutdownContainerRequestWithDefaults instantiates a new ShutdownContainerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShutdownContainerRequestWithDefaults() *ShutdownContainerRequest {
	this := ShutdownContainerRequest{}
	return &this
}

// GetForceStop returns the ForceStop field value if set, zero value otherwise.
func (o *ShutdownContainerRequest) GetForceStop() int32 {
	if o == nil || IsNil(o.ForceStop) {
		var ret int32
		return ret
	}
	return *o.ForceStop
}

// GetForceStopOk returns a tuple with the ForceStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShutdownContainerRequest) GetForceStopOk() (*int32, bool) {
	if o == nil || IsNil(o.ForceStop) {
		return nil, false
	}
	return o.ForceStop, true
}

// HasForceStop returns a boolean if a field has been set.
func (o *ShutdownContainerRequest) HasForceStop() bool {
	if o != nil && !IsNil(o.ForceStop) {
		return true
	}

	return false
}

// SetForceStop gets a reference to the given int32 and assigns it to the ForceStop field.
func (o *ShutdownContainerRequest) SetForceStop(v int32) {
	o.ForceStop = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ShutdownContainerRequest) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShutdownContainerRequest) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ShutdownContainerRequest) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ShutdownContainerRequest) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o ShutdownContainerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShutdownContainerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForceStop) {
		toSerialize["forceStop"] = o.ForceStop
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

type NullableShutdownContainerRequest struct {
	value *ShutdownContainerRequest
	isSet bool
}

func (v NullableShutdownContainerRequest) Get() *ShutdownContainerRequest {
	return v.value
}

func (v *NullableShutdownContainerRequest) Set(val *ShutdownContainerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableShutdownContainerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableShutdownContainerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShutdownContainerRequest(val *ShutdownContainerRequest) *NullableShutdownContainerRequest {
	return &NullableShutdownContainerRequest{value: val, isSet: true}
}

func (v NullableShutdownContainerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShutdownContainerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

