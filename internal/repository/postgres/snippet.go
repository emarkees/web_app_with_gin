package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/emarkees/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SnippetRepo implements repository.SnippetRepository using PostgreSQL.
type SnippetRepo struct {
	Pool *pgxpool.Pool
}

// NewSnippetRepo creates a new PostgreSQL-backed snippet repository.
func NewSnippetRepo(pool *pgxpool.Pool) *SnippetRepo {
	return &SnippetRepo{Pool: pool}
}

// Insert adds a new snippet and returns its ID.
func (r *SnippetRepo) Insert(ctx context.Context, title string, content string, expires int) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
		VALUES ($1, $2, $3, $4) RETURNING id`

	var id int
	err := r.Pool.QueryRow(ctx, query,
		title,
		content,
		time.Now(),
		time.Now().AddDate(0, 0, expires),
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get retrieves a single snippet by its ID.
// Returns repository.ErrNoRecord if no matching snippet is found.
func (r *SnippetRepo) Get(ctx context.Context, id int) (*repository.Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
		WHERE expires > $1 AND id = $2`

	s := &repository.Snippet{}
	err := r.Pool.QueryRow(ctx, query, time.Now(), id).Scan(
		&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrNoRecord
		}
		return nil, err
	}

	return s, nil
}

// Latest returns the 10 most recently created snippets.
func (r *SnippetRepo) Latest(ctx context.Context) ([]*repository.Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
		WHERE expires > $1
		ORDER BY id DESC LIMIT 10`

	rows, err := r.Pool.Query(ctx, query, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*repository.Snippet{}
	for rows.Next() {
		s := &repository.Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
