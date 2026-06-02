package tests

import (
	"log"
	"os"
	"testing"

	firebase "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared//firebase"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	
	err := godotenv.Load("../../../keys/.env.development")
	if err != nil{
		log.Fatalf("Error loading the .env file: %v", err)
	}
	firebase.InitFirebaseDebug(os.Getenv("DEBUG_ADMIN_SDK"))
	code := m.Run()

	os.Exit(code)
}