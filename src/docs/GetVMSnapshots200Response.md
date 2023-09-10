# GetVMSnapshots200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetVMSnapshots200ResponseDataInner**](GetVMSnapshots200ResponseDataInner.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetVMSnapshots200Response

`func NewGetVMSnapshots200Response() *GetVMSnapshots200Response`

NewGetVMSnapshots200Response instantiates a new GetVMSnapshots200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMSnapshots200ResponseWithDefaults

`func NewGetVMSnapshots200ResponseWithDefaults() *GetVMSnapshots200Response`

NewGetVMSnapshots200ResponseWithDefaults instantiates a new GetVMSnapshots200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetVMSnapshots200Response) GetData() []GetVMSnapshots200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVMSnapshots200Response) GetDataOk() (*[]GetVMSnapshots200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVMSnapshots200Response) SetData(v []GetVMSnapshots200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetVMSnapshots200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *GetVMSnapshots200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetVMSnapshots200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetVMSnapshots200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetVMSnapshots200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


