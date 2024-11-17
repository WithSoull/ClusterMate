package db

import (
	"ClusterMate/internal/models"
	"database/sql"
	"fmt"
)

// CreateCluster - добавление нового кластера
func CreateCluster(conn *sql.DB, cluster models.Cluster) (int64, error) {
	result, err := conn.Exec("INSERT INTO clusters (name, description) VALUES (?, ?)", cluster.Name, cluster.Description)
	if err != nil {
		return 0, err
	}
	lastID, _ := result.LastInsertId()
	return lastID, nil
}

// GetClusterByID - получение кластера по ID
func GetClusterByID(conn *sql.DB, id int) (models.Cluster, error) {
	row := conn.QueryRow("SELECT id, name, description FROM clusters WHERE id = ?", id)
	var cluster models.Cluster
	err := row.Scan(&cluster.ID, &cluster.Name, &cluster.Description)
	if err != nil {
		return cluster, err
	}
	return cluster, nil
}

// DeleteCluster - удаление кластера по ID
func DeleteCluster(conn *sql.DB, id int) error {
	_, err := conn.Exec("DELETE FROM clusters WHERE id = ?", id)
	return err
}

// UpdateCluster - обновление кластера по ID
func UpdateCluster(conn *sql.DB, id int, name, description string) error {
	query := `UPDATE clusters 
	          SET name = ?, description = ? 
	          WHERE id = ?`
	_, err := conn.Exec(query, name, description, id)

	if err != nil {
		return fmt.Errorf("Update ERROR: cluster with ID %d: %v", id, err)
	}
	return nil
}
