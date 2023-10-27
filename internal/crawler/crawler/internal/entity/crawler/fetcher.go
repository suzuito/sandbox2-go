package crawler

import (
	"context"
	"io"
)

type Fetcher interface {
	Fetch(ctx context.Context, w io.Writer) error
}
