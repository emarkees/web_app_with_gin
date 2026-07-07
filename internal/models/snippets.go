package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID       int
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func NewSnippetModel(db *pgxpool.Pool) *SnippetModel {
	if db == nil {
		panic("nil db")
	}
	return &SnippetModel{
		DB: db,
	}
}

func (m *SnippetModel) Insert(ctx context.Context, title string, content string, expires int) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
		VALUES ($1, $2, CURRENT_TIMESTAMP + ($3 * INTERVAL '1 day')) RETURNING id
	`
	var id int

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := m.DB.QueryRow(ctx, query,
		title,
		content,
		expires,

	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}