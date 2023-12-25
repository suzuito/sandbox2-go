package queueimpl

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type DispatchCrawlEvent struct {
	CrawlID          crawler.CrawlerID
	CrawlerInputData crawler.CrawlerInputData
}

type TriggerCrawlerQueue struct {
	Cli                     *pubsub.Client
	BaseTopicIDForCrawl     string
	TopicIDForDispatchCrawl string
}

func (t *TriggerCrawlerQueue) PublishDispatchCrawlEvent(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
	crawlerInputData crawler.CrawlerInputData,
) error {
	crawlEvent := DispatchCrawlEvent{
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
	topic := t.Cli.Topic(t.TopicIDForDispatchCrawl).Publish(ctx, &msg)
	_, err = topic.Get(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *TriggerCrawlerQueue) PublishCrawlEvent(
	ctx context.Context,
	crawlerID crawler.CrawlerID,
	crawlerInputData crawler.CrawlerInputData,
	crawlFunctionID crawler.CrawlFunctionID,
) error {
	crawlEvent := DispatchCrawlEvent{
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
	topicID := fmt.Sprintf("%s-%s", t.BaseTopicIDForCrawl, crawlFunctionID)
	topic := t.Cli.Topic(topicID).Publish(ctx, &msg)
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
	crawlEvent := DispatchCrawlEvent{}
	if err := json.Unmarshal(rawBytes, &crawlEvent); err != nil {
		return "", nil, terrors.Wrap(err)
	}
	return crawlEvent.CrawlID, crawlEvent.CrawlerInputData, nil
}
