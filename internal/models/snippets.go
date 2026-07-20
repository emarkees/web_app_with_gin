package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// FIX: Define the custom error so it can be verified by your handlers
// var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID        int
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
	VALUES ($1, $2, timezone('utc', now()), timezone('utc', now()) + ($3 * INTERVAL '1 day')) RETURNING id;`

	var id int

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := m.DB.QueryRow(ctx, query, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(ctx context.Context, id int) (*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
		WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	s := &Snippet{}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := m.DB.QueryRow(ctx, query, id).Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
	if err != nil {
		// FIX: Use the declared ErrNoRecord variable here
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrorRecord
		} else {
			return nil, err
		}

	}

	return s, nil
}

// FIX: Fully implement Latest to grab the 10 most recent unexpired snippets
func (m *SnippetModel) Latest(ctx context.Context) ([]Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
		WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	// Always close the row result set to return connections back to the pool
	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	// Check if any errors occurred during iterations
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
