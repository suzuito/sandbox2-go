package httpclientcache

import (
	"context"
	"io"
	"net/http"
)

type Client interface {
	Get(
		ctx context.Context,
		req *http.Request,
		w io.Writer,
		opt *ClientOption,
	) (bool, error)
	Set(
		ctx context.Context,
		req *http.Request,
		contentType string,
		responseBody io.Reader,
		opt *ClientOption,
	) error
}
