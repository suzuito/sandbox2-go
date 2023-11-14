package infra

import (
	"cloud.google.com/go/pubsub"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/gcp"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
)

func NewTriggerCrawlerQueue(
	cli *pubsub.Client,
	topicIDTriggerCrawlerQueue string,
) queue.TriggerCrawlerQueue {
	return &gcp.TriggerCrawlerQueue{
		Cli:                        cli,
		TopicIDTriggerCrawlerQueue: topicIDTriggerCrawlerQueue,
	}
}
