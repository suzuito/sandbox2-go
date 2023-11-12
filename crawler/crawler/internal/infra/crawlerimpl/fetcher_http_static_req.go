package crawlerimpl

import (
	"context"
	"io"
	"net/http"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
)

type FetcherHTTPStatic struct {
	httpClient             *http.Client
	httpRequest            *http.Request
	isSuccessedRequestFunc func(res *http.Response) bool
}

func (t *FetcherHTTPStatic) Do(ctx context.Context, w io.Writer, _ crawler.CrawlerInputData) error {
	res, err := t.httpClient.Do(t.httpRequest)
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

func NewFetcherHTTPStatic(
	httpClient *http.Client,
	httpRequest *http.Request,
	isSuccessedRequestFunc func(res *http.Response) bool,
) *FetcherHTTPStatic {
	return &FetcherHTTPStatic{
		httpClient:             httpClient,
		httpRequest:            httpRequest,
		isSuccessedRequestFunc: isSuccessedRequestFunc,
	}
}
