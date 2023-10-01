package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	appconfig "github.com/OMO-NOSA/go-application-server/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg = &appconfig.AppConfig{}
)

var rootCmd = &cobra.Command{
	Use:   "go-server",
	Short: "go-server cli",
	Long:  `go-server is the root command of application`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := loadConfig()
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func loadConfig() error {
	viper.SetConfigFile(".env")
	err := viper.Unmarshal(cfg)
	return errors.Wrap(err, "unable to load config")
}
