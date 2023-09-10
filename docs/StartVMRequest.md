# StartVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceCpu** | Pointer to **string** | Override QEMU&#39;s -cpu argument with the given string. | [optional] 
**Machine** | Pointer to **string** | Specifies the QEMU machine type. | [optional] 
**Migratedfrom** | Pointer to **string** | The cluster node name. | [optional] 
**MigrationNetwork** | Pointer to **string** | CIDR of the (sub) network that is used for migration. | [optional] 
**MigrationType** | Pointer to **string** | Migration traffic is encrypted using an SSH tunnel by default. On secure, completely private networks this can be disabled to increase performance. | [optional] 
**Skiplock** | Pointer to **int32** | Ignore locks - only root is allowed to use this option. | [optional] 
**Stateuri** | Pointer to **string** | Some command save/restore state from this location. | [optional] 
**Targetstorage** | Pointer to **string** | Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value &#39;1&#39; will map each source storage to itself. | [optional] 
**Timeout** | Pointer to **int64** | Wait maximal timeout seconds. | [optional] 

## Methods

### NewStartVMRequest

`func NewStartVMRequest() *StartVMRequest`

NewStartVMRequest instantiates a new StartVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartVMRequestWithDefaults

`func NewStartVMRequestWithDefaults() *StartVMRequest`

NewStartVMRequestWithDefaults instantiates a new StartVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceCpu

`func (o *StartVMRequest) GetForceCpu() string`

GetForceCpu returns the ForceCpu field if non-nil, zero value otherwise.

### GetForceCpuOk

`func (o *StartVMRequest) GetForceCpuOk() (*string, bool)`

GetForceCpuOk returns a tuple with the ForceCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCpu

`func (o *StartVMRequest) SetForceCpu(v string)`

SetForceCpu sets ForceCpu field to given value.

### HasForceCpu

`func (o *StartVMRequest) HasForceCpu() bool`

HasForceCpu returns a boolean if a field has been set.

### GetMachine

`func (o *StartVMRequest) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *StartVMRequest) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *StartVMRequest) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *StartVMRequest) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMigratedfrom

`func (o *StartVMRequest) GetMigratedfrom() string`

GetMigratedfrom returns the Migratedfrom field if non-nil, zero value otherwise.

### GetMigratedfromOk

`func (o *StartVMRequest) GetMigratedfromOk() (*string, bool)`

GetMigratedfromOk returns a tuple with the Migratedfrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedfrom

`func (o *StartVMRequest) SetMigratedfrom(v string)`

SetMigratedfrom sets Migratedfrom field to given value.

### HasMigratedfrom

`func (o *StartVMRequest) HasMigratedfrom() bool`

HasMigratedfrom returns a boolean if a field has been set.

### GetMigrationNetwork

`func (o *StartVMRequest) GetMigrationNetwork() string`

GetMigrationNetwork returns the MigrationNetwork field if non-nil, zero value otherwise.

### GetMigrationNetworkOk

`func (o *StartVMRequest) GetMigrationNetworkOk() (*string, bool)`

GetMigrationNetworkOk returns a tuple with the MigrationNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationNetwork

`func (o *StartVMRequest) SetMigrationNetwork(v string)`

SetMigrationNetwork sets MigrationNetwork field to given value.

### HasMigrationNetwork

`func (o *StartVMRequest) HasMigrationNetwork() bool`

HasMigrationNetwork returns a boolean if a field has been set.

### GetMigrationType

`func (o *StartVMRequest) GetMigrationType() string`

GetMigrationType returns the MigrationType field if non-nil, zero value otherwise.

### GetMigrationTypeOk

`func (o *StartVMRequest) GetMigrationTypeOk() (*string, bool)`

GetMigrationTypeOk returns a tuple with the MigrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationType

`func (o *StartVMRequest) SetMigrationType(v string)`

SetMigrationType sets MigrationType field to given value.

### HasMigrationType

`func (o *StartVMRequest) HasMigrationType() bool`

HasMigrationType returns a boolean if a field has been set.

### GetSkiplock

`func (o *StartVMRequest) GetSkiplock() int32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *StartVMRequest) GetSkiplockOk() (*int32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *StartVMRequest) SetSkiplock(v int32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *StartVMRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetStateuri

`func (o *StartVMRequest) GetStateuri() string`

GetStateuri returns the Stateuri field if non-nil, zero value otherwise.

### GetStateuriOk

`func (o *StartVMRequest) GetStateuriOk() (*string, bool)`

GetStateuriOk returns a tuple with the Stateuri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateuri

`func (o *StartVMRequest) SetStateuri(v string)`

SetStateuri sets Stateuri field to given value.

### HasStateuri

`func (o *StartVMRequest) HasStateuri() bool`

HasStateuri returns a boolean if a field has been set.

### GetTargetstorage

`func (o *StartVMRequest) GetTargetstorage() string`

GetTargetstorage returns the Targetstorage field if non-nil, zero value otherwise.

### GetTargetstorageOk

`func (o *StartVMRequest) GetTargetstorageOk() (*string, bool)`

GetTargetstorageOk returns a tuple with the Targetstorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetstorage

`func (o *StartVMRequest) SetTargetstorage(v string)`

SetTargetstorage sets Targetstorage field to given value.

### HasTargetstorage

`func (o *StartVMRequest) HasTargetstorage() bool`

HasTargetstorage returns a boolean if a field has been set.

### GetTimeout

`func (o *StartVMRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *StartVMRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *StartVMRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *StartVMRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


