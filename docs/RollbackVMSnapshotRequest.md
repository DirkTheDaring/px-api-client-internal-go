# RollbackVMSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **bool** | Whether the VM should get started after rolling back successfully. (Note: VMs will be automatically started if the snapshot includes RAM.) | [optional] 

## Methods

### NewRollbackVMSnapshotRequest

`func NewRollbackVMSnapshotRequest() *RollbackVMSnapshotRequest`

NewRollbackVMSnapshotRequest instantiates a new RollbackVMSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackVMSnapshotRequestWithDefaults

`func NewRollbackVMSnapshotRequestWithDefaults() *RollbackVMSnapshotRequest`

NewRollbackVMSnapshotRequestWithDefaults instantiates a new RollbackVMSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *RollbackVMSnapshotRequest) GetStart() bool`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *RollbackVMSnapshotRequest) GetStartOk() (*bool, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *RollbackVMSnapshotRequest) SetStart(v bool)`

SetStart sets Start field to given value.

### HasStart

`func (o *RollbackVMSnapshotRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


