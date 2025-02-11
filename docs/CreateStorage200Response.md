# CreateStorage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**CreateStorage200ResponseData**](CreateStorage200ResponseData.md) |  | [optional] 

## Methods

### NewCreateStorage200Response

`func NewCreateStorage200Response() *CreateStorage200Response`

NewCreateStorage200Response instantiates a new CreateStorage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorage200ResponseWithDefaults

`func NewCreateStorage200ResponseWithDefaults() *CreateStorage200Response`

NewCreateStorage200ResponseWithDefaults instantiates a new CreateStorage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *CreateStorage200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateStorage200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateStorage200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateStorage200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *CreateStorage200Response) GetData() CreateStorage200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateStorage200Response) GetDataOk() (*CreateStorage200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateStorage200Response) SetData(v CreateStorage200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateStorage200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


