package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(dsn string) (*sql.DB, error) {
	// Открываем соединение с MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}

	// Проверяем, что соединение установлено
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	return db, nil
}
