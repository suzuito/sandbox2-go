package fetcher

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
)

type Fetcher interface {
	Do(ctx context.Context, w io.Writer, input crawler.CrawlerInputData) error
}
