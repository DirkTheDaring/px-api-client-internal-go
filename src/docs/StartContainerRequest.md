# StartContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Debug** | Pointer to **int32** | If set, enables very verbose debug log-level on start. | [optional] 
**Skiplock** | Pointer to **int32** | Ignore locks - only root is allowed to use this option. | [optional] 

## Methods

### NewStartContainerRequest

`func NewStartContainerRequest() *StartContainerRequest`

NewStartContainerRequest instantiates a new StartContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartContainerRequestWithDefaults

`func NewStartContainerRequestWithDefaults() *StartContainerRequest`

NewStartContainerRequestWithDefaults instantiates a new StartContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebug

`func (o *StartContainerRequest) GetDebug() int32`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *StartContainerRequest) GetDebugOk() (*int32, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *StartContainerRequest) SetDebug(v int32)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *StartContainerRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetSkiplock

`func (o *StartContainerRequest) GetSkiplock() int32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *StartContainerRequest) GetSkiplockOk() (*int32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *StartContainerRequest) SetSkiplock(v int32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *StartContainerRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


