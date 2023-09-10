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

// checks if the GetVMConfig200ResponseDataNet0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVMConfig200ResponseDataNet0{}

// GetVMConfig200ResponseDataNet0 struct for GetVMConfig200ResponseDataNet0
type GetVMConfig200ResponseDataNet0 struct {
	// Bridge to attach the network device to. The Proxmox VE standard bridge is called 'vmbr0'.  If you do not specify a bridge, we create a kvm user (NATed) network device, which provides DHCP and DNS services. The following addresses are used:   10.0.2.2   Gateway  10.0.2.3   DNS Server  10.0.2.4   SMB Server  The DHCP server assign addresses to the guest starting from 10.0.2.15. 
	Bridge *string `json:"bridge,omitempty"`
	E1000 *string `json:"e1000,omitempty"`
	E100082540em *string `json:"e1000-82540em,omitempty"`
	E100082544gc *string `json:"e1000-82544gc,omitempty"`
	E100082545em *string `json:"e1000-82545em,omitempty"`
	E1000e *string `json:"e1000e,omitempty"`
	// Whether this interface should be protected by the firewall.
	Firewall *int32 `json:"firewall,omitempty"`
	I82551 *string `json:"i82551,omitempty"`
	I82557b *string `json:"i82557b,omitempty"`
	I82559er *string `json:"i82559er,omitempty"`
	// Whether this interface should be disconnected (like pulling the plug).
	LinkDown *int32 `json:"link_down,omitempty"`
	// MAC address. That address must be unique withing your network. This is automatically generated if not specified.
	Macaddr *string `json:"macaddr,omitempty"`
	// Network Card Model. The 'virtio' model provides the best performance with very low CPU overhead. If your guest does not support this driver, it is usually best to use 'e1000'.
	Model *string `json:"model,omitempty"`
	// Force MTU, for VirtIO only. Set to '1' to use the bridge MTU
	Mtu *int64 `json:"mtu,omitempty"`
	Ne2kIsa *string `json:"ne2k_isa,omitempty"`
	Ne2kPci *string `json:"ne2k_pci,omitempty"`
	Pcnet *string `json:"pcnet,omitempty"`
	// Number of packet queues to be used on the device.
	Queues *int64 `json:"queues,omitempty"`
	// Rate limit in mbps (megabytes per second) as floating point number.
	Rate *float32 `json:"rate,omitempty"`
	Rtl8139 *string `json:"rtl8139,omitempty"`
	// VLAN tag to apply to packets on this interface.
	Tag *int64 `json:"tag,omitempty"`
	// VLAN trunks to pass through this interface.
	Trunks *string `json:"trunks,omitempty"`
	Virtio *string `json:"virtio,omitempty"`
	Vmxnet3 *string `json:"vmxnet3,omitempty"`
}

// NewGetVMConfig200ResponseDataNet0 instantiates a new GetVMConfig200ResponseDataNet0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVMConfig200ResponseDataNet0() *GetVMConfig200ResponseDataNet0 {
	this := GetVMConfig200ResponseDataNet0{}
	return &this
}

// NewGetVMConfig200ResponseDataNet0WithDefaults instantiates a new GetVMConfig200ResponseDataNet0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVMConfig200ResponseDataNet0WithDefaults() *GetVMConfig200ResponseDataNet0 {
	this := GetVMConfig200ResponseDataNet0{}
	return &this
}

// GetBridge returns the Bridge field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetBridge() string {
	if o == nil || IsNil(o.Bridge) {
		var ret string
		return ret
	}
	return *o.Bridge
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetBridgeOk() (*string, bool) {
	if o == nil || IsNil(o.Bridge) {
		return nil, false
	}
	return o.Bridge, true
}

// HasBridge returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasBridge() bool {
	if o != nil && !IsNil(o.Bridge) {
		return true
	}

	return false
}

// SetBridge gets a reference to the given string and assigns it to the Bridge field.
func (o *GetVMConfig200ResponseDataNet0) SetBridge(v string) {
	o.Bridge = &v
}

// GetE1000 returns the E1000 field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetE1000() string {
	if o == nil || IsNil(o.E1000) {
		var ret string
		return ret
	}
	return *o.E1000
}

// GetE1000Ok returns a tuple with the E1000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetE1000Ok() (*string, bool) {
	if o == nil || IsNil(o.E1000) {
		return nil, false
	}
	return o.E1000, true
}

