# StopVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepActive** | Pointer to **bool** |  | [optional] 
**Migratedfrom** | Pointer to **string** |  | [optional] 
**OverruleShutdown** | Pointer to **bool** |  | [optional] 
**Skiplock** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 

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

`func (o *StopVMRequest) GetKeepActive() bool`

GetKeepActive returns the KeepActive field if non-nil, zero value otherwise.

### GetKeepActiveOk

`func (o *StopVMRequest) GetKeepActiveOk() (*bool, bool)`

GetKeepActiveOk returns a tuple with the KeepActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepActive

`func (o *StopVMRequest) SetKeepActive(v bool)`

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

### GetOverruleShutdown

`func (o *StopVMRequest) GetOverruleShutdown() bool`

GetOverruleShutdown returns the OverruleShutdown field if non-nil, zero value otherwise.

### GetOverruleShutdownOk

`func (o *StopVMRequest) GetOverruleShutdownOk() (*bool, bool)`

GetOverruleShutdownOk returns a tuple with the OverruleShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverruleShutdown

`func (o *StopVMRequest) SetOverruleShutdown(v bool)`

SetOverruleShutdown sets OverruleShutdown field to given value.

### HasOverruleShutdown

`func (o *StopVMRequest) HasOverruleShutdown() bool`

HasOverruleShutdown returns a boolean if a field has been set.

### GetSkiplock

`func (o *StopVMRequest) GetSkiplock() bool`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *StopVMRequest) GetSkiplockOk() (*bool, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *StopVMRequest) SetSkiplock(v bool)`

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


