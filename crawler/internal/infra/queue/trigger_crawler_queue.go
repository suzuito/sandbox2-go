package queue

import (
	"cloud.google.com/go/pubsub"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/queue/internal/queueimpl"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
)

func NewTriggerCrawlerQueue(
	cli *pubsub.Client,
	baseTopicIDForCrawl string,
	topicIDForDispatchCrawl string,
) queue.TriggerCrawlerQueue {
	return &queueimpl.TriggerCrawlerQueue{
		Cli:                     cli,
		BaseTopicIDForCrawl:     baseTopicIDForCrawl,
		TopicIDForDispatchCrawl: topicIDForDispatchCrawl,
	}
}
