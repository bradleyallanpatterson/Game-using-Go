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
		Index,
	},
	Route{
		"Billups",
		"GET",
		"/testbillups",
		Billups,
	},
	Route{
		"choice",
		"GET",
		"/choice",
		choice,
	},
	Route{
		"choices",
		"GET",
		"/choices",
		choices,
	},
	Route{
		"play",
		"POST",
		"/play",
		play,
	},
	
}
