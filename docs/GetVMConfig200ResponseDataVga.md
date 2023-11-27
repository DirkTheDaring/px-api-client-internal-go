# GetVMConfig200ResponseDataVga

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clipboard** | Pointer to **string** | Enable a specific clipboard. If not set, depending on the display type the SPICE one will be added. | [optional] 
**Memory** | Pointer to **int64** | Sets the VGA memory (in MiB). Has no effect with serial display. | [optional] 
**Type** | Pointer to **string** | Select the VGA type. | [optional] 

## Methods

### NewGetVMConfig200ResponseDataVga

`func NewGetVMConfig200ResponseDataVga() *GetVMConfig200ResponseDataVga`

NewGetVMConfig200ResponseDataVga instantiates a new GetVMConfig200ResponseDataVga object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataVgaWithDefaults

`func NewGetVMConfig200ResponseDataVgaWithDefaults() *GetVMConfig200ResponseDataVga`

NewGetVMConfig200ResponseDataVgaWithDefaults instantiates a new GetVMConfig200ResponseDataVga object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClipboard

`func (o *GetVMConfig200ResponseDataVga) GetClipboard() string`

GetClipboard returns the Clipboard field if non-nil, zero value otherwise.

### GetClipboardOk

`func (o *GetVMConfig200ResponseDataVga) GetClipboardOk() (*string, bool)`

GetClipboardOk returns a tuple with the Clipboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipboard

`func (o *GetVMConfig200ResponseDataVga) SetClipboard(v string)`

SetClipboard sets Clipboard field to given value.

### HasClipboard

`func (o *GetVMConfig200ResponseDataVga) HasClipboard() bool`

HasClipboard returns a boolean if a field has been set.

### GetMemory

`func (o *GetVMConfig200ResponseDataVga) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GetVMConfig200ResponseDataVga) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GetVMConfig200ResponseDataVga) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GetVMConfig200ResponseDataVga) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetType

`func (o *GetVMConfig200ResponseDataVga) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetVMConfig200ResponseDataVga) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetVMConfig200ResponseDataVga) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetVMConfig200ResponseDataVga) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


