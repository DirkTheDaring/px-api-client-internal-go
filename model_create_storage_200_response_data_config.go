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

// checks if the CreateStorage200ResponseDataConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStorage200ResponseDataConfig{}

// CreateStorage200ResponseDataConfig Partial, possible server generated, configuration properties.
type CreateStorage200ResponseDataConfig struct {
	// The, possible auto-generated, encryption-key.
	EncryptionKey *string `json:"encryption-key,omitempty"`
}

// NewCreateStorage200ResponseDataConfig instantiates a new CreateStorage200ResponseDataConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStorage200ResponseDataConfig() *CreateStorage200ResponseDataConfig {
	this := CreateStorage200ResponseDataConfig{}
	return &this
}

// NewCreateStorage200ResponseDataConfigWithDefaults instantiates a new CreateStorage200ResponseDataConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStorage200ResponseDataConfigWithDefaults() *CreateStorage200ResponseDataConfig {
	this := CreateStorage200ResponseDataConfig{}
	return &this
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *CreateStorage200ResponseDataConfig) GetEncryptionKey() string {
	if o == nil || IsNil(o.EncryptionKey) {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage200ResponseDataConfig) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionKey) {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *CreateStorage200ResponseDataConfig) HasEncryptionKey() bool {
	if o != nil && !IsNil(o.EncryptionKey) {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *CreateStorage200ResponseDataConfig) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

func (o CreateStorage200ResponseDataConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStorage200ResponseDataConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EncryptionKey) {
		toSerialize["encryption-key"] = o.EncryptionKey
	}
	return toSerialize, nil
}

type NullableCreateStorage200ResponseDataConfig struct {
	value *CreateStorage200ResponseDataConfig
	isSet bool
}

func (v NullableCreateStorage200ResponseDataConfig) Get() *CreateStorage200ResponseDataConfig {
	return v.value
}

func (v *NullableCreateStorage200ResponseDataConfig) Set(val *CreateStorage200ResponseDataConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStorage200ResponseDataConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStorage200ResponseDataConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStorage200ResponseDataConfig(val *CreateStorage200ResponseDataConfig) *NullableCreateStorage200ResponseDataConfig {
	return &NullableCreateStorage200ResponseDataConfig{value: val, isSet: true}
}

func (v NullableCreateStorage200ResponseDataConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStorage200ResponseDataConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

