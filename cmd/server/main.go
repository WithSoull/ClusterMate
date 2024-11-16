package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"ClusterMate/internal/api/v1"
	"ClusterMate/internal/config"
	"ClusterMate/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var DB *sql.DB
var err error

func main() {
	// Загружаем конфигурацию (например, настройки БД и порты)
	cfg := config.LoadConfig()

	fmt.Println(cfg.GetDSN())
	DB, err = db.OpenDB(cfg.GetDSN())
	if err != nil {
		log.Fatal("DB cant be open cause:", err)
	}

	// Создаем новый роутер
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Инициализируем маршруты
	api.RegisterRoutes(DB, r)

	// Запускаем сервер
	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
