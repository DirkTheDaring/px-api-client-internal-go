# GetVMConfig200ResponseDataNet0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bridge** | Pointer to **string** | Bridge to attach the network device to. The Proxmox VE standard bridge is called &#39;vmbr0&#39;.  If you do not specify a bridge, we create a kvm user (NATed) network device, which provides DHCP and DNS services. The following addresses are used:   10.0.2.2   Gateway  10.0.2.3   DNS Server  10.0.2.4   SMB Server  The DHCP server assign addresses to the guest starting from 10.0.2.15.  | [optional] 
**E1000** | Pointer to **string** |  | [optional] 
**E100082540em** | Pointer to **string** |  | [optional] 
**E100082544gc** | Pointer to **string** |  | [optional] 
**E100082545em** | Pointer to **string** |  | [optional] 
**E1000e** | Pointer to **string** |  | [optional] 
**Firewall** | Pointer to **bool** | Whether this interface should be protected by the firewall. | [optional] 
**I82551** | Pointer to **string** |  | [optional] 
**I82557b** | Pointer to **string** |  | [optional] 
**I82559er** | Pointer to **string** |  | [optional] 
**LinkDown** | Pointer to **bool** | Whether this interface should be disconnected (like pulling the plug). | [optional] 
**Macaddr** | Pointer to **string** | MAC address. That address must be unique withing your network. This is automatically generated if not specified. | [optional] 
**Model** | Pointer to **string** | Network Card Model. The &#39;virtio&#39; model provides the best performance with very low CPU overhead. If your guest does not support this driver, it is usually best to use &#39;e1000&#39;. | [optional] 
**Mtu** | Pointer to **int64** | Force MTU, for VirtIO only. Set to &#39;1&#39; to use the bridge MTU | [optional] 
**Ne2kIsa** | Pointer to **string** |  | [optional] 
**Ne2kPci** | Pointer to **string** |  | [optional] 
**Pcnet** | Pointer to **string** |  | [optional] 
**Queues** | Pointer to **int64** | Number of packet queues to be used on the device. | [optional] 
**Rate** | Pointer to **float32** | Rate limit in mbps (megabytes per second) as floating point number. | [optional] 
**Rtl8139** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **int64** | VLAN tag to apply to packets on this interface. | [optional] 
**Trunks** | Pointer to **string** | VLAN trunks to pass through this interface. | [optional] 
**Virtio** | Pointer to **string** |  | [optional] 
**Vmxnet3** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVMConfig200ResponseDataNet0

`func NewGetVMConfig200ResponseDataNet0() *GetVMConfig200ResponseDataNet0`

NewGetVMConfig200ResponseDataNet0 instantiates a new GetVMConfig200ResponseDataNet0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataNet0WithDefaults

`func NewGetVMConfig200ResponseDataNet0WithDefaults() *GetVMConfig200ResponseDataNet0`

NewGetVMConfig200ResponseDataNet0WithDefaults instantiates a new GetVMConfig200ResponseDataNet0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridge

