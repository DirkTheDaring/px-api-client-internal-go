# GetVMConfig200ResponseDataRng0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxBytes** | Pointer to **int64** | Maximum bytes of entropy allowed to get injected into the guest every &#39;period&#39; milliseconds. Prefer a lower value when using &#39;/dev/random&#39; as source. Use &#x60;0&#x60; to disable limiting (potentially dangerous!). | [optional] 
**Period** | Pointer to **int64** | Every &#39;period&#39; milliseconds the entropy-injection quota is reset, allowing the guest to retrieve another &#39;max_bytes&#39; of entropy. | [optional] 
**Source** | Pointer to **string** | The file on the host to gather entropy from. In most cases &#39;/dev/urandom&#39; should be preferred over &#39;/dev/random&#39; to avoid entropy-starvation issues on the host. Using urandom does *not* decrease security in any meaningful way, as it&#39;s still seeded from real entropy, and the bytes provided will most likely be mixed with real entropy on the guest as well. &#39;/dev/hwrng&#39; can be used to pass through a hardware RNG from the host. | [optional] 

## Methods

### NewGetVMConfig200ResponseDataRng0

`func NewGetVMConfig200ResponseDataRng0() *GetVMConfig200ResponseDataRng0`

NewGetVMConfig200ResponseDataRng0 instantiates a new GetVMConfig200ResponseDataRng0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataRng0WithDefaults

`func NewGetVMConfig200ResponseDataRng0WithDefaults() *GetVMConfig200ResponseDataRng0`

NewGetVMConfig200ResponseDataRng0WithDefaults instantiates a new GetVMConfig200ResponseDataRng0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxBytes

`func (o *GetVMConfig200ResponseDataRng0) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *GetVMConfig200ResponseDataRng0) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *GetVMConfig200ResponseDataRng0) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *GetVMConfig200ResponseDataRng0) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetPeriod

`func (o *GetVMConfig200ResponseDataRng0) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetVMConfig200ResponseDataRng0) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetVMConfig200ResponseDataRng0) SetPeriod(v int64)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetVMConfig200ResponseDataRng0) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetSource

`func (o *GetVMConfig200ResponseDataRng0) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetVMConfig200ResponseDataRng0) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetVMConfig200ResponseDataRng0) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetVMConfig200ResponseDataRng0) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


