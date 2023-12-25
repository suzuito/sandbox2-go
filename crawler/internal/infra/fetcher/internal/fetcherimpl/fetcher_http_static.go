package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"slices"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTPStatic struct {
	Cli                *http.Client
	Req                *http.Request
	StatusCodesSuccess []int
}

func (t *FetcherHTTPStatic) ID() crawler.FetcherID {
	return "fetcher_http_static"
}

func (t *FetcherHTTPStatic) Do(ctx context.Context, logger *slog.Logger, w io.Writer, _ crawler.CrawlerInputData) error {
	LogRequest(logger, t.Req)
	res, err := t.Cli.Do(t.Req)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if !slices.Contains(t.StatusCodesSuccess, res.StatusCode) {
		status := res.StatusCode
		return terrors.Wrapf("HTTP error : status=%d", status)
	}
	if _, err := io.Copy(w, res.Body); err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	return nil
}
