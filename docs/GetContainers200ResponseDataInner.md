# GetContainers200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | Pointer to **float32** | Maximum usable CPUs. | [optional] 
**Disk** | Pointer to **int64** | Root disk image space-usage in bytes. | [optional] 
**Diskread** | Pointer to **int64** | The amount of bytes the guest read from it&#39;s block devices since the guest was started. (Note: This info is not available for all storage types.) | [optional] 
**Diskwrite** | Pointer to **int64** | The amount of bytes the guest wrote from it&#39;s block devices since the guest was started. (Note: This info is not available for all storage types.) | [optional] 
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

### NewGetContainers200ResponseDataInner

`func NewGetContainers200ResponseDataInner() *GetContainers200ResponseDataInner`

NewGetContainers200ResponseDataInner instantiates a new GetContainers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainers200ResponseDataInnerWithDefaults

`func NewGetContainers200ResponseDataInnerWithDefaults() *GetContainers200ResponseDataInner`

NewGetContainers200ResponseDataInnerWithDefaults instantiates a new GetContainers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpus

`func (o *GetContainers200ResponseDataInner) GetCpus() float32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *GetContainers200ResponseDataInner) GetCpusOk() (*float32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *GetContainers200ResponseDataInner) SetCpus(v float32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *GetContainers200ResponseDataInner) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDisk

`func (o *GetContainers200ResponseDataInner) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *GetContainers200ResponseDataInner) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *GetContainers200ResponseDataInner) SetDisk(v int64)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *GetContainers200ResponseDataInner) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetDiskread

`func (o *GetContainers200ResponseDataInner) GetDiskread() int64`

GetDiskread returns the Diskread field if non-nil, zero value otherwise.

### GetDiskreadOk

`func (o *GetContainers200ResponseDataInner) GetDiskreadOk() (*int64, bool)`

GetDiskreadOk returns a tuple with the Diskread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskread

`func (o *GetContainers200ResponseDataInner) SetDiskread(v int64)`

SetDiskread sets Diskread field to given value.

### HasDiskread

`func (o *GetContainers200ResponseDataInner) HasDiskread() bool`

HasDiskread returns a boolean if a field has been set.

### GetDiskwrite

`func (o *GetContainers200ResponseDataInner) GetDiskwrite() int64`

GetDiskwrite returns the Diskwrite field if non-nil, zero value otherwise.

### GetDiskwriteOk

`func (o *GetContainers200ResponseDataInner) GetDiskwriteOk() (*int64, bool)`

GetDiskwriteOk returns a tuple with the Diskwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskwrite

`func (o *GetContainers200ResponseDataInner) SetDiskwrite(v int64)`

SetDiskwrite sets Diskwrite field to given value.

### HasDiskwrite

`func (o *GetContainers200ResponseDataInner) HasDiskwrite() bool`

HasDiskwrite returns a boolean if a field has been set.

### GetLock

`func (o *GetContainers200ResponseDataInner) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetContainers200ResponseDataInner) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetContainers200ResponseDataInner) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetContainers200ResponseDataInner) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMaxdisk

`func (o *GetContainers200ResponseDataInner) GetMaxdisk() int64`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *GetContainers200ResponseDataInner) GetMaxdiskOk() (*int64, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *GetContainers200ResponseDataInner) SetMaxdisk(v int64)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *GetContainers200ResponseDataInner) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetMaxmem

`func (o *GetContainers200ResponseDataInner) GetMaxmem() int64`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *GetContainers200ResponseDataInner) GetMaxmemOk() (*int64, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *GetContainers200ResponseDataInner) SetMaxmem(v int64)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *GetContainers200ResponseDataInner) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetMaxswap

`func (o *GetContainers200ResponseDataInner) GetMaxswap() int64`

GetMaxswap returns the Maxswap field if non-nil, zero value otherwise.

### GetMaxswapOk

`func (o *GetContainers200ResponseDataInner) GetMaxswapOk() (*int64, bool)`

GetMaxswapOk returns a tuple with the Maxswap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxswap

`func (o *GetContainers200ResponseDataInner) SetMaxswap(v int64)`

SetMaxswap sets Maxswap field to given value.

### HasMaxswap

`func (o *GetContainers200ResponseDataInner) HasMaxswap() bool`

HasMaxswap returns a boolean if a field has been set.

### GetName

`func (o *GetContainers200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetContainers200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetContainers200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetContainers200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetin

`func (o *GetContainers200ResponseDataInner) GetNetin() int64`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *GetContainers200ResponseDataInner) GetNetinOk() (*int64, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *GetContainers200ResponseDataInner) SetNetin(v int64)`

SetNetin sets Netin field to given value.

### HasNetin

`func (o *GetContainers200ResponseDataInner) HasNetin() bool`

HasNetin returns a boolean if a field has been set.

### GetNetout

`func (o *GetContainers200ResponseDataInner) GetNetout() int64`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *GetContainers200ResponseDataInner) GetNetoutOk() (*int64, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *GetContainers200ResponseDataInner) SetNetout(v int64)`

SetNetout sets Netout field to given value.

### HasNetout

`func (o *GetContainers200ResponseDataInner) HasNetout() bool`

HasNetout returns a boolean if a field has been set.

### GetStatus

`func (o *GetContainers200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetContainers200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetContainers200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetContainers200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetContainers200ResponseDataInner) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetContainers200ResponseDataInner) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetContainers200ResponseDataInner) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetContainers200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *GetContainers200ResponseDataInner) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetContainers200ResponseDataInner) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetContainers200ResponseDataInner) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetContainers200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUptime

`func (o *GetContainers200ResponseDataInner) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GetContainers200ResponseDataInner) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GetContainers200ResponseDataInner) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GetContainers200ResponseDataInner) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVmid

`func (o *GetContainers200ResponseDataInner) GetVmid() int64`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *GetContainers200ResponseDataInner) GetVmidOk() (*int64, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *GetContainers200ResponseDataInner) SetVmid(v int64)`

SetVmid sets Vmid field to given value.

### HasVmid

`func (o *GetContainers200ResponseDataInner) HasVmid() bool`

HasVmid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


