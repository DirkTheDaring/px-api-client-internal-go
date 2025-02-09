# CreateContainerRequestMp0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **bool** |  | [optional] 
**Backup** | Pointer to **bool** |  | [optional] 
**Mountoptions** | Pointer to **string** |  | [optional] 
**Mp** | **string** |  | 
**Quota** | Pointer to **bool** |  | [optional] 
**Replicate** | Pointer to **bool** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Volume** | **string** |  | 

## Methods

### NewCreateContainerRequestMp0

`func NewCreateContainerRequestMp0(mp string, volume string, ) *CreateContainerRequestMp0`

NewCreateContainerRequestMp0 instantiates a new CreateContainerRequestMp0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerRequestMp0WithDefaults

`func NewCreateContainerRequestMp0WithDefaults() *CreateContainerRequestMp0`

NewCreateContainerRequestMp0WithDefaults instantiates a new CreateContainerRequestMp0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *CreateContainerRequestMp0) GetAcl() bool`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *CreateContainerRequestMp0) GetAclOk() (*bool, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *CreateContainerRequestMp0) SetAcl(v bool)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *CreateContainerRequestMp0) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetBackup

`func (o *CreateContainerRequestMp0) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateContainerRequestMp0) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateContainerRequestMp0) SetBackup(v bool)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *CreateContainerRequestMp0) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetMountoptions

`func (o *CreateContainerRequestMp0) GetMountoptions() string`

GetMountoptions returns the Mountoptions field if non-nil, zero value otherwise.

### GetMountoptionsOk

`func (o *CreateContainerRequestMp0) GetMountoptionsOk() (*string, bool)`

GetMountoptionsOk returns a tuple with the Mountoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountoptions

`func (o *CreateContainerRequestMp0) SetMountoptions(v string)`

SetMountoptions sets Mountoptions field to given value.

### HasMountoptions

`func (o *CreateContainerRequestMp0) HasMountoptions() bool`

HasMountoptions returns a boolean if a field has been set.

### GetMp

`func (o *CreateContainerRequestMp0) GetMp() string`

GetMp returns the Mp field if non-nil, zero value otherwise.

### GetMpOk

`func (o *CreateContainerRequestMp0) GetMpOk() (*string, bool)`

GetMpOk returns a tuple with the Mp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp

`func (o *CreateContainerRequestMp0) SetMp(v string)`

SetMp sets Mp field to given value.


### GetQuota

`func (o *CreateContainerRequestMp0) GetQuota() bool`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateContainerRequestMp0) GetQuotaOk() (*bool, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateContainerRequestMp0) SetQuota(v bool)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateContainerRequestMp0) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetReplicate

`func (o *CreateContainerRequestMp0) GetReplicate() bool`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *CreateContainerRequestMp0) GetReplicateOk() (*bool, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *CreateContainerRequestMp0) SetReplicate(v bool)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *CreateContainerRequestMp0) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetRo

`func (o *CreateContainerRequestMp0) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *CreateContainerRequestMp0) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *CreateContainerRequestMp0) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *CreateContainerRequestMp0) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetShared

`func (o *CreateContainerRequestMp0) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateContainerRequestMp0) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateContainerRequestMp0) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateContainerRequestMp0) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *CreateContainerRequestMp0) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateContainerRequestMp0) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateContainerRequestMp0) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateContainerRequestMp0) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolume

`func (o *CreateContainerRequestMp0) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateContainerRequestMp0) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateContainerRequestMp0) SetVolume(v string)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


