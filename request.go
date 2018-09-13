package util

import (
	"net/http"
	"strconv"
	"encoding/json"
)

func RequestParamInt(r *http.Request, paramName string) int {
	response, _ := strconv.Atoi(r.URL.Query().Get(paramName))
	return response
}

func RequestParamString(r *http.Request, paramName string) string {
	response := r.URL.Query().Get(paramName)
	return response
}

func RequestBody(r *http.Request, v ok) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.Validate()
}
