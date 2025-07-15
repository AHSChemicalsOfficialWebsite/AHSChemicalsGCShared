package company_details

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/gcp"
)

// Basic Company information
const (
	COMPANYNAME       = "Azure Hospitality Supply"
	COMPANYWEBSITEURL = "https://azurehospitalitysupply.com"
)

// Company details
var (
	COMPANYEMAIL           string
	COMPANYPHONE           string
	COMPANYADDRESSLINE1    string
	COMPANYADDRESSLINE2    string
	EMAILINTERNALRECIPENTS map[string]string
	LOGOPATH               string
)

func InitCompanyDetailsDebug() {
	COMPANYPHONE = os.Getenv("COMPANYPHONE")
	COMPANYEMAIL = os.Getenv("COMPANYEMAIL")
	COMPANYADDRESSLINE1 = os.Getenv("COMPANYADDRESSLINE1")
	COMPANYADDRESSLINE2 = os.Getenv("COMPANYADDRESSLINE2")
	rawEmailRecipients := os.Getenv("EMAILINTERNALRECIPIENTS")
	err := json.Unmarshal([]byte(rawEmailRecipients), &EMAILINTERNALRECIPENTS)
	if err != nil {
		log.Fatal(err)
	}
	LOGOPATH = os.Getenv("LOGOPATH")
	if COMPANYPHONE == "" || COMPANYEMAIL == "" || COMPANYADDRESSLINE1 == "" || COMPANYADDRESSLINE2 == "" || LOGOPATH == "" {
		log.Fatal("Company details not initialized. Please check environment variables.")
	}
	log.Println("Company details initialized for DEBUG environment.")
}

func InitCompanyDetailsProd(ctx context.Context) {
	shared.InitCompanyDetails.Do(func() {
		projectID, err := metadata.ProjectIDWithContext(ctx)
		if err != nil {
			log.Fatalf("Error loading Google Cloud project ID: %v", err)
		}

		COMPANYPHONE = gcp.LoadSecretsHelper(projectID, "COMPANYPHONE")
		COMPANYEMAIL = gcp.LoadSecretsHelper(projectID, "COMPANYEMAIL")
		COMPANYADDRESSLINE1 = gcp.LoadSecretsHelper(projectID, "COMPANYADDRESSLINE1")
		COMPANYADDRESSLINE2 = gcp.LoadSecretsHelper(projectID, "COMPANYADDRESSLINE2")
		rawEmailRecipients := gcp.LoadSecretsHelper(projectID, "EMAILINTERNALRECIPIENTS")
		err = json.Unmarshal([]byte(rawEmailRecipients), &EMAILINTERNALRECIPENTS)
		if err != nil {
			log.Fatal(err)
		}
		LOGOPATH = gcp.LoadSecretsHelper(projectID, "LOGOPATH")
		if COMPANYPHONE == "" || COMPANYEMAIL == "" || COMPANYADDRESSLINE1 == "" || COMPANYADDRESSLINE2 == "" || LOGOPATH == "" {
			log.Fatal("Company details not initialized. Please check environment variables.")
		}
		log.Println("Company details initialized for PRODUCTION environment.")
	})
}
