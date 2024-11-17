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

func UserRouter(conn *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Post("/", createUserHandler(conn))
	r.Get("/{id}", getUserHandler(conn))
	r.Put("/{id}", updateUserHandler(conn))
	r.Delete("/{id}", deleteUserHandler(conn))
	return r
}

func createUserHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		// Декодируем тело запроса в структуру User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Создаем пользователя, передаем подключение к базе данных
		lastID, err := db.CreateUser(conn, user) // Передаем conn
		if err != nil {
			log.Printf("Error creating user: %s", err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// Возвращаем статус 201 (создано) и id нового пользователя
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"id": lastID})
	}
}
func getUserHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		user, err := db.GetUserByID(conn, id)
		if err != nil {
			http.Error(w, "User not found:", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func deleteUserHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		if err := db.DeleteUser(conn, id); err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
		}
	}
}

func updateUserHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем ID пользователя из URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Проверяем, что все обязательные поля были заполнены
		if user.Name == "" || user.Email == "" || user.ClusterID == 0 || user.Password == "" || user.RoleID == 0 {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// Вызываем функцию для обновления пользователя
		err = db.UpdateUser(conn, id, user.Name, user.RoleID, user.ClusterID, user.Email, user.Password)

		if err != nil {
			log.Printf("Error: %s", err)
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Отправляем успешный ответ
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "User updated successfully"})
	}
}
