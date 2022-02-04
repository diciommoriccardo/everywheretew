package main

import (
	"everywheretew.it/main/common"
	"everywheretew.it/main/security"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		homePage,
	},
	Route{
		"Login",
		"GET",
		"/api/login",
		security.Login,
	},
	Route{
		"Get all cluster",
		"OPTIONS",
		"/api/clusters/all",
		common.HandleOptions,
	},
	Route{
		"Get all cluster",
		"GET",
		"/api/clusters/all",
		getAllCluster,
	},
	Route{
		"Get cluster",
		"OPTIONS",
		"/api/clusters/{id}",
		getCluster,
	},
	Route{
		"Get cluster",
		"GET",
		"/api/clusters/{id}",
		getCluster,
	},
	Route{
		"Add order data",
		"POST",
		"/api/order",
		GetOrderData,
	},
	Route{
		"Add product data",
		"POST",
		"/api/product",
		GetProductData,
	},
	Route{
		"Get user info",
		"POST",
		"/api/users/{id}",
		GetUserData,
	},
}
