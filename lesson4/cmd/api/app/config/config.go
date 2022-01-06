package config

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port        int        `yaml:"port" json:"port,omitempty"`
	Loglevel    string     `yaml:"loglevel" json:"loglevel,omitempty"`
	StoragePath string     `yaml:"storage_path" json:"storage_path,omitempty"`
	AuthConfig  AuthConfig `yaml:"auth_config" json:"auth_config"`
}

type AuthConfig struct {
	JWTSecret string        `yaml:"jwt_secret" json:"-"`
	JWTTTL    time.Duration `yaml:"jwt_ttl" json:"jwt_ttl,omitempty"`
}

func (c *Config) ReadFromFile(logger echo.Logger) {
	configPath := flag.String("config", "", "path yo yaml config")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		logger.Fatalf("can't read config: %v", err)
	}

	if err = yaml.Unmarshal(data, c); err != nil {
		logger.Fatalf("can't unmarshal config: %v", err)
	}

	//nolint:errcheck
	jsn, _ := json.Marshal(c)
	logger.Infof("have read config %s", string(jsn))
}
