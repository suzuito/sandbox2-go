package gcp

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/suzuito/sandbox2-go/internal/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/entity/crawler"
)

type CrawlEvent struct {
	CrawlID crawler.CrawlerID
}

type Queue struct {
	pcli              *pubsub.Client
	topicIDCrawlEvent string
}

func (t *Queue) PublishCrawlEvent(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
) error {
	crawlEvent := CrawlEvent{
		CrawlID: crawlerID,
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
	topic := t.pcli.Topic(t.topicIDCrawlEvent).Publish(ctx, &msg)
	_, err = topic.Get(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Queue) RecieveCrawlEvent(
	ctx context.Context,
	rawBytes []byte,
) (crawler.CrawlerID, error) {
	crawlEvent := CrawlEvent{}
	if err := json.Unmarshal(rawBytes, &crawlEvent); err != nil {
		return "", terrors.Wrap(err)
	}
	return crawlEvent.CrawlID, nil
}

func NewQueue(pcli *pubsub.Client, topicIDCrawlEvent string) *Queue {
	return &Queue{
		pcli:              pcli,
		topicIDCrawlEvent: topicIDCrawlEvent,
	}
}
