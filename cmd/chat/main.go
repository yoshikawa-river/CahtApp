package main

import (
	"log"
	"net/http"

	adapterHttp "github.com/yoshikawa-river/CahtApp/adapter/http"
)

func main() {
	router := adapterHttp.InitRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
