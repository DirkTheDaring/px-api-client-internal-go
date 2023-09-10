# GetVMConfig200ResponseDataSpiceEnhancements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Foldersharing** | Pointer to **int32** | Enable folder sharing via SPICE. Needs Spice-WebDAV daemon installed in the VM. | [optional] 
**Videostreaming** | Pointer to **string** | Enable video streaming. Uses compression for detected video streams. | [optional] 

## Methods

### NewGetVMConfig200ResponseDataSpiceEnhancements

`func NewGetVMConfig200ResponseDataSpiceEnhancements() *GetVMConfig200ResponseDataSpiceEnhancements`

NewGetVMConfig200ResponseDataSpiceEnhancements instantiates a new GetVMConfig200ResponseDataSpiceEnhancements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataSpiceEnhancementsWithDefaults

`func NewGetVMConfig200ResponseDataSpiceEnhancementsWithDefaults() *GetVMConfig200ResponseDataSpiceEnhancements`

NewGetVMConfig200ResponseDataSpiceEnhancementsWithDefaults instantiates a new GetVMConfig200ResponseDataSpiceEnhancements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFoldersharing

`func (o *GetVMConfig200ResponseDataSpiceEnhancements) GetFoldersharing() int32`

GetFoldersharing returns the Foldersharing field if non-nil, zero value otherwise.

### GetFoldersharingOk

`func (o *GetVMConfig200ResponseDataSpiceEnhancements) GetFoldersharingOk() (*int32, bool)`

GetFoldersharingOk returns a tuple with the Foldersharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldersharing

`func (o *GetVMConfig200ResponseDataSpiceEnhancements) SetFoldersharing(v int32)`

SetFoldersharing sets Foldersharing field to given value.

### HasFoldersharing

`func (o *GetVMConfig200ResponseDataSpiceEnhancements) HasFoldersharing() bool`

HasFoldersharing returns a boolean if a field has been set.

### GetVideostreaming

`func (o *GetVMConfig200ResponseDataSpiceEnhancements) GetVideostreaming() string`

GetVideostreaming returns the Videostreaming field if non-nil, zero value otherwise.

### GetVideostreamingOk

`func (o *GetVMConfig200ResponseDataSpiceEnhancements) GetVideostreamingOk() (*string, bool)`

GetVideostreamingOk returns a tuple with the Videostreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideostreaming

`func (o *GetVMConfig200ResponseDataSpiceEnhancements) SetVideostreaming(v string)`

SetVideostreaming sets Videostreaming field to given value.

### HasVideostreaming

`func (o *GetVMConfig200ResponseDataSpiceEnhancements) HasVideostreaming() bool`

HasVideostreaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


