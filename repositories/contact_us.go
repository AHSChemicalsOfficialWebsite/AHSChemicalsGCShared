package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SaveContactUsToFirestore(c *models.ContactUsForm, ip string, ctx context.Context) error {

	docSnapshot, err := firebase.FirestoreClient.Collection(firebase.ContactUsCollection).Doc(ip).Get(ctx)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			_, err = firebase.FirestoreClient.Collection(firebase.ContactUsCollection).Doc(ip).Set(ctx, c)
			return err
		}
		return err
	}

	if !docSnapshot.Exists() {
		_, err = firebase.FirestoreClient.Collection(firebase.ContactUsCollection).Doc(ip).Set(ctx, c)
		return err
	}

	var oldContactUs models.ContactUsForm
	err = docSnapshot.DataTo(&oldContactUs)
	if err != nil {
		return err
	}

	if time.Since(oldContactUs.Timestamp) < time.Hour *24 {
		return errors.New("You need to wait 24 hours before submitting another contact us request")
	}

	_, err = firebase.FirestoreClient.Collection(firebase.ContactUsCollection).Doc(ip).Set(ctx, c)
	if err != nil {
		return err
	}
	return nil
}