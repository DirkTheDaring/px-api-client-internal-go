# StopVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepActive** | Pointer to **int32** | Do not deactivate storage volumes. | [optional] 
**Migratedfrom** | Pointer to **string** | The cluster node name. | [optional] 
**Skiplock** | Pointer to **int32** | Ignore locks - only root is allowed to use this option. | [optional] 
**Timeout** | Pointer to **int64** | Wait maximal timeout seconds. | [optional] 

## Methods

### NewStopVMRequest

`func NewStopVMRequest() *StopVMRequest`

NewStopVMRequest instantiates a new StopVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopVMRequestWithDefaults

`func NewStopVMRequestWithDefaults() *StopVMRequest`

NewStopVMRequestWithDefaults instantiates a new StopVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepActive

`func (o *StopVMRequest) GetKeepActive() int32`

GetKeepActive returns the KeepActive field if non-nil, zero value otherwise.

### GetKeepActiveOk

`func (o *StopVMRequest) GetKeepActiveOk() (*int32, bool)`

GetKeepActiveOk returns a tuple with the KeepActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepActive

`func (o *StopVMRequest) SetKeepActive(v int32)`

SetKeepActive sets KeepActive field to given value.

### HasKeepActive

`func (o *StopVMRequest) HasKeepActive() bool`

HasKeepActive returns a boolean if a field has been set.

### GetMigratedfrom

`func (o *StopVMRequest) GetMigratedfrom() string`

GetMigratedfrom returns the Migratedfrom field if non-nil, zero value otherwise.

### GetMigratedfromOk

`func (o *StopVMRequest) GetMigratedfromOk() (*string, bool)`

GetMigratedfromOk returns a tuple with the Migratedfrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedfrom

`func (o *StopVMRequest) SetMigratedfrom(v string)`

SetMigratedfrom sets Migratedfrom field to given value.

### HasMigratedfrom

`func (o *StopVMRequest) HasMigratedfrom() bool`

HasMigratedfrom returns a boolean if a field has been set.

### GetSkiplock

`func (o *StopVMRequest) GetSkiplock() int32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *StopVMRequest) GetSkiplockOk() (*int32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *StopVMRequest) SetSkiplock(v int32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *StopVMRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetTimeout

`func (o *StopVMRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *StopVMRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *StopVMRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *StopVMRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


