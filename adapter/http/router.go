package http

import (
	"database/sql"
	"net/http"
)

type MySQLConnector struct {
	Conn *sql.DB
}

func InitRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/health_check", HealthCheck)
	router.HandleFunc("/server_health_check", ServerHealthCheck)
	router.HandleFunc("/db_health_check", DBHealthCheck)

	return router
}
