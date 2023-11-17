package timeseriesdata

import (
	"fmt"
	"time"
)

type TimeSeriesDataEvent struct {
	EventID     int64
	Title       string
	Catch       string
	Description string
	EventURL    string
	StartedAt   time.Time
	EndedAt     time.Time
	Place       string
	Address     string
	Lat         float64
	Lon         float64
	Organizer   *TimeSeriesDataEventOrganizer
}

func (t *TimeSeriesDataEvent) GetID() TimeSeriesDataID {
	return TimeSeriesDataID(fmt.Sprintf("connpass-%d", t.EventID))
}

func (t *TimeSeriesDataEvent) GetPublishedAt() time.Time {
	return t.StartedAt
}

type TimeSeriesDataEventOrganizer struct {
	Name     string
	URL      string
	ImageURL string
}
