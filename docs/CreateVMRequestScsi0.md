# CreateVMRequestScsi0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aio** | Pointer to **string** | AIO type to use. | [optional] 
**Backup** | Pointer to **bool** | Whether the drive should be included when making backups. | [optional] 
**Bps** | Pointer to **int64** | Maximum r/w speed in bytes per second. | [optional] 
**BpsMaxLength** | Pointer to **int64** | Maximum length of I/O bursts in seconds. | [optional] 
**BpsRd** | Pointer to **int64** | Maximum read speed in bytes per second. | [optional] 
**BpsRdLength** | Pointer to **string** |  | [optional] 
**BpsRdMaxLength** | Pointer to **int64** | Maximum length of read I/O bursts in seconds. | [optional] 
**BpsWr** | Pointer to **int64** | Maximum write speed in bytes per second. | [optional] 
**BpsWrLength** | Pointer to **string** |  | [optional] 
**BpsWrMaxLength** | Pointer to **int64** | Maximum length of write I/O bursts in seconds. | [optional] 
**Cache** | Pointer to **string** | The drive&#39;s cache mode | [optional] 
**Cyls** | Pointer to **int64** | Force the drive&#39;s physical geometry to have a specific cylinder count. | [optional] 
**DetectZeroes** | Pointer to **bool** | Controls whether to detect and try to optimize writes of zeroes. | [optional] 
**Discard** | Pointer to **string** | Controls whether to pass discard/trim requests to the underlying storage. | [optional] 
**File** | Pointer to **string** | The drive&#39;s backing volume. | [optional] 
**Format** | Pointer to **string** | The drive&#39;s backing file&#39;s data format. | [optional] 
**Heads** | Pointer to **int64** | Force the drive&#39;s physical geometry to have a specific head count. | [optional] 
**ImportFrom** | Pointer to **string** | Create a new disk, importing from this source (volume ID or absolute path). When an absolute path is specified, it&#39;s up to you to ensure that the source is not actively used by another process during the import! | [optional] 
**Iops** | Pointer to **int64** | Maximum r/w I/O in operations per second. | [optional] 
**IopsMax** | Pointer to **int64** | Maximum unthrottled r/w I/O pool in operations per second. | [optional] 
**IopsMaxLength** | Pointer to **int64** | Maximum length of I/O bursts in seconds. | [optional] 
**IopsRd** | Pointer to **int64** | Maximum read I/O in operations per second. | [optional] 
**IopsRdLength** | Pointer to **string** |  | [optional] 
**IopsRdMax** | Pointer to **int64** | Maximum unthrottled read I/O pool in operations per second. | [optional] 
**IopsRdMaxLength** | Pointer to **int64** | Maximum length of read I/O bursts in seconds. | [optional] 
**IopsWr** | Pointer to **int64** | Maximum write I/O in operations per second. | [optional] 
**IopsWrLength** | Pointer to **string** |  | [optional] 
**IopsWrMax** | Pointer to **int64** | Maximum unthrottled write I/O pool in operations per second. | [optional] 
**IopsWrMaxLength** | Pointer to **int64** | Maximum length of write I/O bursts in seconds. | [optional] 
**Iothread** | Pointer to **bool** | Whether to use iothreads for this drive | [optional] 
**Mbps** | Pointer to **float32** | Maximum r/w speed in megabytes per second. | [optional] 
**MbpsMax** | Pointer to **float32** | Maximum unthrottled r/w pool in megabytes per second. | [optional] 
**MbpsRd** | Pointer to **float32** | Maximum read speed in megabytes per second. | [optional] 
**MbpsRdMax** | Pointer to **float32** | Maximum unthrottled read pool in megabytes per second. | [optional] 
**MbpsWr** | Pointer to **float32** | Maximum write speed in megabytes per second. | [optional] 
**MbpsWrMax** | Pointer to **float32** | Maximum unthrottled write pool in megabytes per second. | [optional] 
**Media** | Pointer to **string** | The drive&#39;s media type. | [optional] 
**Queues** | Pointer to **int64** | Number of queues. | [optional] 
**Replicate** | Pointer to **bool** | Whether the drive should considered for replication jobs. | [optional] 
**Rerror** | Pointer to **string** | Read error action. | [optional] 
**Ro** | Pointer to **bool** | Whether the drive is read-only. | [optional] 
**Scsiblock** | Pointer to **bool** | whether to use scsi-block for full passthrough of host block device  WARNING: can lead to I/O errors in combination with low memory or high memory fragmentation on host | [optional] 
**Secs** | Pointer to **int64** | Force the drive&#39;s physical geometry to have a specific sector count. | [optional] 
**Serial** | Pointer to **string** | The drive&#39;s reported serial number, url-encoded, up to 20 bytes long. | [optional] 
**Shared** | Pointer to **bool** | Mark this locally-managed volume as available on all nodes | [optional] 
**Size** | Pointer to **string** | Disk size. This is purely informational and has no effect. | [optional] 
**Snapshot** | Pointer to **bool** | Controls qemu&#39;s snapshot mode feature. If activated, changes made to the disk are temporary and will be discarded when the VM is shutdown. | [optional] 
**Ssd** | Pointer to **bool** | Whether to expose this drive as an SSD, rather than a rotational hard disk. | [optional] 
**Trans** | Pointer to **string** | Force disk geometry bios translation mode. | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**Werror** | Pointer to **string** | Write error action. | [optional] 
**Wwn** | Pointer to **string** | The drive&#39;s worldwide name, encoded as 16 bytes hex string, prefixed by &#39;0x&#39;. | [optional] 

