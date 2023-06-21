/*
Lite API

lite api hotel booking api

API version: 1.0.0
*/

package openapi

import (
	"encoding/json"
)

// HotelsPostRequest struct for HotelsPostRequest
type HotelsPostRequest struct {
	Checkin          *string   `json:"checkin,omitempty"`
	Checkout         *string   `json:"checkout,omitempty"`
	Latitude         *float32  `json:"latitude,omitempty"`
	Longitude        *float32  `json:"longitude,omitempty"`
	Country          *string   `json:"country,omitempty"`
	Distance         *int32    `json:"distance,omitempty"`
	Adults           *int32    `json:"adults,omitempty"`
	Children         *string   `json:"children,omitempty"`
	Currency         *string   `json:"currency,omitempty"`
	GuestNationality *string   `json:"guestNationality,omitempty"`
	HotelIDs         *[]string `json:"hotelIds,omitempty"`
}

// NewHotelsPostRequest instantiates a new HotelsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHotelsPostRequest() *HotelsPostRequest {
	this := HotelsPostRequest{}
	return &this
}

// NewHotelsPostRequestWithDefaults instantiates a new HotelsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHotelsPostRequestWithDefaults() *HotelsPostRequest {
	this := HotelsPostRequest{}
	return &this
}

// GetCheckin returns the Checkin field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetCheckin() string {
	if o == nil || isNil(o.Checkin) {
		var ret string
		return ret
	}
	return *o.Checkin
}

// GetCheckinOk returns a tuple with the Checkin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetCheckinOk() (*string, bool) {
	if o == nil || isNil(o.Checkin) {
		return nil, false
	}
	return o.Checkin, true
}

// HasCheckin returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasCheckin() bool {
	if o != nil && !isNil(o.Checkin) {
		return true
	}

	return false
}

// SetCheckin gets a reference to the given string and assigns it to the Checkin field.
func (o *HotelsPostRequest) SetCheckin(v string) {
	o.Checkin = &v
}

// GetCheckout returns the Checkout field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetCheckout() string {
	if o == nil || isNil(o.Checkout) {
		var ret string
		return ret
	}
	return *o.Checkout
}

// GetCheckoutOk returns a tuple with the Checkout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetCheckoutOk() (*string, bool) {
	if o == nil || isNil(o.Checkout) {
		return nil, false
	}
	return o.Checkout, true
}

// HasCheckout returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasCheckout() bool {
	if o != nil && !isNil(o.Checkout) {
		return true
	}

	return false
}

// SetCheckout gets a reference to the given string and assigns it to the Checkout field.
func (o *HotelsPostRequest) SetCheckout(v string) {
	o.Checkout = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetLatitude() float32 {
	if o == nil || isNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetLatitudeOk() (*float32, bool) {
	if o == nil || isNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasLatitude() bool {
	if o != nil && !isNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *HotelsPostRequest) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetLongitude() float32 {
	if o == nil || isNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetLongitudeOk() (*float32, bool) {
	if o == nil || isNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasLongitude() bool {
	if o != nil && !isNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *HotelsPostRequest) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *HotelsPostRequest) SetCountry(v string) {
	o.Country = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetDistance() int32 {
	if o == nil || isNil(o.Distance) {
		var ret int32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetDistanceOk() (*int32, bool) {
	if o == nil || isNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasDistance() bool {
	if o != nil && !isNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int32 and assigns it to the Distance field.
func (o *HotelsPostRequest) SetDistance(v int32) {
	o.Distance = &v
}

// GetAdults returns the Adults field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetAdults() int32 {
	if o == nil || isNil(o.Adults) {
		var ret int32
		return ret
	}
	return *o.Adults
}

// GetAdultsOk returns a tuple with the Adults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetAdultsOk() (*int32, bool) {
	if o == nil || isNil(o.Adults) {
		return nil, false
	}
	return o.Adults, true
}

// HasAdults returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasAdults() bool {
	if o != nil && !isNil(o.Adults) {
		return true
	}

	return false
}

// SetAdults gets a reference to the given int32 and assigns it to the Adults field.
func (o *HotelsPostRequest) SetAdults(v int32) {
	o.Adults = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetChildren() string {
	if o == nil || isNil(o.Children) {
		var ret string
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetChildrenOk() (*string, bool) {
	if o == nil || isNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasChildren() bool {
	if o != nil && !isNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given string and assigns it to the Children field.
func (o *HotelsPostRequest) SetChildren(v string) {
	o.Children = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetCurrency() string {
	if o == nil || isNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || isNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasCurrency() bool {
	if o != nil && !isNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *HotelsPostRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetGuestNationality returns the GuestNationality field value if set, zero value otherwise.
func (o *HotelsPostRequest) GetGuestNationality() string {
	if o == nil || isNil(o.GuestNationality) {
		var ret string
		return ret
	}
	return *o.GuestNationality
}

// GetGuestNationalityOk returns a tuple with the GuestNationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HotelsPostRequest) GetGuestNationalityOk() (*string, bool) {
	if o == nil || isNil(o.GuestNationality) {
		return nil, false
	}
	return o.GuestNationality, true
}

// HasGuestNationality returns a boolean if a field has been set.
func (o *HotelsPostRequest) HasGuestNationality() bool {
	if o != nil && !isNil(o.GuestNationality) {
		return true
	}

	return false
}

// SetGuestNationality gets a reference to the given string and assigns it to the GuestNationality field.
func (o *HotelsPostRequest) SetGuestNationality(v string) {
	o.GuestNationality = &v
}

func (o HotelsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if !isNil(o.Checkin) {
		toSerialize["checkin"] = o.Checkin
	}
	if !isNil(o.Checkout) {
		toSerialize["checkout"] = o.Checkout
	}
	if !isNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !isNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !isNil(o.Adults) {
		toSerialize["adults"] = o.Adults
	}
	if !isNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !isNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !isNil(o.GuestNationality) {
		toSerialize["guestNationality"] = o.GuestNationality
	}
	if !isNil(o.HotelIDs) {
		toSerialize["hotelIds"] = o.HotelIDs
	}
	return json.Marshal(toSerialize)
}

type NullableHotelsPostRequest struct {
	value *HotelsPostRequest
	isSet bool
}

func (v NullableHotelsPostRequest) Get() *HotelsPostRequest {
	return v.value
}

func (v *NullableHotelsPostRequest) Set(val *HotelsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelsPostRequest(val *HotelsPostRequest) *NullableHotelsPostRequest {
	return &NullableHotelsPostRequest{value: val, isSet: true}
}

func (v NullableHotelsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
