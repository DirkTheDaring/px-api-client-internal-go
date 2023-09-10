# GetVMConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetVMConfig200ResponseData**](GetVMConfig200ResponseData.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetVMConfig200Response

`func NewGetVMConfig200Response() *GetVMConfig200Response`

NewGetVMConfig200Response instantiates a new GetVMConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseWithDefaults

`func NewGetVMConfig200ResponseWithDefaults() *GetVMConfig200Response`

NewGetVMConfig200ResponseWithDefaults instantiates a new GetVMConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetVMConfig200Response) GetData() GetVMConfig200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVMConfig200Response) GetDataOk() (*GetVMConfig200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVMConfig200Response) SetData(v GetVMConfig200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetVMConfig200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *GetVMConfig200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetVMConfig200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetVMConfig200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetVMConfig200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