## Methods

### NewCreateVMRequestScsi0

`func NewCreateVMRequestScsi0() *CreateVMRequestScsi0`

NewCreateVMRequestScsi0 instantiates a new CreateVMRequestScsi0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMRequestScsi0WithDefaults

`func NewCreateVMRequestScsi0WithDefaults() *CreateVMRequestScsi0`

NewCreateVMRequestScsi0WithDefaults instantiates a new CreateVMRequestScsi0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAio

`func (o *CreateVMRequestScsi0) GetAio() string`

GetAio returns the Aio field if non-nil, zero value otherwise.

### GetAioOk

`func (o *CreateVMRequestScsi0) GetAioOk() (*string, bool)`

GetAioOk returns a tuple with the Aio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAio

`func (o *CreateVMRequestScsi0) SetAio(v string)`

SetAio sets Aio field to given value.

### HasAio

`func (o *CreateVMRequestScsi0) HasAio() bool`

HasAio returns a boolean if a field has been set.

### GetBackup

`func (o *CreateVMRequestScsi0) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateVMRequestScsi0) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateVMRequestScsi0) SetBackup(v bool)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *CreateVMRequestScsi0) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBps

`func (o *CreateVMRequestScsi0) GetBps() int64`

GetBps returns the Bps field if non-nil, zero value otherwise.

### GetBpsOk

`func (o *CreateVMRequestScsi0) GetBpsOk() (*int64, bool)`

GetBpsOk returns a tuple with the Bps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBps

`func (o *CreateVMRequestScsi0) SetBps(v int64)`

SetBps sets Bps field to given value.

### HasBps

`func (o *CreateVMRequestScsi0) HasBps() bool`

HasBps returns a boolean if a field has been set.

### GetBpsMaxLength

`func (o *CreateVMRequestScsi0) GetBpsMaxLength() int64`

GetBpsMaxLength returns the BpsMaxLength field if non-nil, zero value otherwise.

### GetBpsMaxLengthOk

`func (o *CreateVMRequestScsi0) GetBpsMaxLengthOk() (*int64, bool)`

GetBpsMaxLengthOk returns a tuple with the BpsMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsMaxLength

`func (o *CreateVMRequestScsi0) SetBpsMaxLength(v int64)`

SetBpsMaxLength sets BpsMaxLength field to given value.

### HasBpsMaxLength

`func (o *CreateVMRequestScsi0) HasBpsMaxLength() bool`

HasBpsMaxLength returns a boolean if a field has been set.

### GetBpsRd

`func (o *CreateVMRequestScsi0) GetBpsRd() int64`

GetBpsRd returns the BpsRd field if non-nil, zero value otherwise.

### GetBpsRdOk

`func (o *CreateVMRequestScsi0) GetBpsRdOk() (*int64, bool)`

GetBpsRdOk returns a tuple with the BpsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsRd

`func (o *CreateVMRequestScsi0) SetBpsRd(v int64)`

SetBpsRd sets BpsRd field to given value.

### HasBpsRd

`func (o *CreateVMRequestScsi0) HasBpsRd() bool`

HasBpsRd returns a boolean if a field has been set.

### GetBpsRdLength

`func (o *CreateVMRequestScsi0) GetBpsRdLength() string`

GetBpsRdLength returns the BpsRdLength field if non-nil, zero value otherwise.

### GetBpsRdLengthOk

`func (o *CreateVMRequestScsi0) GetBpsRdLengthOk() (*string, bool)`

GetBpsRdLengthOk returns a tuple with the BpsRdLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsRdLength

`func (o *CreateVMRequestScsi0) SetBpsRdLength(v string)`

SetBpsRdLength sets BpsRdLength field to given value.

### HasBpsRdLength

`func (o *CreateVMRequestScsi0) HasBpsRdLength() bool`

HasBpsRdLength returns a boolean if a field has been set.

### GetBpsRdMaxLength

`func (o *CreateVMRequestScsi0) GetBpsRdMaxLength() int64`

GetBpsRdMaxLength returns the BpsRdMaxLength field if non-nil, zero value otherwise.

### GetBpsRdMaxLengthOk

`func (o *CreateVMRequestScsi0) GetBpsRdMaxLengthOk() (*int64, bool)`

GetBpsRdMaxLengthOk returns a tuple with the BpsRdMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsRdMaxLength

`func (o *CreateVMRequestScsi0) SetBpsRdMaxLength(v int64)`

SetBpsRdMaxLength sets BpsRdMaxLength field to given value.

### HasBpsRdMaxLength

`func (o *CreateVMRequestScsi0) HasBpsRdMaxLength() bool`

HasBpsRdMaxLength returns a boolean if a field has been set.

### GetBpsWr

`func (o *CreateVMRequestScsi0) GetBpsWr() int64`

GetBpsWr returns the BpsWr field if non-nil, zero value otherwise.

### GetBpsWrOk

`func (o *CreateVMRequestScsi0) GetBpsWrOk() (*int64, bool)`

GetBpsWrOk returns a tuple with the BpsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsWr

`func (o *CreateVMRequestScsi0) SetBpsWr(v int64)`

SetBpsWr sets BpsWr field to given value.

### HasBpsWr

`func (o *CreateVMRequestScsi0) HasBpsWr() bool`

HasBpsWr returns a boolean if a field has been set.

### GetBpsWrLength

`func (o *CreateVMRequestScsi0) GetBpsWrLength() string`

GetBpsWrLength returns the BpsWrLength field if non-nil, zero value otherwise.

### GetBpsWrLengthOk

`func (o *CreateVMRequestScsi0) GetBpsWrLengthOk() (*string, bool)`

GetBpsWrLengthOk returns a tuple with the BpsWrLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsWrLength

`func (o *CreateVMRequestScsi0) SetBpsWrLength(v string)`

SetBpsWrLength sets BpsWrLength field to given value.

### HasBpsWrLength

`func (o *CreateVMRequestScsi0) HasBpsWrLength() bool`

HasBpsWrLength returns a boolean if a field has been set.

### GetBpsWrMaxLength

`func (o *CreateVMRequestScsi0) GetBpsWrMaxLength() int64`

GetBpsWrMaxLength returns the BpsWrMaxLength field if non-nil, zero value otherwise.

### GetBpsWrMaxLengthOk

`func (o *CreateVMRequestScsi0) GetBpsWrMaxLengthOk() (*int64, bool)`

GetBpsWrMaxLengthOk returns a tuple with the BpsWrMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsWrMaxLength

`func (o *CreateVMRequestScsi0) SetBpsWrMaxLength(v int64)`

SetBpsWrMaxLength sets BpsWrMaxLength field to given value.

### HasBpsWrMaxLength

`func (o *CreateVMRequestScsi0) HasBpsWrMaxLength() bool`

HasBpsWrMaxLength returns a boolean if a field has been set.

### GetCache

`func (o *CreateVMRequestScsi0) GetCache() string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *CreateVMRequestScsi0) GetCacheOk() (*string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *CreateVMRequestScsi0) SetCache(v string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *CreateVMRequestScsi0) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCyls

