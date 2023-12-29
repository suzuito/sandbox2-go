package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTP struct {
	Cli                httpclientwrapper.HTTPClientWrapper
	StatusCodesSuccess []int
}

func (t *FetcherHTTP) ID() crawler.FetcherID {
	return crawler.FetcherID("fetcher_http")
}

func (t *FetcherHTTP) Do(ctx context.Context, logger *slog.Logger, w io.Writer, input crawler.CrawlerInputData) error {
	urlString, exists := input["URL"]
	if !exists {
		return terrors.Wrapf("input[\"URL\"] not found in input")
	}
	method, exists := input["Method"]
	if !exists {
		method = http.MethodGet
	}
	methodAsString := ""
	switch v := method.(type) {
	case string:
		methodAsString = v
	default:
		return terrors.Wrapf("input[\"Method\"] must be string in input")
	}
	u, err := url.Parse(urlString.(string))
	if err != nil {
		return terrors.Wrap(err)
	}
	req, _ := http.NewRequestWithContext(ctx, methodAsString, u.String(), nil)
	return terrors.Wrap(t.Cli.Do(ctx, logger, req, w, t.StatusCodesSuccess))
}
