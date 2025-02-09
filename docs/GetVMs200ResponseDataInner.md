# GetVMs200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | Pointer to **float32** | Maximum usable CPUs. | [optional] 
**Diskread** | Pointer to **int64** | The amount of bytes the guest read from it&#39;s block devices since the guest was started. (Note: This info is not available for all storage types.) | [optional] 
**Diskwrite** | Pointer to **int64** | The amount of bytes the guest wrote from it&#39;s block devices since the guest was started. (Note: This info is not available for all storage types.) | [optional] 
**Lock** | Pointer to **string** | The current config lock, if any. | [optional] 
**Maxdisk** | Pointer to **int64** | Root disk size in bytes. | [optional] 
**Maxmem** | Pointer to **int64** | Maximum memory in bytes. | [optional] 
**Name** | Pointer to **string** | VM (host)name. | [optional] 
**Netin** | Pointer to **int64** | The amount of traffic in bytes that was sent to the guest over the network since it was started. | [optional] 
**Netout** | Pointer to **int64** | The amount of traffic in bytes that was sent from the guest over the network since it was started. | [optional] 
**Pid** | Pointer to **int64** | PID of the QEMU process, if the VM is running. | [optional] 
**Qmpstatus** | Pointer to **string** | VM run state from the &#39;query-status&#39; QMP monitor command. | [optional] 
**RunningMachine** | Pointer to **string** | The currently running machine type (if running). | [optional] 
**RunningQemu** | Pointer to **string** | The QEMU version the VM is currently using (if running). | [optional] 
**Status** | Pointer to **string** | QEMU process status. | [optional] 
**Tags** | Pointer to **string** | The current configured tags, if any | [optional] 
**Template** | Pointer to **int32** | Determines if the guest is a template. | [optional] 
**Uptime** | Pointer to **int64** | Uptime in seconds. | [optional] 
**Vmid** | Pointer to **int64** | The (unique) ID of the VM. | [optional] 

## Methods

### NewGetVMs200ResponseDataInner

`func NewGetVMs200ResponseDataInner() *GetVMs200ResponseDataInner`

NewGetVMs200ResponseDataInner instantiates a new GetVMs200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMs200ResponseDataInnerWithDefaults

`func NewGetVMs200ResponseDataInnerWithDefaults() *GetVMs200ResponseDataInner`

NewGetVMs200ResponseDataInnerWithDefaults instantiates a new GetVMs200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpus

`func (o *GetVMs200ResponseDataInner) GetCpus() float32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *GetVMs200ResponseDataInner) GetCpusOk() (*float32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *GetVMs200ResponseDataInner) SetCpus(v float32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *GetVMs200ResponseDataInner) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDiskread

`func (o *GetVMs200ResponseDataInner) GetDiskread() int64`

GetDiskread returns the Diskread field if non-nil, zero value otherwise.

### GetDiskreadOk

`func (o *GetVMs200ResponseDataInner) GetDiskreadOk() (*int64, bool)`

GetDiskreadOk returns a tuple with the Diskread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskread

`func (o *GetVMs200ResponseDataInner) SetDiskread(v int64)`

SetDiskread sets Diskread field to given value.

### HasDiskread

`func (o *GetVMs200ResponseDataInner) HasDiskread() bool`

HasDiskread returns a boolean if a field has been set.

### GetDiskwrite

`func (o *GetVMs200ResponseDataInner) GetDiskwrite() int64`

GetDiskwrite returns the Diskwrite field if non-nil, zero value otherwise.

### GetDiskwriteOk

`func (o *GetVMs200ResponseDataInner) GetDiskwriteOk() (*int64, bool)`

GetDiskwriteOk returns a tuple with the Diskwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskwrite

`func (o *GetVMs200ResponseDataInner) SetDiskwrite(v int64)`

SetDiskwrite sets Diskwrite field to given value.

### HasDiskwrite

`func (o *GetVMs200ResponseDataInner) HasDiskwrite() bool`

HasDiskwrite returns a boolean if a field has been set.

### GetLock

`func (o *GetVMs200ResponseDataInner) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetVMs200ResponseDataInner) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetVMs200ResponseDataInner) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetVMs200ResponseDataInner) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMaxdisk

`func (o *GetVMs200ResponseDataInner) GetMaxdisk() int64`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *GetVMs200ResponseDataInner) GetMaxdiskOk() (*int64, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *GetVMs200ResponseDataInner) SetMaxdisk(v int64)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *GetVMs200ResponseDataInner) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetMaxmem

