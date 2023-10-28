package crawler

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Publisher interface {
	Publish(ctx context.Context, data ...timeseriesdata.TimeSeriesData) error
}
