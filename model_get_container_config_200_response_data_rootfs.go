/*
ProxMox VE API

ProxMox VE API

API version: 8.0
Contact: baldur@email.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pxapiobject

import (
	"encoding/json"
)

// checks if the GetContainerConfig200ResponseDataRootfs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContainerConfig200ResponseDataRootfs{}

// GetContainerConfig200ResponseDataRootfs struct for GetContainerConfig200ResponseDataRootfs
type GetContainerConfig200ResponseDataRootfs struct {
	// Explicitly enable or disable ACL support.
	Acl *int32 `json:"acl,omitempty"`
	// Extra mount options for rootfs/mps.
	Mountoptions *string `json:"mountoptions,omitempty"`
	// Enable user quotas inside the container (not supported with zfs subvolumes)
	Quota *int32 `json:"quota,omitempty"`
	// Will include this volume to a storage replica job.
	Replicate *int32 `json:"replicate,omitempty"`
	// Read-only mount point
	Ro *int32 `json:"ro,omitempty"`
	// Mark this non-volume mount point as available on multiple nodes (see 'nodes')
	Shared *int32 `json:"shared,omitempty"`
	// Volume size (read only value).
	Size *string `json:"size,omitempty"`
	// Volume, device or directory to mount into the container.
	Volume *string `json:"volume,omitempty"`
}

// NewGetContainerConfig200ResponseDataRootfs instantiates a new GetContainerConfig200ResponseDataRootfs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContainerConfig200ResponseDataRootfs() *GetContainerConfig200ResponseDataRootfs {
	this := GetContainerConfig200ResponseDataRootfs{}
	return &this
}

// NewGetContainerConfig200ResponseDataRootfsWithDefaults instantiates a new GetContainerConfig200ResponseDataRootfs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContainerConfig200ResponseDataRootfsWithDefaults() *GetContainerConfig200ResponseDataRootfs {
	this := GetContainerConfig200ResponseDataRootfs{}
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataRootfs) GetAcl() int32 {
	if o == nil || IsNil(o.Acl) {
		var ret int32
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataRootfs) GetAclOk() (*int32, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataRootfs) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given int32 and assigns it to the Acl field.
func (o *GetContainerConfig200ResponseDataRootfs) SetAcl(v int32) {
	o.Acl = &v
}

// GetMountoptions returns the Mountoptions field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataRootfs) GetMountoptions() string {
	if o == nil || IsNil(o.Mountoptions) {
		var ret string
		return ret
	}
	return *o.Mountoptions
}

// GetMountoptionsOk returns a tuple with the Mountoptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataRootfs) GetMountoptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Mountoptions) {
		return nil, false
	}
	return o.Mountoptions, true
}

// HasMountoptions returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataRootfs) HasMountoptions() bool {
	if o != nil && !IsNil(o.Mountoptions) {
		return true
	}

	return false
}

// SetMountoptions gets a reference to the given string and assigns it to the Mountoptions field.
func (o *GetContainerConfig200ResponseDataRootfs) SetMountoptions(v string) {
	o.Mountoptions = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataRootfs) GetQuota() int32 {
	if o == nil || IsNil(o.Quota) {
		var ret int32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataRootfs) GetQuotaOk() (*int32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataRootfs) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int32 and assigns it to the Quota field.
func (o *GetContainerConfig200ResponseDataRootfs) SetQuota(v int32) {
	o.Quota = &v
}

// GetReplicate returns the Replicate field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataRootfs) GetReplicate() int32 {
	if o == nil || IsNil(o.Replicate) {
		var ret int32
		return ret
	}
	return *o.Replicate
}

// GetReplicateOk returns a tuple with the Replicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataRootfs) GetReplicateOk() (*int32, bool) {
	if o == nil || IsNil(o.Replicate) {
		return nil, false
	}
	return o.Replicate, true
}

// HasReplicate returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataRootfs) HasReplicate() bool {
	if o != nil && !IsNil(o.Replicate) {
		return true
	}

	return false
}

// SetReplicate gets a reference to the given int32 and assigns it to the Replicate field.
func (o *GetContainerConfig200ResponseDataRootfs) SetReplicate(v int32) {
	o.Replicate = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataRootfs) GetRo() int32 {
	if o == nil || IsNil(o.Ro) {
		var ret int32
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataRootfs) GetRoOk() (*int32, bool) {
	if o == nil || IsNil(o.Ro) {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataRootfs) HasRo() bool {
	if o != nil && !IsNil(o.Ro) {
		return true
	}

	return false
}

// SetRo gets a reference to the given int32 and assigns it to the Ro field.
func (o *GetContainerConfig200ResponseDataRootfs) SetRo(v int32) {
	o.Ro = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataRootfs) GetShared() int32 {
	if o == nil || IsNil(o.Shared) {
		var ret int32
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataRootfs) GetSharedOk() (*int32, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataRootfs) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given int32 and assigns it to the Shared field.
func (o *GetContainerConfig200ResponseDataRootfs) SetShared(v int32) {
	o.Shared = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataRootfs) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataRootfs) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataRootfs) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *GetContainerConfig200ResponseDataRootfs) SetSize(v string) {
	o.Size = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *GetContainerConfig200ResponseDataRootfs) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerConfig200ResponseDataRootfs) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *GetContainerConfig200ResponseDataRootfs) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *GetContainerConfig200ResponseDataRootfs) SetVolume(v string) {
	o.Volume = &v
}

func (o GetContainerConfig200ResponseDataRootfs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContainerConfig200ResponseDataRootfs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	if !IsNil(o.Mountoptions) {
		toSerialize["mountoptions"] = o.Mountoptions
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.Replicate) {
		toSerialize["replicate"] = o.Replicate
	}
	if !IsNil(o.Ro) {
		toSerialize["ro"] = o.Ro
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableGetContainerConfig200ResponseDataRootfs struct {
	value *GetContainerConfig200ResponseDataRootfs
	isSet bool
}

func (v NullableGetContainerConfig200ResponseDataRootfs) Get() *GetContainerConfig200ResponseDataRootfs {
	return v.value
}

func (v *NullableGetContainerConfig200ResponseDataRootfs) Set(val *GetContainerConfig200ResponseDataRootfs) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContainerConfig200ResponseDataRootfs) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContainerConfig200ResponseDataRootfs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContainerConfig200ResponseDataRootfs(val *GetContainerConfig200ResponseDataRootfs) *NullableGetContainerConfig200ResponseDataRootfs {
	return &NullableGetContainerConfig200ResponseDataRootfs{value: val, isSet: true}
}

func (v NullableGetContainerConfig200ResponseDataRootfs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContainerConfig200ResponseDataRootfs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

