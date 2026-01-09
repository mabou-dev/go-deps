package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mabou-dev/go-deps/pkg/config"
	"github.com/mabou-dev/go-deps/pkg/logger"
)

var (
	cfgFile string
	path    string
	verbose bool
	log     logger.Logger
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-deps",
		Short: "A Go dependency manager",
		Long: `go-deps is a CLI application for managing Go dependencies.
It helps you track and manage project dependencies efficiently`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to go-deps!")
			fmt.Printf("Config file: %s\n", viper.ConfigFileUsed())

			cfg, err := config.LoadConfig()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error loading config:", err)
				os.Exit(1)
			}
			log = logger.NewLogger(cfg.App.Debug)
		},
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config/config.json)")
	cmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbose output")
	viper.BindPFlag("app.debug", cmd.PersistentFlags().Lookup("verbose"))
	cmd.PersistentFlags().StringVar(&path, "path", ".", "path of the project")

	cmd.AddCommand(ListCmd())
	cmd.AddCommand(FetchCmd())

	return cmd
}

// Execute runs the root command
func Execute() {
	cmd := RootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("json")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
