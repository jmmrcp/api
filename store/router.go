package store

import (
	"github.com/gorilla/mux"
	"net/http"
)

var controller = &Controller {
	Repository: Repository {},
}

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken,
	},
	Route {
		"Index",
		"GET",
		"/",
		controller.Index,
  },
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes { 
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc
		
		router.
		 Methods(route.Method).
		 Path(route.Pattern).
		 Name(route.Name).
		 Handler(handler)
  }
return router
}