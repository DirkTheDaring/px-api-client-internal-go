# CreateVMSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Snapname** | **string** |  | 
**Vmstate** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateVMSnapshotRequest

`func NewCreateVMSnapshotRequest(snapname string, ) *CreateVMSnapshotRequest`

NewCreateVMSnapshotRequest instantiates a new CreateVMSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMSnapshotRequestWithDefaults

`func NewCreateVMSnapshotRequestWithDefaults() *CreateVMSnapshotRequest`

NewCreateVMSnapshotRequestWithDefaults instantiates a new CreateVMSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateVMSnapshotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVMSnapshotRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVMSnapshotRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVMSnapshotRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSnapname

`func (o *CreateVMSnapshotRequest) GetSnapname() string`

GetSnapname returns the Snapname field if non-nil, zero value otherwise.

### GetSnapnameOk

`func (o *CreateVMSnapshotRequest) GetSnapnameOk() (*string, bool)`

GetSnapnameOk returns a tuple with the Snapname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapname

`func (o *CreateVMSnapshotRequest) SetSnapname(v string)`

SetSnapname sets Snapname field to given value.


### GetVmstate

`func (o *CreateVMSnapshotRequest) GetVmstate() bool`

GetVmstate returns the Vmstate field if non-nil, zero value otherwise.

### GetVmstateOk

`func (o *CreateVMSnapshotRequest) GetVmstateOk() (*bool, bool)`

GetVmstateOk returns a tuple with the Vmstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstate

`func (o *CreateVMSnapshotRequest) SetVmstate(v bool)`

SetVmstate sets Vmstate field to given value.

### HasVmstate

`func (o *CreateVMSnapshotRequest) HasVmstate() bool`

HasVmstate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


