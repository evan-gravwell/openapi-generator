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
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/v2/",
		Index,
	},

	{
		"AddPet",
		http.MethodPost,
		"/v2/pet",
		AddPet,
	},

	{
		"DeletePet",
		http.MethodDelete,
		"/v2/pet/:petId",
		DeletePet,
	},

	{
		"FindPetsByStatus",
		http.MethodGet,
		"/v2/pet/findByStatus",
		FindPetsByStatus,
	},

	{
		"FindPetsByTags",
		http.MethodGet,
		"/v2/pet/findByTags",
		FindPetsByTags,
	},

	{
		"GetPetById",
		http.MethodGet,
		"/v2/pet/:petId",
		GetPetById,
	},

	{
		"UpdatePet",
		http.MethodPut,
		"/v2/pet",
		UpdatePet,
	},

	{
		"UpdatePetWithForm",
		http.MethodPost,
		"/v2/pet/:petId",
		UpdatePetWithForm,
	},

	{
		"UploadFile",
		http.MethodPost,
		"/v2/pet/:petId/uploadImage",
		UploadFile,
	},

	{
		"DeleteOrder",
		http.MethodDelete,
		"/v2/store/order/:orderId",
		DeleteOrder,
	},

	{
		"GetInventory",
		http.MethodGet,
		"/v2/store/inventory",
		GetInventory,
	},

	{
		"GetOrderById",
		http.MethodGet,
		"/v2/store/order/:orderId",
		GetOrderById,
	},

	{
		"PlaceOrder",
		http.MethodPost,
		"/v2/store/order",
		PlaceOrder,
	},

	{
		"CreateUser",
		http.MethodPost,
		"/v2/user",
		CreateUser,
	},

	{
		"CreateUsersWithArrayInput",
		http.MethodPost,
		"/v2/user/createWithArray",
		CreateUsersWithArrayInput,
	},

	{
		"CreateUsersWithListInput",
		http.MethodPost,
		"/v2/user/createWithList",
		CreateUsersWithListInput,
	},

	{
		"DeleteUser",
		http.MethodDelete,
		"/v2/user/:username",
		DeleteUser,
	},

	{
		"GetUserByName",
		http.MethodGet,
		"/v2/user/:username",
		GetUserByName,
	},

	{
		"LoginUser",
		http.MethodGet,
		"/v2/user/login",
		LoginUser,
	},

	{
		"LogoutUser",
		http.MethodGet,
		"/v2/user/logout",
		LogoutUser,
	},

	{
		"UpdateUser",
		http.MethodPut,
		"/v2/user/:username",
		UpdateUser,
	},
}
