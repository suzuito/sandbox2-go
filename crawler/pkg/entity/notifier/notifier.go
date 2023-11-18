package notifier

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type NotifierID string

type Notifier interface {
	ID() NotifierID
	NewEmptyTimeSeriesData() timeseriesdata.TimeSeriesData
	Notify(
		ctx context.Context,
		d timeseriesdata.TimeSeriesData,
	) error
}
