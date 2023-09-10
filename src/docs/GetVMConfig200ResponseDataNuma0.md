# GetVMConfig200ResponseDataNuma0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | Pointer to **string** | CPUs accessing this NUMA node. | [optional] 
**Hostnodes** | Pointer to **string** | Host NUMA nodes to use. | [optional] 
**Memory** | Pointer to **float32** | Amount of memory this NUMA node provides. | [optional] 
**Policy** | Pointer to **string** | NUMA allocation policy. | [optional] 

## Methods

### NewGetVMConfig200ResponseDataNuma0

`func NewGetVMConfig200ResponseDataNuma0() *GetVMConfig200ResponseDataNuma0`

NewGetVMConfig200ResponseDataNuma0 instantiates a new GetVMConfig200ResponseDataNuma0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataNuma0WithDefaults

`func NewGetVMConfig200ResponseDataNuma0WithDefaults() *GetVMConfig200ResponseDataNuma0`

NewGetVMConfig200ResponseDataNuma0WithDefaults instantiates a new GetVMConfig200ResponseDataNuma0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpus

`func (o *GetVMConfig200ResponseDataNuma0) GetCpus() string`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *GetVMConfig200ResponseDataNuma0) GetCpusOk() (*string, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *GetVMConfig200ResponseDataNuma0) SetCpus(v string)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *GetVMConfig200ResponseDataNuma0) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetHostnodes

`func (o *GetVMConfig200ResponseDataNuma0) GetHostnodes() string`

GetHostnodes returns the Hostnodes field if non-nil, zero value otherwise.

### GetHostnodesOk

`func (o *GetVMConfig200ResponseDataNuma0) GetHostnodesOk() (*string, bool)`

GetHostnodesOk returns a tuple with the Hostnodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnodes

`func (o *GetVMConfig200ResponseDataNuma0) SetHostnodes(v string)`

SetHostnodes sets Hostnodes field to given value.

### HasHostnodes

`func (o *GetVMConfig200ResponseDataNuma0) HasHostnodes() bool`

HasHostnodes returns a boolean if a field has been set.

### GetMemory

`func (o *GetVMConfig200ResponseDataNuma0) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GetVMConfig200ResponseDataNuma0) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GetVMConfig200ResponseDataNuma0) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GetVMConfig200ResponseDataNuma0) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetPolicy

`func (o *GetVMConfig200ResponseDataNuma0) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetVMConfig200ResponseDataNuma0) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetVMConfig200ResponseDataNuma0) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetVMConfig200ResponseDataNuma0) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