`func (o *GetVMConfig200ResponseDataNet0) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *GetVMConfig200ResponseDataNet0) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *GetVMConfig200ResponseDataNet0) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *GetVMConfig200ResponseDataNet0) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetE1000

`func (o *GetVMConfig200ResponseDataNet0) GetE1000() string`

GetE1000 returns the E1000 field if non-nil, zero value otherwise.

### GetE1000Ok

`func (o *GetVMConfig200ResponseDataNet0) GetE1000Ok() (*string, bool)`

GetE1000Ok returns a tuple with the E1000 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE1000

`func (o *GetVMConfig200ResponseDataNet0) SetE1000(v string)`

SetE1000 sets E1000 field to given value.

### HasE1000

`func (o *GetVMConfig200ResponseDataNet0) HasE1000() bool`

HasE1000 returns a boolean if a field has been set.

### GetE100082540em

`func (o *GetVMConfig200ResponseDataNet0) GetE100082540em() string`

GetE100082540em returns the E100082540em field if non-nil, zero value otherwise.

### GetE100082540emOk

`func (o *GetVMConfig200ResponseDataNet0) GetE100082540emOk() (*string, bool)`

GetE100082540emOk returns a tuple with the E100082540em field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE100082540em

`func (o *GetVMConfig200ResponseDataNet0) SetE100082540em(v string)`

SetE100082540em sets E100082540em field to given value.

### HasE100082540em

`func (o *GetVMConfig200ResponseDataNet0) HasE100082540em() bool`

HasE100082540em returns a boolean if a field has been set.

### GetE100082544gc

`func (o *GetVMConfig200ResponseDataNet0) GetE100082544gc() string`

GetE100082544gc returns the E100082544gc field if non-nil, zero value otherwise.

### GetE100082544gcOk

`func (o *GetVMConfig200ResponseDataNet0) GetE100082544gcOk() (*string, bool)`

GetE100082544gcOk returns a tuple with the E100082544gc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE100082544gc

`func (o *GetVMConfig200ResponseDataNet0) SetE100082544gc(v string)`

SetE100082544gc sets E100082544gc field to given value.

### HasE100082544gc

`func (o *GetVMConfig200ResponseDataNet0) HasE100082544gc() bool`

HasE100082544gc returns a boolean if a field has been set.

### GetE100082545em

`func (o *GetVMConfig200ResponseDataNet0) GetE100082545em() string`

GetE100082545em returns the E100082545em field if non-nil, zero value otherwise.

### GetE100082545emOk

`func (o *GetVMConfig200ResponseDataNet0) GetE100082545emOk() (*string, bool)`

GetE100082545emOk returns a tuple with the E100082545em field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE100082545em

`func (o *GetVMConfig200ResponseDataNet0) SetE100082545em(v string)`

SetE100082545em sets E100082545em field to given value.

### HasE100082545em

`func (o *GetVMConfig200ResponseDataNet0) HasE100082545em() bool`

HasE100082545em returns a boolean if a field has been set.

### GetE1000e

`func (o *GetVMConfig200ResponseDataNet0) GetE1000e() string`

GetE1000e returns the E1000e field if non-nil, zero value otherwise.

### GetE1000eOk

`func (o *GetVMConfig200ResponseDataNet0) GetE1000eOk() (*string, bool)`

GetE1000eOk returns a tuple with the E1000e field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE1000e

`func (o *GetVMConfig200ResponseDataNet0) SetE1000e(v string)`

SetE1000e sets E1000e field to given value.

### HasE1000e

`func (o *GetVMConfig200ResponseDataNet0) HasE1000e() bool`

HasE1000e returns a boolean if a field has been set.

### GetFirewall

`func (o *GetVMConfig200ResponseDataNet0) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *GetVMConfig200ResponseDataNet0) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *GetVMConfig200ResponseDataNet0) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *GetVMConfig200ResponseDataNet0) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetI82551

`func (o *GetVMConfig200ResponseDataNet0) GetI82551() string`

GetI82551 returns the I82551 field if non-nil, zero value otherwise.

### GetI82551Ok

`func (o *GetVMConfig200ResponseDataNet0) GetI82551Ok() (*string, bool)`

GetI82551Ok returns a tuple with the I82551 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI82551

`func (o *GetVMConfig200ResponseDataNet0) SetI82551(v string)`

SetI82551 sets I82551 field to given value.

### HasI82551

`func (o *GetVMConfig200ResponseDataNet0) HasI82551() bool`

HasI82551 returns a boolean if a field has been set.

### GetI82557b

`func (o *GetVMConfig200ResponseDataNet0) GetI82557b() string`

GetI82557b returns the I82557b field if non-nil, zero value otherwise.

### GetI82557bOk

`func (o *GetVMConfig200ResponseDataNet0) GetI82557bOk() (*string, bool)`

GetI82557bOk returns a tuple with the I82557b field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI82557b

`func (o *GetVMConfig200ResponseDataNet0) SetI82557b(v string)`

SetI82557b sets I82557b field to given value.

### HasI82557b

`func (o *GetVMConfig200ResponseDataNet0) HasI82557b() bool`

HasI82557b returns a boolean if a field has been set.

### GetI82559er

