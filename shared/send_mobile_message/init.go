// Package send_mobile_message handles initialization and configuration
// for sending Twilio SMS messages using either debug (local) or production settings.
package send_mobile_message

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/gcp"
)

// Twilio credentials and recipient list.
// These variables are populated during initialization using either
// InitTwilioDebug (for local/debug environments) or InitTwilioProd (for production).
var (
	TWILIO_ACCOUNT_SID      string // Twilio Account SID
	TWILIO_AUTH_TOKEN       string // Twilio Auth Token
	TWILIO_FROM_PHONE       string // Sender phone number registered with Twilio
	TWILIO_RECIPIENTS_PHONE string // Semicolon-separated list of recipient phone numbers
)

// InitTwilioDebug initializes Twilio credentials from environment variables.
//
// This function should be used in local development or debugging scenarios,
// where credentials are stored in environment variables (e.g., via .env or local shell).
//
// Environment variables required:
//   - TWILIO_ACCOUNT_SID
//   - TWILIO_AUTH_TOKEN
//   - TWILIO_FROM_PHONE
//   - TWILIO_RECIPIENTS_PHONE
func InitTwilioDebug() {
	TWILIO_ACCOUNT_SID = os.Getenv("TWILIO_ACCOUNT_SID")
	TWILIO_AUTH_TOKEN = os.Getenv("TWILIO_AUTH_TOKEN")
	TWILIO_FROM_PHONE = os.Getenv("TWILIO_FROM_PHONE")
	TWILIO_RECIPIENTS_PHONE = os.Getenv("TWILIO_RECIPIENTS_PHONE")

	log.Printf("Initialized Twilio debug credentials")
}

// InitTwilioProd initializes Twilio credentials by securely retrieving them
// from Google Cloud Secret Manager using the application's project ID.
//
// This function is intended to be used in production environments on Google Cloud,
// and it ensures secrets are not hardcoded or exposed in environment variables.
//
// It runs only once per application lifecycle using sync.Once from the shared package.
//
// Secrets retrieved from Secret Manager:
//   - TWILIO_ACCOUNT_SID
//   - TWILIO_AUTH_TOKEN
//   - TWILIO_FROM_PHONE
//   - TWILIO_RECIPIENTS_PHONE
//
// Parameters:
//   - ctx: context.Context for metadata and secret retrieval.
//
// Panics:
//   - If the project ID cannot be retrieved from the GCP metadata server.
func InitTwilioProd(ctx context.Context) {
	shared.InitTwilioOnce.Do(func() {
		projectID, err := metadata.ProjectIDWithContext(ctx)
		if err != nil {
			log.Fatalf("Error loading Google Cloud project ID: %v", err)
		}

		TWILIO_ACCOUNT_SID = gcp.LoadSecretsHelper(projectID, "TWILIO_ACCOUNT_SID")
		TWILIO_AUTH_TOKEN = gcp.LoadSecretsHelper(projectID, "TWILIO_AUTH_TOKEN")
		TWILIO_FROM_PHONE = gcp.LoadSecretsHelper(projectID, "TWILIO_FROM_PHONE")
		TWILIO_RECIPIENTS_PHONE = gcp.LoadSecretsHelper(projectID, "TWILIO_RECIPIENTS_PHONE")

		log.Printf("Initialized Twilio production credentials")
	})
}