# GetContainerConfigPending200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delete** | Pointer to **int64** | Indicates a pending delete request if present and not 0. | [optional] 
**Key** | Pointer to **string** | Configuration option name. | [optional] 
**Pending** | Pointer to **string** | Pending value. | [optional] 
**Value** | Pointer to **string** | Current value. | [optional] 

## Methods

### NewGetContainerConfigPending200ResponseDataInner

`func NewGetContainerConfigPending200ResponseDataInner() *GetContainerConfigPending200ResponseDataInner`

NewGetContainerConfigPending200ResponseDataInner instantiates a new GetContainerConfigPending200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerConfigPending200ResponseDataInnerWithDefaults

`func NewGetContainerConfigPending200ResponseDataInnerWithDefaults() *GetContainerConfigPending200ResponseDataInner`

NewGetContainerConfigPending200ResponseDataInnerWithDefaults instantiates a new GetContainerConfigPending200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *GetContainerConfigPending200ResponseDataInner) GetDelete() int64`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *GetContainerConfigPending200ResponseDataInner) GetDeleteOk() (*int64, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *GetContainerConfigPending200ResponseDataInner) SetDelete(v int64)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *GetContainerConfigPending200ResponseDataInner) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetKey

`func (o *GetContainerConfigPending200ResponseDataInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetContainerConfigPending200ResponseDataInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetContainerConfigPending200ResponseDataInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetContainerConfigPending200ResponseDataInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPending

`func (o *GetContainerConfigPending200ResponseDataInner) GetPending() string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GetContainerConfigPending200ResponseDataInner) GetPendingOk() (*string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GetContainerConfigPending200ResponseDataInner) SetPending(v string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *GetContainerConfigPending200ResponseDataInner) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetValue

`func (o *GetContainerConfigPending200ResponseDataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetContainerConfigPending200ResponseDataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetContainerConfigPending200ResponseDataInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetContainerConfigPending200ResponseDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


