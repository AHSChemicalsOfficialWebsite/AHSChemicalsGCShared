package validation

import (
	"net/mail"
	"regexp"
	"strings"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

var (
	nameRegex     = regexp.MustCompile(`^[A-Za-z]+(?:\s+[A-Za-z]+)+$`)
	locationRegex = regexp.MustCompile(`^[a-zA-Z0-9 ,._-]{10,50}$`)
	messageRegex  = regexp.MustCompile(`^[a-zA-Z0-9\s.,'"!?;:\-_()\n\r]{10,300}$`)
)

func ValidateContactUsForm(c *models.ContactUsForm) error {
	if err := validateEmail(c.Email); err != nil {
		return err
	}
	if err := validateName(c.Name); err != nil {
		return err
	}
	if err := validatePhone(c.Phone); err != nil {
		return err
	}
	if err := validateLocation(c.Location); err != nil {
		return err
	}
	if err := validateMessage(c.Message); err != nil {
		return err
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return ErrEmailRequired
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return ErrEmailInvalid
	}
	return nil
}

func validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrNameRequired
	}
	if !nameRegex.MatchString(name) {
		return ErrNameInvalid
	}
	return nil
}

func validatePhone(phone string) error {
	if phone == "" {
		return ErrPhoneRequired
	}
	return nil
}

func validateLocation(location string) error {
	if location == "" {
		return ErrLocationRequired
	}
	if !locationRegex.MatchString(location) {
		return ErrLocationInvalid
	}
	return nil
}

func validateMessage(message string) error {
	if message == "" {
		return ErrMessageRequired
	}
	if !messageRegex.MatchString(message) {
		return ErrMessageInvalid
	}
	return nil
}