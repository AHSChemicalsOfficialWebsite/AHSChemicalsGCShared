package company_details

import (
	"context"
	"encoding/json"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/gcp"
	"log"
	"os"
	"sync"
)

const (
	Name       = "Azure Hospitality Supply"
	WebsiteUrl = "https://azurehospitalitysupply.com"
)

var (
	Email                   string
	Phone                   string
	AddressLine1            string
	AddressLine2            string
	EmailInternalRecipients map[string]string //super admin email map ("email": "name")
	initOnce                sync.Once
)

func validateCompanyDetails() {
	if Phone == "" || Email == "" || AddressLine1 == "" ||
		AddressLine2 == "" || len(EmailInternalRecipients) == 0 {
		log.Fatal("Company details not initialized.")
	}
}

func InitDebug() {
	initOnce.Do(func() {
		Email = os.Getenv("COMPANY_EMAIL")
		Phone = os.Getenv("COMPANY_PHONE")
		AddressLine1 = os.Getenv("COMPANY_ADDRESS_LINE1")
		AddressLine2 = os.Getenv("COMPANY_ADDRESS_LINE2")
		rawEmailRecipients := os.Getenv("EMAIL_INTERNAL_RECIPIENTS")
		err := json.Unmarshal([]byte(rawEmailRecipients), &EmailInternalRecipients)
		if err != nil {
			log.Fatal(err)
		}
		validateCompanyDetails()
		log.Println("Company details initialized for DEBUG environment.")
	})
}

func InitProd(projectID string, ctx context.Context) {
	initOnce.Do(func() {
		Email = gcp.LoadSecretsHelper(projectID, "COMPANY_EMAIL")
		Phone = gcp.LoadSecretsHelper(projectID, "COMPANY_PHONE")
		AddressLine1 = gcp.LoadSecretsHelper(projectID, "COMPANY_ADDRESS_LINE1")
		AddressLine2 = gcp.LoadSecretsHelper(projectID, "COMPANY_ADDRESS_LINE2")
		rawEmailRecipients := gcp.LoadSecretsHelper(projectID, "EMAIL_INTERNAL_RECIPIENTS")
		err := json.Unmarshal([]byte(rawEmailRecipients), &EmailInternalRecipients)
		if err != nil {
			log.Fatal(err)
		}
		validateCompanyDetails()
		log.Println("Company details initialized for PRODUCTION environment.")
	})
}
