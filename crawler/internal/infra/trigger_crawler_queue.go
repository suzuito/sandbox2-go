package infra

import (
	"cloud.google.com/go/pubsub"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/queueimpl/triggercrawlerqueue"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
)

func NewTriggerCrawlerQueue(
	cli *pubsub.Client,
	baseTopicIDForCrawl string,
	topicIDForDispatchCrawl string,
) queue.TriggerCrawlerQueue {
	return &triggercrawlerqueue.TriggerCrawlerQueue{
		Cli:                     cli,
		BaseTopicIDForCrawl:     baseTopicIDForCrawl,
		TopicIDForDispatchCrawl: topicIDForDispatchCrawl,
	}
}
