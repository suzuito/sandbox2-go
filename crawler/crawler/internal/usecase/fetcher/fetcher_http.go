package fetcher

import (
	"context"
	"io"
	"net/http"
)

type FetcherHTTP interface {
	DoRequest(ctx context.Context, request *http.Request, w io.Writer) error
}
