# CreateStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authsupported** | Pointer to **string** |  | [optional] 
**Base** | Pointer to **string** |  | [optional] 
**Blocksize** | Pointer to **string** |  | [optional] 
**Bwlimit** | Pointer to [**CreateStorageRequestBwlimit**](CreateStorageRequestBwlimit.md) |  | [optional] 
**ComstarHg** | Pointer to **string** |  | [optional] 
**ComstarTg** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ContentDirs** | Pointer to **string** |  | [optional] 
**CreateBasePath** | Pointer to **bool** |  | [optional] 
**CreateSubdirs** | Pointer to **bool** |  | [optional] 
**DataPool** | Pointer to **string** |  | [optional] 
**Datastore** | Pointer to **string** |  | [optional] 
**Disable** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**EncryptionKey** | Pointer to **string** |  | [optional] 
**Export** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**FsName** | Pointer to **string** |  | [optional] 
**Fuse** | Pointer to **bool** |  | [optional] 
**IsMountpoint** | Pointer to **string** |  | [optional] 
**Iscsiprovider** | Pointer to **string** |  | [optional] 
**Keyring** | Pointer to **string** |  | [optional] 
**Krbd** | Pointer to **bool** |  | [optional] 
**LioTpg** | Pointer to **string** |  | [optional] 
**MasterPubkey** | Pointer to **string** |  | [optional] 
**MaxProtectedBackups** | Pointer to **int64** |  | [optional] 
**Maxfiles** | Pointer to **int64** |  | [optional] 
**Mkdir** | Pointer to **bool** |  | [optional] 
**Monhost** | Pointer to **string** |  | [optional] 
**Mountpoint** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Nocow** | Pointer to **bool** |  | [optional] 
**Nodes** | Pointer to **string** |  | [optional] 
**Nowritecache** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Pool** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Portal** | Pointer to **string** |  | [optional] 
**Preallocation** | Pointer to **string** |  | [optional] 
**PruneBackups** | Pointer to **string** |  | [optional] 
**Saferemove** | Pointer to **bool** |  | [optional] 
**SaferemoveThroughput** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**Server2** | Pointer to **string** |  | [optional] 
**Share** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**SkipCertVerification** | Pointer to **bool** |  | [optional] 
**Smbversion** | Pointer to **string** |  | [optional] 
**Sparse** | Pointer to **bool** |  | [optional] 
**Storage** | **string** |  | 
**Subdir** | Pointer to **string** |  | [optional] 
**TaggedOnly** | Pointer to **bool** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Thinpool** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 
**Vgname** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateStorageRequest

`func NewCreateStorageRequest(storage string, type_ string, ) *CreateStorageRequest`

NewCreateStorageRequest instantiates a new CreateStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageRequestWithDefaults

`func NewCreateStorageRequestWithDefaults() *CreateStorageRequest`

NewCreateStorageRequestWithDefaults instantiates a new CreateStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthsupported

`func (o *CreateStorageRequest) GetAuthsupported() string`

GetAuthsupported returns the Authsupported field if non-nil, zero value otherwise.

### GetAuthsupportedOk

`func (o *CreateStorageRequest) GetAuthsupportedOk() (*string, bool)`

GetAuthsupportedOk returns a tuple with the Authsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthsupported

`func (o *CreateStorageRequest) SetAuthsupported(v string)`

SetAuthsupported sets Authsupported field to given value.

### HasAuthsupported

`func (o *CreateStorageRequest) HasAuthsupported() bool`

HasAuthsupported returns a boolean if a field has been set.

### GetBase

