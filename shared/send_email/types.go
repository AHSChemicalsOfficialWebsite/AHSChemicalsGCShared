package send_email


type Attachment struct {
	Base64Content string `json:"name"`
	MimeType      string `json:"data"`
	FileName      string `json:"content_type"`
}

type EmailMetaData struct {
	Recipients  map[string]string `json:"recipients"` //Key value pairs of recipient names and their respective emails 
	Data        map[string]any    `json:"data"` //Key value pairs of data that will be used in the send grid dynamic template
	TemplateID  string            `json:"template_id"` //Sendgrid dynamic template 
	Attachments []Attachment      `json:"attachments"`
}
