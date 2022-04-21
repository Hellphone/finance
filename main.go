package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	a := &App{}

	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err = a.Init(cfg); err != nil {
		log.Fatal(err)
	}

	a.Run()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	fmt.Println("server stopped")
}

// TODO: write tests (optional)

func readConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port: os.Getenv("PORT"),
		DB: &DBConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			DBName:   os.Getenv("POSTGRES_DB"),
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		},
	}

	if cfg.Port == "" {
		cfg.Port = defaultServerPort
	}

	if cfg.DB.Host == "" {
		return nil, errors.New("unexpected POSTGRES_HOST' env variable")
	}

	if cfg.DB.DBName == "" {
		return nil, errors.New("unexpected POSTGRES_NAME env variable")
	}

	if cfg.DB.Username == "" {
		return nil, errors.New("unexpected POSTGRES_USER env variable")
	}

	if cfg.DB.Password == "" {
		return nil, errors.New("unexpected POSTGRES_PASSWORD env variable")
	}

	if cfg.DB.Port == "" {
		cfg.DB.Port = defaultPostgresqlPort
	}

	return cfg, nil
}
