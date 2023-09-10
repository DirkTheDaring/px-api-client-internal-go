# GetNodeTaskLog200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetNodeTaskLog200ResponseDataInner**](GetNodeTaskLog200ResponseDataInner.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetNodeTaskLog200Response

`func NewGetNodeTaskLog200Response() *GetNodeTaskLog200Response`

NewGetNodeTaskLog200Response instantiates a new GetNodeTaskLog200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodeTaskLog200ResponseWithDefaults

`func NewGetNodeTaskLog200ResponseWithDefaults() *GetNodeTaskLog200Response`

NewGetNodeTaskLog200ResponseWithDefaults instantiates a new GetNodeTaskLog200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetNodeTaskLog200Response) GetData() []GetNodeTaskLog200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetNodeTaskLog200Response) GetDataOk() (*[]GetNodeTaskLog200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetNodeTaskLog200Response) SetData(v []GetNodeTaskLog200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetNodeTaskLog200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *GetNodeTaskLog200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetNodeTaskLog200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetNodeTaskLog200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetNodeTaskLog200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


