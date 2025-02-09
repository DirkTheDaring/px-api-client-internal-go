# CreateAccessTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewFormat** | Pointer to **bool** |  | [optional] 
**Otp** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 
**Privs** | Pointer to **string** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**TfaChallenge** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewCreateAccessTicketRequest

`func NewCreateAccessTicketRequest(password string, username string, ) *CreateAccessTicketRequest`

NewCreateAccessTicketRequest instantiates a new CreateAccessTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessTicketRequestWithDefaults

`func NewCreateAccessTicketRequestWithDefaults() *CreateAccessTicketRequest`

NewCreateAccessTicketRequestWithDefaults instantiates a new CreateAccessTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewFormat

`func (o *CreateAccessTicketRequest) GetNewFormat() bool`

GetNewFormat returns the NewFormat field if non-nil, zero value otherwise.

### GetNewFormatOk

`func (o *CreateAccessTicketRequest) GetNewFormatOk() (*bool, bool)`

GetNewFormatOk returns a tuple with the NewFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFormat

`func (o *CreateAccessTicketRequest) SetNewFormat(v bool)`

SetNewFormat sets NewFormat field to given value.

### HasNewFormat

`func (o *CreateAccessTicketRequest) HasNewFormat() bool`

HasNewFormat returns a boolean if a field has been set.

### GetOtp

`func (o *CreateAccessTicketRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *CreateAccessTicketRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *CreateAccessTicketRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *CreateAccessTicketRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetPassword

`func (o *CreateAccessTicketRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateAccessTicketRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateAccessTicketRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPath

`func (o *CreateAccessTicketRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateAccessTicketRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateAccessTicketRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateAccessTicketRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPrivs

`func (o *CreateAccessTicketRequest) GetPrivs() string`

GetPrivs returns the Privs field if non-nil, zero value otherwise.

### GetPrivsOk

`func (o *CreateAccessTicketRequest) GetPrivsOk() (*string, bool)`

GetPrivsOk returns a tuple with the Privs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivs

`func (o *CreateAccessTicketRequest) SetPrivs(v string)`

SetPrivs sets Privs field to given value.

### HasPrivs

`func (o *CreateAccessTicketRequest) HasPrivs() bool`

HasPrivs returns a boolean if a field has been set.

### GetRealm

`func (o *CreateAccessTicketRequest) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *CreateAccessTicketRequest) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *CreateAccessTicketRequest) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *CreateAccessTicketRequest) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetTfaChallenge

`func (o *CreateAccessTicketRequest) GetTfaChallenge() string`

GetTfaChallenge returns the TfaChallenge field if non-nil, zero value otherwise.

### GetTfaChallengeOk

`func (o *CreateAccessTicketRequest) GetTfaChallengeOk() (*string, bool)`

GetTfaChallengeOk returns a tuple with the TfaChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaChallenge

`func (o *CreateAccessTicketRequest) SetTfaChallenge(v string)`

SetTfaChallenge sets TfaChallenge field to given value.

### HasTfaChallenge

`func (o *CreateAccessTicketRequest) HasTfaChallenge() bool`

HasTfaChallenge returns a boolean if a field has been set.

### GetUsername

`func (o *CreateAccessTicketRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateAccessTicketRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateAccessTicketRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


