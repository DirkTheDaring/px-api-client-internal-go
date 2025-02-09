# CreateVMRequestVirtio0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aio** | Pointer to **string** |  | [optional] 
**Backup** | Pointer to **bool** |  | [optional] 
**Bps** | Pointer to **int64** |  | [optional] 
**BpsMaxLength** | Pointer to **int64** |  | [optional] 
**BpsRd** | Pointer to **int64** |  | [optional] 
**BpsRdMaxLength** | Pointer to **int64** |  | [optional] 
**BpsWr** | Pointer to **int64** |  | [optional] 
**BpsWrMaxLength** | Pointer to **int64** |  | [optional] 
**Cache** | Pointer to **string** |  | [optional] 
**Cyls** | Pointer to **int64** |  | [optional] 
**DetectZeroes** | Pointer to **bool** |  | [optional] 
**Discard** | Pointer to **string** |  | [optional] 
**File** | **string** |  | 
**Format** | Pointer to **string** |  | [optional] 
**Heads** | Pointer to **int64** |  | [optional] 
**ImportFrom** | Pointer to **string** |  | [optional] 
**Iops** | Pointer to **int64** |  | [optional] 
**IopsMax** | Pointer to **int64** |  | [optional] 
**IopsMaxLength** | Pointer to **int64** |  | [optional] 
**IopsRd** | Pointer to **int64** |  | [optional] 
**IopsRdMax** | Pointer to **int64** |  | [optional] 
**IopsRdMaxLength** | Pointer to **int64** |  | [optional] 
**IopsWr** | Pointer to **int64** |  | [optional] 
**IopsWrMax** | Pointer to **int64** |  | [optional] 
**IopsWrMaxLength** | Pointer to **int64** |  | [optional] 
**Iothread** | Pointer to **bool** |  | [optional] 
**Mbps** | Pointer to **float32** |  | [optional] 
**MbpsMax** | Pointer to **float32** |  | [optional] 
**MbpsRd** | Pointer to **float32** |  | [optional] 
**MbpsRdMax** | Pointer to **float32** |  | [optional] 
**MbpsWr** | Pointer to **float32** |  | [optional] 
**MbpsWrMax** | Pointer to **float32** |  | [optional] 
**Media** | Pointer to **string** |  | [optional] 
**Replicate** | Pointer to **bool** |  | [optional] 
**Rerror** | Pointer to **string** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Secs** | Pointer to **int64** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Snapshot** | Pointer to **bool** |  | [optional] 
**Trans** | Pointer to **string** |  | [optional] 
**Werror** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVMRequestVirtio0

`func NewCreateVMRequestVirtio0(file string, ) *CreateVMRequestVirtio0`

NewCreateVMRequestVirtio0 instantiates a new CreateVMRequestVirtio0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestVirtio0WithDefaults

`func NewCreateVMRequestVirtio0WithDefaults() *CreateVMRequestVirtio0`

NewCreateVMRequestVirtio0WithDefaults instantiates a new CreateVMRequestVirtio0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAio

`func (o *CreateVMRequestVirtio0) GetAio() string`

GetAio returns the Aio field if non-nil, zero value otherwise.

### GetAioOk

`func (o *CreateVMRequestVirtio0) GetAioOk() (*string, bool)`

GetAioOk returns a tuple with the Aio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAio

`func (o *CreateVMRequestVirtio0) SetAio(v string)`

SetAio sets Aio field to given value.

### HasAio

`func (o *CreateVMRequestVirtio0) HasAio() bool`

HasAio returns a boolean if a field has been set.

### GetBackup

