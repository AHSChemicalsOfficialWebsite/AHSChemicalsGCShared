package sendgrid

import (
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/company_details"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func CreateContactUsInternalEmailMetaData(c *models.ContactUsForm) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: company_details.EmailInternalRecipients,
		Data: map[string]any{
			"name":     c.Name,
			"email":    c.Email,
			"phone":    c.Location,
			"location": c.Location,
			"message":  c.Message,
			"year":     time.Now().UTC().Year(),
		},
		TemplateID: CONTACT_US_INTERNAL_TEMPLATE_ID,
	}
	return emailData
}

func CreateContactUsUserEmailMetaData(c *models.ContactUsForm) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: map[string]string{c.Email: c.Name},
		Data: map[string]any{
			"name":     c.Name,
			"email":    c.Email,
			"phone":    c.Location,
			"location": c.Location,
			"message":  c.Message,
		},
		TemplateID: CONTACT_US_USER_TEMPLATE_ID,
	}
	return emailData
}
