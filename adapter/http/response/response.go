package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SimpleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Detail  Detail `json:"detail,omitempty"`
}

type Detail struct {
	ID    uint64 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func RespondJson(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode response by error '%#v'", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
