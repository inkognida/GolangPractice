package response

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}

}
