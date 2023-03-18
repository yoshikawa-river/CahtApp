package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/yoshikawa-river/CahtApp/config"
	"github.com/yoshikawa-river/CahtApp/infla"
)

type SimpleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"string,omitempty"`
}

func ServerHealthCheck(w http.ResponseWriter, r *http.Request) {
	rs := SimpleResponse{
		Status: http.StatusOK,
	}
	respondJson(w, rs, http.StatusOK)
}

func DBHealthCheck(w http.ResponseWriter, r *http.Request) {
	DBInfo := config.LoadConfig().DBInfo
	_, err := infla.NewConnDB(DBInfo)
	if err != nil {
		// DBへのコネクションにてエラーが発生した場合は503レスポンス
		rs := SimpleResponse{
			Status:  http.StatusServiceUnavailable,
			Message: "failed to get connection database",
			Detail:  err.Error(),
		}
		respondJson(w, rs, rs.Status)
		return
	}

	rs := SimpleResponse{
		Status:  http.StatusOK,
		Message: "success to connect server",
	}
	respondJson(w, rs, rs.Status)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	ServerHealthCheck(w, r)
	DBHealthCheck(w, r)
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
