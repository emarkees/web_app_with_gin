package controller

import "net/http"

type Container struct {
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my Homepage"))
}
