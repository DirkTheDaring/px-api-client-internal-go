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

// checks if the GetClusterResources200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetClusterResources200Response{}

// GetClusterResources200Response struct for GetClusterResources200Response
type GetClusterResources200Response struct {
	Data []GetClusterResources200ResponseDataInner `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewGetClusterResources200Response instantiates a new GetClusterResources200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClusterResources200Response() *GetClusterResources200Response {
	this := GetClusterResources200Response{}
	return &this
}

// NewGetClusterResources200ResponseWithDefaults instantiates a new GetClusterResources200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClusterResources200ResponseWithDefaults() *GetClusterResources200Response {
	this := GetClusterResources200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetClusterResources200Response) GetData() []GetClusterResources200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetClusterResources200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterResources200Response) GetDataOk() ([]GetClusterResources200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetClusterResources200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetClusterResources200ResponseDataInner and assigns it to the Data field.
func (o *GetClusterResources200Response) SetData(v []GetClusterResources200ResponseDataInner) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetClusterResources200Response) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterResources200Response) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetClusterResources200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GetClusterResources200Response) SetErrors(v []string) {
	o.Errors = v
}

func (o GetClusterResources200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetClusterResources200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetClusterResources200Response struct {
	value *GetClusterResources200Response
	isSet bool
}

func (v NullableGetClusterResources200Response) Get() *GetClusterResources200Response {
	return v.value
}

func (v *NullableGetClusterResources200Response) Set(val *GetClusterResources200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClusterResources200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClusterResources200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClusterResources200Response(val *GetClusterResources200Response) *NullableGetClusterResources200Response {
	return &NullableGetClusterResources200Response{value: val, isSet: true}
}

func (v NullableGetClusterResources200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClusterResources200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


