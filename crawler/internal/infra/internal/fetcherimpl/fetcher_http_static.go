package fetcherimpl

import (
	"context"
	"io"
	"net/http"
	"slices"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
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

func (t *FetcherHTTPStatic) Do(ctx context.Context, w io.Writer, _ crawler.CrawlerInputData) error {
	res, err := t.Cli.Do(t.Req)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if !slices.Contains(t.StatusCodesSuccess, res.StatusCode) {
		body, _ := io.ReadAll(res.Body)
		status := res.StatusCode
		return terrors.Wrapf("HTTP error : status=%d body=%s", status, body)
	}
	if _, err := io.Copy(w, res.Body); err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	return nil
}

func NewFetcherHTTPStatic(def *crawler.FetcherDefinition, args *factory.NewFuncFetcherArgument) (crawler.Fetcher, error) {
	f := FetcherHTTPStatic{}
	if f.ID() != def.ID {
		return nil, factory.ErrNoMatchedFetcherID
	}
	f.Cli = args.HTTPClient
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
	f.Req = req
	f.StatusCodesSuccess = statusCodesSuccess
	return &f, nil
}
