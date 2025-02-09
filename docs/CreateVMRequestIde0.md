# CreateVMRequestIde0

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
**Mbps** | Pointer to **float32** |  | [optional] 
**MbpsMax** | Pointer to **float32** |  | [optional] 
**MbpsRd** | Pointer to **float32** |  | [optional] 
**MbpsRdMax** | Pointer to **float32** |  | [optional] 
**MbpsWr** | Pointer to **float32** |  | [optional] 
**MbpsWrMax** | Pointer to **float32** |  | [optional] 
**Media** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Replicate** | Pointer to **bool** |  | [optional] 
**Rerror** | Pointer to **string** |  | [optional] 
**Secs** | Pointer to **int64** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Snapshot** | Pointer to **bool** |  | [optional] 
**Ssd** | Pointer to **bool** |  | [optional] 
**Trans** | Pointer to **string** |  | [optional] 
**Werror** | Pointer to **string** |  | [optional] 
**Wwn** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVMRequestIde0

`func NewCreateVMRequestIde0(file string, ) *CreateVMRequestIde0`

NewCreateVMRequestIde0 instantiates a new CreateVMRequestIde0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestIde0WithDefaults

`func NewCreateVMRequestIde0WithDefaults() *CreateVMRequestIde0`

NewCreateVMRequestIde0WithDefaults instantiates a new CreateVMRequestIde0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAio

`func (o *CreateVMRequestIde0) GetAio() string`

GetAio returns the Aio field if non-nil, zero value otherwise.

### GetAioOk

`func (o *CreateVMRequestIde0) GetAioOk() (*string, bool)`

GetAioOk returns a tuple with the Aio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAio

`func (o *CreateVMRequestIde0) SetAio(v string)`

SetAio sets Aio field to given value.

### HasAio

`func (o *CreateVMRequestIde0) HasAio() bool`

HasAio returns a boolean if a field has been set.

### GetBackup