// HasE1000 returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasE1000() bool {
	if o != nil && !IsNil(o.E1000) {
		return true
	}

	return false
}

// SetE1000 gets a reference to the given string and assigns it to the E1000 field.
func (o *GetVMConfig200ResponseDataNet0) SetE1000(v string) {
	o.E1000 = &v
}

// GetE100082540em returns the E100082540em field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetE100082540em() string {
	if o == nil || IsNil(o.E100082540em) {
		var ret string
		return ret
	}
	return *o.E100082540em
}

// GetE100082540emOk returns a tuple with the E100082540em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetE100082540emOk() (*string, bool) {
	if o == nil || IsNil(o.E100082540em) {
		return nil, false
	}
	return o.E100082540em, true
}

// HasE100082540em returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasE100082540em() bool {
	if o != nil && !IsNil(o.E100082540em) {
		return true
	}

	return false
}

// SetE100082540em gets a reference to the given string and assigns it to the E100082540em field.
func (o *GetVMConfig200ResponseDataNet0) SetE100082540em(v string) {
	o.E100082540em = &v
}

// GetE100082544gc returns the E100082544gc field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetE100082544gc() string {
	if o == nil || IsNil(o.E100082544gc) {
		var ret string
		return ret
	}
	return *o.E100082544gc
}

// GetE100082544gcOk returns a tuple with the E100082544gc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetE100082544gcOk() (*string, bool) {
	if o == nil || IsNil(o.E100082544gc) {
		return nil, false
	}
	return o.E100082544gc, true
}

// HasE100082544gc returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasE100082544gc() bool {
	if o != nil && !IsNil(o.E100082544gc) {
		return true
	}

	return false
}

// SetE100082544gc gets a reference to the given string and assigns it to the E100082544gc field.
func (o *GetVMConfig200ResponseDataNet0) SetE100082544gc(v string) {
	o.E100082544gc = &v
}

// GetE100082545em returns the E100082545em field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetE100082545em() string {
	if o == nil || IsNil(o.E100082545em) {
		var ret string
		return ret
	}
	return *o.E100082545em
}

// GetE100082545emOk returns a tuple with the E100082545em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetE100082545emOk() (*string, bool) {
	if o == nil || IsNil(o.E100082545em) {
		return nil, false
	}
	return o.E100082545em, true
}

// HasE100082545em returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasE100082545em() bool {
	if o != nil && !IsNil(o.E100082545em) {
		return true
	}

	return false
}

// SetE100082545em gets a reference to the given string and assigns it to the E100082545em field.
func (o *GetVMConfig200ResponseDataNet0) SetE100082545em(v string) {
	o.E100082545em = &v
}

// GetE1000e returns the E1000e field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetE1000e() string {
	if o == nil || IsNil(o.E1000e) {
		var ret string
		return ret
	}
	return *o.E1000e
}

// GetE1000eOk returns a tuple with the E1000e field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetE1000eOk() (*string, bool) {
	if o == nil || IsNil(o.E1000e) {
		return nil, false
	}
	return o.E1000e, true
}

// HasE1000e returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasE1000e() bool {
	if o != nil && !IsNil(o.E1000e) {
		return true
	}

	return false
}

// SetE1000e gets a reference to the given string and assigns it to the E1000e field.
func (o *GetVMConfig200ResponseDataNet0) SetE1000e(v string) {
	o.E1000e = &v
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetFirewall() int32 {
	if o == nil || IsNil(o.Firewall) {
		var ret int32
		return ret
	}
	return *o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetFirewallOk() (*int32, bool) {
	if o == nil || IsNil(o.Firewall) {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasFirewall() bool {
	if o != nil && !IsNil(o.Firewall) {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given int32 and assigns it to the Firewall field.
func (o *GetVMConfig200ResponseDataNet0) SetFirewall(v int32) {
	o.Firewall = &v
}

// GetI82551 returns the I82551 field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetI82551() string {
	if o == nil || IsNil(o.I82551) {
		var ret string
		return ret
	}
	return *o.I82551
}

// GetI82551Ok returns a tuple with the I82551 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetI82551Ok() (*string, bool) {
	if o == nil || IsNil(o.I82551) {
		return nil, false
	}
	return o.I82551, true
}

// HasI82551 returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasI82551() bool {
	if o != nil && !IsNil(o.I82551) {
		return true
	}

	return false
}

// SetI82551 gets a reference to the given string and assigns it to the I82551 field.
func (o *GetVMConfig200ResponseDataNet0) SetI82551(v string) {
	o.I82551 = &v
}

// GetI82557b returns the I82557b field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetI82557b() string {
	if o == nil || IsNil(o.I82557b) {
		var ret string
		return ret
	}
	return *o.I82557b
}

// GetI82557bOk returns a tuple with the I82557b field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetI82557bOk() (*string, bool) {
	if o == nil || IsNil(o.I82557b) {
		return nil, false
	}
	return o.I82557b, true
}

// HasI82557b returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasI82557b() bool {
	if o != nil && !IsNil(o.I82557b) {
		return true
	}

	return false
}

