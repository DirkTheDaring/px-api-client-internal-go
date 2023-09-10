# GetContainerConfig200ResponseDataFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceRwSys** | Pointer to **int32** | Mount /sys in unprivileged containers as &#x60;rw&#x60; instead of &#x60;mixed&#x60;. This can break networking under newer (&gt;&#x3D; v245) systemd-network use. | [optional] 
**Fuse** | Pointer to **int32** | Allow using &#39;fuse&#39; file systems in a container. Note that interactions between fuse and the freezer cgroup can potentially cause I/O deadlocks. | [optional] 
**Keyctl** | Pointer to **int32** | For unprivileged containers only: Allow the use of the keyctl() system call. This is required to use docker inside a container. By default unprivileged containers will see this system call as non-existent. This is mostly a workaround for systemd-networkd, as it will treat it as a fatal error when some keyctl() operations are denied by the kernel due to lacking permissions. Essentially, you can choose between running systemd-networkd or docker. | [optional] 
**Mknod** | Pointer to **int32** | Allow unprivileged containers to use mknod() to add certain device nodes. This requires a kernel with seccomp trap to user space support (5.3 or newer). This is experimental. | [optional] 
**Mount** | Pointer to **string** | Allow mounting file systems of specific types. This should be a list of file system types as used with the mount command. Note that this can have negative effects on the container&#39;s security. With access to a loop device, mounting a file can circumvent the mknod permission of the devices cgroup, mounting an NFS file system can block the host&#39;s I/O completely and prevent it from rebooting, etc. | [optional] 
**Nesting** | Pointer to **int32** | Allow nesting. Best used with unprivileged containers with additional id mapping. Note that this will expose procfs and sysfs contents of the host to the guest. | [optional] 

## Methods

### NewGetContainerConfig200ResponseDataFeatures

`func NewGetContainerConfig200ResponseDataFeatures() *GetContainerConfig200ResponseDataFeatures`

NewGetContainerConfig200ResponseDataFeatures instantiates a new GetContainerConfig200ResponseDataFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerConfig200ResponseDataFeaturesWithDefaults

`func NewGetContainerConfig200ResponseDataFeaturesWithDefaults() *GetContainerConfig200ResponseDataFeatures`

NewGetContainerConfig200ResponseDataFeaturesWithDefaults instantiates a new GetContainerConfig200ResponseDataFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceRwSys

`func (o *GetContainerConfig200ResponseDataFeatures) GetForceRwSys() int32`

GetForceRwSys returns the ForceRwSys field if non-nil, zero value otherwise.

### GetForceRwSysOk

`func (o *GetContainerConfig200ResponseDataFeatures) GetForceRwSysOk() (*int32, bool)`

GetForceRwSysOk returns a tuple with the ForceRwSys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRwSys

`func (o *GetContainerConfig200ResponseDataFeatures) SetForceRwSys(v int32)`

SetForceRwSys sets ForceRwSys field to given value.

### HasForceRwSys

`func (o *GetContainerConfig200ResponseDataFeatures) HasForceRwSys() bool`

HasForceRwSys returns a boolean if a field has been set.

### GetFuse

`func (o *GetContainerConfig200ResponseDataFeatures) GetFuse() int32`

GetFuse returns the Fuse field if non-nil, zero value otherwise.

### GetFuseOk

`func (o *GetContainerConfig200ResponseDataFeatures) GetFuseOk() (*int32, bool)`

GetFuseOk returns a tuple with the Fuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuse

`func (o *GetContainerConfig200ResponseDataFeatures) SetFuse(v int32)`

SetFuse sets Fuse field to given value.

### HasFuse

`func (o *GetContainerConfig200ResponseDataFeatures) HasFuse() bool`

HasFuse returns a boolean if a field has been set.

### GetKeyctl

`func (o *GetContainerConfig200ResponseDataFeatures) GetKeyctl() int32`

GetKeyctl returns the Keyctl field if non-nil, zero value otherwise.

### GetKeyctlOk

`func (o *GetContainerConfig200ResponseDataFeatures) GetKeyctlOk() (*int32, bool)`

GetKeyctlOk returns a tuple with the Keyctl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyctl

`func (o *GetContainerConfig200ResponseDataFeatures) SetKeyctl(v int32)`

SetKeyctl sets Keyctl field to given value.

### HasKeyctl

`func (o *GetContainerConfig200ResponseDataFeatures) HasKeyctl() bool`

HasKeyctl returns a boolean if a field has been set.

### GetMknod

`func (o *GetContainerConfig200ResponseDataFeatures) GetMknod() int32`

GetMknod returns the Mknod field if non-nil, zero value otherwise.

### GetMknodOk

`func (o *GetContainerConfig200ResponseDataFeatures) GetMknodOk() (*int32, bool)`

GetMknodOk returns a tuple with the Mknod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMknod

`func (o *GetContainerConfig200ResponseDataFeatures) SetMknod(v int32)`

SetMknod sets Mknod field to given value.

### HasMknod

`func (o *GetContainerConfig200ResponseDataFeatures) HasMknod() bool`

HasMknod returns a boolean if a field has been set.

### GetMount

`func (o *GetContainerConfig200ResponseDataFeatures) GetMount() string`

GetMount returns the Mount field if non-nil, zero value otherwise.

### GetMountOk

`func (o *GetContainerConfig200ResponseDataFeatures) GetMountOk() (*string, bool)`

GetMountOk returns a tuple with the Mount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMount

`func (o *GetContainerConfig200ResponseDataFeatures) SetMount(v string)`

SetMount sets Mount field to given value.

### HasMount

`func (o *GetContainerConfig200ResponseDataFeatures) HasMount() bool`

HasMount returns a boolean if a field has been set.

### GetNesting

`func (o *GetContainerConfig200ResponseDataFeatures) GetNesting() int32`

GetNesting returns the Nesting field if non-nil, zero value otherwise.

### GetNestingOk

`func (o *GetContainerConfig200ResponseDataFeatures) GetNestingOk() (*int32, bool)`

GetNestingOk returns a tuple with the Nesting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNesting

`func (o *GetContainerConfig200ResponseDataFeatures) SetNesting(v int32)`

SetNesting sets Nesting field to given value.

### HasNesting

`func (o *GetContainerConfig200ResponseDataFeatures) HasNesting() bool`

HasNesting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


