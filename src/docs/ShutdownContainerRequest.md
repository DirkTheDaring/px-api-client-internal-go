# ShutdownContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceStop** | Pointer to **int32** | Make sure the Container stops. | [optional] 
**Timeout** | Pointer to **int64** | Wait maximal timeout seconds. | [optional] 

## Methods

### NewShutdownContainerRequest

`func NewShutdownContainerRequest() *ShutdownContainerRequest`

NewShutdownContainerRequest instantiates a new ShutdownContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutdownContainerRequestWithDefaults

`func NewShutdownContainerRequestWithDefaults() *ShutdownContainerRequest`

NewShutdownContainerRequestWithDefaults instantiates a new ShutdownContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceStop

`func (o *ShutdownContainerRequest) GetForceStop() int32`

GetForceStop returns the ForceStop field if non-nil, zero value otherwise.

### GetForceStopOk

`func (o *ShutdownContainerRequest) GetForceStopOk() (*int32, bool)`

GetForceStopOk returns a tuple with the ForceStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceStop

`func (o *ShutdownContainerRequest) SetForceStop(v int32)`

SetForceStop sets ForceStop field to given value.

### HasForceStop

`func (o *ShutdownContainerRequest) HasForceStop() bool`

HasForceStop returns a boolean if a field has been set.

### GetTimeout

`func (o *ShutdownContainerRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ShutdownContainerRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ShutdownContainerRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ShutdownContainerRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


