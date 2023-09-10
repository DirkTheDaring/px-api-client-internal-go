# GetContainerSnapshots200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Snapshot description. | [optional] 
**Name** | Pointer to **string** | Snapshot identifier. Value &#39;current&#39; identifies the current VM. | [optional] 
**Parent** | Pointer to **string** | Parent snapshot identifier. | [optional] 
**Snaptime** | Pointer to **int64** | Snapshot creation time | [optional] 

## Methods

### NewGetContainerSnapshots200ResponseDataInner

`func NewGetContainerSnapshots200ResponseDataInner() *GetContainerSnapshots200ResponseDataInner`

NewGetContainerSnapshots200ResponseDataInner instantiates a new GetContainerSnapshots200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerSnapshots200ResponseDataInnerWithDefaults

`func NewGetContainerSnapshots200ResponseDataInnerWithDefaults() *GetContainerSnapshots200ResponseDataInner`

NewGetContainerSnapshots200ResponseDataInnerWithDefaults instantiates a new GetContainerSnapshots200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetContainerSnapshots200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetContainerSnapshots200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetContainerSnapshots200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetContainerSnapshots200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *GetContainerSnapshots200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetContainerSnapshots200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetContainerSnapshots200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetContainerSnapshots200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *GetContainerSnapshots200ResponseDataInner) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetContainerSnapshots200ResponseDataInner) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetContainerSnapshots200ResponseDataInner) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetContainerSnapshots200ResponseDataInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSnaptime

`func (o *GetContainerSnapshots200ResponseDataInner) GetSnaptime() int64`

GetSnaptime returns the Snaptime field if non-nil, zero value otherwise.

### GetSnaptimeOk

`func (o *GetContainerSnapshots200ResponseDataInner) GetSnaptimeOk() (*int64, bool)`

GetSnaptimeOk returns a tuple with the Snaptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnaptime

`func (o *GetContainerSnapshots200ResponseDataInner) SetSnaptime(v int64)`

SetSnaptime sets Snaptime field to given value.

### HasSnaptime

`func (o *GetContainerSnapshots200ResponseDataInner) HasSnaptime() bool`

HasSnaptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


