package fetcherimpl

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"slices"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/httprequestcache"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type FetcherHTTP struct {
	Cli                *http.Client
	HTTPRequestCache   httprequestcache.HTTPRequestCache
	StatusCodesSuccess []int
	UseCache           bool
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
	LogRequest(logger, req)
	res, err := t.Cli.Do(req)
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

func NewFetcherHTTP(def *crawler.FetcherDefinition, args *factory.NewFuncFetcherArgument) (crawler.Fetcher, error) {
	f := FetcherHTTP{}
	if f.ID() != def.ID {
		return nil, factory.ErrNoMatchedFetcherID
	}
	f.Cli = args.HTTPClient
	statusCodesSuccess, err := argument.GetFromArgumentDefinition[[]int](def.Argument, "StatusCodesSuccess")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	useCache, err := argument.GetFromArgumentDefinition[bool](def.Argument, "UseCache")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	f.StatusCodesSuccess = statusCodesSuccess
	f.UseCache = useCache
	f.HTTPRequestCache = args.HTTPRequestCache
	return &f, nil
}
