package main

import (
	"log"
	"net/http"

	"ClusterMate/internal/api/v1"
	"ClusterMate/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Загружаем конфигурацию (например, настройки БД и порты)
	cfg := config.LoadConfig()

	// Создаем новый роутер
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Инициализируем маршруты
	api.RegisterRoutes(r)

	// Запускаем сервер
	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
