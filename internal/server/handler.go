package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
	req, err := decodeRequest(r.Body, new(CreateRequest))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	offset, err := h.Log.Append(req.Record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encodeResponse(w, &CreateResponse{offset})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type GetResponse struct {
	Record Record
}

func (h *handler) handleGet(w http.ResponseWriter, r *http.Request) {
	offset, err := decodeFromPathInt(r, "offset")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	record, err := h.Log.GetByOffset(uint64(offset))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = encodeResponse(w, &GetResponse{record})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func encodeResponse[T any](w http.ResponseWriter, val T) error {
	return json.NewEncoder(w).Encode(val)
}

func decodeFromPathInt(r *http.Request, name string) (int, error) {
	vars := mux.Vars(r)
	offsetStr, ok := vars[name]
	if !ok {
		return 0, fmt.Errorf("missing path parameter '%s'", name)
	}
	return strconv.Atoi(offsetStr)
}

func decodeRequest[T any](source io.ReadCloser, target T) (T, error) {
	err := json.NewDecoder(source).Decode(target)
	return target, err
}
