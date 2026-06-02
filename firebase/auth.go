package firebase

import (
	"errors"
	"net/http"
	"strings"
)

//Custom claim role assigned to each firebase authenticated user
type Role string

func (r Role) String() string {
	return string(r)
}

const (
	UserRole       Role = "user"
	AdminRole      Role = "admin"
	SuperAdminRole Role = "superadmin"
)

var ValidRoles = map[Role]struct{}{
	UserRole:       {},
	AdminRole:      {},
	SuperAdminRole: {},
}

func IsAuthorized(request *http.Request, roles ...Role) (string, error) {
	ctx := request.Context()

	authHeader := request.Header.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid Authorization header format")
	}

	idToken := parts[1]
	token, err := AuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}

	// check for role if any
	for _, role := range roles {
		claimValue, ok := token.Claims["role"].(string)
		if ok && claimValue != "" && claimValue == role.String() {
			return token.UID, nil
		} else {
			return "", errors.New("invalid role")
		}
	}
	return token.UID, nil
}