`func (o *CreateVMRequestIde0) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateVMRequestIde0) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateVMRequestIde0) SetBackup(v bool)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *CreateVMRequestIde0) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBps

`func (o *CreateVMRequestIde0) GetBps() int64`

GetBps returns the Bps field if non-nil, zero value otherwise.

### GetBpsOk

`func (o *CreateVMRequestIde0) GetBpsOk() (*int64, bool)`

GetBpsOk returns a tuple with the Bps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBps

`func (o *CreateVMRequestIde0) SetBps(v int64)`

SetBps sets Bps field to given value.

### HasBps

`func (o *CreateVMRequestIde0) HasBps() bool`

HasBps returns a boolean if a field has been set.

### GetBpsMaxLength

`func (o *CreateVMRequestIde0) GetBpsMaxLength() int64`

GetBpsMaxLength returns the BpsMaxLength field if non-nil, zero value otherwise.

### GetBpsMaxLengthOk

`func (o *CreateVMRequestIde0) GetBpsMaxLengthOk() (*int64, bool)`

GetBpsMaxLengthOk returns a tuple with the BpsMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsMaxLength

`func (o *CreateVMRequestIde0) SetBpsMaxLength(v int64)`

SetBpsMaxLength sets BpsMaxLength field to given value.

### HasBpsMaxLength

`func (o *CreateVMRequestIde0) HasBpsMaxLength() bool`

HasBpsMaxLength returns a boolean if a field has been set.

### GetBpsRd

`func (o *CreateVMRequestIde0) GetBpsRd() int64`

GetBpsRd returns the BpsRd field if non-nil, zero value otherwise.

### GetBpsRdOk

`func (o *CreateVMRequestIde0) GetBpsRdOk() (*int64, bool)`

GetBpsRdOk returns a tuple with the BpsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsRd

`func (o *CreateVMRequestIde0) SetBpsRd(v int64)`

SetBpsRd sets BpsRd field to given value.

### HasBpsRd

`func (o *CreateVMRequestIde0) HasBpsRd() bool`

HasBpsRd returns a boolean if a field has been set.

### GetBpsRdMaxLength

`func (o *CreateVMRequestIde0) GetBpsRdMaxLength() int64`

GetBpsRdMaxLength returns the BpsRdMaxLength field if non-nil, zero value otherwise.

### GetBpsRdMaxLengthOk

`func (o *CreateVMRequestIde0) GetBpsRdMaxLengthOk() (*int64, bool)`

GetBpsRdMaxLengthOk returns a tuple with the BpsRdMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsRdMaxLength

`func (o *CreateVMRequestIde0) SetBpsRdMaxLength(v int64)`

SetBpsRdMaxLength sets BpsRdMaxLength field to given value.

### HasBpsRdMaxLength

`func (o *CreateVMRequestIde0) HasBpsRdMaxLength() bool`

HasBpsRdMaxLength returns a boolean if a field has been set.

### GetBpsWr

`func (o *CreateVMRequestIde0) GetBpsWr() int64`

GetBpsWr returns the BpsWr field if non-nil, zero value otherwise.

### GetBpsWrOk

`func (o *CreateVMRequestIde0) GetBpsWrOk() (*int64, bool)`

GetBpsWrOk returns a tuple with the BpsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsWr

`func (o *CreateVMRequestIde0) SetBpsWr(v int64)`

SetBpsWr sets BpsWr field to given value.

### HasBpsWr

`func (o *CreateVMRequestIde0) HasBpsWr() bool`

HasBpsWr returns a boolean if a field has been set.

### GetBpsWrMaxLength

`func (o *CreateVMRequestIde0) GetBpsWrMaxLength() int64`

GetBpsWrMaxLength returns the BpsWrMaxLength field if non-nil, zero value otherwise.

### GetBpsWrMaxLengthOk

`func (o *CreateVMRequestIde0) GetBpsWrMaxLengthOk() (*int64, bool)`

GetBpsWrMaxLengthOk returns a tuple with the BpsWrMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsWrMaxLength

`func (o *CreateVMRequestIde0) SetBpsWrMaxLength(v int64)`

SetBpsWrMaxLength sets BpsWrMaxLength field to given value.

### HasBpsWrMaxLength

`func (o *CreateVMRequestIde0) HasBpsWrMaxLength() bool`

HasBpsWrMaxLength returns a boolean if a field has been set.

### GetCache

`func (o *CreateVMRequestIde0) GetCache() string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *CreateVMRequestIde0) GetCacheOk() (*string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *CreateVMRequestIde0) SetCache(v string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *CreateVMRequestIde0) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCyls

`func (o *CreateVMRequestIde0) GetCyls() int64`

GetCyls returns the Cyls field if non-nil, zero value otherwise.

### GetCylsOk

`func (o *CreateVMRequestIde0) GetCylsOk() (*int64, bool)`

GetCylsOk returns a tuple with the Cyls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyls

`func (o *CreateVMRequestIde0) SetCyls(v int64)`

SetCyls sets Cyls field to given value.

### HasCyls

`func (o *CreateVMRequestIde0) HasCyls() bool`

HasCyls returns a boolean if a field has been set.

### GetDetectZeroes

`func (o *CreateVMRequestIde0) GetDetectZeroes() bool`

GetDetectZeroes returns the DetectZeroes field if non-nil, zero value otherwise.

### GetDetectZeroesOk

`func (o *CreateVMRequestIde0) GetDetectZeroesOk() (*bool, bool)`

GetDetectZeroesOk returns a tuple with the DetectZeroes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectZeroes

`func (o *CreateVMRequestIde0) SetDetectZeroes(v bool)`

SetDetectZeroes sets DetectZeroes field to given value.

### HasDetectZeroes

`func (o *CreateVMRequestIde0) HasDetectZeroes() bool`

HasDetectZeroes returns a boolean if a field has been set.

### GetDiscard

