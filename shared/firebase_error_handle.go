package shared

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type FirebaseErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
	} `json:"error"`
}

func ExtractFirebaseErrorFromResponse(err error) *FirebaseErrorResponse {
	errString := err.Error()
	start := strings.Index(errString, "{")
	var firebaseError FirebaseErrorResponse

	if start == -1 {
		return nil
	}

	jsonPart := errString[start:]
	if unmarshalErr := json.Unmarshal([]byte(jsonPart), &firebaseError); unmarshalErr != nil {
		return &firebaseError
	}
	return &firebaseError
}

func WriteJSONError(response http.ResponseWriter, statusCode int, message string) {
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(statusCode)
	if err := json.NewEncoder(response).Encode(map[string]any{"code": statusCode, "message": message}); err != nil {
		log.Printf("Error writing the error response: %v", err)
	}
}
