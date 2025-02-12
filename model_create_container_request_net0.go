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

// checks if the CreateContainerRequestNet0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerRequestNet0{}

// CreateContainerRequestNet0 struct for CreateContainerRequestNet0
type CreateContainerRequestNet0 struct {
	Bridge *string `json:"bridge,omitempty"`
	Firewall *bool `json:"firewall,omitempty"`
	Gw *string `json:"gw,omitempty"`
	Gw6 *string `json:"gw6,omitempty"`
	Hwaddr *string `json:"hwaddr,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Ip6 *string `json:"ip6,omitempty"`
	LinkDown *bool `json:"link_down,omitempty"`
	Mtu *int64 `json:"mtu,omitempty"`
	Name string `json:"name"`
	Rate *float32 `json:"rate,omitempty"`
	Tag *int64 `json:"tag,omitempty"`
	Trunks *string `json:"trunks,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCreateContainerRequestNet0 instantiates a new CreateContainerRequestNet0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerRequestNet0(name string) *CreateContainerRequestNet0 {
	this := CreateContainerRequestNet0{}
	this.Name = name
	return &this
}

// NewCreateContainerRequestNet0WithDefaults instantiates a new CreateContainerRequestNet0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerRequestNet0WithDefaults() *CreateContainerRequestNet0 {
	this := CreateContainerRequestNet0{}
	return &this
}

// GetBridge returns the Bridge field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetBridge() string {
	if o == nil || IsNil(o.Bridge) {
		var ret string
		return ret
	}
	return *o.Bridge
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetBridgeOk() (*string, bool) {
	if o == nil || IsNil(o.Bridge) {
		return nil, false
	}
	return o.Bridge, true
}

// HasBridge returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasBridge() bool {
	if o != nil && !IsNil(o.Bridge) {
		return true
	}

	return false
}

// SetBridge gets a reference to the given string and assigns it to the Bridge field.
func (o *CreateContainerRequestNet0) SetBridge(v string) {
	o.Bridge = &v
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetFirewall() bool {
	if o == nil || IsNil(o.Firewall) {
		var ret bool
		return ret
	}
	return *o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetFirewallOk() (*bool, bool) {
	if o == nil || IsNil(o.Firewall) {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasFirewall() bool {
	if o != nil && !IsNil(o.Firewall) {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given bool and assigns it to the Firewall field.
func (o *CreateContainerRequestNet0) SetFirewall(v bool) {
	o.Firewall = &v
}

// GetGw returns the Gw field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetGw() string {
	if o == nil || IsNil(o.Gw) {
		var ret string
		return ret
	}
	return *o.Gw
}

// GetGwOk returns a tuple with the Gw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetGwOk() (*string, bool) {
	if o == nil || IsNil(o.Gw) {
		return nil, false
	}
	return o.Gw, true
}

// HasGw returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasGw() bool {
	if o != nil && !IsNil(o.Gw) {
		return true
	}

	return false
}

// SetGw gets a reference to the given string and assigns it to the Gw field.
func (o *CreateContainerRequestNet0) SetGw(v string) {
	o.Gw = &v
}

// GetGw6 returns the Gw6 field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetGw6() string {
	if o == nil || IsNil(o.Gw6) {
		var ret string
		return ret
	}
	return *o.Gw6
}

// GetGw6Ok returns a tuple with the Gw6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetGw6Ok() (*string, bool) {
	if o == nil || IsNil(o.Gw6) {
		return nil, false
	}
	return o.Gw6, true
}

// HasGw6 returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasGw6() bool {
	if o != nil && !IsNil(o.Gw6) {
		return true
	}

	return false
}

// SetGw6 gets a reference to the given string and assigns it to the Gw6 field.
func (o *CreateContainerRequestNet0) SetGw6(v string) {
	o.Gw6 = &v
}

// GetHwaddr returns the Hwaddr field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetHwaddr() string {
	if o == nil || IsNil(o.Hwaddr) {
		var ret string
		return ret
	}
	return *o.Hwaddr
}

// GetHwaddrOk returns a tuple with the Hwaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetHwaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Hwaddr) {
		return nil, false
	}
	return o.Hwaddr, true
}

// HasHwaddr returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasHwaddr() bool {
	if o != nil && !IsNil(o.Hwaddr) {
		return true
	}

	return false
}