`func (o *CreateVMRequestIde0) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *CreateVMRequestIde0) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *CreateVMRequestIde0) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *CreateVMRequestIde0) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetFile

`func (o *CreateVMRequestIde0) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateVMRequestIde0) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateVMRequestIde0) SetFile(v string)`

SetFile sets File field to given value.


### GetFormat

`func (o *CreateVMRequestIde0) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateVMRequestIde0) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateVMRequestIde0) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateVMRequestIde0) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHeads

`func (o *CreateVMRequestIde0) GetHeads() int64`

GetHeads returns the Heads field if non-nil, zero value otherwise.

### GetHeadsOk

`func (o *CreateVMRequestIde0) GetHeadsOk() (*int64, bool)`

GetHeadsOk returns a tuple with the Heads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeads

`func (o *CreateVMRequestIde0) SetHeads(v int64)`

SetHeads sets Heads field to given value.

### HasHeads

`func (o *CreateVMRequestIde0) HasHeads() bool`

HasHeads returns a boolean if a field has been set.

### GetImportFrom

`func (o *CreateVMRequestIde0) GetImportFrom() string`

GetImportFrom returns the ImportFrom field if non-nil, zero value otherwise.

### GetImportFromOk

`func (o *CreateVMRequestIde0) GetImportFromOk() (*string, bool)`

GetImportFromOk returns a tuple with the ImportFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFrom

`func (o *CreateVMRequestIde0) SetImportFrom(v string)`

SetImportFrom sets ImportFrom field to given value.

### HasImportFrom

`func (o *CreateVMRequestIde0) HasImportFrom() bool`

HasImportFrom returns a boolean if a field has been set.

### GetIops

