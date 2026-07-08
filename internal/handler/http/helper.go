package handler

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (c *RouterContext) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	c.App.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (c *RouterContext) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}