# GetVMConfig200ResponseDataUsb0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The Host USB device or port or the value &#39;spice&#39;. HOSTUSBDEVICE syntax is:   &#39;bus-port(.port)*&#39; (decimal numbers) or  &#39;vendor_id:product_id&#39; (hexadeciaml numbers) or  &#39;spice&#39;  You can use the &#39;lsusb -t&#39; command to list existing usb devices.  NOTE: This option allows direct access to host hardware. So it is no longer possible to migrate such machines - use with special care.  The value &#39;spice&#39; can be used to add a usb redirection devices for spice.  Either this or the &#39;mapping&#39; key must be set.  | [optional] 
**Mapping** | Pointer to **string** | The ID of a cluster wide mapping. Either this or the default-key &#39;host&#39; must be set. | [optional] 
**Usb3** | Pointer to **int32** | Specifies whether if given host option is a USB3 device or port. For modern guests (machine version &gt;&#x3D; 7.1 and ostype l26 and windows &gt; 7), this flag is irrelevant (all devices are plugged into a xhci controller). | [optional] 

## Methods

### NewGetVMConfig200ResponseDataUsb0

`func NewGetVMConfig200ResponseDataUsb0() *GetVMConfig200ResponseDataUsb0`

NewGetVMConfig200ResponseDataUsb0 instantiates a new GetVMConfig200ResponseDataUsb0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVMConfig200ResponseDataUsb0WithDefaults

`func NewGetVMConfig200ResponseDataUsb0WithDefaults() *GetVMConfig200ResponseDataUsb0`

NewGetVMConfig200ResponseDataUsb0WithDefaults instantiates a new GetVMConfig200ResponseDataUsb0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *GetVMConfig200ResponseDataUsb0) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetVMConfig200ResponseDataUsb0) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetVMConfig200ResponseDataUsb0) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetVMConfig200ResponseDataUsb0) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMapping

`func (o *GetVMConfig200ResponseDataUsb0) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *GetVMConfig200ResponseDataUsb0) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *GetVMConfig200ResponseDataUsb0) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *GetVMConfig200ResponseDataUsb0) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetUsb3

`func (o *GetVMConfig200ResponseDataUsb0) GetUsb3() int32`

GetUsb3 returns the Usb3 field if non-nil, zero value otherwise.

### GetUsb3Ok

`func (o *GetVMConfig200ResponseDataUsb0) GetUsb3Ok() (*int32, bool)`

GetUsb3Ok returns a tuple with the Usb3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb3

`func (o *GetVMConfig200ResponseDataUsb0) SetUsb3(v int32)`

SetUsb3 sets Usb3 field to given value.

### HasUsb3

`func (o *GetVMConfig200ResponseDataUsb0) HasUsb3() bool`

HasUsb3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


