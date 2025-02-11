# GetCurrentVMStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**GetCurrentVMStatus200ResponseData**](GetCurrentVMStatus200ResponseData.md) |  | [optional] 

## Methods

### NewGetCurrentVMStatus200Response

`func NewGetCurrentVMStatus200Response() *GetCurrentVMStatus200Response`

NewGetCurrentVMStatus200Response instantiates a new GetCurrentVMStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentVMStatus200ResponseWithDefaults

`func NewGetCurrentVMStatus200ResponseWithDefaults() *GetCurrentVMStatus200Response`

NewGetCurrentVMStatus200ResponseWithDefaults instantiates a new GetCurrentVMStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GetCurrentVMStatus200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetCurrentVMStatus200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetCurrentVMStatus200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetCurrentVMStatus200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *GetCurrentVMStatus200Response) GetData() GetCurrentVMStatus200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCurrentVMStatus200Response) GetDataOk() (*GetCurrentVMStatus200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCurrentVMStatus200Response) SetData(v GetCurrentVMStatus200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetCurrentVMStatus200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


