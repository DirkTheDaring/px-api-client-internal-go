# GetStorageContent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**[]GetStorageContent200ResponseDataInner**](GetStorageContent200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetStorageContent200Response

`func NewGetStorageContent200Response() *GetStorageContent200Response`

NewGetStorageContent200Response instantiates a new GetStorageContent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageContent200ResponseWithDefaults

`func NewGetStorageContent200ResponseWithDefaults() *GetStorageContent200Response`

NewGetStorageContent200ResponseWithDefaults instantiates a new GetStorageContent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GetStorageContent200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetStorageContent200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetStorageContent200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetStorageContent200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *GetStorageContent200Response) GetData() []GetStorageContent200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetStorageContent200Response) GetDataOk() (*[]GetStorageContent200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetStorageContent200Response) SetData(v []GetStorageContent200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetStorageContent200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


