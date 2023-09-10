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

// checks if the CreateAccessTicket200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccessTicket200ResponseData{}

// CreateAccessTicket200ResponseData struct for CreateAccessTicket200ResponseData
type CreateAccessTicket200ResponseData struct {
	CSRFPreventionToken *string `json:"CSRFPreventionToken,omitempty"`
	Clustername *string `json:"clustername,omitempty"`
	Ticket *string `json:"ticket,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewCreateAccessTicket200ResponseData instantiates a new CreateAccessTicket200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccessTicket200ResponseData() *CreateAccessTicket200ResponseData {
	this := CreateAccessTicket200ResponseData{}
	return &this
}

// NewCreateAccessTicket200ResponseDataWithDefaults instantiates a new CreateAccessTicket200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccessTicket200ResponseDataWithDefaults() *CreateAccessTicket200ResponseData {
	this := CreateAccessTicket200ResponseData{}
	return &this
}

// GetCSRFPreventionToken returns the CSRFPreventionToken field value if set, zero value otherwise.
func (o *CreateAccessTicket200ResponseData) GetCSRFPreventionToken() string {
	if o == nil || IsNil(o.CSRFPreventionToken) {
		var ret string
		return ret
	}
	return *o.CSRFPreventionToken
}

// GetCSRFPreventionTokenOk returns a tuple with the CSRFPreventionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicket200ResponseData) GetCSRFPreventionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CSRFPreventionToken) {
		return nil, false
	}
	return o.CSRFPreventionToken, true
}

// HasCSRFPreventionToken returns a boolean if a field has been set.
func (o *CreateAccessTicket200ResponseData) HasCSRFPreventionToken() bool {
	if o != nil && !IsNil(o.CSRFPreventionToken) {
		return true
	}

	return false
}

// SetCSRFPreventionToken gets a reference to the given string and assigns it to the CSRFPreventionToken field.
func (o *CreateAccessTicket200ResponseData) SetCSRFPreventionToken(v string) {
	o.CSRFPreventionToken = &v
}

// GetClustername returns the Clustername field value if set, zero value otherwise.
func (o *CreateAccessTicket200ResponseData) GetClustername() string {
	if o == nil || IsNil(o.Clustername) {
		var ret string
		return ret
	}
	return *o.Clustername
}

// GetClusternameOk returns a tuple with the Clustername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicket200ResponseData) GetClusternameOk() (*string, bool) {
	if o == nil || IsNil(o.Clustername) {
		return nil, false
	}
	return o.Clustername, true
}

// HasClustername returns a boolean if a field has been set.
func (o *CreateAccessTicket200ResponseData) HasClustername() bool {
	if o != nil && !IsNil(o.Clustername) {
		return true
	}

	return false
}

// SetClustername gets a reference to the given string and assigns it to the Clustername field.
func (o *CreateAccessTicket200ResponseData) SetClustername(v string) {
	o.Clustername = &v
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *CreateAccessTicket200ResponseData) GetTicket() string {
	if o == nil || IsNil(o.Ticket) {
		var ret string
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicket200ResponseData) GetTicketOk() (*string, bool) {
	if o == nil || IsNil(o.Ticket) {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *CreateAccessTicket200ResponseData) HasTicket() bool {
	if o != nil && !IsNil(o.Ticket) {
		return true
	}

	return false
}

// SetTicket gets a reference to the given string and assigns it to the Ticket field.
func (o *CreateAccessTicket200ResponseData) SetTicket(v string) {
	o.Ticket = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateAccessTicket200ResponseData) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicket200ResponseData) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateAccessTicket200ResponseData) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateAccessTicket200ResponseData) SetUsername(v string) {
	o.Username = &v
}

func (o CreateAccessTicket200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccessTicket200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CSRFPreventionToken) {
		toSerialize["CSRFPreventionToken"] = o.CSRFPreventionToken
	}
	if !IsNil(o.Clustername) {
		toSerialize["clustername"] = o.Clustername
	}
	if !IsNil(o.Ticket) {
		toSerialize["ticket"] = o.Ticket
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableCreateAccessTicket200ResponseData struct {
	value *CreateAccessTicket200ResponseData
	isSet bool
}

func (v NullableCreateAccessTicket200ResponseData) Get() *CreateAccessTicket200ResponseData {
	return v.value
}

func (v *NullableCreateAccessTicket200ResponseData) Set(val *CreateAccessTicket200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccessTicket200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccessTicket200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccessTicket200ResponseData(val *CreateAccessTicket200ResponseData) *NullableCreateAccessTicket200ResponseData {
	return &NullableCreateAccessTicket200ResponseData{value: val, isSet: true}
}

func (v NullableCreateAccessTicket200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccessTicket200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


