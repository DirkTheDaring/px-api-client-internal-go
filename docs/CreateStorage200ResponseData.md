# CreateStorage200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**CreateStorage200ResponseDataConfig**](CreateStorage200ResponseDataConfig.md) |  | [optional] 
**Storage** | Pointer to **string** | The ID of the created storage. | [optional] 
**Type** | Pointer to **string** | The type of the created storage. | [optional] 

## Methods

### NewCreateStorage200ResponseData

`func NewCreateStorage200ResponseData() *CreateStorage200ResponseData`

NewCreateStorage200ResponseData instantiates a new CreateStorage200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorage200ResponseDataWithDefaults

`func NewCreateStorage200ResponseDataWithDefaults() *CreateStorage200ResponseData`

NewCreateStorage200ResponseDataWithDefaults instantiates a new CreateStorage200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateStorage200ResponseData) GetConfig() CreateStorage200ResponseDataConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateStorage200ResponseData) GetConfigOk() (*CreateStorage200ResponseDataConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateStorage200ResponseData) SetConfig(v CreateStorage200ResponseDataConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateStorage200ResponseData) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorage

`func (o *CreateStorage200ResponseData) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateStorage200ResponseData) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateStorage200ResponseData) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateStorage200ResponseData) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetType

`func (o *CreateStorage200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStorage200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStorage200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateStorage200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


