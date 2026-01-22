package config

import (
	"log"
	"os"
)

type Config struct {
	Port string

	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string

	StaticToken   string
	SupplierToken string
}

func MustLoad() Config {

	cfg := Config{
		Port: getenv("APP_PORT", "8080"),

		DBHost: getenv("DB_HOST", "localhost"),
		DBPort: getenv("DB_PORT", "5432"),
		DBName: getenv("DB_NAME", "vsm_restaurant"),
		DBUser: getenv("DB_USER", "vsm"),
		DBPass: getenv("DB_PASSWORD", "vsm"),

		StaticToken:   getenv("STATIC_TOKEN", "changeme"),
		SupplierToken: getenv("SUPPLIER_TOKEN", "changeme_supplier"),
	}

	if cfg.Port == "" {
		log.Fatal("APP_PORT is empty")
	}
	return cfg
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
