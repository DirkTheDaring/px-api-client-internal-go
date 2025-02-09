# GetStorageContent200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctime** | Pointer to **int64** | Creation time (seconds since the UNIX Epoch). | [optional] 
**Encrypted** | Pointer to **string** | If whole backup is encrypted, value is the fingerprint or &#39;1&#39;  if encrypted. Only useful for the Proxmox Backup Server storage type. | [optional] 
**Format** | Pointer to **string** | Format identifier (&#39;raw&#39;, &#39;qcow2&#39;, &#39;subvol&#39;, &#39;iso&#39;, &#39;tgz&#39; ...) | [optional] 
**Notes** | Pointer to **string** | Optional notes. If they contain multiple lines, only the first one is returned here. | [optional] 
**Parent** | Pointer to **string** | Volume identifier of parent (for linked cloned). | [optional] 
**Protected** | Pointer to **int32** | Protection status. Currently only supported for backups. | [optional] 
**Size** | Pointer to **int64** | Volume size in bytes. | [optional] 
**Used** | Pointer to **int64** | Used space. Please note that most storage plugins do not report anything useful here. | [optional] 
**Verification** | Pointer to [**GetStorageContent200ResponseDataInnerVerification**](GetStorageContent200ResponseDataInnerVerification.md) |  | [optional] 
**Vmid** | Pointer to **int64** | Associated Owner VMID. | [optional] 
**Volid** | Pointer to **string** | Volume identifier. | [optional] 

## Methods

### NewGetStorageContent200ResponseDataInner

`func NewGetStorageContent200ResponseDataInner() *GetStorageContent200ResponseDataInner`

NewGetStorageContent200ResponseDataInner instantiates a new GetStorageContent200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageContent200ResponseDataInnerWithDefaults

`func NewGetStorageContent200ResponseDataInnerWithDefaults() *GetStorageContent200ResponseDataInner`

NewGetStorageContent200ResponseDataInnerWithDefaults instantiates a new GetStorageContent200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtime

`func (o *GetStorageContent200ResponseDataInner) GetCtime() int64`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *GetStorageContent200ResponseDataInner) GetCtimeOk() (*int64, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *GetStorageContent200ResponseDataInner) SetCtime(v int64)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *GetStorageContent200ResponseDataInner) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetEncrypted

`func (o *GetStorageContent200ResponseDataInner) GetEncrypted() string`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *GetStorageContent200ResponseDataInner) GetEncryptedOk() (*string, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *GetStorageContent200ResponseDataInner) SetEncrypted(v string)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *GetStorageContent200ResponseDataInner) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetFormat

`func (o *GetStorageContent200ResponseDataInner) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GetStorageContent200ResponseDataInner) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GetStorageContent200ResponseDataInner) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GetStorageContent200ResponseDataInner) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetNotes

`func (o *GetStorageContent200ResponseDataInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetStorageContent200ResponseDataInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetStorageContent200ResponseDataInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetStorageContent200ResponseDataInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetParent

`func (o *GetStorageContent200ResponseDataInner) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetStorageContent200ResponseDataInner) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetStorageContent200ResponseDataInner) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetStorageContent200ResponseDataInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtected

`func (o *GetStorageContent200ResponseDataInner) GetProtected() int32`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *GetStorageContent200ResponseDataInner) GetProtectedOk() (*int32, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *GetStorageContent200ResponseDataInner) SetProtected(v int32)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *GetStorageContent200ResponseDataInner) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetSize

`func (o *GetStorageContent200ResponseDataInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetStorageContent200ResponseDataInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetStorageContent200ResponseDataInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetStorageContent200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUsed

`func (o *GetStorageContent200ResponseDataInner) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *GetStorageContent200ResponseDataInner) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *GetStorageContent200ResponseDataInner) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *GetStorageContent200ResponseDataInner) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetVerification

`func (o *GetStorageContent200ResponseDataInner) GetVerification() GetStorageContent200ResponseDataInnerVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *GetStorageContent200ResponseDataInner) GetVerificationOk() (*GetStorageContent200ResponseDataInnerVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *GetStorageContent200ResponseDataInner) SetVerification(v GetStorageContent200ResponseDataInnerVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *GetStorageContent200ResponseDataInner) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetVmid

`func (o *GetStorageContent200ResponseDataInner) GetVmid() int64`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *GetStorageContent200ResponseDataInner) GetVmidOk() (*int64, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *GetStorageContent200ResponseDataInner) SetVmid(v int64)`

SetVmid sets Vmid field to given value.

### HasVmid

`func (o *GetStorageContent200ResponseDataInner) HasVmid() bool`

HasVmid returns a boolean if a field has been set.

### GetVolid

`func (o *GetStorageContent200ResponseDataInner) GetVolid() string`

GetVolid returns the Volid field if non-nil, zero value otherwise.

### GetVolidOk

`func (o *GetStorageContent200ResponseDataInner) GetVolidOk() (*string, bool)`

GetVolidOk returns a tuple with the Volid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolid

`func (o *GetStorageContent200ResponseDataInner) SetVolid(v string)`

SetVolid sets Volid field to given value.

### HasVolid

`func (o *GetStorageContent200ResponseDataInner) HasVolid() bool`

HasVolid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


