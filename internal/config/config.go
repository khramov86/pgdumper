package config

import (
	"log"
	"os"

	"flag"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DBConfig       `yaml:"db"`
	BackupConfig   `yaml:"backup"`
	TelegramConfig `yaml:"telegram"`
}

type DBConfig struct {
	Host string `yaml:"host" env:"DB_HOST"`
	Port string `yaml:"port" env:"DB_PORT"`
	User string `yaml:"user" env:"DB_USER"`
	Pass string `yaml:"password" env:"DB_PASS"`
	Name string `yaml:"name" env:"DB_NAME"`
}
type BackupConfig struct {
	Path   string   `yaml:"path" env:"BACKUP_PATH"`
	DBList []string `yaml:"db_list"`
}
type TelegramConfig struct {
	ChatID   string `yaml:"chat_id" env:"TELEGRAM_CHAT_ID"`
	BotToken string `yaml:"bot_token" env:"TELEGRAM_BOT_TOKEN"`
}

func GetConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	var config Config
	if configPath == "" {
		flag.StringVar(&configPath, "config", "", "Path to config file")
		flag.Parse()
	}
	if configPath == "" {
		configPath = "config.yaml"
	}
	log.Printf("Config path: %v", configPath)
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Error opening config file: %v from path: %v", err, configPath)
	}
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v from path: %v", err, configPath)
	}
	return &config
}