`func (o *CreateVMRequestScsi0) GetCyls() int64`

GetCyls returns the Cyls field if non-nil, zero value otherwise.

### GetCylsOk

`func (o *CreateVMRequestScsi0) GetCylsOk() (*int64, bool)`

GetCylsOk returns a tuple with the Cyls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyls

`func (o *CreateVMRequestScsi0) SetCyls(v int64)`

SetCyls sets Cyls field to given value.

### HasCyls

`func (o *CreateVMRequestScsi0) HasCyls() bool`

HasCyls returns a boolean if a field has been set.

### GetDetectZeroes

`func (o *CreateVMRequestScsi0) GetDetectZeroes() bool`

GetDetectZeroes returns the DetectZeroes field if non-nil, zero value otherwise.

### GetDetectZeroesOk

`func (o *CreateVMRequestScsi0) GetDetectZeroesOk() (*bool, bool)`

GetDetectZeroesOk returns a tuple with the DetectZeroes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectZeroes

`func (o *CreateVMRequestScsi0) SetDetectZeroes(v bool)`

SetDetectZeroes sets DetectZeroes field to given value.

### HasDetectZeroes

`func (o *CreateVMRequestScsi0) HasDetectZeroes() bool`

HasDetectZeroes returns a boolean if a field has been set.

