package models

import (

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
)

type UserAccountCreate struct {
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Customers []string      `json:"customers"`
	Brands    []string      `json:"brands"`
	Role      firebase.Role `json:"role"`
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
	Name      string        `json:"name" firestore:"name"`
	Email     string        `json:"email" firestore:"email"`
	Customers []string      `json:"customers" firestore:"customers"`
	Brands    []string      `json:"brands" firestore:"brands"`
	Role      firebase.Role `json:"role" firestore:"role"`
}
