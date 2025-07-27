package models

import (
	"errors"
	"fmt"
)

//PhoneNumber represents a phone number stored in firebase authentication property
type PhoneNumber struct {
	Code   string `json:"code"`
	Number string `json:"number"`
}

//Firebase authentication handles phone number validation so only writing the empty cases
func (p *PhoneNumber) Validate() error {
	if p.Code == "" {
		return errors.New("Country code of the user cannot be empty")
	}
	if p.Number == "" {
		return errors.New("Phone number of the user cannot be empty")
	}
	return nil
}

func (p *PhoneNumber) Format() string {
	return fmt.Sprintf("%s%s", p.Code, p.Number)
}