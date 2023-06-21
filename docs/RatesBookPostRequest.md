# RatesBookPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrebookId** | **string** | prebook id retrived from prebook response | 
**GuestInfo** | [**RatesBookPostRequestGuestInfo**](RatesBookPostRequestGuestInfo.md) |  | 
**Payment** | Pointer to [**RatesBookPostRequestPayment**](RatesBookPostRequestPayment.md) |  | [optional] 

## Methods

### NewRatesBookPostRequest

`func NewRatesBookPostRequest(prebookId string, guestInfo RatesBookPostRequestGuestInfo, ) *RatesBookPostRequest`

NewRatesBookPostRequest instantiates a new RatesBookPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatesBookPostRequestWithDefaults

`func NewRatesBookPostRequestWithDefaults() *RatesBookPostRequest`

NewRatesBookPostRequestWithDefaults instantiates a new RatesBookPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrebookId

`func (o *RatesBookPostRequest) GetPrebookId() string`

GetPrebookId returns the PrebookId field if non-nil, zero value otherwise.

### GetPrebookIdOk

`func (o *RatesBookPostRequest) GetPrebookIdOk() (*string, bool)`

GetPrebookIdOk returns a tuple with the PrebookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrebookId

`func (o *RatesBookPostRequest) SetPrebookId(v string)`

SetPrebookId sets PrebookId field to given value.


### GetGuestInfo

`func (o *RatesBookPostRequest) GetGuestInfo() RatesBookPostRequestGuestInfo`

GetGuestInfo returns the GuestInfo field if non-nil, zero value otherwise.

### GetGuestInfoOk

`func (o *RatesBookPostRequest) GetGuestInfoOk() (*RatesBookPostRequestGuestInfo, bool)`

GetGuestInfoOk returns a tuple with the GuestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestInfo

`func (o *RatesBookPostRequest) SetGuestInfo(v RatesBookPostRequestGuestInfo)`

SetGuestInfo sets GuestInfo field to given value.


### GetPayment

`func (o *RatesBookPostRequest) GetPayment() RatesBookPostRequestPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *RatesBookPostRequest) GetPaymentOk() (*RatesBookPostRequestPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *RatesBookPostRequest) SetPayment(v RatesBookPostRequestPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *RatesBookPostRequest) HasPayment() bool`

HasPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


