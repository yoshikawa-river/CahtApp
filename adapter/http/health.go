package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SimpleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"string,omitempty"`
}

func SimpleHealthCheck(w http.ResponseWriter, r *http.Request) {
	rs := SimpleResponse{
		Status: http.StatusOK,
	}
	respondJson(w, rs, http.StatusOK)
}

func respondJson(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode response by error '%#v'", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
