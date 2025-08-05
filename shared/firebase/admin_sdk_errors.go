// package firebase_shared is used to initialize the Firebase Admin SDK, initializations, and utilities.
package firebase_shared

import (
	"encoding/json"
	"strings"
)

// FirebaseErrorResponse represents the structure of an error response returned by Firebase Admin SDK.
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

// ExtractFirebaseErrorFromResponse attempts to extract a FirebaseErrorResponse from a given error.
//
// Firebase sometimes returns error responses as JSON strings prefixed with additional data.
// This function locates the JSON portion, unmarshals it into a FirebaseErrorResponse struct, and returns it.
//
// Parameters:
//   - err: The original error object returned by Firebase.
//
// Returns:
//   - A pointer to FirebaseErrorResponse if extraction is successful.
//   - nil if the JSON portion is not found.
//   - If JSON is malformed, returns a FirebaseErrorResponse with empty or partial fields.
func ExtractFirebaseErrorFromResponse(err error) *FirebaseErrorResponse {
	errString := err.Error()

	// Locate the start of the JSON object within the error string.
	start := strings.Index(errString, "{")
	if start == -1 {
		return nil
	}

	var firebaseError FirebaseErrorResponse
	jsonPart := errString[start:]
	if unmarshalErr := json.Unmarshal([]byte(jsonPart), &firebaseError); unmarshalErr != nil {
		return &firebaseError // Return partial object even if unmarshaling fails.
	}
	return &firebaseError
}