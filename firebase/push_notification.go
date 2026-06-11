package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/messaging"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/gcp"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func getFCMTokens(ctx context.Context) ([]string, error) {
    docs, err := FirestoreClient.Collection(UsersCollection).Documents(ctx).GetAll()
    if err != nil {
        return nil, err
    }
    var tokens []string
    for _, doc := range docs {
        var user models.UserAccount
        if err := doc.DataTo(&user); err != nil {
        	gcp.LogWarning("FCM Push Notification", "Error converting document to user: "+err.Error())
            continue
        }
        tokens = append(tokens, user.FCMTokens...)
    }
    return tokens, nil
}

func SendNotification(ctx context.Context, n *models.Notifcation) {
	tokens, err := getFCMTokens(ctx)
	if err != nil{
		gcp.LogError("FCM Push Notification", "Error getting tokens: "+err.Error())
		return
	}

	badge := 1 // for APNS
	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title:    n.Title,
			Body:     n.Body,
			ImageURL: "https://azurehospitalitysupply.com/web-app-manifest-192x192.png",
		},
		Webpush: &messaging.WebpushConfig{
			Headers: map[string]string{
				"Urgency": "high",
				"TTL":     "86400", //24 hours
			},
			Notification: &messaging.WebpushNotification{
				Title:              n.Title,
				Body:               n.Body,
				Icon:               "https://azurehospitalitysupply.com/web-app-manifest-192x192.png", // notification icon
				Badge:              "https://azurehospitalitysupply.com/favicon-96x96.png",            // small monochrome icon in status bar
				Image:              "https://azurehospitalitysupply.com/web-app-manifest-512x512.png", // large image below notification n.Body
				Vibrate:            []int{200, 100, 200},
				RequireInteraction: true,
				Renotify:           false, // if true, notify even if same tag exists
				Silent:             false,
				Tag:                "inventory-update", // replaces previous notification with same tag
				TimestampMillis:    nil,                // custom timestamp, nil = now
				Actions: []*messaging.WebpushNotificationAction{
					{
						Action: "view",
						Title:  "View Products",
					},
					{
						Action: "dismiss",
						Title:  "Dismiss",
					},
				},
				CustomData: map[string]interface{}{
					"url": "https://azurehospitalitysupply.com/dasboard/products",
				},
			},
			FCMOptions: &messaging.WebpushFCMOptions{
				Link: "https://azurehospitalitysupply.com/dasboard/products",
			},
		},
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				Title:       n.Title,
				Body:        n.Body,
				Icon:        "ic_notification",
				Color:       "#415391",
				Sound:       "default",
				ClickAction: "FLUTTER_NOTIFICATION_CLICK",
				Priority:    messaging.PriorityHigh,
			},
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Alert: &messaging.ApsAlert{
						Title:    n.Title,
						Body:     n.Body,
						SubTitle: "Inventory Alert",
					},
					Badge:            &badge,
					Sound:            "default",
					ContentAvailable: true,
					MutableContent:   true,
				},
			},
			FCMOptions: &messaging.APNSFCMOptions{
				ImageURL: "https://azurehospitalitysupply.com/web-app-manifest-192x192.png",
			},
		},
	}
	resp, err := MessagingClient.SendEachForMulticast(ctx, message)
	if err != nil {
		gcp.LogError("FCM Push Notification", "Error sending notification: "+err.Error())
	}
	for i, result := range resp.Responses {
		if !result.Success {
			if messaging.IsUnregistered(result.Error) {
				removeInvalidToken(ctx, tokens[i])
			} else {
				gcp.LogError("FCM Push Notification", "Error sending notification: "+result.Error.Error())
			}
		}
	}
}

func removeInvalidToken(ctx context.Context, token string) {

	docs, err := FirestoreClient.Collection(UsersCollection).
		Where("fcmTokens", "array-contains", token).
		Documents(ctx).GetAll()
	if err != nil {
		gcp.LogInfo("FCM Push Notification", "Error getting documents during remove Invalid Token: "+err.Error())
	}
	for _, doc := range docs {
		_, err = doc.Ref.Update(ctx, []firestore.Update{
			{Path: "fcmTokens", Value: firestore.ArrayRemove(token)},
		})
		if err != nil {
			gcp.LogInfo("FCM Push Notification", "Error updating document during remove Invalid Token: "+err.Error())
		}
	}
}
