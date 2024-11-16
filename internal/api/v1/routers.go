package api

import (
	"database/sql"
	"net/http"

	"ClusterMate/internal/api/v1/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(conn *sql.DB, r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", handlers.UserRouter(conn))
		//  TODO: Реализовать хэнделры
		// r.Mount("/clusters", handlers.ClusterRouter())
		// r.Mount("/roles", handlers.RoleRouter())
	})

	// Проверка доступности сервера
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}
