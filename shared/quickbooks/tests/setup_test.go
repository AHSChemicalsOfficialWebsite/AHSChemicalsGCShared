package tests

import (
	"log"
	"os"
	"testing"

	firebase_shared "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/quickbooks"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	
	err := godotenv.Load("../../../keys/.env.development")
	if err != nil{
		log.Fatalf("Error loading the .env file: %v", err)
	}	
	firebase_shared.InitFirebaseDebug(os.Getenv("DEBUG_ADMIN_SDK"))
	quickbooks.InitQuickBooksDebug()
	
	code := m.Run()
	os.Exit(code)
}