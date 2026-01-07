package cli

import (
	"log/slog"
	"os"

	"github.com/abevz/platform-iac-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
	cfg     *config.Config // Глобальный объект конфигурации
)

var RootCmd = &cobra.Command{
	Use:   "platform-cli",
	Short: "Advanced IaC Orchestrator",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogging()
		initConfiguration()
	},
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.platform-cli.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable debug logging")
}

func initLogging() {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(handler))
}

func initConfiguration() {
	var err error
	cfg, err = config.LoadConfig(cfgFile)
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		// Если файл не найден, мы можем либо выйти, либо продолжить с дефолтами
		// Для DevSecOps инструмента лучше выйти при ошибке в конфиге
		os.Exit(1)
	}
	slog.Debug("Config successfully initialized", "project", cfg.Project.Name)
}
