package cmd

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

var (
	port string
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start API server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting api")
		startHTTP()
	},
}

func init() {
	apiCmd.Flags().StringVarP(&port, "port", "p", "8000", "Port to run the server on")
	rootCmd.AddCommand(apiCmd)
}

func startHTTP() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

// func runHTTPServer(ctx context.Context) {
// 	osChan := make(chan os.Signal, 1)
// 	signal.Notify(osChain, syscall.SIGINT, syscall.SIGTREM)

// 	startCtx, startCancel := context.WithCancel(ctx)
// 	server := &http.Server{
// 		Addr:              fmt.Sprintf(":%s", port),
// 		Handler:           nil,
// 		WriteTImeout:      5 * time.Second,
// 		ReadTimeout:       5 * time.Second,
// 		ReadHeaderTimeout: 5 * time.Second,
// 	}

// 	go func ()  {
// 		<-osChan
// 		fmt.Println("shutting down server")

// 	}

// }

// - Using Cobra
// - Chi Routers (Graceful shutdowns)
// - Pointers, Struct, Interfaces
// - Http Server Spinup
// - https://github.com/spf13/viper
// -
