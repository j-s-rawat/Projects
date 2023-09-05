// delivery/http/router.go
package http

import (
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *UserHandler) *mux.Router {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID).Methods("GET")

	return router
}
