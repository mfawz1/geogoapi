package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Host       string `yaml:"host"`
	DBName     string `yaml:"dbName"`
	DBUsername string `yaml:"dbUsername"`
	DBPassword string `yaml:"dbPassword"`
	DBPort     string `yaml:"dbPort"`
}

func loadDatabaseConfig() string {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Config file is missing", err)
	}
	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatal("Error unmarshalling config file", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	return dsn
}
