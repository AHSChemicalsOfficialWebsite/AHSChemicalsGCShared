package send_email

import (
	"errors"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/company_details"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Sends mail using sendgrid api
func SendEmail(metaData EmailMetaData) (*rest.Response, error){

	from := mail.NewEmail("AHSChemicals", company_details.COMPANYEMAIL)

	var recipients []*mail.Email
	for email, name := range metaData.Recipients {
		recipients = append(recipients, mail.NewEmail(name, email))
	}

	if len(recipients) == 0 {
		return nil, errors.New("No recipents found to send a mail")
	}

	if metaData.TemplateID == "" {
		return nil, errors.New("No template ID found for sending the mail")
	}

	// Configure personalization for dynamic template data.
	p := mail.NewPersonalization()
	p.AddTos(recipients...)

	//Set the key and values for the dynamic template data
	for key, value := range metaData.Data {
		p.SetDynamicTemplateData(key, value)
	}

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.AddPersonalizations(p)
	message.SetTemplateID(metaData.TemplateID)

	for _, item := range metaData.Attachments {
		attachment := mail.NewAttachment()
		attachment.SetType(item.MimeType)
		attachment.SetContent(item.Base64Content)
		attachment.SetFilename(item.FileName)
		attachment.SetDisposition("attachment")
		message.AddAttachment(attachment)
	}

	client := sendgrid.NewSendClient(SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		return nil, err
	}

	return response, nil
}
