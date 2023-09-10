# CreateVMRequestTpmstate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to **string** | The drive&#39;s backing volume. | [optional] 
**ImportFrom** | Pointer to **string** | Create a new disk, importing from this source (volume ID or absolute path). When an absolute path is specified, it&#39;s up to you to ensure that the source is not actively used by another process during the import! | [optional] 
**Size** | Pointer to **string** | Disk size. This is purely informational and has no effect. | [optional] 
**Version** | Pointer to **string** | The TPM interface version. v2.0 is newer and should be preferred. Note that this cannot be changed later on. | [optional] 
**Volume** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVMRequestTpmstate0

`func NewCreateVMRequestTpmstate0() *CreateVMRequestTpmstate0`

NewCreateVMRequestTpmstate0 instantiates a new CreateVMRequestTpmstate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestTpmstate0WithDefaults

`func NewCreateVMRequestTpmstate0WithDefaults() *CreateVMRequestTpmstate0`

NewCreateVMRequestTpmstate0WithDefaults instantiates a new CreateVMRequestTpmstate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *CreateVMRequestTpmstate0) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateVMRequestTpmstate0) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateVMRequestTpmstate0) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *CreateVMRequestTpmstate0) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetImportFrom

`func (o *CreateVMRequestTpmstate0) GetImportFrom() string`

GetImportFrom returns the ImportFrom field if non-nil, zero value otherwise.

### GetImportFromOk

`func (o *CreateVMRequestTpmstate0) GetImportFromOk() (*string, bool)`

GetImportFromOk returns a tuple with the ImportFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFrom

`func (o *CreateVMRequestTpmstate0) SetImportFrom(v string)`

SetImportFrom sets ImportFrom field to given value.

### HasImportFrom

`func (o *CreateVMRequestTpmstate0) HasImportFrom() bool`

HasImportFrom returns a boolean if a field has been set.

### GetSize

`func (o *CreateVMRequestTpmstate0) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVMRequestTpmstate0) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVMRequestTpmstate0) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateVMRequestTpmstate0) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVersion

`func (o *CreateVMRequestTpmstate0) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateVMRequestTpmstate0) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateVMRequestTpmstate0) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateVMRequestTpmstate0) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVolume

`func (o *CreateVMRequestTpmstate0) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateVMRequestTpmstate0) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateVMRequestTpmstate0) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CreateVMRequestTpmstate0) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


