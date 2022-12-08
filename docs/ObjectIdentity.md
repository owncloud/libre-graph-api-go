# ObjectIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | domain of the Provider issuing the identity | [optional] 
**IssuerAssignedId** | Pointer to **string** | The unique id assigned by the issuer to the account | [optional] 

## Methods

### NewObjectIdentity

`func NewObjectIdentity() *ObjectIdentity`

NewObjectIdentity instantiates a new ObjectIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectIdentityWithDefaults

`func NewObjectIdentityWithDefaults() *ObjectIdentity`

NewObjectIdentityWithDefaults instantiates a new ObjectIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *ObjectIdentity) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ObjectIdentity) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ObjectIdentity) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ObjectIdentity) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuerAssignedId

`func (o *ObjectIdentity) GetIssuerAssignedId() string`

GetIssuerAssignedId returns the IssuerAssignedId field if non-nil, zero value otherwise.

### GetIssuerAssignedIdOk

`func (o *ObjectIdentity) GetIssuerAssignedIdOk() (*string, bool)`

GetIssuerAssignedIdOk returns a tuple with the IssuerAssignedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAssignedId

`func (o *ObjectIdentity) SetIssuerAssignedId(v string)`

SetIssuerAssignedId sets IssuerAssignedId field to given value.

### HasIssuerAssignedId

`func (o *ObjectIdentity) HasIssuerAssignedId() bool`

HasIssuerAssignedId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


