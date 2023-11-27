# CreateVMRequestEfidisk0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Efitype** | Pointer to **string** | Size and type of the OVMF EFI vars. &#39;4m&#39; is newer and recommended, and required for Secure Boot. For backwards compatibility, &#39;2m&#39; is used if not otherwise specified. Ignored for VMs with arch&#x3D;aarch64 (ARM). | [optional] 
**File** | Pointer to **string** | The drive&#39;s backing volume. | [optional] 
**Format** | Pointer to **string** | The drive&#39;s backing file&#39;s data format. | [optional] 
**ImportFrom** | Pointer to **string** | Create a new disk, importing from this source (volume ID or absolute path). When an absolute path is specified, it&#39;s up to you to ensure that the source is not actively used by another process during the import! | [optional] 
**PreEnrolledKeys** | Pointer to **bool** | Use am EFI vars template with distribution-specific and Microsoft Standard keys enrolled, if used with &#39;efitype&#x3D;4m&#39;. Note that this will enable Secure Boot by default, though it can still be turned off from within the VM. | [optional] 
**Size** | Pointer to **string** | Disk size. This is purely informational and has no effect. | [optional] 
**Volume** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVMRequestEfidisk0

`func NewCreateVMRequestEfidisk0() *CreateVMRequestEfidisk0`

NewCreateVMRequestEfidisk0 instantiates a new CreateVMRequestEfidisk0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestEfidisk0WithDefaults

`func NewCreateVMRequestEfidisk0WithDefaults() *CreateVMRequestEfidisk0`

NewCreateVMRequestEfidisk0WithDefaults instantiates a new CreateVMRequestEfidisk0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEfitype

`func (o *CreateVMRequestEfidisk0) GetEfitype() string`

GetEfitype returns the Efitype field if non-nil, zero value otherwise.

### GetEfitypeOk

`func (o *CreateVMRequestEfidisk0) GetEfitypeOk() (*string, bool)`

GetEfitypeOk returns a tuple with the Efitype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfitype

`func (o *CreateVMRequestEfidisk0) SetEfitype(v string)`

SetEfitype sets Efitype field to given value.

### HasEfitype

`func (o *CreateVMRequestEfidisk0) HasEfitype() bool`

HasEfitype returns a boolean if a field has been set.

### GetFile

`func (o *CreateVMRequestEfidisk0) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateVMRequestEfidisk0) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateVMRequestEfidisk0) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *CreateVMRequestEfidisk0) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFormat

`func (o *CreateVMRequestEfidisk0) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateVMRequestEfidisk0) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateVMRequestEfidisk0) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateVMRequestEfidisk0) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetImportFrom

`func (o *CreateVMRequestEfidisk0) GetImportFrom() string`

GetImportFrom returns the ImportFrom field if non-nil, zero value otherwise.

### GetImportFromOk

`func (o *CreateVMRequestEfidisk0) GetImportFromOk() (*string, bool)`

GetImportFromOk returns a tuple with the ImportFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFrom

`func (o *CreateVMRequestEfidisk0) SetImportFrom(v string)`

SetImportFrom sets ImportFrom field to given value.

### HasImportFrom

`func (o *CreateVMRequestEfidisk0) HasImportFrom() bool`

HasImportFrom returns a boolean if a field has been set.

### GetPreEnrolledKeys

`func (o *CreateVMRequestEfidisk0) GetPreEnrolledKeys() bool`

GetPreEnrolledKeys returns the PreEnrolledKeys field if non-nil, zero value otherwise.

### GetPreEnrolledKeysOk

`func (o *CreateVMRequestEfidisk0) GetPreEnrolledKeysOk() (*bool, bool)`

GetPreEnrolledKeysOk returns a tuple with the PreEnrolledKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreEnrolledKeys

`func (o *CreateVMRequestEfidisk0) SetPreEnrolledKeys(v bool)`

SetPreEnrolledKeys sets PreEnrolledKeys field to given value.

### HasPreEnrolledKeys

`func (o *CreateVMRequestEfidisk0) HasPreEnrolledKeys() bool`

HasPreEnrolledKeys returns a boolean if a field has been set.

### GetSize

`func (o *CreateVMRequestEfidisk0) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVMRequestEfidisk0) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVMRequestEfidisk0) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateVMRequestEfidisk0) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVolume

`func (o *CreateVMRequestEfidisk0) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateVMRequestEfidisk0) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateVMRequestEfidisk0) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CreateVMRequestEfidisk0) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