`func (o *CreateStorageRequest) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *CreateStorageRequest) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *CreateStorageRequest) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *CreateStorageRequest) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBlocksize

`func (o *CreateStorageRequest) GetBlocksize() string`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *CreateStorageRequest) GetBlocksizeOk() (*string, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *CreateStorageRequest) SetBlocksize(v string)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *CreateStorageRequest) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetBwlimit

`func (o *CreateStorageRequest) GetBwlimit() CreateStorageRequestBwlimit`

GetBwlimit returns the Bwlimit field if non-nil, zero value otherwise.

### GetBwlimitOk

`func (o *CreateStorageRequest) GetBwlimitOk() (*CreateStorageRequestBwlimit, bool)`

GetBwlimitOk returns a tuple with the Bwlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwlimit

`func (o *CreateStorageRequest) SetBwlimit(v CreateStorageRequestBwlimit)`

SetBwlimit sets Bwlimit field to given value.

### HasBwlimit

`func (o *CreateStorageRequest) HasBwlimit() bool`

HasBwlimit returns a boolean if a field has been set.

### GetComstarHg

`func (o *CreateStorageRequest) GetComstarHg() string`

GetComstarHg returns the ComstarHg field if non-nil, zero value otherwise.

### GetComstarHgOk

`func (o *CreateStorageRequest) GetComstarHgOk() (*string, bool)`

GetComstarHgOk returns a tuple with the ComstarHg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComstarHg

`func (o *CreateStorageRequest) SetComstarHg(v string)`

SetComstarHg sets ComstarHg field to given value.

### HasComstarHg

`func (o *CreateStorageRequest) HasComstarHg() bool`

HasComstarHg returns a boolean if a field has been set.

### GetComstarTg

`func (o *CreateStorageRequest) GetComstarTg() string`

GetComstarTg returns the ComstarTg field if non-nil, zero value otherwise.

### GetComstarTgOk

`func (o *CreateStorageRequest) GetComstarTgOk() (*string, bool)`

GetComstarTgOk returns a tuple with the ComstarTg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComstarTg

`func (o *CreateStorageRequest) SetComstarTg(v string)`

SetComstarTg sets ComstarTg field to given value.

### HasComstarTg

`func (o *CreateStorageRequest) HasComstarTg() bool`

HasComstarTg returns a boolean if a field has been set.

### GetContent

`func (o *CreateStorageRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateStorageRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateStorageRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreateStorageRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentDirs

`func (o *CreateStorageRequest) GetContentDirs() string`

GetContentDirs returns the ContentDirs field if non-nil, zero value otherwise.

### GetContentDirsOk

`func (o *CreateStorageRequest) GetContentDirsOk() (*string, bool)`

GetContentDirsOk returns a tuple with the ContentDirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDirs

`func (o *CreateStorageRequest) SetContentDirs(v string)`

SetContentDirs sets ContentDirs field to given value.

### HasContentDirs

`func (o *CreateStorageRequest) HasContentDirs() bool`

HasContentDirs returns a boolean if a field has been set.

### GetCreateBasePath

`func (o *CreateStorageRequest) GetCreateBasePath() bool`

GetCreateBasePath returns the CreateBasePath field if non-nil, zero value otherwise.

### GetCreateBasePathOk

`func (o *CreateStorageRequest) GetCreateBasePathOk() (*bool, bool)`

GetCreateBasePathOk returns a tuple with the CreateBasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBasePath

`func (o *CreateStorageRequest) SetCreateBasePath(v bool)`

SetCreateBasePath sets CreateBasePath field to given value.

### HasCreateBasePath

`func (o *CreateStorageRequest) HasCreateBasePath() bool`

HasCreateBasePath returns a boolean if a field has been set.

### GetCreateSubdirs

`func (o *CreateStorageRequest) GetCreateSubdirs() bool`

GetCreateSubdirs returns the CreateSubdirs field if non-nil, zero value otherwise.

### GetCreateSubdirsOk

`func (o *CreateStorageRequest) GetCreateSubdirsOk() (*bool, bool)`

GetCreateSubdirsOk returns a tuple with the CreateSubdirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSubdirs

`func (o *CreateStorageRequest) SetCreateSubdirs(v bool)`

SetCreateSubdirs sets CreateSubdirs field to given value.

### HasCreateSubdirs

`func (o *CreateStorageRequest) HasCreateSubdirs() bool`

HasCreateSubdirs returns a boolean if a field has been set.

### GetDataPool

`func (o *CreateStorageRequest) GetDataPool() string`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *CreateStorageRequest) GetDataPoolOk() (*string, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *CreateStorageRequest) SetDataPool(v string)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *CreateStorageRequest) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### GetDatastore

