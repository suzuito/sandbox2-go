package timeseriesdata

import (
	"strings"
	"time"
)

type TimeSeriesDataRSS struct {
	GUID        string
	PublishedAt time.Time
	Title       string
	URL         string
}

func (t *TimeSeriesDataRSS) GetID() TimeSeriesDataID {
	return TimeSeriesDataID(strings.Replace(
		strings.Replace(t.GUID, ":", "-", -1),
		"/",
		"-",
		-1,
	))
}

func (t *TimeSeriesDataRSS) GetPublishedAt() time.Time {
	return t.PublishedAt
}
