package db

import (
	"ClusterMate/internal/models"
	"database/sql"
	"fmt"
)

// CreateUser - добавление нового пользователя
func CreateUser(conn *sql.DB, user models.User) (int64, error) {
	result, err := conn.Exec("INSERT INTO users (name, role_id, email, password) VALUES (?, ?, ?, ?)", user.Name, user.RoleID, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastID, _ := result.LastInsertId()
	return lastID, nil
}

// GetUserByID - получение пользователя по ID
func GetUserByID(conn *sql.DB, id int) (models.User, error) {
	row := conn.QueryRow("SELECT id, name, role_id, email, password FROM users WHERE id = ?", id)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.RoleID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

// DeleteUser - удаление пользователя по ID
func DeleteUser(conn *sql.DB, id int) error {
	_, err := conn.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

// UpdateUser - полное обновление пользователя по ID
func UpdateUser(conn *sql.DB, id int, name string, role_id int, email, password string) error {
	query := `UPDATE users 
	          SET name = ?, role_id = ?, email = ?, password = ? 
	          WHERE id = ?`
	_, err := conn.Exec(query, name, role_id, email, password, id)

	if err != nil {
		return fmt.Errorf("не удалось обновить пользователя с ID %d: %v", id, err)
	}
	return nil
}
