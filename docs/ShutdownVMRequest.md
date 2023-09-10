# ShutdownVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceStop** | Pointer to **int32** | Make sure the VM stops. | [optional] 
**KeepActive** | Pointer to **int32** | Do not deactivate storage volumes. | [optional] 
**Skiplock** | Pointer to **int32** | Ignore locks - only root is allowed to use this option. | [optional] 
**Timeout** | Pointer to **int64** | Wait maximal timeout seconds. | [optional] 

## Methods

### NewShutdownVMRequest

`func NewShutdownVMRequest() *ShutdownVMRequest`

NewShutdownVMRequest instantiates a new ShutdownVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutdownVMRequestWithDefaults

`func NewShutdownVMRequestWithDefaults() *ShutdownVMRequest`

NewShutdownVMRequestWithDefaults instantiates a new ShutdownVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceStop

`func (o *ShutdownVMRequest) GetForceStop() int32`

GetForceStop returns the ForceStop field if non-nil, zero value otherwise.

### GetForceStopOk

`func (o *ShutdownVMRequest) GetForceStopOk() (*int32, bool)`

GetForceStopOk returns a tuple with the ForceStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceStop

`func (o *ShutdownVMRequest) SetForceStop(v int32)`

SetForceStop sets ForceStop field to given value.

### HasForceStop

`func (o *ShutdownVMRequest) HasForceStop() bool`

HasForceStop returns a boolean if a field has been set.

### GetKeepActive

`func (o *ShutdownVMRequest) GetKeepActive() int32`

GetKeepActive returns the KeepActive field if non-nil, zero value otherwise.

### GetKeepActiveOk

`func (o *ShutdownVMRequest) GetKeepActiveOk() (*int32, bool)`

GetKeepActiveOk returns a tuple with the KeepActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepActive

`func (o *ShutdownVMRequest) SetKeepActive(v int32)`

SetKeepActive sets KeepActive field to given value.

### HasKeepActive

`func (o *ShutdownVMRequest) HasKeepActive() bool`

HasKeepActive returns a boolean if a field has been set.

### GetSkiplock

`func (o *ShutdownVMRequest) GetSkiplock() int32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *ShutdownVMRequest) GetSkiplockOk() (*int32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *ShutdownVMRequest) SetSkiplock(v int32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *ShutdownVMRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetTimeout

`func (o *ShutdownVMRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ShutdownVMRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ShutdownVMRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ShutdownVMRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


