# CreateVMRequestNet0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bridge** | Pointer to **string** |  | [optional] 
**Firewall** | Pointer to **bool** |  | [optional] 
**LinkDown** | Pointer to **bool** |  | [optional] 
**Macaddr** | Pointer to **string** |  | [optional] 
**Model** | **string** |  | 
**Mtu** | Pointer to **int64** |  | [optional] 
**Queues** | Pointer to **int64** |  | [optional] 
**Rate** | Pointer to **float32** |  | [optional] 
**Tag** | Pointer to **int64** |  | [optional] 
**Trunks** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVMRequestNet0

`func NewCreateVMRequestNet0(model string, ) *CreateVMRequestNet0`

NewCreateVMRequestNet0 instantiates a new CreateVMRequestNet0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestNet0WithDefaults

`func NewCreateVMRequestNet0WithDefaults() *CreateVMRequestNet0`

NewCreateVMRequestNet0WithDefaults instantiates a new CreateVMRequestNet0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridge

`func (o *CreateVMRequestNet0) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *CreateVMRequestNet0) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *CreateVMRequestNet0) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *CreateVMRequestNet0) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetFirewall

`func (o *CreateVMRequestNet0) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *CreateVMRequestNet0) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *CreateVMRequestNet0) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *CreateVMRequestNet0) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetLinkDown

`func (o *CreateVMRequestNet0) GetLinkDown() bool`

GetLinkDown returns the LinkDown field if non-nil, zero value otherwise.

### GetLinkDownOk

`func (o *CreateVMRequestNet0) GetLinkDownOk() (*bool, bool)`

GetLinkDownOk returns a tuple with the LinkDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDown

`func (o *CreateVMRequestNet0) SetLinkDown(v bool)`

SetLinkDown sets LinkDown field to given value.

### HasLinkDown

`func (o *CreateVMRequestNet0) HasLinkDown() bool`

HasLinkDown returns a boolean if a field has been set.

### GetMacaddr

`func (o *CreateVMRequestNet0) GetMacaddr() string`

GetMacaddr returns the Macaddr field if non-nil, zero value otherwise.

### GetMacaddrOk

`func (o *CreateVMRequestNet0) GetMacaddrOk() (*string, bool)`

GetMacaddrOk returns a tuple with the Macaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacaddr

`func (o *CreateVMRequestNet0) SetMacaddr(v string)`

SetMacaddr sets Macaddr field to given value.

### HasMacaddr

`func (o *CreateVMRequestNet0) HasMacaddr() bool`

HasMacaddr returns a boolean if a field has been set.

### GetModel

`func (o *CreateVMRequestNet0) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateVMRequestNet0) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateVMRequestNet0) SetModel(v string)`

SetModel sets Model field to given value.


### GetMtu

`func (o *CreateVMRequestNet0) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateVMRequestNet0) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateVMRequestNet0) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateVMRequestNet0) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetQueues

`func (o *CreateVMRequestNet0) GetQueues() int64`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *CreateVMRequestNet0) GetQueuesOk() (*int64, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *CreateVMRequestNet0) SetQueues(v int64)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *CreateVMRequestNet0) HasQueues() bool`

HasQueues returns a boolean if a field has been set.

### GetRate

`func (o *CreateVMRequestNet0) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CreateVMRequestNet0) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CreateVMRequestNet0) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *CreateVMRequestNet0) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTag

`func (o *CreateVMRequestNet0) GetTag() int64`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateVMRequestNet0) GetTagOk() (*int64, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateVMRequestNet0) SetTag(v int64)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateVMRequestNet0) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTrunks

`func (o *CreateVMRequestNet0) GetTrunks() string`

GetTrunks returns the Trunks field if non-nil, zero value otherwise.

### GetTrunksOk

`func (o *CreateVMRequestNet0) GetTrunksOk() (*string, bool)`

GetTrunksOk returns a tuple with the Trunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunks

`func (o *CreateVMRequestNet0) SetTrunks(v string)`

SetTrunks sets Trunks field to given value.

### HasTrunks

`func (o *CreateVMRequestNet0) HasTrunks() bool`

HasTrunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


