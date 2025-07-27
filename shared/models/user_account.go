package models

import (
	"errors"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/utils"
)

//For better readability, creating a struct to represent the phone number
type PhoneNumber struct {
	Code   string `json:"code"`
	Number string `json:"number"`
}

// CreateUserAccount represents a new user account created by the admin 
type CreateUserAccount struct {
	ID          string            `json:"id" firestore:"id"` //Uid of the user given by firebase when the user is created
	Name        string            `json:"name"`
	PhoneNumber PhoneNumber       `json:"phoneNumber"`
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	Customers   []string          `json:"customers" firestore:"customers"`
	Brands      []string          `json:"brands" firestore:"brands"`
}

func (c *CreateUserAccount) Validate() error {
	if c.Name == "" {
		return errors.New("Name of the user cannot be empty")
	}
	if c.PhoneNumber.Code == "" {
		return errors.New("Country code of the user cannot be empty")
	}
	if c.PhoneNumber.Number == "" {
		return errors.New("Phone number of the user cannot be empty")
	}
	if c.Email == "" {
		return errors.New("Email of the user cannot be empty")
	}
	if c.Password == "" {
		return errors.New("Password of the user cannot be empty")
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

func (c *CreateUserAccount) MapToFirestore() map[string]any {
	return map[string]any{
		"customers": c.Customers,
		"brands":    c.Brands,
	}
}