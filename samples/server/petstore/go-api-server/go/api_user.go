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

// UserApiController binds http requests to an api service and writes the service results to the http response
type UserApiController struct {
	service UserApiServicer
	errorHandler ErrorHandler
}

// UserApiOption for how the controller is set up.
type UserApiOption func(*UserApiController)

// WithUserApiErrorHandler inject ErrorHandler into controller
func WithUserApiErrorHandler(h ErrorHandler) UserApiOption {
	return func(c *UserApiController) {
		c.errorHandler = h
	}
}

// NewUserApiController creates a default api controller
func NewUserApiController(s UserApiServicer, opts ...UserApiOption) Router {
	controller := &UserApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the UserApiController
func (c *UserApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateUser",
			strings.ToUpper("Post"),
			"/v2/user",
			c.CreateUser,
		},
		{
			"CreateUsersWithArrayInput",
			strings.ToUpper("Post"),
			"/v2/user/createWithArray",
			c.CreateUsersWithArrayInput,
		},
		{
			"CreateUsersWithListInput",
			strings.ToUpper("Post"),
			"/v2/user/createWithList",
			c.CreateUsersWithListInput,
		},
		{
			"DeleteUser",
			strings.ToUpper("Delete"),
			"/v2/user/{username}",
			c.DeleteUser,
		},
		{
			"GetUserByName",
			strings.ToUpper("Get"),
			"/v2/user/{username}",
			c.GetUserByName,
		},
		{
			"LoginUser",
			strings.ToUpper("Get"),
			"/v2/user/login",
			c.LoginUser,
		},
		{
			"LogoutUser",
			strings.ToUpper("Get"),
			"/v2/user/logout",
			c.LogoutUser,
		},
		{
			"UpdateUser",
			strings.ToUpper("Put"),
			"/v2/user/{username}",
			c.UpdateUser,
		},
	}
}

// CreateUser - Create user
func (c *UserApiController) CreateUser(w http.ResponseWriter, r *http.Request) {
	userParam := User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserRequired(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateUser(r.Context(), userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// CreateUsersWithArrayInput - Creates list of users with given input array
func (c *UserApiController) CreateUsersWithArrayInput(w http.ResponseWriter, r *http.Request) {
	userParam := []User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	for _, el := range userParam {
		if err := AssertUserRequired(el); err != nil {
			c.errorHandler(w, r, err, nil)
			return
		}
	}
	result, err := c.service.CreateUsersWithArrayInput(r.Context(), userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// CreateUsersWithListInput - Creates list of users with given input array
func (c *UserApiController) CreateUsersWithListInput(w http.ResponseWriter, r *http.Request) {
	userParam := []User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	for _, el := range userParam {
		if err := AssertUserRequired(el); err != nil {
			c.errorHandler(w, r, err, nil)
			return
		}
	}
	result, err := c.service.CreateUsersWithListInput(r.Context(), userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// DeleteUser - Delete user
func (c *UserApiController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameParam := params["username"]
	
	result, err := c.service.DeleteUser(r.Context(), usernameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// GetUserByName - Get user by user name
func (c *UserApiController) GetUserByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameParam := params["username"]
	
	result, err := c.service.GetUserByName(r.Context(), usernameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// LoginUser - Logs user into the system
func (c *UserApiController) LoginUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	usernameParam := query.Get("username")
	passwordParam := query.Get("password")
	result, err := c.service.LoginUser(r.Context(), usernameParam, passwordParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// LogoutUser - Logs out current logged in user session
func (c *UserApiController) LogoutUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.LogoutUser(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}

// UpdateUser - Updated user
func (c *UserApiController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameParam := params["username"]
	
	userParam := User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserRequired(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUser(r.Context(), usernameParam, userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONImplResponse(w, result)
}
