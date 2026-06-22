package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/gcp"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func GetFCMTokens(ctx context.Context) ([]string, error) {
	docs, err := FirestoreClient.Collection(UsersCollection).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var tokens []string
	for _, doc := range docs {
		var user models.UserAccount
		if err := doc.DataTo(&user); err != nil {
			gcp.LogWarning("send-notification-fn", "Error converting document to user: "+err.Error())
			continue
		}
		tokens = append(tokens, user.FCMTokens...)
	}
	return tokens, nil
}

func RemoveInvalidToken(ctx context.Context, token string) {

	docs, err := FirestoreClient.Collection(UsersCollection).
		Where("fcmTokens", "array-contains", token).
		Documents(ctx).GetAll()
	if err != nil {
		gcp.LogInfo("send-notification-fn", "Error getting documents during remove Invalid Token: "+err.Error())
		return
	}
	for _, doc := range docs {
		_, err = doc.Ref.Update(ctx, []firestore.Update{
			{Path: "fcmTokens", Value: firestore.ArrayRemove(token)},
		})
		if err != nil {
			gcp.LogInfo("send-notification-fn", "Error updating document during remove Invalid Token: "+err.Error())
		}
	}
}
