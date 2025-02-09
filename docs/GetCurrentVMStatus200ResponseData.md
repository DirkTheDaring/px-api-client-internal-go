# GetCurrentVMStatus200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **int32** | QEMU Guest Agent is enabled in config. | [optional] 
**Clipboard** | Pointer to **string** | Enable a specific clipboard. If not set, depending on the display type the SPICE one will be added. | [optional] 
**Cpus** | Pointer to **float32** | Maximum usable CPUs. | [optional] 
**Diskread** | Pointer to **int64** | The amount of bytes the guest read from it&#39;s block devices since the guest was started. (Note: This info is not available for all storage types.) | [optional] 
**Diskwrite** | Pointer to **int64** | The amount of bytes the guest wrote from it&#39;s block devices since the guest was started. (Note: This info is not available for all storage types.) | [optional] 
**Ha** | Pointer to **map[string]interface{}** | HA manager service status. | [optional] 
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
**Spice** | Pointer to **int32** | QEMU VGA configuration supports spice. | [optional] 
**Status** | Pointer to **string** | QEMU process status. | [optional] 
**Tags** | Pointer to **string** | The current configured tags, if any | [optional] 
**Template** | Pointer to **int32** | Determines if the guest is a template. | [optional] 
**Uptime** | Pointer to **int64** | Uptime in seconds. | [optional] 
**Vmid** | Pointer to **int64** | The (unique) ID of the VM. | [optional] 

## Methods

### NewGetCurrentVMStatus200ResponseData

`func NewGetCurrentVMStatus200ResponseData() *GetCurrentVMStatus200ResponseData`

NewGetCurrentVMStatus200ResponseData instantiates a new GetCurrentVMStatus200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentVMStatus200ResponseDataWithDefaults

`func NewGetCurrentVMStatus200ResponseDataWithDefaults() *GetCurrentVMStatus200ResponseData`

NewGetCurrentVMStatus200ResponseDataWithDefaults instantiates a new GetCurrentVMStatus200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *GetCurrentVMStatus200ResponseData) GetAgent() int32`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *GetCurrentVMStatus200ResponseData) GetAgentOk() (*int32, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *GetCurrentVMStatus200ResponseData) SetAgent(v int32)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *GetCurrentVMStatus200ResponseData) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetClipboard

`func (o *GetCurrentVMStatus200ResponseData) GetClipboard() string`

GetClipboard returns the Clipboard field if non-nil, zero value otherwise.

### GetClipboardOk

`func (o *GetCurrentVMStatus200ResponseData) GetClipboardOk() (*string, bool)`

GetClipboardOk returns a tuple with the Clipboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipboard

`func (o *GetCurrentVMStatus200ResponseData) SetClipboard(v string)`

SetClipboard sets Clipboard field to given value.

### HasClipboard

`func (o *GetCurrentVMStatus200ResponseData) HasClipboard() bool`

HasClipboard returns a boolean if a field has been set.

### GetCpus

`func (o *GetCurrentVMStatus200ResponseData) GetCpus() float32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *GetCurrentVMStatus200ResponseData) GetCpusOk() (*float32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *GetCurrentVMStatus200ResponseData) SetCpus(v float32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *GetCurrentVMStatus200ResponseData) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDiskread

`func (o *GetCurrentVMStatus200ResponseData) GetDiskread() int64`

GetDiskread returns the Diskread field if non-nil, zero value otherwise.

### GetDiskreadOk

`func (o *GetCurrentVMStatus200ResponseData) GetDiskreadOk() (*int64, bool)`

GetDiskreadOk returns a tuple with the Diskread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskread

`func (o *GetCurrentVMStatus200ResponseData) SetDiskread(v int64)`

SetDiskread sets Diskread field to given value.

### HasDiskread

`func (o *GetCurrentVMStatus200ResponseData) HasDiskread() bool`

HasDiskread returns a boolean if a field has been set.

### GetDiskwrite

`func (o *GetCurrentVMStatus200ResponseData) GetDiskwrite() int64`

GetDiskwrite returns the Diskwrite field if non-nil, zero value otherwise.

### GetDiskwriteOk

`func (o *GetCurrentVMStatus200ResponseData) GetDiskwriteOk() (*int64, bool)`

GetDiskwriteOk returns a tuple with the Diskwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskwrite

`func (o *GetCurrentVMStatus200ResponseData) SetDiskwrite(v int64)`

SetDiskwrite sets Diskwrite field to given value.

### HasDiskwrite

`func (o *GetCurrentVMStatus200ResponseData) HasDiskwrite() bool`

HasDiskwrite returns a boolean if a field has been set.

### GetHa

`func (o *GetCurrentVMStatus200ResponseData) GetHa() map[string]interface{}`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *GetCurrentVMStatus200ResponseData) GetHaOk() (*map[string]interface{}, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *GetCurrentVMStatus200ResponseData) SetHa(v map[string]interface{})`

SetHa sets Ha field to given value.

### HasHa

`func (o *GetCurrentVMStatus200ResponseData) HasHa() bool`

HasHa returns a boolean if a field has been set.

### GetLock

`func (o *GetCurrentVMStatus200ResponseData) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetCurrentVMStatus200ResponseData) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetCurrentVMStatus200ResponseData) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetCurrentVMStatus200ResponseData) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMaxdisk

`func (o *GetCurrentVMStatus200ResponseData) GetMaxdisk() int64`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *GetCurrentVMStatus200ResponseData) GetMaxdiskOk() (*int64, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *GetCurrentVMStatus200ResponseData) SetMaxdisk(v int64)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *GetCurrentVMStatus200ResponseData) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetMaxmem

