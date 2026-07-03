package config

import "log"


type Container struct {
	ErrorLog *log.Logger
	InfoLog *log.Logger
}

func NewContainer(infoLog *log.Logger, errorLog *log.Logger) *Container {
	return &Container{
		InfoLog: infoLog,
		ErrorLog: errorLog,
	}
}