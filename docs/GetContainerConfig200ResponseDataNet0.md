# GetContainerConfig200ResponseDataNet0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bridge** | Pointer to **string** | Bridge to attach the network device to. | [optional] 
**Firewall** | Pointer to **bool** | Controls whether this interface&#39;s firewall rules should be used. | [optional] 
**Gw** | Pointer to **string** | Default gateway for IPv4 traffic. | [optional] 
**Gw6** | Pointer to **string** | Default gateway for IPv6 traffic. | [optional] 
**Hwaddr** | Pointer to **string** | The interface MAC address. This is dynamically allocated by default, but you can set that statically if needed, for example to always have the same link-local IPv6 address. (lxc.network.hwaddr) | [optional] 
**Ip** | Pointer to **string** | IPv4 address in CIDR format. | [optional] 
**Ip6** | Pointer to **string** | IPv6 address in CIDR format. | [optional] 
**LinkDown** | Pointer to **bool** | Whether this interface should be disconnected (like pulling the plug). | [optional] 
**Mtu** | Pointer to **int64** | Maximum transfer unit of the interface. (lxc.network.mtu) | [optional] 
**Name** | Pointer to **string** | Name of the network device as seen from inside the container. (lxc.network.name) | [optional] 
**Rate** | Pointer to **float32** | Apply rate limiting to the interface | [optional] 
**Tag** | Pointer to **int64** | VLAN tag for this interface. | [optional] 
**Trunks** | Pointer to **string** | VLAN ids to pass through the interface | [optional] 
**Type** | Pointer to **string** | Network interface type. | [optional] 

## Methods

### NewGetContainerConfig200ResponseDataNet0

`func NewGetContainerConfig200ResponseDataNet0() *GetContainerConfig200ResponseDataNet0`

NewGetContainerConfig200ResponseDataNet0 instantiates a new GetContainerConfig200ResponseDataNet0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerConfig200ResponseDataNet0WithDefaults

`func NewGetContainerConfig200ResponseDataNet0WithDefaults() *GetContainerConfig200ResponseDataNet0`

NewGetContainerConfig200ResponseDataNet0WithDefaults instantiates a new GetContainerConfig200ResponseDataNet0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridge

`func (o *GetContainerConfig200ResponseDataNet0) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *GetContainerConfig200ResponseDataNet0) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *GetContainerConfig200ResponseDataNet0) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *GetContainerConfig200ResponseDataNet0) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetFirewall

`func (o *GetContainerConfig200ResponseDataNet0) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *GetContainerConfig200ResponseDataNet0) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *GetContainerConfig200ResponseDataNet0) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *GetContainerConfig200ResponseDataNet0) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetGw

`func (o *GetContainerConfig200ResponseDataNet0) GetGw() string`

GetGw returns the Gw field if non-nil, zero value otherwise.

### GetGwOk

`func (o *GetContainerConfig200ResponseDataNet0) GetGwOk() (*string, bool)`

GetGwOk returns a tuple with the Gw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw

`func (o *GetContainerConfig200ResponseDataNet0) SetGw(v string)`

SetGw sets Gw field to given value.

### HasGw

`func (o *GetContainerConfig200ResponseDataNet0) HasGw() bool`

HasGw returns a boolean if a field has been set.

### GetGw6

`func (o *GetContainerConfig200ResponseDataNet0) GetGw6() string`

GetGw6 returns the Gw6 field if non-nil, zero value otherwise.

### GetGw6Ok

`func (o *GetContainerConfig200ResponseDataNet0) GetGw6Ok() (*string, bool)`

GetGw6Ok returns a tuple with the Gw6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw6

`func (o *GetContainerConfig200ResponseDataNet0) SetGw6(v string)`

SetGw6 sets Gw6 field to given value.

### HasGw6

`func (o *GetContainerConfig200ResponseDataNet0) HasGw6() bool`

HasGw6 returns a boolean if a field has been set.

### GetHwaddr

`func (o *GetContainerConfig200ResponseDataNet0) GetHwaddr() string`

GetHwaddr returns the Hwaddr field if non-nil, zero value otherwise.

### GetHwaddrOk

`func (o *GetContainerConfig200ResponseDataNet0) GetHwaddrOk() (*string, bool)`

GetHwaddrOk returns a tuple with the Hwaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwaddr

`func (o *GetContainerConfig200ResponseDataNet0) SetHwaddr(v string)`

SetHwaddr sets Hwaddr field to given value.

### HasHwaddr

`func (o *GetContainerConfig200ResponseDataNet0) HasHwaddr() bool`

HasHwaddr returns a boolean if a field has been set.

### GetIp

`func (o *GetContainerConfig200ResponseDataNet0) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetContainerConfig200ResponseDataNet0) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetContainerConfig200ResponseDataNet0) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetContainerConfig200ResponseDataNet0) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *GetContainerConfig200ResponseDataNet0) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *GetContainerConfig200ResponseDataNet0) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *GetContainerConfig200ResponseDataNet0) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *GetContainerConfig200ResponseDataNet0) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetLinkDown

`func (o *GetContainerConfig200ResponseDataNet0) GetLinkDown() bool`

GetLinkDown returns the LinkDown field if non-nil, zero value otherwise.

### GetLinkDownOk

`func (o *GetContainerConfig200ResponseDataNet0) GetLinkDownOk() (*bool, bool)`

GetLinkDownOk returns a tuple with the LinkDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDown

`func (o *GetContainerConfig200ResponseDataNet0) SetLinkDown(v bool)`

SetLinkDown sets LinkDown field to given value.

### HasLinkDown

`func (o *GetContainerConfig200ResponseDataNet0) HasLinkDown() bool`

HasLinkDown returns a boolean if a field has been set.

### GetMtu

`func (o *GetContainerConfig200ResponseDataNet0) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *GetContainerConfig200ResponseDataNet0) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *GetContainerConfig200ResponseDataNet0) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *GetContainerConfig200ResponseDataNet0) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *GetContainerConfig200ResponseDataNet0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetContainerConfig200ResponseDataNet0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetContainerConfig200ResponseDataNet0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetContainerConfig200ResponseDataNet0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRate

`func (o *GetContainerConfig200ResponseDataNet0) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetContainerConfig200ResponseDataNet0) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetContainerConfig200ResponseDataNet0) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetContainerConfig200ResponseDataNet0) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTag

`func (o *GetContainerConfig200ResponseDataNet0) GetTag() int64`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetContainerConfig200ResponseDataNet0) GetTagOk() (*int64, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetContainerConfig200ResponseDataNet0) SetTag(v int64)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetContainerConfig200ResponseDataNet0) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTrunks

`func (o *GetContainerConfig200ResponseDataNet0) GetTrunks() string`

GetTrunks returns the Trunks field if non-nil, zero value otherwise.

### GetTrunksOk

`func (o *GetContainerConfig200ResponseDataNet0) GetTrunksOk() (*string, bool)`

GetTrunksOk returns a tuple with the Trunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunks

`func (o *GetContainerConfig200ResponseDataNet0) SetTrunks(v string)`

SetTrunks sets Trunks field to given value.

### HasTrunks

`func (o *GetContainerConfig200ResponseDataNet0) HasTrunks() bool`

HasTrunks returns a boolean if a field has been set.

### GetType

`func (o *GetContainerConfig200ResponseDataNet0) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetContainerConfig200ResponseDataNet0) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetContainerConfig200ResponseDataNet0) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetContainerConfig200ResponseDataNet0) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