`func (o *CreateVMRequestVirtio0) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateVMRequestVirtio0) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateVMRequestVirtio0) SetBackup(v bool)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *CreateVMRequestVirtio0) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBps

`func (o *CreateVMRequestVirtio0) GetBps() int64`

GetBps returns the Bps field if non-nil, zero value otherwise.

### GetBpsOk

`func (o *CreateVMRequestVirtio0) GetBpsOk() (*int64, bool)`

GetBpsOk returns a tuple with the Bps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBps

`func (o *CreateVMRequestVirtio0) SetBps(v int64)`

SetBps sets Bps field to given value.

### HasBps

`func (o *CreateVMRequestVirtio0) HasBps() bool`

HasBps returns a boolean if a field has been set.

### GetBpsMaxLength

`func (o *CreateVMRequestVirtio0) GetBpsMaxLength() int64`

GetBpsMaxLength returns the BpsMaxLength field if non-nil, zero value otherwise.

### GetBpsMaxLengthOk

`func (o *CreateVMRequestVirtio0) GetBpsMaxLengthOk() (*int64, bool)`

GetBpsMaxLengthOk returns a tuple with the BpsMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsMaxLength

`func (o *CreateVMRequestVirtio0) SetBpsMaxLength(v int64)`

SetBpsMaxLength sets BpsMaxLength field to given value.

### HasBpsMaxLength

`func (o *CreateVMRequestVirtio0) HasBpsMaxLength() bool`

HasBpsMaxLength returns a boolean if a field has been set.

### GetBpsRd

`func (o *CreateVMRequestVirtio0) GetBpsRd() int64`

GetBpsRd returns the BpsRd field if non-nil, zero value otherwise.

### GetBpsRdOk

`func (o *CreateVMRequestVirtio0) GetBpsRdOk() (*int64, bool)`

GetBpsRdOk returns a tuple with the BpsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsRd

`func (o *CreateVMRequestVirtio0) SetBpsRd(v int64)`

SetBpsRd sets BpsRd field to given value.

### HasBpsRd

`func (o *CreateVMRequestVirtio0) HasBpsRd() bool`

HasBpsRd returns a boolean if a field has been set.

### GetBpsRdMaxLength

`func (o *CreateVMRequestVirtio0) GetBpsRdMaxLength() int64`

GetBpsRdMaxLength returns the BpsRdMaxLength field if non-nil, zero value otherwise.

### GetBpsRdMaxLengthOk

`func (o *CreateVMRequestVirtio0) GetBpsRdMaxLengthOk() (*int64, bool)`

GetBpsRdMaxLengthOk returns a tuple with the BpsRdMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsRdMaxLength

`func (o *CreateVMRequestVirtio0) SetBpsRdMaxLength(v int64)`

SetBpsRdMaxLength sets BpsRdMaxLength field to given value.

### HasBpsRdMaxLength

`func (o *CreateVMRequestVirtio0) HasBpsRdMaxLength() bool`

HasBpsRdMaxLength returns a boolean if a field has been set.

### GetBpsWr

`func (o *CreateVMRequestVirtio0) GetBpsWr() int64`

GetBpsWr returns the BpsWr field if non-nil, zero value otherwise.

### GetBpsWrOk

`func (o *CreateVMRequestVirtio0) GetBpsWrOk() (*int64, bool)`

GetBpsWrOk returns a tuple with the BpsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsWr

`func (o *CreateVMRequestVirtio0) SetBpsWr(v int64)`

SetBpsWr sets BpsWr field to given value.

### HasBpsWr

`func (o *CreateVMRequestVirtio0) HasBpsWr() bool`

HasBpsWr returns a boolean if a field has been set.

### GetBpsWrMaxLength

`func (o *CreateVMRequestVirtio0) GetBpsWrMaxLength() int64`

GetBpsWrMaxLength returns the BpsWrMaxLength field if non-nil, zero value otherwise.

### GetBpsWrMaxLengthOk

`func (o *CreateVMRequestVirtio0) GetBpsWrMaxLengthOk() (*int64, bool)`

GetBpsWrMaxLengthOk returns a tuple with the BpsWrMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsWrMaxLength

`func (o *CreateVMRequestVirtio0) SetBpsWrMaxLength(v int64)`

SetBpsWrMaxLength sets BpsWrMaxLength field to given value.

### HasBpsWrMaxLength

`func (o *CreateVMRequestVirtio0) HasBpsWrMaxLength() bool`

HasBpsWrMaxLength returns a boolean if a field has been set.

### GetCache

`func (o *CreateVMRequestVirtio0) GetCache() string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *CreateVMRequestVirtio0) GetCacheOk() (*string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *CreateVMRequestVirtio0) SetCache(v string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *CreateVMRequestVirtio0) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCyls

`func (o *CreateVMRequestVirtio0) GetCyls() int64`

GetCyls returns the Cyls field if non-nil, zero value otherwise.

### GetCylsOk

`func (o *CreateVMRequestVirtio0) GetCylsOk() (*int64, bool)`

GetCylsOk returns a tuple with the Cyls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyls

