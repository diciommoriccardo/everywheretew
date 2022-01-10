package main

import "net/http"

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
		"Get all cluster",
		"GET",
		"/api/cluster/all",
		getAllCluster,
	},
	Route{
		"Get cluster",
		"GET",
		"/api/cluster/{id}",
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
		"/api/products",
		GetProductData,
	},
}
