package config

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	DB       *pgxpool.Pool
}

func NewContainer(infoLog *log.Logger, errorLog *log.Logger, db *pgxpool.Pool) *Container {
	return &Container{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		DB:       db,
	}
}