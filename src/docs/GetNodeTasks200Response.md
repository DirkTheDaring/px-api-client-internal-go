# GetNodeTasks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetNodeTasks200ResponseDataInner**](GetNodeTasks200ResponseDataInner.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetNodeTasks200Response

`func NewGetNodeTasks200Response() *GetNodeTasks200Response`

NewGetNodeTasks200Response instantiates a new GetNodeTasks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodeTasks200ResponseWithDefaults

`func NewGetNodeTasks200ResponseWithDefaults() *GetNodeTasks200Response`

NewGetNodeTasks200ResponseWithDefaults instantiates a new GetNodeTasks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetNodeTasks200Response) GetData() []GetNodeTasks200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetNodeTasks200Response) GetDataOk() (*[]GetNodeTasks200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetNodeTasks200Response) SetData(v []GetNodeTasks200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetNodeTasks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *GetNodeTasks200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetNodeTasks200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetNodeTasks200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetNodeTasks200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


