package timeseriesdata

import "time"

type TimeSeriesDataID string

type TimeSeriesData interface {
	GetID() TimeSeriesDataID
	GetPublishedAt() time.Time
}
