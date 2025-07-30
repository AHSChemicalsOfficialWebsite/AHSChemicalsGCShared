package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase_shared "github.com/HarshMohanSason/AHSChemicalsGCShared/shared/firebase"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
)

//Only adds brands and customers to firestore
func CreateUserAccountInFirestore(userAccount *models.UserAccountUpdate, ctx context.Context) error {
	_, err := firebase_shared.FirestoreClient.Collection("users").Doc(userAccount.UID).Set(ctx, userAccount.ToMap())
	return err
}

func UpdateUserAccountInFirestore(userAccount *models.UserAccountUpdate, ctx context.Context) error {
	_, err := firebase_shared.FirestoreClient.Collection("users").Doc(userAccount.UID).Set(ctx, userAccount.ToMap(), firestore.MergeAll)
	return err
}

func DeleteUserAccountInFirestore(id string, ctx context.Context) error {
	_, err := firebase_shared.FirestoreClient.Collection("users").Doc(id).Delete(ctx)
	return err
}