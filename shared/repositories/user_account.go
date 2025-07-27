package repositories

import (
	"context"

	firebase_shared "github.com/HarshMohanSason/AHSChemicalsGCShared/shared/firebase"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
)

//Only adds brands and customers to firestore
func CreateUserAccountInFirestore(userAccount *models.CreateUserAccount, ctx context.Context) error {
	_, err := firebase_shared.FirestoreClient.Collection("users").Doc(userAccount.Email).Set(ctx, userAccount.MapToFirestore())
	return err
}

func DeleteUserAccountInFirestore(email string, ctx context.Context) error {
	_, err := firebase_shared.FirestoreClient.Collection("users").Doc(email).Delete(ctx)
	return err
}