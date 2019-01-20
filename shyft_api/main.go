package main

//@NOTE SHYFT main func for api, sets up router and spins up a server
//to run server 'go run shyftBlockExplorerApi/*.go'
import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/api"
)

func main() {

	router := NewRouter()
	port := "8080"
	logger.Log("Listening on port " + " " + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

//NewRouter sets up router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range api.Endpoints {

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
