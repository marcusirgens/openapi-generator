/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A StoreApiController binds http requests to an api service and writes the service results to the http response
type StoreApiController struct {
	service StoreApiServicer
}

// NewStoreApiController creates a default api controller
func NewStoreApiController(s StoreApiServicer) Router {
	return &StoreApiController{service: s}
}

// Routes returns all of the api route for the StoreApiController
func (c *StoreApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteOrder",
			strings.ToUpper("Delete"),
			"/v2/store/order/{orderId}",
			c.DeleteOrder,
		},
		{
			"GetInventory",
			strings.ToUpper("Get"),
			"/v2/store/inventory",
			c.GetInventory,
		},
		{
			"GetOrderById",
			strings.ToUpper("Get"),
			"/v2/store/order/{orderId}",
			c.GetOrderById,
		},
		{
			"PlaceOrder",
			strings.ToUpper("Post"),
			"/v2/store/order",
			c.PlaceOrder,
		},
	}
}

// DeleteOrder - Delete purchase order by ID
func (c *StoreApiController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderId := params["orderId"]
	
	result, err := c.service.DeleteOrder(r.Context(), orderId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		er := result
		er.Body = err.Error()
		if er.Code == 0 {
			er.Code = http.StatusInternalServerError
		}
		EncodeJSONImplResponse(w, er)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// GetInventory - Returns pet inventories by status
func (c *StoreApiController) GetInventory(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetInventory(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		er := result
		er.Body = err.Error()
		if er.Code == 0 {
			er.Code = http.StatusInternalServerError
		}
		EncodeJSONImplResponse(w, er)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// GetOrderById - Find purchase order by ID
func (c *StoreApiController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderId, err := parseInt64Parameter(params["orderId"], true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.GetOrderById(r.Context(), orderId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		er := result
		er.Body = err.Error()
		if er.Code == 0 {
			er.Code = http.StatusInternalServerError
		}
		EncodeJSONImplResponse(w, er)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// PlaceOrder - Place an order for a pet
func (c *StoreApiController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	order := &Order{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.PlaceOrder(r.Context(), *order)
	// If an error occurred, encode the error with the status code
	if err != nil {
		er := result
		er.Body = err.Error()
		if er.Code == 0 {
			er.Code = http.StatusInternalServerError
		}
		EncodeJSONImplResponse(w, er)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}