// SetHwaddr gets a reference to the given string and assigns it to the Hwaddr field.
func (o *CreateContainerRequestNet0) SetHwaddr(v string) {
	o.Hwaddr = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CreateContainerRequestNet0) SetIp(v string) {
	o.Ip = &v
}

// GetIp6 returns the Ip6 field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetIp6() string {
	if o == nil || IsNil(o.Ip6) {
		var ret string
		return ret
	}
	return *o.Ip6
}

// GetIp6Ok returns a tuple with the Ip6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetIp6Ok() (*string, bool) {
	if o == nil || IsNil(o.Ip6) {
		return nil, false
	}
	return o.Ip6, true
}

// HasIp6 returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasIp6() bool {
	if o != nil && !IsNil(o.Ip6) {
		return true
	}

	return false
}

// SetIp6 gets a reference to the given string and assigns it to the Ip6 field.
func (o *CreateContainerRequestNet0) SetIp6(v string) {
	o.Ip6 = &v
}

// GetLinkDown returns the LinkDown field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetLinkDown() bool {
	if o == nil || IsNil(o.LinkDown) {
		var ret bool
		return ret
	}
	return *o.LinkDown
}

// GetLinkDownOk returns a tuple with the LinkDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetLinkDownOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkDown) {
		return nil, false
	}
	return o.LinkDown, true
}

// HasLinkDown returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasLinkDown() bool {
	if o != nil && !IsNil(o.LinkDown) {
		return true
	}

	return false
}

// SetLinkDown gets a reference to the given bool and assigns it to the LinkDown field.
func (o *CreateContainerRequestNet0) SetLinkDown(v bool) {
	o.LinkDown = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetMtu() int64 {
	if o == nil || IsNil(o.Mtu) {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetMtuOk() (*int64, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *CreateContainerRequestNet0) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value
func (o *CreateContainerRequestNet0) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateContainerRequestNet0) SetName(v string) {
	o.Name = v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *CreateContainerRequestNet0) SetRate(v float32) {
	o.Rate = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetTag() int64 {
	if o == nil || IsNil(o.Tag) {
		var ret int64
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetTagOk() (*int64, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given int64 and assigns it to the Tag field.
func (o *CreateContainerRequestNet0) SetTag(v int64) {
	o.Tag = &v
}

// GetTrunks returns the Trunks field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetTrunks() string {
	if o == nil || IsNil(o.Trunks) {
		var ret string
		return ret
	}
	return *o.Trunks
}

// GetTrunksOk returns a tuple with the Trunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetTrunksOk() (*string, bool) {
	if o == nil || IsNil(o.Trunks) {
		return nil, false
	}
	return o.Trunks, true
}

// HasTrunks returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasTrunks() bool {
	if o != nil && !IsNil(o.Trunks) {
		return true
	}

	return false
}

// SetTrunks gets a reference to the given string and assigns it to the Trunks field.
func (o *CreateContainerRequestNet0) SetTrunks(v string) {
	o.Trunks = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateContainerRequestNet0) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerRequestNet0) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateContainerRequestNet0) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateContainerRequestNet0) SetType(v string) {
	o.Type = &v
}

func (o CreateContainerRequestNet0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainerRequestNet0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bridge) {
		toSerialize["bridge"] = o.Bridge
	}
	if !IsNil(o.Firewall) {
		toSerialize["firewall"] = o.Firewall
	}
	if !IsNil(o.Gw) {
		toSerialize["gw"] = o.Gw
	}
	if !IsNil(o.Gw6) {
		toSerialize["gw6"] = o.Gw6
	}
	if !IsNil(o.Hwaddr) {
		toSerialize["hwaddr"] = o.Hwaddr
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Ip6) {
		toSerialize["ip6"] = o.Ip6
	}
	if !IsNil(o.LinkDown) {
		toSerialize["link_down"] = o.LinkDown
	}
	if !IsNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Trunks) {
		toSerialize["trunks"] = o.Trunks
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCreateContainerRequestNet0 struct {
	value *CreateContainerRequestNet0
	isSet bool
}

func (v NullableCreateContainerRequestNet0) Get() *CreateContainerRequestNet0 {
	return v.value
}

func (v *NullableCreateContainerRequestNet0) Set(val *CreateContainerRequestNet0) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerRequestNet0) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerRequestNet0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerRequestNet0(val *CreateContainerRequestNet0) *NullableCreateContainerRequestNet0 {
	return &NullableCreateContainerRequestNet0{value: val, isSet: true}
}

func (v NullableCreateContainerRequestNet0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerRequestNet0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