### GetDiscard

`func (o *CreateVMRequestScsi0) GetDiscard() string`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *CreateVMRequestScsi0) GetDiscardOk() (*string, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *CreateVMRequestScsi0) SetDiscard(v string)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *CreateVMRequestScsi0) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetFile

`func (o *CreateVMRequestScsi0) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateVMRequestScsi0) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateVMRequestScsi0) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *CreateVMRequestScsi0) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFormat

`func (o *CreateVMRequestScsi0) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateVMRequestScsi0) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateVMRequestScsi0) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateVMRequestScsi0) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHeads

`func (o *CreateVMRequestScsi0) GetHeads() int64`

GetHeads returns the Heads field if non-nil, zero value otherwise.

### GetHeadsOk

`func (o *CreateVMRequestScsi0) GetHeadsOk() (*int64, bool)`

GetHeadsOk returns a tuple with the Heads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeads

`func (o *CreateVMRequestScsi0) SetHeads(v int64)`

SetHeads sets Heads field to given value.

### HasHeads

`func (o *CreateVMRequestScsi0) HasHeads() bool`

HasHeads returns a boolean if a field has been set.

### GetImportFrom

`func (o *CreateVMRequestScsi0) GetImportFrom() string`

GetImportFrom returns the ImportFrom field if non-nil, zero value otherwise.

### GetImportFromOk

`func (o *CreateVMRequestScsi0) GetImportFromOk() (*string, bool)`

GetImportFromOk returns a tuple with the ImportFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFrom

`func (o *CreateVMRequestScsi0) SetImportFrom(v string)`

SetImportFrom sets ImportFrom field to given value.

### HasImportFrom

`func (o *CreateVMRequestScsi0) HasImportFrom() bool`

HasImportFrom returns a boolean if a field has been set.

### GetIops