`func (o *CreateVMRequestVirtio0) SetCyls(v int64)`

SetCyls sets Cyls field to given value.

### HasCyls

`func (o *CreateVMRequestVirtio0) HasCyls() bool`

HasCyls returns a boolean if a field has been set.

### GetDetectZeroes

`func (o *CreateVMRequestVirtio0) GetDetectZeroes() bool`

GetDetectZeroes returns the DetectZeroes field if non-nil, zero value otherwise.

### GetDetectZeroesOk

`func (o *CreateVMRequestVirtio0) GetDetectZeroesOk() (*bool, bool)`

GetDetectZeroesOk returns a tuple with the DetectZeroes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectZeroes

`func (o *CreateVMRequestVirtio0) SetDetectZeroes(v bool)`

SetDetectZeroes sets DetectZeroes field to given value.

### HasDetectZeroes

`func (o *CreateVMRequestVirtio0) HasDetectZeroes() bool`

HasDetectZeroes returns a boolean if a field has been set.

### GetDiscard

`func (o *CreateVMRequestVirtio0) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *CreateVMRequestVirtio0) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *CreateVMRequestVirtio0) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *CreateVMRequestVirtio0) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetFile

`func (o *CreateVMRequestVirtio0) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateVMRequestVirtio0) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateVMRequestVirtio0) SetFile(v string)`

SetFile sets File field to given value.


### GetFormat

`func (o *CreateVMRequestVirtio0) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateVMRequestVirtio0) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateVMRequestVirtio0) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateVMRequestVirtio0) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHeads

`func (o *CreateVMRequestVirtio0) GetHeads() int64`

GetHeads returns the Heads field if non-nil, zero value otherwise.

### GetHeadsOk

`func (o *CreateVMRequestVirtio0) GetHeadsOk() (*int64, bool)`

GetHeadsOk returns a tuple with the Heads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeads

`func (o *CreateVMRequestVirtio0) SetHeads(v int64)`

SetHeads sets Heads field to given value.

### HasHeads

`func (o *CreateVMRequestVirtio0) HasHeads() bool`

HasHeads returns a boolean if a field has been set.

### GetImportFrom

`func (o *CreateVMRequestVirtio0) GetImportFrom() string`

GetImportFrom returns the ImportFrom field if non-nil, zero value otherwise.

### GetImportFromOk

`func (o *CreateVMRequestVirtio0) GetImportFromOk() (*string, bool)`

GetImportFromOk returns a tuple with the ImportFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFrom

`func (o *CreateVMRequestVirtio0) SetImportFrom(v string)`

SetImportFrom sets ImportFrom field to given value.

### HasImportFrom

`func (o *CreateVMRequestVirtio0) HasImportFrom() bool`

HasImportFrom returns a boolean if a field has been set.

### GetIops

