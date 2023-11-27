# GetContainerConfig200ResponseDataRootfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **bool** | Explicitly enable or disable ACL support. | [optional] 
**Mountoptions** | Pointer to **string** | Extra mount options for rootfs/mps. | [optional] 
**Quota** | Pointer to **bool** | Enable user quotas inside the container (not supported with zfs subvolumes) | [optional] 
**Replicate** | Pointer to **bool** | Will include this volume to a storage replica job. | [optional] 
**Ro** | Pointer to **bool** | Read-only mount point | [optional] 
**Shared** | Pointer to **bool** | Mark this non-volume mount point as available on multiple nodes (see &#39;nodes&#39;) | [optional] 
**Size** | Pointer to **string** | Volume size (read only value). | [optional] 
**Volume** | Pointer to **string** | Volume, device or directory to mount into the container. | [optional] 

## Methods

### NewGetContainerConfig200ResponseDataRootfs

`func NewGetContainerConfig200ResponseDataRootfs() *GetContainerConfig200ResponseDataRootfs`

NewGetContainerConfig200ResponseDataRootfs instantiates a new GetContainerConfig200ResponseDataRootfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerConfig200ResponseDataRootfsWithDefaults

`func NewGetContainerConfig200ResponseDataRootfsWithDefaults() *GetContainerConfig200ResponseDataRootfs`

NewGetContainerConfig200ResponseDataRootfsWithDefaults instantiates a new GetContainerConfig200ResponseDataRootfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *GetContainerConfig200ResponseDataRootfs) GetAcl() bool`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *GetContainerConfig200ResponseDataRootfs) GetAclOk() (*bool, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *GetContainerConfig200ResponseDataRootfs) SetAcl(v bool)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *GetContainerConfig200ResponseDataRootfs) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetMountoptions

`func (o *GetContainerConfig200ResponseDataRootfs) GetMountoptions() string`

GetMountoptions returns the Mountoptions field if non-nil, zero value otherwise.

### GetMountoptionsOk

`func (o *GetContainerConfig200ResponseDataRootfs) GetMountoptionsOk() (*string, bool)`

GetMountoptionsOk returns a tuple with the Mountoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountoptions

`func (o *GetContainerConfig200ResponseDataRootfs) SetMountoptions(v string)`

SetMountoptions sets Mountoptions field to given value.

### HasMountoptions

`func (o *GetContainerConfig200ResponseDataRootfs) HasMountoptions() bool`

HasMountoptions returns a boolean if a field has been set.

### GetQuota

`func (o *GetContainerConfig200ResponseDataRootfs) GetQuota() bool`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetContainerConfig200ResponseDataRootfs) GetQuotaOk() (*bool, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetContainerConfig200ResponseDataRootfs) SetQuota(v bool)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *GetContainerConfig200ResponseDataRootfs) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetReplicate

`func (o *GetContainerConfig200ResponseDataRootfs) GetReplicate() bool`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *GetContainerConfig200ResponseDataRootfs) GetReplicateOk() (*bool, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *GetContainerConfig200ResponseDataRootfs) SetReplicate(v bool)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *GetContainerConfig200ResponseDataRootfs) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetRo

`func (o *GetContainerConfig200ResponseDataRootfs) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *GetContainerConfig200ResponseDataRootfs) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *GetContainerConfig200ResponseDataRootfs) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *GetContainerConfig200ResponseDataRootfs) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetShared

`func (o *GetContainerConfig200ResponseDataRootfs) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *GetContainerConfig200ResponseDataRootfs) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *GetContainerConfig200ResponseDataRootfs) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *GetContainerConfig200ResponseDataRootfs) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *GetContainerConfig200ResponseDataRootfs) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetContainerConfig200ResponseDataRootfs) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetContainerConfig200ResponseDataRootfs) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetContainerConfig200ResponseDataRootfs) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolume

`func (o *GetContainerConfig200ResponseDataRootfs) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetContainerConfig200ResponseDataRootfs) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetContainerConfig200ResponseDataRootfs) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetContainerConfig200ResponseDataRootfs) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