`func (o *CreateVMRequestScsi0) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *CreateVMRequestScsi0) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *CreateVMRequestScsi0) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *CreateVMRequestScsi0) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetIopsMax

`func (o *CreateVMRequestScsi0) GetIopsMax() int64`

GetIopsMax returns the IopsMax field if non-nil, zero value otherwise.

### GetIopsMaxOk

`func (o *CreateVMRequestScsi0) GetIopsMaxOk() (*int64, bool)`

GetIopsMaxOk returns a tuple with the IopsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMax

`func (o *CreateVMRequestScsi0) SetIopsMax(v int64)`

SetIopsMax sets IopsMax field to given value.

### HasIopsMax

`func (o *CreateVMRequestScsi0) HasIopsMax() bool`

HasIopsMax returns a boolean if a field has been set.

### GetIopsMaxLength

`func (o *CreateVMRequestScsi0) GetIopsMaxLength() int64`

GetIopsMaxLength returns the IopsMaxLength field if non-nil, zero value otherwise.

### GetIopsMaxLengthOk

`func (o *CreateVMRequestScsi0) GetIopsMaxLengthOk() (*int64, bool)`

GetIopsMaxLengthOk returns a tuple with the IopsMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMaxLength

`func (o *CreateVMRequestScsi0) SetIopsMaxLength(v int64)`

SetIopsMaxLength sets IopsMaxLength field to given value.

### HasIopsMaxLength

`func (o *CreateVMRequestScsi0) HasIopsMaxLength() bool`

HasIopsMaxLength returns a boolean if a field has been set.

### GetIopsRd

`func (o *CreateVMRequestScsi0) GetIopsRd() int64`

GetIopsRd returns the IopsRd field if non-nil, zero value otherwise.

### GetIopsRdOk

`func (o *CreateVMRequestScsi0) GetIopsRdOk() (*int64, bool)`

GetIopsRdOk returns a tuple with the IopsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRd

`func (o *CreateVMRequestScsi0) SetIopsRd(v int64)`

SetIopsRd sets IopsRd field to given value.

### HasIopsRd

`func (o *CreateVMRequestScsi0) HasIopsRd() bool`

HasIopsRd returns a boolean if a field has been set.

### GetIopsRdLength

`func (o *CreateVMRequestScsi0) GetIopsRdLength() string`

GetIopsRdLength returns the IopsRdLength field if non-nil, zero value otherwise.

### GetIopsRdLengthOk

`func (o *CreateVMRequestScsi0) GetIopsRdLengthOk() (*string, bool)`

GetIopsRdLengthOk returns a tuple with the IopsRdLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRdLength

`func (o *CreateVMRequestScsi0) SetIopsRdLength(v string)`

SetIopsRdLength sets IopsRdLength field to given value.

### HasIopsRdLength

`func (o *CreateVMRequestScsi0) HasIopsRdLength() bool`

HasIopsRdLength returns a boolean if a field has been set.

### GetIopsRdMax

`func (o *CreateVMRequestScsi0) GetIopsRdMax() int64`

GetIopsRdMax returns the IopsRdMax field if non-nil, zero value otherwise.

### GetIopsRdMaxOk

`func (o *CreateVMRequestScsi0) GetIopsRdMaxOk() (*int64, bool)`

GetIopsRdMaxOk returns a tuple with the IopsRdMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRdMax

`func (o *CreateVMRequestScsi0) SetIopsRdMax(v int64)`

SetIopsRdMax sets IopsRdMax field to given value.

### HasIopsRdMax

`func (o *CreateVMRequestScsi0) HasIopsRdMax() bool`

HasIopsRdMax returns a boolean if a field has been set.

### GetIopsRdMaxLength

`func (o *CreateVMRequestScsi0) GetIopsRdMaxLength() int64`

GetIopsRdMaxLength returns the IopsRdMaxLength field if non-nil, zero value otherwise.

### GetIopsRdMaxLengthOk

`func (o *CreateVMRequestScsi0) GetIopsRdMaxLengthOk() (*int64, bool)`

GetIopsRdMaxLengthOk returns a tuple with the IopsRdMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsRdMaxLength

`func (o *CreateVMRequestScsi0) SetIopsRdMaxLength(v int64)`

SetIopsRdMaxLength sets IopsRdMaxLength field to given value.

### HasIopsRdMaxLength

`func (o *CreateVMRequestScsi0) HasIopsRdMaxLength() bool`

HasIopsRdMaxLength returns a boolean if a field has been set.

### GetIopsWr

`func (o *CreateVMRequestScsi0) GetIopsWr() int64`

GetIopsWr returns the IopsWr field if non-nil, zero value otherwise.

### GetIopsWrOk

`func (o *CreateVMRequestScsi0) GetIopsWrOk() (*int64, bool)`

GetIopsWrOk returns a tuple with the IopsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWr

`func (o *CreateVMRequestScsi0) SetIopsWr(v int64)`

SetIopsWr sets IopsWr field to given value.

### HasIopsWr

`func (o *CreateVMRequestScsi0) HasIopsWr() bool`

HasIopsWr returns a boolean if a field has been set.

### GetIopsWrLength

`func (o *CreateVMRequestScsi0) GetIopsWrLength() string`

GetIopsWrLength returns the IopsWrLength field if non-nil, zero value otherwise.

### GetIopsWrLengthOk

`func (o *CreateVMRequestScsi0) GetIopsWrLengthOk() (*string, bool)`

GetIopsWrLengthOk returns a tuple with the IopsWrLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWrLength

`func (o *CreateVMRequestScsi0) SetIopsWrLength(v string)`

SetIopsWrLength sets IopsWrLength field to given value.

### HasIopsWrLength

`func (o *CreateVMRequestScsi0) HasIopsWrLength() bool`

HasIopsWrLength returns a boolean if a field has been set.

### GetIopsWrMax

`func (o *CreateVMRequestScsi0) GetIopsWrMax() int64`

GetIopsWrMax returns the IopsWrMax field if non-nil, zero value otherwise.

### GetIopsWrMaxOk

`func (o *CreateVMRequestScsi0) GetIopsWrMaxOk() (*int64, bool)`

GetIopsWrMaxOk returns a tuple with the IopsWrMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWrMax

`func (o *CreateVMRequestScsi0) SetIopsWrMax(v int64)`

SetIopsWrMax sets IopsWrMax field to given value.

### HasIopsWrMax

`func (o *CreateVMRequestScsi0) HasIopsWrMax() bool`

HasIopsWrMax returns a boolean if a field has been set.

### GetIopsWrMaxLength

`func (o *CreateVMRequestScsi0) GetIopsWrMaxLength() int64`

GetIopsWrMaxLength returns the IopsWrMaxLength field if non-nil, zero value otherwise.

### GetIopsWrMaxLengthOk

`func (o *CreateVMRequestScsi0) GetIopsWrMaxLengthOk() (*int64, bool)`

GetIopsWrMaxLengthOk returns a tuple with the IopsWrMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsWrMaxLength

`func (o *CreateVMRequestScsi0) SetIopsWrMaxLength(v int64)`

SetIopsWrMaxLength sets IopsWrMaxLength field to given value.

### HasIopsWrMaxLength

`func (o *CreateVMRequestScsi0) HasIopsWrMaxLength() bool`

HasIopsWrMaxLength returns a boolean if a field has been set.

### GetIothread

`func (o *CreateVMRequestScsi0) GetIothread() bool`

GetIothread returns the Iothread field if non-nil, zero value otherwise.

### GetIothreadOk

`func (o *CreateVMRequestScsi0) GetIothreadOk() (*bool, bool)`

GetIothreadOk returns a tuple with the Iothread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIothread

`func (o *CreateVMRequestScsi0) SetIothread(v bool)`

SetIothread sets Iothread field to given value.

### HasIothread

`func (o *CreateVMRequestScsi0) HasIothread() bool`

HasIothread returns a boolean if a field has been set.

### GetMbps

`func (o *CreateVMRequestScsi0) GetMbps() float32`

GetMbps returns the Mbps field if non-nil, zero value otherwise.

### GetMbpsOk

`func (o *CreateVMRequestScsi0) GetMbpsOk() (*float32, bool)`

GetMbpsOk returns a tuple with the Mbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbps

`func (o *CreateVMRequestScsi0) SetMbps(v float32)`

SetMbps sets Mbps field to given value.

### HasMbps

`func (o *CreateVMRequestScsi0) HasMbps() bool`

HasMbps returns a boolean if a field has been set.

### GetMbpsMax

`func (o *CreateVMRequestScsi0) GetMbpsMax() float32`

GetMbpsMax returns the MbpsMax field if non-nil, zero value otherwise.

### GetMbpsMaxOk

`func (o *CreateVMRequestScsi0) GetMbpsMaxOk() (*float32, bool)`

GetMbpsMaxOk returns a tuple with the MbpsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsMax

`func (o *CreateVMRequestScsi0) SetMbpsMax(v float32)`

SetMbpsMax sets MbpsMax field to given value.

### HasMbpsMax

`func (o *CreateVMRequestScsi0) HasMbpsMax() bool`

HasMbpsMax returns a boolean if a field has been set.

### GetMbpsRd

`func (o *CreateVMRequestScsi0) GetMbpsRd() float32`

GetMbpsRd returns the MbpsRd field if non-nil, zero value otherwise.

### GetMbpsRdOk

`func (o *CreateVMRequestScsi0) GetMbpsRdOk() (*float32, bool)`

GetMbpsRdOk returns a tuple with the MbpsRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsRd

`func (o *CreateVMRequestScsi0) SetMbpsRd(v float32)`

SetMbpsRd sets MbpsRd field to given value.

### HasMbpsRd

`func (o *CreateVMRequestScsi0) HasMbpsRd() bool`

HasMbpsRd returns a boolean if a field has been set.

### GetMbpsRdMax

`func (o *CreateVMRequestScsi0) GetMbpsRdMax() float32`

GetMbpsRdMax returns the MbpsRdMax field if non-nil, zero value otherwise.

### GetMbpsRdMaxOk

`func (o *CreateVMRequestScsi0) GetMbpsRdMaxOk() (*float32, bool)`

GetMbpsRdMaxOk returns a tuple with the MbpsRdMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsRdMax

`func (o *CreateVMRequestScsi0) SetMbpsRdMax(v float32)`

SetMbpsRdMax sets MbpsRdMax field to given value.

### HasMbpsRdMax

`func (o *CreateVMRequestScsi0) HasMbpsRdMax() bool`

HasMbpsRdMax returns a boolean if a field has been set.

### GetMbpsWr

`func (o *CreateVMRequestScsi0) GetMbpsWr() float32`

GetMbpsWr returns the MbpsWr field if non-nil, zero value otherwise.

### GetMbpsWrOk

`func (o *CreateVMRequestScsi0) GetMbpsWrOk() (*float32, bool)`

GetMbpsWrOk returns a tuple with the MbpsWr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsWr

`func (o *CreateVMRequestScsi0) SetMbpsWr(v float32)`

SetMbpsWr sets MbpsWr field to given value.

### HasMbpsWr

`func (o *CreateVMRequestScsi0) HasMbpsWr() bool`

HasMbpsWr returns a boolean if a field has been set.

### GetMbpsWrMax

`func (o *CreateVMRequestScsi0) GetMbpsWrMax() float32`

GetMbpsWrMax returns the MbpsWrMax field if non-nil, zero value otherwise.

### GetMbpsWrMaxOk

`func (o *CreateVMRequestScsi0) GetMbpsWrMaxOk() (*float32, bool)`

GetMbpsWrMaxOk returns a tuple with the MbpsWrMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbpsWrMax

`func (o *CreateVMRequestScsi0) SetMbpsWrMax(v float32)`

SetMbpsWrMax sets MbpsWrMax field to given value.

### HasMbpsWrMax

`func (o *CreateVMRequestScsi0) HasMbpsWrMax() bool`

HasMbpsWrMax returns a boolean if a field has been set.

### GetMedia

`func (o *CreateVMRequestScsi0) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *CreateVMRequestScsi0) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *CreateVMRequestScsi0) SetMedia(v string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *CreateVMRequestScsi0) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetQueues

