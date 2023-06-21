# HotelsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checkin** | Pointer to **string** |  | [optional] 
**Checkout** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **float32** |  | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **int32** |  | [optional] 
**Adults** | Pointer to **int32** |  | [optional] 
**Children** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**GuestNationality** | Pointer to **string** |  | [optional] 

## Methods

### NewHotelsPostRequest

`func NewHotelsPostRequest() *HotelsPostRequest`

NewHotelsPostRequest instantiates a new HotelsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotelsPostRequestWithDefaults

`func NewHotelsPostRequestWithDefaults() *HotelsPostRequest`

NewHotelsPostRequestWithDefaults instantiates a new HotelsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckin

`func (o *HotelsPostRequest) GetCheckin() string`

GetCheckin returns the Checkin field if non-nil, zero value otherwise.

### GetCheckinOk

`func (o *HotelsPostRequest) GetCheckinOk() (*string, bool)`

GetCheckinOk returns a tuple with the Checkin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckin

`func (o *HotelsPostRequest) SetCheckin(v string)`

SetCheckin sets Checkin field to given value.

### HasCheckin

`func (o *HotelsPostRequest) HasCheckin() bool`

HasCheckin returns a boolean if a field has been set.

### GetCheckout

`func (o *HotelsPostRequest) GetCheckout() string`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *HotelsPostRequest) GetCheckoutOk() (*string, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *HotelsPostRequest) SetCheckout(v string)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *HotelsPostRequest) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### GetLatitude

`func (o *HotelsPostRequest) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *HotelsPostRequest) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *HotelsPostRequest) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *HotelsPostRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *HotelsPostRequest) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *HotelsPostRequest) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *HotelsPostRequest) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *HotelsPostRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetCountry

`func (o *HotelsPostRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *HotelsPostRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *HotelsPostRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *HotelsPostRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDistance

`func (o *HotelsPostRequest) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *HotelsPostRequest) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *HotelsPostRequest) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *HotelsPostRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetAdults

`func (o *HotelsPostRequest) GetAdults() int32`

GetAdults returns the Adults field if non-nil, zero value otherwise.

### GetAdultsOk

`func (o *HotelsPostRequest) GetAdultsOk() (*int32, bool)`

GetAdultsOk returns a tuple with the Adults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdults

`func (o *HotelsPostRequest) SetAdults(v int32)`

SetAdults sets Adults field to given value.

### HasAdults

`func (o *HotelsPostRequest) HasAdults() bool`

HasAdults returns a boolean if a field has been set.

### GetChildren

`func (o *HotelsPostRequest) GetChildren() string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HotelsPostRequest) GetChildrenOk() (*string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HotelsPostRequest) SetChildren(v string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HotelsPostRequest) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCurrency

`func (o *HotelsPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *HotelsPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *HotelsPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *HotelsPostRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGuestNationality

`func (o *HotelsPostRequest) GetGuestNationality() string`

GetGuestNationality returns the GuestNationality field if non-nil, zero value otherwise.

### GetGuestNationalityOk

`func (o *HotelsPostRequest) GetGuestNationalityOk() (*string, bool)`

GetGuestNationalityOk returns a tuple with the GuestNationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNationality

`func (o *HotelsPostRequest) SetGuestNationality(v string)`

SetGuestNationality sets GuestNationality field to given value.

### HasGuestNationality

`func (o *HotelsPostRequest) HasGuestNationality() bool`

HasGuestNationality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


