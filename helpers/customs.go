package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Msg string
}

//JSONError return custom JSON error messages
func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
