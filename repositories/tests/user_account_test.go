package tests

import (
	"context"
	"testing"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/repositories"
)

func TestUserAccount(t *testing.T) {

	userAccounts, err := repositories.FetchAllCustomersFromFirestore(context.Background())
	if err != nil {
		t.Errorf("Error fetching user accounts: %v", err)
		return
	}
	if len(userAccounts) == 0 {
		t.Errorf("No user accounts found")
		return
	}
	for _, acc := range userAccounts {
		t.Logf("User account: %+v", *&acc.Email) // *acc to dereference the pointer
	}
}
