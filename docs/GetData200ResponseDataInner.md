# GetData200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HotelDescription** | Pointer to **string** |  | [optional] 
**HotelImportantInformation** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**GetData200ResponseDataInnerLocation**](GetData200ResponseDataInnerLocation.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**CreditcardRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetData200ResponseDataInner

`func NewGetData200ResponseDataInner() *GetData200ResponseDataInner`

NewGetData200ResponseDataInner instantiates a new GetData200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetData200ResponseDataInnerWithDefaults

`func NewGetData200ResponseDataInnerWithDefaults() *GetData200ResponseDataInner`

NewGetData200ResponseDataInnerWithDefaults instantiates a new GetData200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetData200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetData200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetData200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetData200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetData200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetData200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetData200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetData200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHotelDescription

`func (o *GetData200ResponseDataInner) GetHotelDescription() string`

GetHotelDescription returns the HotelDescription field if non-nil, zero value otherwise.

### GetHotelDescriptionOk

`func (o *GetData200ResponseDataInner) GetHotelDescriptionOk() (*string, bool)`

GetHotelDescriptionOk returns a tuple with the HotelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotelDescription

`func (o *GetData200ResponseDataInner) SetHotelDescription(v string)`

SetHotelDescription sets HotelDescription field to given value.

### HasHotelDescription

`func (o *GetData200ResponseDataInner) HasHotelDescription() bool`

HasHotelDescription returns a boolean if a field has been set.

### GetHotelImportantInformation

`func (o *GetData200ResponseDataInner) GetHotelImportantInformation() string`

GetHotelImportantInformation returns the HotelImportantInformation field if non-nil, zero value otherwise.

### GetHotelImportantInformationOk

`func (o *GetData200ResponseDataInner) GetHotelImportantInformationOk() (*string, bool)`

GetHotelImportantInformationOk returns a tuple with the HotelImportantInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotelImportantInformation

`func (o *GetData200ResponseDataInner) SetHotelImportantInformation(v string)`

SetHotelImportantInformation sets HotelImportantInformation field to given value.

### HasHotelImportantInformation

`func (o *GetData200ResponseDataInner) HasHotelImportantInformation() bool`

HasHotelImportantInformation returns a boolean if a field has been set.

### GetCurrency

`func (o *GetData200ResponseDataInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetData200ResponseDataInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetData200ResponseDataInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetData200ResponseDataInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCountry

`func (o *GetData200ResponseDataInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetData200ResponseDataInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetData200ResponseDataInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetData200ResponseDataInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *GetData200ResponseDataInner) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetData200ResponseDataInner) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetData200ResponseDataInner) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GetData200ResponseDataInner) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetLocation

`func (o *GetData200ResponseDataInner) GetLocation() GetData200ResponseDataInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetData200ResponseDataInner) GetLocationOk() (*GetData200ResponseDataInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetData200ResponseDataInner) SetLocation(v GetData200ResponseDataInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetData200ResponseDataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAddress

`func (o *GetData200ResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetData200ResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetData200ResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetData200ResponseDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetZip

`func (o *GetData200ResponseDataInner) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *GetData200ResponseDataInner) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *GetData200ResponseDataInner) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *GetData200ResponseDataInner) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetIsClosed

`func (o *GetData200ResponseDataInner) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *GetData200ResponseDataInner) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *GetData200ResponseDataInner) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *GetData200ResponseDataInner) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetCreditcardRequired

`func (o *GetData200ResponseDataInner) GetCreditcardRequired() bool`

GetCreditcardRequired returns the CreditcardRequired field if non-nil, zero value otherwise.

### GetCreditcardRequiredOk

`func (o *GetData200ResponseDataInner) GetCreditcardRequiredOk() (*bool, bool)`

GetCreditcardRequiredOk returns a tuple with the CreditcardRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditcardRequired

`func (o *GetData200ResponseDataInner) SetCreditcardRequired(v bool)`

SetCreditcardRequired sets CreditcardRequired field to given value.

### HasCreditcardRequired

`func (o *GetData200ResponseDataInner) HasCreditcardRequired() bool`

HasCreditcardRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


