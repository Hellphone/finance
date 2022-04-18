package main

import (
	"errors"
	"log"
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
}

// TODO: handle errors so that they don't force the application to exit
// TODO: create migrations
// TODO: предусмотреть остановку веб-сервера без потери обрабатываемых запросов

// TODO: read config from .env file
func readConfig() (*Config, error) {
	cfg := &Config{
		Port: "8888",
		DB: &DBConfig{
			//Host:     os.Getenv("DB_HOST"),
			//Port:     os.Getenv("DB_PORT"),
			//DBName:   os.Getenv("DB_NAME"),
			//Username: os.Getenv("DB_USERNAME"),
			//Password: os.Getenv("DB_PASSWORD"),
			Host:     "localhost",
			Port:     "5342",
			DBName:   "finance",
			Username: "postgres",
			Password: "asdjk2j",
		},
	}

	if cfg.DB.Host == "" {
		return nil, errors.New("unexpected DB_HOST' env variable")
	}

	if cfg.DB.DBName == "" {
		return nil, errors.New("unexpected DB_NAME env variable")
	}

	return cfg, nil
}
