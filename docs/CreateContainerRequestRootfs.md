# CreateContainerRequestRootfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **bool** |  | [optional] 
**Mountoptions** | Pointer to **string** |  | [optional] 
**Quota** | Pointer to **bool** |  | [optional] 
**Replicate** | Pointer to **bool** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Volume** | **string** |  | 

## Methods

### NewCreateContainerRequestRootfs

`func NewCreateContainerRequestRootfs(volume string, ) *CreateContainerRequestRootfs`

NewCreateContainerRequestRootfs instantiates a new CreateContainerRequestRootfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerRequestRootfsWithDefaults

`func NewCreateContainerRequestRootfsWithDefaults() *CreateContainerRequestRootfs`

NewCreateContainerRequestRootfsWithDefaults instantiates a new CreateContainerRequestRootfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *CreateContainerRequestRootfs) GetAcl() bool`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *CreateContainerRequestRootfs) GetAclOk() (*bool, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *CreateContainerRequestRootfs) SetAcl(v bool)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *CreateContainerRequestRootfs) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetMountoptions

`func (o *CreateContainerRequestRootfs) GetMountoptions() string`

GetMountoptions returns the Mountoptions field if non-nil, zero value otherwise.

### GetMountoptionsOk

`func (o *CreateContainerRequestRootfs) GetMountoptionsOk() (*string, bool)`

GetMountoptionsOk returns a tuple with the Mountoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountoptions

`func (o *CreateContainerRequestRootfs) SetMountoptions(v string)`

SetMountoptions sets Mountoptions field to given value.

### HasMountoptions

`func (o *CreateContainerRequestRootfs) HasMountoptions() bool`

HasMountoptions returns a boolean if a field has been set.

### GetQuota

`func (o *CreateContainerRequestRootfs) GetQuota() bool`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateContainerRequestRootfs) GetQuotaOk() (*bool, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateContainerRequestRootfs) SetQuota(v bool)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateContainerRequestRootfs) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetReplicate

`func (o *CreateContainerRequestRootfs) GetReplicate() bool`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *CreateContainerRequestRootfs) GetReplicateOk() (*bool, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *CreateContainerRequestRootfs) SetReplicate(v bool)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *CreateContainerRequestRootfs) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetRo

`func (o *CreateContainerRequestRootfs) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *CreateContainerRequestRootfs) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *CreateContainerRequestRootfs) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *CreateContainerRequestRootfs) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetShared

`func (o *CreateContainerRequestRootfs) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateContainerRequestRootfs) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateContainerRequestRootfs) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateContainerRequestRootfs) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *CreateContainerRequestRootfs) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateContainerRequestRootfs) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateContainerRequestRootfs) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateContainerRequestRootfs) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolume

`func (o *CreateContainerRequestRootfs) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateContainerRequestRootfs) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateContainerRequestRootfs) SetVolume(v string)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


