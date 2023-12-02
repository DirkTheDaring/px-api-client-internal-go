# CreateVMRequestMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **int64** | Current amount of online RAM for the VM in MiB. This is the maximum available memory when you use the balloon device. | [optional] 

## Methods

### NewCreateVMRequestMemory

`func NewCreateVMRequestMemory() *CreateVMRequestMemory`

NewCreateVMRequestMemory instantiates a new CreateVMRequestMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestMemoryWithDefaults

`func NewCreateVMRequestMemoryWithDefaults() *CreateVMRequestMemory`

NewCreateVMRequestMemoryWithDefaults instantiates a new CreateVMRequestMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *CreateVMRequestMemory) GetCurrent() int64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CreateVMRequestMemory) GetCurrentOk() (*int64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CreateVMRequestMemory) SetCurrent(v int64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *CreateVMRequestMemory) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


