# GetContainerConfig200ResponseDataMp0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **int32** | Explicitly enable or disable ACL support. | [optional] 
**Backup** | Pointer to **int32** | Whether to include the mount point in backups. | [optional] 
**Mountoptions** | Pointer to **string** | Extra mount options for rootfs/mps. | [optional] 
**Mp** | Pointer to **string** | Path to the mount point as seen from inside the container (must not contain symlinks). | [optional] 
**Quota** | Pointer to **int32** | Enable user quotas inside the container (not supported with zfs subvolumes) | [optional] 
**Replicate** | Pointer to **int32** | Will include this volume to a storage replica job. | [optional] 
**Ro** | Pointer to **int32** | Read-only mount point | [optional] 
**Shared** | Pointer to **int32** | Mark this non-volume mount point as available on multiple nodes (see &#39;nodes&#39;) | [optional] 
**Size** | Pointer to **string** | Volume size (read only value). | [optional] 
**Volume** | Pointer to **string** | Volume, device or directory to mount into the container. | [optional] 

## Methods

### NewGetContainerConfig200ResponseDataMp0

`func NewGetContainerConfig200ResponseDataMp0() *GetContainerConfig200ResponseDataMp0`

NewGetContainerConfig200ResponseDataMp0 instantiates a new GetContainerConfig200ResponseDataMp0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerConfig200ResponseDataMp0WithDefaults

`func NewGetContainerConfig200ResponseDataMp0WithDefaults() *GetContainerConfig200ResponseDataMp0`

NewGetContainerConfig200ResponseDataMp0WithDefaults instantiates a new GetContainerConfig200ResponseDataMp0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *GetContainerConfig200ResponseDataMp0) GetAcl() int32`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *GetContainerConfig200ResponseDataMp0) GetAclOk() (*int32, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *GetContainerConfig200ResponseDataMp0) SetAcl(v int32)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *GetContainerConfig200ResponseDataMp0) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetBackup

`func (o *GetContainerConfig200ResponseDataMp0) GetBackup() int32`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetContainerConfig200ResponseDataMp0) GetBackupOk() (*int32, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetContainerConfig200ResponseDataMp0) SetBackup(v int32)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *GetContainerConfig200ResponseDataMp0) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetMountoptions

`func (o *GetContainerConfig200ResponseDataMp0) GetMountoptions() string`

GetMountoptions returns the Mountoptions field if non-nil, zero value otherwise.

### GetMountoptionsOk

`func (o *GetContainerConfig200ResponseDataMp0) GetMountoptionsOk() (*string, bool)`

GetMountoptionsOk returns a tuple with the Mountoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountoptions

`func (o *GetContainerConfig200ResponseDataMp0) SetMountoptions(v string)`

SetMountoptions sets Mountoptions field to given value.

### HasMountoptions

`func (o *GetContainerConfig200ResponseDataMp0) HasMountoptions() bool`

HasMountoptions returns a boolean if a field has been set.

### GetMp

`func (o *GetContainerConfig200ResponseDataMp0) GetMp() string`

GetMp returns the Mp field if non-nil, zero value otherwise.

### GetMpOk

`func (o *GetContainerConfig200ResponseDataMp0) GetMpOk() (*string, bool)`

GetMpOk returns a tuple with the Mp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp

`func (o *GetContainerConfig200ResponseDataMp0) SetMp(v string)`

SetMp sets Mp field to given value.

### HasMp

`func (o *GetContainerConfig200ResponseDataMp0) HasMp() bool`

HasMp returns a boolean if a field has been set.

### GetQuota

`func (o *GetContainerConfig200ResponseDataMp0) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetContainerConfig200ResponseDataMp0) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetContainerConfig200ResponseDataMp0) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *GetContainerConfig200ResponseDataMp0) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetReplicate

`func (o *GetContainerConfig200ResponseDataMp0) GetReplicate() int32`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *GetContainerConfig200ResponseDataMp0) GetReplicateOk() (*int32, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *GetContainerConfig200ResponseDataMp0) SetReplicate(v int32)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *GetContainerConfig200ResponseDataMp0) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetRo

`func (o *GetContainerConfig200ResponseDataMp0) GetRo() int32`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *GetContainerConfig200ResponseDataMp0) GetRoOk() (*int32, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *GetContainerConfig200ResponseDataMp0) SetRo(v int32)`

SetRo sets Ro field to given value.

### HasRo

`func (o *GetContainerConfig200ResponseDataMp0) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetShared

`func (o *GetContainerConfig200ResponseDataMp0) GetShared() int32`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *GetContainerConfig200ResponseDataMp0) GetSharedOk() (*int32, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *GetContainerConfig200ResponseDataMp0) SetShared(v int32)`

SetShared sets Shared field to given value.

### HasShared

`func (o *GetContainerConfig200ResponseDataMp0) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *GetContainerConfig200ResponseDataMp0) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetContainerConfig200ResponseDataMp0) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetContainerConfig200ResponseDataMp0) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetContainerConfig200ResponseDataMp0) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolume

`func (o *GetContainerConfig200ResponseDataMp0) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetContainerConfig200ResponseDataMp0) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetContainerConfig200ResponseDataMp0) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetContainerConfig200ResponseDataMp0) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


