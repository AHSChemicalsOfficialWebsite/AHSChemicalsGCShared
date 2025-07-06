// Package send_email provides initialization logic for SendGrid credentials,
// supporting both local development (debug) and production environments
// (Google Cloud Functions).
package send_email

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/gcp"
)

// Global SendGrid credentials, initialized based on the environment.
// These are required to authenticate and send emails via the SendGrid API.
var (
	SENDGRID_API_KEY   string // API key used to authenticate with the SendGrid service
	SENDGRID_FROM_MAIL string // Email address used as the sender in outgoing emails
)

// InitSendGridDebug initializes SendGrid credentials for local or debug environments.
//
// It reads the credentials from environment variables:
//   - SENDGRID_API_KEY
//   - SENDGRID_FROM_MAIL
//
// This function should be used in local development or testing environments
// where secrets are stored in local environment variables.
func InitSendGridDebug() {
	SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")
	SENDGRID_FROM_MAIL = os.Getenv("SENDGRID_FROM_MAIL")

	log.Println("Initialized SendGrid credentials in debug mode")
}

// InitSendGridProd initializes SendGrid credentials for production environments.
//
// It uses Google Cloud's metadata server to get the current project ID,
// and then loads secrets using the gcp.LoadSecretsHelper function.
//
// This function is designed to be used in production on Google Cloud (e.g., Cloud Run or GCF),
// where secrets are securely stored in Secret Manager.
//
// The initialization is wrapped in a `sync.Once` construct to ensure it is only run once.
//
// Secrets expected in Secret Manager:
//   - SENDGRID_API_KEY
//   - SENDGRID_FROM_MAIL
//
// Parameters:
//   - ctx: context.Context used to access Google Cloud metadata and Secret Manager.
//
// Panics:
//   - If the Google Cloud project ID cannot be retrieved from the metadata server.
func InitSendGridProd(ctx context.Context) {
	shared.InitSendGridOnce.Do(func() {
		projectID, err := metadata.ProjectIDWithContext(ctx)
		if err != nil {
			log.Fatalf("Error loading Google Cloud project ID: %v", err)
		}

		SENDGRID_API_KEY = gcp.LoadSecretsHelper(projectID, "SENDGRID_API_KEY")
		SENDGRID_FROM_MAIL = gcp.LoadSecretsHelper(projectID, "SENDGRID_FROM_MAIL")
	})
}