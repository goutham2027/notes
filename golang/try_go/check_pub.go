package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func PushToTopic() {
	ctx := context.Background()
	config := GetPubSubConfig()
	client, err := pubsub.NewClient(ctx, config.GCP_PROJECT)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	t := client.Topic(config.PUB_SUB_TOPIC)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello World"),
	})

	id, err := result.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	// fmt.Println("PubSub config")
	// fmt.Printf("%+v", config)
}

// func main() {
// 	PushToTopic()
// }
