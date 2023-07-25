/*
liteAPI

The **liteAPI** can be used to to do the following  Get room rates & availability for a set of hotels Select a specific hotel with room availability and make a booking Manage the bookings - retrieve and cancel existing bookings Get static content for hotels, search hotels by destination

API version: 2.0.0
*/

package liteapi

import (
	"encoding/json"
)

// checks if the BookRequestPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookRequestPayment{}

// BookRequestPayment struct for BookRequestPayment
type BookRequestPayment struct {
	HolderName string  `json:"holderName"`
	Number     *string `json:"number,omitempty"`
	ExpireDate *string `json:"expireDate,omitempty"`
	Cvc        *string `json:"cvc,omitempty"`
	Method     string  `json:"method"`
	Token      *string `json:"token,omitempty"`
}

// NewBookRequestPayment instantiates a new BookRequestPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookRequestPayment(holderName string, method string, number string, expireDate string, cvc string) *BookRequestPayment {
	this := BookRequestPayment{}
	this.HolderName = holderName
	this.Method = method
	this.Number = &number
	this.ExpireDate = &expireDate
	this.Cvc = &cvc
	return &this
}

func NewBookRequestPaymentWithToken(holderName string, method string, token string) *BookRequestPayment {
	this := BookRequestPayment{}
	this.HolderName = holderName
	this.Method = method
	this.Token = &token
	return &this
}

// NewBookRequestPaymentWithDefaults instantiates a new BookRequestPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookRequestPaymentWithDefaults() *BookRequestPayment {
	this := BookRequestPayment{}
	return &this
}

// GetHolderName returns the HolderName field value
func (o *BookRequestPayment) GetHolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value
// and a boolean to check if the value has been set.
func (o *BookRequestPayment) GetHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HolderName, true
}

// SetHolderName sets field value
func (o *BookRequestPayment) SetHolderName(v string) {
	o.HolderName = v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *BookRequestPayment) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookRequestPayment) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *BookRequestPayment) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *BookRequestPayment) SetNumber(v string) {
	o.Number = &v
}

// GetExpireDate returns the ExpireDate field value if set, zero value otherwise.
func (o *BookRequestPayment) GetExpireDate() string {
	if o == nil || IsNil(o.ExpireDate) {
		var ret string
		return ret
	}
	return *o.ExpireDate
}

// GetExpireDateOk returns a tuple with the ExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookRequestPayment) GetExpireDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpireDate) {
		return nil, false
	}
	return o.ExpireDate, true
}

// HasExpireDate returns a boolean if a field has been set.
func (o *BookRequestPayment) HasExpireDate() bool {
	if o != nil && !IsNil(o.ExpireDate) {
		return true
	}

	return false
}

// SetExpireDate gets a reference to the given string and assigns it to the ExpireDate field.
func (o *BookRequestPayment) SetExpireDate(v string) {
	o.ExpireDate = &v
}

// GetCvc returns the Cvc field value if set, zero value otherwise.
func (o *BookRequestPayment) GetCvc() string {
	if o == nil || IsNil(o.Cvc) {
		var ret string
		return ret
	}
	return *o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookRequestPayment) GetCvcOk() (*string, bool) {
	if o == nil || IsNil(o.Cvc) {
		return nil, false
	}
	return o.Cvc, true
}

// HasCvc returns a boolean if a field has been set.
func (o *BookRequestPayment) HasCvc() bool {
	if o != nil && !IsNil(o.Cvc) {
		return true
	}

	return false
}

// SetCvc gets a reference to the given string and assigns it to the Cvc field.
func (o *BookRequestPayment) SetCvc(v string) {
	o.Cvc = &v
}

// GetMethod returns the Method field value
func (o *BookRequestPayment) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *BookRequestPayment) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *BookRequestPayment) SetMethod(v string) {
	o.Method = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *BookRequestPayment) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookRequestPayment) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *BookRequestPayment) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *BookRequestPayment) SetToken(v string) {
	o.Token = &v
}

func (o BookRequestPayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookRequestPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["holderName"] = o.HolderName
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.ExpireDate) {
		toSerialize["expireDate"] = o.ExpireDate
	}
	if !IsNil(o.Cvc) {
		toSerialize["cvc"] = o.Cvc
	}
	toSerialize["method"] = o.Method
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableBookRequestPayment struct {
	value *BookRequestPayment
	isSet bool
}

func (v NullableBookRequestPayment) Get() *BookRequestPayment {
	return v.value
}

func (v *NullableBookRequestPayment) Set(val *BookRequestPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableBookRequestPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableBookRequestPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookRequestPayment(val *BookRequestPayment) *NullableBookRequestPayment {
	return &NullableBookRequestPayment{value: val, isSet: true}
}

func (v NullableBookRequestPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookRequestPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
