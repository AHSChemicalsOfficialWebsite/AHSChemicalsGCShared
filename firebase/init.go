package firebase

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

var (
	App             *firebase.App
	AuthClient      *auth.Client
	FirestoreClient *firestore.Client
	StorageClient   *storage.Client
	MessagingClient *messaging.Client
	initOnce        sync.Once
)

func Init(projectID string, keyPath *string) {
	initOnce.Do(func() {
		ctx := context.Background()
		var err error

		config := &firebase.Config{
			StorageBucket: projectID + ".firebasestorage.app",
		}
		
		if keyPath != nil {
			App, err = firebase.NewApp(ctx, config, option.WithCredentialsFile(*keyPath))
		} else {
			App, err = firebase.NewApp(ctx, config)
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
	})
}