// SetI82557b gets a reference to the given string and assigns it to the I82557b field.
func (o *GetVMConfig200ResponseDataNet0) SetI82557b(v string) {
	o.I82557b = &v
}

// GetI82559er returns the I82559er field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetI82559er() string {
	if o == nil || IsNil(o.I82559er) {
		var ret string
		return ret
	}
	return *o.I82559er
}

// GetI82559erOk returns a tuple with the I82559er field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetI82559erOk() (*string, bool) {
	if o == nil || IsNil(o.I82559er) {
		return nil, false
	}
	return o.I82559er, true
}

// HasI82559er returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasI82559er() bool {
	if o != nil && !IsNil(o.I82559er) {
		return true
	}

	return false
}

// SetI82559er gets a reference to the given string and assigns it to the I82559er field.
func (o *GetVMConfig200ResponseDataNet0) SetI82559er(v string) {
	o.I82559er = &v
}

// GetLinkDown returns the LinkDown field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetLinkDown() int32 {
	if o == nil || IsNil(o.LinkDown) {
		var ret int32
		return ret
	}
	return *o.LinkDown
}

// GetLinkDownOk returns a tuple with the LinkDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetLinkDownOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkDown) {
		return nil, false
	}
	return o.LinkDown, true
}

// HasLinkDown returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasLinkDown() bool {
	if o != nil && !IsNil(o.LinkDown) {
		return true
	}

	return false
}

// SetLinkDown gets a reference to the given int32 and assigns it to the LinkDown field.
func (o *GetVMConfig200ResponseDataNet0) SetLinkDown(v int32) {
	o.LinkDown = &v
}

// GetMacaddr returns the Macaddr field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetMacaddr() string {
	if o == nil || IsNil(o.Macaddr) {
		var ret string
		return ret
	}
	return *o.Macaddr
}

// GetMacaddrOk returns a tuple with the Macaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetMacaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Macaddr) {
		return nil, false
	}
	return o.Macaddr, true
}

// HasMacaddr returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasMacaddr() bool {
	if o != nil && !IsNil(o.Macaddr) {
		return true
	}

	return false
}

// SetMacaddr gets a reference to the given string and assigns it to the Macaddr field.
func (o *GetVMConfig200ResponseDataNet0) SetMacaddr(v string) {
	o.Macaddr = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetVMConfig200ResponseDataNet0) SetModel(v string) {
	o.Model = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetMtu() int64 {
	if o == nil || IsNil(o.Mtu) {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetMtuOk() (*int64, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *GetVMConfig200ResponseDataNet0) SetMtu(v int64) {
	o.Mtu = &v
}

// GetNe2kIsa returns the Ne2kIsa field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetNe2kIsa() string {
	if o == nil || IsNil(o.Ne2kIsa) {
		var ret string
		return ret
	}
	return *o.Ne2kIsa
}

// GetNe2kIsaOk returns a tuple with the Ne2kIsa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetNe2kIsaOk() (*string, bool) {
	if o == nil || IsNil(o.Ne2kIsa) {
		return nil, false
	}
	return o.Ne2kIsa, true
}

// HasNe2kIsa returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasNe2kIsa() bool {
	if o != nil && !IsNil(o.Ne2kIsa) {
		return true
	}

	return false
}

// SetNe2kIsa gets a reference to the given string and assigns it to the Ne2kIsa field.
func (o *GetVMConfig200ResponseDataNet0) SetNe2kIsa(v string) {
	o.Ne2kIsa = &v
}

// GetNe2kPci returns the Ne2kPci field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetNe2kPci() string {
	if o == nil || IsNil(o.Ne2kPci) {
		var ret string
		return ret
	}
	return *o.Ne2kPci
}

