package repositories

import (
	"context"

	firebase "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func FetchAssignedEmailsForCustomer(ctx context.Context, customerID string, roles ...firebase.Role) ([]*models.UserAccount, error) {
	userAccounts := make([]*models.UserAccount, 0)
	for _, role := range roles {
		docSnapshots, err := firebase.FirestoreClient.Collection(firebase.UsersCollection).
			Where("role", "==", role).
			Where("customers", "array-contains", customerID).
			Documents(ctx).GetAll()
		if err != nil {
			return nil, err
		}
		for _, docSnapshot := range docSnapshots {
			var userAccount models.UserAccount
			if err = docSnapshot.DataTo(&userAccount); err != nil {
				return nil, err
			}
			userAccounts = append(userAccounts, &userAccount)
		}
	}
	return userAccounts, nil
}
