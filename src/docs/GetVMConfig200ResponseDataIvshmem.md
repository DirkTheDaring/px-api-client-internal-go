# GetVMConfig200ResponseDataIvshmem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the file. Will be prefixed with &#39;pve-shm-&#39;. Default is the VMID. Will be deleted when the VM is stopped. | [optional] 
**Size** | Pointer to **int64** | The size of the file in MB. | [optional] 

## Methods

### NewGetVMConfig200ResponseDataIvshmem

`func NewGetVMConfig200ResponseDataIvshmem() *GetVMConfig200ResponseDataIvshmem`

NewGetVMConfig200ResponseDataIvshmem instantiates a new GetVMConfig200ResponseDataIvshmem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataIvshmemWithDefaults

`func NewGetVMConfig200ResponseDataIvshmemWithDefaults() *GetVMConfig200ResponseDataIvshmem`

NewGetVMConfig200ResponseDataIvshmemWithDefaults instantiates a new GetVMConfig200ResponseDataIvshmem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetVMConfig200ResponseDataIvshmem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVMConfig200ResponseDataIvshmem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVMConfig200ResponseDataIvshmem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVMConfig200ResponseDataIvshmem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *GetVMConfig200ResponseDataIvshmem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetVMConfig200ResponseDataIvshmem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetVMConfig200ResponseDataIvshmem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetVMConfig200ResponseDataIvshmem) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


