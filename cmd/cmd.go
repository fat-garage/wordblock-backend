package cmd

import (
	"context"
	"github.com/fat-garage/wordblock-backend/pkg/conf"
	"github.com/fat-garage/wordblock-backend/pkg/database"
	"github.com/fat-garage/wordblock-backend/pkg/gipfs"
	"github.com/sirupsen/logrus"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	appName      = "wordblock-backend"
	flagConfig   = "config"
	flagLogLevel = "log_level"
)

var (
	defaultConfigPath = os.ExpandEnv("./config.yaml")
	cfgFile           string
	logLevel          string
)

// NewRootCmd returns the root command.
func NewRootCmd() *cobra.Command {
	// RootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "wordblock",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			conf.LoadConfig(cfgFile)
			logInit(logLevel)
			database.Init()
			gipfs.InitWeb3StorageClient()
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, flagConfig, defaultConfigPath, "Config file path")
	rootCmd.PersistentFlags().StringVar(&logLevel, flagLogLevel, logrus.InfoLevel.String(), "The logging level (trace|debug|info|warn|error|fatal|panic)")

	rootCmd.AddCommand(startApiCmd())

	return rootCmd
}

func Execute() {
	cobra.EnableCommandSorting = false
	rootCmd := NewRootCmd()
	rootCmd.SilenceUsage = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	ctx := context.Background()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Fatalf("execute err: %s", err)
	}
}

func logInit(logLevelStr string) {
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		log.Fatalln(err)
	}
	logrus.SetLevel(logLevel)
	logrus.SetReportCaller(true)
}
