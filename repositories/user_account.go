package repositories

import (
	"context"

	firebase "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func FetchAssignedEmailsForCustomer(ctx context.Context, customerID string, roles ...models.Role) (map[models.Role][]*models.UserAccount, error) {
	userAccounts := make(map[models.Role][]*models.UserAccount)
	for _, role := range roles { //Uncessary to use go routines here since the max number of concurrent calls will be 2 anyways
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
			userAccounts[role] = append(userAccounts[role], &userAccount)
		}
	}
	return userAccounts, nil
}
