package timeseriesdata

import "time"

type TimeSeriesDataBaseID string

type TimeSeriesDataID string

type TimeSeriesData interface {
	GetID() TimeSeriesDataID
	GetPublishedAt() time.Time
}
