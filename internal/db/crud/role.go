package db

import (
	"ClusterMate/internal/models"
	"database/sql"
	"fmt"
)

// CreateRole - добавление новой роли
func CreateRole(conn *sql.DB, role models.Role) (int64, error) {
	result, err := conn.Exec("INSERT INTO roles (name) VALUES (?)", role.Name)
	if err != nil {
		return 0, err
	}
	lastID, _ := result.LastInsertId()
	return lastID, nil
}

// GetRoleByID - получение роли по ID
func GetRoleByID(conn *sql.DB, id int) (models.Role, error) {
	row := conn.QueryRow("SELECT id, name FROM roles WHERE id = ?", id)
	var role models.Role
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return role, err
	}
	return role, nil
}

// DeleteRole - удаление роли по ID
func DeleteRole(conn *sql.DB, id int) error {
	_, err := conn.Exec("DELETE FROM roles WHERE id = ?", id)
	return err
}

// UpdateRole - обновление роли по ID
func UpdateRole(conn *sql.DB, id int, name string) error {
	query := `UPDATE roles 
	          SET name = ? 
	          WHERE id = ?`
	_, err := conn.Exec(query, name, id)

	if err != nil {
		return fmt.Errorf("Update ERROR: role with ID %d: %v", id, err)
	}
	return nil
}
