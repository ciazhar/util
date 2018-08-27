package util

import (
	"net/http"
	"strconv"
)

func RequestParamInt(r *http.Request, paramName string) int {
	response, _ := strconv.Atoi(r.URL.Query().Get(paramName))
	return response
}

func RequestParamString(r *http.Request, paramName string) string {
	response := r.URL.Query().Get(paramName)
	return response
}