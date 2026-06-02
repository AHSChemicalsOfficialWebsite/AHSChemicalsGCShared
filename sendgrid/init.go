package sendgrid

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/gcp"
)

var (
	ApiKey   string
	initOnce sync.Once
)

const (
	ACCOUNT_CREATED_USER_TEMPLATE_ID          = "d-5130ffca9ed749458e64388df9931af8"
	ACCOUNT_DELETED_USER_TEMPLATE_ID          = "d-16455882508f49a69cc64af9df98bc79"
	CONTACT_US_INTERNAL_TEMPLATE_ID           = "d-40d7a73d43044231ab2e3e20d6760db5"
	CONTACT_US_USER_TEMPLATE_ID               = "d-84ee771b07344758a5bd0fe38afd8ae8"
	ORDER_PLACED_INTERNAL_TEMPLATE_ID         = "d-b3585038c1094054a8404423e6051573"
	ORDER_PLACED_USER_TEMPLATE_ID             = "d-5d7698237ee7491c96af123f964b4321"
	ORDER_UPDATED_USER_TEMPLATE_ID            = "d-f88446dbe2cb4b83b7452f89906163ee"
	ORDER_UPDATED_INTERNAL_TEMPLATE_ID        = "d-2df971c9b243495faf98f119c7f25dd1"
	ORDER_DELIVERED_INTERNAL_TEMPLATE_ID      = "d-10e63e3c62ac468da03efee2b15a5e96"
	ORDER_DELIVERED_USER_TEMPLATE_ID          = "d-a6344e4ac9254af28afde74ad2b81bc4"
	FINAL_INVOICE_INTERNAL_TEMPLATE_ID        = "d-7d82f2a245a54afd8272e682d1a29353"
	QUICKBOOKS_QUICKBOOKS_EXPIRED_TEMPLATE_ID = "d-5c6a813945bd4669b99225c9f6d13dca"
)

func Init(projectID string, ctx context.Context) {
	initOnce.Do(func() {
		if projectID == "" { // For local testing
			ApiKey = os.Getenv("SENDGRID_API_KEY")
		} else {
			ApiKey = gcp.LoadSecretsHelper(projectID, "SENDGRID_API_KEY")
		}
		if ApiKey == "" {
			log.Fatal("SENDGRID_API_KEY is not set")
		}
		log.Println("Initialized SendGrid credentials")
	})
}
