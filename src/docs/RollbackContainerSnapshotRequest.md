# RollbackContainerSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** | Whether the container should get started after rolling back successfully | [optional] 

## Methods

### NewRollbackContainerSnapshotRequest

`func NewRollbackContainerSnapshotRequest() *RollbackContainerSnapshotRequest`

NewRollbackContainerSnapshotRequest instantiates a new RollbackContainerSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackContainerSnapshotRequestWithDefaults

`func NewRollbackContainerSnapshotRequestWithDefaults() *RollbackContainerSnapshotRequest`

NewRollbackContainerSnapshotRequestWithDefaults instantiates a new RollbackContainerSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *RollbackContainerSnapshotRequest) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *RollbackContainerSnapshotRequest) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *RollbackContainerSnapshotRequest) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *RollbackContainerSnapshotRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


