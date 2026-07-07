package config

import (
	"log"

	"github.com/emarkees/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Snippets  *models.SnippetModel
	DB       *pgxpool.Pool
}

func NewContainer(infoLog *log.Logger, errorLog *log.Logger, db *pgxpool.Pool, snippets *models.SnippetModel) *Container {
	return &Container{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Snippets: snippets,
		DB:       db,
	}
}
