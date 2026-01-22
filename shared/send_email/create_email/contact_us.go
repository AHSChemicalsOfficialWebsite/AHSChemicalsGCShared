//package create_email is a utility package for creating custom emails for sendgrid emails
package create_email

import (
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/company_details"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/send_email"
)

func CreateContactUsAdminEmail(c *models.ContactUsForm) *send_email.EmailMetaData {
	emailData := &send_email.EmailMetaData{
		Recipients: company_details.EMAILINTERNALRECIPENTS,
		Data: map[string]any{
			"name":     c.Name,
			"email":    c.Email,
			"phone":    c.Location,
			"location": c.Location,
			"message":  c.Message,
			"year":     time.Now().UTC().Year(),
		},
		TemplateID:  send_email.CONTACT_US_ADMIN_TEMPLATE_ID,
	}
	return emailData
}

func CreateContactUsUserEmail(c *models.ContactUsForm) *send_email.EmailMetaData {
	emailData := &send_email.EmailMetaData{
		Recipients: map[string]string{c.Email: c.Name},
		Data: map[string]any{
			"name":     c.Name,
			"email":    c.Email,
			"phone":    c.Location,
			"location": c.Location,
			"message":  c.Message,
		},
		TemplateID:  send_email.CONTACT_US_USER_TEMPLATE_ID,
	}
	return emailData
}