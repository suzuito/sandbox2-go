package timeseriesdata

import (
	"time"
)

type TimeSeriesDataBlogFeed struct {
	ID          TimeSeriesDataID
	PublishedAt time.Time
	Title       string
	URL         string
}

func (t *TimeSeriesDataBlogFeed) GetID() TimeSeriesDataID {
	return t.ID
}

func (t *TimeSeriesDataBlogFeed) GetPublishedAt() time.Time {
	return t.PublishedAt
}
