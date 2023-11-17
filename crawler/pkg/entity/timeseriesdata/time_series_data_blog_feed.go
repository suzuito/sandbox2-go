package timeseriesdata

import (
	"time"
)

type TimeSeriesDataBlogFeed struct {
	ID             TimeSeriesDataID
	URL            string
	PublishedAt    time.Time
	Title          string
	Summary        string
	ArticleContent string
	Thumbnail      *TimeSeriesDataBlogFeedThumbnail
	Author         *TimeSeriesDataBlogFeedAuthor
}

func (t *TimeSeriesDataBlogFeed) GetID() TimeSeriesDataID {
	return t.ID
}

func (t *TimeSeriesDataBlogFeed) GetPublishedAt() time.Time {
	return t.PublishedAt
}

type TimeSeriesDataBlogFeedThumbnail struct {
	ImageURL string
}

type TimeSeriesDataBlogFeedAuthor struct {
	URL      string
	Name     string
	ImageURL string
}
