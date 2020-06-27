package http_utils

import (
	"encoding/json"
	"net/http"

	errors "github.com/kadekchrisna/openbook-utils-go/rest_errors"
)

func ResponseSuccess(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}
func ResponseError(w http.ResponseWriter, error errors.ResErr) {
	ResponseSuccess(w, error.Status(), error)
}
