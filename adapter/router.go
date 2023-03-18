package adapter

import (
	"net/http"

	"github.com/yoshikawa-river/CahtApp/health"
)

func InitRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/health_check", health.SimpleHealthCheck)

	return router
}
