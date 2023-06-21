/*
Lite API

lite api hotel booking api

API version: 1.0.0
*/

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HotelsApiService HotelsApi service
type HotelsApiService service

type ApiGetDataRequest struct {
	ctx         context.Context
	ApiService  *HotelsApiService
	countryCode *string
	hotelName   *string
	cityName    *string
	limit       *float32
	offset      *float32
	latitude    *float32
	longitude   *float32
	distance    *float32
}

// country code Alpha-2 code (example US, RU, CN)
func (r ApiGetDataRequest) CountryCode(countryCode string) ApiGetDataRequest {
	r.countryCode = &countryCode
	return r
}

// hotel name
func (r ApiGetDataRequest) HotelName(hotelName string) ApiGetDataRequest {
	r.hotelName = &hotelName
	return r
}

// city name
func (r ApiGetDataRequest) CityName(cityName string) ApiGetDataRequest {
	r.cityName = &cityName
	return r
}

// limit results (max value 1000)
func (r ApiGetDataRequest) Limit(limit float32) ApiGetDataRequest {
	r.limit = &limit
	return r
}

// results offset
func (r ApiGetDataRequest) Offset(offset float32) ApiGetDataRequest {
	r.offset = &offset
	return r
}

// latitude geo coordinates
func (r ApiGetDataRequest) Latitude(latitude float32) ApiGetDataRequest {
	r.latitude = &latitude
	return r
}

// longtude geo coordinates
func (r ApiGetDataRequest) Longitude(longitude float32) ApiGetDataRequest {
	r.longitude = &longitude
	return r
}

// the distance starting from the selected geopgraphic point
func (r ApiGetDataRequest) Distance(distance float32) ApiGetDataRequest {
	r.distance = &distance
	return r
}

func (r ApiGetDataRequest) Execute() (*GetData200Response, *http.Response, error) {
	return r.ApiService.GetDataExecute(r)
}

