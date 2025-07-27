package create_email

import (
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/send_email"
)

func CreateUserAccountCreatedEmail(createdUser *models.UserAccount) *send_email.EmailMetaData {
	emailData := &send_email.EmailMetaData{
		Recipients: map[string]string{createdUser.Email: createdUser.Name},
		Data: map[string]any{
			"name": createdUser.Name,
			"email": createdUser.Email,
			"password": createdUser.Password,
		},
		TemplateID:  send_email.ACCOUNT_CREATED_USER_TEMPLATE_ID,
		Attachments: []send_email.Attachment{},
	}
	return emailData
}

func CreateDeleteUserAccountEmail(deletedUser *models.UserAccount) *send_email.EmailMetaData {
	emailData := &send_email.EmailMetaData{
		Recipients: map[string]string{deletedUser.Email: deletedUser.Name},
		Data: map[string]any{
			"name": deletedUser.Name,
			"email": deletedUser.Email,
		},
		TemplateID:  send_email.ACCOUNT_DELETED_USER_TEMPLATE_ID,
		Attachments: []send_email.Attachment{},
	}
	return emailData
}