<h1 style="font-weight: 500;">liteAPI SDK</h1>

# Table of Contents
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Installing](#installing)
- [Usage](#usage)
- [Static data](#static-data)
  - [List of cities](#list-of-cities)
  - [List of Countries](#list-of-countries)
  - [List of available currencies](#list-of-available-currencies)
  - [List of hotels](#list-of-hotels)
  - [Hotel details](#hotel-details)
  - [IATA code list](#iata-code-list)
- [Booking flow](#booking-flow)
  - [Search](#search)
    - [Minimum Rates availability](#minimum-rates-availability)
    - [Hotel full rates availability](#hotel-full-rates-availability)
  - [Book](#book)
    - [Hotel rate prebook](#hotel-rate-prebook)
    - [Hotel rate book](#hotel-rate-book)
  - [Booking management](#booking-management)
    - [Booking list](#booking-list)
    - [Booking retrieve](#booking-retrieve)
    - [Booking cancel](#booking-cancel)
- [Guest and loyalty](#guest-and-loyalty)
  - [Guests](#guests)

# Introduction
[liteAPI](https://www.liteapi.travel/) is an innovative and robust collection of infrastructure APIs that cater to the travel industry. It is designed to empower developers, offering them the fastest way to build and launch comprehensive travel applications.

At the heart of LiteAPI's power is its extensive network of over 2 million properties worldwide. By incorporating LiteAPI into an app, developers can effortlessly tap into the vast inventory, allowing users to search and book accommodations at these properties.

But that's not all. With LiteAPI, monetization is made even more simple. Developers can generate revenue through their hospitality products by selling accommodations from LiteAPI's broad portfolio of properties. This means that not only can developers launch their products quickly, they can also start generating profits in no time.

LiteAPI opens up a range of powerful functions for travel and hospitality applications. Its features include:

<h3 style="font-weight: 500; display:inline">Hotel Search:</h3> Developers can incorporate a robust search function that allows users to find hotels based on their preferred destination. This can help users discover accommodations that suit their travel plans.
<br><br>
<h3 style="font-weight: 500; display:inline">Static Content for Hotels:</h3> LiteAPI also provides access to static content for hotels, including descriptions, images, and amenity details. This is essential for developers to present comprehensive and accurate information to the end users, aiding their decision-making process.
<br><br>
<h3 style="font-weight: 500; display:inline">Room Rates & Availability:</h3> One of the most significant features of LiteAPI is the ability to pull data on room rates and availability for a selected set of hotels. This feature ensures users have real-time, accurate information to assist in their booking decisions.
<br><br>
<h3 style="font-weight: 500; display:inline">Hotel Booking:</h3> Beyond just providing information, LiteAPI also allows developers to integrate a seamless booking function. Users can select a specific hotel with room availability and proceed to make a booking directly within the app.
<br><br>
<h3 style="font-weight: 500; display:inline">Booking Management:</h3> With LiteAPI, managing bookings becomes a straightforward task. The booking management functions allow for the tracking and management of all bookings made through the app, ensuring users can keep track of their travel plans.
<br><br>
<h3 style="font-weight: 500; display:inline">Booking Retrieval and Cancellation:</h3> Lastly, LiteAPI offers the capability to retrieve and cancel existing bookings. This added flexibility is crucial for users who might need to alter their travel plans.
<br><br>
All these features make LiteAPI a comprehensive solution for travel app development, offering a plethora of functionalities, from search and booking to management and cancellation. Developers can harness these powerful tools to create high-quality, user-friendly travel applications.
<br><br>

Don't have an account yet?  Sign Up [Here](https://dashboard.liteapi.travel/register/).

# Installing

Make sure your project is using Go Modules (it will have a go.mod file in its root if it already is):

```sh
go mod init
```

Run any of the normal go commands (build/install/test). The Go toolchain will resolve and fetch the liteapi-travel/go-sdk module automatically.

Alternatively, you can also explicitly go get the package into a project:

```sh
go get -u github.com/liteapi-travel/go-sdk
```

Then, reference liteapi in a Go program with import:

```go
import (
	liteapi "github.com/liteapi-travel/go-sdk"
)
```

# Usage

After you have installed the LiteAPI package, you need to configure it with your API key. The API key is available in the [liteAPI Dashboard](https://dashboard.liteapi.travel/apikeys/). Here's the step to add the API key to the package.

```go
	configuration := liteapi.NewConfiguration()
	configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY")
	apiClient := liteapi.NewAPIClient(configuration)
```

# Static data

Static data can be directly fetched from the functions below. Alternatively, LiteAPI also provides an option to download static data directly from the [Github URL](STATICDATA.md).

## List of cities

The GetCitiesByCountryCode function returns a list of city names from a specific country. The country codes must be in ISO-2 format. To get the country codes in ISO-2 for all countries please use the getCountries function. 

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
	countryCode := "US"
	result, res, err := apiClient.StaticDataApi.GetCitiesByCountryCode(context.Background()).CountryCode(countryCode).Execute()
```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>


| Name            | Type       | Description                                | Notes      |
| --------------- | ---------- | ------------------------------------------ | ---------- |
| **countryCode** | **string** | Country code in iso-2 format (example: US) | [required] |

*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array of city objects containing the following properties:

| Field    | Type       | Description           |
| -------- | ---------- | --------------------- |
| **city** | **string** | The name of the city. |

<br>

## List of Countries

The GetCountries function returns the list of countries available along with its ISO-2 code.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
	result, res, err := apiClient.StaticDataApi.GetCountries(context.Background()).Execute()
```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

The function does not need any additional parameter.

*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array of country objects containing the following properties:

| Field    | Type       | Description                       |
| -------- | ---------- | --------------------------------- |
| **code** | **string** | The country code in iso-2 format. |
| **name** | **string** | The name of the country.          |

<br>

## List of available currencies

The GetCurrencies function returns all available currency codes along with its name and the list of supported countries that the currency applies to.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
	result, res, err := apiClient.StaticDataApi.GetCurrencies(context.Background()).Execute()
```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

This function does not need any additional parameters.

*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array of currency objects containing the following properties:

| Name          | Type       | Description                                       |
| ------------- | ---------- | ------------------------------------------------- |
| **code**      | **string** | The currency code.                                |
| **currency**  | **string** | The name of the currency.                         |
| **countries** | **Array**  | An array of countries where the currency is used. |

<br>

## List of hotels

The GetHotels function returns a list of hotels available based on different search criteria. The minimum required information is the country code in ISO-2 format.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
	countryCode := "SG" 
    cityName := "Singapore"

    result, res, err := apiClient.StaticDataApi.GetHotels(context.Background()).CountryCode(countryCode).CityName(cityName).Execute()
```

To utilize optional values, you can invoke the function as follows:

```go
	countryCode := "SG" 
    cityName := "Singapore"
	//Optional values
    offset := int32(0) 
    limit := int32(1000) 
    longitude := float32(-115.16988) 
    latitude := float32(36.12510)
    distance := int32(1000) 
    result, res, err := apiClient.StaticDataApi.GetHotels(context.Background()).CountryCode(countryCode).CityName(cityName).Offset(offset).Limit(limit).Longitude(longitude).Latitude(latitude).Distance(distance).Execute()
```

*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

| Name            | Type       | Description                                                         | Notes      |
| --------------- | ---------- | ------------------------------------------------------------------- | ---------- |
| **countryCode** | **string** | country code ISO-2 code - example (US)                              | [required] |
| **cityName**    | **string** | name of the city                                                    | [required] |
| **offset**      | **number** | specifies the number of rows to skip before starting to return rows | [optional] |
| **limit**       | **number** | limit number of results (max 1000)                                  | [optional] |
| **longitude**   | **string** | longitude geo coordinates                                           | [optional] |
| **latitude**    | **string** | latitude geo coordinates                                            | [optional] |
| **distance**    | **number** | distance in meters (min 1000m)                                      | [optional] |


*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array of hotel objects containing the following properties:  

| Name                 | Type       | Description                                        |
| -------------------- | ---------- | -------------------------------------------------- |
| **id**               | **string** | The unique identifier of the hotel.                |
| **name**             | **string** | The name of the hotel.                             |
| **hotelDescription** | **string** | The description of the hotel.                      |
| **currency**         | **string** | The currency used in the hotel.                    |
| **country**          | **string** | The country code of the hotel.                     |
| **city**             | **string** | The city where the hotel is located.               |
| **latitude**         | **number** | The latitude coordinates of the hotel's location.  |
| **longitude**        | **number** | The longitude coordinates of the hotel's location. |
| **address**          | **string** | The address of the hotel.                          |
| **zip**              | **string** | The postal code of the hotel.                      |
| **main_photo**       | **string** | The URL of the main photo of the hotel.            |
| **stars**            | **number** | The star rating of the hotel.                      |

<br>

## Hotel details

The GetHotelDetails function returns all the static contents details of a hotel or property when given a hotel ID. The static content includes name, description, address, amenities, cancellation policies, images and more.


*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
    hotelId := "lp24373"
	result, res, err := apiClient.StaticDataApi.GetHotelDetails(context.Background()).HotelId(hotelId).Execute()
```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>
| Name        | Type       | Description          | Notes      |
| ----------- | ---------- | -------------------- | ---------- |
| **hotelId** | **string** | Unique ID of a hotel | [required] |

*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

| Name                                                              | Type        | Description                                                                          |
| ----------------------------------------------------------------- | ----------- | ------------------------------------------------------------------------------------ |
| **id**                                                            | **string**  | The unique identifier of the hotel.                                                  |
| **name**                                                          | **string**  | The name of the hotel.                                                               |
| **hotelDescription**                                              | **string**  | The description of the hotel.                                                        |
| **checkinCheckoutTimes**                                          | **Object**  | An object containing the check-in and check-out times of the hotel.                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **checkout**     | **string**  | The check-out time of the hotel.                                                     |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **checkin**      | **string**  | The check-in time of the hotel.                                                      |
| **hotelImages**                                                   | **Array**   | An array of hotel image objects containing the following properties:                 |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **url**          | **string**  | The URL of the hotel image.                                                          |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **thumbnailUrl**       | **string**  | The thumbnail URL of the hotel image.                                                |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **caption**      | **string**  | The caption of the hotel image.                                                      |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **order**        | **string**  | The order of the hotel image (null if not applicable).                               |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **defaultImage** | **boolean** | Indicates whether the hotel image is the default image or not.                       |
| **currency**                                                      | **string**  | The currency used in the hotel.                                                      |
| **country**                                                       | **string**  | The country code of the hotel.                                                       |
| **city**                                                          | **string**  | The city where the hotel is located.                                                 |
| **starRating**                                                    | **number**  | The star rating of the hotel.                                                        |
| **location**                                                      | **Object**  | An object containing the latitude and longitude coordinates of the hotel's location. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **latitude**     | **number**  | The latitude coordinate of the hotel's location.                                     |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **longitude**    | **number**  | The longitude coordinate of the hotel's location.                                    |
| **address**                                                       | **string**  | The address of the hotel.                                                            |
| **zip**                                                           | **string**  | The postal code of the hotel.                                                        |
| **chainId**                                                       | **string**  | The unique identifier of the hotel chain.                                            |
| **hotelFacilities**                                               | **Array**   | An array of hotel facilities offered by the hotel.                                   |

<br>

## IATA code list

The GetIataCodes function returns the IATA (International Air Transport Association) codes for all available airports along with the name of the airport, geographical coordinates and country code in ISO-2 format.


*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
	result, res, err := apiClient.StaticDataApi.GetIataCodes(context.Background()).Execute()
```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

This function does not need any additional parameters.

*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array of IATA objects with the following properties:

| Name            | Type       | Description                            |
| --------------- | ---------- | -------------------------------------- |
| **code**        | **string** | The IATA code.                         |
| **name**        | **string** | The name of the IATA.                  |
| **latitude**    | **number** | The latitude coordinates of the IATA.  |
| **longitude**   | **number** | The longitude coordinates of the IATA. |
| **countryCode** | **string** | The country code of the IATA.          |

<br>

# Booking flow

liteAPI offers a comprehensive and simple way to implement Hotel Booking flow. The booking flow consists of 3 sections: Search, Book, and 
booking management.

<br>

## Search

### Minimum Rates availability
------
The GetMinimumRates function is used to retrieve the minimum room rates that are available for a list of hotel ID's on the specified search dates.

For each hotel ID, the minimum room rate that is available is returned.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
	hotelIds := "lp3803c,lp1f982,lp19b70,lp19e75" 
    checkin := "2023-11-15" 
    checkout := "2023-11-16" 
    currency := "USD" 
    guestNationality := "US" 
    adults := int32(1) 

    result, res, err := apiClient.SearchApi.GetMinimumRates(context.Background()).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).Currency(currency).GuestNationality(guestNationality).Adults(adults).Execute()
```

To utilize optional values, you can invoke the function as follows:
```go
	hotelIds := "lp3803c,lp1f982,lp19b70,lp19e75" 
    checkin := "2023-11-15" 
    checkout := "2023-11-16" 
    currency := "USD" 
    guestNationality := "US" 
    adults := int32(1) 
	//Optional values
    children := "12,9" 
    guestId := "testtraveler1"

    result, res, err := apiClient.SearchApi.GetMinimumRates(context.Background()).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).Currency(currency).GuestNationality(guestNationality).Adults(adults).Children(children).GuestId(guestId).Execute()
```


*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

| Name                 | Type                 | Description                                 | Notes      |
| -------------------- | -------------------- | ------------------------------------------- | ---------- |
| **hotelIds**         | **array of strings** | List of hotelIds                            | [required] |
| **checkin**          | **string**           | Check in data in YYYY-MM-DD format          | [required] |
| **checkout**         | **string**           | Check out data in YYYY-MM-DD format         | [required] |
| **currency**         | **string**           | Currency code - example (USD)               | [required] |
| **guestNationality** | **string**           | Guest nationality ISO-2 code - example (SG) | [required] |
| **adults**           | **number**           | Number of adult guests staying              | [required] |
| **children**         | **array of numbers** | Number of children staying if any           | [optional] |
| **guestId**          | **string**           | Unique traveler ID if available             | [optional] |

*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array of hotel minimum rates objects with the following properties:

| Name           | Type       | Description                      |
| -------------- | ---------- | -------------------------------- |
| **hotelId**    | **string** | The ID of the hotel.             |
| **currency**   | **string** | The currency code for the price. |
| **price**      | **number** | The price of the hotel.          |
| **supplierId** | **number** | The ID of the supplier.          |
| **supplier**   | **string** | The name of the supplier.        |


<br>


### Hotel full rates availability
------
The GetFullRates function return the rates of all available rooms along with its cancellation policies for a list of hotel ID's based on the search dates.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
	hotelIds := "lp3803c,lp1f982,lp19b70,lp19e75" 
    checkin := "2023-11-15" 
    checkout := "2023-11-16" 
    currency := "USD" 
    guestNationality := "US" 
    adults := int32(1) 

    result, res, err := apiClient.SearchApi.GetFullRates(context.Background()).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).Currency(currency).GuestNationality(guestNationality).Adults(adults).Execute()
```

To utilize optional values, you can invoke the function as follows:
```go
	hotelIds := "lp3803c,lp1f982,lp19b70,lp19e75" 
    checkin := "2023-11-15" 
    checkout := "2023-11-16" 
    currency := "USD" 
    guestNationality := "US" 
    adults := int32(1) 
	//Optional values
    children := "12,9" 
    guestId := "testtraveler1"

    result, res, err := apiClient.SearchApi.GetFullRates(context.Background()).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).Currency(currency).GuestNationality(guestNationality).Adults(adults).Children(children).GuestId(guestId).Execute()
```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

| Name                 | Type                 | Description                                 | Notes      |
| -------------------- | -------------------- | ------------------------------------------- | ---------- |
| **hotelIds**         | **array of strings** | List of hotelIds                            | [required] |
| **checkin**          | **string**           | Check in data in YYYY-MM-DD format          | [required] |
| **checkout**         | **string**           | Check out data in YYYY-MM-DD format         | [required] |
| **currency**         | **string**           | Currency code - example (USD)               | [required] |
| **guestNationality** | **string**           | Guest nationality ISO-2 code - example (SG) | [required] |
| **adults**           | **number**           | Number of adult guests staying              | [required] |
| **children**         | **array of numbers** | Number of children staying if any           | [optional] |
| **guestId**          | **string**           | Unique traveler ID if available             | [optional] |

*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array of hotel full rates with the following properties:

| Name                                                                                                              | Type       | Description                                                                                                              |
| ----------------------------------------------------------------------------------------------------------------- | ---------- | ------------------------------------------------------------------------------------------------------------------------ |
| **roomTypeId**                                                                                                    | **string** | The ID of the room type.                                                                                                 |
| **supplier**                                                                                                      | **string** | The name of the supplier.                                                                                                |
| **supplierId**                                                                                                    | **number** | The ID of the supplier.                                                                                                  |
| **rates**                                                                                                         | **Array**  | An array of rate objects containing the pricing and details for each rate within the room type.                          |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **rateId**                                                       | **string** | The ID of the rate.                                                                                                      |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **name**                                                         | **string** | The name of the rate.                                                                                                    |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **maxOccupancy**                                                 | **number** | The maximum occupancy of the room.                                                                                       |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **boardType**                                                    | **string** | The type of board included (e.g., Bed Only).                                                                             |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **boardName**                                                    | **string** | The name of the board (e.g., Bed Only).                                                                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **priceType**                                                    | **string** | The type of pricing (e.g., commission).                                                                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **commission**                                                   | **Array**  | An array of commission objects containing the commission amount and currency.                                            |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **retailRate**                                                   | **Object** | An object containing the retail rate information, including total price, MSP (Marked Selling Price), and taxes and fees. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **total**        | **Array**  | An array of total price objects containing the amount and currency.                                                      |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **msp**          | **Array**  | An array of MSP (Marked Selling Price) objects containing the amount and currency.                                       |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **taxesAndFees** | **Array**  | An array of taxes and fees objects containing information about included or additional charges.                          |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **cancellationPolicies**                                         | **Object** | An object containing cancellation policy information.                                                                    |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **cancelPolicyInfos**                    | **Array**  | An array of cancellation policy info objects containing information about cancellation time, amount, currency, and type. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **hotelRemarks**                         | **Array**  | An array of hotel remarks.                                                                                               |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **refundableTag**                        | **string** | The tag indicating if the rate is refundable or non-refundable.                                                          |


<br>

## Book

### Hotel rate prebook
------

The PreBook function is used to confirm if the room rates are still available before a booking function can be called. The input to the function is a specific rate Id coming from the getFullRates function.
The function returns a prebook Id, a new rate Id and also contains information if the price, cancellation policy or boarding information changed.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
     rateId := "NRYDCZRZHAZHYMRQGIZS2MBXFUYTK7BSGAZDGLJQG4WTCNT4GJ6HYVKTPRDVSWSEJVMVUV2HIUZUOS2OKJKEOWKZKRCU2QSXJVETGRCTJZKEMR2ZGNMFSTKKIRDUSWKEIVGVUUKHGRMVISZXIJJUOQK2IRDU2QSYI5CTGSCZLJGE6TBVJNLEON2DKZFU4NSGJNKTERKQKFNEKQ2NINCFAUK2IRCU6QSUKBJEWVCFKZJVST2KGZCEGTSSLFEEKWKEINHFUV2HKUZVIUKOJJNE6QSELBHVKQSEI5CVURCTJZBFER2BKJKEOTKKJZDUKWSEKNHEEUSHII3EMRKSKNHVAUK2IRAU2USUI5ATGVCDJVJFCR2FLFCECN2CKRIFCWKUJFHEEVCHLFJFMRKULJCEWSSEIU2ESWSTI5AVURCHJRFFCRZUK5KEGTKRPRKVGRD4PR6DCNRWFYYDC7BSGAZDGLJQG4WTCMT4IJHXYMJSHE2DCMD4GI";
    result, res, err := apiClient.BookApi.PreBook(context.Background()).RateId(rateId).Execute()
```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

| Name       | Type       | Description                                  | Notes      |
| ---------- | ---------- | -------------------------------------------- | ---------- |
| **rateId** | **string** | rate id retrieved from getFullRates response | [required] |

*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An object containing prebook information and room type details.                      

| Name                                                                                                                                                                         | Type        | Description                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- | ------------------------------------------------------------------------------------------------------------------------ |
| **prebookId**                                                                                                                                                                | **string**  | The ID of the prebook.                                                                                                   |
| **hotelId**                                                                                                                                                                  | **string**  | The ID of the hotel.                                                                                                     |
| **currency**                                                                                                                                                                 | **string**  | The currency used for pricing.                                                                                           |
| **termsAndConditions**                                                                                                                                                       | **string**  | The terms and conditions (if available).                                                                                 |
| **roomTypes**                                                                                                                                                                | **Array**   | An array of room type objects containing the following properties:                                                       |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **rates**                                                                                                                   | **Array**   | An array of rate objects containing pricing and details for each rate within the room type.                              |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **rateId**                                                                  | **string**  | The ID of the rate.                                                                                                      |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **name**                                                                    | **string**  | The name of the rate.                                                                                                    |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **maxOccupancy**                                                            | **number**  | The maximum occupancy of the room.                                                                                       |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **boardType**                                                               | **string**  | The type of board included (e.g., Bed Only).                                                                             |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **boardName**                                                               | **string**  | The name of the board (e.g., Bed Only).                                                                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **priceType**                                                               | **string**  | The type of pricing (e.g., commission).                                                                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **commission**                                                              | **Object**  | An array of commission objects containing the commission amount and currency.                                            |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **retailRate**                                                              | **Object**  | An object containing the retail rate information, including total price, MSP (Marked Selling Price), and taxes and fees. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **total**             | **Object**  | An array of total price objects containing the amount and currency.                                                      |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **msp**               | **Object**  | An array of MSP (Marked Selling Price) objects containing the amount and currency.                                       |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **taxesAndFees**      | **Object**  | An array of taxes and fees objects containing information about included or additional charges.                          |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **cancellationPolicies**                                                    | **Object**  | An object containing cancellation policy information.                                                                    |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **cancelPolicyInfos** | **Object**  | An array of cancellation policy info objects containing information about cancellation time, amount, and type.           |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **hotelRemarks**      | **Array**   | An array of hotel remarks.                                                                                               |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **refundableTag**     | **string**  | The tag indicating if the rate is refundable or non-refundable.                                                          |
| **msp**                                                                                                                                                                      | **number**  | The Marked Selling Price (MSP) for the prebook.                                                                          |
| **commission**                                                                                                                                                               | **number**  | The commission amount for the prebook.                                                                                   |
| **price**                                                                                                                                                                    | **number**  | The price of the prebook.                                                                                                |
| **priceType**                                                                                                                                                                | **string**  | The type of pricing (e.g., commission).                                                                                  |
| **priceDifferencePercent**                                                                                                                                                   | **number**  | The percentage difference between the retail rate and the Marked Selling Price (MSP).                                    |
| **cancellationChanged**                                                                                                                                                      | **boolean** | Indicates if there were changes to the cancellation policy.                                                              |
| **boardChanged**                                                                                                                                                             | **boolean** | Indicates if there were changes to the board type.                                                                       |
| **supplier**                                                                                                                                                                 | **string**  | The name of the supplier.                                                                                                |
| **supplierId**                                                                                                                                                               | **number**  | The ID of the supplier.                                                                                                  |


<br>

### Hotel rate book
------

The book function confirms a booking when the prebook Id and the rate Id from the pre book stage along with the guest and payment information are passeed.

The guest information is an object that should include the guest first name, last name and email.

The payment information is an object that should include the name, credit card number, expiry and CVC number.

The response will confirm the booking along with a booking Id and a hotel confirmation code. It will also include the booking details including the dates, price and the cancellation policies.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>

If you wish to utilize a credit card, simply invoke the function as follows:
```go
    prebookId := "8iaO7PXBU";
    guestFirstName := "Kim";
    guestLastName := "James"; 
    guestEmail := "test@nlite.ml";
    paymentMethod := "CREDIT_CARD";
    holderName := "Kim James";
    cardNumber := "4242424242424242",;
    expireDate =: "11/23";
    cvc := "123";
    token := nil

    result, res, err := apiClient.BookApi.Book(context.Background()).PrebookId(prebookId).GuestFirstName(guestFirstName).GuestLastName(guestLastName).GuestEmail(guestEmail).PaymentMethod(paymentMethod).HolderName(holderName).CardNumber(cardNumber).ExpireDate(expireDate).Cvc(cvc).Execute()

```

If you prefer to use a Stripe token, you can invoke the function in the following manner:

```go
    prebookId := "8iaO7PXBU";
    guestFirstName := "Kim";
    guestLastName := "James"; 
    guestEmail := "test@nlite.ml";
    paymentMethod := "STRIPE_TOKEN";
    holderName := "Kim James";
    token := "G4WTCNT4GJ6HYVKTPRDVSWSEJVMVUV2"

    result, res, err := apiClient.BookApi.Book(context.Background()).PrebookId(prebookId).GuestFirstName(guestFirstName).GuestLastName(guestLastName).GuestEmail(guestEmail).PaymentMethod(paymentMethod).HolderName(holderName).Token(token).Execute()

```

*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

| Name                                                                | Type       | Description                                                      | Notes                                       |
| ------------------------------------------------------------------- | ---------- | ---------------------------------------------------------------- | ------------------------------------------- |
| **prebookId**                                                       | **string** | prebook id retrieved from prebook response                       | [required]                                  |
| **guestInfo**                                                       | **object** | traveler informations                                            | [required]                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **guestFirstName** | **string** | traveler first name                                              | [required]                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **guestLastName**  | **string** | traveler last name                                               | [required]                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **guestEmail**     | **string** | traveler email                                                   | [required]                                  |
| **paymentMethod**                                                   | **string** | methodEnum: "CREDIT_CARD" or "STRIPE_TOKEN"                      | [required]                                  |
| **holderName**                                                      | **string** | name of the cardholder                                           | [required]                                  |
| **paymentInfo**                                                     | **object** | payment informations                                             | [required]                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **card_number**    | **string** | the card number associated with the credit card                  | [required if paymentMethod is CREDIT_CARD]  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **exp_month**      | **number** | the expiration month of the credit card                          | [required if paymentMethod is CREDIT_CARD]  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **exp_year**       | **number** | the expiration year of the credit card                           | [required if paymentMethod is CREDIT_CARD]  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **cvc**            | **number** | the card verification code (CVC) associated with the credit card | [required if paymentMethod is CREDIT_CARD]  |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **token**          | **string** | token provided by Stripe for the payment method.                 | [required if paymentMethod is STRIPE_TOKEN] |


*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An object containing booking information and room details.   


| Name                                                                                                  | Type       | Description                                                                                                    |
| ----------------------------------------------------------------------------------------------------- | ---------- | -------------------------------------------------------------------------------------------------------------- |
| **bookingId**                                                                                         | **string** | The ID of the booking.                                                                                         |
| **clientReference**                                                                                   | **string** | The client reference.                                                                                          |
| **supplierBookingId**                                                                                 | **string** | The supplier booking ID.                                                                                       |
| **supplierBookingName**                                                                               | **string** | The supplier booking name.                                                                                     |
| **supplier**                                                                                          | **string** | The supplier.                                                                                                  |
| **supplierId**                                                                                        | **number** | The ID of the supplier.                                                                                        |
| **status**                                                                                            | **string** | The status of the booking.                                                                                     |
| **hotelConfirmationCode**                                                                             | **string** | The hotel confirmation code.                                                                                   |
| **checkin**                                                                                           | **string** | The check-in date.                                                                                             |
| **checkout**                                                                                          | **string** | The check-out date.                                                                                            |
| **hotel**                                                                                             | **object** | An object containing hotel details.                                                                            |
| **bookedRooms**                                                                                       | **array**  | An array of booked room objects.                                                                               |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **roomType**                                               | **object** | An object containing room type details.                                                                        |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **adults**                                                 | **number** | The number of adults.                                                                                          |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **children**                                               | **number** | The number of children.                                                                                        |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **rate**                                                   | **object** | An object containing rate details.                                                                             |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **maxOccupancy** | **number** | The maximum occupancy.                                                                                         |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **retailRate**   | **object** | An object containing the retail rate information, including total price.                                       |
| **guestInfo**                                                                                         | **object** | An object containing guest details.                                                                            |
| **createdAt**                                                                                         | **string** | The creation date of the booking.                                                                              |
| **cancellationPolicies**                                                                              | **object** | An object containing cancellation policies information.                                                        |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **cancelPolicyInfos**                                      | **Object** | An array of cancellation policy info objects containing information about cancellation time, amount, and type. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **hotelRemarks**                                           | **Array**  | An array of hotel remarks.                                                                                     |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **refundableTag**                                          | **string** | The tag indicating if the rate is refundable or non-refundable.                                                |
| **price**                                                                                             | **number** | The price of the booking.                                                                                      |
| **msp**                                                                                               | **number** | The MSP (Merchant Service Provider) price.                                                                     |
| **commission**                                                                                        | **number** | The commission amount.                                                                                         |
| **currency**                                                                                          | **string** | The currency of the booking.                                                                                   |


<br>

## Booking management

### Booking list
------
The GetBookingListByGuestId function returns the list of all booking Id's for a given guest Id.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
    guestId := "FrT56hfty" 
    result, res, err  := apiClient.BookingManagementApi.GetBookingListByGuestId(context.Background()).GuestId(guestId).Execute()

```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>


| Name        | Type       | Description              | Notes      |
| ----------- | ---------- | ------------------------ | ---------- |
| **guestId** | **string** | The Guest Id of the user | [required] |


*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array containing objects with the following properties:

| Name          | Type       | Description     |
| ------------- | ---------- | --------------- |
| **bookingId** | **string** | The booking ID. |


<br>

### Booking retrieve
------
The RetrievedBooking function returns the status and the details of a specific booking Id.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
    bookingId := " uSQ6Zsc5R"
    result, res, err := apiClient.BookingManagementApi.RetrievedBooking(context.Background(), bookingId).Execute()

```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>


| Name          | Type       | Description                               | Notes      |
| ------------- | ---------- | ----------------------------------------- | ---------- |
| **bookingId** | **string** | The Booking Id that needs to be retrieved | [required] |


*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An object containing booking information and room details.   

| Name                                                                                                  | Type       | Description                                                                                                    |
| ----------------------------------------------------------------------------------------------------- | ---------- | -------------------------------------------------------------------------------------------------------------- |
| **bookingId**                                                                                         | **string** | The booking ID.                                                                                                |
| **clientReference**                                                                                   | **string** | The client reference.                                                                                          |
| **status**                                                                                            | **string** | The booking status.                                                                                            |
| **hotelConfirmationCode**                                                                             | **string** | The hotel confirmation code.                                                                                   |
| **checkin**                                                                                           | **string** | The check-in date.                                                                                             |
| **checkout**                                                                                          | **string** | The check-out date.                                                                                            |
| **hotel**                                                                                             | **object** | An object containing hotel details.                                                                            |
| **bookedRooms**                                                                                       | **array**  | An array of booked room objects.                                                                               |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **roomType**                                               | **object** | An object containing room type details.                                                                        |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **adults**                                                 | **number** | The number of adults.                                                                                          |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **children**                                               | **number** | The number of children.                                                                                        |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **rate**                                                   | **object** | An object containing rate details.                                                                             |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **maxOccupancy** | **number** | The maximum occupancy.                                                                                         |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **boardType**    | **string** | The board type.                                                                                                |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **boardName**    | **string** | The board name.                                                                                                |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **retailRate**   | **object** | An object containing the retail rate information, including total price.                                       |
| **guestInfo**                                                                                         | **object** | An object containing guest information.                                                                        |
| **createdAt**                                                                                         | **string** | The creation date of the booking.                                                                              |
| **cancellationPolicies**                                                                              | **object** | An object containing cancellation policy details.                                                              |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **cancelPolicyInfos**                                      | **Object** | An array of cancellation policy info objects containing information about cancellation time, amount, and type. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **hotelRemarks**                                           | **Array**  | An array of hotel remarks.                                                                                     |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **refundableTag**                                          | **string** | The tag indicating if the rate is refundable or non-refundable.                                                |
| **currency**                                                                                          | **string** | The currency code.                                                                                             |
| **price**                                                                                             | **number** | The price of the booking.                                                                                      |



<br>

### Booking cancel
------

The CancelBooking function is used to request a cancellation of an existing confirmed booking. Cancellation policies and conditions will be used to determine the success of the cancellation. For example a booking with non-refundable (NRFN) tag or a booking with a cancellation policy that was requested past the cancellation date will not be able to cancel the confirmed booking.


*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
    bookingId := "K8Hvb-85O"
    result, res, err := apiClient.BookingManagementApi.CancelBooking(context.Background(), bookingId).Execute()
```
*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

| Name          | Type       | Description                                             | Notes      |
| ------------- | ---------- | ------------------------------------------------------- | ---------- |
| **bookingId** | **string** | The booking Id of the booking you would like to cancel. | [required] |



*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

| Name                 | Type       | Description                  |
| -------------------- | ---------- | ---------------------------- |
| **bookingId**        | **string** | The booking ID.              |
| **status**           | **string** | The booking status.          |
| **cancellation_fee** | **number** | The cancellation fee.        |
| **refund_amount**    | **number** | The refund amount.           |
| **currency**         | **string** | The currency of the booking. |

<br>

# Guest and loyalty

## Guests

The GetGuestsIds function returns the unique guestID when an email is passed as a parameter. The email that is passed to the book function is used to create a guestID. This guestID can be used to build loyalty programs in the application.

*  <h4 style="color:#9155fd; font-weight: 800;"> Example :</h4>
```go
    result, res, err := apiClient.GuestAndLoyaltyApi.GetGuestsIds(context.Background()).Execute()
```

If you want to retrieve the guest IDs of a specific user based on their email, you can provide the email as an optional parameter:

```go
    email := "johndoe@nlite.ml"
    result, res, err := apiClient.GuestAndLoyaltyApi.GetGuestsIds(context.Background()).Email(email).Execute()
```

*  <h4 style="color:#9155fd; font-weight: 800;"> Parameters :</h4>

| Name      | Type       | Description     | Notes      |
| --------- | ---------- | --------------- | ---------- |
| **email** | **string** | the guest Email | [optional] |


*  <h4 style="color:#9155fd; font-weight: 800;"> Return type :</h4>

An array containing objects with the following properties:

| Name        | Type       | Description   |
| ----------- | ---------- | ------------- |
| **guestId** | **string** | The guest ID. |
