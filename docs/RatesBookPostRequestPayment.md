# RatesBookPostRequestPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HolderName** | **string** |  | 
**Number** | Pointer to **string** |  | [optional] 
**ExpireDate** | Pointer to **string** |  | [optional] 
**Cvc** | Pointer to **string** |  | [optional] 
**Method** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewRatesBookPostRequestPayment

`func NewRatesBookPostRequestPayment(holderName string, method string, ) *RatesBookPostRequestPayment`

NewRatesBookPostRequestPayment instantiates a new RatesBookPostRequestPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatesBookPostRequestPaymentWithDefaults

`func NewRatesBookPostRequestPaymentWithDefaults() *RatesBookPostRequestPayment`

NewRatesBookPostRequestPaymentWithDefaults instantiates a new RatesBookPostRequestPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolderName

`func (o *RatesBookPostRequestPayment) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *RatesBookPostRequestPayment) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *RatesBookPostRequestPayment) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.


### GetNumber

`func (o *RatesBookPostRequestPayment) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *RatesBookPostRequestPayment) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *RatesBookPostRequestPayment) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *RatesBookPostRequestPayment) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetExpireDate

`func (o *RatesBookPostRequestPayment) GetExpireDate() string`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *RatesBookPostRequestPayment) GetExpireDateOk() (*string, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *RatesBookPostRequestPayment) SetExpireDate(v string)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *RatesBookPostRequestPayment) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetCvc

`func (o *RatesBookPostRequestPayment) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *RatesBookPostRequestPayment) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *RatesBookPostRequestPayment) SetCvc(v string)`

SetCvc sets Cvc field to given value.

### HasCvc

`func (o *RatesBookPostRequestPayment) HasCvc() bool`

HasCvc returns a boolean if a field has been set.

### GetMethod

`func (o *RatesBookPostRequestPayment) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RatesBookPostRequestPayment) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RatesBookPostRequestPayment) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *RatesBookPostRequestPayment) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RatesBookPostRequestPayment) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RatesBookPostRequestPayment) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RatesBookPostRequestPayment) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


