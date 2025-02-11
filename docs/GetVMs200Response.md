# GetVMs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**[]GetVMs200ResponseDataInner**](GetVMs200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetVMs200Response

`func NewGetVMs200Response() *GetVMs200Response`

NewGetVMs200Response instantiates a new GetVMs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMs200ResponseWithDefaults

`func NewGetVMs200ResponseWithDefaults() *GetVMs200Response`

NewGetVMs200ResponseWithDefaults instantiates a new GetVMs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GetVMs200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetVMs200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetVMs200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetVMs200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *GetVMs200Response) GetData() []GetVMs200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVMs200Response) GetDataOk() (*[]GetVMs200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVMs200Response) SetData(v []GetVMs200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetVMs200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


