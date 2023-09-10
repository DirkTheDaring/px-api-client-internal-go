# CreateVM200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateVM200Response

`func NewCreateVM200Response() *CreateVM200Response`

NewCreateVM200Response instantiates a new CreateVM200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVM200ResponseWithDefaults

`func NewCreateVM200ResponseWithDefaults() *CreateVM200Response`

NewCreateVM200ResponseWithDefaults instantiates a new CreateVM200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateVM200Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateVM200Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateVM200Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CreateVM200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *CreateVM200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateVM200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateVM200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateVM200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