`func (o *CreateVMRequestScsi0) GetQueues() int64`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *CreateVMRequestScsi0) GetQueuesOk() (*int64, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *CreateVMRequestScsi0) SetQueues(v int64)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *CreateVMRequestScsi0) HasQueues() bool`

HasQueues returns a boolean if a field has been set.

### GetReplicate

`func (o *CreateVMRequestScsi0) GetReplicate() bool`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *CreateVMRequestScsi0) GetReplicateOk() (*bool, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *CreateVMRequestScsi0) SetReplicate(v bool)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *CreateVMRequestScsi0) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetRerror

`func (o *CreateVMRequestScsi0) GetRerror() string`

GetRerror returns the Rerror field if non-nil, zero value otherwise.

### GetRerrorOk

`func (o *CreateVMRequestScsi0) GetRerrorOk() (*string, bool)`

GetRerrorOk returns a tuple with the Rerror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerror

`func (o *CreateVMRequestScsi0) SetRerror(v string)`

SetRerror sets Rerror field to given value.

### HasRerror

`func (o *CreateVMRequestScsi0) HasRerror() bool`

HasRerror returns a boolean if a field has been set.

### GetRo

`func (o *CreateVMRequestScsi0) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *CreateVMRequestScsi0) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *CreateVMRequestScsi0) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *CreateVMRequestScsi0) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetScsiblock

