package main

import (
	"log"
	"net/http"

	adapterHttp "github.com/yoshikawa-river/ChatApp/adapter/http"
	"github.com/yoshikawa-river/ChatApp/config"
	"github.com/yoshikawa-river/ChatApp/infra"
)

func main() {
	appConfig := config.LoadConfig()

	db_info, _ := infra.NewConnDB(appConfig.DBInfo)

	router := adapterHttp.InitRouter(db_info)

	server := http.Server{
		Addr:    ":" + appConfig.HttpInfo.PROTOCOL,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
