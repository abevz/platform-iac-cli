package cli

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize infrastructure workspace",
	Long:  `Scaffolds the directory structure for AWS/Terragrunt environments based on the configuration file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit()
	},
}

func init() {
	// Add init command to the root
	RootCmd.AddCommand(initCmd)
}

func runInit() error {
	slog.Info("Initializing project structure", "project", cfg.Project.Name)

	// Base path from configuration (defaulting to current dir if not specified)
	basePath := "./infrastructure"
	if cfg.Orchestration.Path != "" {
		basePath = cfg.Orchestration.Path
	}

	for _, env := range cfg.Orchestration.Environments {
		envPath := filepath.Join(basePath, env.Name)

		slog.Debug("Creating environment directory", "env", env.Name, "path", envPath)

		// Create directory with 0755 permissions
		if err := os.MkdirAll(envPath, 0o755); err != nil {
			return fmt.Errorf("failed to create directory for %s: %w", env.Name, err)
		}

		// Create a placeholder terragrunt.hcl file
		tgFile := filepath.Join(envPath, "terragrunt.hcl")
		if _, err := os.Stat(tgFile); os.IsNotExist(err) {
			content := fmt.Sprintf("# Terragrunt configuration for %s\n# Account ID: %s\n", env.Name, env.AccountID)
			if err := os.WriteFile(tgFile, []byte(content), 0o644); err != nil {
				slog.Warn("Failed to create terragrunt.hcl", "env", env.Name, "error", err)
			}
		}
	}

	slog.Info("Project initialization complete", "total_environments", len(cfg.Orchestration.Environments))
	return nil
}
