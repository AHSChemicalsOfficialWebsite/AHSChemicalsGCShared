package repositories

import (
	"context"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/constants"
	firebase_shared "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/models"
)

func FetchAssignedAdminsForCustomer(ctx context.Context, customerID string) ([]*models.UserAccount, error) {
	docSnapshots, err := firebase_shared.FirestoreClient.Collection(constants.UsersCollection).
		Where("role", "==", constants.RoleAdmin).
		Where("customers", "array-contains", customerID).
		Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	userAccounts := make([]*models.UserAccount, 0)
	for _, docSnapshot := range docSnapshots {
		var userAccount models.UserAccount
		err = docSnapshot.DataTo(&userAccount)
		if err != nil {
			return nil, err
		}
		userAccounts = append(userAccounts, &userAccount)
	}
	return userAccounts, nil
}