`func (o *CreateVMRequestIde0) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *CreateVMRequestIde0) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *CreateVMRequestIde0) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *CreateVMRequestIde0) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetIopsMax

`func (o *CreateVMRequestIde0) GetIopsMax() int64`

GetIopsMax returns the IopsMax field if non-nil, zero value otherwise.

### GetIopsMaxOk

`func (o *CreateVMRequestIde0) GetIopsMaxOk() (*int64, bool)`

GetIopsMaxOk returns a tuple with the IopsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMax

`func (o *CreateVMRequestIde0) SetIopsMax(v int64)`

SetIopsMax sets IopsMax field to given value.

### HasIopsMax

`func (o *CreateVMRequestIde0) HasIopsMax() bool`

HasIopsMax returns a boolean if a field has been set.

### GetIopsMaxLength

`func (o *CreateVMRequestIde0) GetIopsMaxLength() int64`

GetIopsMaxLength returns the IopsMaxLength field if non-nil, zero value otherwise.

### GetIopsMaxLengthOk

`func (o *CreateVMRequestIde0) GetIopsMaxLengthOk() (*int64, bool)`

GetIopsMaxLengthOk returns a tuple with the IopsMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMaxLength

`func (o *CreateVMRequestIde0) SetIopsMaxLength(v int64)`

SetIopsMaxLength sets IopsMaxLength field to given value.

### HasIopsMaxLength

`func (o *CreateVMRequestIde0) HasIopsMaxLength() bool`

HasIopsMaxLength returns a boolean if a field has been set.

### GetIopsRd

`func (o *CreateVMRequestIde0) GetIopsRd() int64`

GetIopsRd returns the IopsRd field if non-nil, zero value otherwise.

### GetIopsRdOk

`func (o *CreateVMRequestIde0) GetIopsRdOk() (*int64, bool)`

GetIopsRdOk returns a tuple with the IopsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRd

`func (o *CreateVMRequestIde0) SetIopsRd(v int64)`

SetIopsRd sets IopsRd field to given value.

### HasIopsRd

`func (o *CreateVMRequestIde0) HasIopsRd() bool`

HasIopsRd returns a boolean if a field has been set.

### GetIopsRdMax

`func (o *CreateVMRequestIde0) GetIopsRdMax() int64`

GetIopsRdMax returns the IopsRdMax field if non-nil, zero value otherwise.

### GetIopsRdMaxOk

`func (o *CreateVMRequestIde0) GetIopsRdMaxOk() (*int64, bool)`

GetIopsRdMaxOk returns a tuple with the IopsRdMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRdMax

`func (o *CreateVMRequestIde0) SetIopsRdMax(v int64)`

SetIopsRdMax sets IopsRdMax field to given value.

### HasIopsRdMax

`func (o *CreateVMRequestIde0) HasIopsRdMax() bool`

HasIopsRdMax returns a boolean if a field has been set.

### GetIopsRdMaxLength

`func (o *CreateVMRequestIde0) GetIopsRdMaxLength() int64`

GetIopsRdMaxLength returns the IopsRdMaxLength field if non-nil, zero value otherwise.

### GetIopsRdMaxLengthOk

`func (o *CreateVMRequestIde0) GetIopsRdMaxLengthOk() (*int64, bool)`

GetIopsRdMaxLengthOk returns a tuple with the IopsRdMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRdMaxLength

`func (o *CreateVMRequestIde0) SetIopsRdMaxLength(v int64)`

SetIopsRdMaxLength sets IopsRdMaxLength field to given value.

### HasIopsRdMaxLength

`func (o *CreateVMRequestIde0) HasIopsRdMaxLength() bool`

HasIopsRdMaxLength returns a boolean if a field has been set.

### GetIopsWr

`func (o *CreateVMRequestIde0) GetIopsWr() int64`

GetIopsWr returns the IopsWr field if non-nil, zero value otherwise.

### GetIopsWrOk

`func (o *CreateVMRequestIde0) GetIopsWrOk() (*int64, bool)`

GetIopsWrOk returns a tuple with the IopsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWr

`func (o *CreateVMRequestIde0) SetIopsWr(v int64)`

SetIopsWr sets IopsWr field to given value.

### HasIopsWr

`func (o *CreateVMRequestIde0) HasIopsWr() bool`

HasIopsWr returns a boolean if a field has been set.

### GetIopsWrMax

`func (o *CreateVMRequestIde0) GetIopsWrMax() int64`

GetIopsWrMax returns the IopsWrMax field if non-nil, zero value otherwise.

### GetIopsWrMaxOk

`func (o *CreateVMRequestIde0) GetIopsWrMaxOk() (*int64, bool)`

GetIopsWrMaxOk returns a tuple with the IopsWrMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWrMax

`func (o *CreateVMRequestIde0) SetIopsWrMax(v int64)`

SetIopsWrMax sets IopsWrMax field to given value.

### HasIopsWrMax

`func (o *CreateVMRequestIde0) HasIopsWrMax() bool`

HasIopsWrMax returns a boolean if a field has been set.

### GetIopsWrMaxLength

`func (o *CreateVMRequestIde0) GetIopsWrMaxLength() int64`

GetIopsWrMaxLength returns the IopsWrMaxLength field if non-nil, zero value otherwise.

### GetIopsWrMaxLengthOk

`func (o *CreateVMRequestIde0) GetIopsWrMaxLengthOk() (*int64, bool)`

GetIopsWrMaxLengthOk returns a tuple with the IopsWrMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWrMaxLength

`func (o *CreateVMRequestIde0) SetIopsWrMaxLength(v int64)`

SetIopsWrMaxLength sets IopsWrMaxLength field to given value.

### HasIopsWrMaxLength

`func (o *CreateVMRequestIde0) HasIopsWrMaxLength() bool`

HasIopsWrMaxLength returns a boolean if a field has been set.

### GetMbps

`func (o *CreateVMRequestIde0) GetMbps() float32`

GetMbps returns the Mbps field if non-nil, zero value otherwise.

### GetMbpsOk

`func (o *CreateVMRequestIde0) GetMbpsOk() (*float32, bool)`

GetMbpsOk returns a tuple with the Mbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbps

`func (o *CreateVMRequestIde0) SetMbps(v float32)`

SetMbps sets Mbps field to given value.

### HasMbps

`func (o *CreateVMRequestIde0) HasMbps() bool`

HasMbps returns a boolean if a field has been set.

### GetMbpsMax

`func (o *CreateVMRequestIde0) GetMbpsMax() float32`

GetMbpsMax returns the MbpsMax field if non-nil, zero value otherwise.

### GetMbpsMaxOk

`func (o *CreateVMRequestIde0) GetMbpsMaxOk() (*float32, bool)`

GetMbpsMaxOk returns a tuple with the MbpsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsMax

`func (o *CreateVMRequestIde0) SetMbpsMax(v float32)`

SetMbpsMax sets MbpsMax field to given value.

### HasMbpsMax

`func (o *CreateVMRequestIde0) HasMbpsMax() bool`

HasMbpsMax returns a boolean if a field has been set.

### GetMbpsRd

`func (o *CreateVMRequestIde0) GetMbpsRd() float32`

GetMbpsRd returns the MbpsRd field if non-nil, zero value otherwise.

### GetMbpsRdOk

`func (o *CreateVMRequestIde0) GetMbpsRdOk() (*float32, bool)`

GetMbpsRdOk returns a tuple with the MbpsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsRd

`func (o *CreateVMRequestIde0) SetMbpsRd(v float32)`

SetMbpsRd sets MbpsRd field to given value.

### HasMbpsRd

`func (o *CreateVMRequestIde0) HasMbpsRd() bool`

HasMbpsRd returns a boolean if a field has been set.

### GetMbpsRdMax

`func (o *CreateVMRequestIde0) GetMbpsRdMax() float32`

GetMbpsRdMax returns the MbpsRdMax field if non-nil, zero value otherwise.

### GetMbpsRdMaxOk

`func (o *CreateVMRequestIde0) GetMbpsRdMaxOk() (*float32, bool)`

GetMbpsRdMaxOk returns a tuple with the MbpsRdMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsRdMax

`func (o *CreateVMRequestIde0) SetMbpsRdMax(v float32)`

SetMbpsRdMax sets MbpsRdMax field to given value.

### HasMbpsRdMax

`func (o *CreateVMRequestIde0) HasMbpsRdMax() bool`

HasMbpsRdMax returns a boolean if a field has been set.

### GetMbpsWr

`func (o *CreateVMRequestIde0) GetMbpsWr() float32`

GetMbpsWr returns the MbpsWr field if non-nil, zero value otherwise.

### GetMbpsWrOk

`func (o *CreateVMRequestIde0) GetMbpsWrOk() (*float32, bool)`

GetMbpsWrOk returns a tuple with the MbpsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsWr

`func (o *CreateVMRequestIde0) SetMbpsWr(v float32)`

SetMbpsWr sets MbpsWr field to given value.

### HasMbpsWr

`func (o *CreateVMRequestIde0) HasMbpsWr() bool`

HasMbpsWr returns a boolean if a field has been set.

### GetMbpsWrMax

`func (o *CreateVMRequestIde0) GetMbpsWrMax() float32`

GetMbpsWrMax returns the MbpsWrMax field if non-nil, zero value otherwise.

### GetMbpsWrMaxOk

`func (o *CreateVMRequestIde0) GetMbpsWrMaxOk() (*float32, bool)`

GetMbpsWrMaxOk returns a tuple with the MbpsWrMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsWrMax

`func (o *CreateVMRequestIde0) SetMbpsWrMax(v float32)`

SetMbpsWrMax sets MbpsWrMax field to given value.

### HasMbpsWrMax

`func (o *CreateVMRequestIde0) HasMbpsWrMax() bool`

HasMbpsWrMax returns a boolean if a field has been set.

### GetMedia

`func (o *CreateVMRequestIde0) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *CreateVMRequestIde0) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *CreateVMRequestIde0) SetMedia(v string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *CreateVMRequestIde0) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetModel

