package main

import (
	"everywheretew.it/main/security"
	"net/http"
)

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		[]string{"GET", "OPTIONS"},
		"/",
		homePage,
	},
	Route{
		"Login",
		[]string{"GET", "OPTIONS"},
		"/api/login",
		security.Login,
	},
	Route{
		"Get all cluster",
		[]string{"GET", "OPTIONS"},
		"/api/clusters/all",
		getAllCluster,
	},
	Route{
		"Get cluster",
		[]string{"GET", "OPTIONS"},
		"/api/clusters/{id}",
		getCluster,
	},
	Route{
		"Add order data",
		[]string{"POST", "OPTIONS"},
		"/api/order",
		GetOrderData,
	},
	Route{
		"Add product data",
		[]string{"POST", "OPTIONS"},
		"/api/product",
		GetProductData,
	},
	Route{
		"Get user info",
		[]string{"POST", "OPTIONS"},
		"/api/users/{id}",
		GetUserData,
	},
}
