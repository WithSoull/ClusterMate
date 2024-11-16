package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createUser)
	r.Get("/{id}", getUser)
	r.Put("/{id}", updateUser)
	r.Delete("/{id}", deleteUser)
	return r
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	w.Write([]byte("Get user " + userID))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	w.Write([]byte("Update user " + userID))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	w.Write([]byte("Get user " + userID))
}
