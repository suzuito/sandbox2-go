package crawler

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

type Parser interface {
	Parse(ctx context.Context, r io.Reader) ([]timeseriesdata.TimeSeriesData, error)
}
