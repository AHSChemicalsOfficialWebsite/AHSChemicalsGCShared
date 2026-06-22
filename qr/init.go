package qr

import (
	"log"
	"os"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/gcp"
)

var (
	API_KEY string
)

func Init(projectID string){

	if os.Getenv("ENV") == "DEBUG"{
		API_KEY = os.Getenv("QR_SCAN_CODE_DEBUG")
	}else{
		API_KEY = gcp.LoadSecretsHelper(projectID, "QR_SCAN_CODE")
	}

	if API_KEY == ""{
		log.Fatal("QR_SCAN_CODE is not set.")
	}
}