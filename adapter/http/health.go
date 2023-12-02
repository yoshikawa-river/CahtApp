package http

import (
	"net/http"

	"github.com/yoshikawa-river/ChatApp/adapter/http/response"
	"github.com/yoshikawa-river/ChatApp/config"
	"github.com/yoshikawa-river/ChatApp/infra"
)

func ServerHealthCheck(w http.ResponseWriter, r *http.Request) {
	rs := response.SimpleResponse{
		Status: http.StatusOK,
	}
	response.RespondJson(w, rs, http.StatusOK)
}

func DBHealthCheck(w http.ResponseWriter, r *http.Request) {
	DBInfo := config.LoadConfig().DBInfo
	_, err := infra.NewConnDB(DBInfo)
	if err != nil {
		// DBへのコネクションにてエラーが発生した場合は503レスポンス
		rs := response.SimpleResponse{
			Status:  http.StatusServiceUnavailable,
			Message: "failed to get connection database",
			Error:   err.Error(),
		}
		response.RespondJson(w, rs, rs.Status)
		return
	}

	rs := response.SimpleResponse{
		Status:  http.StatusOK,
		Message: "success to connect server",
	}
	response.RespondJson(w, rs, rs.Status)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	ServerHealthCheck(w, r)
	DBHealthCheck(w, r)
}
