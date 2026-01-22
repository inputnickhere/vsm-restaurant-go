package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/inputnickhere/vsm-restaurant-go/internal/config"
	"github.com/inputnickhere/vsm-restaurant-go/internal/db"
	apphttp "github.com/inputnickhere/vsm-restaurant-go/internal/http"
)

func main() {
	// Локально читаем .env в переменные окружения
	_ = godotenv.Load()

	cfg := config.MustLoad()

	pool := db.MustConnect(cfg)
	defer pool.Close()

	handler := apphttp.NewRouter(pool, cfg.StaticToken, cfg.SupplierToken)

	addr := ":" + cfg.Port
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
