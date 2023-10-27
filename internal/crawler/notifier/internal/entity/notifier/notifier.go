package notifier

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

type NotifierID string

type Notifier interface {
	ID() NotifierID
	NewEmptyTimeSeriesData(
		ctx context.Context,
	) timeseriesdata.TimeSeriesData
	Notify(
		ctx context.Context,
		d timeseriesdata.TimeSeriesData,
	) error
}
