package crawlerimpl

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
)

type FetcherHTTPGet struct {
	httpClient             *http.Client
	isSuccessedRequestFunc func(res *http.Response) bool
}

func (t *FetcherHTTPGet) Do(ctx context.Context, w io.Writer, input crawler.CrawlerInputData) error {
	urlString, exists := input["URL"]
	if !exists {
		return terrors.Wrapf("input[\"URL\"] not found in input")
	}
	u, err := url.Parse(urlString.(string))
	if err != nil {
		return terrors.Wrap(err)
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	res, err := t.httpClient.Do(req)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer res.Body.Close()
	if !t.isSuccessedRequestFunc(res) {
		body, _ := io.ReadAll(res.Body)
		status := res.StatusCode
		return terrors.Wrapf("HTTP error : status=%d body=%s", status, body)
	}
	if _, err := io.Copy(w, res.Body); err != nil {
		return terrors.Wrapf("Failed to io.Copy: %+v", err)
	}
	return nil
}

func NewFetcherHTTPGet(
	httpClient *http.Client,
	isSuccessedRequestFunc func(res *http.Response) bool,
) *FetcherHTTPGet {
	return &FetcherHTTPGet{
		httpClient:             httpClient,
		isSuccessedRequestFunc: isSuccessedRequestFunc,
	}
}
