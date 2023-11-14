package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type CrawlerRepository interface {
	GetCrawlerDefinition(
		ctx context.Context,
		id crawler.CrawlerID,
	) (*crawler.CrawlerDefinition, error)
	GetCrawlerDefinitions(
		ctx context.Context,
		crawlerIDs ...crawler.CrawlerID,
	) ([]*crawler.CrawlerDefinition, error)
}
