package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const defaultPath string = "../../configs/dev.yaml"

type Config struct {
	Env        string `yaml:"env" env-default:"dev"`
	Storage    string `yaml:"storage" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustRead(args []string) *Config {
	configPath := defaultPath
	if len(args) == 2 {
		configPath = args[1]
	} else {
		log.Println("Using default config: ", defaultPath)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Failed to find config file: ", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("Failed to read config file: ", err)
	}

	return &cfg
}
