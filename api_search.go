/*
liteAPI

The **liteAPI** can be used to to do the following  Get room rates & availability for a set of hotels Select a specific hotel with room availability and make a booking Manage the bookings - retrieve and cancel existing bookings Get static content for hotels, search hotels by destination

API version: 2.0.0
*/

package liteapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// SearchApiService SearchApi service
type SearchApiService service

type ApiGetMinimumRatesRequest struct {
	ctx              context.Context
	ApiService       *SearchApiService
	hotelIds         *string
	checkin          *string
	checkout         *string
	currency         *string
	guestNationality *string
	adults           *int32
	children         *string
	guestId          *string
}

// List of hotelIds
func (r ApiGetMinimumRatesRequest) HotelIds(hotelIds string) ApiGetMinimumRatesRequest {
	r.hotelIds = &hotelIds
	return r
}

// Check in data in YYYY-MM-DD format
func (r ApiGetMinimumRatesRequest) Checkin(checkin string) ApiGetMinimumRatesRequest {
	r.checkin = &checkin
	return r
}

// Check out data in YYYY-MM-DD format
func (r ApiGetMinimumRatesRequest) Checkout(checkout string) ApiGetMinimumRatesRequest {
	r.checkout = &checkout
	return r
}

// Currency code - example (USD)
func (r ApiGetMinimumRatesRequest) Currency(currency string) ApiGetMinimumRatesRequest {
	r.currency = &currency
	return r
}

// Guest nationality ISO-2 code - example (SG)
func (r ApiGetMinimumRatesRequest) GuestNationality(guestNationality string) ApiGetMinimumRatesRequest {
	r.guestNationality = &guestNationality
	return r
}

// Number of adult guests staying
func (r ApiGetMinimumRatesRequest) Adults(adults int32) ApiGetMinimumRatesRequest {
	r.adults = &adults
	return r
}

// Number of children staying if any
func (r ApiGetMinimumRatesRequest) Children(children string) ApiGetMinimumRatesRequest {
	r.children = &children
	return r
}

// Unique traveler ID if available
func (r ApiGetMinimumRatesRequest) GuestId(guestId string) ApiGetMinimumRatesRequest {
	r.guestId = &guestId
	return r
}

func (r ApiGetMinimumRatesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetMinimumRatesExecute(r)
}

