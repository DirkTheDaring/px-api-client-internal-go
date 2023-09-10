# CreateNodesSingleStorageSingleContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | ***os.File** | The name of the file to create. | 
**Format** | Pointer to **string** |  | [optional] 
**Size** | **string** | Size in kilobyte (1024 bytes). Optional suffixes &#39;M&#39; (megabyte, 1024K) and &#39;G&#39; (gigabyte, 1024M) | 
**Vmid** | **int64** | Specify owner VM | 

## Methods

### NewCreateNodesSingleStorageSingleContentRequest

`func NewCreateNodesSingleStorageSingleContentRequest(filename *os.File, size string, vmid int64, ) *CreateNodesSingleStorageSingleContentRequest`

NewCreateNodesSingleStorageSingleContentRequest instantiates a new CreateNodesSingleStorageSingleContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNodesSingleStorageSingleContentRequestWithDefaults

`func NewCreateNodesSingleStorageSingleContentRequestWithDefaults() *CreateNodesSingleStorageSingleContentRequest`

NewCreateNodesSingleStorageSingleContentRequestWithDefaults instantiates a new CreateNodesSingleStorageSingleContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *CreateNodesSingleStorageSingleContentRequest) GetFilename() *os.File`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateNodesSingleStorageSingleContentRequest) GetFilenameOk() (**os.File, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateNodesSingleStorageSingleContentRequest) SetFilename(v *os.File)`

SetFilename sets Filename field to given value.


### GetFormat

`func (o *CreateNodesSingleStorageSingleContentRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateNodesSingleStorageSingleContentRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateNodesSingleStorageSingleContentRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateNodesSingleStorageSingleContentRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSize

`func (o *CreateNodesSingleStorageSingleContentRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateNodesSingleStorageSingleContentRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateNodesSingleStorageSingleContentRequest) SetSize(v string)`

SetSize sets Size field to given value.


### GetVmid

`func (o *CreateNodesSingleStorageSingleContentRequest) GetVmid() int64`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *CreateNodesSingleStorageSingleContentRequest) GetVmidOk() (*int64, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *CreateNodesSingleStorageSingleContentRequest) SetVmid(v int64)`

SetVmid sets Vmid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