`func (o *GetVMConfig200ResponseDataNet0) GetI82559er() string`

GetI82559er returns the I82559er field if non-nil, zero value otherwise.

### GetI82559erOk

`func (o *GetVMConfig200ResponseDataNet0) GetI82559erOk() (*string, bool)`

GetI82559erOk returns a tuple with the I82559er field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI82559er

`func (o *GetVMConfig200ResponseDataNet0) SetI82559er(v string)`

SetI82559er sets I82559er field to given value.

### HasI82559er

`func (o *GetVMConfig200ResponseDataNet0) HasI82559er() bool`

HasI82559er returns a boolean if a field has been set.

### GetLinkDown

`func (o *GetVMConfig200ResponseDataNet0) GetLinkDown() bool`

GetLinkDown returns the LinkDown field if non-nil, zero value otherwise.

### GetLinkDownOk

`func (o *GetVMConfig200ResponseDataNet0) GetLinkDownOk() (*bool, bool)`

GetLinkDownOk returns a tuple with the LinkDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDown

`func (o *GetVMConfig200ResponseDataNet0) SetLinkDown(v bool)`

SetLinkDown sets LinkDown field to given value.

### HasLinkDown

`func (o *GetVMConfig200ResponseDataNet0) HasLinkDown() bool`

HasLinkDown returns a boolean if a field has been set.

### GetMacaddr

`func (o *GetVMConfig200ResponseDataNet0) GetMacaddr() string`

GetMacaddr returns the Macaddr field if non-nil, zero value otherwise.

### GetMacaddrOk

`func (o *GetVMConfig200ResponseDataNet0) GetMacaddrOk() (*string, bool)`

GetMacaddrOk returns a tuple with the Macaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacaddr

`func (o *GetVMConfig200ResponseDataNet0) SetMacaddr(v string)`

SetMacaddr sets Macaddr field to given value.

### HasMacaddr

`func (o *GetVMConfig200ResponseDataNet0) HasMacaddr() bool`

HasMacaddr returns a boolean if a field has been set.

### GetModel

`func (o *GetVMConfig200ResponseDataNet0) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetVMConfig200ResponseDataNet0) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetVMConfig200ResponseDataNet0) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetVMConfig200ResponseDataNet0) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMtu

`func (o *GetVMConfig200ResponseDataNet0) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *GetVMConfig200ResponseDataNet0) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *GetVMConfig200ResponseDataNet0) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *GetVMConfig200ResponseDataNet0) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNe2kIsa

`func (o *GetVMConfig200ResponseDataNet0) GetNe2kIsa() string`

GetNe2kIsa returns the Ne2kIsa field if non-nil, zero value otherwise.

### GetNe2kIsaOk

`func (o *GetVMConfig200ResponseDataNet0) GetNe2kIsaOk() (*string, bool)`

GetNe2kIsaOk returns a tuple with the Ne2kIsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNe2kIsa

`func (o *GetVMConfig200ResponseDataNet0) SetNe2kIsa(v string)`

SetNe2kIsa sets Ne2kIsa field to given value.

### HasNe2kIsa

`func (o *GetVMConfig200ResponseDataNet0) HasNe2kIsa() bool`

HasNe2kIsa returns a boolean if a field has been set.

### GetNe2kPci

`func (o *GetVMConfig200ResponseDataNet0) GetNe2kPci() string`

GetNe2kPci returns the Ne2kPci field if non-nil, zero value otherwise.

### GetNe2kPciOk

`func (o *GetVMConfig200ResponseDataNet0) GetNe2kPciOk() (*string, bool)`

GetNe2kPciOk returns a tuple with the Ne2kPci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNe2kPci

`func (o *GetVMConfig200ResponseDataNet0) SetNe2kPci(v string)`

SetNe2kPci sets Ne2kPci field to given value.

### HasNe2kPci

`func (o *GetVMConfig200ResponseDataNet0) HasNe2kPci() bool`

HasNe2kPci returns a boolean if a field has been set.

### GetPcnet

`func (o *GetVMConfig200ResponseDataNet0) GetPcnet() string`

GetPcnet returns the Pcnet field if non-nil, zero value otherwise.