`func (o *CreateStorageRequest) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *CreateStorageRequest) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *CreateStorageRequest) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *CreateStorageRequest) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDisable

`func (o *CreateStorageRequest) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *CreateStorageRequest) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *CreateStorageRequest) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *CreateStorageRequest) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDomain

`func (o *CreateStorageRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateStorageRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateStorageRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateStorageRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *CreateStorageRequest) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *CreateStorageRequest) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *CreateStorageRequest) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *CreateStorageRequest) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetExport

`func (o *CreateStorageRequest) GetExport() string`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *CreateStorageRequest) GetExportOk() (*string, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *CreateStorageRequest) SetExport(v string)`

SetExport sets Export field to given value.

### HasExport

`func (o *CreateStorageRequest) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetFingerprint

`func (o *CreateStorageRequest) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CreateStorageRequest) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CreateStorageRequest) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CreateStorageRequest) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFormat

`func (o *CreateStorageRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateStorageRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateStorageRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateStorageRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFsName

`func (o *CreateStorageRequest) GetFsName() string`

GetFsName returns the FsName field if non-nil, zero value otherwise.

### GetFsNameOk

`func (o *CreateStorageRequest) GetFsNameOk() (*string, bool)`

GetFsNameOk returns a tuple with the FsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsName

`func (o *CreateStorageRequest) SetFsName(v string)`

SetFsName sets FsName field to given value.

### HasFsName

`func (o *CreateStorageRequest) HasFsName() bool`

HasFsName returns a boolean if a field has been set.

### GetFuse

`func (o *CreateStorageRequest) GetFuse() bool`

GetFuse returns the Fuse field if non-nil, zero value otherwise.

### GetFuseOk

`func (o *CreateStorageRequest) GetFuseOk() (*bool, bool)`

GetFuseOk returns a tuple with the Fuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuse

`func (o *CreateStorageRequest) SetFuse(v bool)`

SetFuse sets Fuse field to given value.

### HasFuse

`func (o *CreateStorageRequest) HasFuse() bool`

HasFuse returns a boolean if a field has been set.

### GetIsMountpoint

`func (o *CreateStorageRequest) GetIsMountpoint() string`

GetIsMountpoint returns the IsMountpoint field if non-nil, zero value otherwise.

### GetIsMountpointOk

`func (o *CreateStorageRequest) GetIsMountpointOk() (*string, bool)`

GetIsMountpointOk returns a tuple with the IsMountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMountpoint

`func (o *CreateStorageRequest) SetIsMountpoint(v string)`

SetIsMountpoint sets IsMountpoint field to given value.

### HasIsMountpoint

`func (o *CreateStorageRequest) HasIsMountpoint() bool`

HasIsMountpoint returns a boolean if a field has been set.

### GetIscsiprovider

`func (o *CreateStorageRequest) GetIscsiprovider() string`

GetIscsiprovider returns the Iscsiprovider field if non-nil, zero value otherwise.

### GetIscsiproviderOk

`func (o *CreateStorageRequest) GetIscsiproviderOk() (*string, bool)`

GetIscsiproviderOk returns a tuple with the Iscsiprovider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiprovider

`func (o *CreateStorageRequest) SetIscsiprovider(v string)`

SetIscsiprovider sets Iscsiprovider field to given value.

### HasIscsiprovider

`func (o *CreateStorageRequest) HasIscsiprovider() bool`

HasIscsiprovider returns a boolean if a field has been set.

### GetKeyring

`func (o *CreateStorageRequest) GetKeyring() string`

GetKeyring returns the Keyring field if non-nil, zero value otherwise.

### GetKeyringOk

`func (o *CreateStorageRequest) GetKeyringOk() (*string, bool)`

GetKeyringOk returns a tuple with the Keyring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyring

`func (o *CreateStorageRequest) SetKeyring(v string)`

SetKeyring sets Keyring field to given value.

### HasKeyring

`func (o *CreateStorageRequest) HasKeyring() bool`

HasKeyring returns a boolean if a field has been set.

### GetKrbd

`func (o *CreateStorageRequest) GetKrbd() bool`

GetKrbd returns the Krbd field if non-nil, zero value otherwise.

### GetKrbdOk

`func (o *CreateStorageRequest) GetKrbdOk() (*bool, bool)`

GetKrbdOk returns a tuple with the Krbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbd

`func (o *CreateStorageRequest) SetKrbd(v bool)`

SetKrbd sets Krbd field to given value.

### HasKrbd

`func (o *CreateStorageRequest) HasKrbd() bool`

HasKrbd returns a boolean if a field has been set.

### GetLioTpg

`func (o *CreateStorageRequest) GetLioTpg() string`

GetLioTpg returns the LioTpg field if non-nil, zero value otherwise.

### GetLioTpgOk

`func (o *CreateStorageRequest) GetLioTpgOk() (*string, bool)`

GetLioTpgOk returns a tuple with the LioTpg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLioTpg

`func (o *CreateStorageRequest) SetLioTpg(v string)`

SetLioTpg sets LioTpg field to given value.

### HasLioTpg

`func (o *CreateStorageRequest) HasLioTpg() bool`

HasLioTpg returns a boolean if a field has been set.

### GetMasterPubkey

`func (o *CreateStorageRequest) GetMasterPubkey() string`

GetMasterPubkey returns the MasterPubkey field if non-nil, zero value otherwise.

### GetMasterPubkeyOk

`func (o *CreateStorageRequest) GetMasterPubkeyOk() (*string, bool)`

GetMasterPubkeyOk returns a tuple with the MasterPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPubkey

`func (o *CreateStorageRequest) SetMasterPubkey(v string)`

SetMasterPubkey sets MasterPubkey field to given value.

### HasMasterPubkey

`func (o *CreateStorageRequest) HasMasterPubkey() bool`

HasMasterPubkey returns a boolean if a field has been set.

### GetMaxProtectedBackups

`func (o *CreateStorageRequest) GetMaxProtectedBackups() int64`

GetMaxProtectedBackups returns the MaxProtectedBackups field if non-nil, zero value otherwise.

### GetMaxProtectedBackupsOk

`func (o *CreateStorageRequest) GetMaxProtectedBackupsOk() (*int64, bool)`

GetMaxProtectedBackupsOk returns a tuple with the MaxProtectedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProtectedBackups

`func (o *CreateStorageRequest) SetMaxProtectedBackups(v int64)`

SetMaxProtectedBackups sets MaxProtectedBackups field to given value.

### HasMaxProtectedBackups

`func (o *CreateStorageRequest) HasMaxProtectedBackups() bool`

HasMaxProtectedBackups returns a boolean if a field has been set.

### GetMaxfiles

`func (o *CreateStorageRequest) GetMaxfiles() int64`

GetMaxfiles returns the Maxfiles field if non-nil, zero value otherwise.

### GetMaxfilesOk

`func (o *CreateStorageRequest) GetMaxfilesOk() (*int64, bool)`

GetMaxfilesOk returns a tuple with the Maxfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxfiles

`func (o *CreateStorageRequest) SetMaxfiles(v int64)`

SetMaxfiles sets Maxfiles field to given value.

### HasMaxfiles

`func (o *CreateStorageRequest) HasMaxfiles() bool`

HasMaxfiles returns a boolean if a field has been set.

### GetMkdir

`func (o *CreateStorageRequest) GetMkdir() bool`

GetMkdir returns the Mkdir field if non-nil, zero value otherwise.

### GetMkdirOk

`func (o *CreateStorageRequest) GetMkdirOk() (*bool, bool)`

GetMkdirOk returns a tuple with the Mkdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMkdir

`func (o *CreateStorageRequest) SetMkdir(v bool)`

SetMkdir sets Mkdir field to given value.

### HasMkdir

`func (o *CreateStorageRequest) HasMkdir() bool`

HasMkdir returns a boolean if a field has been set.

### GetMonhost

`func (o *CreateStorageRequest) GetMonhost() string`

GetMonhost returns the Monhost field if non-nil, zero value otherwise.

### GetMonhostOk

`func (o *CreateStorageRequest) GetMonhostOk() (*string, bool)`

GetMonhostOk returns a tuple with the Monhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonhost

`func (o *CreateStorageRequest) SetMonhost(v string)`

SetMonhost sets Monhost field to given value.

### HasMonhost

`func (o *CreateStorageRequest) HasMonhost() bool`

HasMonhost returns a boolean if a field has been set.

### GetMountpoint

`func (o *CreateStorageRequest) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *CreateStorageRequest) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *CreateStorageRequest) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *CreateStorageRequest) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetNamespace

