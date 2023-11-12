package parser

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Parser interface {
	Do(ctx context.Context, r io.Reader, input crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error)
}
