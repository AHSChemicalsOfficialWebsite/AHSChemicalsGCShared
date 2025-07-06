// Package send_mobile_message provides functionality to send SMS messages 
// to one or more recipients using Twilio's messaging API.
package send_mobile_message

import (
	"errors"
	"log"
	"strings"

	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// MobileMessageMetaData represents metadata required to send a mobile message.
// 
// Fields:
//   - UID: Unique identifier of the user triggering the notification (optional; can be nil for admin-triggered messages).
//   - Message: The content/body of the message to be sent.
type MobileMessageMetaData struct {
	UID     *string `json:"uid"`     // UID of the user triggering the notification. Can be nil if it’s an admin-triggered message.
	Message string  `json:"message"` // Message body to be sent.
}

// SendMobileMessage is a helper function which sends an SMS message to a list of predefined recipients using Twilio.
//
// It reads the recipient phone numbers and Twilio credentials from global configuration variables.
// If any required credential is missing or the message content is empty, an error is returned.
//
// Parameters:
//   - metaData: MobileMessageMetaData struct containing message details.
//
// Returns:
//   - error: Non-nil if the message fails validation or sending fails for all recipients.
//
// Notes:
//   - Recipients are expected to be defined in the global variable TWILIO_RECIPIENTS_PHONE, separated by semicolons.
//   - It logs the success/failure status for each recipient.
//
// Example:
//
//	err := SendMobileMessage(MobileMessageMetaData{
//		UID:     nil,
//		Message: "Hello from Twilio!",
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
func SendMobileMessage(metaData MobileMessageMetaData) error {
	if metaData.UID == nil {
		log.Printf("No UID provided for the Twilio message. Sending message without a UID")
	}

	if metaData.Message == "" {
		return errors.New("no message provided")
	}

	if TWILIO_RECIPIENTS_PHONE == "" || TWILIO_ACCOUNT_SID == "" || TWILIO_AUTH_TOKEN == "" || TWILIO_FROM_PHONE == "" {
		return errors.New("missing Twilio credentials. Check the credentials")
	}

	toPhones := strings.Split(TWILIO_RECIPIENTS_PHONE, ";")

	// Initialize Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})

	var failedRecipients []string

	// Send message to each recipient
	for _, receiverPhone := range toPhones {
		params := &twilioApi.CreateMessageParams{}
		params.SetTo(receiverPhone)
		params.SetFrom(TWILIO_FROM_PHONE)
		params.SetBody(metaData.Message)

		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			log.Printf("Error sending message to %s: %v", receiverPhone, err)
			failedRecipients = append(failedRecipients, receiverPhone)
		} else {
			log.Printf("Successfully sent message to %s. SID: %s", receiverPhone, *resp.Sid)
		}
	}

	// Log result
	if len(failedRecipients) > 0 {
		log.Printf("Failed to send mobile messages to some recipients: %v", failedRecipients)
	}

	return nil
}