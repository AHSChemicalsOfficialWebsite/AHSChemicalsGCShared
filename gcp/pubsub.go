package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
)

var (
	PubSubClient           *pubsub.Client
	initPubSubOnce        	sync.Once
)

func InitPubSub(projectID string, ctx context.Context) {
	initPubSubOnce.Do(func() {
		client, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			log.Fatalf("Failed to create PubSub client: %v", err)
		}
		PubSubClient = client
		log.Println("PubSub initialized")
	})
}

type SubMessage struct {
	Message struct {
		Data        []byte            `json:"data"`        
		MessageID   string            `json:"messageId"`
		PublishTime string            `json:"publishTime"`
		Attributes  map[string]string `json:"attributes"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// DecodeSubMessageData decodes the message data as a T struct. Returns a pointer of
// type T struct. Using T for future usecases.
func DecodeSubMessageData[T any](msg *SubMessage) (*T, error) {
	var result T
	if err := json.Unmarshal(msg.Message.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func PublishMessage(ctx context.Context, topicID string, data []byte) error {
	if PubSubClient == nil {
		return fmt.Errorf("PubSub client not initialized")
	}
	topic := PubSubClient.Topic(topicID)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	_, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	return nil
}
