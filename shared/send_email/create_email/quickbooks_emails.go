package create_email

import (
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/company_details"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/send_email"
)

func CreateQuickBooksSessionExpiredEmail(reconnectURL string) *send_email.EmailMetaData{
	emailData := &send_email.EmailMetaData{
		Recipients: company_details.EMAILINTERNALRECIPENTS,
		Data: map[string]any{
			"reconnect_url": reconnectURL,
		},
		TemplateID:  send_email.QUICKBOOKS_SESSION_EXPIRED_TEMPLATE_ID,
		Attachments: []send_email.Attachment{},
	}
	return emailData
}