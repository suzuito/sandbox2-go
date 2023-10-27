package timeseriesdata

import (
	"testing"
	"time"
)

func TestTimeSeriesDataConnpassEvent_GetID(t *testing.T) {
	event := &TimeSeriesDataConnpassEvent{
		EventID: 12345,
		Title:   "Sample Event",
	}

	expectedID := TimeSeriesDataID("connpass-12345")
	if got := event.GetID(); got != expectedID {
		t.Errorf("GetID() = %v; want %v", got, expectedID)
	}
}

func TestTimeSeriesDataConnpassEvent_GetPublishedAt(t *testing.T) {
	date := time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC)
	event := &TimeSeriesDataConnpassEvent{
		EventID:   12345,
		Title:     "Sample Event",
		StartedAt: date,
	}

	if got := event.GetPublishedAt(); !got.Equal(date) {
		t.Errorf("GetPublishedAt() = %v; want %v", got, date)
	}
}
