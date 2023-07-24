package main

import (
	"log"
	"net/http"

	"github.com/omo-nosa/go-application-server/server"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
