package handlers

import (
	db "ClusterMate/internal/db/crud"
	"ClusterMate/internal/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Роутер для ролей
func RoleRouter(conn *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Post("/", createRoleHandler(conn))       // Создание роли
	r.Get("/{id}", getRoleHandler(conn))       // Получение роли по ID
	r.Put("/{id}", updateRoleHandler(conn))    // Обновление роли
	r.Delete("/{id}", deleteRoleHandler(conn)) // Удаление роли
	return r
}

// Создание роли
func createRoleHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var role models.Role
		// Декодируем тело запроса в структуру Role
		if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Создаем роль, передаем подключение к базе данных
		lastID, err := db.CreateRole(conn, role) // Передаем conn
		if err != nil {
			log.Printf("Error creating role: %s", err)
			http.Error(w, "Error creating role", http.StatusInternalServerError)
			return
		}

		// Возвращаем статус 201 (создано) и id новой роли
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"id": lastID})
	}
}

// Получение роли по ID
func getRoleHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		role, err := db.GetRoleByID(conn, id)
		if err != nil {
			http.Error(w, "Role not found:", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(role)
	}
}

// Удаление роли
func deleteRoleHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		if err := db.DeleteRole(conn, id); err != nil {
			http.Error(w, "Role not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// Обновление роли
func updateRoleHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем ID роли из URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid role ID", http.StatusBadRequest)
			return
		}

		var role models.Role
		if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Проверяем, что имя роли заполнено
		if role.Name == "" {
			http.Error(w, "Role name is required", http.StatusBadRequest)
			return
		}

		// Вызываем функцию для обновления роли
		err = db.UpdateRole(conn, id, role.Name)

		if err != nil {
			log.Printf("Error: %s", err)
			http.Error(w, "Role not found", http.StatusNotFound)
			return
		}

		// Отправляем успешный ответ
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Role updated successfully"})
	}
}
