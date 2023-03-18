package main

import (
	"log"
	"net/http"

	adapterHttp "github.com/yoshikawa-river/CahtApp/adapter/http"
	"github.com/yoshikawa-river/CahtApp/config"
	"github.com/yoshikawa-river/CahtApp/infla"
)

func main() {
	appConfig := config.LoadConfig()

	router := adapterHttp.InitRouter()

	infla.NewConnDB(appConfig.DBInfo)

	server := http.Server{
		Addr:    ":" + appConfig.HttpInfo.PROTOCOL,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
