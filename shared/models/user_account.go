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
	if utils.HasDuplicateStrings(c.Customers) {
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

type UpdatedAccount struct {
	UID          string   `json:"uid"`
	Brands       []string `json:"brands"`
	Customers    []string `json:"customers"`
}

func (u *UpdatedAccount) Validate() error{
	if u.UID == ""{
		return errors.New("No uid found for the updated account")
	}
	if len(u.Brands) == 0{
		return errors.New("No brands found for the updated account")
	}
	if utils.HasDuplicateStrings(u.Brands){
		return errors.New("Duplicate brands found for the updated account")
	}
	if len(u.Customers) == 0{
		return errors.New("No customers found for the updated account")
	}
	if utils.HasDuplicateStrings(u.Customers){
		return errors.New("Duplicate customers found for the updated account")
	}
	return nil
}

func (u *UpdatedAccount) ToMap() map[string]any {
	return map[string]any{
		"customers": u.Customers,
		"brands":    u.Brands,
	}
}