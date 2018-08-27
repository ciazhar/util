package util

import (
	"net/http"
	"github.com/gorilla/mux"
)

func RequestPathVariable(request *http.Request, paramName string) string {
	vars := mux.Vars(request)
	var p string
	p = vars[paramName]
	return p
}