`func (o *CreateVMRequestVirtio0) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *CreateVMRequestVirtio0) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *CreateVMRequestVirtio0) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *CreateVMRequestVirtio0) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetIopsMax

`func (o *CreateVMRequestVirtio0) GetIopsMax() int64`

GetIopsMax returns the IopsMax field if non-nil, zero value otherwise.

### GetIopsMaxOk

`func (o *CreateVMRequestVirtio0) GetIopsMaxOk() (*int64, bool)`

GetIopsMaxOk returns a tuple with the IopsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMax

`func (o *CreateVMRequestVirtio0) SetIopsMax(v int64)`

SetIopsMax sets IopsMax field to given value.

### HasIopsMax

`func (o *CreateVMRequestVirtio0) HasIopsMax() bool`

HasIopsMax returns a boolean if a field has been set.

### GetIopsMaxLength

`func (o *CreateVMRequestVirtio0) GetIopsMaxLength() int64`

GetIopsMaxLength returns the IopsMaxLength field if non-nil, zero value otherwise.

### GetIopsMaxLengthOk

`func (o *CreateVMRequestVirtio0) GetIopsMaxLengthOk() (*int64, bool)`

GetIopsMaxLengthOk returns a tuple with the IopsMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMaxLength

`func (o *CreateVMRequestVirtio0) SetIopsMaxLength(v int64)`

SetIopsMaxLength sets IopsMaxLength field to given value.

### HasIopsMaxLength

`func (o *CreateVMRequestVirtio0) HasIopsMaxLength() bool`

HasIopsMaxLength returns a boolean if a field has been set.

### GetIopsRd

`func (o *CreateVMRequestVirtio0) GetIopsRd() int64`

GetIopsRd returns the IopsRd field if non-nil, zero value otherwise.

### GetIopsRdOk

`func (o *CreateVMRequestVirtio0) GetIopsRdOk() (*int64, bool)`

GetIopsRdOk returns a tuple with the IopsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRd

`func (o *CreateVMRequestVirtio0) SetIopsRd(v int64)`

SetIopsRd sets IopsRd field to given value.

### HasIopsRd

`func (o *CreateVMRequestVirtio0) HasIopsRd() bool`

HasIopsRd returns a boolean if a field has been set.

### GetIopsRdMax

`func (o *CreateVMRequestVirtio0) GetIopsRdMax() int64`

GetIopsRdMax returns the IopsRdMax field if non-nil, zero value otherwise.

### GetIopsRdMaxOk

`func (o *CreateVMRequestVirtio0) GetIopsRdMaxOk() (*int64, bool)`

GetIopsRdMaxOk returns a tuple with the IopsRdMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRdMax

`func (o *CreateVMRequestVirtio0) SetIopsRdMax(v int64)`

SetIopsRdMax sets IopsRdMax field to given value.

### HasIopsRdMax

`func (o *CreateVMRequestVirtio0) HasIopsRdMax() bool`

HasIopsRdMax returns a boolean if a field has been set.

### GetIopsRdMaxLength

`func (o *CreateVMRequestVirtio0) GetIopsRdMaxLength() int64`

GetIopsRdMaxLength returns the IopsRdMaxLength field if non-nil, zero value otherwise.

### GetIopsRdMaxLengthOk

`func (o *CreateVMRequestVirtio0) GetIopsRdMaxLengthOk() (*int64, bool)`

GetIopsRdMaxLengthOk returns a tuple with the IopsRdMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRdMaxLength

`func (o *CreateVMRequestVirtio0) SetIopsRdMaxLength(v int64)`

SetIopsRdMaxLength sets IopsRdMaxLength field to given value.

### HasIopsRdMaxLength

`func (o *CreateVMRequestVirtio0) HasIopsRdMaxLength() bool`

HasIopsRdMaxLength returns a boolean if a field has been set.

### GetIopsWr

`func (o *CreateVMRequestVirtio0) GetIopsWr() int64`

GetIopsWr returns the IopsWr field if non-nil, zero value otherwise.

### GetIopsWrOk

`func (o *CreateVMRequestVirtio0) GetIopsWrOk() (*int64, bool)`

GetIopsWrOk returns a tuple with the IopsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWr

`func (o *CreateVMRequestVirtio0) SetIopsWr(v int64)`

SetIopsWr sets IopsWr field to given value.

### HasIopsWr

`func (o *CreateVMRequestVirtio0) HasIopsWr() bool`

HasIopsWr returns a boolean if a field has been set.

### GetIopsWrMax

`func (o *CreateVMRequestVirtio0) GetIopsWrMax() int64`

GetIopsWrMax returns the IopsWrMax field if non-nil, zero value otherwise.

### GetIopsWrMaxOk

`func (o *CreateVMRequestVirtio0) GetIopsWrMaxOk() (*int64, bool)`

GetIopsWrMaxOk returns a tuple with the IopsWrMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWrMax

`func (o *CreateVMRequestVirtio0) SetIopsWrMax(v int64)`

SetIopsWrMax sets IopsWrMax field to given value.

### HasIopsWrMax

`func (o *CreateVMRequestVirtio0) HasIopsWrMax() bool`

HasIopsWrMax returns a boolean if a field has been set.

### GetIopsWrMaxLength

`func (o *CreateVMRequestVirtio0) GetIopsWrMaxLength() int64`

GetIopsWrMaxLength returns the IopsWrMaxLength field if non-nil, zero value otherwise.

### GetIopsWrMaxLengthOk

`func (o *CreateVMRequestVirtio0) GetIopsWrMaxLengthOk() (*int64, bool)`

GetIopsWrMaxLengthOk returns a tuple with the IopsWrMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWrMaxLength

`func (o *CreateVMRequestVirtio0) SetIopsWrMaxLength(v int64)`

SetIopsWrMaxLength sets IopsWrMaxLength field to given value.

### HasIopsWrMaxLength

`func (o *CreateVMRequestVirtio0) HasIopsWrMaxLength() bool`

HasIopsWrMaxLength returns a boolean if a field has been set.

### GetIothread

`func (o *CreateVMRequestVirtio0) GetIothread() bool`

GetIothread returns the Iothread field if non-nil, zero value otherwise.

### GetIothreadOk

`func (o *CreateVMRequestVirtio0) GetIothreadOk() (*bool, bool)`

GetIothreadOk returns a tuple with the Iothread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIothread

`func (o *CreateVMRequestVirtio0) SetIothread(v bool)`

SetIothread sets Iothread field to given value.

### HasIothread

`func (o *CreateVMRequestVirtio0) HasIothread() bool`

HasIothread returns a boolean if a field has been set.

### GetMbps

`func (o *CreateVMRequestVirtio0) GetMbps() float32`

GetMbps returns the Mbps field if non-nil, zero value otherwise.

### GetMbpsOk

`func (o *CreateVMRequestVirtio0) GetMbpsOk() (*float32, bool)`

GetMbpsOk returns a tuple with the Mbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbps

`func (o *CreateVMRequestVirtio0) SetMbps(v float32)`

SetMbps sets Mbps field to given value.

### HasMbps

`func (o *CreateVMRequestVirtio0) HasMbps() bool`

HasMbps returns a boolean if a field has been set.

### GetMbpsMax

`func (o *CreateVMRequestVirtio0) GetMbpsMax() float32`

GetMbpsMax returns the MbpsMax field if non-nil, zero value otherwise.

### GetMbpsMaxOk

`func (o *CreateVMRequestVirtio0) GetMbpsMaxOk() (*float32, bool)`

GetMbpsMaxOk returns a tuple with the MbpsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsMax

`func (o *CreateVMRequestVirtio0) SetMbpsMax(v float32)`

SetMbpsMax sets MbpsMax field to given value.

### HasMbpsMax

`func (o *CreateVMRequestVirtio0) HasMbpsMax() bool`

HasMbpsMax returns a boolean if a field has been set.

### GetMbpsRd

`func (o *CreateVMRequestVirtio0) GetMbpsRd() float32`

GetMbpsRd returns the MbpsRd field if non-nil, zero value otherwise.

### GetMbpsRdOk

`func (o *CreateVMRequestVirtio0) GetMbpsRdOk() (*float32, bool)`

GetMbpsRdOk returns a tuple with the MbpsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsRd

`func (o *CreateVMRequestVirtio0) SetMbpsRd(v float32)`

SetMbpsRd sets MbpsRd field to given value.

### HasMbpsRd

`func (o *CreateVMRequestVirtio0) HasMbpsRd() bool`

HasMbpsRd returns a boolean if a field has been set.

### GetMbpsRdMax

`func (o *CreateVMRequestVirtio0) GetMbpsRdMax() float32`

GetMbpsRdMax returns the MbpsRdMax field if non-nil, zero value otherwise.

### GetMbpsRdMaxOk

`func (o *CreateVMRequestVirtio0) GetMbpsRdMaxOk() (*float32, bool)`

GetMbpsRdMaxOk returns a tuple with the MbpsRdMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsRdMax

`func (o *CreateVMRequestVirtio0) SetMbpsRdMax(v float32)`

SetMbpsRdMax sets MbpsRdMax field to given value.

### HasMbpsRdMax

`func (o *CreateVMRequestVirtio0) HasMbpsRdMax() bool`

HasMbpsRdMax returns a boolean if a field has been set.

### GetMbpsWr

`func (o *CreateVMRequestVirtio0) GetMbpsWr() float32`

GetMbpsWr returns the MbpsWr field if non-nil, zero value otherwise.

### GetMbpsWrOk

`func (o *CreateVMRequestVirtio0) GetMbpsWrOk() (*float32, bool)`

GetMbpsWrOk returns a tuple with the MbpsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsWr

`func (o *CreateVMRequestVirtio0) SetMbpsWr(v float32)`

SetMbpsWr sets MbpsWr field to given value.

### HasMbpsWr

`func (o *CreateVMRequestVirtio0) HasMbpsWr() bool`

HasMbpsWr returns a boolean if a field has been set.

### GetMbpsWrMax

`func (o *CreateVMRequestVirtio0) GetMbpsWrMax() float32`

GetMbpsWrMax returns the MbpsWrMax field if non-nil, zero value otherwise.

### GetMbpsWrMaxOk

`func (o *CreateVMRequestVirtio0) GetMbpsWrMaxOk() (*float32, bool)`

GetMbpsWrMaxOk returns a tuple with the MbpsWrMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsWrMax

`func (o *CreateVMRequestVirtio0) SetMbpsWrMax(v float32)`

SetMbpsWrMax sets MbpsWrMax field to given value.

### HasMbpsWrMax

`func (o *CreateVMRequestVirtio0) HasMbpsWrMax() bool`

HasMbpsWrMax returns a boolean if a field has been set.

### GetMedia

`func (o *CreateVMRequestVirtio0) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *CreateVMRequestVirtio0) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *CreateVMRequestVirtio0) SetMedia(v string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *CreateVMRequestVirtio0) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetReplicate

