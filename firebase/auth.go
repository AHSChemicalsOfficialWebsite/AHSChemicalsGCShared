package firebase

import (
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

//Custom claim role assigned to each firebase authenticated user
type Role string

func (r Role) String() string {
	return string(r)
}

const (
	UserRole       Role = "user"
	AdminRole      Role = "admin"
	SuperAdminRole Role = "super_admin"
)

var ValidRoles = map[Role]struct{}{
	UserRole:       {},
	AdminRole:      {},
	SuperAdminRole: {},
}

func IsAuthorized(request *http.Request, roles ...Role) (*auth.Token, error) {
	ctx := request.Context()

	authHeader := request.Header.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, ErrInvalidHeaderFormat
	}

	idToken := parts[1]
	token, err := AuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	// check for role if any
	for _, role := range roles {
		claimValue, ok := token.Claims["role"].(string)
		if ok && claimValue != "" && claimValue == role.String() {
			return token, nil
		} else {
			return nil, ErrInvalidRole
		}
	}
	return token, nil
}
