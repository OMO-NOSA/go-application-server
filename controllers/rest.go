package controllers

import (
	"github.com/OMO-NOSA/go-application-server/config"
)

var globalHTTPServer httpServer

type httpServer struct {}

func newHttpServer() *httpServer{
	return &httpServer{}
}

func InitGlobalHTTPServer() {
	globalHTTPServer = newHttpServer()
}

func runHTTPServer(ctx context.Context) {
		server := &http.Server{
		Addr:             ":" + config.AppConfig.PORT,
		Handler:           nil,
		WriteTImeout:      5 * time.Second,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func InitRouter(ctx context.Context) *mux.Router {
	router := mux.NewRouter()
	// r.Use() middleware 
	return router
}
