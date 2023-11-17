package timeseriesdata

import (
	"fmt"
	"time"
)

type TimeSeriesDataConnpassEvent struct {
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
	Organizer   *TimeSeriesDataConnpassEventOrganizer
}

func (t *TimeSeriesDataConnpassEvent) GetID() TimeSeriesDataID {
	return TimeSeriesDataID(fmt.Sprintf("connpass-%d", t.EventID))
}

func (t *TimeSeriesDataConnpassEvent) GetPublishedAt() time.Time {
	return t.StartedAt
}

type TimeSeriesDataConnpassEventOrganizer struct {
	Name     string
	URL      string
	ImageURL string
}
