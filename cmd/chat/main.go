package main

import (
	"log"
	"net/http"

	"github.com/yoshikawa-river/CahtApp/adapter"
)

func main() {
	router := adapter.InitRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
