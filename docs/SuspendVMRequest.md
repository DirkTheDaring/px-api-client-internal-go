# SuspendVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skiplock** | Pointer to **bool** | Ignore locks - only root is allowed to use this option. | [optional] 
**Statestorage** | Pointer to **string** | The storage for the VM state | [optional] 
**Todisk** | Pointer to **bool** | If set, suspends the VM to disk. Will be resumed on next VM start. | [optional] 

## Methods

### NewSuspendVMRequest

`func NewSuspendVMRequest() *SuspendVMRequest`

NewSuspendVMRequest instantiates a new SuspendVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuspendVMRequestWithDefaults

`func NewSuspendVMRequestWithDefaults() *SuspendVMRequest`

NewSuspendVMRequestWithDefaults instantiates a new SuspendVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkiplock

`func (o *SuspendVMRequest) GetSkiplock() bool`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *SuspendVMRequest) GetSkiplockOk() (*bool, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *SuspendVMRequest) SetSkiplock(v bool)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *SuspendVMRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetStatestorage

`func (o *SuspendVMRequest) GetStatestorage() string`

GetStatestorage returns the Statestorage field if non-nil, zero value otherwise.

### GetStatestorageOk

`func (o *SuspendVMRequest) GetStatestorageOk() (*string, bool)`

GetStatestorageOk returns a tuple with the Statestorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatestorage

`func (o *SuspendVMRequest) SetStatestorage(v string)`

SetStatestorage sets Statestorage field to given value.

### HasStatestorage

`func (o *SuspendVMRequest) HasStatestorage() bool`

HasStatestorage returns a boolean if a field has been set.

### GetTodisk

`func (o *SuspendVMRequest) GetTodisk() bool`

GetTodisk returns the Todisk field if non-nil, zero value otherwise.

### GetTodiskOk

`func (o *SuspendVMRequest) GetTodiskOk() (*bool, bool)`

GetTodiskOk returns a tuple with the Todisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodisk

`func (o *SuspendVMRequest) SetTodisk(v bool)`

SetTodisk sets Todisk field to given value.

### HasTodisk

`func (o *SuspendVMRequest) HasTodisk() bool`

HasTodisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


