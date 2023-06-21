/*
Lite API

The **Lite API** can be used to to do the following  Get room rates & availability for a set of hotels Select a specific hotel with room availability and make a booking Manage the bookings - retrieve and cancel existing bookings Get static content for hotels, search hotels by destination

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RatesBookPostRequestGuestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatesBookPostRequestGuestInfo{}

// RatesBookPostRequestGuestInfo struct for RatesBookPostRequestGuestInfo
type RatesBookPostRequestGuestInfo struct {
	// traveler first name
	GuestFirstName string `json:"guestFirstName"`
	// traveler last name
	GuestLastName string `json:"guestLastName"`
	// traveler email
	GuestEmail string `json:"guestEmail"`
}

// NewRatesBookPostRequestGuestInfo instantiates a new RatesBookPostRequestGuestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatesBookPostRequestGuestInfo(guestFirstName string, guestLastName string, guestEmail string) *RatesBookPostRequestGuestInfo {
	this := RatesBookPostRequestGuestInfo{}
	this.GuestFirstName = guestFirstName
	this.GuestLastName = guestLastName
	this.GuestEmail = guestEmail
	return &this
}

// NewRatesBookPostRequestGuestInfoWithDefaults instantiates a new RatesBookPostRequestGuestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatesBookPostRequestGuestInfoWithDefaults() *RatesBookPostRequestGuestInfo {
	this := RatesBookPostRequestGuestInfo{}
	return &this
}

// GetGuestFirstName returns the GuestFirstName field value
func (o *RatesBookPostRequestGuestInfo) GetGuestFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestFirstName
}

// GetGuestFirstNameOk returns a tuple with the GuestFirstName field value
// and a boolean to check if the value has been set.
func (o *RatesBookPostRequestGuestInfo) GetGuestFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestFirstName, true
}

// SetGuestFirstName sets field value
func (o *RatesBookPostRequestGuestInfo) SetGuestFirstName(v string) {
	o.GuestFirstName = v
}

// GetGuestLastName returns the GuestLastName field value
func (o *RatesBookPostRequestGuestInfo) GetGuestLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestLastName
}

// GetGuestLastNameOk returns a tuple with the GuestLastName field value
// and a boolean to check if the value has been set.
func (o *RatesBookPostRequestGuestInfo) GetGuestLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestLastName, true
}

// SetGuestLastName sets field value
func (o *RatesBookPostRequestGuestInfo) SetGuestLastName(v string) {
	o.GuestLastName = v
}

// GetGuestEmail returns the GuestEmail field value
func (o *RatesBookPostRequestGuestInfo) GetGuestEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestEmail
}

// GetGuestEmailOk returns a tuple with the GuestEmail field value
// and a boolean to check if the value has been set.
func (o *RatesBookPostRequestGuestInfo) GetGuestEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestEmail, true
}

// SetGuestEmail sets field value
func (o *RatesBookPostRequestGuestInfo) SetGuestEmail(v string) {
	o.GuestEmail = v
}

func (o RatesBookPostRequestGuestInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatesBookPostRequestGuestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guestFirstName"] = o.GuestFirstName
	toSerialize["guestLastName"] = o.GuestLastName
	toSerialize["guestEmail"] = o.GuestEmail
	return toSerialize, nil
}

type NullableRatesBookPostRequestGuestInfo struct {
	value *RatesBookPostRequestGuestInfo
	isSet bool
}

func (v NullableRatesBookPostRequestGuestInfo) Get() *RatesBookPostRequestGuestInfo {
	return v.value
}

func (v *NullableRatesBookPostRequestGuestInfo) Set(val *RatesBookPostRequestGuestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRatesBookPostRequestGuestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRatesBookPostRequestGuestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatesBookPostRequestGuestInfo(val *RatesBookPostRequestGuestInfo) *NullableRatesBookPostRequestGuestInfo {
	return &NullableRatesBookPostRequestGuestInfo{value: val, isSet: true}
}

func (v NullableRatesBookPostRequestGuestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatesBookPostRequestGuestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

