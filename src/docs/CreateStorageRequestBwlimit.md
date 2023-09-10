# CreateStorageRequestBwlimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clone** | Pointer to **float32** | bandwidth limit in KiB/s for cloning disks | [optional] 
**Default** | Pointer to **float32** | default bandwidth limit in KiB/s | [optional] 
**Migration** | Pointer to **float32** | bandwidth limit in KiB/s for migrating guests (including moving local disks) | [optional] 
**Move** | Pointer to **float32** | bandwidth limit in KiB/s for moving disks | [optional] 
**Restore** | Pointer to **float32** | bandwidth limit in KiB/s for restoring guests from backups | [optional] 

## Methods

### NewCreateStorageRequestBwlimit

`func NewCreateStorageRequestBwlimit() *CreateStorageRequestBwlimit`

NewCreateStorageRequestBwlimit instantiates a new CreateStorageRequestBwlimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageRequestBwlimitWithDefaults

`func NewCreateStorageRequestBwlimitWithDefaults() *CreateStorageRequestBwlimit`

NewCreateStorageRequestBwlimitWithDefaults instantiates a new CreateStorageRequestBwlimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClone

`func (o *CreateStorageRequestBwlimit) GetClone() float32`

GetClone returns the Clone field if non-nil, zero value otherwise.

### GetCloneOk

`func (o *CreateStorageRequestBwlimit) GetCloneOk() (*float32, bool)`

GetCloneOk returns a tuple with the Clone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClone

`func (o *CreateStorageRequestBwlimit) SetClone(v float32)`

SetClone sets Clone field to given value.

### HasClone

`func (o *CreateStorageRequestBwlimit) HasClone() bool`

HasClone returns a boolean if a field has been set.

### GetDefault

`func (o *CreateStorageRequestBwlimit) GetDefault() float32`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CreateStorageRequestBwlimit) GetDefaultOk() (*float32, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CreateStorageRequestBwlimit) SetDefault(v float32)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CreateStorageRequestBwlimit) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetMigration

`func (o *CreateStorageRequestBwlimit) GetMigration() float32`

GetMigration returns the Migration field if non-nil, zero value otherwise.

### GetMigrationOk

`func (o *CreateStorageRequestBwlimit) GetMigrationOk() (*float32, bool)`

GetMigrationOk returns a tuple with the Migration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigration

`func (o *CreateStorageRequestBwlimit) SetMigration(v float32)`

SetMigration sets Migration field to given value.

### HasMigration

`func (o *CreateStorageRequestBwlimit) HasMigration() bool`

HasMigration returns a boolean if a field has been set.

### GetMove

`func (o *CreateStorageRequestBwlimit) GetMove() float32`

GetMove returns the Move field if non-nil, zero value otherwise.

### GetMoveOk

`func (o *CreateStorageRequestBwlimit) GetMoveOk() (*float32, bool)`

GetMoveOk returns a tuple with the Move field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMove

`func (o *CreateStorageRequestBwlimit) SetMove(v float32)`

SetMove sets Move field to given value.

### HasMove

`func (o *CreateStorageRequestBwlimit) HasMove() bool`

HasMove returns a boolean if a field has been set.

### GetRestore

`func (o *CreateStorageRequestBwlimit) GetRestore() float32`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *CreateStorageRequestBwlimit) GetRestoreOk() (*float32, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *CreateStorageRequestBwlimit) SetRestore(v float32)`

SetRestore sets Restore field to given value.

### HasRestore

`func (o *CreateStorageRequestBwlimit) HasRestore() bool`

HasRestore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


