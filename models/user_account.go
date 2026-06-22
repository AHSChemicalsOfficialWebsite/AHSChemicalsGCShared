package models

// Custom claim role assigned to each firebase authenticated user
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

type UserAccountCreate struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Customers []string `json:"customers"`
	Brands    []string `json:"brands"`
	Role      Role     `json:"role"`
}

func (c *UserAccountCreate) ToFirestoreMap() map[string]any {
	return map[string]any{
		"name":      c.Name,
		"email":     c.Email,
		"customers": c.Customers,
		"brands":    c.Brands,
		"role":      c.Role,
	}
}

// Used when storing/retrieving from Firestore (no password)
type UserAccount struct {
	Name      string   `json:"name" firestore:"name"`
	Email     string   `json:"email" firestore:"email"`
	Customers []string `json:"customers" firestore:"customers"`
	Brands    []string `json:"brands" firestore:"brands"`
	Role      Role     `json:"role" firestore:"role"`
	FCMTokens []string `json:"fcmtokens" firestore:"fcmtokens"`
}
