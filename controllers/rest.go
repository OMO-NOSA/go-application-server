package controllers

import (
	"time"
	"context"
	"net/http"
	"github.com/gorilla/mux"

)

var globalHTTPServer httpServer

type httpServer struct {}

func newHttpServer() httpServer{
	return httpServer{}
}

func InitGlobalHTTPServer() {
	globalHTTPServer = newHttpServer()
}

func RunHTTPServer(ctx context.Context) {
	globalHTTPServer = newHttpServer()
	_ = &http.Server{
		Addr:             ":8000",
		Handler:           InitRouter(ctx),
		WriteTimeout:      5 * time.Second,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func InitRouter(ctx context.Context) *mux.Router {
	router := mux.NewRouter()
	// r.Use() middleware 
	return router
}
