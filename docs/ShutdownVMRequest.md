# ShutdownVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceStop** | Pointer to **bool** |  | [optional] 
**KeepActive** | Pointer to **bool** |  | [optional] 
**Skiplock** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 

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

`func (o *ShutdownVMRequest) GetForceStop() bool`

GetForceStop returns the ForceStop field if non-nil, zero value otherwise.

### GetForceStopOk

`func (o *ShutdownVMRequest) GetForceStopOk() (*bool, bool)`

GetForceStopOk returns a tuple with the ForceStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceStop

`func (o *ShutdownVMRequest) SetForceStop(v bool)`

SetForceStop sets ForceStop field to given value.

### HasForceStop

`func (o *ShutdownVMRequest) HasForceStop() bool`

HasForceStop returns a boolean if a field has been set.

### GetKeepActive

`func (o *ShutdownVMRequest) GetKeepActive() bool`

GetKeepActive returns the KeepActive field if non-nil, zero value otherwise.

### GetKeepActiveOk

`func (o *ShutdownVMRequest) GetKeepActiveOk() (*bool, bool)`

GetKeepActiveOk returns a tuple with the KeepActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepActive

`func (o *ShutdownVMRequest) SetKeepActive(v bool)`

SetKeepActive sets KeepActive field to given value.

### HasKeepActive

`func (o *ShutdownVMRequest) HasKeepActive() bool`

HasKeepActive returns a boolean if a field has been set.

### GetSkiplock

`func (o *ShutdownVMRequest) GetSkiplock() bool`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *ShutdownVMRequest) GetSkiplockOk() (*bool, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *ShutdownVMRequest) SetSkiplock(v bool)`

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


