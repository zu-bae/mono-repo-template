package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env       string `mapstructure:"env"`
	Port      string `mapstructure:"port"`
	ClientDir string `mapstructure:"client_dir"`
}

func Load() (*Config, error) {
	v := viper.New()

	v.SetDefault("env", "dev")
	v.SetDefault("port", "8080")
	v.SetDefault("client_dir", "")
	v.SetDefault("frontend_dir", "")

	// Explicit binds so real environment variables always win,
	// even if they are empty strings.
	for _, key := range []string{"env", "port", "client_dir", "frontend_dir"} {
		if err := v.BindEnv(key); err != nil {
			return nil, err
		}
	}

	// .env is optional
	if _, err := os.Stat(".env"); err == nil {
		v.SetConfigFile(".env")
		v.SetConfigType("env")
		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	clientDir := strings.TrimSpace(v.GetString("client_dir"))
	if clientDir == "" {
		if frontendDir := strings.TrimSpace(v.GetString("frontend_dir")); frontendDir != "" {
			clientDir = frontendDir
		} else {
			clientDir = "./web/build"
		}
	}
	v.Set("client_dir", clientDir)

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// IsDev accepts both "dev" so ENV=prod
func (c *Config) IsDev() bool {
	e := strings.ToLower(c.Env)
	return e == "dev"
}

func (c *Config) Addr() string {
	return ":" + c.Port
}
