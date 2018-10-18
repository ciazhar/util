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
	response, _ := json.Marshal(map[string]string{"error": msg})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
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

func RespondWithJson(w http.ResponseWriter, code int, data interface{}) {
	payload := ResponseData{
		Data:data,
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
