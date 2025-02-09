# CreateVMRequestIvshmem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Size** | **int64** |  | 

## Methods

### NewCreateVMRequestIvshmem

`func NewCreateVMRequestIvshmem(size int64, ) *CreateVMRequestIvshmem`

NewCreateVMRequestIvshmem instantiates a new CreateVMRequestIvshmem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestIvshmemWithDefaults

`func NewCreateVMRequestIvshmemWithDefaults() *CreateVMRequestIvshmem`

NewCreateVMRequestIvshmemWithDefaults instantiates a new CreateVMRequestIvshmem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVMRequestIvshmem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVMRequestIvshmem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVMRequestIvshmem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateVMRequestIvshmem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *CreateVMRequestIvshmem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVMRequestIvshmem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVMRequestIvshmem) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


