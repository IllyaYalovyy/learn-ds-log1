package server

import "net/http"

func New(endpoint string) *http.Server {
	log := new(Log)
	handler := &handler{log}
	router := createRouter(handler)

	return &http.Server{
		Addr:    endpoint,
		Handler: router,
	}
}
