package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTPStatic struct {
	Cli                httpclientwrapper.HTTPClientWrapper
	Req                *http.Request
	StatusCodesSuccess []int
}

func (t *FetcherHTTPStatic) ID() crawler.FetcherID {
	return "fetcher_http_static"
}

func (t *FetcherHTTPStatic) Do(ctx context.Context, logger *slog.Logger, w io.Writer, _ crawler.CrawlerInputData) error {
	return terrors.Wrap(t.Cli.Do(ctx, logger, t.Req, w, t.StatusCodesSuccess))
}
