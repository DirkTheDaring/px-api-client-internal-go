# StopContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skiplock** | Pointer to **int32** | Ignore locks - only root is allowed to use this option. | [optional] 

## Methods

### NewStopContainerRequest

`func NewStopContainerRequest() *StopContainerRequest`

NewStopContainerRequest instantiates a new StopContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopContainerRequestWithDefaults

`func NewStopContainerRequestWithDefaults() *StopContainerRequest`

NewStopContainerRequestWithDefaults instantiates a new StopContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkiplock

`func (o *StopContainerRequest) GetSkiplock() int32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *StopContainerRequest) GetSkiplockOk() (*int32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *StopContainerRequest) SetSkiplock(v int32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *StopContainerRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


