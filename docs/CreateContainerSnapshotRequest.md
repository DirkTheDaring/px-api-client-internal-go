# CreateContainerSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Snapname** | **string** |  | 

## Methods

### NewCreateContainerSnapshotRequest

`func NewCreateContainerSnapshotRequest(snapname string, ) *CreateContainerSnapshotRequest`

NewCreateContainerSnapshotRequest instantiates a new CreateContainerSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerSnapshotRequestWithDefaults

`func NewCreateContainerSnapshotRequestWithDefaults() *CreateContainerSnapshotRequest`

NewCreateContainerSnapshotRequestWithDefaults instantiates a new CreateContainerSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateContainerSnapshotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateContainerSnapshotRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateContainerSnapshotRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateContainerSnapshotRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSnapname

`func (o *CreateContainerSnapshotRequest) GetSnapname() string`

GetSnapname returns the Snapname field if non-nil, zero value otherwise.

### GetSnapnameOk

`func (o *CreateContainerSnapshotRequest) GetSnapnameOk() (*string, bool)`

GetSnapnameOk returns a tuple with the Snapname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapname

`func (o *CreateContainerSnapshotRequest) SetSnapname(v string)`

SetSnapname sets Snapname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


