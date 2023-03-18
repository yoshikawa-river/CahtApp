package http

import (
	"net/http"
)

func InitRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/health_check", SimpleHealthCheck)

	return router
}