// GetNe2kPciOk returns a tuple with the Ne2kPci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetNe2kPciOk() (*string, bool) {
	if o == nil || IsNil(o.Ne2kPci) {
		return nil, false
	}
	return o.Ne2kPci, true
}

// HasNe2kPci returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasNe2kPci() bool {
	if o != nil && !IsNil(o.Ne2kPci) {
		return true
	}

	return false
}

// SetNe2kPci gets a reference to the given string and assigns it to the Ne2kPci field.
func (o *GetVMConfig200ResponseDataNet0) SetNe2kPci(v string) {
	o.Ne2kPci = &v
}

// GetPcnet returns the Pcnet field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetPcnet() string {
	if o == nil || IsNil(o.Pcnet) {
		var ret string
		return ret
	}
	return *o.Pcnet
}

// GetPcnetOk returns a tuple with the Pcnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetPcnetOk() (*string, bool) {
	if o == nil || IsNil(o.Pcnet) {
		return nil, false
	}
	return o.Pcnet, true
}

// HasPcnet returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasPcnet() bool {
	if o != nil && !IsNil(o.Pcnet) {
		return true
	}

	return false
}

// SetPcnet gets a reference to the given string and assigns it to the Pcnet field.
func (o *GetVMConfig200ResponseDataNet0) SetPcnet(v string) {
	o.Pcnet = &v
}

// GetQueues returns the Queues field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetQueues() int64 {
	if o == nil || IsNil(o.Queues) {
		var ret int64
		return ret
	}
	return *o.Queues
}

// GetQueuesOk returns a tuple with the Queues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetQueuesOk() (*int64, bool) {
	if o == nil || IsNil(o.Queues) {
		return nil, false
	}
	return o.Queues, true
}

// HasQueues returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasQueues() bool {
	if o != nil && !IsNil(o.Queues) {
		return true
	}

	return false
}

// SetQueues gets a reference to the given int64 and assigns it to the Queues field.
func (o *GetVMConfig200ResponseDataNet0) SetQueues(v int64) {
	o.Queues = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *GetVMConfig200ResponseDataNet0) SetRate(v float32) {
	o.Rate = &v
}

// GetRtl8139 returns the Rtl8139 field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetRtl8139() string {
	if o == nil || IsNil(o.Rtl8139) {
		var ret string
		return ret
	}
	return *o.Rtl8139
}

// GetRtl8139Ok returns a tuple with the Rtl8139 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetRtl8139Ok() (*string, bool) {
	if o == nil || IsNil(o.Rtl8139) {
		return nil, false
	}
	return o.Rtl8139, true
}

// HasRtl8139 returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasRtl8139() bool {
	if o != nil && !IsNil(o.Rtl8139) {
		return true
	}

	return false
}

// SetRtl8139 gets a reference to the given string and assigns it to the Rtl8139 field.
func (o *GetVMConfig200ResponseDataNet0) SetRtl8139(v string) {
	o.Rtl8139 = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetTag() int64 {
	if o == nil || IsNil(o.Tag) {
		var ret int64
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetTagOk() (*int64, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given int64 and assigns it to the Tag field.
func (o *GetVMConfig200ResponseDataNet0) SetTag(v int64) {
	o.Tag = &v
}

// GetTrunks returns the Trunks field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetTrunks() string {
	if o == nil || IsNil(o.Trunks) {
		var ret string
		return ret
	}
	return *o.Trunks
}

// GetTrunksOk returns a tuple with the Trunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetTrunksOk() (*string, bool) {
	if o == nil || IsNil(o.Trunks) {
		return nil, false
	}
	return o.Trunks, true
}

// HasTrunks returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasTrunks() bool {
	if o != nil && !IsNil(o.Trunks) {
		return true
	}

	return false
}

// SetTrunks gets a reference to the given string and assigns it to the Trunks field.
func (o *GetVMConfig200ResponseDataNet0) SetTrunks(v string) {
	o.Trunks = &v
}

// GetVirtio returns the Virtio field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetVirtio() string {
	if o == nil || IsNil(o.Virtio) {
		var ret string
		return ret
	}
	return *o.Virtio
}

// GetVirtioOk returns a tuple with the Virtio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetVirtioOk() (*string, bool) {
	if o == nil || IsNil(o.Virtio) {
		return nil, false
	}
	return o.Virtio, true
}

