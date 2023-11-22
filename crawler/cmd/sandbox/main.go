package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/suzuito/sandbox2-go/common/test_helper"
)

func main() {
	ctx := context.Background()
	cli, err := test_helper.NewPubSubClient(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	topic := cli.Topic("foo1")
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	fmt.Println(topic.ID())

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("hello"),
	})
	sid, err := result.Get(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	fmt.Println(sid)
}
