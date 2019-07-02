package helpers

import (
	"encoding/json"
	"net/http"
)

// APIResponse  - render response json output
func APIResponse(w http.ResponseWriter, httpStatus int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(body)
}
