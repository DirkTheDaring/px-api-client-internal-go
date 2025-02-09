# GetCurrentContainerStatus200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | Pointer to **float32** | Maximum usable CPUs. | [optional] 
**Disk** | Pointer to **int64** | Root disk image space-usage in bytes. | [optional] 
**Diskread** | Pointer to **int64** | The amount of bytes the guest read from it&#39;s block devices since the guest was started. (Note: This info is not available for all storage types.) | [optional] 
**Diskwrite** | Pointer to **int64** | The amount of bytes the guest wrote from it&#39;s block devices since the guest was started. (Note: This info is not available for all storage types.) | [optional] 
**Ha** | Pointer to **map[string]interface{}** | HA manager service status. | [optional] 
**Lock** | Pointer to **string** | The current config lock, if any. | [optional] 
**Maxdisk** | Pointer to **int64** | Root disk image size in bytes. | [optional] 
**Maxmem** | Pointer to **int64** | Maximum memory in bytes. | [optional] 
**Maxswap** | Pointer to **int64** | Maximum SWAP memory in bytes. | [optional] 
**Name** | Pointer to **string** | Container name. | [optional] 
**Netin** | Pointer to **int64** | The amount of traffic in bytes that was sent to the guest over the network since it was started. | [optional] 
**Netout** | Pointer to **int64** | The amount of traffic in bytes that was sent from the guest over the network since it was started. | [optional] 
**Status** | Pointer to **string** | LXC Container status. | [optional] 
**Tags** | Pointer to **string** | The current configured tags, if any. | [optional] 
**Template** | Pointer to **int32** | Determines if the guest is a template. | [optional] 
**Uptime** | Pointer to **int64** | Uptime in seconds. | [optional] 
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

### GetDisk

`func (o *GetCurrentContainerStatus200ResponseData) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *GetCurrentContainerStatus200ResponseData) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *GetCurrentContainerStatus200ResponseData) SetDisk(v int64)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *GetCurrentContainerStatus200ResponseData) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetDiskread

`func (o *GetCurrentContainerStatus200ResponseData) GetDiskread() int64`

GetDiskread returns the Diskread field if non-nil, zero value otherwise.

### GetDiskreadOk

`func (o *GetCurrentContainerStatus200ResponseData) GetDiskreadOk() (*int64, bool)`

GetDiskreadOk returns a tuple with the Diskread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskread

`func (o *GetCurrentContainerStatus200ResponseData) SetDiskread(v int64)`

SetDiskread sets Diskread field to given value.

### HasDiskread

`func (o *GetCurrentContainerStatus200ResponseData) HasDiskread() bool`

HasDiskread returns a boolean if a field has been set.

### GetDiskwrite

`func (o *GetCurrentContainerStatus200ResponseData) GetDiskwrite() int64`

GetDiskwrite returns the Diskwrite field if non-nil, zero value otherwise.

### GetDiskwriteOk

`func (o *GetCurrentContainerStatus200ResponseData) GetDiskwriteOk() (*int64, bool)`

GetDiskwriteOk returns a tuple with the Diskwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskwrite

`func (o *GetCurrentContainerStatus200ResponseData) SetDiskwrite(v int64)`

SetDiskwrite sets Diskwrite field to given value.

### HasDiskwrite

`func (o *GetCurrentContainerStatus200ResponseData) HasDiskwrite() bool`

HasDiskwrite returns a boolean if a field has been set.

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

### GetNetin

`func (o *GetCurrentContainerStatus200ResponseData) GetNetin() int64`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *GetCurrentContainerStatus200ResponseData) GetNetinOk() (*int64, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *GetCurrentContainerStatus200ResponseData) SetNetin(v int64)`

SetNetin sets Netin field to given value.

### HasNetin

`func (o *GetCurrentContainerStatus200ResponseData) HasNetin() bool`

HasNetin returns a boolean if a field has been set.

### GetNetout

`func (o *GetCurrentContainerStatus200ResponseData) GetNetout() int64`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *GetCurrentContainerStatus200ResponseData) GetNetoutOk() (*int64, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *GetCurrentContainerStatus200ResponseData) SetNetout(v int64)`

SetNetout sets Netout field to given value.

### HasNetout

`func (o *GetCurrentContainerStatus200ResponseData) HasNetout() bool`

HasNetout returns a boolean if a field has been set.

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

### GetTemplate

`func (o *GetCurrentContainerStatus200ResponseData) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetCurrentContainerStatus200ResponseData) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetCurrentContainerStatus200ResponseData) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetCurrentContainerStatus200ResponseData) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

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


