package firebase

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"firebase.google.com/go/v4/storage"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/gcp"
	"google.golang.org/api/option"
)

var (
	App             *firebase.App
	AuthClient      *auth.Client
	StorageClient   *storage.Client
	FirestoreClient *firestore.Client
	MessagingClient *messaging.Client
	StorageBucket   string
	initOnce        sync.Once
)

func Init(projectID string, keyPath *string) {
	initOnce.Do(func() {
		ctx := context.Background()
		var err error

		if keyPath != nil {
			opt := option.WithCredentialsFile(*keyPath)
			App, err = firebase.NewApp(ctx, nil, opt)
		} else {
			App, err = firebase.NewApp(ctx, nil)
		}
		if err != nil {
			log.Fatalf("Failed to initialize App: %v", err)
		}

		AuthClient, err = App.Auth(ctx)
		if err != nil {
			log.Fatalf("Failed to initialize Auth client: %v", err)
		}

		FirestoreClient, err = App.Firestore(ctx)
		if err != nil {
			log.Fatalf("Failed to initialize Firestore client: %v", err)
		}

		StorageClient, err = App.Storage(ctx)
		if err != nil {
			log.Fatalf("Failed to initialize Storage client: %v", err)
		}

		MessagingClient, err = App.Messaging(ctx)
		if err != nil {
			log.Fatalf("Failed to initialize Messaging client: %v", err)
		}
		
		if os.Getenv("ENV") == "DEBUG" {
			StorageBucket = os.Getenv("STORAGE_BUCKET")
		} else {
			StorageBucket = gcp.LoadSecretsHelper(projectID, "STORAGE_BUCKET")
		}
		if StorageBucket == "" {
			log.Fatalf("STORAGE_BUCKET environment variable not set")
		}
	})
}
