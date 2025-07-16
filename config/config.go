package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type database struct {
	HostDB string
	PortDB string
	UserDB string
	PassDB string
	NameDB string
}

type Config struct {
	PG database
	PortServer string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PG: database{
			HostDB: envLoad("HOST_DB"),
			PortDB: envLoad("PORT_DB"),
			UserDB: envLoad("USER_DB"),
			PassDB: envLoad("PASS_DB"),
			NameDB: envLoad("NAME_DB"),
		},
		PortServer: envLoad("PORT_SERVER"),
	}
}

func envLoad(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return ""
	}
	return value
}
