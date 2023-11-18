package factory

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type CrawlerFactory struct {
	FetcherFactory   FetcherFactory
	ParserFactory    ParserFactory
	PublisherFactory PublisherFactory
}

func (t *CrawlerFactory) Get(ctx context.Context, def *crawler.CrawlerDefinition) (*crawler.Crawler, error) {
	fetcher, err := t.FetcherFactory.Get(ctx, &def.FetcherDefinition)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	parser, err := t.ParserFactory.Get(ctx, &def.ParserDefinition)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	publisher, err := t.PublisherFactory.Get(ctx, &def.PublisherDefinition)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &crawler.Crawler{
		ID:        def.ID,
		Fetcher:   fetcher,
		Parser:    parser,
		Publisher: publisher,
	}, nil
}