### GetPcnetOk

`func (o *GetVMConfig200ResponseDataNet0) GetPcnetOk() (*string, bool)`

GetPcnetOk returns a tuple with the Pcnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcnet

`func (o *GetVMConfig200ResponseDataNet0) SetPcnet(v string)`

SetPcnet sets Pcnet field to given value.

### HasPcnet

`func (o *GetVMConfig200ResponseDataNet0) HasPcnet() bool`

HasPcnet returns a boolean if a field has been set.

### GetQueues

`func (o *GetVMConfig200ResponseDataNet0) GetQueues() int64`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *GetVMConfig200ResponseDataNet0) GetQueuesOk() (*int64, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *GetVMConfig200ResponseDataNet0) SetQueues(v int64)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *GetVMConfig200ResponseDataNet0) HasQueues() bool`

HasQueues returns a boolean if a field has been set.

### GetRate

`func (o *GetVMConfig200ResponseDataNet0) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetVMConfig200ResponseDataNet0) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetVMConfig200ResponseDataNet0) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetVMConfig200ResponseDataNet0) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRtl8139

`func (o *GetVMConfig200ResponseDataNet0) GetRtl8139() string`

GetRtl8139 returns the Rtl8139 field if non-nil, zero value otherwise.

### GetRtl8139Ok

`func (o *GetVMConfig200ResponseDataNet0) GetRtl8139Ok() (*string, bool)`

GetRtl8139Ok returns a tuple with the Rtl8139 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtl8139

`func (o *GetVMConfig200ResponseDataNet0) SetRtl8139(v string)`

SetRtl8139 sets Rtl8139 field to given value.

### HasRtl8139

`func (o *GetVMConfig200ResponseDataNet0) HasRtl8139() bool`

HasRtl8139 returns a boolean if a field has been set.

### GetTag

`func (o *GetVMConfig200ResponseDataNet0) GetTag() int64`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetVMConfig200ResponseDataNet0) GetTagOk() (*int64, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetVMConfig200ResponseDataNet0) SetTag(v int64)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetVMConfig200ResponseDataNet0) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTrunks

`func (o *GetVMConfig200ResponseDataNet0) GetTrunks() string`

GetTrunks returns the Trunks field if non-nil, zero value otherwise.

### GetTrunksOk

`func (o *GetVMConfig200ResponseDataNet0) GetTrunksOk() (*string, bool)`

GetTrunksOk returns a tuple with the Trunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunks

`func (o *GetVMConfig200ResponseDataNet0) SetTrunks(v string)`

SetTrunks sets Trunks field to given value.

### HasTrunks

`func (o *GetVMConfig200ResponseDataNet0) HasTrunks() bool`

HasTrunks returns a boolean if a field has been set.

### GetVirtio

`func (o *GetVMConfig200ResponseDataNet0) GetVirtio() string`

GetVirtio returns the Virtio field if non-nil, zero value otherwise.

### GetVirtioOk

`func (o *GetVMConfig200ResponseDataNet0) GetVirtioOk() (*string, bool)`

GetVirtioOk returns a tuple with the Virtio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio

`func (o *GetVMConfig200ResponseDataNet0) SetVirtio(v string)`

SetVirtio sets Virtio field to given value.

### HasVirtio

`func (o *GetVMConfig200ResponseDataNet0) HasVirtio() bool`

HasVirtio returns a boolean if a field has been set.

### GetVmxnet3

`func (o *GetVMConfig200ResponseDataNet0) GetVmxnet3() string`

GetVmxnet3 returns the Vmxnet3 field if non-nil, zero value otherwise.

### GetVmxnet3Ok

`func (o *GetVMConfig200ResponseDataNet0) GetVmxnet3Ok() (*string, bool)`

GetVmxnet3Ok returns a tuple with the Vmxnet3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmxnet3

`func (o *GetVMConfig200ResponseDataNet0) SetVmxnet3(v string)`

SetVmxnet3 sets Vmxnet3 field to given value.

### HasVmxnet3

`func (o *GetVMConfig200ResponseDataNet0) HasVmxnet3() bool`

HasVmxnet3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