`func (o *CreateVMRequestIde0) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateVMRequestIde0) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateVMRequestIde0) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CreateVMRequestIde0) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetReplicate

`func (o *CreateVMRequestIde0) GetReplicate() bool`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *CreateVMRequestIde0) GetReplicateOk() (*bool, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *CreateVMRequestIde0) SetReplicate(v bool)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *CreateVMRequestIde0) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetRerror

`func (o *CreateVMRequestIde0) GetRerror() string`

GetRerror returns the Rerror field if non-nil, zero value otherwise.

### GetRerrorOk

`func (o *CreateVMRequestIde0) GetRerrorOk() (*string, bool)`

GetRerrorOk returns a tuple with the Rerror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerror

`func (o *CreateVMRequestIde0) SetRerror(v string)`

SetRerror sets Rerror field to given value.

### HasRerror

`func (o *CreateVMRequestIde0) HasRerror() bool`

HasRerror returns a boolean if a field has been set.

### GetSecs

`func (o *CreateVMRequestIde0) GetSecs() int64`

GetSecs returns the Secs field if non-nil, zero value otherwise.

### GetSecsOk

`func (o *CreateVMRequestIde0) GetSecsOk() (*int64, bool)`

GetSecsOk returns a tuple with the Secs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecs

`func (o *CreateVMRequestIde0) SetSecs(v int64)`

SetSecs sets Secs field to given value.

### HasSecs

`func (o *CreateVMRequestIde0) HasSecs() bool`

HasSecs returns a boolean if a field has been set.

### GetSerial

`func (o *CreateVMRequestIde0) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CreateVMRequestIde0) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CreateVMRequestIde0) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CreateVMRequestIde0) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetShared

