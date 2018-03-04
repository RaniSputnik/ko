package handle

import (
	"encoding/json"
	"net/http"
)

type resError struct {
	Type    string
	Message string
}

func badRequest(w http.ResponseWriter, err error) {
	send(w, http.StatusBadRequest, resError{
		Type:    "BadRequest",
		Message: err.Error(),
	})
}

var unauthorizedError = resError{
	Type:    "Unauthorized",
	Message: "You must be authenticated to make API requests.",
}

func unauthorized(w http.ResponseWriter) {
	send(w, http.StatusUnauthorized, unauthorizedError)
}

func ok(w http.ResponseWriter, v interface{}) {
	send(w, http.StatusOK, v)
}

func send(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	must(encoder.Encode(v))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
