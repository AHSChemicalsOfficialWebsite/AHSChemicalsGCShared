package tests

import (
	"context"
	"testing"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/quickbooks/qbservices"
)

func TestGetTokenFromFirestore(t *testing.T) {

	tokenResp, err := qbservices.GetTokenUIDFromFirestore(context.Background())
	if err != nil{
		t.Error(err)
	}
	t.Log(tokenResp)
}

