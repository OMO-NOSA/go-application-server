package cmd

import (
	"fmt"
	"github.com/OMO-NOSA/go-application-server/controllers"
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
		controllers.InitGlobalHTTPServer()
	},
}

func init() {
	apiCmd.Flags().StringVarP(&port, "port", "p", "8000", "Port to run the server on")
	rootCmd.AddCommand(apiCmd)
}