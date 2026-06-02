package sendgrid

import (
	"log"
	"net/http"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/company_details"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(e *EmailMetaData) {
	from := mail.NewEmail("AHSChemicals", company_details.Email)

	var recipients []*mail.Email
	for email, name := range e.Recipients {
		recipients = append(recipients, mail.NewEmail(name, email))
	}

	if len(recipients) == 0 {
		log.Println("No recipients provided")
		return
	}

	if e.TemplateID == "" {
		log.Println("No template ID provided")
		return
	}

	// Personalization for dynamic template data.
	p := mail.NewPersonalization()
	p.AddTos(recipients...)

	// Key-value pairs used in the send grid dynamic template
	for key, value := range e.Data {
		p.SetDynamicTemplateData(key, value)
	}

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.AddPersonalizations(p)
	message.SetTemplateID(e.TemplateID)

	for _, item := range e.Attachments {
		attachment := mail.NewAttachment()
		attachment.SetType(item.MimeType)
		attachment.SetContent(item.Base64Content)
		attachment.SetFilename(item.FileName)
		attachment.SetDisposition("attachment")
		message.AddAttachment(attachment)
	}

	client := sendgrid.NewSendClient(ApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println("Email sending failed: " + err.Error())
		return
	}

	if response.StatusCode != http.StatusAccepted {
		log.Printf("Email sending failed with status code: %d, header: %v, body: %s", response.StatusCode, response.Headers, response.Body)
	}

	log.Printf("Email sent with status code: %d", response.StatusCode)
}