/*
GetMinimumRates hotel minimum rates availability

**Hotel Minimum Rates API** is to search and return the minimum room rates that are available for a list of hotel ID's on the specified search dates.

For each hotel ID, the minimum room rate that is available is returned.

The API also has a built in loyalty rewards system. The system rewards return users who have made prior bookings.

If the search is coming from a known guest ID, the guest level is also returned along with pricing has more discounts.

If it is a new user, the guest ID will be generated at the time of the first confirmed booking.

	Example API key for test: sand_c0155ab8-c683-4f26-8f94-b5e92c5797b9

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetMinimumRatesRequest
*/
func (a *SearchApiService) GetMinimumRates(ctx context.Context) ApiGetMinimumRatesRequest {
	return ApiGetMinimumRatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *SearchApiService) GetMinimumRatesExecute(r ApiGetMinimumRatesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.GetMinimumRates")
	if err != nil {
		return localVarReturnValue, nil, &GenericLiteAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hotels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.hotelIds == nil {
		return localVarReturnValue, nil, reportError("hotelIds is required and must be specified")
	}
	if r.checkin == nil {
		return localVarReturnValue, nil, reportError("checkin is required and must be specified")
	}
	if r.checkout == nil {
		return localVarReturnValue, nil, reportError("checkout is required and must be specified")
	}
	if r.currency == nil {
		return localVarReturnValue, nil, reportError("currency is required and must be specified")
	}
	if r.guestNationality == nil {
		return localVarReturnValue, nil, reportError("guestNationality is required and must be specified")
	}
	if r.adults == nil {
		return localVarReturnValue, nil, reportError("adults is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "hotelIds", r.hotelIds, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "checkin", r.checkin, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "checkout", r.checkout, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "guestNationality", r.guestNationality, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "adults", r.adults, "")
	if r.children != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "children", r.children, "")
	}
	if r.guestId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "guestId", r.guestId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericLiteAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericLiteAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFullRatesRequest struct {
	ctx              context.Context
	ApiService       *SearchApiService
	hotelIds         *string
	checkin          *string
	checkout         *string
	currency         *string
	guestNationality *string
	adults           *int32
	children         *string
	guestId          *string
}

// List of hotelIds
func (r ApiGetFullRatesRequest) HotelIds(hotelIds string) ApiGetFullRatesRequest {
	r.hotelIds = &hotelIds
	return r
}

// Check in data in YYYY-MM-DD format
func (r ApiGetFullRatesRequest) Checkin(checkin string) ApiGetFullRatesRequest {
	r.checkin = &checkin
	return r
}

// Check out data in YYYY-MM-DD format
func (r ApiGetFullRatesRequest) Checkout(checkout string) ApiGetFullRatesRequest {
	r.checkout = &checkout
	return r
}

// Guest nationality ISO-2 code - example (SG)
func (r ApiGetFullRatesRequest) GuestNationality(guestNationality string) ApiGetFullRatesRequest {
	r.guestNationality = &guestNationality
	return r
}

// Currency code - example (USD)
func (r ApiGetFullRatesRequest) Currency(currency string) ApiGetFullRatesRequest {
	r.currency = &currency
	return r
}

// Number of adult guests staying
func (r ApiGetFullRatesRequest) Adults(adults int32) ApiGetFullRatesRequest {
	r.adults = &adults
	return r
}

// Number of children staying if any
func (r ApiGetFullRatesRequest) Children(children string) ApiGetFullRatesRequest {
	r.children = &children
	return r
}

// Unique traveler ID if available
func (r ApiGetFullRatesRequest) GuestId(guestId string) ApiGetFullRatesRequest {
	r.guestId = &guestId
	return r
}

func (r ApiGetFullRatesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetFullRatesExecute(r)
}

/*
GetFullRates hotel full rates availability

The Full Rates  API is to search and return all available rooms along with its rates, cancellation policies for a list of hotel ID's based on the search dates.

For each hotel ID, all available room information is returned.

The API also has a built in loyalty rewards system. The system rewards return users who have made prior bookings.

If the search is coming from a known guest ID, the guest level is also returned along with the pricing that's appropriate for the guest level.

If it is a new user, the guest ID will be generated at the time of the first confirmed booking.

	Example API key for test: sand_c0155ab8-c683-4f26-8f94-b5e92c5797b9

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetFullRatesRequest
*/
func (a *SearchApiService) GetFullRates(ctx context.Context) ApiGetFullRatesRequest {
	return ApiGetFullRatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *SearchApiService) GetFullRatesExecute(r ApiGetFullRatesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.GetFullRates")
	if err != nil {
		return localVarReturnValue, nil, &GenericLiteAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hotels/rates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.hotelIds == nil {
		return localVarReturnValue, nil, reportError("hotelIds is required and must be specified")
	}
	if r.checkin == nil {
		return localVarReturnValue, nil, reportError("checkin is required and must be specified")
	}
	if r.checkout == nil {
		return localVarReturnValue, nil, reportError("checkout is required and must be specified")
	}
	if r.guestNationality == nil {
		return localVarReturnValue, nil, reportError("guestNationality is required and must be specified")
	}
	if r.currency == nil {
		return localVarReturnValue, nil, reportError("currency is required and must be specified")
	}
	if r.adults == nil {
		return localVarReturnValue, nil, reportError("adults is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "hotelIds", r.hotelIds, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "checkin", r.checkin, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "checkout", r.checkout, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "guestNationality", r.guestNationality, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "adults", r.adults, "")
	if r.children != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "children", r.children, "")
	}
	if r.guestId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "guestId", r.guestId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericLiteAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericLiteAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
