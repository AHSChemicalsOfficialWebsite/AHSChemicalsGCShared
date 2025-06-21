package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared"
)

func TestGetSecretFromGCP(t *testing.T){
	
	secretPath := fmt.Sprintf("projects/%s/secrets/%s/versions/latest",os.Getenv("GCP_PROD_PROJECT_ID"),"TWILIO_RECIPENTS_PHONE")
	
	secret, err := shared.GetSecretFromGCP(secretPath)
	if err != nil{
		t.Fatalf("Error occurred while getting the secret %v", err)
	}
	t.Logf("Secret successfully fetched: %s", secret)
}