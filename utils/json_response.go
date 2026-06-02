package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessPaylod[T any] struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    T      `json:"data"`
}

func WriteJSONSuccess(response http.ResponseWriter, statusCode int, message string, data any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	payload := SuccessPaylod[any]{
		Code:    statusCode,
		Message: message,
	}
	if data != nil {
		payload.Data = data
	}
	if err := json.NewEncoder(response).Encode(payload); err != nil {
		log.Printf("JSON encode response error: %v", err)
	}
}

func WriteJSONError(response http.ResponseWriter, statusCode int, message string) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	if err := json.NewEncoder(response).Encode(map[string]any{"code": statusCode, "message": message}); err != nil {
		log.Printf("Error writing the error response: %v", err)
	}
}