`func (o *CreateStorageRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateStorageRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateStorageRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateStorageRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNocow

`func (o *CreateStorageRequest) GetNocow() bool`

GetNocow returns the Nocow field if non-nil, zero value otherwise.

### GetNocowOk

`func (o *CreateStorageRequest) GetNocowOk() (*bool, bool)`

GetNocowOk returns a tuple with the Nocow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocow

`func (o *CreateStorageRequest) SetNocow(v bool)`

SetNocow sets Nocow field to given value.

### HasNocow

`func (o *CreateStorageRequest) HasNocow() bool`

HasNocow returns a boolean if a field has been set.

### GetNodes

`func (o *CreateStorageRequest) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CreateStorageRequest) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CreateStorageRequest) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CreateStorageRequest) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNowritecache

`func (o *CreateStorageRequest) GetNowritecache() bool`

GetNowritecache returns the Nowritecache field if non-nil, zero value otherwise.

### GetNowritecacheOk

`func (o *CreateStorageRequest) GetNowritecacheOk() (*bool, bool)`

GetNowritecacheOk returns a tuple with the Nowritecache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowritecache

`func (o *CreateStorageRequest) SetNowritecache(v bool)`

SetNowritecache sets Nowritecache field to given value.

### HasNowritecache

