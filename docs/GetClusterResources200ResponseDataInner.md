# GetClusterResources200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CgroupMode** | Pointer to **int64** | The cgroup mode the node operates under (for type &#39;node&#39;). | [optional] 
**Content** | Pointer to **string** | Allowed storage content types (for type &#39;storage&#39;). | [optional] 
**Cpu** | Pointer to **float32** | CPU utilization (for types &#39;node&#39;, &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 
**Disk** | Pointer to **int64** | Used disk space in bytes (for type &#39;storage&#39;), used root image space for VMs (for types &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 
**Diskread** | Pointer to **int64** | The amount of bytes the guest read from its block devices since the guest was started. This info is not available for all storage types. (for types &#39;qemu&#39; and &#39;lxc&#39;) | [optional] 
**Diskwrite** | Pointer to **int64** | The amount of bytes the guest wrote to its block devices since the guest was started. This info is not available for all storage types. (for types &#39;qemu&#39; and &#39;lxc&#39;) | [optional] 
**Hastate** | Pointer to **string** | HA service status (for HA managed VMs). | [optional] 
**Id** | Pointer to **string** | Resource id. | [optional] 
**Level** | Pointer to **string** | Support level (for type &#39;node&#39;). | [optional] 
**Lock** | Pointer to **string** | The guest&#39;s current config lock (for types &#39;qemu&#39; and &#39;lxc&#39;) | [optional] 
**Maxcpu** | Pointer to **float32** | Number of available CPUs (for types &#39;node&#39;, &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 
**Maxdisk** | Pointer to **int64** | Storage size in bytes (for type &#39;storage&#39;), root image size for VMs (for types &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 
**Maxmem** | Pointer to **int64** | Number of available memory in bytes (for types &#39;node&#39;, &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 
**Mem** | Pointer to **int64** | Used memory in bytes (for types &#39;node&#39;, &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Netin** | Pointer to **int64** | The amount of traffic in bytes that was sent to the guest over the network since it was started. (for types &#39;qemu&#39; and &#39;lxc&#39;) | [optional] 
**Netout** | Pointer to **int64** | The amount of traffic in bytes that was sent from the guest over the network since it was started. (for types &#39;qemu&#39; and &#39;lxc&#39;) | [optional] 
**Node** | Pointer to **string** | The cluster node name (for types &#39;node&#39;, &#39;storage&#39;, &#39;qemu&#39;, and &#39;lxc&#39;). | [optional] 
**Plugintype** | Pointer to **string** | More specific type, if available. | [optional] 
**Pool** | Pointer to **string** | The pool name (for types &#39;pool&#39;, &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 
**Status** | Pointer to **string** | Resource type dependent status. | [optional] 
**Storage** | Pointer to **string** | The storage identifier (for type &#39;storage&#39;). | [optional] 
**Tags** | Pointer to **string** | The guest&#39;s tags (for types &#39;qemu&#39; and &#39;lxc&#39;) | [optional] 
**Template** | Pointer to **int64** | Determines if the guest is a template. (for types &#39;qemu&#39; and &#39;lxc&#39;) | [optional] 
**Type** | Pointer to **string** | Resource type. | [optional] 
**Uptime** | Pointer to **int64** | Uptime of node or virtual guest in seconds (for types &#39;node&#39;, &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 
**Vmid** | Pointer to **int64** | The numerical vmid (for types &#39;qemu&#39; and &#39;lxc&#39;). | [optional] 

## Methods

### NewGetClusterResources200ResponseDataInner

`func NewGetClusterResources200ResponseDataInner() *GetClusterResources200ResponseDataInner`

NewGetClusterResources200ResponseDataInner instantiates a new GetClusterResources200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterResources200ResponseDataInnerWithDefaults

`func NewGetClusterResources200ResponseDataInnerWithDefaults() *GetClusterResources200ResponseDataInner`

NewGetClusterResources200ResponseDataInnerWithDefaults instantiates a new GetClusterResources200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCgroupMode

`func (o *GetClusterResources200ResponseDataInner) GetCgroupMode() int64`

GetCgroupMode returns the CgroupMode field if non-nil, zero value otherwise.

### GetCgroupModeOk

`func (o *GetClusterResources200ResponseDataInner) GetCgroupModeOk() (*int64, bool)`

GetCgroupModeOk returns a tuple with the CgroupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupMode

`func (o *GetClusterResources200ResponseDataInner) SetCgroupMode(v int64)`

SetCgroupMode sets CgroupMode field to given value.

### HasCgroupMode

`func (o *GetClusterResources200ResponseDataInner) HasCgroupMode() bool`

HasCgroupMode returns a boolean if a field has been set.

### GetContent

`func (o *GetClusterResources200ResponseDataInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetClusterResources200ResponseDataInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetClusterResources200ResponseDataInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetClusterResources200ResponseDataInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCpu

`func (o *GetClusterResources200ResponseDataInner) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GetClusterResources200ResponseDataInner) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GetClusterResources200ResponseDataInner) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *GetClusterResources200ResponseDataInner) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisk

`func (o *GetClusterResources200ResponseDataInner) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *GetClusterResources200ResponseDataInner) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *GetClusterResources200ResponseDataInner) SetDisk(v int64)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *GetClusterResources200ResponseDataInner) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetDiskread

`func (o *GetClusterResources200ResponseDataInner) GetDiskread() int64`

GetDiskread returns the Diskread field if non-nil, zero value otherwise.

### GetDiskreadOk

`func (o *GetClusterResources200ResponseDataInner) GetDiskreadOk() (*int64, bool)`

GetDiskreadOk returns a tuple with the Diskread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskread

`func (o *GetClusterResources200ResponseDataInner) SetDiskread(v int64)`

SetDiskread sets Diskread field to given value.

### HasDiskread

`func (o *GetClusterResources200ResponseDataInner) HasDiskread() bool`

HasDiskread returns a boolean if a field has been set.

### GetDiskwrite

`func (o *GetClusterResources200ResponseDataInner) GetDiskwrite() int64`

GetDiskwrite returns the Diskwrite field if non-nil, zero value otherwise.

### GetDiskwriteOk

`func (o *GetClusterResources200ResponseDataInner) GetDiskwriteOk() (*int64, bool)`

GetDiskwriteOk returns a tuple with the Diskwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskwrite

`func (o *GetClusterResources200ResponseDataInner) SetDiskwrite(v int64)`

SetDiskwrite sets Diskwrite field to given value.

### HasDiskwrite

`func (o *GetClusterResources200ResponseDataInner) HasDiskwrite() bool`

HasDiskwrite returns a boolean if a field has been set.

### GetHastate

`func (o *GetClusterResources200ResponseDataInner) GetHastate() string`

GetHastate returns the Hastate field if non-nil, zero value otherwise.

### GetHastateOk

`func (o *GetClusterResources200ResponseDataInner) GetHastateOk() (*string, bool)`

GetHastateOk returns a tuple with the Hastate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHastate

`func (o *GetClusterResources200ResponseDataInner) SetHastate(v string)`

SetHastate sets Hastate field to given value.

### HasHastate

`func (o *GetClusterResources200ResponseDataInner) HasHastate() bool`

HasHastate returns a boolean if a field has been set.

### GetId

`func (o *GetClusterResources200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterResources200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterResources200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterResources200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *GetClusterResources200ResponseDataInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetClusterResources200ResponseDataInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetClusterResources200ResponseDataInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetClusterResources200ResponseDataInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLock

`func (o *GetClusterResources200ResponseDataInner) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetClusterResources200ResponseDataInner) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetClusterResources200ResponseDataInner) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetClusterResources200ResponseDataInner) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMaxcpu

`func (o *GetClusterResources200ResponseDataInner) GetMaxcpu() float32`

GetMaxcpu returns the Maxcpu field if non-nil, zero value otherwise.

### GetMaxcpuOk

`func (o *GetClusterResources200ResponseDataInner) GetMaxcpuOk() (*float32, bool)`

GetMaxcpuOk returns a tuple with the Maxcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxcpu

`func (o *GetClusterResources200ResponseDataInner) SetMaxcpu(v float32)`

SetMaxcpu sets Maxcpu field to given value.

### HasMaxcpu

`func (o *GetClusterResources200ResponseDataInner) HasMaxcpu() bool`

HasMaxcpu returns a boolean if a field has been set.

### GetMaxdisk

`func (o *GetClusterResources200ResponseDataInner) GetMaxdisk() int64`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *GetClusterResources200ResponseDataInner) GetMaxdiskOk() (*int64, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *GetClusterResources200ResponseDataInner) SetMaxdisk(v int64)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *GetClusterResources200ResponseDataInner) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetMaxmem

`func (o *GetClusterResources200ResponseDataInner) GetMaxmem() int64`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *GetClusterResources200ResponseDataInner) GetMaxmemOk() (*int64, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *GetClusterResources200ResponseDataInner) SetMaxmem(v int64)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *GetClusterResources200ResponseDataInner) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetMem

`func (o *GetClusterResources200ResponseDataInner) GetMem() int64`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *GetClusterResources200ResponseDataInner) GetMemOk() (*int64, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *GetClusterResources200ResponseDataInner) SetMem(v int64)`

SetMem sets Mem field to given value.

### HasMem

`func (o *GetClusterResources200ResponseDataInner) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetName

`func (o *GetClusterResources200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterResources200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterResources200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterResources200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetin

`func (o *GetClusterResources200ResponseDataInner) GetNetin() int64`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *GetClusterResources200ResponseDataInner) GetNetinOk() (*int64, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *GetClusterResources200ResponseDataInner) SetNetin(v int64)`

SetNetin sets Netin field to given value.

### HasNetin

`func (o *GetClusterResources200ResponseDataInner) HasNetin() bool`

HasNetin returns a boolean if a field has been set.

### GetNetout

`func (o *GetClusterResources200ResponseDataInner) GetNetout() int64`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *GetClusterResources200ResponseDataInner) GetNetoutOk() (*int64, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *GetClusterResources200ResponseDataInner) SetNetout(v int64)`

SetNetout sets Netout field to given value.

### HasNetout

`func (o *GetClusterResources200ResponseDataInner) HasNetout() bool`

HasNetout returns a boolean if a field has been set.

### GetNode

`func (o *GetClusterResources200ResponseDataInner) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *GetClusterResources200ResponseDataInner) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *GetClusterResources200ResponseDataInner) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *GetClusterResources200ResponseDataInner) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPlugintype

`func (o *GetClusterResources200ResponseDataInner) GetPlugintype() string`

GetPlugintype returns the Plugintype field if non-nil, zero value otherwise.

### GetPlugintypeOk

`func (o *GetClusterResources200ResponseDataInner) GetPlugintypeOk() (*string, bool)`

GetPlugintypeOk returns a tuple with the Plugintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugintype

`func (o *GetClusterResources200ResponseDataInner) SetPlugintype(v string)`

SetPlugintype sets Plugintype field to given value.

### HasPlugintype

`func (o *GetClusterResources200ResponseDataInner) HasPlugintype() bool`

HasPlugintype returns a boolean if a field has been set.

### GetPool

`func (o *GetClusterResources200ResponseDataInner) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GetClusterResources200ResponseDataInner) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GetClusterResources200ResponseDataInner) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GetClusterResources200ResponseDataInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetStatus

`func (o *GetClusterResources200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterResources200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterResources200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClusterResources200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorage

`func (o *GetClusterResources200ResponseDataInner) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *GetClusterResources200ResponseDataInner) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *GetClusterResources200ResponseDataInner) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *GetClusterResources200ResponseDataInner) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTags

`func (o *GetClusterResources200ResponseDataInner) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetClusterResources200ResponseDataInner) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetClusterResources200ResponseDataInner) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetClusterResources200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *GetClusterResources200ResponseDataInner) GetTemplate() int64`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetClusterResources200ResponseDataInner) GetTemplateOk() (*int64, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetClusterResources200ResponseDataInner) SetTemplate(v int64)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetClusterResources200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *GetClusterResources200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClusterResources200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClusterResources200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetClusterResources200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *GetClusterResources200ResponseDataInner) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GetClusterResources200ResponseDataInner) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GetClusterResources200ResponseDataInner) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GetClusterResources200ResponseDataInner) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVmid

`func (o *GetClusterResources200ResponseDataInner) GetVmid() int64`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *GetClusterResources200ResponseDataInner) GetVmidOk() (*int64, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *GetClusterResources200ResponseDataInner) SetVmid(v int64)`

SetVmid sets Vmid field to given value.

### HasVmid

`func (o *GetClusterResources200ResponseDataInner) HasVmid() bool`

HasVmid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


