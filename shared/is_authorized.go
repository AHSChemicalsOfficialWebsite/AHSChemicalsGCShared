package shared

import (
	"errors"
	"net/http"
	"strings"
)

// IsAuthorized verifies if the incoming HTTP request is authenticated and authorized as an admin.
//
// This function extracts and validates the Bearer token from the Authorization header,
// verifies the token using Firebase Authentication, and checks for a custom "admin" claim.
//
// Parameters:
//   - request: *http.Request representing the incoming HTTP request.
//
// Returns:
//   - nil if the request is authorized as an admin.
//   - error if the Authorization header is invalid, token verification fails, or if the user is not an admin.
//
// Expected Authorization header format:
//   Authorization: Bearer <idToken>
func IsAuthorized(request *http.Request) error {
	ctx := request.Context()

	// Extract and validate the Authorization header.
	authHeader := request.Header.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return errors.New("invalid Authorization header format")
	}

	// Extract the ID token from the Authorization header.
	idToken := parts[1]
	token, err := AuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return err
	}

	// Check if the "admin" claim exists and is set to true.
	adminClaim, ok := token.Claims["admin"].(bool)
	if !ok || !adminClaim {
		return errors.New("unauthorized: only admins are allowed to perform this operation")
	}

	return nil
}