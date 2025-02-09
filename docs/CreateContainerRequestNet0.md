# CreateContainerRequestNet0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bridge** | Pointer to **string** |  | [optional] 
**Firewall** | Pointer to **bool** |  | [optional] 
**Gw** | Pointer to **string** |  | [optional] 
**Gw6** | Pointer to **string** |  | [optional] 
**Hwaddr** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Ip6** | Pointer to **string** |  | [optional] 
**LinkDown** | Pointer to **bool** |  | [optional] 
**Mtu** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Rate** | Pointer to **float32** |  | [optional] 
**Tag** | Pointer to **int64** |  | [optional] 
**Trunks** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateContainerRequestNet0

`func NewCreateContainerRequestNet0(name string, ) *CreateContainerRequestNet0`

NewCreateContainerRequestNet0 instantiates a new CreateContainerRequestNet0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerRequestNet0WithDefaults

`func NewCreateContainerRequestNet0WithDefaults() *CreateContainerRequestNet0`

NewCreateContainerRequestNet0WithDefaults instantiates a new CreateContainerRequestNet0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridge

`func (o *CreateContainerRequestNet0) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *CreateContainerRequestNet0) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *CreateContainerRequestNet0) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *CreateContainerRequestNet0) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetFirewall

`func (o *CreateContainerRequestNet0) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *CreateContainerRequestNet0) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *CreateContainerRequestNet0) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *CreateContainerRequestNet0) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetGw

`func (o *CreateContainerRequestNet0) GetGw() string`

GetGw returns the Gw field if non-nil, zero value otherwise.

### GetGwOk

`func (o *CreateContainerRequestNet0) GetGwOk() (*string, bool)`

GetGwOk returns a tuple with the Gw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw

`func (o *CreateContainerRequestNet0) SetGw(v string)`

SetGw sets Gw field to given value.

### HasGw

`func (o *CreateContainerRequestNet0) HasGw() bool`

HasGw returns a boolean if a field has been set.

### GetGw6

`func (o *CreateContainerRequestNet0) GetGw6() string`

GetGw6 returns the Gw6 field if non-nil, zero value otherwise.

### GetGw6Ok

`func (o *CreateContainerRequestNet0) GetGw6Ok() (*string, bool)`

GetGw6Ok returns a tuple with the Gw6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw6

`func (o *CreateContainerRequestNet0) SetGw6(v string)`

SetGw6 sets Gw6 field to given value.

### HasGw6

`func (o *CreateContainerRequestNet0) HasGw6() bool`

HasGw6 returns a boolean if a field has been set.

### GetHwaddr

`func (o *CreateContainerRequestNet0) GetHwaddr() string`

GetHwaddr returns the Hwaddr field if non-nil, zero value otherwise.

### GetHwaddrOk

`func (o *CreateContainerRequestNet0) GetHwaddrOk() (*string, bool)`

GetHwaddrOk returns a tuple with the Hwaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwaddr

`func (o *CreateContainerRequestNet0) SetHwaddr(v string)`

SetHwaddr sets Hwaddr field to given value.

### HasHwaddr

`func (o *CreateContainerRequestNet0) HasHwaddr() bool`

HasHwaddr returns a boolean if a field has been set.

### GetIp

`func (o *CreateContainerRequestNet0) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateContainerRequestNet0) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateContainerRequestNet0) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateContainerRequestNet0) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *CreateContainerRequestNet0) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *CreateContainerRequestNet0) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *CreateContainerRequestNet0) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *CreateContainerRequestNet0) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetLinkDown

`func (o *CreateContainerRequestNet0) GetLinkDown() bool`

GetLinkDown returns the LinkDown field if non-nil, zero value otherwise.

### GetLinkDownOk

`func (o *CreateContainerRequestNet0) GetLinkDownOk() (*bool, bool)`

GetLinkDownOk returns a tuple with the LinkDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDown

`func (o *CreateContainerRequestNet0) SetLinkDown(v bool)`

SetLinkDown sets LinkDown field to given value.

### HasLinkDown

`func (o *CreateContainerRequestNet0) HasLinkDown() bool`

HasLinkDown returns a boolean if a field has been set.

### GetMtu

`func (o *CreateContainerRequestNet0) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateContainerRequestNet0) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateContainerRequestNet0) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateContainerRequestNet0) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *CreateContainerRequestNet0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContainerRequestNet0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContainerRequestNet0) SetName(v string)`

SetName sets Name field to given value.


### GetRate

`func (o *CreateContainerRequestNet0) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CreateContainerRequestNet0) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CreateContainerRequestNet0) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *CreateContainerRequestNet0) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTag

`func (o *CreateContainerRequestNet0) GetTag() int64`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateContainerRequestNet0) GetTagOk() (*int64, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateContainerRequestNet0) SetTag(v int64)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateContainerRequestNet0) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTrunks

`func (o *CreateContainerRequestNet0) GetTrunks() string`

GetTrunks returns the Trunks field if non-nil, zero value otherwise.

### GetTrunksOk

`func (o *CreateContainerRequestNet0) GetTrunksOk() (*string, bool)`

GetTrunksOk returns a tuple with the Trunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunks

`func (o *CreateContainerRequestNet0) SetTrunks(v string)`

SetTrunks sets Trunks field to given value.

### HasTrunks

`func (o *CreateContainerRequestNet0) HasTrunks() bool`

HasTrunks returns a boolean if a field has been set.

### GetType

`func (o *CreateContainerRequestNet0) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateContainerRequestNet0) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateContainerRequestNet0) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateContainerRequestNet0) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


