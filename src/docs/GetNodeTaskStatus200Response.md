# GetNodeTaskStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetNodeTaskStatus200ResponseData**](GetNodeTaskStatus200ResponseData.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetNodeTaskStatus200Response

`func NewGetNodeTaskStatus200Response() *GetNodeTaskStatus200Response`

NewGetNodeTaskStatus200Response instantiates a new GetNodeTaskStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodeTaskStatus200ResponseWithDefaults

`func NewGetNodeTaskStatus200ResponseWithDefaults() *GetNodeTaskStatus200Response`

NewGetNodeTaskStatus200ResponseWithDefaults instantiates a new GetNodeTaskStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetNodeTaskStatus200Response) GetData() GetNodeTaskStatus200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetNodeTaskStatus200Response) GetDataOk() (*GetNodeTaskStatus200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetNodeTaskStatus200Response) SetData(v GetNodeTaskStatus200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetNodeTaskStatus200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *GetNodeTaskStatus200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetNodeTaskStatus200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetNodeTaskStatus200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetNodeTaskStatus200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


