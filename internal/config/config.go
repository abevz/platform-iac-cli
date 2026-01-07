package config

import (
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

// Config defines the root configuration structure for platform-iac-cli
type Config struct {
	Project       ProjectConfig       `mapstructure:"project"`
	Orchestration OrchestrationConfig `mapstructure:"orchestration"`
	Security      SecurityConfig      `mapstructure:"security"`
}

type ProjectConfig struct {
	Name   string `mapstructure:"name"`
	Owner  string `mapstructure:"owner"`
	Region string `mapstructure:"region"`
}

type OrchestrationConfig struct {
	Tool         string              `mapstructure:"tool"`
	Path         string              `mapstructure:"path"`
	Environments []EnvironmentConfig `mapstructure:"environments"`
}

type EnvironmentConfig struct {
	Name      string `mapstructure:"name"`
	AccountID string `mapstructure:"account_id"`
}

type SecurityConfig struct {
	Ansible AnsibleConfig `mapstructure:"ansible"`
}

type AnsibleConfig struct {
	// ForceFQCN ensures usage of ansible.build.systemd as per user mandate
	ForceFQCN bool `mapstructure:"force_fqcn"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(cfgFile string) (*Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".platform-cli")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("PLATFORM")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	slog.Debug("Configuration loaded successfully", "project", config.Project.Name)
	return &config, nil
}
