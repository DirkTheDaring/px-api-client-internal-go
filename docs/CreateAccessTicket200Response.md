# CreateAccessTicket200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**CreateAccessTicket200ResponseData**](CreateAccessTicket200ResponseData.md) |  | [optional] 

## Methods

### NewCreateAccessTicket200Response

`func NewCreateAccessTicket200Response() *CreateAccessTicket200Response`

NewCreateAccessTicket200Response instantiates a new CreateAccessTicket200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessTicket200ResponseWithDefaults

`func NewCreateAccessTicket200ResponseWithDefaults() *CreateAccessTicket200Response`

NewCreateAccessTicket200ResponseWithDefaults instantiates a new CreateAccessTicket200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *CreateAccessTicket200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateAccessTicket200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateAccessTicket200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateAccessTicket200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *CreateAccessTicket200Response) GetData() CreateAccessTicket200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateAccessTicket200Response) GetDataOk() (*CreateAccessTicket200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateAccessTicket200Response) SetData(v CreateAccessTicket200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateAccessTicket200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


