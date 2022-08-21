package server

import "net/http"

type handler struct {
	Log *Log
}

type CreateRequest struct {
	Record Record
}

type CreateResponse struct {
	Offset uint64
}

func (h *handler) handleCreate(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) handleGet(w http.ResponseWriter, r *http.Request) {

}
