package main

//@ NOTE Shyft setting up router
import (
	"net/http"

	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
	"github.com/gorilla/mux"
)

//NewRouter sets up router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
