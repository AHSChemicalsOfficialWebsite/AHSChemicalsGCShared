package firebase

import (
	"encoding/json"
	"errors"
	"strings"
)

var (
	ErrInvalidHeaderFormat = errors.New("invalid Authorization header format")
	ErrInvalidRole         = errors.New("invalid auth role")
)

// Reference: https://firebase.google.com/docs/reference/rest/auth
var RestApiErrorsMap = map[string]string{
	"EMAIL_EXISTS":                "The email address is already in use by another account.",
	"INVALID_EMAIL":               "The email address is badly formatted.",
	"WEAK_PASSWORD":               "The password must be 6 characters long or more.",
	"OPERATION_NOT_ALLOWED":       "Password sign-in is disabled for this project.",
	"TOO_MANY_ATTEMPTS_TRY_LATER": "We have blocked all requests from this device due to unusual activity. Try again later.",
	"USER_NOT_FOUND":              "There is no user record corresponding to this identifier. The user may have been deleted.",
	"USER_DISABLED":               "The user account has been disabled by an administrator.",
	"INVALID_ID_TOKEN":            "The user's credential is no longer valid. The user must sign in again.",
}

// FirebaseErrorResponse represents the structure of an error response returned by Firebase RESI API errors.
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
func ExtractFirebaseErrorFromResponse(errString string) *FirebaseErrorResponse {
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
