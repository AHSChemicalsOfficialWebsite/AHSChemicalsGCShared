package models

import (
	"errors"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/utils"
)

// UserAccount represents a new user account created by the admin 
type UserAccount struct {
	ID          string            `json:"id"` //Uid of the user given by firebase when the user is created
	Name        string            `json:"name"`
	PhoneNumber PhoneNumber       `json:"phoneNumber"`
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	Customers   []string          `json:"customers" firestore:"customers"`
	Brands      []string          `json:"brands" firestore:"brands"`
}

// Validate validates the user account. Email and password are validated by firebase authentication, so only the necessary fields are validated
func (c *UserAccount) Validate() error {
	if c.Name == "" {
		return errors.New("Name of the user cannot be empty")
	}
	if err := c.PhoneNumber.Validate(); err != nil {
		return err
	}
	if len(c.Customers) == 0 {
		return errors.New("At least one customer is required for the user")
	}
	if utils.HasDuplicates(c.Customers) {
		return errors.New("Customers of the user cannot have duplicates")
	}
	if len(c.Brands) == 0 {
		return errors.New("Brands of the user cannot be empty")
	}
	return nil
}

func (c *UserAccount) MapToFirestore() map[string]any {
	return map[string]any{
		"customers": c.Customers,
		"brands":    c.Brands,
	}
}

func (c *UserAccount) MapForFrontend() map[string]any {
	return map[string]any{
		"uid":       c.ID,
		"name":      c.Name,
		"email":     c.Email,
		"customers": c.Customers,
		"brands":    c.Brands,
	}
}