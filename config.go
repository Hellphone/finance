package main

const (
	defaultServerPort     = "1234"
	defaultPostgresqlPort = "5432"
)

type Config struct {
	Port string
	DB   *DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}
