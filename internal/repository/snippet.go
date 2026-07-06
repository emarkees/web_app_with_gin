package repository

import (
	"context"
	"time"
)

// Snippet represents a code snippet in the application.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetRepository defines the interface for snippet data access.
// Any storage backend (PostgreSQL, MySQL, in-memory) can implement this.
type SnippetRepository interface {
	Insert(ctx context.Context, title string, content string, expires int) (int, error)
	Get(ctx context.Context, id int) (*Snippet, error)
	Latest(ctx context.Context) ([]*Snippet, error)
}
