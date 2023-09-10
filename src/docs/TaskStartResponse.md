# TaskStartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | the task ID. | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTaskStartResponse

`func NewTaskStartResponse() *TaskStartResponse`

NewTaskStartResponse instantiates a new TaskStartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStartResponseWithDefaults

`func NewTaskStartResponseWithDefaults() *TaskStartResponse`

NewTaskStartResponseWithDefaults instantiates a new TaskStartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TaskStartResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaskStartResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaskStartResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TaskStartResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *TaskStartResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TaskStartResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TaskStartResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TaskStartResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


