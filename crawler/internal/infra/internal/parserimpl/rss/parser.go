package rss

import (
	"context"
	"io"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Parser struct {
	FP *gofeed.Parser
}

func (t *Parser) ID() crawler.ParserID {
	return "rss"
}

func (t *Parser) Do(ctx context.Context, r io.Reader, _ crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	feed, err := t.FP.Parse(r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	returned := []timeseriesdata.TimeSeriesData{}
	for _, item := range feed.Items {
		publishedAt, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", item.Published)
		if err != nil {
			clog.L.Errorf(ctx, "%+v\n", err)
			continue
		}
		data := timeseriesdata.TimeSeriesDataRSS{
			GUID:        item.GUID,
			PublishedAt: publishedAt,
			Title:       item.Title,
			URL:         item.Link,
		}
		returned = append(returned, &data)
	}
	return returned, nil
}

func New(def *crawler.ParserDefinition, _ *factory.NewFuncParserArgument) (crawler.Parser, error) {
	parser := Parser{
		FP: gofeed.NewParser(),
	}
	if def.ID != parser.ID() {
		return nil, factory.ErrNoMatchedParserID
	}
	return &parser, nil
}