`func (o *CreateStorageRequest) HasNowritecache() bool`

HasNowritecache returns a boolean if a field has been set.

### GetOptions

`func (o *CreateStorageRequest) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateStorageRequest) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateStorageRequest) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateStorageRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPassword

`func (o *CreateStorageRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateStorageRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateStorageRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateStorageRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *CreateStorageRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateStorageRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateStorageRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateStorageRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPool

`func (o *CreateStorageRequest) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CreateStorageRequest) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CreateStorageRequest) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CreateStorageRequest) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPort

`func (o *CreateStorageRequest) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateStorageRequest) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateStorageRequest) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateStorageRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortal

`func (o *CreateStorageRequest) GetPortal() string`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *CreateStorageRequest) GetPortalOk() (*string, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *CreateStorageRequest) SetPortal(v string)`

SetPortal sets Portal field to given value.

### HasPortal

`func (o *CreateStorageRequest) HasPortal() bool`

HasPortal returns a boolean if a field has been set.

### GetPreallocation

`func (o *CreateStorageRequest) GetPreallocation() string`

GetPreallocation returns the Preallocation field if non-nil, zero value otherwise.

### GetPreallocationOk

`func (o *CreateStorageRequest) GetPreallocationOk() (*string, bool)`

GetPreallocationOk returns a tuple with the Preallocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreallocation

`func (o *CreateStorageRequest) SetPreallocation(v string)`

SetPreallocation sets Preallocation field to given value.

### HasPreallocation

`func (o *CreateStorageRequest) HasPreallocation() bool`

HasPreallocation returns a boolean if a field has been set.

### GetPruneBackups

`func (o *CreateStorageRequest) GetPruneBackups() string`

GetPruneBackups returns the PruneBackups field if non-nil, zero value otherwise.

### GetPruneBackupsOk

`func (o *CreateStorageRequest) GetPruneBackupsOk() (*string, bool)`

GetPruneBackupsOk returns a tuple with the PruneBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruneBackups

`func (o *CreateStorageRequest) SetPruneBackups(v string)`

SetPruneBackups sets PruneBackups field to given value.

### HasPruneBackups

`func (o *CreateStorageRequest) HasPruneBackups() bool`

HasPruneBackups returns a boolean if a field has been set.

### GetSaferemove

`func (o *CreateStorageRequest) GetSaferemove() bool`

GetSaferemove returns the Saferemove field if non-nil, zero value otherwise.

### GetSaferemoveOk

`func (o *CreateStorageRequest) GetSaferemoveOk() (*bool, bool)`

GetSaferemoveOk returns a tuple with the Saferemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaferemove

`func (o *CreateStorageRequest) SetSaferemove(v bool)`

SetSaferemove sets Saferemove field to given value.

### HasSaferemove

`func (o *CreateStorageRequest) HasSaferemove() bool`

HasSaferemove returns a boolean if a field has been set.

### GetSaferemoveThroughput

`func (o *CreateStorageRequest) GetSaferemoveThroughput() string`

GetSaferemoveThroughput returns the SaferemoveThroughput field if non-nil, zero value otherwise.

### GetSaferemoveThroughputOk

`func (o *CreateStorageRequest) GetSaferemoveThroughputOk() (*string, bool)`

GetSaferemoveThroughputOk returns a tuple with the SaferemoveThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaferemoveThroughput

`func (o *CreateStorageRequest) SetSaferemoveThroughput(v string)`

SetSaferemoveThroughput sets SaferemoveThroughput field to given value.

### HasSaferemoveThroughput

`func (o *CreateStorageRequest) HasSaferemoveThroughput() bool`

HasSaferemoveThroughput returns a boolean if a field has been set.

### GetServer

`func (o *CreateStorageRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateStorageRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateStorageRequest) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *CreateStorageRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetServer2