`func (o *CreateVMRequestVirtio0) GetReplicate() bool`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *CreateVMRequestVirtio0) GetReplicateOk() (*bool, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *CreateVMRequestVirtio0) SetReplicate(v bool)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *CreateVMRequestVirtio0) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetRerror

`func (o *CreateVMRequestVirtio0) GetRerror() string`

GetRerror returns the Rerror field if non-nil, zero value otherwise.

### GetRerrorOk

`func (o *CreateVMRequestVirtio0) GetRerrorOk() (*string, bool)`

GetRerrorOk returns a tuple with the Rerror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerror

`func (o *CreateVMRequestVirtio0) SetRerror(v string)`

SetRerror sets Rerror field to given value.

### HasRerror

`func (o *CreateVMRequestVirtio0) HasRerror() bool`

HasRerror returns a boolean if a field has been set.

### GetRo

`func (o *CreateVMRequestVirtio0) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *CreateVMRequestVirtio0) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *CreateVMRequestVirtio0) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *CreateVMRequestVirtio0) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetSecs

`func (o *CreateVMRequestVirtio0) GetSecs() int64`

GetSecs returns the Secs field if non-nil, zero value otherwise.

### GetSecsOk

`func (o *CreateVMRequestVirtio0) GetSecsOk() (*int64, bool)`

GetSecsOk returns a tuple with the Secs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecs

`func (o *CreateVMRequestVirtio0) SetSecs(v int64)`

SetSecs sets Secs field to given value.

### HasSecs

`func (o *CreateVMRequestVirtio0) HasSecs() bool`

HasSecs returns a boolean if a field has been set.

### GetSerial

`func (o *CreateVMRequestVirtio0) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CreateVMRequestVirtio0) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CreateVMRequestVirtio0) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CreateVMRequestVirtio0) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetShared

