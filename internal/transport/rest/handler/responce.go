package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponce struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func respond(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(payload)
}

func errorRespond(w http.ResponseWriter, r *http.Request, code int, err error) {
	respond(w, r, code, &errorResponce{
		Status:  "error",
		Message: err.Error(),
	})
}

func successRespond(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	respond(w, r, code, payload)
}
