# GetContainerSnapshots200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**[]GetContainerSnapshots200ResponseDataInner**](GetContainerSnapshots200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetContainerSnapshots200Response

`func NewGetContainerSnapshots200Response() *GetContainerSnapshots200Response`

NewGetContainerSnapshots200Response instantiates a new GetContainerSnapshots200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerSnapshots200ResponseWithDefaults

`func NewGetContainerSnapshots200ResponseWithDefaults() *GetContainerSnapshots200Response`

NewGetContainerSnapshots200ResponseWithDefaults instantiates a new GetContainerSnapshots200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GetContainerSnapshots200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetContainerSnapshots200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetContainerSnapshots200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetContainerSnapshots200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *GetContainerSnapshots200Response) GetData() []GetContainerSnapshots200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContainerSnapshots200Response) GetDataOk() (*[]GetContainerSnapshots200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContainerSnapshots200Response) SetData(v []GetContainerSnapshots200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetContainerSnapshots200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


