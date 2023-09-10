# GetContainerConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetContainerConfig200ResponseData**](GetContainerConfig200ResponseData.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetContainerConfig200Response

`func NewGetContainerConfig200Response() *GetContainerConfig200Response`

NewGetContainerConfig200Response instantiates a new GetContainerConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerConfig200ResponseWithDefaults

`func NewGetContainerConfig200ResponseWithDefaults() *GetContainerConfig200Response`

NewGetContainerConfig200ResponseWithDefaults instantiates a new GetContainerConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetContainerConfig200Response) GetData() GetContainerConfig200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContainerConfig200Response) GetDataOk() (*GetContainerConfig200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContainerConfig200Response) SetData(v GetContainerConfig200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetContainerConfig200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *GetContainerConfig200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetContainerConfig200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetContainerConfig200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetContainerConfig200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


