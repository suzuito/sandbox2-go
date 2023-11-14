package gcp

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type CrawlerRepository struct {
	Crawlers map[crawler.CrawlerID]*crawler.CrawlerDefinition
}

func (t *CrawlerRepository) GetCrawlerDefinition(
	ctx context.Context,
	id crawler.CrawlerID,
) (*crawler.CrawlerDefinition, error) {
	crawlerDefinition, exists := t.Crawlers[id]
	if !exists {
		return nil, terrors.Wrapf("CrawlerDefinition[%s] is not found", id)
	}
	return crawlerDefinition, nil
}

func (t *CrawlerRepository) GetCrawlerDefinitions(
	ctx context.Context,
	crawlerIDs ...crawler.CrawlerID,
) ([]*crawler.CrawlerDefinition, error) {
	defs := []*crawler.CrawlerDefinition{}
	for _, id := range crawlerIDs {
		crawlerDefinition, exists := t.Crawlers[id]
		if !exists {
			return nil, terrors.Wrapf("CrawlerDefinition[%s] is not found", id)
		}
		defs = append(defs, crawlerDefinition)
	}
	return defs, nil
}