`func (o *CreateStorageRequest) GetServer2() string`

GetServer2 returns the Server2 field if non-nil, zero value otherwise.

### GetServer2Ok

`func (o *CreateStorageRequest) GetServer2Ok() (*string, bool)`

GetServer2Ok returns a tuple with the Server2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer2

`func (o *CreateStorageRequest) SetServer2(v string)`

SetServer2 sets Server2 field to given value.

### HasServer2

`func (o *CreateStorageRequest) HasServer2() bool`

HasServer2 returns a boolean if a field has been set.

### GetShare

`func (o *CreateStorageRequest) GetShare() string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *CreateStorageRequest) GetShareOk() (*string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *CreateStorageRequest) SetShare(v string)`

SetShare sets Share field to given value.

### HasShare

`func (o *CreateStorageRequest) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetShared

`func (o *CreateStorageRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateStorageRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateStorageRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateStorageRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSkipCertVerification

`func (o *CreateStorageRequest) GetSkipCertVerification() bool`

GetSkipCertVerification returns the SkipCertVerification field if non-nil, zero value otherwise.

### GetSkipCertVerificationOk

`func (o *CreateStorageRequest) GetSkipCertVerificationOk() (*bool, bool)`

GetSkipCertVerificationOk returns a tuple with the SkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCertVerification

`func (o *CreateStorageRequest) SetSkipCertVerification(v bool)`

SetSkipCertVerification sets SkipCertVerification field to given value.

### HasSkipCertVerification

`func (o *CreateStorageRequest) HasSkipCertVerification() bool`

HasSkipCertVerification returns a boolean if a field has been set.

### GetSmbversion

`func (o *CreateStorageRequest) GetSmbversion() string`

GetSmbversion returns the Smbversion field if non-nil, zero value otherwise.

### GetSmbversionOk

`func (o *CreateStorageRequest) GetSmbversionOk() (*string, bool)`

GetSmbversionOk returns a tuple with the Smbversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbversion

`func (o *CreateStorageRequest) SetSmbversion(v string)`

SetSmbversion sets Smbversion field to given value.

### HasSmbversion

`func (o *CreateStorageRequest) HasSmbversion() bool`

HasSmbversion returns a boolean if a field has been set.

### GetSparse

`func (o *CreateStorageRequest) GetSparse() bool`

GetSparse returns the Sparse field if non-nil, zero value otherwise.

### GetSparseOk

`func (o *CreateStorageRequest) GetSparseOk() (*bool, bool)`

GetSparseOk returns a tuple with the Sparse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparse

`func (o *CreateStorageRequest) SetSparse(v bool)`

SetSparse sets Sparse field to given value.

### HasSparse

`func (o *CreateStorageRequest) HasSparse() bool`

HasSparse returns a boolean if a field has been set.

### GetStorage

`func (o *CreateStorageRequest) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateStorageRequest) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateStorageRequest) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetSubdir

`func (o *CreateStorageRequest) GetSubdir() string`

GetSubdir returns the Subdir field if non-nil, zero value otherwise.

### GetSubdirOk

`func (o *CreateStorageRequest) GetSubdirOk() (*string, bool)`

GetSubdirOk returns a tuple with the Subdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdir

`func (o *CreateStorageRequest) SetSubdir(v string)`

SetSubdir sets Subdir field to given value.

### HasSubdir

`func (o *CreateStorageRequest) HasSubdir() bool`

HasSubdir returns a boolean if a field has been set.

### GetTaggedOnly

`func (o *CreateStorageRequest) GetTaggedOnly() bool`

GetTaggedOnly returns the TaggedOnly field if non-nil, zero value otherwise.

### GetTaggedOnlyOk

`func (o *CreateStorageRequest) GetTaggedOnlyOk() (*bool, bool)`

GetTaggedOnlyOk returns a tuple with the TaggedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedOnly

`func (o *CreateStorageRequest) SetTaggedOnly(v bool)`

SetTaggedOnly sets TaggedOnly field to given value.

### HasTaggedOnly

`func (o *CreateStorageRequest) HasTaggedOnly() bool`

HasTaggedOnly returns a boolean if a field has been set.

### GetTarget

`func (o *CreateStorageRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateStorageRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateStorageRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CreateStorageRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetThinpool

`func (o *CreateStorageRequest) GetThinpool() string`

GetThinpool returns the Thinpool field if non-nil, zero value otherwise.

### GetThinpoolOk

`func (o *CreateStorageRequest) GetThinpoolOk() (*string, bool)`

GetThinpoolOk returns a tuple with the Thinpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinpool

`func (o *CreateStorageRequest) SetThinpool(v string)`

SetThinpool sets Thinpool field to given value.

### HasThinpool

`func (o *CreateStorageRequest) HasThinpool() bool`

HasThinpool returns a boolean if a field has been set.

### GetTransport

`func (o *CreateStorageRequest) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *CreateStorageRequest) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *CreateStorageRequest) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *CreateStorageRequest) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetType

`func (o *CreateStorageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStorageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStorageRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUsername

`func (o *CreateStorageRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateStorageRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateStorageRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateStorageRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVgname

`func (o *CreateStorageRequest) GetVgname() string`

GetVgname returns the Vgname field if non-nil, zero value otherwise.

### GetVgnameOk

`func (o *CreateStorageRequest) GetVgnameOk() (*string, bool)`

GetVgnameOk returns a tuple with the Vgname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgname

`func (o *CreateStorageRequest) SetVgname(v string)`

SetVgname sets Vgname field to given value.

### HasVgname

`func (o *CreateStorageRequest) HasVgname() bool`

HasVgname returns a boolean if a field has been set.

### GetVolume

`func (o *CreateStorageRequest) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateStorageRequest) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateStorageRequest) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CreateStorageRequest) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


