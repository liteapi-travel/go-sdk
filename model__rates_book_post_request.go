/*
liteAPI

The **liteAPI** can be used to to do the following  Get room rates & availability for a set of hotels Select a specific hotel with room availability and make a booking Manage the bookings - retrieve and cancel existing bookings Get static content for hotels, search hotels by destination

API version: 2.0.0
*/

package liteapi

import (
	"encoding/json"
)

// checks if the BookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookRequest{}

// BookRequest struct for BookRequest
type BookRequest struct {
	// prebook id retrived from prebook response
	PrebookId string               `json:"prebookId"`
	GuestInfo BookRequestGuestInfo `json:"guestInfo"`
	Payment   BookRequestPayment   `json:"payment"`
}

// NewBookRequest instantiates a new BookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookRequest(prebookId string, guestInfo BookRequestGuestInfo, payment BookRequestPayment) *BookRequest {
	this := BookRequest{}
	this.PrebookId = prebookId
	this.GuestInfo = guestInfo
	this.Payment = payment
	return &this
}

// NewBookRequestWithDefaults instantiates a new BookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookRequestWithDefaults() *BookRequest {
	this := BookRequest{}
	return &this
}

// GetPrebookId returns the PrebookId field value
func (o *BookRequest) GetPrebookId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrebookId
}

// GetPrebookIdOk returns a tuple with the PrebookId field value
// and a boolean to check if the value has been set.
func (o *BookRequest) GetPrebookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrebookId, true
}

// SetPrebookId sets field value
func (o *BookRequest) SetPrebookId(v string) {
	o.PrebookId = v
}

// GetGuestInfo returns the GuestInfo field value
func (o *BookRequest) GetGuestInfo() BookRequestGuestInfo {
	if o == nil {
		var ret BookRequestGuestInfo
		return ret
	}

	return o.GuestInfo
}

// GetGuestInfoOk returns a tuple with the GuestInfo field value
// and a boolean to check if the value has been set.
func (o *BookRequest) GetGuestInfoOk() (*BookRequestGuestInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestInfo, true
}

// SetGuestInfo sets field value
func (o *BookRequest) SetGuestInfo(v BookRequestGuestInfo) {
	o.GuestInfo = v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *BookRequest) GetPayment() BookRequestPayment {
	if o == nil || IsNil(o.Payment) {
		var ret BookRequestPayment
		return ret
	}
	return o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookRequest) GetPaymentOk() (*BookRequestPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return &o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *BookRequest) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given BookRequestPayment and assigns it to the Payment field.
func (o *BookRequest) SetPayment(v BookRequestPayment) {
	o.Payment = v
}

func (o BookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prebookId"] = o.PrebookId
	toSerialize["guestInfo"] = o.GuestInfo
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	return toSerialize, nil
}

type NullableBookRequest struct {
	value *BookRequest
	isSet bool
}

func (v NullableBookRequest) Get() *BookRequest {
	return v.value
}

func (v *NullableBookRequest) Set(val *BookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookRequest(val *BookRequest) *NullableBookRequest {
	return &NullableBookRequest{value: val, isSet: true}
}

func (v NullableBookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