`func (o *GetVMs200ResponseDataInner) GetMaxmem() int64`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *GetVMs200ResponseDataInner) GetMaxmemOk() (*int64, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *GetVMs200ResponseDataInner) SetMaxmem(v int64)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *GetVMs200ResponseDataInner) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetName

`func (o *GetVMs200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVMs200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVMs200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVMs200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetin

`func (o *GetVMs200ResponseDataInner) GetNetin() int64`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *GetVMs200ResponseDataInner) GetNetinOk() (*int64, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *GetVMs200ResponseDataInner) SetNetin(v int64)`

SetNetin sets Netin field to given value.

### HasNetin

`func (o *GetVMs200ResponseDataInner) HasNetin() bool`

HasNetin returns a boolean if a field has been set.

### GetNetout

`func (o *GetVMs200ResponseDataInner) GetNetout() int64`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *GetVMs200ResponseDataInner) GetNetoutOk() (*int64, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *GetVMs200ResponseDataInner) SetNetout(v int64)`

SetNetout sets Netout field to given value.

### HasNetout

`func (o *GetVMs200ResponseDataInner) HasNetout() bool`

HasNetout returns a boolean if a field has been set.

### GetPid

`func (o *GetVMs200ResponseDataInner) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *GetVMs200ResponseDataInner) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *GetVMs200ResponseDataInner) SetPid(v int64)`

SetPid sets Pid field to given value.

### HasPid

`func (o *GetVMs200ResponseDataInner) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetQmpstatus

`func (o *GetVMs200ResponseDataInner) GetQmpstatus() string`

GetQmpstatus returns the Qmpstatus field if non-nil, zero value otherwise.

### GetQmpstatusOk

`func (o *GetVMs200ResponseDataInner) GetQmpstatusOk() (*string, bool)`

GetQmpstatusOk returns a tuple with the Qmpstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQmpstatus

`func (o *GetVMs200ResponseDataInner) SetQmpstatus(v string)`

SetQmpstatus sets Qmpstatus field to given value.

### HasQmpstatus

`func (o *GetVMs200ResponseDataInner) HasQmpstatus() bool`

HasQmpstatus returns a boolean if a field has been set.

### GetRunningMachine

`func (o *GetVMs200ResponseDataInner) GetRunningMachine() string`

GetRunningMachine returns the RunningMachine field if non-nil, zero value otherwise.

### GetRunningMachineOk

`func (o *GetVMs200ResponseDataInner) GetRunningMachineOk() (*string, bool)`

GetRunningMachineOk returns a tuple with the RunningMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMachine

`func (o *GetVMs200ResponseDataInner) SetRunningMachine(v string)`

SetRunningMachine sets RunningMachine field to given value.

### HasRunningMachine

`func (o *GetVMs200ResponseDataInner) HasRunningMachine() bool`

HasRunningMachine returns a boolean if a field has been set.

### GetRunningQemu

`func (o *GetVMs200ResponseDataInner) GetRunningQemu() string`

GetRunningQemu returns the RunningQemu field if non-nil, zero value otherwise.

### GetRunningQemuOk

`func (o *GetVMs200ResponseDataInner) GetRunningQemuOk() (*string, bool)`

GetRunningQemuOk returns a tuple with the RunningQemu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningQemu

`func (o *GetVMs200ResponseDataInner) SetRunningQemu(v string)`

SetRunningQemu sets RunningQemu field to given value.

### HasRunningQemu

`func (o *GetVMs200ResponseDataInner) HasRunningQemu() bool`

HasRunningQemu returns a boolean if a field has been set.

### GetStatus

`func (o *GetVMs200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVMs200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVMs200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVMs200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetVMs200ResponseDataInner) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetVMs200ResponseDataInner) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetVMs200ResponseDataInner) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetVMs200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *GetVMs200ResponseDataInner) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetVMs200ResponseDataInner) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetVMs200ResponseDataInner) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetVMs200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUptime

`func (o *GetVMs200ResponseDataInner) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GetVMs200ResponseDataInner) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GetVMs200ResponseDataInner) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GetVMs200ResponseDataInner) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVmid

`func (o *GetVMs200ResponseDataInner) GetVmid() int64`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *GetVMs200ResponseDataInner) GetVmidOk() (*int64, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *GetVMs200ResponseDataInner) SetVmid(v int64)`

SetVmid sets Vmid field to given value.

### HasVmid

`func (o *GetVMs200ResponseDataInner) HasVmid() bool`

HasVmid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


