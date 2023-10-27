package timeseriesdata

import (
	"testing"
	"time"
)

func TestGetID(t *testing.T) {
	id := TimeSeriesDataID("123")
	feed := TimeSeriesDataBlogFeed{
		ID: id,
	}

	if got := feed.GetID(); got != id {
		t.Errorf("GetID() = %v, want %v", got, id)
	}
}

func TestGetPublishedAt(t *testing.T) {
	pubDate := time.Now()
	feed := TimeSeriesDataBlogFeed{
		PublishedAt: pubDate,
	}

	if got := feed.GetPublishedAt(); !got.Equal(pubDate) {
		t.Errorf("GetPublishedAt() = %v, want %v", got, pubDate)
	}
}
