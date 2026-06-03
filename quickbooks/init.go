package quickbooks

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/gcp"
)

var (
	CLIENT_ID                  string
	CLIENT_SECRET              string
	AUTH_CALLBACK_URL          string //Redirect URL. When the first authentication via initial auth url completes, it redirects to this url
	AUTH_CALLBACK_REDIRECT_URL string //Redirects user to a successful login page when quickbooks is authenticated
	API_URL                    string //base url which contains either the debug or production api url.
	WEBHOOK_VERIFY_TOKEN       string
	initOnce                   sync.Once
)

func validateQuickBooksCredentials() {
	if CLIENT_ID == "" || CLIENT_SECRET == "" || AUTH_CALLBACK_URL == "" || AUTH_CALLBACK_REDIRECT_URL == "" || API_URL == "" || WEBHOOK_VERIFY_TOKEN == ""{
		log.Fatalf("Error initializing QuickBooks credentials: missing required environment variables")
	}
}

func InitDebug() {
	initOnce.Do(func() {
		CLIENT_ID = os.Getenv("QUICKBOOKS_DEBUG_CLIENT_ID")
		CLIENT_SECRET = os.Getenv("QUICKBOOKS_DEBUG_CLIENT_SECRET")
		AUTH_CALLBACK_URL = os.Getenv("QUICKBOOKS_DEBUG_AUTH_CALLBACK_URL")
		AUTH_CALLBACK_REDIRECT_URL = os.Getenv("QUICKBOOKS_DEBUG_AUTH_CALLBACK_REDIRECT_URL")
		API_URL = os.Getenv("QUICKBOOKS_DEBUG_API_URL")
		WEBHOOK_VERIFY_TOKEN = os.Getenv("QUICKBOOKS_DEBUG_WEBHOOK_VERIFY_TOKEN")
		validateQuickBooksCredentials()
		log.Println("Initialized quickbooks credentials in debug...")
	})
}

func InitProd(projectID string, ctx context.Context) {
	initOnce.Do(func() {
		CLIENT_ID = gcp.LoadSecretsHelper(projectID, "QUICKBOOKS_CLIENT_ID")
		CLIENT_SECRET = gcp.LoadSecretsHelper(projectID, "QUICKBOOKS_CLIENT_SECRET")
		AUTH_CALLBACK_URL = gcp.LoadSecretsHelper(projectID, "QUICKBOOKS_AUTH_CALLBACK_URL")
		AUTH_CALLBACK_REDIRECT_URL = gcp.LoadSecretsHelper(projectID, "QUICKBOOKS_AUTH_CALLBACK_REDIRECT_URL")
		API_URL = gcp.LoadSecretsHelper(projectID, "QUICKBOOKS_API_URL")
		WEBHOOK_VERIFY_TOKEN = gcp.LoadSecretsHelper(projectID, "QUICKBOOKS_WEBHOOK_VERIFY_TOKEN")
		validateQuickBooksCredentials()
		log.Println("QuickBooks credentials initialized for PRODUCTION environment.")
	})
}
