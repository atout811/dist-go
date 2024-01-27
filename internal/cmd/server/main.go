package main

import (
	"log"

	"github.com/atout811/dist-go/internal/server"
)

func main() {

	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
