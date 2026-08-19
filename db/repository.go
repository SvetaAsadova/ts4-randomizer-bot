package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Repository предоставляет методы доступа к данным.
type Repository struct {
	db *sql.DB
}

// NewRepository создаёт новый репозиторий.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// RandomEvent возвращает любое случайное событие из БД.
func (r *Repository) RandomEvent(ctx context.Context) (*Event, error) {
	query := `SELECT id, category, title, description, variables, dlc, impact, created_at
			  FROM events ORDER BY RANDOM() LIMIT 1`

	row := r.db.QueryRowContext(ctx, query)
	return scanEvent(row)
}

// HasAnyEvents проверяет, есть ли вообще события в БД.
func (r *Repository) HasAnyEvents(ctx context.Context) (bool, error) {
	query := `SELECT COUNT(*) FROM events`
	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("check events: %w", err)
	}
	return count > 0, nil
}

// scanEvent — обобщённая функция сканирования строки в Event.
func scanEvent(row scanner) (*Event, error) {
	var e Event
	var dlc sql.NullString

	err := row.Scan(&e.ID, &e.Category, &e.Title, &e.Description,
		&e.Variables, &dlc, &e.Impact, &e.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("scan event: %w", err)
	}

	e.DLC = sqlNullString{Value: dlc.String, Valid: dlc.Valid}
	return &e, nil
}

// scanner — интерфейс для строк, совместимый с sql.Row и sql.Rows.
type scanner interface {
	Scan(dest ...any) error
}
