# ResizeContainerDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** |  | [optional] 
**Disk** | **string** |  | 
**Size** | **string** |  | 

## Methods

### NewResizeContainerDiskRequest

`func NewResizeContainerDiskRequest(disk string, size string, ) *ResizeContainerDiskRequest`

NewResizeContainerDiskRequest instantiates a new ResizeContainerDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeContainerDiskRequestWithDefaults

`func NewResizeContainerDiskRequestWithDefaults() *ResizeContainerDiskRequest`

NewResizeContainerDiskRequestWithDefaults instantiates a new ResizeContainerDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *ResizeContainerDiskRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ResizeContainerDiskRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ResizeContainerDiskRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ResizeContainerDiskRequest) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDisk

`func (o *ResizeContainerDiskRequest) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ResizeContainerDiskRequest) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ResizeContainerDiskRequest) SetDisk(v string)`

SetDisk sets Disk field to given value.


### GetSize

`func (o *ResizeContainerDiskRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResizeContainerDiskRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResizeContainerDiskRequest) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


