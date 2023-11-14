package gcp

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type TriggerCrawlerQueueEvent struct {
	CrawlID          crawler.CrawlerID
	CrawlerInputData crawler.CrawlerInputData
}

type TriggerCrawlerQueue struct {
	Cli                        *pubsub.Client
	TopicIDTriggerCrawlerQueue string
}

func (t *TriggerCrawlerQueue) PublishCrawlEvent(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
	crawlerInputData crawler.CrawlerInputData,
) error {
	crawlEvent := TriggerCrawlerQueueEvent{
		CrawlID:          crawlerID,
		CrawlerInputData: crawlerInputData,
	}
	rawBytes, err := json.Marshal(crawlEvent)
	if err != nil {
		return terrors.Wrap(err)
	}
	msg := pubsub.Message{
		Data:       rawBytes,
		Attributes: map[string]string{},
	}
	switch v := ctx.Value("traceId").(type) {
	case string:
		msg.Attributes["traceId"] = v
	}
	topic := t.Cli.Topic(t.TopicIDTriggerCrawlerQueue).Publish(ctx, &msg)
	_, err = topic.Get(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *TriggerCrawlerQueue) RecieveCrawlEvent(
	ctx context.Context,
	rawBytes []byte,
) (crawler.CrawlerID, crawler.CrawlerInputData, error) {
	crawlEvent := TriggerCrawlerQueueEvent{}
	if err := json.Unmarshal(rawBytes, &crawlEvent); err != nil {
		return "", nil, terrors.Wrap(err)
	}
	return crawlEvent.CrawlID, crawlEvent.CrawlerInputData, nil
}
