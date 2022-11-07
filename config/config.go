package config

import "fmt"

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func Load() Config {
	config := Config{}
	config.PostgresHost = "localhost"
	config.PostgresPort = 5432
	config.PostgresUser = "postgres"
	config.PostgresPassword = "12345"
	config.PostgresDatabase = "exam"
	return config
}

func ConnStr() string {
	constr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", Load().PostgresUser, Load().PostgresPassword, Load().PostgresDatabase)
	return constr
}
