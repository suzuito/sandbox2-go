package infra

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type FunctionTrigger struct {
	Cli     *pubsub.Client
	TopicID string
}

func (t *FunctionTrigger) Publish(
	ctx context.Context,
	msg any,
) error {
	rawBytes, err := json.Marshal(msg)
	if err != nil {
		return terrors.Wrap(err)
	}
	m := pubsub.Message{
		Data:       rawBytes,
		Attributes: map[string]string{},
	}
	switch v := ctx.Value("traceId").(type) {
	case string:
		m.Attributes["traceId"] = v
	}
	topic := t.Cli.Topic(t.TopicID).Publish(ctx, &m)
	_, err = topic.Get(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
