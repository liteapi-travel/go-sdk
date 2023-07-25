/*
liteAPI

The **liteAPI** can be used to to do the following  Get room rates & availability for a set of hotels Select a specific hotel with room availability and make a booking Manage the bookings - retrieve and cancel existing bookings Get static content for hotels, search hotels by destination

API version: 2.0.0
*/

package liteapi

import (
	"encoding/json"
)

// checks if the BookRequestGuestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookRequestGuestInfo{}

// BookRequestGuestInfo struct for BookRequestGuestInfo
type BookRequestGuestInfo struct {
	// traveler first name
	GuestFirstName string `json:"guestFirstName"`
	// traveler last name
	GuestLastName string `json:"guestLastName"`
	// traveler email
	GuestEmail string `json:"guestEmail"`
}

// NewBookRequestGuestInfo instantiates a new BookRequestGuestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookRequestGuestInfo(guestFirstName string, guestLastName string, guestEmail string) *BookRequestGuestInfo {
	this := BookRequestGuestInfo{}
	this.GuestFirstName = guestFirstName
	this.GuestLastName = guestLastName
	this.GuestEmail = guestEmail
	return &this
}

// NewBookRequestGuestInfoWithDefaults instantiates a new BookRequestGuestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookRequestGuestInfoWithDefaults() *BookRequestGuestInfo {
	this := BookRequestGuestInfo{}
	return &this
}

// GetGuestFirstName returns the GuestFirstName field value
func (o *BookRequestGuestInfo) GetGuestFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestFirstName
}

// GetGuestFirstNameOk returns a tuple with the GuestFirstName field value
// and a boolean to check if the value has been set.
func (o *BookRequestGuestInfo) GetGuestFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestFirstName, true
}

// SetGuestFirstName sets field value
func (o *BookRequestGuestInfo) SetGuestFirstName(v string) {
	o.GuestFirstName = v
}

// GetGuestLastName returns the GuestLastName field value
func (o *BookRequestGuestInfo) GetGuestLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestLastName
}

// GetGuestLastNameOk returns a tuple with the GuestLastName field value
// and a boolean to check if the value has been set.
func (o *BookRequestGuestInfo) GetGuestLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestLastName, true
}

// SetGuestLastName sets field value
func (o *BookRequestGuestInfo) SetGuestLastName(v string) {
	o.GuestLastName = v
}

// GetGuestEmail returns the GuestEmail field value
func (o *BookRequestGuestInfo) GetGuestEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestEmail
}

// GetGuestEmailOk returns a tuple with the GuestEmail field value
// and a boolean to check if the value has been set.
func (o *BookRequestGuestInfo) GetGuestEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestEmail, true
}

// SetGuestEmail sets field value
func (o *BookRequestGuestInfo) SetGuestEmail(v string) {
	o.GuestEmail = v
}

func (o BookRequestGuestInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookRequestGuestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guestFirstName"] = o.GuestFirstName
	toSerialize["guestLastName"] = o.GuestLastName
	toSerialize["guestEmail"] = o.GuestEmail
	return toSerialize, nil
}

type NullableBookRequestGuestInfo struct {
	value *BookRequestGuestInfo
	isSet bool
}

func (v NullableBookRequestGuestInfo) Get() *BookRequestGuestInfo {
	return v.value
}

func (v *NullableBookRequestGuestInfo) Set(val *BookRequestGuestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBookRequestGuestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBookRequestGuestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookRequestGuestInfo(val *BookRequestGuestInfo) *NullableBookRequestGuestInfo {
	return &NullableBookRequestGuestInfo{value: val, isSet: true}
}

func (v NullableBookRequestGuestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookRequestGuestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
