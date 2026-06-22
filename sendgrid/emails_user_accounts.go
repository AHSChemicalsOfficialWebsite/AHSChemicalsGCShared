package sendgrid

import (
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func CreateUserAccountCreatedEmailMetaData(createdUser *models.UserAccountCreate) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: map[string]string{createdUser.Email: createdUser.Name},
		Data: map[string]any{
			"name":     createdUser.Name,
			"email":    createdUser.Email,
			"password": createdUser.Password,
		},
		TemplateID:  ACCOUNT_CREATED_USER_TEMPLATE_ID,
	}
	return emailData
}

func CreateDeleteUserAccountEmailMetaData(email, name string) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: map[string]string{email: name},
		Data: map[string]any{
			"name":  name,
			"email": email,
		},
		TemplateID: ACCOUNT_DELETED_USER_TEMPLATE_ID,
	}
	return emailData
}
