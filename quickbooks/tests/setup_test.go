package tests

import (
	"log"
	"os"
	"testing"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	
	err := godotenv.Load("../../../keys/.env.development")
	if err != nil{
		log.Fatalf("Error loading the .env file: %v", err)
	}	
	var debugSDK string = os.Getenv("DEBUG_ADMIN_SDK")
	firebase.Init("", &debugSDK)
	quickbooks.InitDebug()
	
	code := m.Run()
	os.Exit(code)
}