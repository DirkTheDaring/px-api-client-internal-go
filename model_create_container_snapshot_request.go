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

// checks if the CreateContainerSnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerSnapshotRequest{}

// CreateContainerSnapshotRequest struct for CreateContainerSnapshotRequest
type CreateContainerSnapshotRequest struct {
	// A textual description or comment.
	Description *string `json:"description,omitempty"`
	// The name of the snapshot.
	Snapname string `json:"snapname"`
}

// NewCreateContainerSnapshotRequest instantiates a new CreateContainerSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerSnapshotRequest(snapname string) *CreateContainerSnapshotRequest {
	this := CreateContainerSnapshotRequest{}
	this.Snapname = snapname
	return &this
}

// NewCreateContainerSnapshotRequestWithDefaults instantiates a new CreateContainerSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerSnapshotRequestWithDefaults() *CreateContainerSnapshotRequest {
	this := CreateContainerSnapshotRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateContainerSnapshotRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerSnapshotRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateContainerSnapshotRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateContainerSnapshotRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSnapname returns the Snapname field value
func (o *CreateContainerSnapshotRequest) GetSnapname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Snapname
}

// GetSnapnameOk returns a tuple with the Snapname field value
// and a boolean to check if the value has been set.
func (o *CreateContainerSnapshotRequest) GetSnapnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snapname, true
}

// SetSnapname sets field value
func (o *CreateContainerSnapshotRequest) SetSnapname(v string) {
	o.Snapname = v
}

func (o CreateContainerSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainerSnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["snapname"] = o.Snapname
	return toSerialize, nil
}

type NullableCreateContainerSnapshotRequest struct {
	value *CreateContainerSnapshotRequest
	isSet bool
}

func (v NullableCreateContainerSnapshotRequest) Get() *CreateContainerSnapshotRequest {
	return v.value
}

func (v *NullableCreateContainerSnapshotRequest) Set(val *CreateContainerSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerSnapshotRequest(val *CreateContainerSnapshotRequest) *NullableCreateContainerSnapshotRequest {
	return &NullableCreateContainerSnapshotRequest{value: val, isSet: true}
}

func (v NullableCreateContainerSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


