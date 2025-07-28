package tests

import (
	"log"
	"os"
	"testing"

	firebase_shared "github.com/HarshMohanSason/AHSChemicalsGCShared/shared/firebase"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	
	err := godotenv.Load("../../../keys/.env.development")
	if err != nil{
		log.Fatalf("Error loading the .env file: %v", err)
	}
	firebase_shared.InitFirebaseDebug(os.Getenv("DEBUG_ADMIN_SDK"))
	code := m.Run()

	os.Exit(code)
}