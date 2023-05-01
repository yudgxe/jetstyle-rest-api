package handler

import (
	"net/http"
)

func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, http.StatusNotFound, &errorResponce{
			Status:  "error",
			Message: "page not found",
		})
	})
}

func MethodNotAllowed() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, http.StatusMethodNotAllowed, &errorResponce{
			Status:  "error",
			Message: "method not allowed",
		})
	})
}
