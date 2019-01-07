package service

import "github.com/gorilla/mux"

//NewRouter returns a pointer to a mux router that we use as a handler
func NewRouter() *mux.Router {
	// Instance of router
	router := mux.NewRouter().StrictSlash(true)

	//Iterarte over the available routes and attache each route using a builder like pattern
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}
