# CreateVMRequestAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**FreezeFsOnBackup** | Pointer to **bool** |  | [optional] 
**FstrimClonedDisks** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVMRequestAgent

`func NewCreateVMRequestAgent(enabled bool, ) *CreateVMRequestAgent`

NewCreateVMRequestAgent instantiates a new CreateVMRequestAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestAgentWithDefaults

`func NewCreateVMRequestAgentWithDefaults() *CreateVMRequestAgent`

NewCreateVMRequestAgentWithDefaults instantiates a new CreateVMRequestAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateVMRequestAgent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateVMRequestAgent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateVMRequestAgent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFreezeFsOnBackup

`func (o *CreateVMRequestAgent) GetFreezeFsOnBackup() bool`

GetFreezeFsOnBackup returns the FreezeFsOnBackup field if non-nil, zero value otherwise.

### GetFreezeFsOnBackupOk

`func (o *CreateVMRequestAgent) GetFreezeFsOnBackupOk() (*bool, bool)`

GetFreezeFsOnBackupOk returns a tuple with the FreezeFsOnBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeFsOnBackup

`func (o *CreateVMRequestAgent) SetFreezeFsOnBackup(v bool)`

SetFreezeFsOnBackup sets FreezeFsOnBackup field to given value.

### HasFreezeFsOnBackup

`func (o *CreateVMRequestAgent) HasFreezeFsOnBackup() bool`

HasFreezeFsOnBackup returns a boolean if a field has been set.

### GetFstrimClonedDisks

`func (o *CreateVMRequestAgent) GetFstrimClonedDisks() bool`

GetFstrimClonedDisks returns the FstrimClonedDisks field if non-nil, zero value otherwise.

### GetFstrimClonedDisksOk

`func (o *CreateVMRequestAgent) GetFstrimClonedDisksOk() (*bool, bool)`

GetFstrimClonedDisksOk returns a tuple with the FstrimClonedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFstrimClonedDisks

`func (o *CreateVMRequestAgent) SetFstrimClonedDisks(v bool)`

SetFstrimClonedDisks sets FstrimClonedDisks field to given value.

### HasFstrimClonedDisks

`func (o *CreateVMRequestAgent) HasFstrimClonedDisks() bool`

HasFstrimClonedDisks returns a boolean if a field has been set.

### GetType

`func (o *CreateVMRequestAgent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateVMRequestAgent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateVMRequestAgent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateVMRequestAgent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


