# ResizeVMDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** |  | [optional] 
**Disk** | **string** |  | 
**Size** | **string** |  | 
**Skiplock** | Pointer to **bool** |  | [optional] 

## Methods

### NewResizeVMDiskRequest

`func NewResizeVMDiskRequest(disk string, size string, ) *ResizeVMDiskRequest`

NewResizeVMDiskRequest instantiates a new ResizeVMDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeVMDiskRequestWithDefaults

`func NewResizeVMDiskRequestWithDefaults() *ResizeVMDiskRequest`

NewResizeVMDiskRequestWithDefaults instantiates a new ResizeVMDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *ResizeVMDiskRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ResizeVMDiskRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ResizeVMDiskRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ResizeVMDiskRequest) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDisk

`func (o *ResizeVMDiskRequest) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ResizeVMDiskRequest) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ResizeVMDiskRequest) SetDisk(v string)`

SetDisk sets Disk field to given value.


### GetSize

`func (o *ResizeVMDiskRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResizeVMDiskRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResizeVMDiskRequest) SetSize(v string)`

SetSize sets Size field to given value.


### GetSkiplock

`func (o *ResizeVMDiskRequest) GetSkiplock() bool`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *ResizeVMDiskRequest) GetSkiplockOk() (*bool, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *ResizeVMDiskRequest) SetSkiplock(v bool)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *ResizeVMDiskRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