`func (o *CreateVMRequestScsi0) GetScsiblock() bool`

GetScsiblock returns the Scsiblock field if non-nil, zero value otherwise.

### GetScsiblockOk

`func (o *CreateVMRequestScsi0) GetScsiblockOk() (*bool, bool)`

GetScsiblockOk returns a tuple with the Scsiblock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsiblock

`func (o *CreateVMRequestScsi0) SetScsiblock(v bool)`

SetScsiblock sets Scsiblock field to given value.

### HasScsiblock

`func (o *CreateVMRequestScsi0) HasScsiblock() bool`

HasScsiblock returns a boolean if a field has been set.

### GetSecs

`func (o *CreateVMRequestScsi0) GetSecs() int64`

GetSecs returns the Secs field if non-nil, zero value otherwise.

### GetSecsOk

`func (o *CreateVMRequestScsi0) GetSecsOk() (*int64, bool)`

GetSecsOk returns a tuple with the Secs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecs

`func (o *CreateVMRequestScsi0) SetSecs(v int64)`

SetSecs sets Secs field to given value.

### HasSecs

`func (o *CreateVMRequestScsi0) HasSecs() bool`

HasSecs returns a boolean if a field has been set.

### GetSerial

`func (o *CreateVMRequestScsi0) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CreateVMRequestScsi0) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CreateVMRequestScsi0) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CreateVMRequestScsi0) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetShared

`func (o *CreateVMRequestScsi0) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateVMRequestScsi0) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateVMRequestScsi0) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateVMRequestScsi0) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSize

