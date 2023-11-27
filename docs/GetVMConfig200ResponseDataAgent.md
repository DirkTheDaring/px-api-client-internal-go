# GetVMConfig200ResponseDataAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable/disable communication with a QEMU Guest Agent (QGA) running in the VM. | [optional] 
**FreezeFsOnBackup** | Pointer to **bool** | Freeze/thaw guest filesystems on backup for consistency. | [optional] 
**FstrimClonedDisks** | Pointer to **bool** | Run fstrim after moving a disk or migrating the VM. | [optional] 
**Type** | Pointer to **string** | Select the agent type | [optional] 

## Methods

### NewGetVMConfig200ResponseDataAgent

`func NewGetVMConfig200ResponseDataAgent() *GetVMConfig200ResponseDataAgent`

NewGetVMConfig200ResponseDataAgent instantiates a new GetVMConfig200ResponseDataAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataAgentWithDefaults

`func NewGetVMConfig200ResponseDataAgentWithDefaults() *GetVMConfig200ResponseDataAgent`

NewGetVMConfig200ResponseDataAgentWithDefaults instantiates a new GetVMConfig200ResponseDataAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetVMConfig200ResponseDataAgent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetVMConfig200ResponseDataAgent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetVMConfig200ResponseDataAgent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetVMConfig200ResponseDataAgent) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFreezeFsOnBackup

`func (o *GetVMConfig200ResponseDataAgent) GetFreezeFsOnBackup() bool`

GetFreezeFsOnBackup returns the FreezeFsOnBackup field if non-nil, zero value otherwise.

### GetFreezeFsOnBackupOk

`func (o *GetVMConfig200ResponseDataAgent) GetFreezeFsOnBackupOk() (*bool, bool)`

GetFreezeFsOnBackupOk returns a tuple with the FreezeFsOnBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeFsOnBackup

`func (o *GetVMConfig200ResponseDataAgent) SetFreezeFsOnBackup(v bool)`

SetFreezeFsOnBackup sets FreezeFsOnBackup field to given value.

### HasFreezeFsOnBackup

`func (o *GetVMConfig200ResponseDataAgent) HasFreezeFsOnBackup() bool`

HasFreezeFsOnBackup returns a boolean if a field has been set.

### GetFstrimClonedDisks

`func (o *GetVMConfig200ResponseDataAgent) GetFstrimClonedDisks() bool`

GetFstrimClonedDisks returns the FstrimClonedDisks field if non-nil, zero value otherwise.

### GetFstrimClonedDisksOk

`func (o *GetVMConfig200ResponseDataAgent) GetFstrimClonedDisksOk() (*bool, bool)`

GetFstrimClonedDisksOk returns a tuple with the FstrimClonedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFstrimClonedDisks

`func (o *GetVMConfig200ResponseDataAgent) SetFstrimClonedDisks(v bool)`

SetFstrimClonedDisks sets FstrimClonedDisks field to given value.

### HasFstrimClonedDisks

`func (o *GetVMConfig200ResponseDataAgent) HasFstrimClonedDisks() bool`

HasFstrimClonedDisks returns a boolean if a field has been set.

### GetType

`func (o *GetVMConfig200ResponseDataAgent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetVMConfig200ResponseDataAgent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetVMConfig200ResponseDataAgent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetVMConfig200ResponseDataAgent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


