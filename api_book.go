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

// BookApiService BookApi service
type BookApiService service

type ApiBookRequest struct {
	ctx        context.Context
	ApiService *BookApiService

	prebookId      string
	guestFirstName string
	guestLastName  string
	guestEmail     string
	paymentMethod  string
	holderName     string
	cardNumber     *string
	expireDate     *string
	cvc            *string
	token          *string
}

func (r ApiBookRequest) PrebookId(prebookId string) ApiBookRequest {
	r.prebookId = prebookId
	return r
}

func (r ApiBookRequest) GuestFirstName(guestFirstName string) ApiBookRequest {
	r.guestFirstName = guestFirstName
	return r
}

func (r ApiBookRequest) GuestLastName(guestLastName string) ApiBookRequest {
	r.guestLastName = guestLastName
	return r
}

func (r ApiBookRequest) GuestEmail(guestEmail string) ApiBookRequest {
	r.guestEmail = guestEmail
	return r
}

func (r ApiBookRequest) PaymentMethod(paymentMethod string) ApiBookRequest {
	r.paymentMethod = paymentMethod
	return r
}

func (r ApiBookRequest) HolderName(holderName string) ApiBookRequest {
	r.holderName = holderName
	return r
}

func (r ApiBookRequest) CardNumber(cardNumber string) ApiBookRequest {
	r.cardNumber = &cardNumber
	return r
}

func (r ApiBookRequest) ExpireDate(expireDate string) ApiBookRequest {
	r.expireDate = &expireDate
	return r
}

func (r ApiBookRequest) Cvc(cvc string) ApiBookRequest {
	r.cvc = &cvc
	return r
}

func (r ApiBookRequest) Token(token string) ApiBookRequest {
	r.token = &token
	return r
}

func (r ApiBookRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.BookExecute(r)
}

/*
Book hotel rate book

This API confirms a booking when the prebook Id and the rate Id from the pre book stage along with the guest and payment information are passed.

The guest information is an object that should include the guest first name, last name and email.

The payment information is an object that should include the name, credit card number, expiry and CVC number.

The response will confirm the booking along with a booking Id and a hotel confirmation code. It will also include the booking details including the dates, price and the cancellation policies.

	Example API key for test: sand_c0155ab8-c683-4f26-8f94-b5e92c5797b9

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiBookRequest
*/
func (a *BookApiService) Book(ctx context.Context) ApiBookRequest {
	return ApiBookRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *BookApiService) BookExecute(r ApiBookRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerBookURLWithContext(r.ctx, "BookApiService.Book")
	if err != nil {
		return localVarReturnValue, nil, &GenericLiteAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rates/book"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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

	var bookRequest *BookRequest

	if r.paymentMethod == "CREDIT_CARD" {
		bookRequest = NewBookRequest(r.prebookId, *NewBookRequestGuestInfo(r.guestFirstName, r.guestLastName, r.guestEmail), *NewBookRequestPayment(r.holderName, r.paymentMethod, *r.cardNumber, *r.expireDate, *r.cvc))
	} else if r.paymentMethod == "STRIPE_TOKEN" {
		bookRequest = NewBookRequest(r.prebookId, *NewBookRequestGuestInfo(r.guestFirstName, r.guestLastName, r.guestEmail), *NewBookRequestPaymentWithToken(r.holderName, r.paymentMethod, *r.token))
	}
	// body params
	localVarPostBody = bookRequest
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

type ApiPreBookRequest struct {
	ctx        context.Context
	ApiService *BookApiService
	rateId     *string
}

func (r ApiPreBookRequest) RateId(rateId string) ApiPreBookRequest {
	r.rateId = &rateId
	return r
}

func (r ApiPreBookRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreBookExecute(r)
}

/*
PreBook hotel rate prebook

This API is used to confirm if the room and rates for the search criterion. The input to the endpoint is a specific rate Id coming from the **GET** hotel full rates availability API. In response, the API generates a prebook Id, a new rate Id and contains information if  price, cancellation policy or boarding information has changed. Example API key for test sand_c0155ab8-c683-4f26-8f94-b5e92c5797b9

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPreBookRequest
*/
func (a *BookApiService) PreBook(ctx context.Context) ApiPreBookRequest {
	return ApiPreBookRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *BookApiService) PreBookExecute(r ApiPreBookRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerBookURLWithContext(r.ctx, "BookApiService.PreBook")
	if err != nil {
		return localVarReturnValue, nil, &GenericLiteAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rates/prebook"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = map[string]interface{}{"rateId": r.rateId}
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
