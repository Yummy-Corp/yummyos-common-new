package pubsub

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func PublishProtoMessages(topicID string, payload protoreflect.ProtoMessage) error {
	projectID := os.Getenv("GCP_PROJECT")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("pubsub.NewClient: %v", err)
		return nil
	}

	// Get the topic encoding type.
	t := client.Topic(topicID)
	cfg, err := t.Config(ctx)
	if err != nil {
		log.Printf("topic.Config err: %v", err)
		return nil
	}
	encoding := cfg.SchemaSettings.Encoding

	var msg []byte
	switch encoding {
	case pubsub.EncodingBinary:
		msg, err = proto.Marshal(payload)
		if err != nil {
			log.Printf("proto.Marshal err: %v", err)
			return nil
		}
	case pubsub.EncodingJSON:
		pjson := protojson.MarshalOptions{
			EmitUnpopulated: true,
		}
		msg, err = pjson.Marshal(payload)
		if err != nil {
			log.Printf("protojson.Marshal err: %v", err)
			return nil
		}
	default:
		log.Printf("invalid encoding: %v", encoding)
		return nil
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: msg,
	})
	_, err = result.Get(ctx)
	if err != nil {
		log.Printf("result.Get: %v", err)
		return nil
	}
	return nil
}

func SubscribeProtoMessages(subID string, payload protoreflect.ProtoMessage) error {
	projectID := os.Getenv("GCP_PROJECT")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	sub := client.Subscription(subID)
	ctx2, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var mu sync.Mutex
	sub.Receive(ctx2, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		encoding := msg.Attributes["googclient_schemaencoding"]

		if encoding == "BINARY" {
			if err := proto.Unmarshal(msg.Data, payload); err != nil {
				log.Printf("proto.Unmarshal err: %v", err)
				return
			}
		} else if encoding == "JSON" {
			if err := protojson.Unmarshal(msg.Data, payload); err != nil {
				log.Printf("proto.Unmarshal err: %v", err)
				return
			}
		} else {
			log.Printf("invalid encoding: %s", encoding)
		}
		msg.Ack()
	})
	return nil
}