// HasVirtio returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasVirtio() bool {
	if o != nil && !IsNil(o.Virtio) {
		return true
	}

	return false
}

// SetVirtio gets a reference to the given string and assigns it to the Virtio field.
func (o *GetVMConfig200ResponseDataNet0) SetVirtio(v string) {
	o.Virtio = &v
}

// GetVmxnet3 returns the Vmxnet3 field value if set, zero value otherwise.
func (o *GetVMConfig200ResponseDataNet0) GetVmxnet3() string {
	if o == nil || IsNil(o.Vmxnet3) {
		var ret string
		return ret
	}
	return *o.Vmxnet3
}

// GetVmxnet3Ok returns a tuple with the Vmxnet3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVMConfig200ResponseDataNet0) GetVmxnet3Ok() (*string, bool) {
	if o == nil || IsNil(o.Vmxnet3) {
		return nil, false
	}
	return o.Vmxnet3, true
}

// HasVmxnet3 returns a boolean if a field has been set.
func (o *GetVMConfig200ResponseDataNet0) HasVmxnet3() bool {
	if o != nil && !IsNil(o.Vmxnet3) {
		return true
	}

	return false
}

// SetVmxnet3 gets a reference to the given string and assigns it to the Vmxnet3 field.
func (o *GetVMConfig200ResponseDataNet0) SetVmxnet3(v string) {
	o.Vmxnet3 = &v
}

func (o GetVMConfig200ResponseDataNet0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVMConfig200ResponseDataNet0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bridge) {
		toSerialize["bridge"] = o.Bridge
	}
	if !IsNil(o.E1000) {
		toSerialize["e1000"] = o.E1000
	}
	if !IsNil(o.E100082540em) {
		toSerialize["e1000-82540em"] = o.E100082540em
	}
	if !IsNil(o.E100082544gc) {
		toSerialize["e1000-82544gc"] = o.E100082544gc
	}
	if !IsNil(o.E100082545em) {
		toSerialize["e1000-82545em"] = o.E100082545em
	}
	if !IsNil(o.E1000e) {
		toSerialize["e1000e"] = o.E1000e
	}
	if !IsNil(o.Firewall) {
		toSerialize["firewall"] = o.Firewall
	}
	if !IsNil(o.I82551) {
		toSerialize["i82551"] = o.I82551
	}
	if !IsNil(o.I82557b) {
		toSerialize["i82557b"] = o.I82557b
	}
	if !IsNil(o.I82559er) {
		toSerialize["i82559er"] = o.I82559er
	}
	if !IsNil(o.LinkDown) {
		toSerialize["link_down"] = o.LinkDown
	}
	if !IsNil(o.Macaddr) {
		toSerialize["macaddr"] = o.Macaddr
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	if !IsNil(o.Ne2kIsa) {
		toSerialize["ne2k_isa"] = o.Ne2kIsa
	}
	if !IsNil(o.Ne2kPci) {
		toSerialize["ne2k_pci"] = o.Ne2kPci
	}
	if !IsNil(o.Pcnet) {
		toSerialize["pcnet"] = o.Pcnet
	}
	if !IsNil(o.Queues) {
		toSerialize["queues"] = o.Queues
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Rtl8139) {
		toSerialize["rtl8139"] = o.Rtl8139
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Trunks) {
		toSerialize["trunks"] = o.Trunks
	}
	if !IsNil(o.Virtio) {
		toSerialize["virtio"] = o.Virtio
	}
	if !IsNil(o.Vmxnet3) {
		toSerialize["vmxnet3"] = o.Vmxnet3
	}
	return toSerialize, nil
}

type NullableGetVMConfig200ResponseDataNet0 struct {
	value *GetVMConfig200ResponseDataNet0
	isSet bool
}

func (v NullableGetVMConfig200ResponseDataNet0) Get() *GetVMConfig200ResponseDataNet0 {
	return v.value
}

func (v *NullableGetVMConfig200ResponseDataNet0) Set(val *GetVMConfig200ResponseDataNet0) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVMConfig200ResponseDataNet0) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVMConfig200ResponseDataNet0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVMConfig200ResponseDataNet0(val *GetVMConfig200ResponseDataNet0) *NullableGetVMConfig200ResponseDataNet0 {
	return &NullableGetVMConfig200ResponseDataNet0{value: val, isSet: true}
}

func (v NullableGetVMConfig200ResponseDataNet0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVMConfig200ResponseDataNet0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


