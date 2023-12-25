package fetchercache

import (
	"context"
	"io"
	"net/http"
)

type Client interface {
	Do(ctx context.Context, req *http.Request) (io.Reader, error)
}
