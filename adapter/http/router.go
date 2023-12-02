package http

import (
	"database/sql"
	"net/http"

	"github.com/yoshikawa-river/ChatApp/infra"
	"github.com/yoshikawa-river/ChatApp/usecase"
)

func InitRouter(db_info *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/health_check", HealthCheck)
	router.HandleFunc("/server_health_check", ServerHealthCheck)
	router.HandleFunc("/db_health_check", DBHealthCheck)

	userRepository := infra.NewUserRepository(db_info)
	userUsecase := usecase.NewUserInteractor(userRepository)
	userHandler := NewUserHandler(userUsecase)

	router.HandleFunc("/user/create", userHandler.createUser)

	return router
}
