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

// ClusterRouter - роутер для кластеров
func ClusterRouter(conn *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Post("/", createClusterHandler(conn))       // Создание кластера
	r.Get("/{id}", getClusterHandler(conn))       // Получение кластера по ID
	r.Put("/{id}", updateClusterHandler(conn))    // Обновление кластера
	r.Delete("/{id}", deleteClusterHandler(conn)) // Удаление кластера
	return r
}

// Создание кластера
func createClusterHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cluster models.Cluster
		// Декодируем тело запроса в структуру Cluster
		if err := json.NewDecoder(r.Body).Decode(&cluster); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Создаем кластер, передаем подключение к базе данных
		lastID, err := db.CreateCluster(conn, cluster) // Передаем conn
		if err != nil {
			log.Printf("Error creating cluster: %s", err)
			http.Error(w, "Error creating cluster", http.StatusInternalServerError)
			return
		}

		// Возвращаем статус 201 (создано) и id нового кластера
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"id": lastID})
	}
}

// Получение кластера по ID
func getClusterHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		cluster, err := db.GetClusterByID(conn, id)
		if err != nil {
			log.Printf("Error: %s", err)
			http.Error(w, "Cluster not found:", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(cluster)
	}
}

// Удаление кластера
func deleteClusterHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		if err := db.DeleteCluster(conn, id); err != nil {
			http.Error(w, "Cluster not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// Обновление кластера
func updateClusterHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем ID кластера из URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid cluster ID", http.StatusBadRequest)
			return
		}

		var cluster models.Cluster
		if err := json.NewDecoder(r.Body).Decode(&cluster); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Проверяем, что имя кластера заполнено
		if cluster.Name == "" {
			http.Error(w, "Cluster name is required", http.StatusBadRequest)
			return
		}

		// Вызываем функцию для обновления кластера
		err = db.UpdateCluster(conn, id, cluster.Name, *cluster.Description)

		if err != nil {
			log.Printf("Error: %s", err)
			http.Error(w, "Cluster not found", http.StatusNotFound)
			return
		}

		// Отправляем успешный ответ
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Cluster updated successfully"})
	}
}
