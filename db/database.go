package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// Open открывает базу данных SQLite и применяет миграции.
func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// Настройки пула
	db.SetMaxOpenConns(1)       // SQLite не поддерживает конкурентные записи
	db.SetMaxIdleConns(1)

	if err := initSchema(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("init schema: %w", err)
	}

	return db, nil
}

// initSchema применяет SQL-миграции из migrations/schema.sql.
func initSchema(db *sql.DB) error {
	schema, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		return fmt.Errorf("read migration file: %w", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("exec migration: %w", err)
	}

	return nil
}
