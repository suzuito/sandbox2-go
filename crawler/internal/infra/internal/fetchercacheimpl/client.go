package fetchercacheimpl

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type Client struct {
}

func (t *Client) Do(ctx context.Context, req *http.Request) (io.Reader, error) {
	return nil, errors.New("not implemented")
}
