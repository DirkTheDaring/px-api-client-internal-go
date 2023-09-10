# GetStorages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetStorages200ResponseDataInner**](GetStorages200ResponseDataInner.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetStorages200Response

`func NewGetStorages200Response() *GetStorages200Response`

NewGetStorages200Response instantiates a new GetStorages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorages200ResponseWithDefaults

`func NewGetStorages200ResponseWithDefaults() *GetStorages200Response`

NewGetStorages200ResponseWithDefaults instantiates a new GetStorages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetStorages200Response) GetData() []GetStorages200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetStorages200Response) GetDataOk() (*[]GetStorages200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetStorages200Response) SetData(v []GetStorages200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetStorages200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *GetStorages200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetStorages200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetStorages200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetStorages200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


