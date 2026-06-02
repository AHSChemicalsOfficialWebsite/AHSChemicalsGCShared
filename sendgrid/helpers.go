package sendgrid

// Note: Filenames must have the extension with the name. In some applications such as outlook mobile, it will
// not render the attachment if the extension is not provided along with the filename.
func CreateEmailAttachments(base64Contents, mimeTypes, filenames []string) []EmailAttachment {
	attachments := make([]EmailAttachment, 0)
	for i := range base64Contents {
		attachment := EmailAttachment{
			Base64Content: base64Contents[i],
			MimeType:      mimeTypes[i],
			FileName:      filenames[i],
		}
		attachments = append(attachments, attachment)
	}
	return attachments
}

// Note: Filenames must have the extension with the name. In some applications such as outlook mobile, it will
// not render the attachment if the extension is not provided along with the filename.
func CreateSingleEmailAttachment(base64Content, mimeType, filename string) EmailAttachment {
	return EmailAttachment{
		Base64Content: base64Content,
		MimeType:      mimeType,
		FileName:      filename,
	}
}
