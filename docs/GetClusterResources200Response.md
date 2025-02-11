# GetClusterResources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to [**[]GetClusterResources200ResponseDataInner**](GetClusterResources200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetClusterResources200Response

`func NewGetClusterResources200Response() *GetClusterResources200Response`

NewGetClusterResources200Response instantiates a new GetClusterResources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterResources200ResponseWithDefaults

`func NewGetClusterResources200ResponseWithDefaults() *GetClusterResources200Response`

NewGetClusterResources200ResponseWithDefaults instantiates a new GetClusterResources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GetClusterResources200Response) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetClusterResources200Response) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetClusterResources200Response) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetClusterResources200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetData

`func (o *GetClusterResources200Response) GetData() []GetClusterResources200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetClusterResources200Response) GetDataOk() (*[]GetClusterResources200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetClusterResources200Response) SetData(v []GetClusterResources200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetClusterResources200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