`func (o *CreateVMRequestScsi0) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVMRequestScsi0) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVMRequestScsi0) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateVMRequestScsi0) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapshot

`func (o *CreateVMRequestScsi0) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CreateVMRequestScsi0) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CreateVMRequestScsi0) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CreateVMRequestScsi0) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetSsd

`func (o *CreateVMRequestScsi0) GetSsd() bool`

GetSsd returns the Ssd field if non-nil, zero value otherwise.

### GetSsdOk

`func (o *CreateVMRequestScsi0) GetSsdOk() (*bool, bool)`

GetSsdOk returns a tuple with the Ssd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsd

`func (o *CreateVMRequestScsi0) SetSsd(v bool)`

SetSsd sets Ssd field to given value.

### HasSsd

`func (o *CreateVMRequestScsi0) HasSsd() bool`

HasSsd returns a boolean if a field has been set.

### GetTrans

`func (o *CreateVMRequestScsi0) GetTrans() string`

GetTrans returns the Trans field if non-nil, zero value otherwise.

### GetTransOk

`func (o *CreateVMRequestScsi0) GetTransOk() (*string, bool)`

GetTransOk returns a tuple with the Trans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrans

`func (o *CreateVMRequestScsi0) SetTrans(v string)`

SetTrans sets Trans field to given value.

### HasTrans

`func (o *CreateVMRequestScsi0) HasTrans() bool`

HasTrans returns a boolean if a field has been set.

### GetVolume

`func (o *CreateVMRequestScsi0) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateVMRequestScsi0) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateVMRequestScsi0) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CreateVMRequestScsi0) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWerror

`func (o *CreateVMRequestScsi0) GetWerror() string`

GetWerror returns the Werror field if non-nil, zero value otherwise.

### GetWerrorOk

`func (o *CreateVMRequestScsi0) GetWerrorOk() (*string, bool)`

GetWerrorOk returns a tuple with the Werror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWerror

`func (o *CreateVMRequestScsi0) SetWerror(v string)`

SetWerror sets Werror field to given value.

### HasWerror

`func (o *CreateVMRequestScsi0) HasWerror() bool`

HasWerror returns a boolean if a field has been set.

### GetWwn

`func (o *CreateVMRequestScsi0) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *CreateVMRequestScsi0) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *CreateVMRequestScsi0) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *CreateVMRequestScsi0) HasWwn() bool`

HasWwn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