`func (o *CreateVMRequestVirtio0) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateVMRequestVirtio0) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateVMRequestVirtio0) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateVMRequestVirtio0) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *CreateVMRequestVirtio0) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVMRequestVirtio0) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVMRequestVirtio0) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateVMRequestVirtio0) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapshot

`func (o *CreateVMRequestVirtio0) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CreateVMRequestVirtio0) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CreateVMRequestVirtio0) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CreateVMRequestVirtio0) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetTrans

`func (o *CreateVMRequestVirtio0) GetTrans() string`

GetTrans returns the Trans field if non-nil, zero value otherwise.

### GetTransOk

`func (o *CreateVMRequestVirtio0) GetTransOk() (*string, bool)`

GetTransOk returns a tuple with the Trans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrans

`func (o *CreateVMRequestVirtio0) SetTrans(v string)`

SetTrans sets Trans field to given value.

### HasTrans

`func (o *CreateVMRequestVirtio0) HasTrans() bool`

HasTrans returns a boolean if a field has been set.

### GetWerror

`func (o *CreateVMRequestVirtio0) GetWerror() string`

GetWerror returns the Werror field if non-nil, zero value otherwise.

### GetWerrorOk

`func (o *CreateVMRequestVirtio0) GetWerrorOk() (*string, bool)`

GetWerrorOk returns a tuple with the Werror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWerror

`func (o *CreateVMRequestVirtio0) SetWerror(v string)`

SetWerror sets Werror field to given value.

### HasWerror

`func (o *CreateVMRequestVirtio0) HasWerror() bool`

HasWerror returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