`func (o *GetCurrentVMStatus200ResponseData) GetMaxmem() int64`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *GetCurrentVMStatus200ResponseData) GetMaxmemOk() (*int64, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *GetCurrentVMStatus200ResponseData) SetMaxmem(v int64)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *GetCurrentVMStatus200ResponseData) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetName

`func (o *GetCurrentVMStatus200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCurrentVMStatus200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCurrentVMStatus200ResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCurrentVMStatus200ResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetin

`func (o *GetCurrentVMStatus200ResponseData) GetNetin() int64`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *GetCurrentVMStatus200ResponseData) GetNetinOk() (*int64, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *GetCurrentVMStatus200ResponseData) SetNetin(v int64)`

SetNetin sets Netin field to given value.

### HasNetin

`func (o *GetCurrentVMStatus200ResponseData) HasNetin() bool`

HasNetin returns a boolean if a field has been set.

### GetNetout

`func (o *GetCurrentVMStatus200ResponseData) GetNetout() int64`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *GetCurrentVMStatus200ResponseData) GetNetoutOk() (*int64, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *GetCurrentVMStatus200ResponseData) SetNetout(v int64)`

SetNetout sets Netout field to given value.

### HasNetout

`func (o *GetCurrentVMStatus200ResponseData) HasNetout() bool`

HasNetout returns a boolean if a field has been set.

### GetPid

`func (o *GetCurrentVMStatus200ResponseData) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *GetCurrentVMStatus200ResponseData) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *GetCurrentVMStatus200ResponseData) SetPid(v int64)`

SetPid sets Pid field to given value.

### HasPid

`func (o *GetCurrentVMStatus200ResponseData) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetQmpstatus

`func (o *GetCurrentVMStatus200ResponseData) GetQmpstatus() string`

GetQmpstatus returns the Qmpstatus field if non-nil, zero value otherwise.

### GetQmpstatusOk

`func (o *GetCurrentVMStatus200ResponseData) GetQmpstatusOk() (*string, bool)`

GetQmpstatusOk returns a tuple with the Qmpstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQmpstatus

`func (o *GetCurrentVMStatus200ResponseData) SetQmpstatus(v string)`

SetQmpstatus sets Qmpstatus field to given value.

### HasQmpstatus

`func (o *GetCurrentVMStatus200ResponseData) HasQmpstatus() bool`

HasQmpstatus returns a boolean if a field has been set.

### GetRunningMachine

`func (o *GetCurrentVMStatus200ResponseData) GetRunningMachine() string`

GetRunningMachine returns the RunningMachine field if non-nil, zero value otherwise.

### GetRunningMachineOk

`func (o *GetCurrentVMStatus200ResponseData) GetRunningMachineOk() (*string, bool)`

GetRunningMachineOk returns a tuple with the RunningMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMachine

`func (o *GetCurrentVMStatus200ResponseData) SetRunningMachine(v string)`

SetRunningMachine sets RunningMachine field to given value.

### HasRunningMachine

`func (o *GetCurrentVMStatus200ResponseData) HasRunningMachine() bool`

HasRunningMachine returns a boolean if a field has been set.

### GetRunningQemu

`func (o *GetCurrentVMStatus200ResponseData) GetRunningQemu() string`

GetRunningQemu returns the RunningQemu field if non-nil, zero value otherwise.

### GetRunningQemuOk

`func (o *GetCurrentVMStatus200ResponseData) GetRunningQemuOk() (*string, bool)`

GetRunningQemuOk returns a tuple with the RunningQemu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningQemu

`func (o *GetCurrentVMStatus200ResponseData) SetRunningQemu(v string)`

SetRunningQemu sets RunningQemu field to given value.

### HasRunningQemu

`func (o *GetCurrentVMStatus200ResponseData) HasRunningQemu() bool`

HasRunningQemu returns a boolean if a field has been set.

### GetSpice

`func (o *GetCurrentVMStatus200ResponseData) GetSpice() int32`

GetSpice returns the Spice field if non-nil, zero value otherwise.

### GetSpiceOk

`func (o *GetCurrentVMStatus200ResponseData) GetSpiceOk() (*int32, bool)`

GetSpiceOk returns a tuple with the Spice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpice

`func (o *GetCurrentVMStatus200ResponseData) SetSpice(v int32)`

SetSpice sets Spice field to given value.

### HasSpice

`func (o *GetCurrentVMStatus200ResponseData) HasSpice() bool`

HasSpice returns a boolean if a field has been set.

### GetStatus

`func (o *GetCurrentVMStatus200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCurrentVMStatus200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCurrentVMStatus200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCurrentVMStatus200ResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetCurrentVMStatus200ResponseData) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetCurrentVMStatus200ResponseData) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetCurrentVMStatus200ResponseData) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetCurrentVMStatus200ResponseData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *GetCurrentVMStatus200ResponseData) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetCurrentVMStatus200ResponseData) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetCurrentVMStatus200ResponseData) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetCurrentVMStatus200ResponseData) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUptime

`func (o *GetCurrentVMStatus200ResponseData) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GetCurrentVMStatus200ResponseData) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GetCurrentVMStatus200ResponseData) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GetCurrentVMStatus200ResponseData) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVmid

`func (o *GetCurrentVMStatus200ResponseData) GetVmid() int64`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *GetCurrentVMStatus200ResponseData) GetVmidOk() (*int64, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *GetCurrentVMStatus200ResponseData) SetVmid(v int64)`

SetVmid sets Vmid field to given value.

### HasVmid

`func (o *GetCurrentVMStatus200ResponseData) HasVmid() bool`

HasVmid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


