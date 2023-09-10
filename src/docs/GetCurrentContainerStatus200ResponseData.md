# GetCurrentContainerStatus200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | Pointer to **float32** | Maximum usable CPUs. | [optional] 
**Ha** | Pointer to **map[string]interface{}** | HA manager service status. | [optional] 
**Lock** | Pointer to **string** | The current config lock, if any. | [optional] 
**Maxdisk** | Pointer to **int64** | Root disk size in bytes. | [optional] 
**Maxmem** | Pointer to **int64** | Maximum memory in bytes. | [optional] 
**Maxswap** | Pointer to **int64** | Maximum SWAP memory in bytes. | [optional] 
**Name** | Pointer to **string** | Container name. | [optional] 
**Status** | Pointer to **string** | LXC Container status. | [optional] 
**Tags** | Pointer to **string** | The current configured tags, if any. | [optional] 
**Uptime** | Pointer to **int64** | Uptime. | [optional] 
**Vmid** | Pointer to **int64** | The (unique) ID of the VM. | [optional] 

## Methods

### NewGetCurrentContainerStatus200ResponseData

`func NewGetCurrentContainerStatus200ResponseData() *GetCurrentContainerStatus200ResponseData`

NewGetCurrentContainerStatus200ResponseData instantiates a new GetCurrentContainerStatus200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentContainerStatus200ResponseDataWithDefaults

`func NewGetCurrentContainerStatus200ResponseDataWithDefaults() *GetCurrentContainerStatus200ResponseData`

NewGetCurrentContainerStatus200ResponseDataWithDefaults instantiates a new GetCurrentContainerStatus200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpus

`func (o *GetCurrentContainerStatus200ResponseData) GetCpus() float32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *GetCurrentContainerStatus200ResponseData) GetCpusOk() (*float32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *GetCurrentContainerStatus200ResponseData) SetCpus(v float32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *GetCurrentContainerStatus200ResponseData) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetHa

`func (o *GetCurrentContainerStatus200ResponseData) GetHa() map[string]interface{}`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *GetCurrentContainerStatus200ResponseData) GetHaOk() (*map[string]interface{}, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *GetCurrentContainerStatus200ResponseData) SetHa(v map[string]interface{})`

SetHa sets Ha field to given value.

### HasHa

`func (o *GetCurrentContainerStatus200ResponseData) HasHa() bool`

HasHa returns a boolean if a field has been set.

### GetLock

`func (o *GetCurrentContainerStatus200ResponseData) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetCurrentContainerStatus200ResponseData) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetCurrentContainerStatus200ResponseData) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetCurrentContainerStatus200ResponseData) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMaxdisk

`func (o *GetCurrentContainerStatus200ResponseData) GetMaxdisk() int64`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *GetCurrentContainerStatus200ResponseData) GetMaxdiskOk() (*int64, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *GetCurrentContainerStatus200ResponseData) SetMaxdisk(v int64)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *GetCurrentContainerStatus200ResponseData) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetMaxmem

`func (o *GetCurrentContainerStatus200ResponseData) GetMaxmem() int64`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *GetCurrentContainerStatus200ResponseData) GetMaxmemOk() (*int64, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *GetCurrentContainerStatus200ResponseData) SetMaxmem(v int64)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *GetCurrentContainerStatus200ResponseData) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetMaxswap

`func (o *GetCurrentContainerStatus200ResponseData) GetMaxswap() int64`

GetMaxswap returns the Maxswap field if non-nil, zero value otherwise.

### GetMaxswapOk

`func (o *GetCurrentContainerStatus200ResponseData) GetMaxswapOk() (*int64, bool)`

GetMaxswapOk returns a tuple with the Maxswap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxswap

`func (o *GetCurrentContainerStatus200ResponseData) SetMaxswap(v int64)`

SetMaxswap sets Maxswap field to given value.

### HasMaxswap

`func (o *GetCurrentContainerStatus200ResponseData) HasMaxswap() bool`

HasMaxswap returns a boolean if a field has been set.

### GetName

`func (o *GetCurrentContainerStatus200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCurrentContainerStatus200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCurrentContainerStatus200ResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCurrentContainerStatus200ResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GetCurrentContainerStatus200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCurrentContainerStatus200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCurrentContainerStatus200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCurrentContainerStatus200ResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetCurrentContainerStatus200ResponseData) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetCurrentContainerStatus200ResponseData) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetCurrentContainerStatus200ResponseData) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetCurrentContainerStatus200ResponseData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUptime

`func (o *GetCurrentContainerStatus200ResponseData) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GetCurrentContainerStatus200ResponseData) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GetCurrentContainerStatus200ResponseData) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GetCurrentContainerStatus200ResponseData) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVmid

`func (o *GetCurrentContainerStatus200ResponseData) GetVmid() int64`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *GetCurrentContainerStatus200ResponseData) GetVmidOk() (*int64, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *GetCurrentContainerStatus200ResponseData) SetVmid(v int64)`

SetVmid sets Vmid field to given value.

### HasVmid

`func (o *GetCurrentContainerStatus200ResponseData) HasVmid() bool`

HasVmid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


