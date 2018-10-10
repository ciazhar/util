package util

import (
	"net/http"
	"encoding/json"
)

type ResponseData struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func Json(w http.ResponseWriter, code int, message string ,data interface{}) {
	payload := ResponseData{
		Message:message,
		Data:data,
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
