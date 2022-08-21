package server

import "github.com/gorilla/mux"

func createRouter(handler *handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v1/records", handler.handleCreate).Methods("POST")
	r.HandleFunc("/v1/records/{offset}", handler.handleGet).Methods("GET")

	return r
}
