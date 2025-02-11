# GetVM200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**[]GetVM200ResponseDataInner**](GetVM200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetVM200Response

`func NewGetVM200Response() *GetVM200Response`

NewGetVM200Response instantiates a new GetVM200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVM200ResponseWithDefaults

`func NewGetVM200ResponseWithDefaults() *GetVM200Response`

NewGetVM200ResponseWithDefaults instantiates a new GetVM200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GetVM200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetVM200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetVM200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetVM200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *GetVM200Response) GetData() []GetVM200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVM200Response) GetDataOk() (*[]GetVM200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVM200Response) SetData(v []GetVM200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetVM200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


