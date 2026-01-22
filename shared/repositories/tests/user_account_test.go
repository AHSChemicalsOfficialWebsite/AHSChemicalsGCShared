package tests

import (
	"context"
	"testing"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/repositories"
)

func TestUserAccount(t *testing.T) {

	userAccounts, err := repositories.FetchAssignedAdminsForCustomer(context.Background(), "1")
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
