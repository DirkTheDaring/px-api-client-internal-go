# GetVMConfigPending200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**[]GetVMConfigPending200ResponseDataInner**](GetVMConfigPending200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetVMConfigPending200Response

`func NewGetVMConfigPending200Response() *GetVMConfigPending200Response`

NewGetVMConfigPending200Response instantiates a new GetVMConfigPending200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfigPending200ResponseWithDefaults

`func NewGetVMConfigPending200ResponseWithDefaults() *GetVMConfigPending200Response`

NewGetVMConfigPending200ResponseWithDefaults instantiates a new GetVMConfigPending200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GetVMConfigPending200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetVMConfigPending200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetVMConfigPending200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetVMConfigPending200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *GetVMConfigPending200Response) GetData() []GetVMConfigPending200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVMConfigPending200Response) GetDataOk() (*[]GetVMConfigPending200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVMConfigPending200Response) SetData(v []GetVMConfigPending200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetVMConfigPending200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


