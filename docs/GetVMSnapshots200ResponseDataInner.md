# GetVMSnapshots200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Snapshot description. | [optional] 
**Name** | Pointer to **string** | Snapshot identifier. Value &#39;current&#39; identifies the current VM. | [optional] 
**Parent** | Pointer to **string** | Parent snapshot identifier. | [optional] 
**Snaptime** | Pointer to **int64** | Snapshot creation time | [optional] 
**Vmstate** | Pointer to **int32** | Snapshot includes RAM. | [optional] 

## Methods

### NewGetVMSnapshots200ResponseDataInner

`func NewGetVMSnapshots200ResponseDataInner() *GetVMSnapshots200ResponseDataInner`

NewGetVMSnapshots200ResponseDataInner instantiates a new GetVMSnapshots200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMSnapshots200ResponseDataInnerWithDefaults

`func NewGetVMSnapshots200ResponseDataInnerWithDefaults() *GetVMSnapshots200ResponseDataInner`

NewGetVMSnapshots200ResponseDataInnerWithDefaults instantiates a new GetVMSnapshots200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetVMSnapshots200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVMSnapshots200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVMSnapshots200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVMSnapshots200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *GetVMSnapshots200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVMSnapshots200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVMSnapshots200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVMSnapshots200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *GetVMSnapshots200ResponseDataInner) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetVMSnapshots200ResponseDataInner) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetVMSnapshots200ResponseDataInner) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetVMSnapshots200ResponseDataInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSnaptime

`func (o *GetVMSnapshots200ResponseDataInner) GetSnaptime() int64`

GetSnaptime returns the Snaptime field if non-nil, zero value otherwise.

### GetSnaptimeOk

`func (o *GetVMSnapshots200ResponseDataInner) GetSnaptimeOk() (*int64, bool)`

GetSnaptimeOk returns a tuple with the Snaptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnaptime

`func (o *GetVMSnapshots200ResponseDataInner) SetSnaptime(v int64)`

SetSnaptime sets Snaptime field to given value.

### HasSnaptime

`func (o *GetVMSnapshots200ResponseDataInner) HasSnaptime() bool`

HasSnaptime returns a boolean if a field has been set.

### GetVmstate

`func (o *GetVMSnapshots200ResponseDataInner) GetVmstate() int32`

GetVmstate returns the Vmstate field if non-nil, zero value otherwise.

### GetVmstateOk

`func (o *GetVMSnapshots200ResponseDataInner) GetVmstateOk() (*int32, bool)`

GetVmstateOk returns a tuple with the Vmstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstate

`func (o *GetVMSnapshots200ResponseDataInner) SetVmstate(v int32)`

SetVmstate sets Vmstate field to given value.

### HasVmstate

`func (o *GetVMSnapshots200ResponseDataInner) HasVmstate() bool`

HasVmstate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