`func (o *CreateVMRequestIde0) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateVMRequestIde0) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateVMRequestIde0) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateVMRequestIde0) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *CreateVMRequestIde0) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVMRequestIde0) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVMRequestIde0) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateVMRequestIde0) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapshot

`func (o *CreateVMRequestIde0) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CreateVMRequestIde0) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CreateVMRequestIde0) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CreateVMRequestIde0) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetSsd

`func (o *CreateVMRequestIde0) GetSsd() bool`

GetSsd returns the Ssd field if non-nil, zero value otherwise.

### GetSsdOk

`func (o *CreateVMRequestIde0) GetSsdOk() (*bool, bool)`

GetSsdOk returns a tuple with the Ssd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsd

`func (o *CreateVMRequestIde0) SetSsd(v bool)`

SetSsd sets Ssd field to given value.

### HasSsd

`func (o *CreateVMRequestIde0) HasSsd() bool`

HasSsd returns a boolean if a field has been set.

### GetTrans

`func (o *CreateVMRequestIde0) GetTrans() string`

GetTrans returns the Trans field if non-nil, zero value otherwise.

### GetTransOk

`func (o *CreateVMRequestIde0) GetTransOk() (*string, bool)`

GetTransOk returns a tuple with the Trans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrans

`func (o *CreateVMRequestIde0) SetTrans(v string)`

SetTrans sets Trans field to given value.

### HasTrans

`func (o *CreateVMRequestIde0) HasTrans() bool`

HasTrans returns a boolean if a field has been set.

### GetWerror

`func (o *CreateVMRequestIde0) GetWerror() string`

GetWerror returns the Werror field if non-nil, zero value otherwise.

### GetWerrorOk

`func (o *CreateVMRequestIde0) GetWerrorOk() (*string, bool)`

GetWerrorOk returns a tuple with the Werror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWerror

`func (o *CreateVMRequestIde0) SetWerror(v string)`

SetWerror sets Werror field to given value.

### HasWerror

`func (o *CreateVMRequestIde0) HasWerror() bool`

HasWerror returns a boolean if a field has been set.

### GetWwn

`func (o *CreateVMRequestIde0) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *CreateVMRequestIde0) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *CreateVMRequestIde0) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *CreateVMRequestIde0) HasWwn() bool`

HasWwn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


