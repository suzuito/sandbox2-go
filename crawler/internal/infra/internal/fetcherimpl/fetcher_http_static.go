package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"slices"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/fetchercache"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTPStatic struct {
	Cli                *http.Client
	FetcherCacheClient fetchercache.Client
	Req                *http.Request
	StatusCodesSuccess []int
	UseCache           bool
}

func (t *FetcherHTTPStatic) ID() crawler.FetcherID {
	return "fetcher_http_static"
}

func (t *FetcherHTTPStatic) Do(ctx context.Context, logger *slog.Logger, w io.Writer, _ crawler.CrawlerInputData) error {
	LogRequest(logger, t.Req)
	var r io.Reader
	if t.UseCache {
		var err error
		r, err = t.FetcherCacheClient.Do(ctx, t.Req)
		if err != nil {
			return terrors.Wrap(err)
		}
	} else {
		res, err := t.Cli.Do(t.Req)
		if err != nil {
			return terrors.Wrap(err)
		}
		defer res.Body.Close()
		if !slices.Contains(t.StatusCodesSuccess, res.StatusCode) {
			status := res.StatusCode
			return terrors.Wrapf("HTTP error : status=%d", status)
		}
		r = res.Body
	}
	if _, err := io.Copy(w, r); err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	return nil
}

func NewFetcherHTTPStatic(def *crawler.FetcherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
	f := FetcherHTTPStatic{}
	if f.ID() != def.ID {
		return nil, factory.ErrNoMatchedFetcherID
	}
	urlString, err := argument.GetFromArgumentDefinition[string](def.Argument, "URL")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	method, err := argument.GetFromArgumentDefinition[string](def.Argument, "Method")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	statusCodesSuccess, err := argument.GetFromArgumentDefinition[[]int](def.Argument, "StatusCodesSuccess")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	req, err := http.NewRequest(method, urlString, nil)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	useCache, err := argument.GetFromArgumentDefinition[bool](def.Argument, "UseCache")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	f.Req = req
	f.StatusCodesSuccess = statusCodesSuccess
	f.UseCache = useCache
	f.Cli = setting.FetcherFactorySetting.HTTPClient
	f.FetcherCacheClient = setting.FetcherFactorySetting.FetcherCacheClient
	return &f, nil
}
