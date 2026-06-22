package firebase

import (
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func IsAuthorized(request *http.Request, roles ...models.Role) (*auth.Token, error) {
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
	//In case no role is specified, return the token
    if len(roles) == 0 {
        return token, nil
    }
	for _, role := range roles {
		claimValue, ok := token.Claims["role"].(string)
		if ok && claimValue != "" && claimValue == role.String() {
			return token, nil
		}
	}
	return nil, ErrInvalidRole
}
