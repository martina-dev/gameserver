package service

import "net/http"

//Route datastructure with input required by routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes slice of type route
type Routes []Route

var routes = Routes{
	Route{
		"GetScore",
		"GET",
		"/player/{id}",
		GetScore,
	},
}
