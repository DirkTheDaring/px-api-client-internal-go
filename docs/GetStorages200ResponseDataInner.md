# GetStorages200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | Set when storage is accessible. | [optional] 
**Avail** | Pointer to **int64** | Available storage space in bytes. | [optional] 
**Content** | Pointer to **string** | Allowed storage content types. | [optional] 
**Enabled** | Pointer to **int32** | Set when storage is enabled (not disabled). | [optional] 
**Shared** | Pointer to **int32** | Shared flag from storage configuration. | [optional] 
**Storage** | Pointer to **string** | The storage identifier. | [optional] 
**Total** | Pointer to **int64** | Total storage space in bytes. | [optional] 
**Type** | Pointer to **string** | Storage type. | [optional] 
**Used** | Pointer to **int64** | Used storage space in bytes. | [optional] 
**UsedFraction** | Pointer to **float32** | Used fraction (used/total). | [optional] 

## Methods

### NewGetStorages200ResponseDataInner

`func NewGetStorages200ResponseDataInner() *GetStorages200ResponseDataInner`

NewGetStorages200ResponseDataInner instantiates a new GetStorages200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorages200ResponseDataInnerWithDefaults

`func NewGetStorages200ResponseDataInnerWithDefaults() *GetStorages200ResponseDataInner`

NewGetStorages200ResponseDataInnerWithDefaults instantiates a new GetStorages200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GetStorages200ResponseDataInner) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetStorages200ResponseDataInner) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetStorages200ResponseDataInner) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetStorages200ResponseDataInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAvail

`func (o *GetStorages200ResponseDataInner) GetAvail() int64`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *GetStorages200ResponseDataInner) GetAvailOk() (*int64, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *GetStorages200ResponseDataInner) SetAvail(v int64)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *GetStorages200ResponseDataInner) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetContent

`func (o *GetStorages200ResponseDataInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetStorages200ResponseDataInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetStorages200ResponseDataInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetStorages200ResponseDataInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEnabled

`func (o *GetStorages200ResponseDataInner) GetEnabled() int32`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetStorages200ResponseDataInner) GetEnabledOk() (*int32, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetStorages200ResponseDataInner) SetEnabled(v int32)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetStorages200ResponseDataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetShared

`func (o *GetStorages200ResponseDataInner) GetShared() int32`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *GetStorages200ResponseDataInner) GetSharedOk() (*int32, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *GetStorages200ResponseDataInner) SetShared(v int32)`

SetShared sets Shared field to given value.

### HasShared

`func (o *GetStorages200ResponseDataInner) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetStorage

`func (o *GetStorages200ResponseDataInner) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *GetStorages200ResponseDataInner) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *GetStorages200ResponseDataInner) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *GetStorages200ResponseDataInner) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTotal

`func (o *GetStorages200ResponseDataInner) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetStorages200ResponseDataInner) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetStorages200ResponseDataInner) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetStorages200ResponseDataInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *GetStorages200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetStorages200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetStorages200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetStorages200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsed

`func (o *GetStorages200ResponseDataInner) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *GetStorages200ResponseDataInner) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *GetStorages200ResponseDataInner) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *GetStorages200ResponseDataInner) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUsedFraction

`func (o *GetStorages200ResponseDataInner) GetUsedFraction() float32`

GetUsedFraction returns the UsedFraction field if non-nil, zero value otherwise.

### GetUsedFractionOk

`func (o *GetStorages200ResponseDataInner) GetUsedFractionOk() (*float32, bool)`

GetUsedFractionOk returns a tuple with the UsedFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFraction

`func (o *GetStorages200ResponseDataInner) SetUsedFraction(v float32)`

SetUsedFraction sets UsedFraction field to given value.

### HasUsedFraction

`func (o *GetStorages200ResponseDataInner) HasUsedFraction() bool`

HasUsedFraction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


