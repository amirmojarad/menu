package config

import "github.com/kelseyhightower/envconfig"

var cfg *AppConfig

type AppConfig struct {
	App struct {
		Port  int    `envconfig:"PORT" default:"8080"`
		Host  string `envconfig:"HOST" default:"localhost"`
		Debug bool   `envconfig:"DEBUG" default:"true"`
		Name  string `envconfig:"NAME" default:"menu_generator"`
	} `envconfig:"APP"`

	Database struct {
		MigrationPath string `envconfig:"MIGRATION_PATH"`
		User          string `envconfig:"USERNAME"`
		Name          string `envconfig:"NAME"`
		Port          int    `envconfig:"PORT"`
		Host          string `envconfig:"HOST"`
		Password      string `envconfig:"PASSWORD"`
	} `envconfig:"POSTGRES"`
}

func GetConfig() (*AppConfig, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &AppConfig{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
