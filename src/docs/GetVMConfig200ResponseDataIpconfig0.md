# GetVMConfig200ResponseDataIpconfig0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gw** | Pointer to **string** | gw | [optional] 
**Gw6** | Pointer to **string** | gw6 | [optional] 
**Ip** | Pointer to **string** | IP addresses use CIDR notation. The special string dhcp can be used for IP addresses to use DHCP in which case no explicit gateway should be provided. If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4 | [optional] 
**Ip6** | Pointer to **string** | For IPv6 the special string auto can be used to use stateless autoconfiguration | [optional] 

## Methods

### NewGetVMConfig200ResponseDataIpconfig0

`func NewGetVMConfig200ResponseDataIpconfig0() *GetVMConfig200ResponseDataIpconfig0`

NewGetVMConfig200ResponseDataIpconfig0 instantiates a new GetVMConfig200ResponseDataIpconfig0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataIpconfig0WithDefaults

`func NewGetVMConfig200ResponseDataIpconfig0WithDefaults() *GetVMConfig200ResponseDataIpconfig0`

NewGetVMConfig200ResponseDataIpconfig0WithDefaults instantiates a new GetVMConfig200ResponseDataIpconfig0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGw

`func (o *GetVMConfig200ResponseDataIpconfig0) GetGw() string`

GetGw returns the Gw field if non-nil, zero value otherwise.

### GetGwOk

`func (o *GetVMConfig200ResponseDataIpconfig0) GetGwOk() (*string, bool)`

GetGwOk returns a tuple with the Gw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw

`func (o *GetVMConfig200ResponseDataIpconfig0) SetGw(v string)`

SetGw sets Gw field to given value.

### HasGw

`func (o *GetVMConfig200ResponseDataIpconfig0) HasGw() bool`

HasGw returns a boolean if a field has been set.

### GetGw6

`func (o *GetVMConfig200ResponseDataIpconfig0) GetGw6() string`

GetGw6 returns the Gw6 field if non-nil, zero value otherwise.

### GetGw6Ok

`func (o *GetVMConfig200ResponseDataIpconfig0) GetGw6Ok() (*string, bool)`

GetGw6Ok returns a tuple with the Gw6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGw6

`func (o *GetVMConfig200ResponseDataIpconfig0) SetGw6(v string)`

SetGw6 sets Gw6 field to given value.

### HasGw6

`func (o *GetVMConfig200ResponseDataIpconfig0) HasGw6() bool`

HasGw6 returns a boolean if a field has been set.

### GetIp

`func (o *GetVMConfig200ResponseDataIpconfig0) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetVMConfig200ResponseDataIpconfig0) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetVMConfig200ResponseDataIpconfig0) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetVMConfig200ResponseDataIpconfig0) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *GetVMConfig200ResponseDataIpconfig0) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *GetVMConfig200ResponseDataIpconfig0) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *GetVMConfig200ResponseDataIpconfig0) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *GetVMConfig200ResponseDataIpconfig0) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


