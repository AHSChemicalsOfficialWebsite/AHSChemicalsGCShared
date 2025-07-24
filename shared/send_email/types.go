package send_email

//Atachment for the SendGrid Email. 
type Attachment struct {
	Base64Content string `json:"name"`  //The encoded base64 content 
	MimeType      string `json:"data"` //MimeType of the attachment being sent. Eg "application/pdf"
	FileName      string `json:"content_type"` //FileName of the attachment. Must have an extension because on some mail applications, the attachment cannot be viewed if it does not have an extension. Eg: "purchase_order.pdf"
}

//EmailMetaData represents the entire email object that will be sent to the SendGrid API
type EmailMetaData struct {
	Recipients  map[string]string `json:"recipients"` //Map of recipent emails as keys and names as values 
	Data        map[string]any    `json:"data"` //Key value pairs of data that will be used in the send grid dynamic template
	TemplateID  string            `json:"template_id"` //Sendgrid dynamic template 
	Attachments []Attachment      `json:"attachments"` //List of attachments for the email
}