/*
GetData Search by Destination/Hotel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDataRequest
*/
func (a *HotelsApiService) GetData(ctx context.Context) ApiGetDataRequest {
	return ApiGetDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetData200Response
func (a *HotelsApiService) GetDataExecute(r ApiGetDataRequest) (*GetData200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *GetData200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HotelsApiService.GetData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.countryCode == nil {
		return localVarReturnValue, nil, ReportError("countryCode is required and must be specified")
	}

	if r.hotelName != nil {
		localVarQueryParams.Add("hotelName", parameterToString(*r.hotelName, ""))
	}
	if r.cityName != nil {
		localVarQueryParams.Add("cityName", parameterToString(*r.cityName, ""))
	}
	localVarQueryParams.Add("countryCode", parameterToString(*r.countryCode, ""))
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.latitude != nil {
		localVarQueryParams.Add("latitude", parameterToString(*r.latitude, ""))
	}
	if r.longitude != nil {
		localVarQueryParams.Add("longitude", parameterToString(*r.longitude, ""))
	}
	if r.distance != nil {
		localVarQueryParams.Add("distance", parameterToString(*r.distance, ""))
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
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHotelsHotelIdGetRequest struct {
	ctx              context.Context
	ApiService       *HotelsApiService
	hotelId          int32
	rid              *string
	xid              *string
	checkin          *string
	checkout         *string
	adults           *int32
	guestNationality *string
	currency         *string
	sessionId        *string
}

func (r ApiHotelsHotelIdGetRequest) Rid(rid string) ApiHotelsHotelIdGetRequest {
	r.rid = &rid
	return r
}

func (r ApiHotelsHotelIdGetRequest) Xid(xid string) ApiHotelsHotelIdGetRequest {
	r.xid = &xid
	return r
}

func (r ApiHotelsHotelIdGetRequest) Checkin(checkin string) ApiHotelsHotelIdGetRequest {
	r.checkin = &checkin
	return r
}

func (r ApiHotelsHotelIdGetRequest) Checkout(checkout string) ApiHotelsHotelIdGetRequest {
	r.checkout = &checkout
	return r
}

func (r ApiHotelsHotelIdGetRequest) Adults(adults int32) ApiHotelsHotelIdGetRequest {
	r.adults = &adults
	return r
}

func (r ApiHotelsHotelIdGetRequest) GuestNationality(guestNationality string) ApiHotelsHotelIdGetRequest {
	r.guestNationality = &guestNationality
	return r
}

func (r ApiHotelsHotelIdGetRequest) Currency(currency string) ApiHotelsHotelIdGetRequest {
	r.currency = &currency
	return r
}

func (r ApiHotelsHotelIdGetRequest) SessionId(sessionId string) ApiHotelsHotelIdGetRequest {
	r.sessionId = &sessionId
	return r
}

func (r ApiHotelsHotelIdGetRequest) Execute() (map[string]any, *http.Response, error) {
	return r.ApiService.HotelsHotelIdGetExecute(r)
}

/*
HotelsHotelIdGet Get Room Availability & Rates for a Hotel ID

This endpoint allows you to send a hotel ID with a specific date range and in response receive all the rooms, rates that are available along with the cancelllation policies.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param hotelId
	@return ApiHotelsHotelIdGetRequest
*/
func (a *HotelsApiService) HotelsHotelIdGet(ctx context.Context, hotelId int32) ApiHotelsHotelIdGetRequest {
	return ApiHotelsHotelIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		hotelId:    hotelId,
	}
}

// Execute executes the request
//
//	@return map[string]any
func (a *HotelsApiService) HotelsHotelIdGetExecute(r ApiHotelsHotelIdGetRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HotelsApiService.HotelsHotelIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hotels/{hotelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"hotelId"+"}", url.PathEscape(parameterToString(r.hotelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.checkin != nil {
		localVarQueryParams.Add("checkin", parameterToString(*r.checkin, ""))
	}
	if r.checkout != nil {
		localVarQueryParams.Add("checkout", parameterToString(*r.checkout, ""))
	}
	if r.adults != nil {
		localVarQueryParams.Add("adults", parameterToString(*r.adults, ""))
	}
	if r.guestNationality != nil {
		localVarQueryParams.Add("guestNationality", parameterToString(*r.guestNationality, ""))
	}
	if r.currency != nil {
		localVarQueryParams.Add("currency", parameterToString(*r.currency, ""))
	}
	if r.sessionId != nil {
		localVarQueryParams.Add("sessionId", parameterToString(*r.sessionId, ""))
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
	if r.rid != nil {
		localVarHeaderParams["rid"] = parameterToString(*r.rid, "")
	}
	if r.xid != nil {
		localVarHeaderParams["xid"] = parameterToString(*r.xid, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]any
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHotelsPostRequest struct {
	ctx               context.Context
	ApiService        *HotelsApiService
	rid               *string
	hotelsPostRequest *HotelsPostRequest
}

func (r ApiHotelsPostRequest) Rid(rid string) ApiHotelsPostRequest {
	r.rid = &rid
	return r
}

func (r ApiHotelsPostRequest) HotelsPostRequest(hotelsPostRequest HotelsPostRequest) ApiHotelsPostRequest {
	r.hotelsPostRequest = &hotelsPostRequest
	return r
}

func (r ApiHotelsPostRequest) Execute() (map[string]any, *http.Response, error) {
	return r.ApiService.HotelsPostExecute(r)
}

/*
HotelsPost Get Minimum Price for Available Hotels

This endpoint allows you to send a list of hotel ID's for a specific date range and in response receive the best rate available for each of the hotel ID's. The limit is set to 200 hotels.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiHotelsPostRequest
*/
func (a *HotelsApiService) HotelsPost(ctx context.Context) ApiHotelsPostRequest {
	return ApiHotelsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]any
func (a *HotelsApiService) HotelsPostExecute(r ApiHotelsPostRequest) (map[string]any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue map[string]any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HotelsApiService.HotelsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hotels"

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
	if r.rid != nil {
		localVarHeaderParams["rid"] = parameterToString(*r.rid, "")
	}
	// body params
	localVarPostBody = r.hotelsPostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-KEY"] = key
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v map[string]any
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
			var v map[string]any
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
